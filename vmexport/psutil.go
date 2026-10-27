package vmexport

import (
	"context"
	"io"
	"net"
	"strconv"
	"sync"
	"time"

	"github.com/VictoriaMetrics/metrics"
	pscpu "github.com/shirou/gopsutil/v4/cpu"
	psdisk "github.com/shirou/gopsutil/v4/disk"
	psmem "github.com/shirou/gopsutil/v4/mem"
	psnet "github.com/shirou/gopsutil/v4/net"
)

func NewPsutil() *Psutil {
	return &Psutil{}
}

type Psutil struct {
	mutex  sync.Mutex
	lastAt time.Time
	names  map[string]bool
}

func (p *Psutil) WriteMetric(ctx context.Context, w io.Writer) {
	// 内存
	if stat, _ := psmem.VirtualMemoryWithContext(ctx); stat != nil {
		metrics.WriteCounterUint64(w, "psutil_memory_bytes_total", stat.Total)
		metrics.WriteCounterUint64(w, "psutil_memory_bytes_used", stat.Used)
	}

	// CPU
	if stats, _ := pscpu.PercentWithContext(ctx, 0, false); len(stats) != 0 {
		metrics.WriteGaugeFloat64(w, "psutil_cpu_usage_percent", stats[0])
	}

	// 磁盘 I/O
	if stats, _ := psdisk.IOCountersWithContext(ctx); len(stats) != 0 {
		for device, stat := range stats {
			name := strconv.Quote(device)
			metrics.WriteCounterUint64(w, "psutil_disk_read_bytes_total{device="+name+"}", stat.ReadBytes)
			metrics.WriteCounterUint64(w, "psutil_disk_written_bytes_total{device="+name+"}", stat.WriteBytes)
		}
	}

	// 磁盘空间
	{
		parts, _ := psdisk.PartitionsWithContext(ctx, true)
		for _, part := range parts {
			device := strconv.Quote(part.Device)
			fstype := strconv.Quote(part.Fstype)
			if stat, _ := psdisk.UsageWithContext(ctx, part.Mountpoint); stat != nil {
				metrics.WriteGaugeUint64(w, "psutil_disk_bytes_total{device="+device+",fstype="+fstype+"}", stat.Total)
				metrics.WriteGaugeUint64(w, "psutil_disk_bytes_free{device="+device+",fstype="+fstype+"}", stat.Free)
				metrics.WriteGaugeUint64(w, "psutil_disk_bytes_used{device="+device+",fstype="+fstype+"}", stat.Used)
			}
		}
	}

	// 网卡 I/O
	{
		interfaceNames := p.interfaceNames()
		if stats, _ := psnet.IOCountersWithContext(ctx, true); stats != nil {
			for _, stat := range stats {
				name := stat.Name
				if _, exists := interfaceNames[name]; !exists {
					continue
				}

				// 网卡名字转义
				name = strconv.Quote(name)
				metrics.WriteCounterUint64(w, "psutil_network_receive_bytes_total{interface="+name+"}", stat.BytesRecv)
				metrics.WriteCounterUint64(w, "psutil_network_transmit_bytes_total{interface="+name+"}", stat.BytesSent)
			}
		}
	}
}

// interfaceNames 获取网卡名。
//
// 通过缓存避免每次都要查询，过期时间保证网卡名不会永久缓存，防止网卡热插拔/运行中启停。
func (p *Psutil) interfaceNames() map[string]bool {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	now := time.Now()
	if now.Sub(p.lastAt) <= 10*time.Minute {
		return p.names
	}

	names := make(map[string]bool, 4)
	faces, _ := net.Interfaces()
	for _, face := range faces {
		if face.Flags&net.FlagLoopback == 0 &&
			face.Flags&net.FlagUp != 0 &&
			face.Flags&net.FlagRunning != 0 { // 忽略环回网卡/不工作的网卡
			names[face.Name] = true
		}
	}

	p.lastAt = now
	p.names = names

	return names
}

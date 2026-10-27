package vmexport

import (
	"context"
	"io"

	"github.com/VictoriaMetrics/metrics"
)

type MetricWriter interface {
	WriteMetric(ctx context.Context, w io.Writer)
}

// VictoriaPushConfigLoader 加载 VictoriaMetrics 配置文件。
type VictoriaPushConfigLoader interface {
	LoadVictoriaPushConfig(ctx context.Context) (pushURL string, opts *metrics.PushOptions, err error)
}

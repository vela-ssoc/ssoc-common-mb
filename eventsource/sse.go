package eventsource

import (
	"compress/gzip"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"sync"
)

// Accept 将请求转为 sse，如果不符合协议就返回 nil。
func Accept(w http.ResponseWriter, r *http.Request) *Writer {
	f, ok := w.(http.Flusher)
	if !ok {
		return nil
	}
	if !strings.Contains(r.Header.Get("Accept"), "text/event-stream") {
		return nil
	}
	if !strings.Contains(strings.ToLower(r.Header.Get("Cache-Control")), "no-cache") {
		return nil
	}

	parent := r.Context()
	ctx, cancel := context.WithCancel(parent)
	sse := &Writer{ctx: ctx, cancel: cancel}

	var compress bool
	encodings := r.Header.Get("Accept-Encoding")
	for str := range strings.SplitSeq(encodings, ",") {
		if strings.TrimSpace(str) == "gzip" {
			w.Header().Set("Content-Encoding", "gzip")
			gzw := gzip.NewWriter(w)
			sse.write = gzw
			sse.close = gzw
			sse.flush = &gzipFlusher{write: gzw, flush: f}
			compress = true
			break
		}
	}
	if !compress {
		sse.write = w
		sse.flush = f
	}

	w.Header().Set("Content-Type", "text/event-stream; charset=utf-8")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("X-Accel-Buffering", "no") // https://nginx.org/en/docs/http/ngx_http_proxy_module.html
	w.WriteHeader(http.StatusOK)

	return sse
}

type Writer struct {
	mutex  sync.Mutex
	write  io.Writer
	flush  http.Flusher
	close  io.Closer
	ctx    context.Context
	cancel context.CancelFunc
}

func (sw *Writer) Write(p []byte) (int, error) {
	sw.mutex.Lock()
	defer sw.mutex.Unlock()

	n := len(p)
	if err := sw.writeData(p); err != nil {
		return 0, err
	}

	return n, nil
}

func (sw *Writer) Close() error {
	var err error
	if c := sw.close; c != nil {
		err = c.Close()
	}
	sw.cancel()

	return err
}

func (sw *Writer) JSON(event string, data any) error {
	msg, err := json.Marshal(data)
	if err != nil {
		return err
	}

	sw.mutex.Lock()
	defer sw.mutex.Unlock()

	if event != "" {
		if err = sw.writeEvent([]byte(event)); err != nil {
			return err
		}
	}
	err = sw.writeData(msg)

	return err
}

func (sw *Writer) Done() <-chan struct{} {
	return sw.ctx.Done()
}

func (sw *Writer) Ping() error {
	sw.mutex.Lock()
	_, err := sw.write.Write([]byte(": ping\n"))
	sw.mutex.Unlock()

	sw.flush.Flush()

	return err
}

func (sw *Writer) writeEvent(evt []byte) error {
	_, err := sw.write.Write([]byte("event: "))
	if err == nil {
		if _, err = sw.write.Write(evt); err == nil {
			_, err = sw.write.Write([]byte{'\n'})
		}
	}
	if err != nil {
		_ = sw.Close()
	}

	return err
}

func (sw *Writer) writeData(data []byte) error {
	_, err := sw.write.Write([]byte("data: "))
	if err == nil {
		if _, err = sw.write.Write(data); err == nil {
			_, err = sw.write.Write([]byte{'\n', '\n'})
		}
	}
	if err != nil {
		_ = sw.Close()
		return err
	}
	sw.flush.Flush()

	return nil
}

type gzipFlusher struct {
	write *gzip.Writer
	flush http.Flusher
}

func (gf *gzipFlusher) Flush() {
	_ = gf.write.Flush()
	gf.flush.Flush()
}

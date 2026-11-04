package victoria

import (
	"io"
	"log/slog"
	"net/http"

	"github.com/VictoriaMetrics/metrics"
)

type PushProxy struct {
	cfg PushConfigurer
	log *slog.Logger
}

func (pp *PushProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	pushURL, opts, err := pp.cfg.PushConfig(ctx)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		pp.log.Error("加载 VictoriaMetrics 配置出错", "err", err)
		return
	}

	bw := pp.newBodyWriter(r)
	if err = metrics.PushMetricsExt(ctx, pushURL, bw.writeTo, opts); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		pp.log.Error("推送 VictoriaMetrics 出错", "err", err)
		return
	}
	if _, err = bw.result(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		pp.log.Error("复制 VictoriaMetrics 数据流出错", "err", err)
		return
	}
}

func (pp *PushProxy) newBodyWriter(r *http.Request) *bodyWriter {
	return &bodyWriter{r: r}
}

type bodyWriter struct {
	r *http.Request
	n int64
	e error
}

func (bw *bodyWriter) writeTo(w io.Writer) {
	bw.n, bw.e = io.Copy(w, bw.r.Body)
}

func (bw *bodyWriter) result() (int64, error) {
	return bw.n, bw.e
}

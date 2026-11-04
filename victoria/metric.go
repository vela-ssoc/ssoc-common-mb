package victoria

import (
	"context"
	"io"

	"github.com/VictoriaMetrics/metrics"
)

type MetricWriter interface {
	WriteMetric(ctx context.Context, w io.Writer)
}

type PushConfigurer interface {
	PushConfig(ctx context.Context) (pushURL string, opts *metrics.PushOptions, err error)
}

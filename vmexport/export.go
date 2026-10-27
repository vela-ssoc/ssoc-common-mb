package vmexport

import (
	"context"
	"io"
)

type MetricWriter interface {
	WriteMetric(ctx context.Context, w io.Writer)
}

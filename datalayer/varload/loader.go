package varload

import (
	"context"
	"os"
)

type Loader interface {
	Load(ctx context.Context, v any) error
}

type JSONC string

func (j JSONC) Load(_ context.Context, v any) error {
	raw, err := os.ReadFile(string(j))
	if err != nil {
		return err
	}

	return unmarshalJSONC(raw, v)
}

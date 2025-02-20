package shipx

import "github.com/xgfone/ship/v5"

type RouteBinder interface {
	BindRoute(r *ship.RouteGroupBuilder) error
}

func BindRouters(r *ship.RouteGroupBuilder, rbs []RouteBinder) error {
	for _, rb := range rbs {
		if err := rb.BindRoute(r); err != nil {
			return err
		}
	}

	return nil
}

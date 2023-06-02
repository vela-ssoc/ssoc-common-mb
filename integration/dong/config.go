package dong

import "github.com/vela-ssoc/vela-common-mb/dal/model"

type Configurer interface {
	Emc() (*model.Emc, error)
	Reset()
}

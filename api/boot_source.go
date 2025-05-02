package api

import (
	"github.com/elegantwalk/gomaasclient/entity"
)

// BootSource is an interface defining API behaviour for
// boot sources
type BootSource interface {
	Get(id int) (*entity.BootSource, error)
	Update(id int, params *entity.BootSourceParams) (*entity.BootSource, error)
	Delete(id int) error
}

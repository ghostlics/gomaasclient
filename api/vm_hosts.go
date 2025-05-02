package api

import (
	"github.com/elegantwalk/gomaasclient/entity"
)

// VMHosts is an interface for listing and creating VMHost objects
type VMHosts interface {
	Get() ([]entity.VMHost, error)
	Create(params *entity.VMHostParams) (*entity.VMHost, error)
}

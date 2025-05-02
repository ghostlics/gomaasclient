package api

import (
	"github.com/elegantwalk/gomaasclient/entity"
	"github.com/elegantwalk/gomaasclient/entity/subnet"
)

// Subnet represents the MAAS Subnet endpoint
type Subnet interface {
	Delete(id int) error
	Get(id int) (*entity.Subnet, error)
	GetIPAddresses(id int) ([]subnet.IPAddress, error)
	GetReservedIPRanges(id int) ([]subnet.ReservedIPRange, error)
	GetStatistics(id int) (*subnet.Statistics, error)
	GetUnreservedIPRanges(id int) ([]subnet.IPRange, error)
	Update(id int, params *entity.SubnetParams) (*entity.Subnet, error)
}

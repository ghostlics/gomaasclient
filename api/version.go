package api

import (
	"github.com/elegantwalk/gomaasclient/entity"
)

// Version is an interface for listing MAAS version details
type Version interface {
	Get() (*entity.Version, error)
}

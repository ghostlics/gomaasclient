package api

import (
	"github.com/elegantwalk/gomaasclient/entity"
)

// SSLKey is an interface defining API behaviour for SSLKey objects
type SSLKey interface {
	Get(id int) (*entity.SSLKey, error)
	Delete(id int) error
}

package api

import (
	"github.com/elegantwalk/gomaasclient/entity"
)

// Tags is an interface for listing and creating Tag objects
type Tags interface {
	Get() ([]entity.Tag, error)
	Create(tagParams *entity.TagParams) (*entity.Tag, error)
}

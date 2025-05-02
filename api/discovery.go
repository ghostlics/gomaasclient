package api

import "github.com/elegantwalk/gomaasclient/entity"

type Discovery interface {
	Get(id string) (*entity.Discovery, error)
}

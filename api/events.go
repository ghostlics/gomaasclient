package api

import "github.com/elegantwalk/gomaasclient/entity"

// Events is an interface for node events
type Events interface {
	Get(params *entity.EventParams) (*entity.EventsResp, error)
}

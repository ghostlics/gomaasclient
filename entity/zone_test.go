package entity

import (
	"testing"

	"github.com/elegantwalk/gomaasclient/test/helper"
)

func TestZonet(t *testing.T) {
	zone := new(Zone)
	if err := helper.TestdataFromJSON("maas/zone.json", zone); err != nil {
		t.Fatal(err)
	}
}

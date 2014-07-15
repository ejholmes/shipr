package api

import (
	"testing"

	"github.com/remind101/shipr"
)

func testShipr(t *testing.T) *shipr.Shipr {
	c, err := shipr.New(&shipr.Options{Env: "test", DBDir: "../../db"})
	if err != nil {
		t.Fatal(err)
	}
	return c
}

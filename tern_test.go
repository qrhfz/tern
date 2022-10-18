package tern_test

import (
	"testing"

	"github.com/qrhfz/tern"
)

func TestBasic(t *testing.T) {
	x := tern.E(5 > 3, "true", "false")

	if x != "true" {
		t.Errorf("should return \"true\"")
	}
}

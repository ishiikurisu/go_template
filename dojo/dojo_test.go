package dojo

import (
	"testing"
)

func TestIsTrue(t *testing.T) {
	var := false
	if !var {
		t.Fatal("false is not true")
		return
	}
}


package clmconv_test

import (
	"testing"

	"github.com/takuoki/clmconv"
)

func TestDummy(t *testing.T) {
	if clmconv.Dummy(0) != 2 {
		t.Error("something wrong")
	}
}

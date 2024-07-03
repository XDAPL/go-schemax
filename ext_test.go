package schemax

import (
	//"fmt"
	"testing"
)

func TestExtensions_codecov(t *testing.T) {
	x := NewExtensions()
	_ = x.String()
	x.Push(nil)
	x.canPush()
	x.canPush(rune(3), 3.14, -1, map[string]string{})
	x.Set(`X-ORIGIN`, `RFCNNNN`)
	x.Set(`X-ORIGIN`, `RFCNNNX`)
	_ = x.String()
	x.Exists(`X-oRiGiN`)
	x.Get(`X-oriGIN`)

	lx := &extension{
		hindent: true,
	}
	X := Extension{lx}
	_ = X.String()
}

package schemax

import (
	"testing"
)

func TestMacros_codecov(t *testing.T) {
	var m Macros
	m.Resolve(`1.3.6.1.4.1.56521`)
	m.ReverseResolve(`n`)
	m = newMacros()
	m.Set(`1.3.6.1.4.1.56521`, `jesse`)
	m.ReverseResolve(`jesse`)
	m.Keys()
}

package schemax

import (
	"testing"
)

func TestDescription(t *testing.T) {
	var desc string = `this is descriptive text`
	if !isValidDescription(desc) {
		t.Errorf("%s failed: valid text caused an error", t.Name())
		return
	}
}

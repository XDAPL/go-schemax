package schemax

import (
	"testing"
)

func TestBoolean_Set(t *testing.T) {
	b := Boolean(0)

	b.set(Obsolete)
	if x := b.Obsolete(); x != `OBSOLETE` {
		t.Errorf("TestBoolean_Set failed: unable to set Obsolete")
	}
	b.unset(Obsolete)
	if !b.IsZero() {
		t.Errorf("TestBoolean_Set failed: unable to ")
	}

	b.set(SingleValue)
	if x := b.SingleValue(); x != `SINGLE-VALUE` {
		t.Errorf("TestBoolean_Set failed: unable to set SingleValue")
	}
	b.unset(SingleValue)

	b.set(Collective)
	if x := b.Collective(); x != `COLLECTIVE` {
		t.Errorf("TestBoolean_Set failed: unable to set Collective")
	}
	b.unset(Collective)

	b.set(NoUserModification)
	if x := b.NoUserModification(); x != `NO-USER-MODIFICATION` {
		t.Errorf("TestBoolean_Set failed: unable to set NoUserModification")
	}
	b.unset(NoUserModification)

	b.set(Collective)
	b.set(SingleValue)
	if x, y := b.is(SingleValue), b.is(Collective); !x || !y {
		t.Errorf("TestBoolean_Set failed: unable to set compound boolean values")
	}

}

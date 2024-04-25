package schemax

import (
	"testing"
)

var tsch Schema

func TestLoadSyntaxes(t *testing.T) {
	want := 67
	if got := tsch.LDAPSyntaxes().Len(); got != want {
		t.Errorf("%s failed: want '%d' ldapSyntaxes, got '%d'",
			t.Name(), want, got)
		return
	}
}

func TestLoadMatchingRules(t *testing.T) {
	want := 44
	if got := tsch.MatchingRules().Len(); got != want {
		t.Errorf("%s failed: want '%d' matchingRules, got '%d'",
			t.Name(), want, got)
		return
	}
}

func TestLoadAttributeTypes(t *testing.T) {
	want := 158
	if got := tsch.AttributeTypes().Len(); got != want {
		t.Errorf("%s failed: want '%d' attributeTypes, got '%d'",
			t.Name(), want, got)
		return
	}
}

func TestLoadObjectClasses(t *testing.T) {
	want := 38
	if got := tsch.ObjectClasses().Len(); got != want {
		t.Errorf("%s failed: want '%d' objectClasses, got '%d'",
			t.Name(), want, got)
		return
	}
}

/*
func TestParseFile(t *testing.T) {
	path := `/home/jc/dev/schema/oiddir.schema`
	if err := tsch.ParseFile(path); err != nil {
		t.Errorf("%s failed: %v", t.Name(), err)
		return
	}

	want := 260
	if got := tsch.AttributeTypes().Len(); got != want {
		t.Errorf("%s failed: want '%d' attributeTypes, got '%d'",
			t.Name(), want, got)
	}
}
*/

func init() {
	tsch = NewEmptySchema()
	tsch.LoadLDAPSyntaxes()
	tsch.LoadMatchingRules()
	tsch.LoadAttributeTypes()
	tsch.LoadObjectClasses()
}

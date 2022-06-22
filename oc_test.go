package schemax

import (
	"testing"
)

/*
Assemble an ObjectClass manually and evaluate unmarshal reliability.
*/
func TestCompositeObjectClass001(t *testing.T) {
	schema := testSubschema()
	schema.PopulateDefaultAttributeTypes()

	def := NewObjectClass()
	def.OID = NewOID(`1.3.6.1.4.1.56521.999.1.2.4.1`)
	def.Name = NewName(`genericTestClass`, `classAltName`)
	def.Description = Description(`Generic test class`)
	def.Kind = Auxiliary

	for _, must := range []string{`cn`, `o`, `mail`} {
		if at := schema.GetAttributeType(must); !at.IsZero() {
			def.Must.Set(at)
		}
	}

	want := 3
	got := def.Must.Len()

	if want != got {
		t.Errorf("%s validation failed: bad %T length (want %d, got %d)", t.Name(), def.Must, want, got)
		return
	}

	for _, may := range []string{`c`, `description`, `givenName`, `sn`} {
		if at := schema.GetAttributeType(may); !at.IsZero() {
			def.May.Set(at)
		}
	}

	want = 4
	got = def.May.Len()

	if want != got {
		t.Errorf("%s validation failed: bad %T length (want %d, got %d)", t.Name(), def.May, want, got)
		return
	}

	raw, err := Unmarshal(def)
	if err != nil {
		t.Errorf("%s unmarshal failed: %s", t.Name(), err.Error())
		return
	}

	want = 176
	got = len(raw)

	if want != got {
		t.Errorf("%s failed: unexpected raw length (want %d, got %d)", t.Name(), want, got)
	}
}

func TestParseObjectClass001(t *testing.T) {
        def := `( 1.3.6.1.4.1.56521.999.100.2.2.37 NAME 'testClass' DESC 'Generic test class' SUP top STRUCTURAL MUST cn MAY ( l $ c $ o ) )`

        var lsc LDAPSyntaxCollection = PopulateDefaultLDAPSyntaxes()
        var mrc MatchingRuleCollection = PopulateDefaultMatchingRules()
        var atc AttributeTypeCollection = PopulateDefaultAttributeTypes()
        var occ ObjectClassCollection = PopulateDefaultObjectClasses()

        var x ObjectClass

        err := Marshal(def, &x, atc, occ, lsc, mrc, nil, nil, nil, nil)
        if err != nil {
                t.Errorf("%s failed: %s\n", t.Name(), err.Error())
                return
        }

        var um string
        if um, err = Unmarshal(&x); err != nil {
                t.Errorf("%s failed: %s", t.Name(), err.Error())
                return
        }

        // What went in should match
        // what comes out.
        want := len(def)
        got := len(um)

        if want != got {
                t.Errorf("%s failed: unexpected raw length (want %d, got %d)", t.Name(), want, got)
        }
}

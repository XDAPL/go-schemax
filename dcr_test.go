package schemax

import (
	"testing"
)

/*
Assemble a DITContentRule manually and evaluate unmarshal reliability.
*/
func TestCompositeDITContentRule001(t *testing.T) {
	schema := testSubschema()
	schema.PopulateDefaultAttributeTypes()
	schema.PopulateDefaultObjectClasses()

	def := NewDITContentRule()
	def.OID = NewOID(`1.3.6.1.4.1.56521.999.1.2.5.1`)
	def.Name = NewName(`genericTestRule`, `ruleAltName`)
	def.Description = Description(`Generic test rule`)

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

	for _, not := range []string{`co`, `l`} {
		if at := schema.GetAttributeType(not); !at.IsZero() {
			def.Not.Set(at)
		}
	}

	want = 2
	got = def.Not.Len()

	if want != got {
		t.Errorf("%s validation failed: bad %T length (want %d, got %d)", t.Name(), def.Not, want, got)
		return
	}

	for _, aux := range []string{`posixAccount`, `shadowAccount`} {
		if at := schema.GetObjectClass(aux); !at.IsZero() {
			def.Aux.Set(at)
		}
	}

	want = 2
	got = def.Not.Len()

	if want != got {
		t.Errorf("%s validation failed: bad %T length (want %d, got %d)", t.Name(), def.Aux, want, got)
		return
	}

	raw, err := Unmarshal(def)
	if err != nil {
		t.Errorf("%s unmarshal failed: %s", t.Name(), err.Error())
		return
	}

	want = 214
	got = len(raw)

	if want != got {
		t.Errorf("%s failed: unexpected raw length (want %d, got %d)", t.Name(), want, got)
	}
}

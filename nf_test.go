package schemax

import (
	"testing"
)

/*
Assemble a NameForm manually and evaluate unmarshal reliability.
*/
func TestCompositeNameForm001(t *testing.T) {
	schema := testSubschema()
	schema.PopulateDefaultAttributeTypes()
	schema.PopulateDefaultObjectClasses()

	person := schema.GetObjectClass(`person`)
	cn := schema.GetAttributeType(`cn`)
	ou := schema.GetAttributeType(`ou`)

	if cn.IsZero() || ou.IsZero() || person.IsZero() {
		t.Errorf("%s failed: lookup error for supporting AttributeType and/or ObjectClass definition values", t.Name())
		return
	}

	def := NewNameForm()
	def.OID = NewOID(`1.3.6.1.4.1.56521.999.1.2.5.1`)
	def.Name = NewName(`genericNameForm`)
	def.Description = Description(`Generic test nameForm`)
	def.OC = StructuralObjectClass{person}
	def.Must.Set(cn)
	def.May.Set(ou)

	raw, err := Unmarshal(def)
	if err != nil {
		t.Errorf("%s unmarshal failed: %s", t.Name(), err.Error())
		return
	}

	want := 247
	got := len(raw)

	if want != got {
		t.Errorf("%s failed: unexpected raw length (want %d, got %d)", t.Name(), want, got)
	}
}

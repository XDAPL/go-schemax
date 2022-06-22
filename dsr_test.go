package schemax

import (
	"testing"
)

/*
Assemble a DITStructureRule manually and evaluate unmarshal reliability.
*/
func TestCompositeDITStructureRule001(t *testing.T) {
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

	nf := NewNameForm()
	nf.OID = NewOID(`1.3.6.1.4.1.56521.999.1.2.5.1`)
	nf.Name = NewName(`genericNameForm`)
	nf.Description = Description(`Generic test nameForm`)
	nf.OC = StructuralObjectClass{person}
	nf.Must.Set(cn)
	nf.May.Set(ou)

	def := NewDITStructureRule()
	def.ID = NewRuleID(0)
	def.Name = NewName(`genericTestRule`, `ruleAltName`)
	def.Description = Description(`Generic test rule`)
	def.Form = nf

	raw, err := Unmarshal(def)
	if err != nil {
		t.Errorf("%s unmarshal failed: %s", t.Name(), err.Error())
		return
	}

	want := 324
	got := len(raw)

	if want != got {
		t.Errorf("%s failed: unexpected raw length (want %d, got %d)", t.Name(), want, got)
	}
}

package schemax

import (
	"github.com/JesseCoretta/go-schemax/rfc4519"
	"testing"
)

/*
Assemble an AttributeType manually (NO parsing) and evaluate unmarshal reliability.
*/
func TestCompositeAttributeTypeMarshal001(t *testing.T) {
	schema := testSubschema()

	cim := schema.GetMatchingRule(`caseIgnoreMatch`)

	def := new(AttributeType)
	def.OID = NewOID(`1.3.6.1.4.1.56521.999.1.2.3.1`)
	def.Name = NewName(`genericTestAttribute`, `attrAltName`)
	def.Syntax = schema.GetLDAPSyntax(`1.3.6.1.4.1.1466.115.121.1.15`)
	def.Equality = Equality{cim}
	def.Description = Description(`Generic test attribute`)
	def.SetSingleValue()
	def.SetMaxLength(32)

	if err := def.Validate(); err != nil {
		t.Errorf("TestCompositeMarshal failed: %s", err.Error())
		return
	}

	if ok := schema.SetAttributeType(def); !ok {
		t.Errorf("TestCompositeMarshal failed: collection registration failed")
		return
	}

	if um := schema.GetAttributeType(`genericTestAttribute`); um.IsZero() {
		t.Errorf("TestCompositeMarshal failed: test attribute not successfully registered in collection")
	}
}

func TestUnmarshal001(t *testing.T) {
	schema := testSubschema()
	schema.PopulateDefaultAttributeTypes()

	name := schema.GetAttributeType(`name`)
	if name.IsZero() {
		t.Errorf("TestUnmarshal failed: lookup error for 'name' attributeType definition")
		return
	}

	raw, err := Unmarshal(name)
	if err != nil {
		t.Errorf("TestUnmarshal failed: %s", err.Error())
		return
	}

	want := 138
	got := len(raw)
	if got != want {
		t.Errorf("TestUnmarshal failed: unexpected raw length (want %d, got %d)", want, got)
	}
}

/*
TestAttributeMarshal001 tests a variety of marshal operations.
*/
func TestRawAttributeTypeMarshal001(t *testing.T) {
	schema := testSubschema()
	if err := schema.MarshalAttributeType(string(rfc4519.Name)); err != nil {
		t.Errorf("TestAttributeMarshal001 failed: %s", err.Error())
	}

	bogusDef := `( 1.3.6.1.4.1.56521.999.100.2.1.13 NAME failme DESC 'Doomed attribute type' SYNTAX 1.3.6.1.4.1.56521.999.100.44.6 EQUALITY whatsAnEquality ORDERING dontTellMeWhatToDo SUBSTRATE oops SINGLE-VALUE COLLECTIVE X-NOT-HUMAN-READABLE fargus )`
	if err := schema.MarshalAttributeType(bogusDef); err == nil {
		t.Errorf("TestAttributeMarshal002 failed: bogus definition wrongfully accepted")
	}

	if err := schema.MarshalAttributeType(``); err == nil {
		t.Errorf("TestAttributeMarshal003 failed: nil definition wrongfully accepted")
	}
}

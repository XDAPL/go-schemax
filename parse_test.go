package schemax

import (
	"github.com/JesseCoretta/go-schemax/rfc4519"
	"testing"
)

func testSubschema() (x *Subschema) {
	x = NewSubschema()
	x.LSM = PopulateDefaultLDAPSyntaxesManifest()
	x.MRM = PopulateDefaultMatchingRulesManifest()
	x.ATM = NewAttributeTypesManifest()
	x.OCM = NewObjectClassesManifest()

	return
}

/*
TestAttributeMarshal001 tests a variety of marshal operations.
*/
func TestAttributeMarshal001(t *testing.T) {
	schema := testSubschema()
	if err := schema.MarshalAttributeType(string(rfc4519.Name)); err != nil {
		t.Errorf("TestAttributeMarshal001 failed: %s", err.Error())
	}

	bogusDef := `( 1.3.6.1.4.1.56521.999.100.2.1.13 NAME failme DESC 'Doomed attribute type' SYNTAX 1.3.6.1.4.1.56521.999.100.44.6 EQUALITY whatsAnEquality ORDERING dontTellMeWhatToDo SUBSTRATE oops SINGLE-VALUE COLLECTIVE X-NOT-HUMAN-READABLE fargus )`
	if err := schema.MarshalAttributeType(bogusDef); err == nil {
		t.Errorf("TestAttributeMarshal002 failed: bogus definition accepted")
	}

	if err := schema.MarshalAttributeType(``); err == nil {
		t.Errorf("TestAttributeMarshal003 failed: nil definition accepted")
	}
}

/*
TestEquals001 tests equality checks between separate types.
*/
func TestEquals001(t *testing.T) {
	schema := testSubschema()

	if schema.MRM.Equals(schema.LSM) {
		t.Errorf("TestEqualsManifest001 failed: obviously different content wrongly reported to be equal")
	}
}

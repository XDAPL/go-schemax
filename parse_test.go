package schemax

import (
	"github.com/JesseCoretta/go-schemax/rfc4519"
	"testing"
)

func testSubschema() (x *Subschema) {
	x = NewSubschema()
	x.LSC = PopulateDefaultLDAPSyntaxes()
	x.MRC = PopulateDefaultMatchingRules()
	x.ATC = NewAttributeTypes()
	x.OCC = NewObjectClasses()

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

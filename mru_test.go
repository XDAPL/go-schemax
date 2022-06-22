package schemax

import (
	"testing"
)

/*
Assemble a MatchingRule manually and evaluate unmarshal reliability.

Please note it is very unusual for MatchingRuleUse instances to be
created in this fashion. This example is strictly for educational
purposes only. The correct methods for creating MRU instances are:

 - Parse from a raw definition, such as those generated and known by a DSA, or ...

 - Initialize MRU instances using the MatchingRuleUses.Refresh() method.
*/
func TestCompositeMatchingRuleUse001(t *testing.T) {
	sch := NewSubschema()
	sch.LSC = PopulateDefaultLDAPSyntaxes()
	sch.MRC = PopulateDefaultMatchingRules()
	sch.ATC = PopulateDefaultAttributeTypes()

	// Create a new instance of *MatchingRuleUse
	// and populate as needed.
	def := NewMatchingRuleUse()
	def.Name = NewName(`objectIdentifierFirstComponentMatch`)
	def.OID = NewOID(`2.5.13.30`)

	// specify a list of attributetypes applicable to the
	// matchingRule identified by 2.5.13.30.
	applied := []string{
		`objectClasses`,
		`attributeTypes`,
		`matchingRules`,
		`matchingRuleUse`,
		`ldapSyntaxes`,
		`dITContentRules`,
		`nameForms`,
	}

	// Cycle through the above list of applied
	// *AttributeType names, looking-up and
	// setting each one within the MRU.
	for _, apply := range applied {
		if at := sch.GetAttributeType(apply); !at.IsZero() {
			def.Applies.Set(at)
		}
	}

	// Make sure validation checks return NO
	// errors before using your definition!
	um, err := Unmarshal(def)
	if err != nil {
		t.Errorf("%s failed: %s", t.Name(), err.Error())
		return
	}

	want := 178
	got := len(um)

	if want != got {
		t.Errorf("%s failed: unexpected raw length (want %d, got %d)", t.Name(), want, got)
	}
}

/*
test the population of the MatchingRuleUses collection type using the Refresh method.
*/
func TestMatchingRuleUsesRefresh001(t *testing.T) {
	sch := NewSubschema()
	sch.LSC = PopulateDefaultLDAPSyntaxes()
	sch.MRC = PopulateDefaultMatchingRules()
	sch.ATC = PopulateDefaultAttributeTypes()

	if err := sch.MRUC.Refresh(sch.ATC); err != nil {
		t.Errorf("%s failed: %s", t.Name(), err.Error())
		return
	}

	if sch.MRUC.Len() == 0 {
		t.Errorf("%s failed: %T collection is empty, Refresh failed", t.Name(), sch.MRUC)
	}
}

package schemax

import (
	"testing"
)

var testSchema Schema

func TestSchema_Stats(t *testing.T) {
	ct := testSchema.Counters()

	// at the absolute minimum, ldapSyntaxes and
	// matchingRules should be present.
	if ct.LS == 0 || ct.MR == 0 {
		t.Errorf("No statistical %T output", testSchema)
		return
	}
}

func TestSchema_Get(t *testing.T) {
	class := `2.5.6.0`
	if oc := testSchema.ObjectClasses().get(class); oc.IsZero() {
		t.Errorf("Unable to lookup objectClass by numeric OID '%s'", class)
		return
	}
	class = `alias`
	if oc := testSchema.ObjectClasses().get(class); oc.IsZero() {
		t.Errorf("Unable to lookup objectClass by name '%s'", class)
		return
	}
}

// certain high-level parser handlers (which are often unneeded)
// need not be tested since they all execute a common function
// that is already already tested.  This test exists merely to
// satisfy code coverage.
func TestSchema_highLevelParsers(t *testing.T) {
	_ = testSchema.ParseLDAPSyntax(
		`( 1.2.3.4 DESC 'NOT A SYNTAX' X-NOT-HUMAN-READABLE 'TRUE' )`)
	_ = testSchema.ParseMatchingRule(
		`( 1.3.4.5 NAME 'matchingRoool' DESC 'FAKE MATCH' SYNTAX 1.2.3.4 )`)
	_ = testSchema.ParseAttributeType(
		`( 1.4.4.5 NAME 'attriboot' DESC 'lolz' EQUALITY matchingRoool SYNTAX 1.2.3.4 )`)
	_ = testSchema.ParseObjectClass(
		`( 1.5.4.5 NAME 'classy' DESC 'extreme classiness' SUP top AUXILIARY MAY attriboot )`)
	_ = testSchema.ParseDITContentRule(
		`( 0.9.2342.19200300.100.4.13 NAME 'domainContent' DESC 'content control for domain objects' AUX ( dcObject $ domainRelatedObject ) MUST ( dc $ associatedDomain ) MAY description NOT street )`)
	_ = testSchema.ParseNameForm(
		`( 1.3.6.1.4.1.56521.999.5 NAME 'anotherFormidableForm' OC domain MUST dc MAY associatedDomain )`)
	_ = testSchema.ParseDITStructureRule(
		`( 0 NAME 'structureRule' DESC 'structural control' FORM formidableForm )`)
}

// This test will attempt to cause panics by utilizing instances of
// types not initialized properly (or at all).
func TestSchema_panics(t *testing.T) {
	var (
		at AttributeType
		oc ObjectClass
		ls LDAPSyntax
		mr MatchingRule
		mu MatchingRuleUse
		dc DITContentRule
		nf NameForm
		ds DITStructureRule
	)

	_ = ls.OID()
	_ = mr.OID()
	_ = at.OID()
	_ = mu.OID()
	_ = oc.OID()
	_ = dc.OID()
	_ = nf.OID()
	_ = ds.RuleID()
	_ = ds.ID()
}

/*
TestObjectClass_Error cycles a series of bogus objectClass definitions
and incremental portions of bogus objectClass definitions.  During each
iteration, a check is conducted to see whether an error occurred as we
would expect.

Lack of errors indicates serious problems with the underlying antlr4512
package, and/or the manner in which this package uses antlr4512.
*/
func TestObjectClass_Error(t *testing.T) {
	for _, bogus := range []string{
		`( 1.2.3.4.5.6.7.8 NAME ( 'bogusClass' 'cl^assHole' ) DESC 'I am full of shit' )`,
		`( 1.2.3.4.5.6.7.8 NAME ( 'bogusClass' 'classHole' ) DESC 'I am full of shit' SOP top )`,
		`( 1.2.3.4.5.6.7.8 NAME ( 'bogusClass' 'classHole' ) DESC 'I am full of shit' AUXSMILITARY )`,
		`( 1.2.3.4.5.6.7.8 NAME ( 'bogusClass' 'classHole' ) DESC 'I am full of shit' MUSK sucks )`,
	} {

		if err := testSchema.ParseObjectClass(bogus); err == nil {
			t.Errorf("%s error: bogus Objectclass produced no error [BAD]", t.Name())
			return
		}
	}
}

func init() {
	testSchema = NewSchema()
	//printf("%s\n", testSchema.Counters())
}

package schemax

import (
	"fmt"
	"testing"
)

var mySchema Schema

func ExampleSchema_Options() {
	opts := mySchema.Options()
	opts.Shift(AllowOverride)
	fmt.Println(opts.Positive(AllowOverride))
	// Output: true
}

func TestLoadSyntaxes(t *testing.T) {
	want := 67
	if got := mySchema.LDAPSyntaxes().Len(); got != want {
		t.Errorf("%s failed: want '%d' ldapSyntaxes, got '%d'",
			t.Name(), want, got)
		return
	}
}

func TestLoadMatchingRules(t *testing.T) {
	want := 44
	if got := mySchema.MatchingRules().Len(); got != want {
		t.Errorf("%s failed: want '%d' matchingRules, got '%d'",
			t.Name(), want, got)
		return
	}
}

func TestLoadAttributeTypes(t *testing.T) {
	want := 164 // includes supplementals
	if got := mySchema.AttributeTypes().Len(); got != want {
		t.Errorf("%s failed: want '%d' attributeTypes, got '%d'",
			t.Name(), want, got)
		return
	}
}

func TestLoadObjectClasses(t *testing.T) {
	want := 52
	if got := mySchema.ObjectClasses().Len(); got != want {
		t.Errorf("%s failed: want '%d' objectClasses, got '%d'",
			t.Name(), want, got)
		return
	}
}

func TestLoads_codecov(t *testing.T) {
	coolSchema := NewEmptySchema()
	coolSchema.LoadRFC4517Syntaxes()
	coolSchema.LoadRFC2307Syntaxes()
	coolSchema.LoadRFC4523Syntaxes()
	coolSchema.LoadRFC4530Syntaxes()

	coolSchema.LoadRFC4517MatchingRules()
	coolSchema.LoadRFC2307MatchingRules()
	coolSchema.LoadRFC4523MatchingRules()
	coolSchema.LoadRFC4530MatchingRules()

	coolSchema.LoadRFC4512AttributeTypes()
	coolSchema.LoadRFC2079AttributeTypes()
	coolSchema.LoadRFC2798AttributeTypes()
	coolSchema.LoadRFC3045AttributeTypes()
	coolSchema.LoadRFC3672AttributeTypes()
	coolSchema.LoadRFC4519AttributeTypes()
	coolSchema.LoadRFC2307AttributeTypes()
	coolSchema.LoadRFC3671AttributeTypes()
	coolSchema.LoadRFC4523AttributeTypes()
	coolSchema.LoadRFC4524AttributeTypes()
	coolSchema.LoadRFC4530AttributeTypes()

	coolSchema.LoadRFC4512ObjectClasses()
	coolSchema.LoadRFC2079ObjectClasses()
	coolSchema.LoadRFC2798ObjectClasses()
	coolSchema.LoadRFC2307ObjectClasses()
	coolSchema.LoadRFC4512ObjectClasses()
	coolSchema.LoadRFC3671ObjectClasses()
	coolSchema.LoadRFC3672ObjectClasses()
	coolSchema.LoadRFC4519ObjectClasses()
	coolSchema.LoadRFC4523ObjectClasses()
	coolSchema.LoadRFC4524ObjectClasses()

}

/*
func TestParseFile(t *testing.T) {
	path := `/home/jc/dev/schema/oiddir.schema`
	if err := mySchema.ParseFile(path); err != nil {
		t.Errorf("%s failed: %v", t.Name(), err)
		return
	}

	want := 260
	if got := mySchema.AttributeTypes().Len(); got != want {
		t.Errorf("%s failed: want '%d' attributeTypes, got '%d'",
			t.Name(), want, got)
	}
}
*/

// supplemental attributeTypes not sourced from an official doc, but
// are useful in UTs, et al.
var suplATs []string = []string{
	`( 2.5.18.9 NAME 'hasSubordinates' DESC 'X.501: entry has children' EQUALITY booleanMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.7 SINGLE-VALUE NO-USER-MODIFICATION USAGE directoryOperation )`,
}

func init() {
	// Prepare our UT/Example reference schema
	mySchema = NewEmptySchema(
		SortExtensions,
		SortLists,
		HangingIndents,
	)

	funks := []func() error{
		mySchema.LoadLDAPSyntaxes,   // import all built-in Syntaxes
		mySchema.LoadMatchingRules,  // import all built-in Matching Rules
		mySchema.LoadAttributeTypes, // import all built-in Attribute Types
		mySchema.LoadObjectClasses,  // import all built-in Object Classes
	}

	var err error
	for i := 0; i < len(funks) && err == nil; i++ {
		if err = funks[i](); err != nil {
			panic(err)
		}
	}

	// load some supplemental attributeTypes used for special cases
	for _, at := range suplATs {
		if err := mySchema.ParseAttributeType(at); err != nil {
			panic(err)
		}
	}

	// Refresh our matching rules
	mySchema.updateMatchingRuleUses(mySchema.AttributeTypes())
}

package schemax

import (
	"fmt"
	"testing"
)

var mySchema Schema

/*
This example demonstrates the so-called "Quick Start Schema" initialization.
The [NewSchema] function imports all built-in definitions instantly, allowing
the user to start their activities with no fuss.
*/
func ExampleNewSchema() {
	mySchema := NewSchema()
	fmt.Printf("%d types parsed", mySchema.Counters().AT)
	// Output: 164 types parsed
}

func ExampleNewBasicSchema() {
	mySchema := NewBasicSchema()
	fmt.Printf("%d syntaxes parsed", mySchema.Counters().LS)
	// Output: 67 syntaxes parsed
}

func ExampleNewEmptySchema() {
	mySchema := NewEmptySchema()
	fmt.Printf("%d syntaxes parsed", mySchema.Counters().LS)
	// Output: 0 syntaxes parsed
}

func ExampleSchema_Options() {
	opts := mySchema.Options()
	opts.Shift(AllowOverride)
	fmt.Println(opts.Positive(AllowOverride))
	// Output: true
}

func ExampleSchema_Replace_objectClass() {
	mySchema.Options().Shift(AllowOverride)

	gon := mySchema.ObjectClasses().Get(`groupOfNames`)
	ngon := mySchema.NewObjectClass().
		SetNumericOID(gon.NumericOID()).
		SetName(gon.Name()).
		SetDescription(gon.Description()).
		SetKind(gon.Kind()).
		SetSuperClass(`top`).
		SetMust(`cn`).
		SetMay(`member`,
			`businessCategory`,
			`seeAlso`,
			`owner`,
			`ou`,
			`o`,
			`description`).
		SetExtension(`X-ORIGIN`, `RFC4519`).
		SetExtension(`X-WARNING`, `MODIFIED`). // optional
		SetStringer()

	mySchema.Replace(ngon)

	fmt.Println(mySchema.ObjectClasses().Get(`groupOfNames`))
	// Output: ( 2.5.6.9
	//     NAME 'groupOfNames'
	//     SUP top
	//     STRUCTURAL
	//     MUST cn
	//     MAY ( member
	//         $ businessCategory
	//         $ seeAlso
	//         $ owner
	//         $ ou
	//         $ o
	//         $ description )
	//     X-ORIGIN 'RFC4519'
	//     X-WARNING 'MODIFIED' )
}

/*
This example demonstrates refreshing the [MatchingRuleUses] collection
within the receiver instance of [Schema]. The result of this operation
is influence by any new [AttributeType] instances that have been added
since the last refresh.
*/
func ExampleSchema_UpdateMatchingRuleUses() {
	mySchema.UpdateMatchingRuleUses()
	fmt.Printf("%d matchingRuleUses present", mySchema.Counters().MU)
	// Output: 32 matchingRuleUses present
}

/*
This example demonstrates obtaining a non thread-safe [Counters] instance,
which outlines the number of [Definition] instances in categorical fashion.
*/
func ExampleSchema_Counters() {
	fmt.Printf("%d types present", mySchema.Counters().AT)
	// Output: 165 types present
}

/*
This example demonstrates accessing the [Schema] instance's distinguished
name, if set.
*/
func ExampleSchema_DN() {
	fmt.Println(mySchema.DN())
	// Output: cn=schema
}

/*
This example demonstrates specifying a non-standard distinguished name
for use by the [Schema] instance.
*/
func ExampleSchema_SetDN() {
	mySchema := NewEmptySchema()
	mySchema.SetDN(`cn=subschema`)

	fmt.Println(mySchema.DN())
	// Output: cn=subschema
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
	coolSchema.LoadRFC4517Syntaxes()
	coolSchema.LoadRFC2307Syntaxes()
	coolSchema.LoadRFC4523Syntaxes()
	coolSchema.LoadRFC4530Syntaxes()

	coolSchema.LoadRFC4517MatchingRules()
	coolSchema.LoadRFC2307MatchingRules()
	coolSchema.LoadRFC4523MatchingRules()
	coolSchema.LoadRFC4530MatchingRules()

	coolSchema.LoadX501AttributeTypes()
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

// supplemental attributeTypes not sourced from an official doc, but
// are useful in UTs, et al.
var suplATs []string = []string{
	`( 2.5.18.9 NAME 'hasSubordinates' DESC 'X.501: entry has children' EQUALITY booleanMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.7 SINGLE-VALUE NO-USER-MODIFICATION USAGE directoryOperation )`,
}

func init() {
	// Prepare our UT/Example reference schema
	mySchema = NewEmptySchema(
		AllowOverride,
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

package schemax

import (
	"fmt"
)

/*
This example demonstrates the creation of a [DITStructureRule].
*/
func ExampleNewDITStructureRule() {
	// First create a name form that requires an
	// RDN of uid=<val>, or (optionally) an RDN
	// of uid=<val>+gidNumber=<val>
	nf := mySchema.NewNameForm().
		SetNumericOID(`1.3.6.1.4.1.56521.999.16.7`).
		SetName(`personForm`).
		SetDescription(`generalized person name form`).
		SetOC(`person`).
		SetMust(`uid`).
		SetMay(`gidnumber`).
		SetStringer()

	// Create the structure rule and assign the
	// new nameform
	ds := mySchema.NewDITStructureRule().
		SetRuleID(0).
		SetName(`personStructure`).
		SetDescription(`person structure rule`).
		SetForm(nf).
		SetStringer()

	fmt.Println(ds)
	// Output: ( 0
	//     NAME 'personStructure'
	//     DESC 'person structure rule'
	//     FORM personForm )
}

func ExampleDITStructureRule_Compliant() {
	nf := mySchema.NewNameForm().
		SetNumericOID(`1.3.6.1.4.1.56521.999.16.7`).
		SetName(`personForm`).
		SetDescription(`generalized person name form`).
		SetOC(`person`).
		SetMust(`uid`).
		SetMay(`gidnumber`).
		SetStringer()

		//mySchema.NameForms().Push(nf)

	// Create the structure rule and assign the
	// new nameform
	ds := mySchema.NewDITStructureRule().
		SetRuleID(0).
		SetName(`personStructure`).
		SetDescription(`person structure rule`).
		SetForm(nf).
		SetStringer()

	fmt.Println(ds.Compliant())
	// Output: true
}

/*
This example demonstrates the assignment of arbitrary data to an instance
of [AttributeType].

Note: this example assumes a legitimate schema variable is defined
in place of the fictional "mySchema" var shown here for simplicity.
*/
func ExampleDITStructureRule_SetData() {
	nf := mySchema.NewNameForm().
		SetNumericOID(`1.3.6.1.4.1.56521.999.16.7`).
		SetName(`personForm`).
		SetDescription(`generalized person name form`).
		SetOC(`person`).
		SetMust(`uid`).
		SetMay(`gidnumber`).
		SetStringer()

	//mySchema.NameForms().Push(nf)

	// Create the structure rule and assign the
	// new nameform
	ds := mySchema.NewDITStructureRule().
		SetRuleID(0).
		SetName(`personStructure`).
		SetDescription(`person structure rule`).
		SetForm(nf).
		SetExtension(`X-ORIGIN`, `NOWHERE`).
		SetStringer()

	// The value can be any type, but we'll
	// use a string for simplicity.
	documentation := `Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.`

	// Set it.
	ds.SetData(documentation)

	// Get it and compare to the original.
	equal := documentation == ds.Data().(string)

	fmt.Printf("Values are equal: %t", equal)
	// Output: Values are equal: true
}

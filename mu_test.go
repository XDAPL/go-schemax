package schemax

import "fmt"

/*
This example demonstrates manual assembly of a new MatchingRuleUse
instance. Note this is provided for demonstration purposes only and
in context does not perform anything useful.

In general it is not necessary for end-users to manually define
this kind of instance.  Instances of this type are normally created
by automated processes when new [AttributeType] definitions are created
or introduced which make use of a given [MatchingRule] instance.
*/
func ExampleNewMatchingRuleUse() {
	var def MatchingRuleUse = NewMatchingRuleUse().SetSchema(mySchema)

	def.SetNumericOID(`2.5.13.16`).
		SetName(`fakeBitStringMatch`).
		SetExtension(`X-ORIGIN`, `NOWHERE`)

	for _, apl := range []AttributeType{
		mySchema.AttributeTypes().Get(`cn`),
		mySchema.AttributeTypes().Get(`sn`),
		mySchema.AttributeTypes().Get(`l`),
	} {
		def.SetApplies(apl)
	}

	// We're done and ready, set the stringer
	def.SetStringer()

	fmt.Printf("%s", def)
	// Output: ( 2.5.13.16
	//     NAME 'fakeBitStringMatch'
	//     APPLIES ( cn
	//             $ sn
	//             $ l )
	//     X-ORIGIN 'NOWHERE' )
}

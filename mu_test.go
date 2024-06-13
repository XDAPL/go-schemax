package schemax

import "fmt"

/*
This example demonstrates the process of parsing a raw string-based
matchingRule definition into a proper instance of [MatchingRule].

Note: this example assumes a legitimate schema variable is defined
in place of the fictional "mySchema" var shown here for simplicity.
*/
func ExampleMatchingRuleUse_Parse() {
	mu := mySchema.MatchingRuleUses().Index(0)
	def := mySchema.NewMatchingRuleUse()
	if err := def.Parse(mu.String()); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(def)
	// Output: ( 2.5.13.1
	//     NAME 'distinguishedNameMatch'
	//     APPLIES ( aliasedObjectName
	//             $ associatedName
	//             $ collectiveAttributeSubentries
	//             $ creatorsName
	//             $ distinguishedName
	//             $ documentAuthor
	//             $ manager
	//             $ modifiersName
	//             $ secretary
	//             $ subschemaSubentry )
	//     X-ORIGIN 'RFC4517' )
}

/*
This example demonstrates manual assembly of a new MatchingRuleUse
instance. Note this is provided for demonstration purposes only.

In general it is not necessary for end-users to manually define
this kind of instance.  Instances of this type are normally created
by automated processes when new [AttributeType] definitions are created
or introduced which make use of a given [MatchingRule] instance.
*/
func ExampleNewMatchingRuleUse() {

	UseHangingIndents = true

	var def MatchingRuleUse = NewMatchingRuleUse()

	def.SetNumericOID(`1.3.6.1.4.1.56521.999.88.5`).
		SetName(`pulsarMatchingRuleUses`).
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
	// Output: ( 1.3.6.1.4.1.56521.999.88.5
	//     NAME 'pulsarMatchingRuleUses'
	//     APPLIES ( cn
	//             $ sn
	//             $ l )
	//     X-ORIGIN 'NOWHERE' )
}

package schemax

import "fmt"

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

	for _, apl := range []AttributeType {
		tsch.AttributeTypes().Get(`cn`),
		tsch.AttributeTypes().Get(`sn`),
		tsch.AttributeTypes().Get(`l`),
	} {
		def.SetApplies(apl)
	}

	fmt.Printf("%s", def)
	// Output: ( 1.3.6.1.4.1.56521.999.88.5
	//     NAME 'pulsarMatchingRuleUses'
	//     APPLIES ( cn $ sn $ l )
	//     X-ORIGIN 'NOWHERE' )
}

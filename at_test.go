package schemax

import "fmt"

func ExampleNewAttributeType() {
	UseHangingIndents = true

	// lookup and get the Directory String syntax
	dStr := tsch.LDAPSyntaxes().Get(`1.3.6.1.4.1.1466.115.121.1.15`)
	if dStr.IsZero() {
		return
	}

	// lookup and get the caseIgnoreMatch equality matching rule
	cIM := tsch.MatchingRules().Get(`caseIgnoreMatch`)
	if cIM.IsZero() {
		return
	}

	// prepare new var instance
	var def AttributeType = NewAttributeType()

	// set values in fluent form
	def.SetNumericOID(`1.3.6.1.4.1.56521.999.5`).
		SetName(`cb`).
		SetDescription(`Celestial Body`).
		SetSyntax(dStr).
		SetEquality(cIM).
		SetSingleValue(true).
		SetExtension(`X-ORIGIN`, `NOWHERE`)

	fmt.Printf("%s", def)
	// Output: ( 1.3.6.1.4.1.56521.999.5
	//     NAME 'cb'
	//     DESC 'Celestial Body'
	//     EQUALITY caseIgnoreMatch
	//     SYNTAX 1.3.6.1.4.1.1466.115.121.1.15
	//     SINGLE-VALUE
	//     X-ORIGIN 'NOWHERE' )
}

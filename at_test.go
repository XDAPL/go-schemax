package schemax

import (
	"fmt"
	"strings"
	//"testing"
)

/*
This example demonstrates the means of gathering references to every
superior [AttributeType] in the relevant super type chain.
*/
func ExampleAttributeType_SuperChain() {
	cn := mySchema.AttributeTypes().Get(`nisMapName`)
	chain := cn.SuperChain()

	for i := 0; i < chain.Len(); i++ {
		fmt.Println(chain.Index(i).OID())
	}

	// Output: name
}

/*
This example demonstrates the means of walking the super type chain to
determine the effective [LDAPSyntax] instance held by an [AttributeType]
instance, whether direct or indirect.
*/
func ExampleAttributeType_EffectiveSyntax() {
	at := mySchema.AttributeTypes().Get(`roleOccupant`)
	fmt.Println(at.EffectiveSyntax().Description())
	// Output: DN
}

/*
This example demonstrates the means of walking the super type chain to
determine the effective EQUALITY [MatchingRule] instance held by an
[AttributeType] instance, whether direct or indirect.
*/
func ExampleAttributeType_EffectiveEquality() {
	at := mySchema.AttributeTypes().Get(`registeredAddress`)
	fmt.Println(at.EffectiveEquality().OID())
	// Output: caseIgnoreListMatch
}

/*
This example demonstrates the means of walking the super type chain
to determine the effective SUBSTR [MatchingRule] instance held by an
[AttributeType] instance, whether direct or indirect.
*/
func ExampleAttributeType_EffectiveSubstring() {
	at := mySchema.AttributeTypes().Get(`registeredAddress`)
	fmt.Println(at.EffectiveSubstring().OID())
	// Output: caseIgnoreListSubstringsMatch
}

/*
This example demonstrates the means of walking the super type chain
to determine the effective ORDERING [MatchingRule] instance held by
an [AttributeType] instance, whether direct or indirect.
*/
func ExampleAttributeType_EffectiveOrdering() {
	at := mySchema.AttributeTypes().Get(`createTimestamp`)
	fmt.Println(at.EffectiveOrdering().OID())
	// Output: generalizedTimeOrderingMatch
}

/*
This example demonstrates a conventional means of checking a given
value under the terms of a specific [AttributeType]'s assigned
[ValueQualifier].

Naturally this example is overly simplified, with support extended
for nil value states purely for educational purposes only.  A real
life implementation would likely be more stringent.

Note: this example assumes a legitimate schema variable is defined
in place of the fictional "mySchema" var shown here for simplicity.
*/
func ExampleAttributeType_QualifyValue_withSet() {
	// Let's use "hasSubordinates" due to its common
	// use within multiple popular directory products.
	hS := mySchema.AttributeTypes().Get(`hasSubordinates`)
	if hS.IsZero() {
		fmt.Println("hasSubordinates not found!")
		return
	}

	// hasSubordinates is a BOOLEAN type, so let's make
	// a very "forgiving" handler for such values, both
	// for string and (Go) bool/*bool values.
	hS.SetValueQualifier(func(x any) error {
		var err error

		switch tv := x.(type) {
		case string:
			switch strings.ToLower(tv) {
			case `true`, `false`, `undefined`, ``:
				// OK: Match all valid string values in one shot
				// in a manner compliant with the caseIgnoreMatch
				// equality matching rule.
			default:
				// BAD: No other string value is applicable here.
				err = ErrInvalidSyntax
			}
		case bool, *bool, nil:
			// OK: Guaranteed to be valid, with a nil instance
			// equivalent to the LDAP "Undefined" BOOLEAN state.
		default:
			// BAD: no other type is applicable here.
			err = ErrInvalidType
		}

		return err
	})

	// Let's subject our newly-assigned SyntaxQualifier to
	// a series of valid values of supported types.
	for _, possibleValue := range []any{
		`True`,
		false,
		`False`,
		true,
		`FALSE`,
		`fALse`,
		``,
		`undefineD`,
		nil,
	} {
		// None of these should return errors.
		if err := hS.QualifyValue(possibleValue); err != nil {
			fmt.Println(err)
			return
		}
	}

	// Let's pass a known bogus value just to
	// make sure this thing is indeed working.
	err := hS.QualifyValue(`falsch`) // No German support :(
	fmt.Println(err)
	// Output: Value does not meet the prescribed syntax qualifications
}

/*
This example demonstrates a conventional means of defining a new
[AttributeType] instance.

Note: this example assumes a legitimate schema variable is defined
in place of the fictional "mySchema" var shown here for simplicity.
*/
func ExampleNewAttributeType() {
	// lookup and get the Directory String syntax
	dStr := mySchema.LDAPSyntaxes().Get(`1.3.6.1.4.1.1466.115.121.1.15`)
	if dStr.IsZero() {
		return
	}

	// lookup and get the caseIgnoreMatch equality matching rule
	cIM := mySchema.MatchingRules().Get(`caseIgnoreMatch`)
	if cIM.IsZero() {
		return
	}

	// prepare new var instance
	var def AttributeType = NewAttributeType()

	// set values in fluent form
	def.SetSchema(mySchema).
		SetNumericOID(`1.3.6.1.4.1.56521.999.5`).
		SetName(`cb`).
		SetDescription(`Celestial Body`).
		SetMinimumUpperBounds(64).
		SetSyntax(dStr).
		SetEquality(cIM).
		SetSingleValue(true).
		SetExtension(`X-ORIGIN`, `NOWHERE`).
		SetStringer() // use default stringer

	fmt.Printf("%s", def)
	// Output: ( 1.3.6.1.4.1.56521.999.5
	//     NAME 'cb'
	//     DESC 'Celestial Body'
	//     EQUALITY caseIgnoreMatch
	//     SYNTAX 1.3.6.1.4.1.1466.115.121.1.15{64}
	//     SINGLE-VALUE
	//     X-ORIGIN 'NOWHERE' )
}

package schemax

import (
	"fmt"
	"strings"
	//"testing"
)

/*
This example demonstrates a conventional means of checking a given
value under the terms of a specific [AttributeType]'s assigned
[SyntaxQualifier].

Naturally this example is overly simplified, with support extended
for nil value states purely for educational purposes only.  A real
life implementation would likely be more stringent.

Note: this example assumes a legitimate schema variable is defined
in place of the fictional "mySchema" var shown here for simplicity.
*/
func ExampleAttributeType_CheckValueSyntax_withSet() {
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
	hS.SetSyntaxQualifier(func(x any) error {
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
		if err := hS.CheckValueSyntax(possibleValue); err != nil {
			fmt.Println(err)
			return
		}
	}

	// Let's pass a known bogus value just to
	// make sure this thing is indeed working.
	err := hS.CheckValueSyntax(`falsch`) // No German support :(
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
	UseHangingIndents = true

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
	def.SetNumericOID(`1.3.6.1.4.1.56521.999.5`).
		SetName(`cb`).
		SetDescription(`Celestial Body`).
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
	//     SYNTAX 1.3.6.1.4.1.1466.115.121.1.15
	//     SINGLE-VALUE
	//     X-ORIGIN 'NOWHERE' )
}

package schemax

import "fmt"

func ExampleNewLDAPSyntax() {

	UseHangingIndents = true

	// prepare new var instance
	var def LDAPSyntax = NewLDAPSyntax()

	// set values in fluent form
	def.SetNumericOID(`1.3.6.1.4.1.56521.999.5`).
		SetDescription(`pulsarFrequencySyntax`).
		SetExtension(`X-ORIGIN`, `NOWHERE`).
		SetExtension(`X-NOT-HUMAN-READABLE`, `TRUE`)

	fmt.Printf("%s", def)
	// Output: ( 1.3.6.1.4.1.56521.999.5
	//     DESC 'pulsarFrequencySyntax'
	//     X-ORIGIN 'NOWHERE'
	//     X-NOT-HUMAN-READABLE 'TRUE' )
}

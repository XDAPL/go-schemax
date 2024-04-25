package schemax

import "fmt"

func ExampleNewMatchingRule() {

	UseHangingIndents = true

        // lookup and get the Directory String syntax                   
        integer := tsch.LDAPSyntaxes().Get(`1.3.6.1.4.1.1466.115.121.1.27`) // INTEGER
        if integer.IsZero() {                                              
                return                                                  
        }                                                               

	var def MatchingRule = NewMatchingRule()
                                                                        
	// set values in fluent form
	def.SetNumericOID(`1.3.6.1.4.1.56521.999.88.5`).
		SetName(`pulsarMatchingRule`).
		SetSyntax(integer).
		SetExtension(`X-ORIGIN`, `NOWHERE`)

	fmt.Printf("%s", def)
	// Output: ( 1.3.6.1.4.1.56521.999.88.5
	//     NAME 'pulsarMatchingRule'
	//     SYNTAX 1.3.6.1.4.1.1466.115.121.1.27
	//     X-ORIGIN 'NOWHERE' )
}

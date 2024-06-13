package schemax

import "fmt"

func ExampleNewNameForm() {

	UseHangingIndents = true

	// lookup and get the Directory String syntax
	dvc := mySchema.ObjectClasses().Get(`device`)
	if dvc.IsZero() {
		return
	}

	var def NameForm = NewNameForm()

	// set values in fluent form
	def.SetSchema(mySchema).
		SetNumericOID(`1.3.6.1.4.1.56521.999.97.7`).
		SetName(`deviceNameForm`).
		SetOC(dvc).
		SetMust(`cn`).
		SetStringer()

	fmt.Printf("%s", def)
	// Output: ( 1.3.6.1.4.1.56521.999.97.7
	//     NAME 'deviceNameForm'
	//     OC device
	//     MUST cn )
}

package schemax

import (
	"fmt"
	"testing"
)

/*
This example demonstrates the creation of a new [LDAPSyntax]
instance which will be replaced in memory by another. This change
will be recognized in any and all stacks in which the replaced
[LDAPSyntax] resides.

Note: this example assumes a legitimate schema variable is defined
in place of the fictional "mySchema" var shown here for simplicity.
*/
func ExampleLDAPSyntax_Replace() {

	// Here is our bad version
	orig := mySchema.NewLDAPSyntax().
		SetNumericOID(`1.3.6.1.4.1.56521.999.5`).
		SetDescription(`freakwency`).
		SetExtension(`X-OERIGIN`, `NOWHERE`).
		SetStringer()

	// Here is our good version
	good := NewLDAPSyntax().
		SetSchema(mySchema).
		SetNumericOID(`1.3.6.1.4.1.56521.999.5`).
		SetDescription(`Frequency`).
		SetExtension(`X-ORIGIN`, `NOWHERE`).
		SetStringer()

	// Swap orig for good, but while preserving
	// the same pointer address to keep our
	// references valid.
	orig.Replace(good)

	fmt.Printf("%s", orig)
	// Output: ( 1.3.6.1.4.1.56521.999.5
	//     DESC 'Frequency'
	//     X-ORIGIN 'NOWHERE' )
}

func ExampleLDAPSyntax_IsIdentifiedAs() {
	ls := mySchema.LDAPSyntaxes().Get(`integer`)
	fmt.Println(ls.IsIdentifiedAs(`1.3.6.1.4.1.1466.115.121.1.27`))
	// Output: true
}

func ExampleLDAPSyntax_Description() {
	integer := mySchema.LDAPSyntaxes().Get(`integer`)
	fmt.Println(integer.Description())
	// Output: INTEGER
}

func ExampleLDAPSyntax_NumericOID() {
	integer := mySchema.LDAPSyntaxes().Get(`integer`)
	fmt.Println(integer.NumericOID())
	// Output: 1.3.6.1.4.1.1466.115.121.1.27
}

func ExampleLDAPSyntax_OID() {
	integer := mySchema.LDAPSyntaxes().Get(`integer`)
	fmt.Println(integer.OID())
	// Output: 1.3.6.1.4.1.1466.115.121.1.27
}

func ExampleLDAPSyntax_Extensions() {
	integer := mySchema.LDAPSyntaxes().Get(`integer`)
	fmt.Println(integer.Extensions())
	// Output: X-ORIGIN 'RFC4517'
}

func ExampleLDAPSyntax_Map() {
	integer := mySchema.LDAPSyntaxes().Get(`integer`)
	fmt.Println(integer.Map()[`NUMERICOID`][0])
	// Output: 1.3.6.1.4.1.1466.115.121.1.27
}

func ExampleLDAPSyntaxes_Maps() {
	maps := mySchema.LDAPSyntaxes().Maps()
	fmt.Println(maps[3][`NUMERICOID`][0])
	// Output: 1.3.6.1.4.1.1466.115.121.1.4
}

/*
This example demonstrates the process of parsing a raw string-based
ldapSyntax definition into a proper instance of [LDAPSyntax].

Note: this example assumes a legitimate schema variable is defined
in place of the fictional "mySchema" var shown here for simplicity.
*/
func ExampleLDAPSyntax_Parse() {
	var raw string = `( 1.3.6.1.4.1.56521.999.5 DESC 'pulsarFrequencySyntax' X-NOT-HUMAN-READABLE 'TRUE' X-ORIGIN 'NOWHERE' )`
	var def LDAPSyntax = mySchema.NewLDAPSyntax()
	if err := def.Parse(raw); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(def.SetStringer())
	// Output: ( 1.3.6.1.4.1.56521.999.5
	//     DESC 'pulsarFrequencySyntax'
	//     X-NOT-HUMAN-READABLE 'TRUE'
	//     X-ORIGIN 'NOWHERE' )
}

/*
This example demonstrates the creation of a new [LDAPSyntax]
instance for manual assembly in a fluent manner.

Note: this example assumes a legitimate schema variable is defined
in place of the fictional "mySchema" var shown here for simplicity.
*/
func ExampleNewLDAPSyntax_fluent() {
	// prepare new var instance and
	// set values in fluent form
	def := NewLDAPSyntax().
		SetSchema(mySchema).
		SetNumericOID(`1.3.6.1.4.1.56521.999.5`).
		SetDescription(`pulsarFrequencySyntax`).
		SetExtension(`X-NOT-HUMAN-READABLE`, `TRUE`).
		SetExtension(`X-ORIGIN`, `NOWHERE`).
		SetStringer() // default closure

	fmt.Printf("%s", def)
	// Output: ( 1.3.6.1.4.1.56521.999.5
	//     DESC 'pulsarFrequencySyntax'
	//     X-NOT-HUMAN-READABLE 'TRUE'
	//     X-ORIGIN 'NOWHERE' )
}

/*
This example demonstrates the creation of a new [LDAPSyntax]
instance for manual assembly piecemeal.

Note: this example assumes a legitimate schema variable is defined
in place of the fictional "mySchema" var shown here for simplicity.
*/
func ExampleNewLDAPSyntax_piecemeal() {
	// prepare new var instance and
	// set values in fluent form
	var def LDAPSyntax = NewLDAPSyntax() // initialization always required

	def.SetNumericOID(`1.3.6.1.4.1.56521.999.5`)

	// ... do other things ...

	def.SetExtension(`X-ORIGIN`, `NOWHERE`)
	def.SetDescription(`pulsarFrequencySyntax`)
	def.SetExtension(`X-NOT-HUMAN-READABLE`, `TRUE`)

	// Set default closure if, and only if, the definition
	// is deemed to be RFC compliant. In this case, print
	// the string representation as our final act.
	if def.Compliant() {
		def.SetStringer()
		fmt.Printf("%s", def)
	}

	// Output: ( 1.3.6.1.4.1.56521.999.5
	//     DESC 'pulsarFrequencySyntax'
	//     X-ORIGIN 'NOWHERE'
	//     X-NOT-HUMAN-READABLE 'TRUE' )
}

/*
This example demonstrates instant compliance checks for all [LDAPSyntax]
instances present within an instance of [LDAPSyntaxes].
*/
func ExampleLDAPSyntaxes_Compliant() {
	syns := mySchema.LDAPSyntaxes()
	fmt.Printf("All %d %s are compliant: %t", syns.Len(), syns.Type(), syns.Compliant())
	// Output: All 67 ldapSyntaxes are compliant: true
}

func ExampleLDAPSyntax_Data() {
	syn := mySchema.LDAPSyntaxes().Get(`integer`)

	// Let's pretend img ([]uint8) represents
	// some JPEG data (e.g.: a diagram)
	var img []uint8 = []uint8{0x1, 0x2, 0x3, 0x4}
	syn.SetData(img)

	got := syn.Data().([]uint8)

	fmt.Printf("%T, Len:%d", got, len(got))
	// Output: []uint8, Len:4
}

func ExampleLDAPSyntax_SetData() {
	syn := mySchema.LDAPSyntaxes().Get(`integer`)

	// Let's pretend img ([]uint8) represents
	// some JPEG data (e.g.: a diagram)
	var img []uint8 = []uint8{0x1, 0x2, 0x3, 0x4}
	syn.SetData(img)

	got := syn.Data().([]uint8)

	fmt.Printf("%T, Len:%d", got, len(got))
	// Output: []uint8, Len:4
}

/*
This example demonstrates use of the [LDAPSyntaxes.Type] method to determine
the type of stack defined within the receiver. This is mainly useful in cases
where multiple stacks are being iterated in [Definitions] interface contexts
and is more efficient when compared to manual type assertion.
*/
func ExampleLDAPSyntaxes_Type() {
	syns := mySchema.LDAPSyntaxes()
	fmt.Printf("We have %d %s", syns.Len(), syns.Type())
	// Output: We have 67 ldapSyntaxes
}

/*
This example demonstrates the means of accessing the integer length of
an [LDAPSyntaxes] stack instance.
*/
func ExampleLDAPSyntaxes_Len() {
	syns := mySchema.LDAPSyntaxes()
	fmt.Printf("We have %d %s", syns.Len(), syns.Type())
	// Output: We have 67 ldapSyntaxes
}

/*
This example demonstrates the means of accessing a specific slice value
within an instance of [LDAPSyntaxes] by way of its associated integer
index.
*/
func ExampleLDAPSyntaxes_Index() {
	slice := mySchema.LDAPSyntaxes().Index(3)
	fmt.Println(slice)
	// Output: ( 1.3.6.1.4.1.1466.115.121.1.4
	//     DESC 'Audio'
	//     X-NOT-HUMAN-READABLE 'TRUE'
	//     X-ORIGIN 'RFC4517' )
}

/*
Do stupid things to make schemax panic, gain additional
coverage in the process.
*/
func TestLDAPSyntax_codecov(t *testing.T) {
	_ = mySchema.LDAPSyntaxes().SetStringer().Contains(``)
	mySchema.LDAPSyntaxes().Push(rune(10))
	mySchema.LDAPSyntaxes().IsZero()
	mySchema.LDAPSyntaxes().String()
	cim := mySchema.LDAPSyntaxes().Get(`caseIgnoreMatch`)
	mySchema.LDAPSyntaxes().canPush()
	mySchema.LDAPSyntaxes().canPush(``, ``, ``, ``, cim)
	mySchema.LDAPSyntaxes().canPush(cim, cim)
	bmr := newCollection(``)
	LDAPSyntaxes(bmr.cast()).Push(NewLDAPSyntax().SetSchema(mySchema))
	LDAPSyntaxes(bmr.cast()).Push(NewLDAPSyntax().SetSchema(mySchema).SetNumericOID(`1.2.3.4.5`))
	bmr.cast().Push(NewLDAPSyntax().SetSchema(mySchema))
	bmr.cast().Push(NewLDAPSyntax().SetSchema(mySchema).SetNumericOID(`1.2.3.4.5`))
	LDAPSyntaxes(bmr).Compliant()

	var def LDAPSyntax

	_ = def.String()
	_ = def.SetStringer()
	_ = def.Description()
	_ = def.Name()
	_ = def.Names()
	_ = def.Extensions()
	_ = def.Schema()
	_ = def.Map()
	_ = def.Compliant()
	_ = def.macro()
	_ = def.Obsolete()

	def.setOID(`4.3.2.1`)
	var raw string = `( 1.3.6.1.4.1.56521.999.88.5 DESC 'frequency' X-ORIGIN 'NOWHERE' )`
	if err := def.Parse(raw); err != ErrNilReceiver {
		t.Errorf("%s failed: expected ErrNilReceiver, got %v", t.Name(), err)
		return
	}

	def = NewLDAPSyntax()
	def.SetDescription(`'a`)
	def.SetDescription(`'Unnecessary quoted value to be overwritten'`)

	if err := def.Parse(raw); err != ErrNilSchemaRef {
		t.Errorf("%s failed: expected ErrNilSchemaRef, got %v", t.Name(), err)
		return
	}

	// Try again. Properly.
	def.SetSchema(mySchema)
	if def.Schema().IsZero() {
		t.Errorf("%s failed: no schema reference!", t.Name())
		return
	}
	def.setStringer(func() string {
		return "blarg"
	})

	def.SetData(`fake`)
	def.SetData(nil)
	def.Data()

	if err := def.Parse(raw); err != nil {
		t.Errorf("%s failed: expected success, got %v", t.Name(), err)
		return
	}
	_ = def.macro()
	def.setOID(`1.3.6.1.4.1.56521.999.88.5`)

	var def2 LDAPSyntax
	_ = def2.Replace(def) // will fail

}

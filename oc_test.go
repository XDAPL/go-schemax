package schemax

import (
	"fmt"
	"testing"
)

/*
This example demonstrates the means of checking superiority of a class
over another class by way of the [ObjectClass.SuperClassOf] method.

Note: this example assumes a legitimate schema variable is defined
in place of the fictional "mySchema" var shown here for simplicity.
*/
func ExampleObjectClass_SuperClassOf() {
	top := mySchema.ObjectClasses().Get(`top`)
	acc := mySchema.ObjectClasses().Get(`account`)

	fmt.Println(top.SuperClassOf(acc))
	// Output: true
}

/*
This example demonstrates the means of gathering references to every
superior [ObjectClass] in the relevant super class chain.
*/
func ExampleObjectClass_SuperChain() {
	inet := mySchema.ObjectClasses().Get(`inetOrgPerson`)

	oc := inet.SuperChain()
	for i := 0; i < oc.Len(); i++ {
		fmt.Println(oc.Index(i).OID())
	}

	// Output: organizationalPerson
	// person
	// top
}

func ExampleObjectClass_IsIdentifiedAs() {
	oc := mySchema.ObjectClasses().Get(`account`)
	fmt.Println(oc.IsIdentifiedAs(`0.9.2342.19200300.100.4.5`))
	// Output: true
}

/*
This example demonstrates a common (and most unfortunate) modification to an
OFFICIAL [ObjectClass] definition -- "groupOfNames", found within Section 3.5
of RFC 4519.

The design of this particular class is widely considered to be inconvenient
due to its mandate that at least one (1) instance of the "member" [AttributeType]
(from Section 2.17 of RFC 4519).

As such, this has forced many LDAP architects to literally modify this [ObjectClass]
definition within the given directory schema, moving the "member" [AttributeType]
from the MUST clause to the MAY clause.

For reasons of oversight, we've added the RFC of origin as an X-ORIGIN extension, and
a custom extension X-WARNING to remind users and admin alike that we've resorted to
this risky trick.
*/
func ExampleObjectClass_Replace() {
	// make sure we enable the AllowOverride bit in
	// our Schema instance early in its initialization
	//mySchema.Options().Shift(AllowOverride)

	// Same for HangingIndents - must be done prior
	// to the parsing/loading of ANY definitions
	// in a given Schema instance.
	//mySchema.Options().Shift(HangingIndents)

	// Obtain the groupOfNames (gon) ObjectClass so
	// we can copy some of its values.
	gon := mySchema.ObjectClasses().Get(`groupOfNames`)

	// Craft a near identical groupOfNames instance,
	// save for the one change we intend to make.
	ngon := mySchema.NewObjectClass().
		SetName(gon.Name()).
		SetNumericOID(gon.NumericOID()).
		SetKind(gon.Kind()).
		SetSuperClass(`top`).
		SetMust(`cn`).
		SetMay(`member`, `businessCategory`, `seeAlso`, `owner`, `ou`, `o`, `description`).
		SetExtension(`X-ORIGIN`, `RFC4519`).
		SetExtension(`X-WARNING`, `MODIFIED`). // optional
		SetStringer()

	// Replace gon with ngon, while preserving its pointer
	// address so that references within stacks do not fail.
	gon.Replace(ngon)

	// call the new one (just to be sure)
	fmt.Println(mySchema.ObjectClasses().Get(`groupOfNames`))
	// Output: ( 2.5.6.9
	//     NAME 'groupOfNames'
	//     SUP top
	//     STRUCTURAL
	//     MUST cn
	//     MAY ( member
	//         $ businessCategory
	//         $ seeAlso
	//         $ owner
	//         $ ou
	//         $ o
	//         $ description )
	//     X-ORIGIN 'RFC4519'
	//     X-WARNING 'MODIFIED' )
}

func ExampleObjectClass_SetObsolete() {
	fake := NewObjectClass().
		SetNumericOID(`1.3.6.1.4.1.56521.999.108.4`).
		SetName(`obsoleteClass`).
		SetObsolete()

	fmt.Println(fake.Obsolete())
	// Output: true
}

func ExampleObjectClasses_Inventory() {
	oc := mySchema.ObjectClasses().Inventory()
	fmt.Println(oc[`2.5.6.7`][0])
	// Output: organizationalPerson
}

func ExampleObjectClasses_Type() {
	oc := mySchema.ObjectClasses()
	fmt.Println(oc.Type())
	// Output: objectClasses
}

func ExampleObjectClass_Type() {
	var def ObjectClass
	fmt.Println(def.Type())
	// Output: objectClass
}

func ExampleObjectClass_Map() {
	def := mySchema.ObjectClasses().Get(`account`)
	fmt.Println(def.Map()[`NUMERICOID`][0]) // risky, just for simplicity
	// Output: 0.9.2342.19200300.100.4.5
}

/*
This example demonstrates use of the [ObjectClasses.Maps] method, which
produces slices of [DefinitionMap] instances born of the [ObjectClasses]
stack in which they reside.  We (quite recklessly) call index three (3)
and reference index zero (0) of its `SYNTAX` key to obtain the relevant
[LDAPSyntax] OID string value.
*/
func ExampleObjectClasses_Maps() {
	defs := mySchema.ObjectClasses().Maps()
	fmt.Println(defs[3][`NUMERICOID`][0]) // risky, just for simplicity
	// Output: 1.3.6.1.4.1.1466.101.120.111
}

func ExampleObjectClass_Attributes() {
	ats := mySchema.ObjectClasses().Get(`posixAccount`).Attributes()
	for i := 0; i < ats.Len(); i++ {
		at := ats.Index(i)
		fmt.Println(at.OID())
	}
	// Output: cn
	// gidNumber
	// homeDirectory
	// uid
	// uidNumber
	// description
	// gecos
	// loginShell
	// userPassword
}

/*
This example demonstrates the means of gathering a list of all possible
[AttributeType] instances -- by OID -- that are either required or allowed
by an [ObjectClass] instance.
*/
func ExampleObjectClass_AllAttributes() {
	ats := mySchema.ObjectClasses().Get(`posixAccount`).AllAttributes()
	for i := 0; i < ats.Len(); i++ {
		at := ats.Index(i)
		fmt.Println(at.OID())
	}
	// Output: description
	// gecos
	// loginShell
	// userPassword
	// objectClass
	// cn
	// gidNumber
	// homeDirectory
	// uid
	// uidNumber
}

/*
This example demonstrates the means of gathering a list of all possible
[AttributeType] instances -- by OID -- that are considered OPTIONAL per
an [ObjectClass] instance.
*/
func ExampleObjectClass_AllMay() {
	ats := mySchema.ObjectClasses().Get(`posixAccount`).AllMay()
	for i := 0; i < ats.Len(); i++ {
		at := ats.Index(i)
		fmt.Println(at.OID())
	}
	// Output: description
	// gecos
	// loginShell
	// userPassword
}

/*
This example demonstrates the means of gathering a list of all possible
[AttributeType] instances -- by OID -- that are considered OPTIONAL per
an [ObjectClass] instance.
*/
func ExampleObjectClass_AllMust() {
	ats := mySchema.ObjectClasses().Get(`posixAccount`).AllMust()
	for i := 0; i < ats.Len(); i++ {
		at := ats.Index(i)
		fmt.Println(at.OID())
	}
	// Output: objectClass
	// cn
	// gidNumber
	// homeDirectory
	// uid
	// uidNumber
}

/*
This example demonstrates the manual (non-parsed) assembly of a new
[ObjectClass] instance.

Note: this example assumes a legitimate schema variable is defined
in place of the fictional "mySchema" var shown here for simplicity.
*/
func ExampleNewObjectClass() {
	oc := NewObjectClass() // Important! Initializes internal stacks

	// Conveniently input values in fluent form ...
	oc.SetSchema(mySchema).
		SetName(`engineeringPersonnel`).
		SetDescription(`EP-46: Engineering employee`).
		SetKind(`AUXILIARY`).
		SetNumericOID(`1.3.6.1.4.1.56521.999.12.5`).
		SetSuperClass(`account`, `organizationalPerson`).
		SetMust(`uid`).
		SetMay(`sn`, `cn`, `l`, `st`, `c`, `co`).
		SetExtension(`X-ORIGIN`, `NOWHERE`).
		SetStringer() // use default stringer

	fmt.Println(oc)
	// Output: ( 1.3.6.1.4.1.56521.999.12.5
	//     NAME 'engineeringPersonnel'
	//     DESC 'EP-46: Engineering employee'
	//     SUP ( account
	//         $ organizationalPerson )
	//     AUXILIARY
	//     MUST uid
	//     MAY ( sn
	//         $ cn
	//         $ l
	//         $ st
	//         $ c
	//         $ co )
	//     X-ORIGIN 'NOWHERE' )
}

/*
Do stupid things to make schemax panic, gain additional
coverage in the process.
*/
func TestObjectClass_codecov(t *testing.T) {
	_ = mySchema.ObjectClasses().SetStringer().Contains(``)
	mySchema.ObjectClasses().Push(rune(10))
	mySchema.ObjectClasses().IsZero()
	_ = mySchema.ObjectClasses().String()
	cim := mySchema.ObjectClasses().Get(`account`)
	mySchema.ObjectClasses().canPush()
	mySchema.ObjectClasses().canPush(``, ``, ``, ``, cim)
	mySchema.ObjectClasses().canPush(cim, cim)
	bmr := newCollection(``)
	ObjectClasses(bmr.cast()).Push(NewObjectClass().SetSchema(mySchema))
	ObjectClasses(bmr.cast()).Push(NewObjectClass().SetSchema(mySchema).SetNumericOID(`1.2.3.4.5`))
	bmr.cast().Push(NewObjectClass().SetSchema(mySchema))
	bmr.cast().Push(NewObjectClass().SetSchema(mySchema).SetNumericOID(`1.2.3.4.5`))
	var bad ObjectClass
	bmr.cast().Push(bad)

	ObjectClasses(bmr.cast()).oIDsStringerPretty(0)
	ObjectClasses(bmr.cast()).oIDsStringerStd()
	ObjectClasses(bmr.cast()).canPush()
	ObjectClasses(bmr.cast()).canPush(`things`)
	var ocs ObjectClasses
	ocs.oIDsStringerPretty(0)
	ocs.oIDsStringerStd()
	ocs.canPush(`forks`)
	ocs.Push(NewObjectClass().SetSchema(mySchema))

	ocs.Push(bad)

	ObjectClasses(bmr).Push(NewObjectClass().SetSchema(mySchema))
	ObjectClasses(bmr).Push(NewObjectClass().SetSchema(mySchema).SetNumericOID(`1.2.3.4.5`))
	ObjectClasses(bmr).Compliant()

	var def ObjectClass

	_ = def.String()
	_ = def.SetStringer()
	_ = def.Description()
	_ = def.Name()
	_ = def.Names()
	_ = def.Extensions()
	_ = def.Must()
	_ = def.May()
	_ = def.SuperClasses()
	_ = def.Schema()
	_ = def.Map()
	_ = def.Compliant()
	_ = def.macro()
	_ = def.Obsolete()

	def.setOID(`4.3.2.1`)
	var raw string = `( 2.999.6.11 NAME 'fakeApplicationProcess' SUP top STRUCTURAL MUST cn MAY ( seeAlso $ ou $ l $ description ) X-ORIGIN 'RFC4519' )`
	if err := def.Parse(raw); err != ErrNilReceiver {
		t.Errorf("%s failed: expected ErrNilReceiver, got %v", t.Name(), err)
		return
	}

	def = NewObjectClass()
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

	def.SetKind(0)
	def.SetKind(1)
	def.SetKind(2)
	def.SetKind(`structural`)
	def.SetKind(`auxiliary`)
	def.SetKind(`abstract`)
	def.SetKind(StructuralKind)
	def.SetKind(AuxiliaryKind)
	def.SetKind(AbstractKind)
	def.SetMust(mySchema.AttributeTypes().Get(`cn`))
	def.SetMust(rune(11))
	def.SetMay(mySchema.AttributeTypes().Get(`cn`))
	def.SetMay(rune(11))
	def.SetSuperClass(mySchema.ObjectClasses().Get(`top`))
	def.SetSuperClass(rune(11))

	if err := def.Parse(raw); err != nil {
		t.Errorf("%s failed: expected success, got %v", t.Name(), err)
		return
	}
	_ = def.macro()
	def.setOID(`2.5.13.2`)

	def.SetData(`fake`)
	def.SetData(nil)
	def.Data()

	var def2 ObjectClass
	_ = def2.Replace(def) // will fail

}

package schemax

import (
	"fmt"
)

func ExampleNewSubschema() {
	sch := NewSubschema()
	fmt.Printf("%T\n", sch)
	// Output: *schemax.Subschema
}

func ExampleNewApplies() {
	aps := NewApplies()
	fmt.Printf("%T is zero: %t\n", aps, aps.IsZero())
	// Output: schemax.Applies is zero: true
}

func ExampleNewSuperiorDITStructureRules() {
	sups := NewSuperiorDITStructureRules()
	fmt.Printf("%T is zero: %t\n", sups, sups.IsZero())
	// Output: schemax.SuperiorDITStructureRules is zero: true
}

func ExampleNewAuxiliaryObjectClasses() {
	auxs := NewAuxiliaryObjectClasses()
	fmt.Printf("%T is zero: %t\n", auxs, auxs.IsZero())
	// Output: schemax.AuxiliaryObjectClasses is zero: true
}

func ExampleNewSuperiorObjectClasses() {
	sups := NewSuperiorObjectClasses()
	fmt.Printf("%T is zero: %t\n", sups, sups.IsZero())
	// Output: schemax.SuperiorObjectClasses is zero: true
}

func ExampleNewRequiredAttributeTypes() {
	reqs := NewRequiredAttributeTypes()
	fmt.Printf("%T is zero: %t\n", reqs, reqs.IsZero())
	// Output: schemax.RequiredAttributeTypes is zero: true
}

func ExampleNewPermittedAttributeTypes() {
	pats := NewPermittedAttributeTypes()
	fmt.Printf("%T is zero: %t\n", pats, pats.IsZero())
	// Output: schemax.PermittedAttributeTypes is zero: true
}

func ExampleNewProhibitedAttributeTypes() {
	pats := NewProhibitedAttributeTypes()
	fmt.Printf("%T is zero: %t\n", pats, pats.IsZero())
	// Output: schemax.ProhibitedAttributeTypes is zero: true
}

func ExampleNewDITStructureRulesManifest() {
	dsrm := NewDITStructureRulesManifest()
	fmt.Printf("%T is zero: %t\n", dsrm, dsrm.IsZero())
	// Output: schemax.DITStructureRulesManifest is zero: true
}

func ExampleNewDITContentRulesManifest() {
	dcrm := NewDITContentRulesManifest()
	fmt.Printf("%T is zero: %t\n", dcrm, dcrm.IsZero())
	// Output: schemax.DITContentRulesManifest is zero: true
}

func ExampleNewAttributeTypesManifest() {
	atm := NewAttributeTypesManifest()
	fmt.Printf("%T is zero: %t", atm, atm.IsZero())
	// Output: schemax.AttributeTypesManifest is zero: true
}

func ExampleNewObjectClassesManifest() {
	ocm := NewObjectClassesManifest()
	fmt.Printf("%T is zero: %t", ocm, ocm.IsZero())
	// Output: schemax.ObjectClassesManifest is zero: true
}

func ExampleNewLDAPSyntaxesManifest() {
	lsm := NewLDAPSyntaxesManifest()
	fmt.Printf("%T is zero: %t", lsm, lsm.IsZero())
	// Output: schemax.LDAPSyntaxesManifest is zero: true
}

func ExampleNewMatchingRulesManifest() {
	mrm := NewMatchingRulesManifest()
	fmt.Printf("%T is zero: %t", mrm, mrm.IsZero())
	// Output: schemax.MatchingRulesManifest is zero: true
}

func ExampleNewMatchingRuleUsesManifest() {
	mrum := NewMatchingRuleUsesManifest()
	fmt.Printf("%T is zero: %t", mrum, mrum.IsZero())
	// Output: schemax.MatchingRuleUsesManifest is zero: true
}

func ExampleNewNameFormsManifest() {
	nfm := NewNameFormsManifest()
	fmt.Printf("%T is zero: %t", nfm, nfm.IsZero())
	// Output: schemax.NameFormsManifest is zero: true
}

func ExampleNewExtensions() {
	ext := NewExtensions()
	fmt.Printf("%T is zero: %t", ext, ext.IsZero())
	// Output: schemax.Extensions is zero: true
}

func ExampleOrdering_Equals() {
	sch := NewSubschema()
	sch.LSM = PopulateDefaultLDAPSyntaxesManifest()
	sch.MRM = PopulateDefaultMatchingRulesManifest()

	iOM := sch.MRM.Get(`integerOrderingMatch`, nil)
	ord1 := &Ordering{}
	ord2 := &Ordering{iOM}

	fmt.Printf("%T instances are equal: %t\n", ord1, ord1.Equals(ord2))
}

func ExampleSubstring_Equals() {
	sch := NewSubschema()
	sch.LSM = PopulateDefaultLDAPSyntaxesManifest()
	sch.MRM = PopulateDefaultMatchingRulesManifest()

	cISM := sch.MRM.Get(`caseIgnoreSubstringsMatch`, nil)
	ss1 := &Substring{}
	ss2 := &Substring{cISM}

	fmt.Printf("%T instances are equal: %t\n", ss1, ss1.Equals(ss2))
}

func ExampleEquality_Equals() {
	sch := NewSubschema()
	sch.LSM = PopulateDefaultLDAPSyntaxesManifest()
	sch.MRM = PopulateDefaultMatchingRulesManifest()

	cEIM := sch.MRM.Get(`caseExactIA5Match`, nil)
	eq1 := &Equality{}
	eq2 := &Equality{cEIM}

	fmt.Printf("%T instances are equal: %t\n", eq1, eq1.Equals(eq2))
}

func ExampleSubschema_MarshalAttributeType() {
	sch := NewSubschema()
	sch.LSM = PopulateDefaultLDAPSyntaxesManifest()
	sch.MRM = PopulateDefaultMatchingRulesManifest()
	sch.ATM = NewAttributeTypesManifest()

	myDef := `( 1.3.6.1.4.1.56521.999.104.17.2.1.10
		NAME ( 'jessesNewAttr' 'jesseAltName' )
		DESC 'This is an example AttributeType definition'
		OBSOLETE SYNTAX 1.3.6.1.4.1.1466.115.121.1.15
		EQUALITY caseIgnoreMatch X-ORIGIN 'Jesse Coretta' )`

	if err := sch.MarshalAttributeType(myDef); err != nil {
		panic(err)
	}
}

func ExampleSubschema_MarshalObjectClass() {
	sch := NewSubschema()
	sch.LSM = PopulateDefaultLDAPSyntaxesManifest()
	sch.MRM = PopulateDefaultMatchingRulesManifest()
	sch.ATM = PopulateDefaultAttributeTypesManifest()
	sch.OCM = NewObjectClassesManifest()

	myDef := `( 1.3.6.1.4.1.56521.999.104.17.2.2.91
                NAME ( 'jesseOC' 'specialJesseClass' )
                DESC 'This is an example ObjectClass definition'
		SUP account
		MUST ( sn $ givenName $ carLicense )
		MAY ( description $ mobile )
                X-ORIGIN 'Jesse Coretta' )`

	if err := sch.MarshalObjectClass(myDef); err != nil {
		panic(err)
	}
}

func ExampleSubschema_MarshalNameForm() {
	sch := NewSubschema()
	sch.LSM = PopulateDefaultLDAPSyntaxesManifest()
	sch.MRM = PopulateDefaultMatchingRulesManifest()
	sch.ATM = PopulateDefaultAttributeTypesManifest()
	sch.OCM = PopulateDefaultObjectClassesManifest()
	sch.NFM = NewNameFormsManifest()

	myDef := `( 1.3.6.1.4.1.56521.999.104.17.2.16.4
		NAME 'jesseNameForm'
		DESC 'this is an example Name Form definition'
		OC account
		MUST ( cn $ givenName $ carLicense )
		MAY ( description $ mobile )
		X-ORIGIN 'Jesse Coretta' )`

	if err := sch.MarshalNameForm(myDef); err != nil {
		panic(err)
	}
}

func ExampleSubschema_MarshalDITStructureRule() {
	sch := NewSubschema()
	sch.LSM = PopulateDefaultLDAPSyntaxesManifest()
	sch.MRM = PopulateDefaultMatchingRulesManifest()
	sch.ATM = PopulateDefaultAttributeTypesManifest()
	sch.OCM = PopulateDefaultObjectClassesManifest()
	sch.NFM = NewNameFormsManifest()
	sch.DSRM = NewDITStructureRulesManifest()

	nf := `( 1.3.6.1.4.1.56521.999.104.17.2.16.4
                NAME 'jesseNameForm'
		DESC 'this is an example Name Form definition'
                OC account
                MUST ( cn $ givenName $ carLicense )
                MAY ( description $ mobile )
                X-ORIGIN 'Jesse Coretta' )`

	if err := sch.MarshalNameForm(nf); err != nil {
		panic(err)
	}

	dsr := `( 0 
		NAME 'jesseDSR'
		DESC 'this is an example DIT Structure Rule'
		FORM jesseNameForm
		X-ORIGIN 'Jesse Coretta' )`

	if err := sch.MarshalDITStructureRule(dsr); err != nil {
		panic(err)
	}
}

func ExampleSubschema_MarshalDITContentRule() {
	sch := NewSubschema()
	sch.LSM = PopulateDefaultLDAPSyntaxesManifest()
	sch.MRM = PopulateDefaultMatchingRulesManifest()
	sch.ATM = PopulateDefaultAttributeTypesManifest()
	sch.OCM = PopulateDefaultObjectClassesManifest()
	sch.DCRM = NewDITContentRulesManifest()

	dcr := `( 0.9.2342.19200300.100.4.5
		NAME 'accountContentRule'
		AUX ( posixAccount $ shadowAccount )
		MUST uid
		MAY description
		NOT gecos )`

	if err := sch.MarshalDITContentRule(dcr); err != nil {
		panic(err)
	}
}

func ExampleSubschema_MarshalLDAPSyntax() {
	sch := NewSubschema()
	sch.LSM = NewLDAPSyntaxesManifest()

	ls := `( 1.3.6.1.4.1.1466.115.121.1.2
		DESC 'Access Point'
		X-ORIGIN 'RFC4517' )`

	if err := sch.MarshalLDAPSyntax(ls); err != nil {
		panic(err)
	}
}

func ExampleSubschema_MarshalMatchingRule() {
	sch := NewSubschema()
	sch.LSM = PopulateDefaultLDAPSyntaxesManifest()
	sch.MRM = NewMatchingRulesManifest()

	mr := `( 1.3.6.1.4.1.1466.109.114.3
		NAME 'caseIgnoreIA5SubstringsMatch'
		SYNTAX 1.3.6.1.4.1.1466.115.121.1.58
		X-ORIGIN 'RFC4517' )`

	if err := sch.MarshalMatchingRule(mr); err != nil {
		panic(err)
	}
}

func ExampleUnmarshal() {
	sch := NewSubschema()
	sch.LSM = PopulateDefaultLDAPSyntaxesManifest()
	sch.MRM = PopulateDefaultMatchingRulesManifest()
	sch.ATM = PopulateDefaultAttributeTypesManifest()

	name := sch.ATM.Get(`name`, nil)
	def, err := Unmarshal(name)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", def)
	// Output: ( 2.5.4.41 NAME 'name' SYNTAX 1.3.6.1.4.1.1466.115.121.1.15 EQUALITY caseIgnoreMatch SUBSTR caseIgnoreSubstringsMatch X-ORIGIN 'RFC4519' )
}

func ExampleAttributeTypesManifest_Get() {
	sch := NewSubschema()
	sch.LSM = PopulateDefaultLDAPSyntaxesManifest()
	sch.MRM = PopulateDefaultMatchingRulesManifest()
	sch.ATM = PopulateDefaultAttributeTypesManifest()

	if attr := sch.ATM.Get(`name`, nil); !attr.IsZero() {
		fmt.Printf("Name OID: %s\n", attr.OID)
	}
	// Output: Name OID: 2.5.4.41
}

func ExampleObjectClassesManifest_Get() {
	sch := NewSubschema()
	sch.LSM = PopulateDefaultLDAPSyntaxesManifest()
	sch.MRM = PopulateDefaultMatchingRulesManifest()
	sch.ATM = PopulateDefaultAttributeTypesManifest()
	sch.OCM = PopulateDefaultObjectClassesManifest()

	if class := sch.OCM.Get(`account`, nil); !class.IsZero() {
		fmt.Printf("account OID: %s\n", class.OID)
	}
	// Output: account OID: 0.9.2342.19200300.100.4.5
}

func ExampleLDAPSyntaxesManifest_Get() {
	sch := NewSubschema()
	sch.LSM = PopulateDefaultLDAPSyntaxesManifest()

	if ls := sch.LSM.Get(`1.3.6.1.4.1.1466.115.121.1.4`, nil); !ls.IsZero() {
		fmt.Printf("%T OID: %s\n", ls, ls.OID)
	}
	// Output: *schemax.LDAPSyntax OID: 1.3.6.1.4.1.1466.115.121.1.4
}

func ExampleAttributeTypesManifest_IsZero() {
	sch := NewSubschema()
	sch.LSM = PopulateDefaultLDAPSyntaxesManifest()
	sch.MRM = PopulateDefaultMatchingRulesManifest()
	sch.ATM = PopulateDefaultAttributeTypesManifest()

	isZero := sch.ATM.IsZero()
	fmt.Printf("%T is zero: %t\n", sch.ATM, isZero)
	// Output: schemax.AttributeTypesManifest is zero: false
}

func ExampleObjectClassesManifest_IsZero() {
	sch := NewSubschema()
	sch.LSM = PopulateDefaultLDAPSyntaxesManifest()
	sch.MRM = PopulateDefaultMatchingRulesManifest()
	sch.ATM = PopulateDefaultAttributeTypesManifest()
	sch.OCM = PopulateDefaultObjectClassesManifest()

	isZero := sch.OCM.IsZero()
	fmt.Printf("%T is zero: %t\n", sch.OCM, isZero)
	// Output: schemax.ObjectClassesManifest is zero: false
}

func ExampleName_IsZero() {
	n := &Name{}
	fmt.Printf("%T is zero: %t\n", n, n.IsZero())
	// Output *schemax.Name is zero: true
}

func ExampleAttributeType_IsZero() {
	at := &AttributeType{}
	fmt.Printf("%T is zero: %t\n", at, at.IsZero())
	// Output *schemax.AttributeType is zero: true
}

func ExampleObjectClass_IsZero() {
	oc := &ObjectClass{}
	fmt.Printf("%T is zero: %t\n", oc, oc.IsZero())
	// Output *schemax.ObjectClass is zero: true
}

func ExampleLDAPSyntax_IsZero() {
	ls := &LDAPSyntax{}
	fmt.Printf("%T is zero: %t\n", ls, ls.IsZero())
	// Output *schemax.LDAPSyntax is zero: true
}

func ExampleMatchingRule_IsZero() {
	mr := &MatchingRule{}
	fmt.Printf("%T is zero: %t\n", mr, mr.IsZero())
	// Output *schemax.MatchingRule is zero: true
}

func ExampleMatchingRuleUse_IsZero() {
	mru := &MatchingRuleUse{}
	fmt.Printf("%T is zero: %t\n", mru, mru.IsZero())
	// Output *schemax.MatchingRuleUse is zero: true
}

func ExampleNameForm_IsZero() {
	nf := &NameForm{}
	fmt.Printf("%T is zero: %t\n", nf, nf.IsZero())
	// Output *schemax.NameForm is zero: true
}

func ExampleDITStructureRule_IsZero() {
	dsr := &DITStructureRule{}
	fmt.Printf("%T is zero: %t\n", dsr, dsr.IsZero())
	// Output *schemax.DITStructureRule is zero: true
}

func ExampleDITContentRule_IsZero() {
	dcr := &DITContentRule{}
	fmt.Printf("%T is zero: %t\n", dcr, dcr.IsZero())
	// Output *schemax.DITContentRule is zero: true
}

func ExampleApplies_IsZero() {
	aps := &Applies{}
	fmt.Printf("%T is zero: %t\n", aps, aps.IsZero())
	// Output: *schemax.Applies is zero: true
}

func ExampleProhibitedAttributeTypes_IsZero() {
	aps := &ProhibitedAttributeTypes{}
	fmt.Printf("%T is zero: %t\n", aps, aps.IsZero())
	// Output: *schemax.ProhibitedAttributeTypes is zero: true
}

func ExampleRequiredAttributeTypes_IsZero() {
	aps := &RequiredAttributeTypes{}
	fmt.Printf("%T is zero: %t\n", aps, aps.IsZero())
	// Output: *schemax.RequiredAttributeTypes is zero: true
}

func ExamplePermittedAttributeTypes_IsZero() {
	aps := &PermittedAttributeTypes{}
	fmt.Printf("%T is zero: %t\n", aps, aps.IsZero())
	// Output: *schemax.PermittedAttributeTypes is zero: true
}

func ExampleDITStructureRule_Equals() {
	dsr1 := &DITStructureRule{}
	dsr2 := &DITStructureRule{ID: NewRuleID(0)}
	fmt.Printf("%T instances are equal: %t\n", dsr1, dsr1.Equals(dsr2))
	// Output: *schemax.DITStructureRule instances are equal: true
}

func ExampleLDAPSyntax_Equals() {
	sch := NewSubschema()
	sch.LSM = PopulateDefaultLDAPSyntaxesManifest()

	syn1 := &LDAPSyntax{}
	syn2 := sch.LSM.Get(`1.3.6.1.4.1.1466.115.121.1.4`, nil) // Audio per RFC4517
	fmt.Printf("%T instances are equal: %t\n", syn1, syn1.Equals(syn2))
	// Output: *schemax.LDAPSyntax instances are equal: false
}

func ExampleAuxiliaryObjectClasses_IsZero() {
	aux := &AuxiliaryObjectClasses{}
	fmt.Printf("%T is zero: %t\n", aux, aux.IsZero())
	// Output: *schemax.AuxiliaryObjectClasses is zero: true
}

func ExampleSuperiorObjectClasses_IsZero() {
	sup := &SuperiorObjectClasses{}
	fmt.Printf("%T is zero: %t\n", sup, sup.IsZero())
	// Output: *schemax.SuperiorObjectClasses is zero: true
}

func ExampleSuperiorDITStructureRules_IsZero() {
	sups := &SuperiorDITStructureRules{}
	fmt.Printf("%T is zero: %t\n", sups, sups.IsZero())
	// Output: *schemax.SuperiorDITStructureRules is zero: true
}

func ExampleLDAPSyntax_IsHumanReadable() {
	sch := NewSubschema()
	sch.LSM = PopulateDefaultLDAPSyntaxesManifest()

	audio := sch.LSM.Get(`1.3.6.1.4.1.1466.115.121.1.4`, nil) // Audio per RFC4517
	fmt.Printf("%T (%s) is Human Readable: %t\n", audio, audio.OID, audio.IsHumanReadable())
	// Output: *schemax.LDAPSyntax (1.3.6.1.4.1.1466.115.121.1.4) is Human Readable: false
}

func ExampleAuxiliaryObjectClasses_Equals() {
	sch := NewSubschema()
	sch.LSM = PopulateDefaultLDAPSyntaxesManifest()
	sch.MRM = PopulateDefaultMatchingRulesManifest()
	sch.ATM = PopulateDefaultAttributeTypesManifest()
	sch.OCM = PopulateDefaultObjectClassesManifest()

	aux1 := &AuxiliaryObjectClasses{sch.OCM.Get(`posixAccount`, nil)}
	aux2 := &AuxiliaryObjectClasses{}
	fmt.Printf("%T instances are equal: %t\n", aux1, aux1.Equals(aux2))
	// Output *schemax.AuxiliaryObjectClasses instances are equal: false
}

func ExamplePermittedAttributeTypes_Equals() {
	sch := NewSubschema()
	sch.LSM = PopulateDefaultLDAPSyntaxesManifest()
	sch.MRM = PopulateDefaultMatchingRulesManifest()
	sch.ATM = PopulateDefaultAttributeTypesManifest()

	may1 := &PermittedAttributeTypes{sch.ATM.Get(`givenName`, nil)}
	may2 := &PermittedAttributeTypes{}
	fmt.Printf("%T instances are equal: %t\n", may1, may1.Equals(may2))
	// Output *schemax.PermittedAttributeTypes instances are equal: false
}

func ExampleRequiredAttributeTypes_Equals() {
	sch := NewSubschema()
	sch.LSM = PopulateDefaultLDAPSyntaxesManifest()
	sch.MRM = PopulateDefaultMatchingRulesManifest()
	sch.ATM = PopulateDefaultAttributeTypesManifest()

	must1 := &RequiredAttributeTypes{sch.ATM.Get(`cn`, nil)}
	must2 := &RequiredAttributeTypes{}
	fmt.Printf("%T instances are equal: %t\n", must1, must1.Equals(must2))
	// Output *schemax.RequiredAttributeTypes instances are equal: false
}

func ExampleProhibitedAttributeTypes_Equals() {
	sch := NewSubschema()
	sch.LSM = PopulateDefaultLDAPSyntaxesManifest()
	sch.MRM = PopulateDefaultMatchingRulesManifest()
	sch.ATM = PopulateDefaultAttributeTypesManifest()

	not1 := &ProhibitedAttributeTypes{sch.ATM.Get(`sn`, nil)}
	not2 := &ProhibitedAttributeTypes{}
	fmt.Printf("%T instances are equal: %t\n", not1, not1.Equals(not2))
	// Output *schemax.ProhibitedAttributeTypes instances are equal: false
}

func ExampleProhibitedAttributeTypes_Index() {
	sch := NewSubschema()
	sch.LSM = PopulateDefaultLDAPSyntaxesManifest()
	sch.MRM = PopulateDefaultMatchingRulesManifest()
	sch.ATM = PopulateDefaultAttributeTypesManifest()

	attr := sch.ATM.Get(`cn`, nil)
	not := &ProhibitedAttributeTypes{attr}

	idx, exists := not.Index(`2.5.4.3`)
	if exists {
		fmt.Printf("Index for %s attr in %T: %d\n", attr.Name.Value(0), not, idx)
	}
	// Output: Index for cn attr in *schemax.ProhibitedAttributeTypes: 0
}

func ExampleAuxiliaryObjectClasses_Index() {
	sch := NewSubschema()
	sch.LSM = PopulateDefaultLDAPSyntaxesManifest()
	sch.MRM = PopulateDefaultMatchingRulesManifest()
	sch.ATM = PopulateDefaultAttributeTypesManifest()
	sch.OCM = PopulateDefaultObjectClassesManifest()

	oc := sch.OCM.Get(`posixAccount`, nil)
	aux := &AuxiliaryObjectClasses{oc}

	idx, exists := aux.Index(`1.3.6.1.1.1.2.0`)
	if exists {
		fmt.Printf("Index for %s class in %T: %d\n", oc.Name.Value(0), aux, idx)
	}
	// Output: Index for posixAccount class in *schemax.AuxiliaryObjectClasses: 0
}

func ExampleRequiredAttributeTypes_Index() {
	sch := NewSubschema()
	sch.LSM = PopulateDefaultLDAPSyntaxesManifest()
	sch.MRM = PopulateDefaultMatchingRulesManifest()
	sch.ATM = PopulateDefaultAttributeTypesManifest()

	attr := sch.ATM.Get(`cn`, nil)
	must := &RequiredAttributeTypes{attr}

	idx, exists := must.Index(`2.5.4.3`)
	if exists {
		fmt.Printf("Index for %s attr in %T: %d\n", attr.Name.Value(0), must, idx)
	}
	// Output: Index for cn attr in *schemax.RequiredAttributeTypes: 0
}

func ExamplePermittedAttributeTypes_Index() {
	sch := NewSubschema()
	sch.LSM = PopulateDefaultLDAPSyntaxesManifest()
	sch.MRM = PopulateDefaultMatchingRulesManifest()
	sch.ATM = PopulateDefaultAttributeTypesManifest()

	attr := sch.ATM.Get(`cn`, nil)
	may := &PermittedAttributeTypes{attr}

	idx, exists := may.Index(`2.5.4.3`)
	if exists {
		fmt.Printf("Index for %s attr in %T: %d\n", attr.Name.Value(0), may, idx)
	}
	// Output: Index for cn attr in *schemax.PermittedAttributeTypes: 0
}

func ExampleApplies_Index() {
	sch := NewSubschema()
	sch.LSM = PopulateDefaultLDAPSyntaxesManifest()
	sch.MRM = PopulateDefaultMatchingRulesManifest()
	sch.ATM = PopulateDefaultAttributeTypesManifest()

	attr := sch.ATM.Get(`cn`, nil)
	applies := &Applies{attr}

	idx, exists := applies.Index(`2.5.4.3`)
	if exists {
		fmt.Printf("Index for %s attr in %T: %d\n", attr.Name.Value(0), applies, idx)
	}
	// Output: Index for cn attr in *schemax.Applies: 0
}

func ExampleAttributeType_SetMaxLength() {
	sch := NewSubschema()
	sch.LSM = PopulateDefaultLDAPSyntaxesManifest()

	at := &AttributeType{
		Name:   NewName(`fakeName`),
		OID:    OID(`1.3.6.1.4.1.56521.999.104.17.2.1.55`),
		Syntax: sch.LSM.Get(`1.3.6.1.4.1.1466.115.121.1.15`, nil),
	}
	at.SetMaxLength(128)
	fmt.Printf("Max allowed length for %T value: %d\n", at, at.MaxLength())
	// Output: Max allowed length for *schemax.AttributeType value: 128
}

func ExampleAttributeType_MaxLength() {
	sch := NewSubschema()
	sch.LSM = PopulateDefaultLDAPSyntaxesManifest()

	at := &AttributeType{
		Name:   NewName(`fakeName`),
		OID:    OID(`1.3.6.1.4.1.56521.999.104.17.2.1.55`),
		Syntax: sch.LSM.Get(`1.3.6.1.4.1.1466.115.121.1.15`, nil),
	}
	at.SetMaxLength(128)
	fmt.Printf("Max allowed length for %T value: %d\n", at, at.MaxLength())
	// Output: Max allowed length for *schemax.AttributeType value: 128
}

func ExampleAttributeType_Equals() {
	sch := NewSubschema()
	sch.LSM = PopulateDefaultLDAPSyntaxesManifest()

	at1 := &AttributeType{
		Name:   NewName(`fakeName`),
		OID:    OID(`1.3.6.1.4.1.56521.999.104.17.2.1.55`),
		Syntax: sch.LSM.Get(`1.3.6.1.4.1.1466.115.121.1.15`, nil),
	}

	at2 := &AttributeType{
		Name:   NewName(`fakeName`),
		OID:    OID(`1.3.6.1.4.1.56521.999.104.17.2.1.55`),
		Syntax: sch.LSM.Get(`1.3.6.1.4.1.1466.115.121.1.15`, nil),
	}

	fmt.Printf("at1 is equal to at2: %t\n", at1.Equals(at2))
	// Output: at1 is equal to at2: true
}

func ExampleObjectClass_Equals() {
	oc1 := &ObjectClass{
		Name: NewName(`fakeName`),
		OID:  OID(`1.3.6.1.4.1.56521.999.104.17.2.2.84`),
		Kind: Auxiliary,
	}

	oc2 := &ObjectClass{
		Name: NewName(`fakeName`),
		OID:  OID(`1.3.6.1.4.1.56521.999.104.17.2.2.81`),
		Kind: Structural,
	}

	fmt.Printf("%T instances are equal: %t\n", oc1, oc1.Equals(oc2))
	// Output: *schemax.ObjectClass instances are equal: false
}

func ExampleLDAPSyntaxesManifest_Exists() {
	sch := NewSubschema()
	sch.LSM = PopulateDefaultLDAPSyntaxesManifest()

	exists := sch.LSM.Exists(`1.3.6.1.4.1.1466.115.121.1.4`, nil)
	fmt.Printf("syntax exists: %t\n", exists)
	// Output: syntax exists: true
}

func ExampleObjectClassesManifest_Exists() {
	sch := NewSubschema()
	sch.LSM = PopulateDefaultLDAPSyntaxesManifest()
	sch.MRM = PopulateDefaultMatchingRulesManifest()
	sch.ATM = PopulateDefaultAttributeTypesManifest()
	sch.OCM = PopulateDefaultObjectClassesManifest()

	exists := sch.OCM.Exists(`account`, nil)
	fmt.Printf("account ObjectClass exists: %t\n", exists)
	// Output: account ObjectClass exists: true
}

func ExampleAttributeTypesManifest_Set() {

	sch := NewSubschema()
	sch.LSM = PopulateDefaultLDAPSyntaxesManifest()
	sch.MRM = PopulateDefaultMatchingRulesManifest()
	sch.ATM = PopulateDefaultAttributeTypesManifest()

	sch.ATM.Set(&AttributeType{
		Name:   NewName(`fakeName`),
		OID:    OID(`1.3.6.1.4.1.56521.999.104.17.2.1.55`),
		Syntax: sch.LSM.Get(`1.3.6.1.4.1.1466.115.121.1.15`, nil),
	})

	if get := sch.ATM.Get(`fakeName`, nil); !get.IsZero() {
		fmt.Printf("fakeName OID: %s\n", get.OID)
	}
	// Output: fakeName OID: 1.3.6.1.4.1.56521.999.104.17.2.1.55
}

func ExampleAttributeTypesManifest_Exists() {

	sch := NewSubschema()
	sch.LSM = PopulateDefaultLDAPSyntaxesManifest()
	sch.MRM = PopulateDefaultMatchingRulesManifest()
	sch.ATM = PopulateDefaultAttributeTypesManifest()

	exists := sch.ATM.Exists(`name`, nil)
	fmt.Printf("name AttributeType exists: %t\n", exists)
	// Output: name AttributeType exists: true
}

func ExampleAttributeTypesManifest_Equals() {

	sch := NewSubschema()
	sch.LSM = PopulateDefaultLDAPSyntaxesManifest()
	sch.MRM = PopulateDefaultMatchingRulesManifest()
	sch.ATM = PopulateDefaultAttributeTypesManifest()

	otherATM := NewAttributeTypesManifest()
	same := sch.ATM.Equals(otherATM)
	fmt.Printf("%T instances are equal: %t\n", sch.ATM, same)
	// Output: schemax.AttributeTypesManifest instances are equal: false
}

func ExampleObjectClassesManifest_Equals() {

	sch := NewSubschema()
	sch.LSM = PopulateDefaultLDAPSyntaxesManifest()
	sch.MRM = PopulateDefaultMatchingRulesManifest()
	sch.ATM = PopulateDefaultAttributeTypesManifest()
	sch.OCM = PopulateDefaultObjectClassesManifest()

	otherOCM := NewObjectClassesManifest()
	same := sch.OCM.Equals(otherOCM)
	fmt.Printf("%T instances are equal: %t\n", sch.OCM, same)
	// Output: schemax.ObjectClassesManifest instances are equal: false
}

func ExampleLDAPSyntaxesManifest_Equals() {

	sch := NewSubschema()
	sch.LSM = PopulateDefaultLDAPSyntaxesManifest()

	otherLSM := NewLDAPSyntaxesManifest()
	same := sch.LSM.Equals(otherLSM)
	fmt.Printf("%T instances are equal: %t\n", sch.LSM, same)
	// Output: schemax.LDAPSyntaxesManifest instances are equal: false
}

func ExampleMatchingRulesManifest_Equals() {

	sch := NewSubschema()
	sch.LSM = PopulateDefaultLDAPSyntaxesManifest()
	sch.MRM = PopulateDefaultMatchingRulesManifest()

	otherMRM := NewMatchingRulesManifest()
	same := sch.MRM.Equals(otherMRM)
	fmt.Printf("%T instances are equal: %t\n", sch.MRM, same)
	// Output: schemax.MatchingRulesManifest instances are equal: false
}

func ExampleDITContentRule_Equals() {

	sch := NewSubschema()
	sch.LSM = PopulateDefaultLDAPSyntaxesManifest()
	sch.MRM = PopulateDefaultMatchingRulesManifest()
	sch.ATM = PopulateDefaultAttributeTypesManifest()
	sch.OCM = PopulateDefaultObjectClassesManifest()
	sch.DCRM = NewDITContentRulesManifest()

	posixAcct := sch.OCM.Get(`posixAccount`, nil)
	shadowAcct := sch.OCM.Get(`posixAccount`, nil)
	uidAttr := sch.ATM.Get(`uid`, nil)
	descAttr := sch.ATM.Get(`description`, nil)
	givenNameAttr := sch.ATM.Get(`givenName`, nil)

	if posixAcct.IsZero() || shadowAcct.IsZero() {
		panic("missing key ObjectClasse instances")
	}

	if uidAttr.IsZero() || descAttr.IsZero() || givenNameAttr.IsZero() {
		panic("missing key AttributeType instances")
	}

	dcr1 := &DITContentRule{
		OID:  OID(`0.9.2342.19200300.100.4.5`),
		Name: NewName(`jesseAccountContentRule`),
		Aux:  AuxiliaryObjectClasses{posixAcct, shadowAcct},
		Must: RequiredAttributeTypes{uidAttr},
		May:  PermittedAttributeTypes{descAttr},
		Not:  ProhibitedAttributeTypes{givenNameAttr},
	}

	dcr2 := &DITContentRule{
		OID:  OID(`0.9.2342.19200300.100.4.5`),
		Name: NewName(`jesseAccountContentRule2`),
		Aux:  AuxiliaryObjectClasses{posixAcct, shadowAcct},
		Must: RequiredAttributeTypes{uidAttr},
		May:  PermittedAttributeTypes{descAttr},
	}

	same := dcr1.Equals(dcr2)
	fmt.Printf("The two %T instances are equal: %t\n", dcr1, same)
	// Output: The two *schemax.DITContentRule instances are equal: false

}

func ExampleDITContentRulesManifest_Equals() {

	dcrm1 := NewDITContentRulesManifest()
	dcrm2 := NewDITContentRulesManifest()

	fmt.Printf("%T instances are equal: %t\n", dcrm1, dcrm1.Equals(dcrm2))
	// Output: schemax.DITContentRulesManifest instances are equal: true
}

func ExampleDITContentRulesManifest_Get() {

	sch := NewSubschema()
	sch.LSM = PopulateDefaultLDAPSyntaxesManifest()
	sch.MRM = PopulateDefaultMatchingRulesManifest()
	sch.ATM = PopulateDefaultAttributeTypesManifest()
	sch.OCM = PopulateDefaultObjectClassesManifest()
	sch.DCRM = NewDITContentRulesManifest()

	posixAcct := sch.OCM.Get(`posixAccount`, nil)
	shadowAcct := sch.OCM.Get(`posixAccount`, nil)
	uidAttr := sch.ATM.Get(`uid`, nil)
	descAttr := sch.ATM.Get(`description`, nil)
	givenNameAttr := sch.ATM.Get(`givenName`, nil)

	if posixAcct.IsZero() || shadowAcct.IsZero() {
		panic("missing key ObjectClasse instances")
	}

	if uidAttr.IsZero() || descAttr.IsZero() || givenNameAttr.IsZero() {
		panic("missing key AttributeType instances")
	}

	sch.DCRM.Set(&DITContentRule{
		OID:  OID(`0.9.2342.19200300.100.4.5`),
		Name: NewName(`jesseAccountContentRule`),
		Aux:  AuxiliaryObjectClasses{posixAcct, shadowAcct},
		Must: RequiredAttributeTypes{uidAttr},
		May:  PermittedAttributeTypes{descAttr},
		Not:  ProhibitedAttributeTypes{givenNameAttr},
	})

	dcr := sch.DCRM.Get(`jesseAccountContentRule`, nil)
	fmt.Printf("%T.Name: %s\n", *dcr, dcr.Name.Value(0))
	// Output: schemax.DITContentRule.Name: jesseAccountContentRule

}

func ExampleDITContentRulesManifest_Set() {
	sch := NewSubschema()
	sch.LSM = PopulateDefaultLDAPSyntaxesManifest()
	sch.MRM = PopulateDefaultMatchingRulesManifest()
	sch.ATM = PopulateDefaultAttributeTypesManifest()
	sch.OCM = PopulateDefaultObjectClassesManifest()
	sch.DCRM = NewDITContentRulesManifest()

	posixAcct := sch.OCM.Get(`posixAccount`, nil)
	shadowAcct := sch.OCM.Get(`posixAccount`, nil)
	uidAttr := sch.ATM.Get(`uid`, nil)
	descAttr := sch.ATM.Get(`description`, nil)
	givenNameAttr := sch.ATM.Get(`givenName`, nil)

	if posixAcct.IsZero() || shadowAcct.IsZero() {
		panic("missing key ObjectClass instances")
	}

	if uidAttr.IsZero() || descAttr.IsZero() || givenNameAttr.IsZero() {
		panic("missing key AttributeType instances")
	}

	sch.DCRM.Set(&DITContentRule{
		OID:  OID(`0.9.2342.19200300.100.4.5`),
		Name: NewName(`jesseAccountContentRule`),
		Aux:  AuxiliaryObjectClasses{posixAcct, shadowAcct},
		Must: RequiredAttributeTypes{uidAttr},
		May:  PermittedAttributeTypes{descAttr},
		Not:  ProhibitedAttributeTypes{givenNameAttr},
	})

	dcr := sch.DCRM.Get(`jesseAccountContentRule`, nil)
	fmt.Printf("%T.Name: %s\n", *dcr, dcr.Name.Value(0))
	// Output: schemax.DITContentRule.Name: jesseAccountContentRule
}

func ExampleDITContentRulesManifest_Exists() {
	sch := NewSubschema()
	sch.DCRM = NewDITContentRulesManifest()

	dcr := &DITContentRule{}

	exists := sch.DCRM.Exists(dcr, nil)
	fmt.Printf("%T exists within %T: %t\n", dcr, sch.DCRM, exists)
	// Output: *schemax.DITContentRule exists within schemax.DITContentRulesManifest: false
}

func ExamplePopulateDefaultAttributeTypesManifest() {
	atm := NewAttributeTypesManifest()
	fmt.Printf("%T is zero: %t", atm, atm.IsZero())
	// Output: schemax.AttributeTypesManifest is zero: true
}

func ExampleApplies_Label() {
	aps := Applies{}
	lab := aps.Label()
	fmt.Printf("%T label: %s (len:%d)\n", aps, lab, len(lab))
	// Output: schemax.Applies label: APPLIES (len:7)
}

func ExampleAuxiliaryObjectClasses_Label() {
	aux := AuxiliaryObjectClasses{}
	lab := aux.Label()
	fmt.Printf("%T label: %s (len:%d)\n", aux, lab, len(lab))
	// Output: schemax.AuxiliaryObjectClasses label: AUX (len:3)
}

func ExamplePermittedAttributeTypes_Label() {
	may := PermittedAttributeTypes{}
	lab := may.Label()
	fmt.Printf("%T label: %s (len:%d)\n", may, lab, len(lab))
	// Output: schemax.PermittedAttributeTypes label: MAY (len:3)
}

func ExampleRequiredAttributeTypes_Label() {
	must := RequiredAttributeTypes{}
	lab := must.Label()
	fmt.Printf("%T label: %s (len:%d)\n", must, lab, len(lab))
	// Output: schemax.RequiredAttributeTypes label: MUST (len:4)
}

func ExampleProhibitedAttributeTypes_Label() {
	not := ProhibitedAttributeTypes{}
	lab := not.Label()
	fmt.Printf("%T label: %s (len:%d)\n", not, lab, len(lab))
	// Output: schemax.ProhibitedAttributeTypes label: NOT (len:3)
}

func ExampleUsage_Label() {
	usage := DSAOperation
	lab := usage.Label()
	fmt.Printf("%T label: %s (len:%d)\n", usage, lab, len(lab))
	// Output: schemax.Usage label: USAGE (len:5)
}

func ExampleKind_Label() {
	// Remember! Some labels are invisible, like
	// this one ...
	kind := Auxiliary
	lab := kind.Label()
	fmt.Printf("%T label: %s (len:%d)\n", kind, lab, len(lab))
	// Output: schemax.Kind label:  (len:0)
}

func ExampleLDAPSyntax_Label() {
	ls := LDAPSyntax{}
	lab := ls.Label()
	fmt.Printf("%T label: %s (len:%d)\n", ls, lab, len(lab))
	// Output: schemax.LDAPSyntax label: SYNTAX (len:6)
}

func ExampleEquality_Label() {
	eq := Equality{}
	lab := eq.Label()
	fmt.Printf("%T label: %s (len:%d)\n", eq, lab, len(lab))
	// Output: schemax.Equality label: EQUALITY (len:8)
}

func ExampleSubstring_Label() {
	ss := Substring{}
	lab := ss.Label()
	fmt.Printf("%T label: %s (len:%d)\n", ss, lab, len(lab))
	// Output: schemax.Substring label: SUBSTR (len:6)
}

func ExampleOrdering_Label() {
	ord := Ordering{}
	lab := ord.Label()
	fmt.Printf("%T label: %s (len:%d)\n", ord, lab, len(lab))
	// Output: schemax.Ordering label: ORDERING (len:8)
}

func ExampleSuperiorAttributeType_Label() {
	sup := SuperiorAttributeType{}
	lab := sup.Label()
	fmt.Printf("%T label: %s (len:%d)\n", sup, lab, len(lab))
	// Output: schemax.SuperiorAttributeType label: SUP (len:3)
}

func ExampleSuperiorObjectClasses_Label() {
	sups := SuperiorObjectClasses{}
	lab := sups.Label()
	fmt.Printf("%T label: %s (len:%d)\n", sups, lab, len(lab))
	// Output: schemax.SuperiorObjectClasses label: SUP (len:3)
}

func ExampleSuperiorDITStructureRules_Label() {
	sups := SuperiorDITStructureRules{}
	lab := sups.Label()
	fmt.Printf("%T label: %s (len:%d)\n", sups, lab, len(lab))
	// Output: schemax.SuperiorDITStructureRules label: SUP (len:3)
}

func ExampleExtensions_Label() {
	// Remember! Some labels are invisible, like
	// this one ...
	ext := Extensions{}
	lab := ext.Label()
	fmt.Printf("%T label: %s (len:%d)\n", ext, lab, len(lab))
	// Output: schemax.Extensions label:  (len:0)
}

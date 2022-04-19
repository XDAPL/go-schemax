package schemax

import (
	"fmt"
)

func ExampleNewSubschema() {
	sch := NewSubschema()
	fmt.Printf("%T\n", sch)
	// Output: *schemax.Subschema
}

func ExampleNewApplicableAttributeTypes() {
	aps := NewApplicableAttributeTypes()
	fmt.Printf("%T is zero: %t\n", aps, aps.IsZero())
	// Output: *schemax.ApplicableAttributeTypes is zero: true
}

func ExampleNewDITStructureRules() {
	dsrs := NewDITStructureRules()
	fmt.Printf("%T is zero: %t\n", dsrs, dsrs.IsZero())
	// Output: *schemax.DITStructureRules is zero: true
}

func ExampleNewSuperiorDITStructureRules() {
	sups := NewSuperiorDITStructureRules()
	fmt.Printf("%T is zero: %t\n", sups, sups.IsZero())
	// Output: *schemax.SuperiorDITStructureRules is zero: true
}

func ExampleNewAuxiliaryObjectClasses() {
	auxs := NewAuxiliaryObjectClasses()
	fmt.Printf("%T is zero: %t\n", auxs, auxs.IsZero())
	// Output: *schemax.AuxiliaryObjectClasses is zero: true
}

func ExampleNewSuperiorObjectClasses() {
	sups := NewSuperiorObjectClasses()
	fmt.Printf("%T is zero: %t\n", sups, sups.IsZero())
	// Output: *schemax.SuperiorObjectClasses is zero: true
}

func ExampleNewRequiredAttributeTypes() {
	reqs := NewRequiredAttributeTypes()
	fmt.Printf("%T is zero: %t\n", reqs, reqs.IsZero())
	// Output: *schemax.RequiredAttributeTypes is zero: true
}

func ExampleNewPermittedAttributeTypes() {
	pats := NewPermittedAttributeTypes()
	fmt.Printf("%T is zero: %t\n", pats, pats.IsZero())
	// Output: *schemax.PermittedAttributeTypes is zero: true
}

func ExampleNewProhibitedAttributeTypes() {
	pats := NewProhibitedAttributeTypes()
	fmt.Printf("%T is zero: %t\n", pats, pats.IsZero())
	// Output: *schemax.ProhibitedAttributeTypes is zero: true
}

func ExampleNewDITContentRules() {
	dcrm := NewDITContentRules()
	fmt.Printf("%T is zero: %t\n", dcrm, dcrm.IsZero())
	// Output: *schemax.DITContentRules is zero: true
}

func ExampleNewAttributeTypes() {
	atm := NewAttributeTypes()
	fmt.Printf("%T is zero: %t", atm, atm.IsZero())
	// Output: *schemax.AttributeTypes is zero: true
}

func ExampleNewObjectClasses() {
	ocm := NewObjectClasses()
	fmt.Printf("%T is zero: %t", ocm, ocm.IsZero())
	// Output: *schemax.ObjectClasses is zero: true
}

func ExampleNewLDAPSyntaxes() {
	lsm := NewLDAPSyntaxes()
	fmt.Printf("%T is zero: %t", lsm, lsm.IsZero())
	// Output: *schemax.LDAPSyntaxes is zero: true
}

func ExampleNewMatchingRules() {
	mrm := NewMatchingRules()
	fmt.Printf("%T is zero: %t", mrm, mrm.IsZero())
	// Output: *schemax.MatchingRules is zero: true
}

func ExampleNewMatchingRuleUses() {
	mrum := NewMatchingRuleUses()
	fmt.Printf("%T is zero: %t", mrum, mrum.IsZero())
	// Output: *schemax.MatchingRuleUses is zero: true
}

func ExampleNewNameForms() {
	nfm := NewNameForms()
	fmt.Printf("%T is zero: %t", nfm, nfm.IsZero())
	// Output: *schemax.NameForms is zero: true
}

func ExampleNewExtensions() {
	ext := NewExtensions()
	fmt.Printf("%T is zero: %t", ext, ext.IsZero())
	// Output: schemax.Extensions is zero: true
}

func ExampleMatchingRule_Equal() {
	sch := NewSubschema()
	sch.LSC = PopulateDefaultLDAPSyntaxes()
	sch.MRC = PopulateDefaultMatchingRules()

	mr1 := sch.MRC.Get(`caseIgnoreMatch`)
	mr2 := sch.MRC.Get(`caseExactMatch`)

	fmt.Printf("%T instances are equal: %t\n", mr1, mr1.Equal(mr2))
	// Output: *schemax.MatchingRule instances are equal: false
}

func ExampleOrdering_Equal() {
	sch := NewSubschema()
	sch.LSC = PopulateDefaultLDAPSyntaxes()
	sch.MRC = PopulateDefaultMatchingRules()

	iOM := sch.MRC.Get(`integerOrderingMatch`)
	ord1 := &Ordering{}
	ord2 := &Ordering{iOM}

	fmt.Printf("%T instances are equal: %t\n", ord1, ord1.Equal(ord2))
}

func ExampleSubstring_Equal() {
	sch := NewSubschema()
	sch.LSC = PopulateDefaultLDAPSyntaxes()
	sch.MRC = PopulateDefaultMatchingRules()

	cISM := sch.MRC.Get(`caseIgnoreSubstringsMatch`)
	ss1 := &Substring{}
	ss2 := &Substring{cISM}

	fmt.Printf("%T instances are equal: %t\n", ss1, ss1.Equal(ss2))
}

func ExampleEquality_Equal() {
	sch := NewSubschema()
	sch.LSC = PopulateDefaultLDAPSyntaxes()
	sch.MRC = PopulateDefaultMatchingRules()

	cEIM := sch.MRC.Get(`caseExactIA5Match`)
	eq1 := &Equality{}
	eq2 := &Equality{cEIM}

	fmt.Printf("%T instances are equal: %t\n", eq1, eq1.Equal(eq2))
}

func ExampleSubschema_MarshalAttributeType() {
	sch := NewSubschema()
	sch.LSC = PopulateDefaultLDAPSyntaxes()
	sch.MRC = PopulateDefaultMatchingRules()
	sch.ATC = NewAttributeTypes()

	myDef := `( 1.3.6.1.4.1.56521.999.104.17.2.1.10
			NAME ( 'jessesNewAttr' 'jesseAltName' )
			DESC 'This is an example AttributeType definition'
			OBSOLETE SYNTAX 1.3.6.1.4.1.1466.115.121.1.15
			EQUALITY caseIgnoreMatch X-ORIGIN 'Jesse Coretta' )`

	if err := sch.MarshalAttributeType(myDef); err != nil {
		panic(err)
	}

	jattr := sch.ATC.Get(`jessesNewAttr`)
	if !jattr.IsZero() {
		fmt.Printf("%T.OID: %s\n", jattr, jattr.OID)
	}
	// Output: *schemax.AttributeType.OID: 1.3.6.1.4.1.56521.999.104.17.2.1.10
}

func ExampleSubschema_MarshalObjectClass() {
	sch := NewSubschema()
	sch.LSC = PopulateDefaultLDAPSyntaxes()
	sch.MRC = PopulateDefaultMatchingRules()
	sch.ATC = PopulateDefaultAttributeTypes()
	sch.OCC = PopulateDefaultObjectClasses()

	myDef := `( 1.3.6.1.4.1.56521.999.104.17.2.2.91
			NAME ( 'jesseOC' 'specialJesseClass' )
			DESC 'This is an example ObjectClass definition'
			SUP account
			AUXILIARY
			MUST ( sn $ givenName $ carLicense )
			MAY ( description $ mobile )
			X-ORIGIN 'Jesse Coretta' )`

	if err := sch.MarshalObjectClass(myDef); err != nil {
		panic(err)
	}

	joc := sch.OCC.Get(`jesseOC`)
	if !joc.IsZero() {
		fmt.Printf("%T.OID: %s\n", joc, joc.OID)
	}
	// Output: *schemax.ObjectClass.OID: 1.3.6.1.4.1.56521.999.104.17.2.2.91
}

func ExampleSubschema_MarshalNameForm() {
	sch := NewSubschema()
	sch.LSC = PopulateDefaultLDAPSyntaxes()
	sch.MRC = PopulateDefaultMatchingRules()
	sch.ATC = PopulateDefaultAttributeTypes()
	sch.OCC = PopulateDefaultObjectClasses()
	sch.NFC = NewNameForms()

	myDef := `( 1.3.6.1.4.1.56521.999.104.17.2.16.4
                NAME 'jesseNameForm'
                DESC 'this is an example Name Form definition'
                OC account
                MUST ( cn $ o $ uid )
                MAY ( description $ l )
                X-ORIGIN 'Jesse Coretta' )`

	if err := sch.MarshalNameForm(myDef); err != nil {
		panic(err)
	}

	jnf := sch.NFC.Get(`jesseNameForm`)
	if !jnf.IsZero() {
		fmt.Printf("%T.OID: %s\n", jnf, jnf.OID)
	}
	// Output: *schemax.NameForm.OID: 1.3.6.1.4.1.56521.999.104.17.2.16.4
}

func ExampleSubschema_MarshalDITStructureRule() {
	sch := NewSubschema()
	sch.LSC = PopulateDefaultLDAPSyntaxes()
	sch.MRC = PopulateDefaultMatchingRules()
	sch.ATC = PopulateDefaultAttributeTypes()
	sch.OCC = PopulateDefaultObjectClasses()

	nf := `( 1.3.6.1.4.1.56521.999.104.17.2.16.4
                NAME 'jesseNameForm'
                DESC 'this is an example Name Form definition'
                OC account
                MUST ( cn $ o $ uid )
                MAY ( description $ l )
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

	jdsr := sch.DSRC.Get(`jesseDSR`)
	if !jdsr.IsZero() {
		fmt.Printf("%T.ID: %d\n", jdsr, jdsr.ID)
	}
	// Output: *schemax.DITStructureRule.ID: 0
}

func ExampleSubschema_MarshalDITContentRule() {
	sch := NewSubschema()
	sch.LSC = PopulateDefaultLDAPSyntaxes()
	sch.MRC = PopulateDefaultMatchingRules()
	sch.ATC = PopulateDefaultAttributeTypes()
	sch.OCC = PopulateDefaultObjectClasses()
	sch.DCRC = NewDITContentRules()

	dcr := `( 0.9.2342.19200300.100.4.5
                NAME 'accountContentRule'
                AUX ( posixAccount $ shadowAccount )
                MUST uid
                MAY description
                NOT gecos )`

	if err := sch.MarshalDITContentRule(dcr); err != nil {
		panic(err)
	}

	aDCR := sch.DCRC.Get(`accountContentRule`)
	if !aDCR.IsZero() {
		fmt.Printf("%T.OID: %s\n", aDCR, aDCR.OID)
	}
	// Output: *schemax.DITContentRule.OID: 0.9.2342.19200300.100.4.5
}

func ExampleSubschema_MarshalLDAPSyntax() {
	sch := NewSubschema()

	ls := `( 1.3.6.1.4.1.1466.115.121.1.2
                DESC 'Access Point'
                X-ORIGIN 'RFC4517' )`

	if err := sch.MarshalLDAPSyntax(ls); err != nil {
		panic(err)
	}

	ap := sch.LSC.Get(`Access Point`)
	if !ap.IsZero() {
		fmt.Printf("%T.OID: %s (%s)\n", ap, ap.OID, ap.Description)
	}
	// Output: *schemax.LDAPSyntax.OID: 1.3.6.1.4.1.1466.115.121.1.2 ('Access Point')
}

func ExampleSubschema_MarshalMatchingRule() {
	sch := NewSubschema()
	sch.LSC = PopulateDefaultLDAPSyntaxes()
	sch.MRC = NewMatchingRules()

	mr := `( 1.3.6.1.4.1.1466.109.114.3
                NAME 'caseIgnoreIA5SubstringsMatch'
                SYNTAX 1.3.6.1.4.1.1466.115.121.1.58
                X-ORIGIN 'RFC4517' )`

	if err := sch.MarshalMatchingRule(mr); err != nil {
		panic(err)
	}

	mrd := sch.MRC.Get(`caseIgnoreIA5SubstringsMatch`)
	if !mrd.IsZero() {
		fmt.Printf("%T.OID: %s\n", mrd, mrd.OID)
	}
	// Output: *schemax.MatchingRule.OID: 1.3.6.1.4.1.1466.109.114.3
}

func ExampleMarshal_newAttributeType() {
	lsc := PopulateDefaultLDAPSyntaxes()
	mrc := PopulateDefaultMatchingRules()

	def := `( 1.3.6.1.4.1.56521.999.104.17.2.1.10
                NAME ( 'jessesNewAttr' 'jesseAltName' )
                DESC 'This is an example AttributeType definition'
                OBSOLETE SYNTAX 1.3.6.1.4.1.1466.115.121.1.15
                EQUALITY caseIgnoreMatch X-ORIGIN 'Jesse Coretta' )`

	var x AttributeType
	err := Marshal(def, &x, nil, nil, lsc, mrc, nil, nil, nil, nil)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%T is obsolete: %t\n", x, x.Obsolete())
	// Output: schemax.AttributeType is obsolete: true
}

/*
func ExampleMarshal_newAttributeTypeWithFormatting() {
        lsc := PopulateDefaultLDAPSyntaxes()
        mrc := PopulateDefaultMatchingRules()

        def := `( 1.3.6.1.4.1.56521.999.104.17.2.1.10
                NAME ( 'jessesNewAttr' 'jesseAltName' )
                DESC 'This is an example AttributeType definition'
                OBSOLETE SYNTAX 1.3.6.1.4.1.1466.115.121.1.15
                EQUALITY caseIgnoreMatch X-ORIGIN 'Jesse Coretta' )`

        var x AttributeType
	x.SetUnmarshalFunc(x.AttributeTypeUnmarshalFunc)
        err := Marshal(def, &x, nil, nil, lsc, mrc, nil, nil, nil, nil)
        if err != nil {
                panic(err)
        }

	var back string
	if back, err = Unmarshal(&x); err != nil {
                panic(err)
        }

        fmt.Printf("%T is obsolete: %t\n", x, x.Obsolete())
        // Output: schemax.AttributeType is obsolete: true
}
*/

func ExampleMarshal_newObjectClass() {
	atc := PopulateDefaultAttributeTypes()
	occ := PopulateDefaultObjectClasses()

	def := `( 1.3.6.1.4.1.56521.999.107.12.1.2.11
			NAME ( 'testClass' )
			DESC 'This is an example ObjectClass definition'
			STRUCTURAL
			MUST ( uid $ cn $ o )
			MAY ( description $ l $ c )
			X-ORIGIN 'Jesse Coretta' )`

	var x ObjectClass
	err := Marshal(def, &x, atc, occ, nil, nil, nil, nil, nil, nil)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%T requires: %s\n", x, x.Must)
	// Output: schemax.ObjectClass requires: ( uid $ cn $ o )
}

func ExampleMarshal_newLDAPSyntax() {

	def := `( 1.3.6.1.4.1.56521.999.43.1.2.71
                DESC 'This is an example LDAPSyntax definition'
                X-ORIGIN 'Jesse Coretta' )`

	var x LDAPSyntax
	err := Marshal(def, &x, nil, nil, nil, nil, nil, nil, nil, nil)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%T OID: %s (%s)\n", x, x.OID, x.Description)
	// Output: schemax.LDAPSyntax OID: 1.3.6.1.4.1.56521.999.43.1.2.71 ('This is an example LDAPSyntax definition')
}

func ExampleMarshal_newMatchingRule() {
	lsc := PopulateDefaultLDAPSyntaxes()

	def := `( 1.3.6.1.4.1.56521.999.314.21.1
                NAME 'alternativeCaseIgnoreStringMatch'
                DESC 'This is an example MatchingRule definition'
                SYNTAX 1.3.6.1.4.1.1466.115.121.1.15
                X-ORIGIN 'Jesse Coretta' )`

	var x MatchingRule
	err := Marshal(def, &x, nil, nil, lsc, nil, nil, nil, nil, nil)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%T syntax OID: %s\n", x, x.Syntax.OID)
	// Output: schemax.MatchingRule syntax OID: 1.3.6.1.4.1.1466.115.121.1.15
}

func ExampleMarshal_newDITContentRule() {
	atc := PopulateDefaultAttributeTypes()
	occ := PopulateDefaultObjectClasses()

	def := `( 2.16.840.1.113730.3.2.2
		NAME 'inetOrgPerson-content-rule'
		AUX strongAuthenticationUser
		MUST uid
		MAY ( c $ l )
		NOT telexNumber )`

	var x DITContentRule
	err := Marshal(def, &x, atc, occ, nil, nil, nil, nil, nil, nil)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%T allows: %s\n", x, x.May)
	// Output: schemax.DITContentRule allows: ( c $ l )
}

func ExampleMarshal_newNameForm() {
	atc := PopulateDefaultAttributeTypes()
	occ := PopulateDefaultObjectClasses()

	nf := `( 1.3.6.1.4.1.56521.999.104.17.2.16.4
                NAME 'jesseNameForm'
                DESC 'this is an example Name Form definition'
                OC account
                MUST ( cn $ o $ uid )
                MAY ( description $ l )
                X-ORIGIN 'Jesse Coretta' )`

	var z NameForm
	err := Marshal(nf, &z, atc, occ, nil, nil, nil, nil, nil, nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%T requires: %s\n", z, z.Must)
	// Output: schemax.NameForm requires: ( cn $ o $ uid )
}

func ExampleMarshal_newDITStructureRule() {
	atc := PopulateDefaultAttributeTypes()
	occ := PopulateDefaultObjectClasses()

	nf := `( 1.3.6.1.4.1.56521.999.104.17.2.16.4
                NAME 'jesseNameForm'
                DESC 'this is an example Name Form definition'
                OC account
                MUST ( cn $ o $ uid )
                MAY ( description $ l )
                X-ORIGIN 'Jesse Coretta' )`

	var z NameForm
	err := Marshal(nf, &z, atc, occ, nil, nil, nil, nil, nil, nil)
	if err != nil {
		panic(err)
	}

	nfc := NewNameForms()
	nfc.Set(&z)

	dsr := `( 0
                NAME 'jesseDSR'
                DESC 'this is an example DIT Structure Rule'
                FORM jesseNameForm
                X-ORIGIN 'Jesse Coretta' )`

	var x DITStructureRule
	err = Marshal(dsr, &x, nil, nil, nil, nil, nil, nil, nil, nfc)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%T NameForm: %s\n", x, x.Form.Name.Index(0))
	// Output: schemax.DITStructureRule NameForm: jesseNameForm
}

func ExampleUnmarshal() {
	sch := NewSubschema()
	sch.LSC = PopulateDefaultLDAPSyntaxes()
	sch.MRC = PopulateDefaultMatchingRules()
	sch.ATC = PopulateDefaultAttributeTypes()

	name := sch.ATC.Get(`name`)
	def, err := Unmarshal(name)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", def)
	// Output: ( 2.5.4.41 NAME 'name' SYNTAX 1.3.6.1.4.1.1466.115.121.1.15 EQUALITY caseIgnoreMatch SUBSTR caseIgnoreSubstringsMatch X-ORIGIN 'RFC4519' )
}

func ExampleAttributeTypes_Get() {
	sch := NewSubschema()
	sch.LSC = PopulateDefaultLDAPSyntaxes()
	sch.MRC = PopulateDefaultMatchingRules()
	sch.ATC = PopulateDefaultAttributeTypes()

	if attr := sch.ATC.Get(`name`); !attr.IsZero() {
		fmt.Printf("Name OID: %s\n", attr.OID)
	}
	// Output: Name OID: 2.5.4.41
}

func ExampleObjectClasses_Get() {
	sch := NewSubschema()
	sch.LSC = PopulateDefaultLDAPSyntaxes()
	sch.MRC = PopulateDefaultMatchingRules()
	sch.ATC = PopulateDefaultAttributeTypes()
	sch.OCC = PopulateDefaultObjectClasses()

	if class := sch.OCC.Get(`account`); !class.IsZero() {
		fmt.Printf("account OID: %s\n", class.OID)
	}
	// Output: account OID: 0.9.2342.19200300.100.4.5
}

func ExampleLDAPSyntaxes_Get() {
	sch := NewSubschema()
	sch.LSC = PopulateDefaultLDAPSyntaxes()

	if ls := sch.LSC.Get(`1.3.6.1.4.1.1466.115.121.1.4`); !ls.IsZero() {
		fmt.Printf("%T OID: %s\n", ls, ls.OID)
	}
	// Output: *schemax.LDAPSyntax OID: 1.3.6.1.4.1.1466.115.121.1.4
}

func ExampleAttributeTypes_IsZero() {
	sch := NewSubschema()
	sch.LSC = PopulateDefaultLDAPSyntaxes()
	sch.MRC = PopulateDefaultMatchingRules()
	sch.ATC = PopulateDefaultAttributeTypes()

	isZero := sch.ATC.IsZero()
	fmt.Printf("%T is zero: %t\n", sch.ATC, isZero)
	// Output: *schemax.AttributeTypes is zero: false
}

func ExampleObjectClasses_IsZero() {
	sch := NewSubschema()
	sch.LSC = PopulateDefaultLDAPSyntaxes()
	sch.MRC = PopulateDefaultMatchingRules()
	sch.ATC = PopulateDefaultAttributeTypes()
	sch.OCC = PopulateDefaultObjectClasses()

	isZero := sch.OCC.IsZero()
	fmt.Printf("%T is zero: %t\n", sch.OCC, isZero)
	// Output: *schemax.ObjectClasses is zero: false
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

func ExampleApplicableAttributeTypes_IsZero() {
	aps := NewApplicableAttributeTypes()
	fmt.Printf("%T is zero: %t\n", aps, aps.IsZero())
	// Output: *schemax.ApplicableAttributeTypes is zero: true
}

func ExampleProhibitedAttributeTypes_IsZero() {
	aps := NewProhibitedAttributeTypes()
	fmt.Printf("%T is zero: %t\n", aps, aps.IsZero())
	// Output: *schemax.ProhibitedAttributeTypes is zero: true
}

func ExampleRequiredAttributeTypes_IsZero() {
	aps := NewRequiredAttributeTypes()
	fmt.Printf("%T is zero: %t\n", aps, aps.IsZero())
	// Output: *schemax.RequiredAttributeTypes is zero: true
}

func ExamplePermittedAttributeTypes_IsZero() {
	aps := NewPermittedAttributeTypes()
	fmt.Printf("%T is zero: %t\n", aps, aps.IsZero())
	// Output: *schemax.PermittedAttributeTypes is zero: true
}

func ExampleDITStructureRule_Equal() {
	dsr1 := &DITStructureRule{}
	dsr2 := &DITStructureRule{ID: NewRuleID(0)}
	fmt.Printf("%T instances are equal: %t\n", dsr1, dsr1.Equal(dsr2))
	// Output: *schemax.DITStructureRule instances are equal: false
}

func ExampleLDAPSyntax_Equal() {
	sch := NewSubschema()
	sch.LSC = PopulateDefaultLDAPSyntaxes()

	syn1 := &LDAPSyntax{}
	syn2 := sch.LSC.Get(`1.3.6.1.4.1.1466.115.121.1.4`) // Audio per RFC4517
	fmt.Printf("%T instances are equal: %t\n", syn1, syn1.Equal(syn2))
	// Output: *schemax.LDAPSyntax instances are equal: false
}

func ExampleAuxiliaryObjectClasses_IsZero() {
	var aux AuxiliaryObjectClasses
	fmt.Printf("%T is zero: %t\n", aux, aux.IsZero())
	// Output: schemax.AuxiliaryObjectClasses is zero: true
}

func ExampleSuperiorObjectClasses_IsZero() {
	sup := SuperiorObjectClasses{}
	fmt.Printf("%T is zero: %t\n", sup, sup.IsZero())
	// Output: schemax.SuperiorObjectClasses is zero: true
}

func ExampleSuperiorDITStructureRules_IsZero() {
	sup := SuperiorDITStructureRules{}
	fmt.Printf("%T is zero: %t\n", sup, sup.IsZero())
	// Output: schemax.SuperiorDITStructureRules is zero: true
}

func ExampleLDAPSyntax_HumanReadable() {
	sch := NewSubschema()
	sch.LSC = PopulateDefaultLDAPSyntaxes()

	audio := sch.LSC.Get(`1.3.6.1.4.1.1466.115.121.1.4`) // Audio per RFC4517
	fmt.Printf("%T (%s) is Human Readable: %t\n", audio, audio.OID, audio.HumanReadable())
	// Output: *schemax.LDAPSyntax (1.3.6.1.4.1.1466.115.121.1.4) is Human Readable: false
}

func ExampleAuxiliaryObjectClasses_Equal() {
	sch := NewSubschema()
	sch.LSC = PopulateDefaultLDAPSyntaxes()
	sch.MRC = PopulateDefaultMatchingRules()
	sch.ATC = PopulateDefaultAttributeTypes()
	sch.OCC = PopulateDefaultObjectClasses()

	aux1 := NewAuxiliaryObjectClasses()
	_ = aux1.Set(sch.OCC.Get(`posixAccount`))

	aux2 := NewAuxiliaryObjectClasses()
	fmt.Printf("%T instances are equal: %t\n", aux1, aux1.Equal(aux2))
	// Output *schemax.AuxiliaryObjectClasses instances are equal: false
}

func ExamplePermittedAttributeTypes_Equal() {
	sch := NewSubschema()
	sch.LSC = PopulateDefaultLDAPSyntaxes()
	sch.MRC = PopulateDefaultMatchingRules()
	sch.ATC = PopulateDefaultAttributeTypes()

	may1 := NewPermittedAttributeTypes()
	_ = may1.Set(sch.ATC.Get(`givenName`))

	may2 := NewPermittedAttributeTypes()
	fmt.Printf("%T instances are equal: %t\n", may1, may1.Equal(may2))
	// Output *schemax.PermittedAttributeTypes instances are equal: false
}

func ExampleRequiredAttributeTypes_Equal() {
	sch := NewSubschema()
	sch.LSC = PopulateDefaultLDAPSyntaxes()
	sch.MRC = PopulateDefaultMatchingRules()
	sch.ATC = PopulateDefaultAttributeTypes()

	must1 := NewRequiredAttributeTypes()
	_ = must1.Set(sch.ATC.Get(`cn`))

	must2 := NewRequiredAttributeTypes()
	fmt.Printf("%T instances are equal: %t\n", must1, must1.Equal(must2))
	// Output *schemax.RequiredAttributeTypes instances are equal: false
}

func ExampleProhibitedAttributeTypes_Equal() {
	sch := NewSubschema()
	sch.LSC = PopulateDefaultLDAPSyntaxes()
	sch.MRC = PopulateDefaultMatchingRules()
	sch.ATC = PopulateDefaultAttributeTypes()

	not1 := NewProhibitedAttributeTypes()
	_ = not1.Set(sch.ATC.Get(`sn`))

	not2 := NewProhibitedAttributeTypes()
	fmt.Printf("%T instances are equal: %t\n", not1, not1.Equal(not2))
	// Output *schemax.ProhibitedAttributeTypes instances are equal: false
}

func ExampleProhibitedAttributeTypes_Index() {
	sch := NewSubschema()
	sch.LSC = PopulateDefaultLDAPSyntaxes()
	sch.MRC = PopulateDefaultMatchingRules()
	sch.ATC = PopulateDefaultAttributeTypes()

	attr := sch.ATC.Get(`cn`)
	not := NewProhibitedAttributeTypes()
	not.Set(attr)

	idx, exists := not.Contains(`2.5.4.3`)
	if exists {
		fmt.Printf("Index for %s attr in %T: %d\n", attr.Name.Index(0), not, idx)
	}
	// Output: Index for cn attr in *schemax.ProhibitedAttributeTypes: 0
}

func ExampleAuxiliaryObjectClasses_Index() {
	sch := NewSubschema()
	sch.LSC = PopulateDefaultLDAPSyntaxes()
	sch.MRC = PopulateDefaultMatchingRules()
	sch.ATC = PopulateDefaultAttributeTypes()
	sch.OCC = PopulateDefaultObjectClasses()

	class := sch.OCC.Get(`posixAccount`)
	aux := NewAuxiliaryObjectClasses()
	aux.Set(class)

	idx, exists := aux.Contains(`1.3.6.1.1.1.2.0`)
	if exists {
		fmt.Printf("Index for %s class in %T: %d\n", class.Name.Index(0), aux, idx)
	}
	// Output: Index for posixAccount class in *schemax.AuxiliaryObjectClasses: 0
}

func ExampleRequiredAttributeTypes_Index() {
	sch := NewSubschema()
	sch.LSC = PopulateDefaultLDAPSyntaxes()
	sch.MRC = PopulateDefaultMatchingRules()
	sch.ATC = PopulateDefaultAttributeTypes()

	attr := sch.ATC.Get(`cn`)
	must := NewRequiredAttributeTypes()
	must.Set(attr)

	idx, exists := must.Contains(`2.5.4.3`)
	if exists {
		fmt.Printf("Index for %s attr in %T: %d\n", attr.Name.Index(0), must, idx)
	}
	// Output: Index for cn attr in *schemax.RequiredAttributeTypes: 0
}

func ExamplePermittedAttributeTypes_Index() {
	sch := NewSubschema()
	sch.LSC = PopulateDefaultLDAPSyntaxes()
	sch.MRC = PopulateDefaultMatchingRules()
	sch.ATC = PopulateDefaultAttributeTypes()

	attr := sch.ATC.Get(`cn`)
	may := NewPermittedAttributeTypes()
	may.Set(attr)

	idx, exists := may.Contains(`2.5.4.3`)
	if exists {
		fmt.Printf("Index for %s attr in %T: %d\n", attr.Name.Index(0), may, idx)
	}
	// Output: Index for cn attr in *schemax.PermittedAttributeTypes: 0
}

func ExampleApplicableAttributeTypes_Index() {
	sch := NewSubschema()
	sch.LSC = PopulateDefaultLDAPSyntaxes()
	sch.MRC = PopulateDefaultMatchingRules()
	sch.ATC = PopulateDefaultAttributeTypes()

	attr := sch.ATC.Get(`cn`)
	applies := NewApplicableAttributeTypes()
	applies.Set(attr)

	idx, exists := applies.Contains(`2.5.4.3`)
	if exists {
		fmt.Printf("Index for %s attr in %T: %d\n", attr.Name.Index(0), applies, idx)
	}
	// Output: Index for cn attr in *schemax.ApplicableAttributeTypes: 0
}

func ExampleAttributeType_SetMaxLength() {
	sch := NewSubschema()
	sch.LSC = PopulateDefaultLDAPSyntaxes()

	at := &AttributeType{
		Name:   NewName(`fakeName`),
		OID:    NewOID(`1.3.6.1.4.1.56521.999.104.17.2.1.55`),
		Syntax: sch.LSC.Get(`1.3.6.1.4.1.1466.115.121.1.15`),
	}
	at.SetMaxLength(128)
	fmt.Printf("Max allowed length for %T value: %d\n", at, at.MaxLength())
	// Output: Max allowed length for *schemax.AttributeType value: 128
}

func ExampleAttributeType_MaxLength() {
	sch := NewSubschema()
	sch.LSC = PopulateDefaultLDAPSyntaxes()

	at := &AttributeType{
		Name:   NewName(`fakeName`),
		OID:    NewOID(`1.3.6.1.4.1.56521.999.104.17.2.1.55`),
		Syntax: sch.LSC.Get(`1.3.6.1.4.1.1466.115.121.1.15`),
	}
	at.SetMaxLength(128)
	fmt.Printf("Max allowed length for %T value: %d\n", at, at.MaxLength())
	// Output: Max allowed length for *schemax.AttributeType value: 128
}

func ExampleAttributeType_Equal() {
	sch := NewSubschema()
	sch.LSC = PopulateDefaultLDAPSyntaxes()

	at1 := &AttributeType{
		Name:   NewName(`fakeName`),
		OID:    NewOID(`1.3.6.1.4.1.56521.999.104.17.2.1.55`),
		Syntax: sch.LSC.Get(`1.3.6.1.4.1.1466.115.121.1.15`),
	}

	at2 := &AttributeType{
		Name:   NewName(`fakeName`),
		OID:    NewOID(`1.3.6.1.4.1.56521.999.104.17.2.1.55`),
		Syntax: sch.LSC.Get(`1.3.6.1.4.1.1466.115.121.1.15`),
	}

	fmt.Printf("at1 is equal to at2: %t\n", at1.Equal(at2))
	// Output: at1 is equal to at2: true
}

func ExampleObjectClass_Equal() {
	oc1 := &ObjectClass{
		Name: NewName(`fakeName`),
		OID:  NewOID(`1.3.6.1.4.1.56521.999.104.17.2.2.84`),
		Kind: Auxiliary,
	}

	oc2 := &ObjectClass{
		Name: NewName(`fakeName`),
		OID:  NewOID(`1.3.6.1.4.1.56521.999.104.17.2.2.81`),
		Kind: Structural,
	}

	fmt.Printf("%T instances are equal: %t\n", oc1, oc1.Equal(oc2))
	// Output: *schemax.ObjectClass instances are equal: false
}

func ExampleStructuralObjectClass_Equal() {
	oc1 := &ObjectClass{
		Name: NewName(`fakeName`),
		OID:  NewOID(`1.3.6.1.4.1.56521.999.104.17.2.2.84`),
		Kind: Auxiliary,
	}
	soc1 := StructuralObjectClass{oc1}

	oc2 := &ObjectClass{
		Name: NewName(`fakeName`),
		OID:  NewOID(`1.3.6.1.4.1.56521.999.104.17.2.2.81`),
		Kind: Structural,
	}

	fmt.Printf("%T instances are equal: %t\n", oc1, soc1.Equal(oc2))
	// Output: *schemax.ObjectClass instances are equal: false
}

func ExampleLDAPSyntaxes_Contains() {
	sch := NewSubschema()
	sch.LSC = PopulateDefaultLDAPSyntaxes()

	_, found := sch.LSC.Contains(`1.3.6.1.4.1.1466.115.121.1.4`)
	fmt.Printf("syntax exists: %t\n", found)
	// Output: syntax exists: true
}

func ExampleObjectClasses_Contains() {
	sch := NewSubschema()
	sch.LSC = PopulateDefaultLDAPSyntaxes()
	sch.MRC = PopulateDefaultMatchingRules()
	sch.ATC = PopulateDefaultAttributeTypes()
	sch.OCC = PopulateDefaultObjectClasses()

	_, found := sch.OCC.Contains(`account`)
	fmt.Printf("account ObjectClass exists: %t\n", found)
	// Output: account ObjectClass exists: true
}

func ExampleAttributeTypes_Set() {

	sch := NewSubschema()
	sch.LSC = PopulateDefaultLDAPSyntaxes()
	sch.MRC = PopulateDefaultMatchingRules()
	sch.ATC = PopulateDefaultAttributeTypes()

	sch.ATC.Set(&AttributeType{
		Name:   NewName(`fakeName`),
		OID:    NewOID(`1.3.6.1.4.1.56521.999.104.17.2.1.55`),
		Syntax: sch.LSC.Get(`1.3.6.1.4.1.1466.115.121.1.15`),
	})

	if get := sch.ATC.Get(`fakeName`); !get.IsZero() {
		fmt.Printf("fakeName OID: %s\n", get.OID)
	}
	// Output: fakeName OID: 1.3.6.1.4.1.56521.999.104.17.2.1.55
}

func ExampleAttributeTypes_Contains() {

	sch := NewSubschema()
	sch.LSC = PopulateDefaultLDAPSyntaxes()
	sch.MRC = PopulateDefaultMatchingRules()
	sch.ATC = PopulateDefaultAttributeTypes()

	_, exists := sch.ATC.Contains(`name`)
	fmt.Printf("name AttributeType exists: %t\n", exists)
	// Output: name AttributeType exists: true
}

func ExampleAttributeTypes_Equal() {

	sch := NewSubschema()
	sch.LSC = PopulateDefaultLDAPSyntaxes()
	sch.MRC = PopulateDefaultMatchingRules()
	sch.ATC = PopulateDefaultAttributeTypes()

	otherATC := NewAttributeTypes()
	same := sch.ATC.Equal(otherATC)
	fmt.Printf("%T instances are equal: %t\n", sch.ATC, same)
	// Output: *schemax.AttributeTypes instances are equal: false
}

func ExampleObjectClasses_Equal() {

	sch := NewSubschema()
	sch.LSC = PopulateDefaultLDAPSyntaxes()
	sch.MRC = PopulateDefaultMatchingRules()
	sch.ATC = PopulateDefaultAttributeTypes()
	sch.OCC = PopulateDefaultObjectClasses()

	otherOCC := NewObjectClasses()
	same := sch.OCC.Equal(otherOCC)
	fmt.Printf("%T instances are equal: %t\n", sch.OCC, same)
	// Output: *schemax.ObjectClasses instances are equal: false
}

func ExampleLDAPSyntaxes_Equal() {

	sch := NewSubschema()
	sch.LSC = PopulateDefaultLDAPSyntaxes()

	otherLSC := NewLDAPSyntaxes()
	same := sch.LSC.Equal(otherLSC)
	fmt.Printf("%T instances are equal: %t\n", sch.LSC, same)
	// Output: *schemax.LDAPSyntaxes instances are equal: false
}

func ExampleMatchingRules_Equal() {

	sch := NewSubschema()
	sch.LSC = PopulateDefaultLDAPSyntaxes()
	sch.MRC = PopulateDefaultMatchingRules()

	otherMRC := NewMatchingRules()
	same := sch.MRC.Equal(otherMRC)
	fmt.Printf("%T instances are equal: %t\n", sch.MRC, same)
	// Output: *schemax.MatchingRules instances are equal: false
}

func ExampleDITContentRule_Equal() {

	sch := NewSubschema()
	sch.LSC = PopulateDefaultLDAPSyntaxes()
	sch.MRC = PopulateDefaultMatchingRules()
	sch.ATC = PopulateDefaultAttributeTypes()
	sch.OCC = PopulateDefaultObjectClasses()
	sch.DCRC = NewDITContentRules()

	posixAcct := sch.OCC.Get(`posixAccount`)
	shadowAcct := sch.OCC.Get(`posixAccount`)
	uidAttr := sch.ATC.Get(`uid`)
	descAttr := sch.ATC.Get(`description`)
	givenNameAttr := sch.ATC.Get(`givenName`)

	if posixAcct.IsZero() || shadowAcct.IsZero() {
		panic("missing key ObjectClasse instances")
	}

	if uidAttr.IsZero() || descAttr.IsZero() || givenNameAttr.IsZero() {
		panic("missing key AttributeType instances")
	}

	aux := NewAuxiliaryObjectClasses()
	aux.Set(posixAcct)
	aux.Set(shadowAcct)

	must := NewRequiredAttributeTypes()
	must.Set(uidAttr)

	may := NewPermittedAttributeTypes()
	may.Set(descAttr)

	not := NewProhibitedAttributeTypes()
	not.Set(givenNameAttr)

	dcr1 := &DITContentRule{
		OID:  NewOID(`0.9.2342.19200300.100.4.5`),
		Name: NewName(`jesseAccountContentRule`),
		Aux:  aux,
		Must: must,
		May:  may,
		Not:  not,
	}

	dcr2 := &DITContentRule{
		OID:  NewOID(`0.9.2342.19200300.100.4.5`),
		Name: NewName(`jesseAccountContentRule2`),
		Aux:  aux,
		Must: must,
		May:  may,
		Not:  not,
	}

	same := dcr1.Equal(dcr2)
	fmt.Printf("The two %T instances are equal: %t\n", dcr1, same)
	// Output: The two *schemax.DITContentRule instances are equal: false

}

func ExampleDITContentRules_Equal() {

	dcrm1 := NewDITContentRules()
	dcrm2 := NewDITContentRules()

	fmt.Printf("%T instances are equal: %t\n", dcrm1, dcrm1.Equal(dcrm2))
	// Output: *schemax.DITContentRules instances are equal: true
}

func ExampleDITContentRules_Get() {

	sch := NewSubschema()
	sch.LSC = PopulateDefaultLDAPSyntaxes()
	sch.MRC = PopulateDefaultMatchingRules()
	sch.ATC = PopulateDefaultAttributeTypes()
	sch.OCC = PopulateDefaultObjectClasses()
	sch.DCRC = NewDITContentRules()

	posixAcct := sch.OCC.Get(`posixAccount`)
	shadowAcct := sch.OCC.Get(`posixAccount`)
	uidAttr := sch.ATC.Get(`uid`)
	descAttr := sch.ATC.Get(`description`)
	givenNameAttr := sch.ATC.Get(`givenName`)

	if posixAcct.IsZero() || shadowAcct.IsZero() {
		panic("missing key ObjectClasse instances")
	}

	if uidAttr.IsZero() || descAttr.IsZero() || givenNameAttr.IsZero() {
		panic("missing key AttributeType instances")
	}

	aux := NewAuxiliaryObjectClasses()
	aux.Set(posixAcct)
	aux.Set(shadowAcct)

	must := NewRequiredAttributeTypes()
	must.Set(uidAttr)

	may := NewPermittedAttributeTypes()
	may.Set(descAttr)

	not := NewProhibitedAttributeTypes()
	not.Set(givenNameAttr)

	sch.DCRC.Set(&DITContentRule{
		OID:  NewOID(`0.9.2342.19200300.100.4.5`),
		Name: NewName(`jesseAccountContentRule`),
		Aux:  aux,
		Must: must,
		May:  may,
		Not:  not,
	})

	dcr := sch.DCRC.Get(`jesseAccountContentRule`)
	fmt.Printf("%T.Name: %s\n", *dcr, dcr.Name.Index(0))
	// Output: schemax.DITContentRule.Name: jesseAccountContentRule

}

func ExampleDITContentRules_Set() {
	sch := NewSubschema()
	sch.LSC = PopulateDefaultLDAPSyntaxes()
	sch.MRC = PopulateDefaultMatchingRules()
	sch.ATC = PopulateDefaultAttributeTypes()
	sch.OCC = PopulateDefaultObjectClasses()
	sch.DCRC = NewDITContentRules()

	posixAcct := sch.OCC.Get(`posixAccount`)
	shadowAcct := sch.OCC.Get(`posixAccount`)
	uidAttr := sch.ATC.Get(`uid`)
	descAttr := sch.ATC.Get(`description`)
	givenNameAttr := sch.ATC.Get(`givenName`)

	if posixAcct.IsZero() || shadowAcct.IsZero() {
		panic("missing key ObjectClass instances")
	}

	if uidAttr.IsZero() || descAttr.IsZero() || givenNameAttr.IsZero() {
		panic("missing key AttributeType instances")
	}

	aux := NewAuxiliaryObjectClasses()
	aux.Set(posixAcct)
	aux.Set(shadowAcct)

	must := NewRequiredAttributeTypes()
	must.Set(uidAttr)

	may := NewPermittedAttributeTypes()
	may.Set(descAttr)

	not := NewProhibitedAttributeTypes()
	not.Set(givenNameAttr)

	sch.DCRC.Set(&DITContentRule{
		OID:  NewOID(`0.9.2342.19200300.100.4.5`),
		Name: NewName(`jesseAccountContentRule`),
		Aux:  aux,
		Must: must,
		May:  may,
		Not:  not,
	})

	dcr := sch.DCRC.Get(`jesseAccountContentRule`)
	fmt.Printf("%T.Name: %s\n", *dcr, dcr.Name.Index(0))
	// Output: schemax.DITContentRule.Name: jesseAccountContentRule
}

func ExampleDITContentRules_Contains() {
	sch := NewSubschema()
	dcr := &DITContentRule{}
	_, found := sch.DCRC.Contains(dcr)
	fmt.Printf("%T exists within %T: %t\n", dcr, sch.DCRC, found)
	// Output: *schemax.DITContentRule exists within *schemax.DITContentRules: false
}

func ExamplePopulateDefaultAttributeTypes() {
	atm := NewAttributeTypes()
	fmt.Printf("%T is zero: %t", atm, atm.IsZero())
	// Output: *schemax.AttributeTypes is zero: true
}

func ExampleApplicableAttributeTypes_Label() {
	aps := ApplicableAttributeTypes{}
	lab := aps.Label()
	fmt.Printf("%T label: %s (len:%d)\n", aps, lab, len(lab))
	// Output: schemax.ApplicableAttributeTypes label: APPLIES (len:7)
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

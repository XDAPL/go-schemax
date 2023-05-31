package schemax

/*
init.go provided initialization routines for certain global variables and recommended content.
*/

import (
	"github.com/JesseCoretta/go-schemax/rfc2079"
	"github.com/JesseCoretta/go-schemax/rfc2307"
	"github.com/JesseCoretta/go-schemax/rfc2798"
	"github.com/JesseCoretta/go-schemax/rfc3045"
	"github.com/JesseCoretta/go-schemax/rfc3672"
	"github.com/JesseCoretta/go-schemax/rfc4512"
	"github.com/JesseCoretta/go-schemax/rfc4517"
	"github.com/JesseCoretta/go-schemax/rfc4519"
	"github.com/JesseCoretta/go-schemax/rfc4523"
	"github.com/JesseCoretta/go-schemax/rfc4524"
	"github.com/JesseCoretta/go-schemax/rfc4530"
)

/*
for internal go test runs only!
*/
func testSubschema() (x *Subschema) {
	x = NewSubschema()
	x.LSC = PopulateDefaultLDAPSyntaxes()
	x.MRC = PopulateDefaultMatchingRules()
	x.ATC = NewAttributeTypes()
	x.OCC = NewObjectClasses()

	return
}

////////////////
// LDAP Syntaxes

/*
PopulateDefaultLDAPSyntaxes returns a new auto-populated instance of LDAPSyntaxes.

Within this structure are all of the following:

• Syntax definitions from RFC2307 (imported from go-schemax/rfc2307)

• Syntax definitions from RFC4517 (imported from go-schemax/rfc4517)

• Syntax definitions from RFC4523 (imported from go-schemax/rfc4523)

• Syntax definitions from RFC4530 (imported from go-schemax/rfc4530)
*/
func PopulateDefaultLDAPSyntaxes() (lsc LDAPSyntaxCollection) {
	lsc = NewLDAPSyntaxes()
	PopulateRFC2307Syntaxes(lsc)
	PopulateRFC4517Syntaxes(lsc)
	PopulateRFC4523Syntaxes(lsc)
	PopulateRFC4530Syntaxes(lsc)

	return
}

/*
PopulateRFC4523Syntaxes only populates the provided LDAPSyntaxes (lsc) with LDAPSyntax definitions defined in this RFC.

Definitions are imported from go-schemax/rfc4523.
*/
func PopulateRFC4523Syntaxes(lsc LDAPSyntaxCollection) {
	for _, ls := range rfc4523.AllLDAPSyntaxes {
		var nls LDAPSyntax
		if err := Marshal(string(ls), &nls, nil, nil, nil, nil, nil, nil, nil, nil); err != nil {
			panic(err)
		}
		lsc.Set(&nls)
	}
}

/*
PopulateRFC2307Syntaxes only populates the provided LDAPSyntaxes (lsc) with LDAPSyntax definitions defined in this RFC.

Definitions are imported from go-schemax/rfc2307.
*/
func PopulateRFC2307Syntaxes(lsc LDAPSyntaxCollection) {
	for _, ls := range rfc2307.AllLDAPSyntaxes {
		var nls LDAPSyntax
		if err := Marshal(string(ls), &nls, nil, nil, nil, nil, nil, nil, nil, nil); err != nil {
			panic(err)
		}
		lsc.Set(&nls)
	}
}

/*
PopulateRFC4517Syntaxes only populates the provided LDAPSyntaxes (lsc) with LDAPSyntax definitions defined in this RFC.

Definitions are imported from go-schemax/rfc4517.
*/
func PopulateRFC4517Syntaxes(lsc LDAPSyntaxCollection) {
	for _, ls := range rfc4517.AllLDAPSyntaxes {
		var nls LDAPSyntax
		if err := Marshal(string(ls), &nls, nil, nil, nil, nil, nil, nil, nil, nil); err != nil {
			panic(sprintf("%s: %s", err.Error(), string(ls)))
		}
		lsc.Set(&nls)
	}
}

/*
PopulateRFC4530Syntaxes only populates the provided LDAPSyntaxes (lsc) with LDAPSyntax definitions defined in this RFC.

Definitions are imported from go-schemax/rfc4530.
*/
func PopulateRFC4530Syntaxes(lsc LDAPSyntaxCollection) {
	for _, ls := range rfc4530.AllLDAPSyntaxes {
		var nls LDAPSyntax
		if err := Marshal(string(ls), &nls, nil, nil, nil, nil, nil, nil, nil, nil); err != nil {
			panic(sprintf("%s: %s", err.Error(), string(ls)))
		}
		lsc.Set(&nls)
	}
}

/////////////////
// Matching Rules

/*
PopulateDefaultMatchingRules returns a new auto-populated instance of MatchingRules. Within this structure are all of the following:

• Matching rules from RFC2307 (imported from go-schemax/rfc2307)

• Matching rules from RFC4517 (imported from go-schemax/rfc4517)

• Matching rules from RFC4523 (imported from go-schemax/rfc4523)

• Matching rules from RFC4530 (imported from go-schemax/rfc4530)
*/
func PopulateDefaultMatchingRules() (mrc MatchingRuleCollection) {
	mrc = NewMatchingRules()
	PopulateRFC2307MatchingRules(DefaultLDAPSyntaxes, mrc)
	PopulateRFC4517MatchingRules(DefaultLDAPSyntaxes, mrc)
	PopulateRFC4523MatchingRules(DefaultLDAPSyntaxes, mrc)
	PopulateRFC4530MatchingRules(DefaultLDAPSyntaxes, mrc)

	return
}

/*
PopulateRFC2307MatchingRules only populates the provided MatchingRules (mrc) with MatchingRule definitions defined in this RFC (which is currently one (1) definition).

Definitions are imported from go-schemax/rfc2307.

A valid instance of LDAPSyntaxes that contains all referenced LDAPSyntax instances must be provided as the lsc argument.
*/
func PopulateRFC2307MatchingRules(lsc LDAPSyntaxCollection, mrc MatchingRuleCollection) {
	for _, mr := range rfc2307.AllMatchingRules {
		var nmr MatchingRule
		if err := Marshal(string(mr), &nmr, nil, nil, lsc, mrc, nil, nil, nil, nil); err != nil {
			panic(sprintf("%s: %s", err.Error(), string(mr)))
		}
		mrc.Set(&nmr)
	}
}

/*
PopulateRFC4517MatchingRules only populates the provided MatchingRules (mrc) with MatchingRule definitions defined in this RFC.

Definitions are imported from go-schemax/rfc4517.

A valid instance of LDAPSyntaxes that contains all referenced LDAPSyntax instances must be provided as the lsc argument.
*/
func PopulateRFC4517MatchingRules(lsc LDAPSyntaxCollection, mrc MatchingRuleCollection) {
	for _, mr := range rfc4517.AllMatchingRules {
		var nmr MatchingRule
		if err := Marshal(string(mr), &nmr, nil, nil, lsc, mrc, nil, nil, nil, nil); err != nil {
			panic(sprintf("%s: %s", err.Error(), string(mr)))
		}
		mrc.Set(&nmr)
	}
}

/*
PopulateRFC4530MatchingRules only populates the provided MatchingRules (mrc) with MatchingRule definitions defined in this RFC.

Definitions are imported from go-schemax/rfc4530.

A valid instance of LDAPSyntaxes that contains all referenced LDAPSyntax instances must be provided as the lsc argument.
*/
func PopulateRFC4530MatchingRules(lsc LDAPSyntaxCollection, mrc MatchingRuleCollection) {
	for _, mr := range rfc4530.AllMatchingRules {
		var nmr MatchingRule
		if err := Marshal(string(mr), &nmr, nil, nil, lsc, mrc, nil, nil, nil, nil); err != nil {
			panic(sprintf("%s: %s", err.Error(), string(mr)))
		}
		mrc.Set(&nmr)
	}
}

/*
PopulateRFC4523MatchingRules only populates the provided MatchingRules (mrc) with MatchingRule definitions defined in this RFC.

Definitions are imported from go-schemax/rfc4523.

A valid instance of LDAPSyntaxes that contains all referenced LDAPSyntax instances must be provided as the lsc argument.
*/
func PopulateRFC4523MatchingRules(lsc LDAPSyntaxCollection, mrc MatchingRuleCollection) {
	for _, mr := range rfc4523.AllMatchingRules {
		var nmr MatchingRule
		if err := Marshal(string(mr), &nmr, nil, nil, lsc, mrc, nil, nil, nil, nil); err != nil {
			panic(sprintf("%s: %s", err.Error(), string(mr)))
		}
		mrc.Set(&nmr)
	}
}

/////////////////
// Object Classes

/*
PopulateDefaultObjectClasses returns a new auto-populated instance of ObjectClasses.

Within this structure are all of the following:

• Object classes from RFC2079 (imported from go-schemax/rfc2079)

• Object classes from RFC2307 (imported from go-schemax/rfc2307)

• Object classes from RFC4512 (imported from go-schemax/rfc4512)

• Object classes from RFC4519 (imported from go-schemax/rfc4519)

• Object classes from RFC4523 (imported from go-schemax/rfc4523)

• Object classes from RFC4524 (imported from go-schemax/rfc4524)

• Object classes from RFC3672 (imported from go-schemax/rfc3672)
*/
func PopulateDefaultObjectClasses() ObjectClassCollection {
	occ := NewObjectClasses()

	PopulateRFC4512ObjectClasses(occ,
		DefaultAttributeTypes,
		DefaultLDAPSyntaxes,
		DefaultMatchingRules)

	PopulateRFC4519ObjectClasses(occ,
		DefaultAttributeTypes,
		DefaultLDAPSyntaxes,
		DefaultMatchingRules)

	PopulateRFC4524ObjectClasses(occ,
		DefaultAttributeTypes,
		DefaultLDAPSyntaxes,
		DefaultMatchingRules)

	PopulateRFC2798ObjectClasses(occ,
		DefaultAttributeTypes,
		DefaultLDAPSyntaxes,
		DefaultMatchingRules)

	PopulateRFC2307ObjectClasses(occ,
		DefaultAttributeTypes,
		DefaultLDAPSyntaxes,
		DefaultMatchingRules)

	PopulateRFC4523ObjectClasses(occ,
		DefaultAttributeTypes,
		DefaultLDAPSyntaxes,
		DefaultMatchingRules)

	PopulateRFC2079ObjectClasses(occ,
		DefaultAttributeTypes,
		DefaultLDAPSyntaxes,
		DefaultMatchingRules)

	PopulateRFC3672ObjectClasses(occ,
		DefaultAttributeTypes,
		DefaultLDAPSyntaxes,
		DefaultMatchingRules)

	return occ
}

/*
PopulateRFC4524ObjectClasses only populates the provided ObjectClasses (occ) with ObjectClass definitions defined in this RFC.

Definitions are imported from go-schemax/rfc4524.

Valid instances of LDAPSyntaxes (containing all referenced LDAPSyntax definitions), MatchingRules (containing all referenced MatchingRules definitions), and AttributeTypes (containing all registered AttributeType definitions) must be provided as the atc, lsc and mrc arguments respectively.
*/
func PopulateRFC4524ObjectClasses(
	occ ObjectClassCollection,
	atc AttributeTypeCollection,
	lsc LDAPSyntaxCollection,
	mrc MatchingRuleCollection) {

	for _, oc := range rfc4524.AllObjectClasses {
		var noc ObjectClass
		if err := Marshal(string(oc), &noc, atc, occ, lsc, mrc, nil, nil, nil, nil); err != nil {
			panic(sprintf("%s: %s", err.Error(), string(oc)))
		}
		occ.Set(&noc)
	}
}

/*
PopulateRFC2079ObjectClasses only populates the provided ObjectClasses (occ) with ObjectClass definitions defined in this RFC.

Definitions are imported from go-schemax/rfc2079.

Valid instances of LDAPSyntaxes (containing all referenced LDAPSyntax definitions), MatchingRules (containing all referenced MatchingRules definitions), and AttributeTypes (containing all registered AttributeType definitions) must be provided as the atc, lsc and mrc arguments respectively.
*/
func PopulateRFC2079ObjectClasses(
	occ ObjectClassCollection,
	atc AttributeTypeCollection,
	lsc LDAPSyntaxCollection,
	mrc MatchingRuleCollection) {

	for _, oc := range rfc2079.AllObjectClasses {
		var noc ObjectClass
		if err := Marshal(string(oc), &noc, atc, occ, lsc, mrc, nil, nil, nil, nil); err != nil {
			panic(sprintf("%s: %s", err.Error(), string(oc)))
		}
		occ.Set(&noc)
	}
}

/*
PopulateRFC3672ObjectClasses only populates the provided ObjectClasses (occ) with ObjectClass definitions defined in this RFC.

Definitions are imported from go-schemax/rfc3672.

Valid instances of LDAPSyntaxes (containing all referenced LDAPSyntax definitions), MatchingRules (containing all referenced MatchingRules definitions), and AttributeTypes (containing all registered AttributeType definitions) must be provided as the atc, lsc and mrc arguments respectively.
*/
func PopulateRFC3672ObjectClasses(
	occ ObjectClassCollection,
	atc AttributeTypeCollection,
	lsc LDAPSyntaxCollection,
	mrc MatchingRuleCollection) {

	for _, oc := range rfc3672.AllObjectClasses {
		var noc ObjectClass
		if err := Marshal(string(oc), &noc, atc, occ, lsc, mrc, nil, nil, nil, nil); err != nil {
			panic(sprintf("%s: %s", err.Error(), string(oc)))
		}
		occ.Set(&noc)
	}
}

/*
PopulateRFC2798ObjectClasses only populates the provided ObjectClasses (occ) with ObjectClass definitions defined in this RFC.

Definitions are imported from go-schemax/rfc2798.

Valid instances of LDAPSyntaxes (containing all referenced LDAPSyntax definitions), MatchingRules (containing all referenced MatchingRules definitions), and AttributeTypes (containing all registered AttributeType definitions) must be provided as the atc, lsc and mrc arguments respectively.
*/
func PopulateRFC2798ObjectClasses(
	occ ObjectClassCollection,
	atc AttributeTypeCollection,
	lsc LDAPSyntaxCollection,
	mrc MatchingRuleCollection) {

	for _, oc := range rfc2798.AllObjectClasses {
		var noc ObjectClass
		if err := Marshal(string(oc), &noc, atc, occ, lsc, mrc, nil, nil, nil, nil); err != nil {
			panic(sprintf("%s: %s", err.Error(), string(oc)))
		}
		occ.Set(&noc)
	}
}

/*
PopulateRFC2307ObjectClasses only populates the provided ObjectClasses (occ) with ObjectClass definitions defined in this RFC.

Definitions are imported from go-schemax/rfc2307.

Valid instances of LDAPSyntaxes (containing all referenced LDAPSyntax definitions), MatchingRules (containing all referenced MatchingRules definitions), and AttributeTypes (containing all registered AttributeType definitions) must be provided as the atc, lsc and mrc arguments respectively.
*/
func PopulateRFC2307ObjectClasses(
	occ ObjectClassCollection,
	atc AttributeTypeCollection,
	lsc LDAPSyntaxCollection,
	mrc MatchingRuleCollection) {

	for _, oc := range rfc2307.AllObjectClasses {
		var noc ObjectClass
		if err := Marshal(string(oc), &noc, atc, occ, lsc, mrc, nil, nil, nil, nil); err != nil {
			panic(sprintf("%s: %s", err.Error(), string(oc)))
		}
		occ.Set(&noc)
	}
}

/*
PopulateRFC4512ObjectClasses only populates the provided ObjectClasses (occ) with ObjectClass definitions defined in this RFC.

Definitions are imported from go-schemax/rfc4512.

Valid instances of LDAPSyntaxes (containing all referenced LDAPSyntax definitions), MatchingRules (containing all referenced MatchingRules definitions), and AttributeTypes (containing all registered AttributeType definitions) must be provided as the atc, lsc, and mrc arguments respectively.
*/
func PopulateRFC4512ObjectClasses(
	occ ObjectClassCollection,
	atc AttributeTypeCollection,
	lsc LDAPSyntaxCollection,
	mrc MatchingRuleCollection) {

	for _, oc := range rfc4512.AllObjectClasses {
		var noc ObjectClass
		if err := Marshal(string(oc), &noc, atc, occ, lsc, mrc, nil, nil, nil, nil); err != nil {
			panic(sprintf("%s: %s", err.Error(), string(oc)))
		}
		occ.Set(&noc)
	}
}

/*
PopulateRFC4519ObjectClasses only populates the provided ObjectClasses (occ) with ObjectClass definitions defined in this RFC.

Definitions are imported from go-schemax/rfc4519.

Valid instances of LDAPSyntaxes (containing all referenced LDAPSyntax definitions), MatchingRules (containing all referenced MatchingRules definitions), and AttributeTypes (containing all registered AttributeType definitions) must be provided as the atc, lsc, and mrc arguments respectively.
*/
func PopulateRFC4519ObjectClasses(
	occ ObjectClassCollection,
	atc AttributeTypeCollection,
	lsc LDAPSyntaxCollection,
	mrc MatchingRuleCollection) {

	for _, oc := range rfc4519.AllObjectClasses {
		var noc ObjectClass
		if err := Marshal(string(oc), &noc, atc, occ, lsc, mrc, nil, nil, nil, nil); err != nil {
			panic(sprintf("%s: %s", err.Error(), string(oc)))
		}
		occ.Set(&noc)
	}
}

/*
PopulateRFC4523ObjectClasses only populates the provided ObjectClasses (occ) with ObjectClass definitions defined in this RFC.

Definitions are imported from go-schemax/rfc4523.

Valid instances of LDAPSyntaxes (containing all referenced LDAPSyntax definitions), MatchingRules (containing all referenced MatchingRules definitions), and AttributeTypes (containing all registered AttributeType definitions) must be provided as the atc, lsc, and mrc arguments respectively.
*/
func PopulateRFC4523ObjectClasses(
	occ ObjectClassCollection,
	atc AttributeTypeCollection,
	lsc LDAPSyntaxCollection,
	mrc MatchingRuleCollection) {

	for _, oc := range rfc4523.AllObjectClasses {
		var noc ObjectClass
		if err := Marshal(string(oc), &noc, atc, occ, lsc, mrc, nil, nil, nil, nil); err != nil {
			panic(sprintf("%s: %s", err.Error(), string(oc)))
		}
		occ.Set(&noc)
	}
}

//////////////////
// Attribute Types

/*
PopulateDefaultAttributeTypes returns a new auto-populated instance of AttributeTypes. Within this structure are all of the following:

• Attribute types from RFC2079 (imported from go-schemax/rfc2079)

• Attribute types from RFC2307 (imported from go-schemax/rfc2307)

• Attribute types from RFC2798 (imported from go-schemax/rfc2798)

• Attribute types from RFC3045 (imported from go-schemax/rfc3045)

• Attribute types from RFC4512 (imported from go-schemax/rfc4512)

• Attribute types from RFC4519 (imported from go-schemax/rfc4519)

• Attribute types from RFC4523 (imported from go-schemax/rfc4523)

• Attribute types from RFC4524 (imported from go-schemax/rfc4524)

• Attribute types from RFC4530 (imported from go-schemax/rfc4530)

• Attribute types from RFC3672 (imported from go-schemax/rfc3672)
*/
func PopulateDefaultAttributeTypes() (atc AttributeTypeCollection) {
	atc = NewAttributeTypes()

	PopulateRFC4512AttributeTypes(atc,
		DefaultLDAPSyntaxes,
		DefaultMatchingRules)

	PopulateRFC4519AttributeTypes(atc,
		DefaultLDAPSyntaxes,
		DefaultMatchingRules)

	PopulateRFC4524AttributeTypes(atc,
		DefaultLDAPSyntaxes,
		DefaultMatchingRules)

	PopulateRFC4530AttributeTypes(atc,
		DefaultLDAPSyntaxes,
		DefaultMatchingRules)

	PopulateRFC2798AttributeTypes(atc,
		DefaultLDAPSyntaxes,
		DefaultMatchingRules)

	PopulateRFC2307AttributeTypes(atc,
		DefaultLDAPSyntaxes,
		DefaultMatchingRules)

	PopulateRFC4523AttributeTypes(atc,
		DefaultLDAPSyntaxes,
		DefaultMatchingRules)

	PopulateRFC2079AttributeTypes(atc,
		DefaultLDAPSyntaxes,
		DefaultMatchingRules)

	PopulateRFC3672AttributeTypes(atc,
		DefaultLDAPSyntaxes,
		DefaultMatchingRules)

        PopulateRFC3045AttributeTypes(atc,
                DefaultLDAPSyntaxes,
                DefaultMatchingRules)

	return
}

/*
PopulateRFC4524AttributeTypes only populates the provided AttributeTypes (atc) with AttributeType definitions defined in this RFC.

Definitions are imported from go-schemax/rfc4524.

Valid instances of LDAPSyntaxes (containing all referenced LDAPSyntax definitions) and MatchingRules (containing all referenced MatchingRules definitions), must be provided as the lsc and mrc arguments respectively.
*/
func PopulateRFC4524AttributeTypes(
	atc AttributeTypeCollection,
	lsc LDAPSyntaxCollection,
	mrc MatchingRuleCollection) {

	for _, at := range rfc4524.AllAttributeTypes {
		var nat AttributeType
		if err := Marshal(string(at), &nat, atc, nil, lsc, mrc, nil, nil, nil, nil); err != nil {
			panic(sprintf("%s: %s", err.Error(), string(at)))
		}
		atc.Set(&nat)
	}
}

/*
PopulateRFC2079AttributeTypes only populates the provided AttributeTypes (atc) with AttributeType definitions defined in this RFC.

Definitions are imported from go-schemax/rfc2079.

Valid instances of LDAPSyntaxes (containing all referenced LDAPSyntax definitions) and MatchingRules (containing all referenced MatchingRules definitions), must be provided as the lsc and mrc arguments respectively.
*/
func PopulateRFC2079AttributeTypes(
	atc AttributeTypeCollection,
	lsc LDAPSyntaxCollection,
	mrc MatchingRuleCollection) {

	for _, at := range rfc2079.AllAttributeTypes {
		var nat AttributeType
		if err := Marshal(string(at), &nat, atc, nil, lsc, mrc, nil, nil, nil, nil); err != nil {
			panic(sprintf("%s: %s", err.Error(), string(at)))
		}
		atc.Set(&nat)
	}
}

/*
PopulateRFC3045AttributeTypes only populates the provided AttributeTypes (atc) with AttributeType definitions defined in this RFC.

Definitions are imported from go-schemax/rfc3045.

Valid instances of LDAPSyntaxes (containing all referenced LDAPSyntax definitions) and MatchingRules (containing all referenced MatchingRules definitions), must be provided as the lsc and mrc arguments respectively.
*/
func PopulateRFC3045AttributeTypes(
        atc AttributeTypeCollection,
        lsc LDAPSyntaxCollection,
        mrc MatchingRuleCollection) {

        for _, at := range rfc3045.AllAttributeTypes {
                var nat AttributeType
                if err := Marshal(string(at), &nat, atc, nil, lsc, mrc, nil, nil, nil, nil); err != nil {
                        panic(sprintf("%s: %s", err.Error(), string(at)))
                }
                atc.Set(&nat)
        }
}

/*
PopulateRFC3672AttributeTypes only populates the provided AttributeTypes (atc) with AttributeType definitions defined in this RFC.

Definitions are imported from go-schemax/rfc3672.

Valid instances of LDAPSyntaxes (containing all referenced LDAPSyntax definitions) and MatchingRules (containing all referenced MatchingRules definitions), must be provided as the lsc and mrc arguments respectively.
*/
func PopulateRFC3672AttributeTypes(
	atc AttributeTypeCollection,
	lsc LDAPSyntaxCollection,
	mrc MatchingRuleCollection) {

	for _, at := range rfc3672.AllAttributeTypes {
		var nat AttributeType
		if err := Marshal(string(at), &nat, atc, nil, lsc, mrc, nil, nil, nil, nil); err != nil {
			panic(sprintf("%s: %s", err.Error(), string(at)))
		}
		atc.Set(&nat)
	}
}

/*
PopulateRFC2307AttributeTypes only populates the provided AttributeTypes (atc) with AttributeType definitions defined in this RFC.

Definitions are imported from go-schemax/rfc2307.

Valid instances of LDAPSyntaxes (containing all referenced LDAPSyntax definitions) and MatchingRules (containing all referenced MatchingRules definitions), must be provided as the lsc and mrc arguments respectively.
*/
func PopulateRFC2307AttributeTypes(
	atc AttributeTypeCollection,
	lsc LDAPSyntaxCollection,
	mrc MatchingRuleCollection) {

	for _, at := range rfc2307.AllAttributeTypes {
		var nat AttributeType
		if err := Marshal(string(at), &nat, atc, nil, lsc, mrc, nil, nil, nil, nil); err != nil {
			panic(sprintf("%s: %s", err.Error(), string(at)))
		}
		atc.Set(&nat)
	}
}

/*
PopulateRFC2798AttributeTypes only populates the provided AttributeTypes (atc) with AttributeType definitions defined in this RFC.

Definitions are imported from go-schemax/rfc2798.

Valid instances of LDAPSyntaxes (containing all referenced LDAPSyntax definitions) and MatchingRules (containing all referenced MatchingRules definitions), must be provided as the lsc and mrc arguments respectively.
*/
func PopulateRFC2798AttributeTypes(
	atc AttributeTypeCollection,
	lsc LDAPSyntaxCollection,
	mrc MatchingRuleCollection) {

	for _, at := range rfc2798.AllAttributeTypes {
		var nat AttributeType
		if err := Marshal(string(at), &nat, atc, nil, lsc, mrc, nil, nil, nil, nil); err != nil {
			panic(sprintf("%s: %s", err.Error(), string(at)))
		}
		atc.Set(&nat)
	}
}

/*
PopulateRFC4512AttributeTypes only populates the provided AttributeTypes (atc) with AttributeType definitions defined in this RFC.

Definitions are imported from go-schemax/rfc4512.

Valid instances of LDAPSyntaxes (containing all referenced LDAPSyntax definitions) and MatchingRules (containing all referenced MatchingRules definitions), must be provided as the lsc and mrc arguments respectively.
*/
func PopulateRFC4512AttributeTypes(
	atc AttributeTypeCollection,
	lsc LDAPSyntaxCollection,
	mrc MatchingRuleCollection) {

	for _, at := range rfc4512.AllAttributeTypes {
		var nat AttributeType
		if err := Marshal(string(at), &nat, atc, nil, lsc, mrc, nil, nil, nil, nil); err != nil {
			panic(sprintf("%s: %s", err.Error(), string(at)))
		}
		atc.Set(&nat)
	}
}

/*
PopulateRFC4519AttributeTypes only populates the provided AttributeTypes (atc) with AttributeType definitions defined in this RFC.

Definitions are imported from go-schemax/rfc4519.

Valid instances of LDAPSyntaxes (containing all referenced LDAPSyntax definitions) and MatchingRules (containing all referenced MatchingRules definitions), must be provided as the lsc and mrc arguments respectively.
*/
func PopulateRFC4519AttributeTypes(
	atc AttributeTypeCollection,
	lsc LDAPSyntaxCollection,
	mrc MatchingRuleCollection) {

	for _, at := range rfc4519.AllAttributeTypes {
		var nat AttributeType
		if err := Marshal(string(at), &nat, atc, nil, lsc, mrc, nil, nil, nil, nil); err != nil {
			panic(sprintf("%s: %s", err.Error(), string(at)))
		}
		atc.Set(&nat)
	}
}

/*
PopulateRFC4523AttributeTypes only populates the provided AttributeTypes (atc) with AttributeType definitions defined in this RFC.

Definitions are imported from go-schemax/rfc4523.

Valid instances of LDAPSyntaxes (containing all referenced LDAPSyntax definitions) and MatchingRules (containing all referenced MatchingRules definitions), must be provided as the lsc and mrc arguments respectively.
*/
func PopulateRFC4523AttributeTypes(
	atc AttributeTypeCollection,
	lsc LDAPSyntaxCollection,
	mrc MatchingRuleCollection) {

	for _, at := range rfc4523.AllAttributeTypes {
		var nat AttributeType
		if err := Marshal(string(at), &nat, atc, nil, lsc, mrc, nil, nil, nil, nil); err != nil {
			panic(sprintf("%s: %s", err.Error(), string(at)))
		}
		atc.Set(&nat)
	}
}

/*
PopulateRFC4530AttributeTypes only populates the provided AttributeTypes (atc) with AttributeType definitions defined in this RFC.

Definitions are imported from go-schemax/rfc4530.

Valid instances of LDAPSyntaxes (containing all referenced LDAPSyntax definitions) and MatchingRules (containing all referenced MatchingRules definitions), must be provided as the lsc and mrc arguments respectively.
*/
func PopulateRFC4530AttributeTypes(
	atc AttributeTypeCollection,
	lsc LDAPSyntaxCollection,
	mrc MatchingRuleCollection) {

	for _, at := range rfc4530.AllAttributeTypes {
		var nat AttributeType
		if err := Marshal(string(at), &nat, atc, nil, lsc, mrc, nil, nil, nil, nil); err != nil {
			panic(sprintf("%s: %s", err.Error(), string(at)))
		}
		atc.Set(&nat)
	}
}

func init() {

	// default collections
	DefaultLDAPSyntaxes = PopulateDefaultLDAPSyntaxes()
	DefaultMatchingRules = PopulateDefaultMatchingRules()
	DefaultAttributeTypes = PopulateDefaultAttributeTypes()
	DefaultObjectClasses = PopulateDefaultObjectClasses()

}

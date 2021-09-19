package schemax

/*
init.go provided initialization routines for certain global variables and recommended content.
*/

import (
	"github.com/JesseCoretta/go-schemax/rfc2079"
	"github.com/JesseCoretta/go-schemax/rfc2307"
	"github.com/JesseCoretta/go-schemax/rfc2798"
	"github.com/JesseCoretta/go-schemax/rfc4512"
	"github.com/JesseCoretta/go-schemax/rfc4517"
	"github.com/JesseCoretta/go-schemax/rfc4519"
	"github.com/JesseCoretta/go-schemax/rfc4523"
	"github.com/JesseCoretta/go-schemax/rfc4524"
)

////////////////
// LDAP Syntaxes

/*
PopulateDefaultLDAPSyntaxesManifest returns a new auto-populated instance of LDAPSyntaxesManifest.

Within this structure are all of the following:

 - Syntax definitions from RFC2307 (imported from go-schemax/rfc2307)
 - Syntax definitions from RFC4517 (imported from go-schemax/rfc4517)
 - Syntax definitions from RFC4523 (imported from go-schemax/rfc4523)
*/
func PopulateDefaultLDAPSyntaxesManifest() (lsm LDAPSyntaxesManifest) {
	lsm = NewLDAPSyntaxesManifest()
	PopulateRFC2307Syntaxes(lsm)
	PopulateRFC4517Syntaxes(lsm)
	PopulateRFC4523Syntaxes(lsm)

	return
}

/*
PopulateRFC4523Syntaxes only populates the provided LDAPSyntaxesManifest (lsm) with LDAPSyntax definitions defined in this RFC.

Definitions are imported from go-schemax/rfc4523.
*/
func PopulateRFC4523Syntaxes(lsm LDAPSyntaxesManifest) {
	for _, ls := range rfc4523.AllLDAPSyntaxes {
		var nls LDAPSyntax
		if err := Marshal(string(ls), &nls, nil, nil, nil, nil, nil, nil, nil, nil, nil); err != nil {
			panic(err)
		}
		lsm.Set(&nls)
	}
}

/*
PopulateRFC2307Syntaxes only populates the provided LDAPSyntaxesManifest (lsm) with LDAPSyntax definitions defined in this RFC.

Definitions are imported from go-schemax/rfc2307.
*/
func PopulateRFC2307Syntaxes(lsm LDAPSyntaxesManifest) {
	for _, ls := range rfc2307.AllLDAPSyntaxes {
		var nls LDAPSyntax
		if err := Marshal(string(ls), &nls, nil, nil, nil, nil, nil, nil, nil, nil, nil); err != nil {
			panic(err)
		}
		lsm.Set(&nls)
	}
}

/*
PopulateRFC4517Syntaxes only populates the provided LDAPSyntaxesManifest (lsm) with LDAPSyntax definitions defined in this RFC.

Definitions are imported from go-schemax/rfc4517.
*/
func PopulateRFC4517Syntaxes(lsm LDAPSyntaxesManifest) {
	for _, ls := range rfc4517.AllLDAPSyntaxes {
		var nls LDAPSyntax
		if err := Marshal(string(ls), &nls, nil, nil, nil, nil, nil, nil, nil, nil, nil); err != nil {
			panic(sprintf("%s: %s", err.Error(), string(ls)))
		}
		lsm.Set(&nls)
	}
}

/////////////////
// Matching Rules

/*
PopulateDefaultMatchingRulesManifest returns a new auto-populated instance of MatchingRulesManifest. Within this structure are all of the following:

 - Matching rules from RFC4517 (imported from go-schemax/rfc4517)
 - Matching rules from RFC4523 (imported from go-schemax/rfc4523)
*/
func PopulateDefaultMatchingRulesManifest() (mrm MatchingRulesManifest) {
	mrm = NewMatchingRulesManifest()
	PopulateRFC2307MatchingRules(DefaultLDAPSyntaxesManifest, mrm)
	PopulateRFC4517MatchingRules(DefaultLDAPSyntaxesManifest, mrm)
	PopulateRFC4523MatchingRules(DefaultLDAPSyntaxesManifest, mrm)

	return
}

/*
PopulateRFC2307MatchingRules only populates the provided MatchingRulesManifest (mrm) with MatchingRule definitions defined in this RFC (which is currently one (1) definition).

Definitions are imported from go-schemax/rfc2307.

A valid instance of LDAPSyntaxesManifest that contains all referenced LDAPSyntax instances must be provided as the lsm argument.
*/
func PopulateRFC2307MatchingRules(lsm LDAPSyntaxesManifest, mrm MatchingRulesManifest) {
	for _, mr := range rfc2307.AllMatchingRules {
		var nmr MatchingRule
		if err := Marshal(string(mr), &nmr, nil, nil, nil, lsm, mrm, nil, nil, nil, nil); err != nil {
			panic(sprintf("%s: %s", err.Error(), string(mr)))
		}
		mrm.Set(&nmr)
	}
}

/*
PopulateRFC4517MatchingRules only populates the provided MatchingRulesManifest (mrm) with MatchingRule definitions defined in this RFC.

Definitions are imported from go-schemax/rfc4517.

A valid instance of LDAPSyntaxesManifest that contains all referenced LDAPSyntax instances must be provided as the lsm argument.
*/
func PopulateRFC4517MatchingRules(lsm LDAPSyntaxesManifest, mrm MatchingRulesManifest) {
	for _, mr := range rfc4517.AllMatchingRules {
		var nmr MatchingRule
		if err := Marshal(string(mr), &nmr, nil, nil, nil, lsm, mrm, nil, nil, nil, nil); err != nil {
			panic(sprintf("%s: %s", err.Error(), string(mr)))
		}
		mrm.Set(&nmr)
	}
}

/*
PopulateRFC4523MatchingRules only populates the provided MatchingRulesManifest (mrm) with MatchingRule definitions defined in this RFC.

Definitions are imported from go-schemax/rfc4523.

A valid instance of LDAPSyntaxesManifest that contains all referenced LDAPSyntax instances must be provided as the lsm argument.
*/
func PopulateRFC4523MatchingRules(lsm LDAPSyntaxesManifest, mrm MatchingRulesManifest) {
	for _, mr := range rfc4523.AllMatchingRules {
		var nmr MatchingRule
		if err := Marshal(string(mr), &nmr, nil, nil, nil, lsm, mrm, nil, nil, nil, nil); err != nil {
			panic(sprintf("%s: %s", err.Error(), string(mr)))
		}
		mrm.Set(&nmr)
	}
}

/////////////////
// Object Classes

/*
PopulateDefaultObjectClassesManifest returns a new auto-populated instance of ObjectClassesManifest.

Within this structure are all of the following:

 - Object classes from RFC2079 (imported from go-schemax/rfc2079)
 - Object classes from RFC2307 (imported from go-schemax/rfc2307)
 - Object classes from RFC4512 (imported from go-schemax/rfc4512)
 - Object classes from RFC4519 (imported from go-schemax/rfc4519)
 - Object classes from RFC4523 (imported from go-schemax/rfc4523)
 - Object classes from RFC4524 (imported from go-schemax/rfc4524)
*/
func PopulateDefaultObjectClassesManifest() (ocm ObjectClassesManifest) {
	ocm = NewObjectClassesManifest()

	PopulateRFC4512ObjectClasses(ocm,
		DefaultAttributeTypesManifest,
		DefaultLDAPSyntaxesManifest,
		DefaultMatchingRulesManifest)

	PopulateRFC4519ObjectClasses(ocm,
		DefaultAttributeTypesManifest,
		DefaultLDAPSyntaxesManifest,
		DefaultMatchingRulesManifest)

	PopulateRFC4524ObjectClasses(ocm,
		DefaultAttributeTypesManifest,
		DefaultLDAPSyntaxesManifest,
		DefaultMatchingRulesManifest)

	PopulateRFC2798ObjectClasses(ocm,
		DefaultAttributeTypesManifest,
		DefaultLDAPSyntaxesManifest,
		DefaultMatchingRulesManifest)

	PopulateRFC2307ObjectClasses(ocm,
		DefaultAttributeTypesManifest,
		DefaultLDAPSyntaxesManifest,
		DefaultMatchingRulesManifest)

	PopulateRFC4523ObjectClasses(ocm,
		DefaultAttributeTypesManifest,
		DefaultLDAPSyntaxesManifest,
		DefaultMatchingRulesManifest)

	PopulateRFC2079ObjectClasses(ocm,
		DefaultAttributeTypesManifest,
		DefaultLDAPSyntaxesManifest,
		DefaultMatchingRulesManifest)

	return
}

/*
PopulateRFC4524ObjectClasses only populates the provided ObjectClassesManifest (ocm) with ObjectClass definitions defined in this RFC.

Definitions are imported from go-schemax/rfc4524.

Valid instances of LDAPSyntaxesManifest (containing all referenced LDAPSyntax definitions), MatchingRulesManifest (containing all referenced MatchingRules definitions), and AttributeTypesManifest (containing all registered AttributeType definitions) must be provided as the atm, lsm and mrm arguments respectively.
*/
func PopulateRFC4524ObjectClasses(
	ocm ObjectClassesManifest,
	atm AttributeTypesManifest,
	lsm LDAPSyntaxesManifest,
	mrm MatchingRulesManifest) {

	for _, oc := range rfc4524.AllObjectClasses {
		var noc ObjectClass
		if err := Marshal(string(oc), &noc, nil, atm, ocm, lsm, mrm, nil, nil, nil, nil); err != nil {
			panic(sprintf("%s: %s", err.Error(), string(oc)))
		}
		ocm.Set(&noc)
	}
}

/*
PopulateRFC2079ObjectClasses only populates the provided ObjectClassesManifest (ocm) with ObjectClass definitions defined in this RFC.

Definitions are imported from go-schemax/rfc2079.

Valid instances of LDAPSyntaxesManifest (containing all referenced LDAPSyntax definitions), MatchingRulesManifest (containing all referenced MatchingRules definitions), and AttributeTypesManifest (containing all registered AttributeType definitions) must be provided as the atm, lsm and mrm arguments respectively.
*/
func PopulateRFC2079ObjectClasses(
	ocm ObjectClassesManifest,
	atm AttributeTypesManifest,
	lsm LDAPSyntaxesManifest,
	mrm MatchingRulesManifest) {

	for _, oc := range rfc2079.AllObjectClasses {
		var noc ObjectClass
		if err := Marshal(string(oc), &noc, nil, atm, ocm, lsm, mrm, nil, nil, nil, nil); err != nil {
			panic(sprintf("%s: %s", err.Error(), string(oc)))
		}
		ocm.Set(&noc)
	}
}

/*
PopulateRFC2798ObjectClasses only populates the provided ObjectClassesManifest (ocm) with ObjectClass definitions defined in this RFC.

Definitions are imported from go-schemax/rfc2798.

Valid instances of LDAPSyntaxesManifest (containing all referenced LDAPSyntax definitions), MatchingRulesManifest (containing all referenced MatchingRules definitions), and AttributeTypesManifest (containing all registered AttributeType definitions) must be provided as the atm, lsm and mrm arguments respectively.
*/
func PopulateRFC2798ObjectClasses(
	ocm ObjectClassesManifest,
	atm AttributeTypesManifest,
	lsm LDAPSyntaxesManifest,
	mrm MatchingRulesManifest) {

	for _, oc := range rfc2798.AllObjectClasses {
		var noc ObjectClass
		if err := Marshal(string(oc), &noc, nil, atm, ocm, lsm, mrm, nil, nil, nil, nil); err != nil {
			panic(sprintf("%s: %s", err.Error(), string(oc)))
		}
		ocm.Set(&noc)
	}
}

/*
PopulateRFC2307ObjectClasses only populates the provided ObjectClassesManifest (ocm) with ObjectClass definitions defined in this RFC.

Definitions are imported from go-schemax/rfc2307.

Valid instances of LDAPSyntaxesManifest (containing all referenced LDAPSyntax definitions), MatchingRulesManifest (containing all referenced MatchingRules definitions), and AttributeTypesManifest (containing all registered AttributeType definitions) must be provided as the atm, lsm and mrm arguments respectively.
*/
func PopulateRFC2307ObjectClasses(
	ocm ObjectClassesManifest,
	atm AttributeTypesManifest,
	lsm LDAPSyntaxesManifest,
	mrm MatchingRulesManifest) {

	for _, oc := range rfc2307.AllObjectClasses {
		var noc ObjectClass
		if err := Marshal(string(oc), &noc, nil, atm, ocm, lsm, mrm, nil, nil, nil, nil); err != nil {
			panic(sprintf("%s: %s", err.Error(), string(oc)))
		}
		ocm.Set(&noc)
	}
}

/*
PopulateRFC4512ObjectClasses only populates the provided ObjectClassesManifest (ocm) with ObjectClass definitions defined in this RFC.

Definitions are imported from go-schemax/rfc4512.

Valid instances of LDAPSyntaxesManifest (containing all referenced LDAPSyntax definitions), MatchingRulesManifest (containing all referenced MatchingRules definitions), and AttributeTypesManifest (containing all registered AttributeType definitions) must be provided as the atm, lsm, and mrm arguments respectively.
*/
func PopulateRFC4512ObjectClasses(
	ocm ObjectClassesManifest,
	atm AttributeTypesManifest,
	lsm LDAPSyntaxesManifest,
	mrm MatchingRulesManifest) {

	for _, oc := range rfc4512.AllObjectClasses {
		var noc ObjectClass
		if err := Marshal(string(oc), &noc, nil, atm, ocm, lsm, mrm, nil, nil, nil, nil); err != nil {
			panic(sprintf("%s: %s", err.Error(), string(oc)))
		}
		ocm.Set(&noc)
	}
}

/*
PopulateRFC4519ObjectClasses only populates the provided ObjectClassesManifest (ocm) with ObjectClass definitions defined in this RFC.

Definitions are imported from go-schemax/rfc4519.

Valid instances of LDAPSyntaxesManifest (containing all referenced LDAPSyntax definitions), MatchingRulesManifest (containing all referenced MatchingRules definitions), and AttributeTypesManifest (containing all registered AttributeType definitions) must be provided as the atm, lsm, and mrm arguments respectively.
*/
func PopulateRFC4519ObjectClasses(
	ocm ObjectClassesManifest,
	atm AttributeTypesManifest,
	lsm LDAPSyntaxesManifest,
	mrm MatchingRulesManifest) {

	for _, oc := range rfc4519.AllObjectClasses {
		var noc ObjectClass
		if err := Marshal(string(oc), &noc, nil, atm, ocm, lsm, mrm, nil, nil, nil, nil); err != nil {
			panic(sprintf("%s: %s", err.Error(), string(oc)))
		}
		ocm.Set(&noc)
	}
}

/*
PopulateRFC4523ObjectClasses only populates the provided ObjectClassesManifest (ocm) with ObjectClass definitions defined in this RFC.

Definitions are imported from go-schemax/rfc4523.

Valid instances of LDAPSyntaxesManifest (containing all referenced LDAPSyntax definitions), MatchingRulesManifest (containing all referenced MatchingRules definitions), and AttributeTypesManifest (containing all registered AttributeType definitions) must be provided as the atm, lsm, and mrm arguments respectively.
*/
func PopulateRFC4523ObjectClasses(
	ocm ObjectClassesManifest,
	atm AttributeTypesManifest,
	lsm LDAPSyntaxesManifest,
	mrm MatchingRulesManifest) {

	for _, oc := range rfc4523.AllObjectClasses {
		var noc ObjectClass
		if err := Marshal(string(oc), &noc, nil, atm, ocm, lsm, mrm, nil, nil, nil, nil); err != nil {
			panic(sprintf("%s: %s", err.Error(), string(oc)))
		}
		ocm.Set(&noc)
	}
}

//////////////////
// Attribute Types

/*
PopulateDefaultAttributeTypesManifest returns a new auto-populated instance of AttributeTypesManifest. Within this structure are all of the following:

 - Attribute types from RFC2079 (imported from go-schemax/rfc2079)
 - Attribute types from RFC2307 (imported from go-schemax/rfc2307)
 - Attribute types from RFC2798 (imported from go-schemax/rfc2798)
 - Attribute types from RFC4512 (imported from go-schemax/rfc4512)
 - Attribute types from RFC4519 (imported from go-schemax/rfc4519)
 - Attribute types from RFC4523 (imported from go-schemax/rfc4523)
 - Attribute types from RFC4524 (imported from go-schemax/rfc4524)
*/
func PopulateDefaultAttributeTypesManifest() (atm AttributeTypesManifest) {
	atm = NewAttributeTypesManifest()

	PopulateRFC4512AttributeTypes(atm,
		DefaultLDAPSyntaxesManifest,
		DefaultMatchingRulesManifest)

	PopulateRFC4519AttributeTypes(atm,
		DefaultLDAPSyntaxesManifest,
		DefaultMatchingRulesManifest)

	PopulateRFC4524AttributeTypes(atm,
		DefaultLDAPSyntaxesManifest,
		DefaultMatchingRulesManifest)

	PopulateRFC2798AttributeTypes(atm,
		DefaultLDAPSyntaxesManifest,
		DefaultMatchingRulesManifest)

	PopulateRFC2307AttributeTypes(atm,
		DefaultLDAPSyntaxesManifest,
		DefaultMatchingRulesManifest)

	PopulateRFC4523AttributeTypes(atm,
		DefaultLDAPSyntaxesManifest,
		DefaultMatchingRulesManifest)

	PopulateRFC2079AttributeTypes(atm,
		DefaultLDAPSyntaxesManifest,
		DefaultMatchingRulesManifest)

	return
}

/*
PopulateRFC4524AttributeTypes only populates the provided AttributeTypesManifest (atm) with AttributeType definitions defined in this RFC.

Definitions are imported from go-schemax/rfc4524.

Valid instances of LDAPSyntaxesManifest (containing all referenced LDAPSyntax definitions) and MatchingRulesManifest (containing all referenced MatchingRules definitions), must be provided as the lsm and mrm arguments respectively.
*/
func PopulateRFC4524AttributeTypes(
	atm AttributeTypesManifest,
	lsm LDAPSyntaxesManifest,
	mrm MatchingRulesManifest) {

	for _, at := range rfc4524.AllAttributeTypes {
		var nat AttributeType
		if err := Marshal(string(at), &nat, nil, atm, nil, lsm, mrm, nil, nil, nil, nil); err != nil {
			panic(sprintf("%s: %s", err.Error(), string(at)))
		}
		atm.Set(&nat)
	}
}

/*
PopulateRFC2079AttributeTypes only populates the provided AttributeTypesManifest (atm) with AttributeType definitions defined in this RFC.

Definitions are imported from go-schemax/rfc2079.

Valid instances of LDAPSyntaxesManifest (containing all referenced LDAPSyntax definitions) and MatchingRulesManifest (containing all referenced MatchingRules definitions), must be provided as the lsm and mrm arguments respectively.
*/
func PopulateRFC2079AttributeTypes(
	atm AttributeTypesManifest,
	lsm LDAPSyntaxesManifest,
	mrm MatchingRulesManifest) {

	for _, at := range rfc2079.AllAttributeTypes {
		var nat AttributeType
		if err := Marshal(string(at), &nat, nil, atm, nil, lsm, mrm, nil, nil, nil, nil); err != nil {
			panic(sprintf("%s: %s", err.Error(), string(at)))
		}
		atm.Set(&nat)
	}
}

/*
PopulateRFC2307AttributeTypes only populates the provided AttributeTypesManifest (atm) with AttributeType definitions defined in this RFC.

Definitions are imported from go-schemax/rfc2307.

Valid instances of LDAPSyntaxesManifest (containing all referenced LDAPSyntax definitions) and MatchingRulesManifest (containing all referenced MatchingRules definitions), must be provided as the lsm and mrm arguments respectively.
*/
func PopulateRFC2307AttributeTypes(
	atm AttributeTypesManifest,
	lsm LDAPSyntaxesManifest,
	mrm MatchingRulesManifest) {

	for _, at := range rfc2307.AllAttributeTypes {
		var nat AttributeType
		if err := Marshal(string(at), &nat, nil, atm, nil, lsm, mrm, nil, nil, nil, nil); err != nil {
			panic(sprintf("%s: %s", err.Error(), string(at)))
		}
		atm.Set(&nat)
	}
}

/*
PopulateRFC2798AttributeTypes only populates the provided AttributeTypesManifest (atm) with AttributeType definitions defined in this RFC.

Definitions are imported from go-schemax/rfc2798.

Valid instances of LDAPSyntaxesManifest (containing all referenced LDAPSyntax definitions) and MatchingRulesManifest (containing all referenced MatchingRules definitions), must be provided as the lsm and mrm arguments respectively.
*/
func PopulateRFC2798AttributeTypes(
	atm AttributeTypesManifest,
	lsm LDAPSyntaxesManifest,
	mrm MatchingRulesManifest) {

	for _, at := range rfc2798.AllAttributeTypes {
		var nat AttributeType
		if err := Marshal(string(at), &nat, nil, atm, nil, lsm, mrm, nil, nil, nil, nil); err != nil {
			panic(sprintf("%s: %s", err.Error(), string(at)))
		}
		atm.Set(&nat)
	}
}

/*
PopulateRFC4512AttributeTypes only populates the provided AttributeTypesManifest (atm) with AttributeType definitions defined in this RFC.

Definitions are imported from go-schemax/rfc4512.

Valid instances of LDAPSyntaxesManifest (containing all referenced LDAPSyntax definitions) and MatchingRulesManifest (containing all referenced MatchingRules definitions), must be provided as the lsm and mrm arguments respectively.
*/
func PopulateRFC4512AttributeTypes(
	atm AttributeTypesManifest,
	lsm LDAPSyntaxesManifest,
	mrm MatchingRulesManifest) {

	for _, at := range rfc4512.AllAttributeTypes {
		var nat AttributeType
		if err := Marshal(string(at), &nat, nil, atm, nil, lsm, mrm, nil, nil, nil, nil); err != nil {
			panic(sprintf("%s: %s", err.Error(), string(at)))
		}
		atm.Set(&nat)
	}
}

/*
PopulateRFC4519AttributeTypes only populates the provided AttributeTypesManifest (atm) with AttributeType definitions defined in this RFC.

Definitions are imported from go-schemax/rfc4519.

Valid instances of LDAPSyntaxesManifest (containing all referenced LDAPSyntax definitions) and MatchingRulesManifest (containing all referenced MatchingRules definitions), must be provided as the lsm and mrm arguments respectively.
*/
func PopulateRFC4519AttributeTypes(
	atm AttributeTypesManifest,
	lsm LDAPSyntaxesManifest,
	mrm MatchingRulesManifest) {

	for _, at := range rfc4519.AllAttributeTypes {
		var nat AttributeType
		if err := Marshal(string(at), &nat, nil, atm, nil, lsm, mrm, nil, nil, nil, nil); err != nil {
			panic(sprintf("%s: %s", err.Error(), string(at)))
		}
		atm.Set(&nat)
	}
}

/*
PopulateRFC4523AttributeTypes only populates the provided AttributeTypesManifest (atm) with AttributeType definitions defined in this RFC.

Definitions are imported from go-schemax/rfc4523.

Valid instances of LDAPSyntaxesManifest (containing all referenced LDAPSyntax definitions) and MatchingRulesManifest (containing all referenced MatchingRules definitions), must be provided as the lsm and mrm arguments respectively.
*/
func PopulateRFC4523AttributeTypes(
	atm AttributeTypesManifest,
	lsm LDAPSyntaxesManifest,
	mrm MatchingRulesManifest) {

	for _, at := range rfc4523.AllAttributeTypes {
		var nat AttributeType
		if err := Marshal(string(at), &nat, nil, atm, nil, lsm, mrm, nil, nil, nil, nil); err != nil {
			panic(sprintf("%s: %s", err.Error(), string(at)))
		}
		atm.Set(&nat)
	}
}

func init() {

	// default manifests
	DefaultLDAPSyntaxesManifest = PopulateDefaultLDAPSyntaxesManifest()
	DefaultMatchingRulesManifest = PopulateDefaultMatchingRulesManifest()
	DefaultAttributeTypesManifest = PopulateDefaultAttributeTypesManifest()
	DefaultObjectClassesManifest = PopulateDefaultObjectClassesManifest()

}

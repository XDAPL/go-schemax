/*
Author: Jesse Coretta - for github.com/JesseCoretta/go-schemax/_examples

This example describes a basic interaction with an LDAP DSA via the go-ldap/v3 package.

What happens:

 - Anonymous bind to DSA
 - Search of subschema subentry distinguished name
 - (Try to) obtain all known definition types:
   - attributeType
   - objectClass
   - ldapSyntaxes
   - matchingRule
   - matchingRuleUse
   - dITContentRule
   - dITStructureRule
   - nameForm
 - Initialize schemax.Subschema instance
 - Marshal schema content (if any) into schemax.Subschema instance
 - Disconnect

This example can be modified to actually do something *useful* with the produced data as opposed to simply printing it.
*/
package main

import (
	"fmt"
	"log"

	"github.com/go-ldap/ldap/v3"
	"go-schemax"
)

func main() {

	// Search params -- update as desired
	uri := `ldap://ldap.example.com/`
	scope := 0
	schemaDN := `cn=subschema`
	filter := `(objectClass=*)`
	attrs := []string{
		`ldapSyntaxes`,
		`matchingRules`,
		`attributeTypes`,
		`matchingRuleUses`,
		`objectClasses`,
		`dITContentRules`,
		`nameForms`,
		`dITStructureRules`,
	}

	// quick error handler
	chkerr := func(err error) {
		if err != nil {
			log.Fatalln(err)
		}
	}

	// CONNECT
	// -------
	dua, err := ldap.DialURL(uri)
	chkerr(err)
	defer dua.Close()

	// BIND
	// ----
	// This is an ANONYMOUS bind, which generally
	// should be sufficient to access your DSA's
	// subschema subentry. This may vary depending
	// on the security posture in your respective
	// network architecture.
	chkerr(dua.UnauthenticatedBind(``))

	// PREPARE REQUEST
	// ---------------
	// Define the parameters of our intended
	// LDAP Search Operation.
	req := &ldap.SearchRequest{
		BaseDN:     schemaDN,
		Scope:      scope,
		Filter:     filter,
		Attributes: attrs,
	}

	// SEARCH
	// ------
	// Pass our REQUEST object to the DUA
	// and get back an *ldap.SearchResults
	// instance.
	sr, err := dua.Search(req)
	chkerr(err)
	if len(sr.Entries) == 0 {
		log.Fatalln("No entries found, no error")
	}

	// VERIFY PAYLOAD
	// --------------
	// We only expect ONE entry in this situation,
	// so we'll target the index directly, which
	// contains an *ldap.Entry instance.
	entry := sr.Entries[0]
	if entry == nil {
		log.Fatalf("%T is nil", entry)
	}

	// PREPARE ITERABLE
	// ----------------
	// Prepare values for iteration in the desired
	// order specified earlier.
	var values [][]string = make([][]string, 8, 8)
	for idx, attr := range attrs {
		values[idx] = entry.GetEqualFoldAttributeValues(attr)
	}

	// PREPARE SUBSCHEMA
	// -----------------
	// Initialize our schemax.Subschema instance
	// for storage of marshaled definitions.
	subschema := schemax.NewSubschema()
	subschema.DN = schemaDN // optional

	// PREPARE SUPPLEMENTAL DEFINITIONS
	// --------------------------------
	// Sometimes, certain definition types are referenced
	// within a directory schema without actually being
	// present.
	//
	// One example of this is the subtreeSpecification
	// attributeType (RFC3672), which is referenced in
	// OpenLDAP by the subentry OC but doesn't exist :/.
	//
	// Users may choose to only import specific definitions
	// from specific RFCs. In this case, we'll import EVERY
	// definition we have, and then add whatever our DSA
	// provides on its own, assuming it is unique.
	subschema.LSC = schemax.PopulateDefaultLDAPSyntaxes()
	subschema.MRC = schemax.PopulateDefaultMatchingRules()
	subschema.ATC = schemax.PopulateDefaultAttributeTypes()

	// ITERATE and MARSHAL
	// -------------------
	// Loop through all values of each type of definition, and
	// parse each (string) value into proper type instances.
	for typ, defs := range values {
		if len(defs) == 0 {
			continue
		}
		for _, def := range defs {
			var err error
			switch attrs[typ] {
			case `ldapSyntaxes`:
				err = subschema.MarshalLDAPSyntax(def)
			case `matchingRules`:
				err = subschema.MarshalMatchingRule(def)
			case `attributeTypes`:
				err = subschema.MarshalAttributeType(def)
			case `matchingRuleUses`:
				//err = subschema.MarshalMatchingRuleUse(def)
			case `objectClasses`:
				err = subschema.MarshalObjectClass(def)
			case `dITContentRules`:
				err = subschema.MarshalDITContentRule(def)
			case `nameForms`:
				err = subschema.MarshalNameForm(def)
			case `dITStructureRules`:
				err = subschema.MarshalDITStructureRule(def)
			}
			chkerr(err)
		}
	}

	// Update our MatchingRuleUseCollection instance based
	// on all registered *AttributeType instances thus far.
	if subschema.ATC.Len() > 0 {
		chkerr(subschema.MRUC.Refresh(subschema.ATC))
	}

	fmt.Printf("############################################################\n")
	fmt.Printf("## BEGIN SCHEMA %s\n\n", subschema.DN)

	fmt.Printf("\n############################################################\n")
	fmt.Printf("## Parsed ldapSyntaxes: %d\n\n", subschema.LSC.Len())
	for i := 0; i < subschema.LSC.Len(); i++ {
		def := subschema.LSC.Index(i)
		def.SetSpecifier(`ldapsyntax`)
		def.SetUnmarshalFunc(def.LDAPSyntaxUnmarshalFunc)
		fmt.Printf("%s\n\n", def.String())
	}

	fmt.Printf("\n############################################################\n")
	fmt.Printf("## Parsed matchingRules: %d\n\n", subschema.MRC.Len())
	for i := 0; i < subschema.MRC.Len(); i++ {
		def := subschema.MRC.Index(i)
		def.SetSpecifier(`matchingrule`)
		def.SetUnmarshalFunc(def.MatchingRuleUnmarshalFunc)
		fmt.Printf("%s\n\n", def.String())
	}

	fmt.Printf("\n############################################################\n")
	fmt.Printf("## Parsed attributeTypes: %d\n\n", subschema.ATC.Len())
	for i := 0; i < subschema.ATC.Len(); i++ {
		def := subschema.ATC.Index(i)
		def.SetSpecifier(`attributetype`)
		def.SetUnmarshalFunc(def.AttributeTypeUnmarshalFunc)
		fmt.Printf("%s\n\n", def.String())
	}

	fmt.Printf("\n############################################################\n")
	fmt.Printf("## Parsed matchingRuleUses:%d\n\n", subschema.MRUC.Len())
	for i := 0; i < subschema.MRUC.Len(); i++ {
		def := subschema.MRUC.Index(i)
		def.SetSpecifier(`matchingruleuse`)
		def.SetUnmarshalFunc(def.MatchingRuleUseUnmarshalFunc)
		fmt.Printf("%s\n\n", def.String())
	}

	fmt.Printf("\n############################################################\n")
	fmt.Printf("## Parsed objectClasses: %d\n\n", subschema.OCC.Len())
	for i := 0; i < subschema.OCC.Len(); i++ {
		def := subschema.OCC.Index(i)
		def.SetSpecifier(`objectclass`)
		def.SetUnmarshalFunc(def.ObjectClassUnmarshalFunc)
		fmt.Printf("%s\n\n", def.String())
	}

	fmt.Printf("\n############################################################\n")
	fmt.Printf("## Parsed dITContentRules: %d\n\n", subschema.DCRC.Len())
	for i := 0; i < subschema.DCRC.Len(); i++ {
		def := subschema.DCRC.Index(i)
		def.SetSpecifier(`ditcontentrule`)
		def.SetUnmarshalFunc(def.DITContentRuleUnmarshalFunc)
		fmt.Printf("%s\n\n", def.String())
	}

	fmt.Printf("\n############################################################\n")
	fmt.Printf("## Parsed nameForms: %d\n\n", subschema.NFC.Len())
	for i := 0; i < subschema.NFC.Len(); i++ {
		def := subschema.NFC.Index(i)
		def.SetSpecifier(`nameform`)
		def.SetUnmarshalFunc(def.NameFormUnmarshalFunc)
		fmt.Printf("%s\n\n", def.String())
	}

	fmt.Printf("\n############################################################\n")
	fmt.Printf("## Parsed dITStructureRules: %d\n\n", subschema.DSRC.Len())
	for i := 0; i < subschema.DSRC.Len(); i++ {
		def := subschema.DSRC.Index(i)
		def.SetSpecifier(`ditstructurerule`)
		def.SetUnmarshalFunc(def.DITStructureRuleUnmarshalFunc)
		fmt.Printf("%s\n\n", def.String())
	}

	fmt.Printf("\n## END SCHEMA\n")
	fmt.Printf("##########################\n\n")
}

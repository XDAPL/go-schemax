/*
Author: Jesse Coretta - for github.com/JesseCoretta/go-schemax/_examples

This example describes a basic file parse of an LDAP schema in RFC4512 compliant format.

What happens:

 - Each definition is parsed as specific type instances (e.g.: *AttributeType, *DITContentRule, et al)
 - Each type is then assigned a desired unmarshaler function and then printed back to string
*/
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/JesseCoretta/go-schemax"
)

func main() {

	file := `/tmp/my.schema` // update as needed

	// quick error handler
	chkerr := func(err error) {
		if err != nil {
			log.Fatalln(err)
		}
	}
	hasPrefix := func(line, pfx string) bool {
		if strings.HasPrefix(strings.ToLower(line), pfx) {
			return true
		}
		return false
	}

	// read file into bytes
	data, err := ioutil.ReadFile(file)
	chkerr(err)

	proc := strings.ReplaceAll(string(data), "\n\t", " ")
	data = []byte(proc)

	// Make a schema object
	sch := schemax.NewSubschema()
	sch.PopulateDefaultLDAPSyntaxes()
	sch.PopulateDefaultMatchingRules()

	// Split "data" ([]byte) into string slices, delimited on the
	// newline character. Each slice represents a single line and
	// (possibly) a schema definition ...
	lines := strings.Split(string(data), "\n")

	// Iterate over each perceived line, and evaluate the raw text
	// to ascertain if it is a known definition type. If not recognized,
	// then take no action. If recognized, then unmarshal.
	for idx, line := range lines {

		// Ignore any purely empty lines
		if len(line) == 0 {
			continue
		}

		// Ignore any lines we believe are comments
		if line[0] == '#' {
			continue
		}

		// Look for a definition specifier (e.g.: "attributetype")
		// as the first component of each line. Please keep in mind
		// that these specifiers are known to vary across schema
		// standards adopted in different LDAP products, such as
		// Netscape vs. OpenLDAP ...
		switch {
		case hasPrefix(line, `ldapsyntax`):
			err = sch.MarshalLDAPSyntax(line)
		case hasPrefix(line, `matchingrule`):
			err = sch.MarshalMatchingRule(line)
		case hasPrefix(line, `attributetype`):
			err = sch.MarshalAttributeType(line)
		case hasPrefix(line, `objectclass`):
			err = sch.MarshalObjectClass(line)
		case hasPrefix(line, `ditcontentrule`):
			err = sch.MarshalDITContentRule(line)
		case hasPrefix(line, `nameform`):
			err = sch.MarshalNameForm(line)
		case hasPrefix(line, `ditstructurerule`):
			err = sch.MarshalDITStructureRule(line)
		default:
			err = fmt.Errorf("Unrecognized definition on, or around, line #%d: %s", idx+1, line)
		}

		chkerr(err)
	}

	// Now that all ATs are loaded, refresh the
	// manifest of applied MatchingRuleUses ...
	sch.MRUC.Refresh(sch.ATC)

	// Use the pretty package-provided unmarshaler
	// funcs to print with nice indenting and linebreaks,
	// and set our desired specifier for each collection.
	var raw string

	sch.LSC.SetSpecifier(`ldapsyntax`)
	sch.LSC.SetUnmarshaler(schemax.LDAPSyntaxUnmarshaler)
	raw, err = schemax.Unmarshal(sch.LSC)
	chkerr(err)
	fmt.Printf("%s\n", raw)

	sch.MRC.SetSpecifier(`matchingrule`)
	sch.MRC.SetUnmarshaler(schemax.MatchingRuleUnmarshaler)
	raw, err = schemax.Unmarshal(sch.MRC)
	chkerr(err)
	fmt.Printf("%s\n", raw)

	sch.ATC.SetSpecifier(`attributetype`)
	sch.ATC.SetUnmarshaler(schemax.AttributeTypeUnmarshaler)
	raw, err = schemax.Unmarshal(sch.ATC)
	chkerr(err)
	fmt.Printf("%s\n", raw)

	sch.MRUC.SetSpecifier(`matchingruleuse`)
	sch.MRUC.SetUnmarshaler(schemax.MatchingRuleUseUnmarshaler)
	raw, err = schemax.Unmarshal(sch.MRUC)
	chkerr(err)
	fmt.Printf("%s\n", raw)

	sch.OCC.SetSpecifier(`objectclass`)
	sch.OCC.SetUnmarshaler(schemax.ObjectClassUnmarshaler)
	raw, err = schemax.Unmarshal(sch.OCC)
	chkerr(err)
	fmt.Printf("%s\n", raw)

	sch.DCRC.SetSpecifier(`ditcontentrule`)
	sch.DCRC.SetUnmarshaler(schemax.DITContentRuleUnmarshaler)
	raw, err = schemax.Unmarshal(sch.DCRC)
	chkerr(err)
	fmt.Printf("%s\n", raw)

	sch.NFC.SetSpecifier(`nameform`)
	sch.NFC.SetUnmarshaler(schemax.NameFormUnmarshaler)
	raw, err = schemax.Unmarshal(sch.NFC)
	chkerr(err)
	fmt.Printf("%s\n", raw)

	sch.DSRC.SetSpecifier(`ditstructurerule`)
	sch.DSRC.SetUnmarshaler(schemax.DITStructureRuleUnmarshaler)
	raw, err = schemax.Unmarshal(sch.DSRC)
	chkerr(err)
	fmt.Printf("%s\n", raw)

	return
}

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

/*
flatten will (crudely) do all of the following in the order specified:

 - Take the provided byte slice and remove all linebreaks EXCEPT those that act as boundaries between definitions
 - Collapse all WHSP (\s+ and \t+) to a singular WHSP instances
 - Remove comments
 - Return the resultant set of lines as an instance of []string
 - ... all without regex :)

This function has been tested with all of the following file "conditions" in varying permutations:

 - Definitions contain "\n\t" or "\n\s" hanging indents for each field EXCEPT the first
 - Definitions are single-line definitions WITH OR WITHOUT an empty newline in between
 - Definitions are accompanied by comments of varying styles that MAY OR MAY NOT span multiple lines
*/
func flatten(data []byte) []string {

	// I don't want regex, so let's do this
	// a different way ...
	var clean string
	var next, last rune
	var activeComment bool
	for idx, char := range data {
		ch := rune(char)

		if idx+1 < len(data) {
			next = rune(data[idx+1])
		} else {
			next = rune(0)
		}

		switch ch {
		case '#':
			if last == '\n' || last != '#' {
				clean += string(ch)
				activeComment = true
			}
		case ' ':
			if ch != last && last != '\n' {
				clean += string(ch)
			}
		case '\n':
			if !activeComment {
				if last == ')' && ( next != ' ' && next != '\t' ) {
					clean += "<SPLITHERE>"
				}
			} else {
				activeComment = false
				if last != '\'' {
					clean += "<SPLITHERE>"
				}
			}
		case '\t':
			clean += " "
		default:
			clean += string(ch)
		}

		last = ch
	}

	// Split at the desired split points
	tlines := strings.Split(clean, "<SPLITHERE>")

	// Begin final processing, cleaning up each
	// delimited line to remove characters that
	// are no longer needed.
	var lines []string = make([]string, 0)
	for _, line := range tlines {
		if len(line) == 0 {
			continue
		}
		if line[0] == '#' {
			continue
		}

		var newline string = line
		if newline[0] == ' ' {
			newline = newline[1:]
		}

		if strings.HasSuffix(line, `\n`) {
			newline = newline[:len(newline)-2]
		}

		if strings.HasSuffix(line, " ") {
			newline = newline[:len(newline)-1]
		}

		lines = append(lines, newline) // seems ok
	}

	return lines
}

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

	// Make a schema object
	sch := schemax.NewSubschema()
	sch.PopulateDefaultLDAPSyntaxes()
	sch.PopulateDefaultMatchingRules()

	// Create a set of lines, one definition each
	lines := flatten(data)
	//fmt.Printf("%#v\n", lines)
	//panic("yo")

	// Iterate over each perceived line, and evaluate the raw text
	// to ascertain if it is a known definition type. If recognized,
	// then unmarshal.
	for idx, line := range lines {

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

	// OPTIONAL: Now that all ATs are loaded, refresh the
	// manifest of applied MatchingRuleUses ...
	sch.MRUC.Refresh(sch.ATC)
	if sch.MRUC.Len() > 0 {
		sch.MRUC.SetSpecifier(`matchingruleuse`)
		sch.MRUC.SetUnmarshaler(schemax.MatchingRuleUseUnmarshaler)
		raw, err = schemax.Unmarshal(sch.MRUC)
		chkerr(err)
		fmt.Printf("%s\n", raw)
	}

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

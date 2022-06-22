package schemax

/*
stripTags simply converts (for example) "userCertificate;binary" to "userCertificate" so that lookups and comparisons are performed properly. This package does not really care about tags, but the presence of such values should not have a negative impact.
*/
func stripTags(x string) (name string) {
	idx := indexRune(x, ';')
	if idx == -1 {
		return x
	}

	return x[:idx]
}

var parsers map[string]parseMeth

/*
lfindex (label/field) returns an index number indicating the correct field number associated with the search term.
*/
func lfindex(term string, def definition) (idx int) {
	if len(term) < 2 {
		return -1
	}

	// Get extensions out of the way ..
	if term[:2] == `X-` {
		for idx = range def.fields {
			if def.labels[idx] == `EXT` {
				return idx
			}
		}

		return -1 // weird
	}

	for idx = range def.fields {
		fname := def.fields[idx].Name
		if fname == term {
			return idx
		} else if def.labels[idx] == term {
			return idx
		} else if def.labels[idx] == `KIND` {
			switch term {
			case Structural.String(), Abstract.String(),
				Auxiliary.String():
				return idx
			}
		} else if def.labels[idx] == `FLAGS` {
			switch term {
			case SingleValue.String(), NoUserModification.String(), Collective.String():
				return idx
			}
		} else {
			// Some schemas use the full "SUBSTRINGS" term as opposed to
			// the far-more-popular abbreviation ...
			if term == `SUBSTRINGS` && def.labels[idx] == `SUBSTR` {
				return idx
			}
		}
	}

	return -1
}

func parse_qdstring(def string) (name []string, rest string, ok bool) {
	idx := indexRune(def, ' ')
	if idx == -1 {
		return
	}

	closer := `' `
	idx = index(def[1:], closer)
	if idx == -1 {
		return
	}

	name = []string{def[1 : idx+1]}
	rest = def[idx+3:]
	ok = len(name) > 0

	return
}

func parse_extensions(def string) (name []string, rest string, ok bool) {
	return parse_qdescrs(def) // extensions are just like NAMEs
}

/*
parse_oids_ids returns one or more object identifier values or, in the case of DITStructureRule instances, an integer value.  The untouched remainder of the definition is returned, along with a success-indicative boolean value.
*/
func parse_oids_ids(def string) (name []string, rest string, ok bool) {
	idx := indexRune(def, ' ')
	if idx == -1 {
		return
	}
	ndef := def[idx:]

	switch {
	case def[0] == '(':
		// *POSSIBLY* multi-valued
		closer := ')'
		idx = indexRune(ndef[2:], closer)
		if idx == -1 {
			return
		}
		rawname := def[2 : idx+2]
		var vals string
		for c := range rawname {
			ch := rune(rawname[c])
			if ch == '\'' {
				continue
			}
			vals += string(ch)
		}
		name = split(vals, ` $ `)
		rest = def[idx+5:]
	default:
		// *DEFINITELY* single-valued
		closer := ' '
		idx = indexRune(def[1:], closer)
		if idx == -1 {
			return
		}

		name = []string{def[:idx+1]}
		rest = def[idx+2:]
	}
	ok = len(name) > 0

	return
}

func parse_qdescrs(def string) (name []string, rest string, ok bool) {
	idx := indexRune(def, ' ')
	if idx == -1 {
		return
	}
	ndef := def[idx:]

	switch def[0] {
	case '(':
		// *POSSIBLY* multi-valued
		closer := ')'
		idx = indexRune(ndef[2:], closer)
		if idx == -1 {
			return
		}
		rawname := def[2 : idx+1]
		var vals string
		for c := range rawname {
			ch := rune(rawname[c])
			if ch == '\'' {
				continue
			}
			vals += string(ch)
		}
		name = split(vals, ` `)
		rest = def[idx+5:]
	case '\'':
		// *DEFINITELY* single-valued
		closer := '\''
		idx = indexRune(def[1:], closer)
		if idx == -1 {
			return
		}

		name = []string{def[1 : idx+1]}
		rest = def[idx+3:]
	}
	ok = len(name) > 0

	return
}

func parse_atflags(def string) ([]string, string, bool) {
	return []string{`true`}, def, true
}

func parse_obsolete(def string) ([]string, string, bool) {
	return []string{`true`}, def, true
}

/*
func parse_boolean(def string) ([]string, string, bool) {
	return []string{`true`}, def, true
}
*/

func parse_kind(def string) (name []string, rest string, ok bool) {
	if len(def) == 0 {
		return
	}

	idx := indexRune(def, ' ')
	if idx == -1 {
		return
	}
	n := def[:idx]
	name = []string{n}
	rest = def[idx+1:]
	ok = len(n) > 0

	return
}

func parse_usage(def string) (name []string, rest string, ok bool) {
	if len(def) == 0 {
		return
	}

	idx := indexRune(def, ' ')
	if idx == -1 {
		return
	}
	n := def[:idx]
	name = []string{n}
	rest = def[idx+1:]
	ok = len(n) > 0

	return
}

func parse_definition_label(def string) (label []string, rest string, ok bool) {

	if len(def) == 0 {
		rest = def
		return
	}

	idx := indexRune(def, ' ')
	if idx == -1 {
		rest = def
		return
	}

	testlabel := def[:idx]
	for c := range testlabel {
		ch := rune(testlabel[c])
		if !runeIsUpper(ch) && ch != '-' {
			rest = def
			return
		}
	}

	switch testlabel {
	case Auxiliary.String(), Structural.String(), Abstract.String():
		label = []string{`KIND`}
		rest = def
	case NoUserModification.String(), Collective.String(), SingleValue.String():
		label = []string{testlabel}
		rest = def[idx+1:]
	default:
		label = []string{testlabel}
		rest = def[len(testlabel)+1:]
	}
	ok = len(testlabel) > 0

	return
}

/*
parse is the first field-parsing routine and returns three (3) values:

  - Two (2) strings
  - One (1) boolean

The first string is the parsed OID or RuleID, while the second string is the remainding (unparsed) text.

The boolean value indicates whether parsing was successful.
*/
func parse(def string) (id []string, rest string, ok bool) {
	// Remove any leading or trailing WHSP
	def = trimSpace(def)

	// Remove any newline chars with a
	// zero string ...
	def = replaceAll(def, `\n`, ``)

	// start by finding the first opening parenthesis ('('). This
	// is considered the absolute start of the definition. Anything
	// before this character, e.g.: "attributetype", is discarded.
	idx := indexRune(def, '(')
	if idx == -1 {
		return
	}
	def = def[idx:]

	// Ensure the first and last runes are the opening
	// and closing parenthesis characters respectively
	if def[0] != '(' || def[len(def)-1] != ')' {
		return
	}

	// Ensure the second and second-to-last runes
	// are WHSP characters only.
	if def[1] != ' ' || def[len(def)-2] != ' ' {
		return
	}

	// Parse (and possibly resolve an alias to) an OID
	if id, rest, ok = parse_numericoid(def[2 : len(def)-1]); !ok {
		if id, rest, ok = parse_ruleid(def[2 : len(def)-1]); !ok {
			return
		}
	}

	ok = (len(id) > 0 && len(rest) > 0)

	return
}

/*
parse_ruleid returns a DIT Structure Rule number, the untouched remainder of the definition and a success-indicative boolean value.  This function is intended only for definitions destined to marshal into a DITStructureRule instance.
*/
func parse_ruleid(def string) (id []string, rest string, ok bool) {
	idx := indexRune(def, ' ')
	if idx == -1 {
		return
	}
	id = []string{def[:idx]}
	rest = def[idx+1:]
	ok = len(id) > 0

	return
}

/*
parse_mub takes a syntax OID with a minimum upper-bound value, such as `1.3.6.1.4.1.1466.115.121.1.15{128}`, and discards the oid and curly braces. A single string number is returned in a single-member slice of strings.
*/
func parse_mub(def string) (mub []string, rest string, ok bool) {
	idx := indexRune(def, '{')
	if idx == -1 {
		return
	}

	oid := def[:idx]
	part := def[idx+1:]
	idx = indexRune(part, '}')
	if idx == -1 {
		return
	}

	x := part[:idx]
	rest = ``
	mub = []string{oid, x}
	ok = len(x) > 0

	return
}

/*
parse_numericoid returns a dotted ASN.1 object identifier, the untouched remainder of the definition and a success-indicative boolean value.  This function is intended for use on any definition *except* for those destined to marshal intoa DITStructureRule instance.
*/
func parse_numericoid(def string) (oid []string, rest string, ok bool) {
	idx := indexRune(def, ' ')
	if idx == -1 {
		return
	}

	var part string
	if def[0] == '\'' {
		if def[idx-1] != '\'' {
			return
		}
		part = def[1 : idx-1]
	} else {
		part = def[:idx]
	}

	var lastSpec bool
	var o string

	for c := range part {
		ch := rune(part[c])
		switch {
		case runeIsLetter(ch) || runeIsDigit(ch) ||
			ch == ':' || ch == '.' || ch == '-' || ch == '{' || ch == '}':
			if c == 0 && (!runeIsLetter(ch) && !runeIsDigit(ch)) {
				return // first char must be letter or number
			}
			if c == len(oid)-1 && (!runeIsLetter(ch) &&
				!runeIsDigit(ch)) {
				return // last char must be letter or number
			}
			if !runeIsLetter(ch) && !runeIsDigit(ch) {
				if lastSpec {
					return // contiguous delimiters
				}
				lastSpec = true
			} else {
				lastSpec = false
			}
		default:
			return
		}

		// Finally add the rune to the
		// incremental value
		o += string(ch)
	}

	oid = []string{o}
	rest = def[idx+1:]

	ok = len(o) > 0
	return
}

/*
sanitize takes the user-provided string-based schema definition and will remove extraneous content not needed for parsing.

Such content includes repeated WHSP chars (e.g.: tabs/spaces), which are collapsed down to single spaces as needed for delimitation, and the outright removal of any newlines.
*/
func sanitize(def string) (processed string) {
	def = replaceAll(def, `\n`, ``)
	var lastWasSpace bool
	for c := range def {
		ch := rune(def[c])
		switch ch {
		case '\n':
			continue
		case ' ', '\t':
			if !lastWasSpace {
				lastWasSpace = true
				processed += string(' ')
			}
			continue
		default:
			processed += string(ch)
			lastWasSpace = false
		}
	}
	processed = trimSpace(processed)

	return
}

func isNumericalOID(x interface{}) (is bool) {
	var val string
	switch tv := x.(type) {
	case string:
		val = stripTags(tv)
	case OID:
		val = tv.String()
	default:
		return
	}

	oid := split(val, `.`)
	if len(oid) < 2 {
		return false
	}

	switch oid[0] {
	case `0`, `1`, `2`:
		// OK!
	default:
		return
	}

	var dot bool
	for arc := 1; arc < len(oid); arc++ {
		for c := range oid[arc] {
			ch := rune(oid[arc][c])
			switch {
			case isDigit(string(ch)):
				dot = false
				continue
			case ch == '.' && !dot &&
				arc != 0 && arc != len(oid)-1:
				dot = true
			default:
				return
			}
		}
	}
	is = true

	return
}

func strInSlice(str string, slice []string) bool {
	if len(str) == 0 || len(slice) == 0 {
		return false
	}

	for el := range slice {
		if slice[el] == str {
			return true
		}
	}

	return false
}

func isDigit(val string) (is bool) {
	if len(val) == 0 {
		return
	}
	for c := range val {
		ch := rune(val[c])
		switch {
		case '0' <= ch && ch <= '9':
			continue
		default:
			return
		}
	}
	is = true

	return
}

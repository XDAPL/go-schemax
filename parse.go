package schemax

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
		} else if def.labels[idx] == `BOOLS` {
			switch term {
			case Obsolete.string(), SingleValue.string(),
				NoUserModification.string(), Collective.string():
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

func definitionType(def definition) (n string) {
	n = `unknown`
	t := def.typ.Name()
	switch t {
	case `ObjectClass`:
		n = `oc`
	case `StructuralObjectClass`:
		n = `soc`
	case `AttributeType`:
		n = `at`
	case `SuperiorAttributeType`:
		n = `sat`
	case `LDAPSyntax`:
		n = `ls`
	case `Equality`:
		n = `eq`
	case `Substring`:
		n = `ss`
	case `Ordering`:
		n = `ord`
	case `MatchingRule`:
		n = `mr`
	case `MatchingRuleUse`:
		n = `mru`
	case `DITContentRule`:
		n = `dcr`
	case `DITStructureRule`:
		n = `dsr`
	case `NameForm`:
		n = `nf`
	}

	return
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

	switch {
	case def[0] == '(':
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
	case def[0] == '\'':
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

func parse_boolean(def string) ([]string, string, bool) {
	return []string{`true`}, def, true
}

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
	case Obsolete.string(), NoUserModification.string(),
		Collective.string(), SingleValue.string():
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

	//if alias {
	// TODO
	//oid = at.resolveOID(oid)
	//}
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
Unmarshal takes an instance of one (1) of the following types and (if valid) and returns the textual form of the definition:

 - *ObjectClass
 - *AttributeType
 - *LDAPSyntax
 - *MatchingRule
 - *MatchingRuleUse
 - *DITContentRule
 - *DITStructureRule
 - *NameForm

Should any validation errors occur, a non-nil instance of error is returned.
*/
func Unmarshal(x interface{}) (def string, err error) {
	switch tv := x.(type) {
	case *ObjectClass:
		def, err = tv.unmarshal(true)
	case *AttributeType:
		def, err = tv.unmarshal(true)
	case *LDAPSyntax:
		def, err = tv.unmarshal(true)
	case *MatchingRule:
		def, err = tv.unmarshal(true)
	case *MatchingRuleUse:
		def, err = tv.unmarshal(true)
	case *DITContentRule:
		def, err = tv.unmarshal(true)
	case *DITStructureRule:
		def, err = tv.unmarshal(true)
	case *NameForm:
		def, err = tv.unmarshal(true)
	default:
		err = raise(invalidUnmarshal,
			"unknown or unsupported type %T", tv)
	}

	if err != nil {
		err = raise(invalidUnmarshal, err.Error())
	} else if len(def) == 0 {
		err = raise(invalidUnmarshal,
			"zero-length definition returned from Unmarshal (of %T)", x)
	}

	return
}

/*
Marshal takes the provided schema definition (def) and attempts to marshal it into x.  x MUST be one of the following types:

 - *AttributeType
 - *ObjectClass
 - *LDAPSyntax
 - *MatchingRule
 - *MatchingRuleUse
 - *DITContentRule
 - *DITStructureRule
 - *NameForm

Should any validation errors occur, a non-nil instance of error is returned.

Note that it is far more convenient to use the Subschema.Marshal wrapper, as it only requires a single argument (the raw definition).
*/
func Marshal(raw string, x interface{},
	alm AliasesManifest,
	atm AttributeTypesManifest,
	ocm ObjectClassesManifest,
	lsm LDAPSyntaxesManifest,
	mrm MatchingRulesManifest,
	mrum MatchingRuleUsesManifest,
	dcrm DITContentRulesManifest,
	dsrm DITStructureRulesManifest,
	nfm NameFormsManifest) (err error) {
	// I am so sorry.

	if len(raw) == 0 {
		return raise(emptyDefinition, "no length")
	}

	// Remove all outer WHSP, collapse all successive inner
	// WHSP to single space, and purge all linebreaks.
	raw = sanitize(raw)

	def, ok := newDefinition(x, alm)
	if !ok {
		return raise(invalidMarshal, "newDefinition: assembly failure")
	}

	// here we parse the OID, which is a constant
	// for all schema types, as our first order ...

	id, rest, ok := parse(raw)
	if !ok {
		return raise(invalidMarshal, invalidOID.Error())
	}

	if _, asserted := x.(*DITStructureRule); asserted {
		def.values[0].Set(valueOf(NewRuleID(id[0])))
	} else {
		isnumoid := isNumericalOID(id[0])
		if alm.IsZero() {
			if !isnumoid {
				return raise(invalidMarshal, "unresolvable alias '%s' (nil manifest)", id[0])
			}
			def.values[0].Set(valueOf(OID(id[0])))
		} else {
			oid, ok := alm.Resolve(id[0])
			if !ok {
				return raise(invalidMarshal, "unresolvable alias '%s'", id[0])
			}
			def.values[0].Set(valueOf(oid))
		}
	}

	// Now we'll parse all KEY WHSP VALUE [VALUE...] instances
	for {
		if len(rest) <= 1 || err != nil {
			break
		}

		// parseDefLabel receives one chunk of information
		// (which should be a single, raw schema definition).
		// We then attempt to extract a "label" from this,
		// and then parse the remainder ("rest") based on
		// the appropriate known actions for said type of
		// value (e.g.: `NAME` vs. `SYNTAX`).
		var label []string
		if label, rest, ok = parse_definition_label(rest); !ok {
			return raise(invalidLabel,
				"failed parse for label was: '%s', raw def: '%s'",
				label, rest)
		} else {
			idx := def.lfindex(label[0])
			if idx == -1 {
				return raise(invalidLabel,
					"failed index localization (lfindex) for label: '%s', raw def: '%s'",
					label[0], rest)
			}

			var value []string
			if value, rest, ok = def.meths[idx](rest); !ok {
				return raise(invalidValue,
					"failed value localization for label: '%s' (deflabel:%s), raw def: '%s'",
					label[0], def.labels[idx], rest)
			}

			switch def.labels[idx] {
			case `KIND`:
				err = def.setKind(value[0], idx)
			case `EXT`:
				err = def.setExtensions(label[0], value, idx)
			case `NAME`:
				err = def.setName(idx, value...)
			case `DESC`:
				err = def.setDesc(idx, value[0])
			case `BOOLS`:
				err = def.setBoolean(label[0], x)
			case `USAGE`:
				err = def.setUsage(value[0], idx)
			case `FORM`:
				err = def.setNameForm(nfm, x, value[0], idx)
			case `MAY`:
				err = def.setPermittedAttributeTypes(atm, x, value, idx)
			case `NOT`:
				err = def.setProhibitedAttributeTypes(atm, x, value, idx)
			case `MUST`:
				err = def.setRequiredAttributeTypes(atm, x, value, idx)
			case `APPLIES`:
				err = def.setApplies(atm, x, value, idx)
			case `AUX`:
				err = def.setAuxiliaryObjectClasses(ocm, x, value, idx)
			case `OC`:
				err = def.setStructuralObjectClass(ocm, x, value[0], idx)
			case `SUP`:
				switch def.definitionType() {
				case `oc`:
					err = def.setSuperiorObjectClasses(ocm, x, value, idx)
				case `sat`, `at`:
					err = def.setSuperiorAttributeType(atm, x, value[0], idx)
				case `dsr`:
					err = def.setSuperiorDITStructureRules(dsrm, x, value, idx)
				}
			case `SYNTAX`:
				err = def.setSyntax(lsm, x, value[0], idx)
			case `EQUALITY`, `SUBSTR`, `SUBSTRINGS`, `ORDERING`:
				err = def.setEqSubOrd(mrm, x, value[0], idx)
			default:
				return raise(invalidLabel,
					"Field '%s'(def.label:'%s') unhandled; would have set '%s', raw def: '%s')",
					label[0], def.labels[idx], value, rest)
			}
		}
		label = []string{} // reset our label
	}

	if err != nil {
		return
	}

	// Our instance has been populated with the
	// marshaled bytes. Now we conduct validation
	// checks to ensure said bytes were sane.
	switch tv := x.(type) {
	case *LDAPSyntax:
		// we'll take an extra step to identify
		// any syntax that is considered to be
		// human readable either through a value
		// of 'FALSE' for the X-NOT-HUMAN-READABLE
		// well-known extension, or absence of said
		// extension altogether.
		if tv.Extensions.Exists(`X-NOT-HUMAN-READABLE`) {
			if strInSlice(`FALSE`, tv.Extensions[`X-NOT-HUMAN-READABLE`]) {
				tv.bools.set(HumanReadable)
			}
		} else {
			tv.bools.set(HumanReadable)
		}
		err = tv.validate()
	case *AttributeType:
		err = tv.validate()
	case *ObjectClass:
		err = tv.validate()
	case *NameForm:
		err = tv.validate()
	case *MatchingRule:
		err = tv.validate()
	case *MatchingRuleUse:
		err = tv.validate()
	case *DITContentRule:
		err = tv.validate()
	case *DITStructureRule:
		err = tv.validate()
	default:
		err = raise(unexpectedType,
			"No validator for %T", tv)
	}

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

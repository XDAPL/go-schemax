package schemax

/*
Macros is a map structure associating a string value known as "macro" name with a dot-delimited ASN.1 object identifier.  Use of this is limited to scenarios involving LDAP implementations that support OID macros for certain schema definition elements.
*/
type Macros map[string]OID

func (r *Macros) IsZero() bool {
	return r == nil
}

func isAnAlias(key string) (alias, rest string, is, ok bool) {
	if 0 == len(key) || len(key) > 512 {
		return
	}

	if isNumericalOID(key) {
		ok = true
		rest = key
		return
	}

	var done bool
	var idx int = -1
	for _, c := range key {
		if done {
			break
		}

		ch := rune(c)
		switch {
		case ch == ':' || ch == '.':
			idx = indexRune(key, ch)
			done = true
		case runeIsLetter(ch) || runeIsDigit(ch) || ch == '-':
			alias += string(ch)
		default:
			// unsupported char
			alias = ``
			return
		}
	}
	is = len(alias) > 0
	ok = (len(rest) > 0 || is)

	// no delim detected but there was
	// an alias (could be literal alias
	// to oid w/o leaf node) ...
	if idx >= 0 {
		rest = key[idx+1:]
	}

	return
}

/*
Set assigns the provided key and oid to the receiver instance.  The arguments must both be strings.  The first argument must be the alias name, while the second must be its ASN.1 dotted object identifier.  Subsequent values are ignored.
*/
func (r *Macros) Set(als ...interface{}) {
	if r.IsZero() {
		*r = NewMacros()
	}

	if len(als) < 2 {
		return
	}

	var alias string
	var oid string
	var ok bool

	if alias, ok = als[0].(string); !ok {
		return
	}

	if oid, ok = als[1].(string); !ok {
		return
	}

	if !isNumericalOID(oid) {
		return
	}

	// avoid duplicate alias or oid entries ...
	for i, v := range *r {
		if v.String() == oid || equalFold(alias, i) {
			continue
		}
	}

	R := *r
	R[alias] = NewOID(oid)
	*r = R

	return
}

/*
Resolve returns a registered OID and a boolean value indicative of a successful lookup. A search is conducted using the provided alias key name, with or without the colon:number suffix included.
*/
func (r Macros) Resolve(x interface{}) (oid OID, ok bool) {
	// If its already an OID, don't bother
	// doing a lookup, just return it as-is.
	if oid, ok = x.(OID); ok {
		return
	}

	// If it isn't a string-based name,
	// then we can't go any further.
	key, assert := x.(string)
	if !assert {
		return
	}

	var alias string
	var rest string
	var is bool

	alias, rest, is, ok = isAnAlias(key)
	if !is {
		return
	}

	switch {
	case is:
		oid, ok = r[alias]
		if !ok {
			return
		}
		if len(rest) > 0 {
			oid = NewOID(oid.String() + `.` + rest)
		}
	default:
		if len(rest) == 0 {
			return
		}
		oid = NewOID(rest)
	}
	ok = isNumericalOID(oid.String())

	return
}

/*
NewMacros returns a new instance of Macros, intended for use in resolving OID aliases.
*/
func NewMacros() Macros {
	return make(Macros, 0)
}

/*
Extensions is a map structure associating a string extension name (e.g.: X-ORIGIN) with a non-zero string value.
*/
type Extensions map[string][]string

/*
Exists returns a boolean value indicative of whether the named label exists within the receiver instance.
*/
func (r Extensions) Exists(label string) (exists bool) {
	for k, _ := range r {
		if k == label {
			exists = true
			return
		}
	}

	return
}

func (r Extensions) IsZero() bool {
	if r == nil {
		return true
	}
	return len(r) == 0
}

/*
Set assigns the provided label and value(s) to the receiver instance.  The first argument must be a string and is interpreted as a label (e.g.: X-ORIGIN).

All subsequent values (strings or slices of strings) are interpreted as values to be assigned to said label.
*/
func (r Extensions) Set(x ...interface{}) {
	if len(x) < 2 {
		return
	}

	label, ok := x[0].(string)
	if !ok {
		return
	}

	if r.Exists(label) || !r.labelIsValid(label) {
		return
	}

	var values []string
	for i := 1; i < len(x); i++ {
		switch tv := x[i].(type) {
		case string:
			if len(tv) == 0 {
				return
			}
			values = append(values, tv)
		case []string:
			for _, z := range tv {
				if len(z) == 0 {
					return
				}
				values = append(values, z)
			}
		default:
			return
		}
	}

	if ok = len(values) > 0; ok {
		r[label] = values
	}

	return
}

/*
Equals compares the receiver to the provided interface value (which must be of the same effective type).  A comparison is performed, and an equality-indicative boolean value is returned.
*/
func (r Extensions) Equal(x interface{}) (equals bool) {
	if assert, ok := x.(Extensions); ok {
		if len(assert) != len(r) {
			return
		}

		for k, v := range r {
			v2, exists := assert[k]
			if !exists {
				return
			}
			if len(v2) != len(v) {
				return
			}
			for i, z := range v2 {
				if equals = equalFold(z, v[i]); !equals {
					return
				}
			}
		}
		equals = true
	}

	return
}

/*
String is a stringer method that returns the receiver data as a compliant schema definition component.
*/
func (r Extensions) String() (exts string) {
	for _, v := range r.strings() {
		exts += v + ` `
	}
	if len(exts) == 0 {
		return
	}
	if exts[len(exts)-1] == ' ' {
		exts = exts[:len(exts)-1]
	}

	return
}

func (r Extensions) strings() (exts []string) {
	exts = make([]string, len(r))
	ct := 0
	for k, v := range r {
		if len(v) == 1 {
			exts[ct] = k + ` '` + v[0] + `'`
		} else if len(v) > 1 {
			vals := make([]string, len(v))
			for mi, mv := range v {
				val := `'` + mv + `'`
				vals[mi] = val
			}
			exts[ct] = k + ` ( ` + join(vals, ` `) + ` )`
		}
		ct++
	}

	return
}

/*
NewExtensions returns a new instance of Extensions, intended for assignment to any definition type.
*/
func NewExtensions() Extensions {
	return make(Extensions, 0)
}

package schemax

/*
macros.go handles OID->Name macros.
*/

/*
Macros is a map structure associating a string value known as a "macro" name with a dot-delimited ASN.1 object identifier.

Use of instances of this type is limited to scenarios involving LDAP implementations that support OID macros for certain schema definition elements. Generally speaking, it is best to avoid macros if at all possible.
*/
type Macros map[string]OID

// IsZero returns a boolean value indicative of whether the receiver is considered zero, or undefined.
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
NewMacros returns a new instance of Macros.
*/
func NewMacros() Macros {
	return make(Macros, 0)
}

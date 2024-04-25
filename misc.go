package schemax

func isNumericOID(id string) bool {
	if len(id) == 0 {
		return false
	}

	if !('0' <= rune(id[0]) && rune(id[0]) <= '2') || id[len(id)-1] == '.' {
		return false
	}

	var last rune
	for _, c := range id {
		switch {
		case c == '.':
			if last == c {
				return false
			}
			last = '.'
		case isDigit(c):
			last = c
			continue
		}
	}

	return true
}

func isAlnum(x rune) bool {
	return isDigit(x) || isAlpha(x)
}

func bool2str(x bool) string {
	if x {
		return `true`
	}

	return `false`
}

/*
IsDescriptor scans the input string val and judges whether
it qualifies a descriptor, in that all of the following
evaluate as true:

  - non-zero in length
  - begins with an alphabetical character
  - ends in an alphanumeric character
  - contains only alphanumeric characters or hyphens
  - no contiguous hyphens

This function is an alternative to engaging the antlr4512
parser.
*/
func IsDescriptor(val string) bool {
	return isDescriptor(val)
}

func isDescriptor(val string) bool {
	if len(val) == 0 {
		return false
	}

	// must begin with an alpha.
	if !isAlpha(rune(val[0])) {
		return false
	}

	// can only end in alnum.
	if !isAlnum(rune(val[len(val)-1])) {
		return false
	}

	// watch hyphens to avoid contiguous use
	var lastHyphen bool

	// iterate all characters in val, checking
	// each one for "descr" validity.
	for i := 0; i < len(val); i++ {
		ch := rune(val[i])
		switch {
		case isAlnum(ch):
			lastHyphen = false
		case ch == '-':
			if lastHyphen {
				// cannot use consecutive hyphens
				return false
			}
			lastHyphen = true
		default:
			// invalid character (none of [a-zA-Z0-9\-])
			return false
		}
	}

	return true
}

/*
mapTransferExtensions returns the provided dest instance of DefinitionMap,
following an attempt to copy all extensions found within src into dest.

This is mainly used to keep cyclomatics low during presentation and marshaling
procedures and may be used for any Definition qualifier.

The dest input value must be initialized else go will panic.
*/
func mapTransferExtensions(src Definition, dest DefinitionMap) DefinitionMap {
	exts := src.Extensions()
	for _, k := range exts.Keys() {
		if ext, found := exts.get(k); found {
			dest[k] = ext.List()
		}
	}

	return dest
}

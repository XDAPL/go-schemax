package schemax

func newMacros() Macros {
	return Macros{
		make(macros, 0),
	}
}

/*
Resolve returns string y (a name) following an attempt to forward-resolve
string x (numeric OID).  The Boolean value (found) is returned indicative
of a successful resolution attempt.

Case is not significant in the matching process.
*/
func (r Macros) Resolve(x string) (y string, found bool) {
	if r.macros == nil {
		return
	}

	for k, v := range r.macros {
		if eq(x, k) {
			y = v
			found = true
			break
		}
	}

	return
}

/*
ReverseResolve returns string x (a numeric OID) following an attempt to
reverse-resolve string y (name).  The Boolean value (found) is returned
indicative of a successful resolution attempt.

Case is not significant in the matching process.
*/
func (r Macros) ReverseResolve(y string) (x string, found bool) {
	if r.macros == nil {
		return
	}

	for k, v := range r.macros {
		if eq(y, v) {
			x = k
			found = true
			break
		}
	}

	return
}

/*
Set assigns value y (macro name) to key x (numeric OID).

This is a fluent method.
*/
func (r Macros) Set(x, y string) Macros {
	r.macros[x] = y
	return r
}

/*
Keys returns all numeric OID keys present within the receiver instance.
*/
func (r Macros) Keys() []string {
	var s []string
	for k := range r.macros {
		s = append(s, k)
	}

	return s
}

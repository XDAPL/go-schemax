package schemax

/*
MaxLength returns the integer value, if one was specified, that defines the maximum acceptable value size supported by this *AttributeType per its associated *LDAPSyntax.  If not applicable, a -1 is returned.
*/
func (r AttributeType) MaxLength() int {
	return int(r.mub)
}

/*
SetMaxLength sets the minimum upper bounds, or maximum length, of the receiver instance. The argument must be a positive, non-zero integer.

This will only apply to *AttributeTypes that use a human-readable syntax.
*/
func (r *AttributeType) SetMaxLength(max int) {
	r.setMUB(max)
}

/*
setBoolean is a private method used by reflect to set the minimum upper bounds.
*/
func (r *AttributeType) setMUB(mub interface{}) {
	if r.IsZero() {
		return
	}

	if r.Syntax.IsZero() {
		return
	}

	if !r.Syntax.IsHumanReadable() {
		return
	}

	switch tv := mub.(type) {
	case string:
		n, err := atoi(tv)
		if err != nil || n < 0 {
			return
		}
		r.mub = uint(n)
	case int:
		if tv > 0 {
			r.mub = uint(tv)
		}
	case uint:
		r.mub = tv
	}

}

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

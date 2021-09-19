package schemax

func (a *AttributeType) validate() (err error) {
	if a.IsZero() {
		return raise(isZero, "%T.validate", a)
	}

	if err = validateBool(a.bools); err != nil {
		return
	}

	var ls *LDAPSyntax
	if ls, err = a.validateSyntax(); err != nil {
		return
	}

	if err = a.validateMatchingRules(ls); err != nil {
		return
	}

	if err = validateNames(a.Name.strings()...); err != nil {
		return
	}

	if err = validateDesc(a.Description); err != nil {
		return
	}

	if !a.SuperType.IsZero() {
		if a.SuperType.Syntax.IsZero() {
			err = raise(invalidUnmarshal, "%T.unmarshal: %T.%T: %s (sub-typed)",
				a.SuperType, a.SuperType, a.SuperType.Syntax, isZero.Error())
		}
	} else {
		if a.Syntax.IsZero() {
			err = raise(invalidUnmarshal, "%T.unmarshal: %T.%T: %s (not sub-typed)",
				a, a, a.Syntax, isZero.Error())
		}
	}

	return
}

func (o *ObjectClass) validate() (err error) {
	if o.IsZero() {
		return raise(isZero, "%T.validate", o)
	}

	if err = validateBool(o.bools); err != nil {
		return
	}

	if err = o.validateKind(); err != nil {
		return
	}

	if err = validateNames(o.Name.strings()...); err != nil {
		return
	}

	if err = validateDesc(o.Description); err != nil {
		return
	}

	return
}

func (r *NameForm) validate() (err error) {
	if r.IsZero() {
		return raise(isZero, "%T.validate", r)
	}

	if err = validateBool(r.bools); err != nil {
		return
	}

	if err = validateNames(r.Name.strings()...); err != nil {
		return
	}

	if err = r.validateMustMay(); err != nil {
		return
	}

	if err = r.validateStructuralObjectClass(); err != nil {
		return
	}

	if err = validateDesc(r.Description); err != nil {
		return
	}

	return
}

func (r *MatchingRuleUse) validate() (err error) {
	if r.IsZero() {
		return raise(isZero, "%T.validate", r)
	}

	if r.OID.IsZero() {
		return raise(isZero, "%T.validate: no %T",
			r, r.OID)
	}

	if err = validateNames(r.Name.strings()...); err != nil {
		return
	}

	if err = validateDesc(r.Description); err != nil {
		return
	}

	if r.Applies.IsZero() {
		return raise(isZero, "%T.validate: no %T",
			r, r.Applies)
	}

	if err = validateBool(r.bools); err != nil {
		return
	}

	return
}

func (r *MatchingRule) validate() (err error) {
	if r.IsZero() {
		return raise(isZero, "%T.validate", r)
	}

	if err = r.validateSyntax(); err != nil {
		return
	}

	if err = r.validateSyntax(); err != nil {
		return
	}

	if err = validateBool(r.bools); err != nil {
		return
	}

	if err = validateNames(r.Name.strings()...); err != nil {
		return
	}

	if err = validateDesc(r.Description); err != nil {
		return
	}

	return
}

func (r *DITContentRule) validate() (err error) {
	if r.IsZero() {
		return raise(isZero, "%T.validate", r)
	}

	if err = validateDesc(r.Description); err != nil {
		return
	}

	return
}

func (r *LDAPSyntax) validate() (err error) {
	if r.IsZero() {
		return raise(isZero, "%T.validate", r)
	}

	if err = validateDesc(r.Description); err != nil {
		return
	}

	return
}

func (r *DITStructureRule) validate() (err error) {
	if r.IsZero() {
		err = raise(isZero, "%T.validate", r, r)
	}

	if r.Form.IsZero() {
		err = raise(invalidNameForm,
			"%T.validate: missing %T",
			r, r.Form)
	}

	if err = validateNames(r.Name.strings()...); err != nil {
		return
	}

	if err = validateDesc(r.Description); err != nil {
		return
	}

	if err = validateBool(r.bools); err != nil {
		return
	}

	if !r.SuperiorRules.IsZero() {
		for _, sup := range r.SuperiorRules {
			if err = sup.(*DITStructureRule).validate(); err != nil {
				return err
			}
		}
	}

	return
}

func (r *MatchingRule) validateSyntax() (err error) {
	if r.Syntax.IsZero() {
		err = raise(invalidSyntax,
			"%T.validateSyntax: zero syntax", r)
	}

	return
}

func (r *ObjectClass) getMay(m PermittedAttributeTypes) (ok PermittedAttributeTypes) {
	ok.Set(m)

	if !r.SuperClass.IsZero() {
		for _, sup := range r.SuperClass {
			ok.Set(sup.(*ObjectClass).getMay(ok))
		}
	}
	if !r.May.IsZero() {
		for _, may := range r.May {
			ok.Set(may)
		}
	}

	return
}

func (r *ObjectClass) getMust(m RequiredAttributeTypes) (req RequiredAttributeTypes) {
	req.Set(m)

	if !r.SuperClass.IsZero() {
		for _, sup := range r.SuperClass {
			req.Set(sup.(*ObjectClass).getMust(req))
		}
	}
	if !r.Must.IsZero() {
		for _, must := range r.Must {
			req.Set(must)
		}
	}

	return
}

func (r *NameForm) validateMustMay() (err error) {
	if r.Must.IsZero() {
		err = raise(invalidNameForm,
			"validateMustMay: missing %T.Must (%T)",
			r, r.Must)
	}

	return
}

func (r *NameForm) validateStructuralObjectClass() (err error) {
	if r.OC.IsZero() {
		err = raise(invalidNameForm,
			"validateOC: Missing %T", r.OC)
	} else if !r.OC.Kind.is(Structural) {
		err = raise(invalidNameForm,
			"validateOC: %T lacks an embedded STRUCTURAL type (is '%s')",
			r.OC, r.OC.Kind.String())
	}

	if err == nil {
		err = r.OC.validate() // generalized validation
	}

	return
}

/*
getSyntax will traverse the supertype chain upwards until it finds an explicit SYNTAX definition
*/
func (r *AttributeType) getSyntax() *LDAPSyntax {
	if r.IsZero() {
		return nil
	}
	if r.Syntax.IsZero() {
		return r.SuperType.getSyntax()
	}

	return r.Syntax
}

func (r *AttributeType) validateSyntax() (ls *LDAPSyntax, err error) {
	ls = r.getSyntax()
	if ls.IsZero() {
		err = raise(invalidSyntax,
			"checkMatchingRules: %T is missing a syntax", r)
	}

	return
}

func (r *AttributeType) validateMatchingRules(ls *LDAPSyntax) (err error) {
	if err = r.validateEquality(ls); err != nil {
		return err
	}

	if err = r.validateOrdering(ls); err != nil {
		return err
	}

	if err = r.validateSubstr(ls); err != nil {
		return err
	}

	return
}

func (r *AttributeType) validateEquality(ls *LDAPSyntax) error {
	if !r.Equality.IsZero() {
		if contains(toLower(r.Equality.Name.Value(0)), `ordering`) ||
			contains(toLower(r.Equality.Name.Value(0)), `substring`) {
			return raise(invalidMatchingRule,
				"validateEquality: %T.Equality uses non-equality %T syntax (%s)",
				r, r.Equality, r.Equality.Syntax.OID.string())
		}
	}

	return nil
}

func (r *AttributeType) validateSubstr(ls *LDAPSyntax) error {
	if !r.Substring.IsZero() {
		if !contains(toLower(r.Substring.Name.Value(0)), `substring`) {
			return raise(invalidMatchingRule,
				"validateSubstr: %T.Substring uses non-substring %T syntax (%s)",
				r, r.Substring, r.Substring.Syntax.OID.string())
		}
	}

	return nil
}

func (r *AttributeType) validateOrdering(ls *LDAPSyntax) error {
	if !r.Ordering.IsZero() {
		if !contains(toLower(r.Ordering.Name.Value(0)), `ordering`) {
			return raise(invalidMatchingRule,
				"validateOrdering: %T.Ordering uses non-substring %T syntax (%s)",
				r, r.Ordering, r.Ordering.Syntax.OID.string())
		}
	}

	return nil
}

func validateBool(b Boolean) (err error) {
	if b.is(Collective) && b.is(SingleValue) {
		return raise(invalidBoolean,
			"Cannot have single-valued collective attribute")
	}

	return
}

func (r *ObjectClass) validateKind() (err error) {
	if newKind(r.Kind.String()) == badKind {
		err = invalidObjectClassKind
	}

	return
}

func validateNames(value ...string) (err error) {
	// check slice as a whole
	switch {
	case len(value) > nameListMaxLen:
		err = raise(invalidName,
			"validation error: length of slices exceeds limit (%d)",
			nameListMaxLen)
	case len(value) == 0:
		err = raise(invalidName,
			"validation error: zero-length slices detected")
	}

	if err != nil {
		return
	}

	// check slice members
	for _, n := range value {
		if err = validateName(n); err != nil {
			return
		}
	}

	return
}

func validateName(value string) (err error) {
	switch {
	case len(value) > nameMaxLen:
		err = raise(invalidName,
			"validation error: length of slice exceeds limit (%d)",
			nameMaxLen)
	case len(value) == 0:
		err = raise(invalidName, "validateName: zero-length name")
	}

	for _, c := range value {
		if err != nil {
			break
		}

		switch {
		case !runeIsLetter(c) && !runeIsDigit(c) && c != '-':
			err = raise(invalidName,
				"validation error: bad char '%c'", c)
		}
	}

	return
}

func validateDesc(x interface{}) (err error) {
	var value string
	switch tv := x.(type) {
	case Description:
		value = string(tv)
	case []byte:
		value = string(tv)
	case string:
		value = tv
	default:
		return raise(invalidDescription,
			"validateDesc: unexpected type %T", tv)
	}

	switch {
	case len(value) > descMaxLen:
		err = raise(invalidDescription,
			"validation error: length exceeds limit (%d)",
			descMaxLen)
	case len(value) > 0 && !isUTF8([]byte(value)):
		err = raise(invalidDescription,
			"validation error: not UTF-8")
	}

	return
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

func isNumericalOID(x interface{}) (is bool) {
	var val string
	switch tv := x.(type) {
	case string:
		val = stripTags(tv)
	case OID:
		val = string(tv)
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

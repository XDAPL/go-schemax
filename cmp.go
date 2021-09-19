package schemax

/*
Equals compares the receiver to the provided interface value (which must be of the same effective type).  A comparison is performed, and an equality-indicative boolean value is returned.
*/
func (r Equality) Equals(x interface{}) (equals bool) {
	z, ok := x.(Equality)
	if !ok {
		return
	}

	if z.IsZero() && r.IsZero() {
		equals = true
		return
	} else if z.IsZero() || r.IsZero() {
		return
	}

	if z.OID != r.OID {
		return
	}

	if !z.Name.Equals(r.Name) {
		return
	}

	if !z.Syntax.Equals(r.Syntax) {
		return
	}

	if z.bools != r.bools {
		return
	}
	equals = r.Extensions.Equals(z.Extensions)

	return
}

/*
Equals compares the receiver to the provided interface value (which must be of the same effective type).  A comparison is performed, and an equality-indicative boolean value is returned.
*/
func (r Ordering) Equals(x interface{}) (equals bool) {
	z, ok := x.(Ordering)
	if !ok {
		return
	}

	if z.IsZero() && r.IsZero() {
		equals = true
		return
	} else if z.IsZero() || r.IsZero() {
		return
	}

	if z.OID != r.OID {
		return
	}

	if !z.Name.Equals(r.Name) {
		return
	}

	if !z.Syntax.Equals(r.Syntax) {
		return
	}

	if z.bools != r.bools {
		return
	}
	equals = r.Extensions.Equals(z.Extensions)

	return
}

/*
Equals compares the receiver to the provided interface value (which must be of the same effective type).  A comparison is performed, and an equality-indicative boolean value is returned.
*/
func (r Substring) Equals(x interface{}) (equals bool) {
	z, ok := x.(Substring)
	if !ok {
		return
	}

	if z.IsZero() && r.IsZero() {
		equals = true
		return
	} else if z.IsZero() || r.IsZero() {
		return
	}

	if z.OID != r.OID {
		return
	}

	if !z.Name.Equals(r.Name) {
		return
	}

	if !z.Syntax.Equals(r.Syntax) {
		return
	}

	if z.bools != r.bools {
		return
	}
	equals = r.Extensions.Equals(z.Extensions)

	return
}

/*
Equals compares the receiver to the provided interface value (which must be of the same effective type).  A comparison is performed, and an equality-indicative boolean value is returned.
*/
func (r *MatchingRuleUse) Equals(x interface{}) (equals bool) {
	z, ok := x.(MatchingRuleUse)
	if !ok {
		return
	}

	if z.IsZero() && r.IsZero() {
		equals = true
		return
	} else if z.IsZero() || r.IsZero() {
		return
	}

	if z.OID != r.OID {
		return
	}

	if !z.Name.Equals(r.Name) {
		return
	}

	if z.bools != r.bools {
		return
	}

	if !z.Applies.Equals(r.Applies) {
		return
	}
	equals = r.Extensions.Equals(z.Extensions)

	return
}

/*
Equals returns a boolean value indicative of the test result involving the receiver instance and x.  This test is meant to ascertain if the two instances represent the same type and have the same effective values.
*/
func (r *SuperiorDITStructureRules) Equals(x interface{}) (equals bool) {
	z, ok := x.(*SuperiorDITStructureRules)
	if !ok {
		return
	}

	if z.IsZero() && r.IsZero() {
		equals = true
		return
	} else if z.IsZero() || r.IsZero() {
		return
	}

	if len(*z) != len(*r) {
		return
	}

	R := *r
	for i := 0; i < len(R); i++ {
		if z2, ok := R[i].(*DITStructureRule); ok {
			A := *z
			if !z2.Equals(A[i].(*DITStructureRule)) {
				return
			}
		}
	}
	equals = true

	return
}

/*
Equals returns a boolean value indicative of the test result involving the receiver instance and x.  This test is meant to ascertain if the two instances represent the same type and have the same effective values.
*/
func (r *AuxiliaryObjectClasses) Equals(o interface{}) (equals bool) {
	z, ok := o.(*AuxiliaryObjectClasses)
	if !ok {
		return
	}

	if z.IsZero() && r.IsZero() {
		equals = true
		return
	} else if z.IsZero() || r.IsZero() {
		return
	}

	if len(*z) != len(*r) {
		return
	}

	R := *r
	for i := 0; i < len(R); i++ {
		if z2, ok := R[i].(*ObjectClass); ok {
			A := *z
			if !z2.Equals(A[i].(*ObjectClass)) {
				return
			}
		}
	}
	equals = true

	return
}

/*
Equals returns a boolean value indicative of the test result involving the receiver instance and x.  This test is meant to ascertain if the two instances represent the same type and have the same effective values.
*/
func (r *SuperiorObjectClasses) Equals(o interface{}) (equals bool) {
	z, ok := o.(*SuperiorObjectClasses)
	if !ok {
		return
	}

	if z.IsZero() && r.IsZero() {
		equals = true
		return
	} else if z.IsZero() || r.IsZero() {
		return
	}

	if len(*z) != len(*r) {
		return
	}

	R := *r
	for i := 0; i < len(R); i++ {
		if z2, ok := R[i].(*ObjectClass); ok {
			A := *z
			if !z2.Equals(A[i].(*ObjectClass)) {
				return
			}
		}
	}
	equals = true

	return
}

/*
Equals returns a boolean value indicative of the test result involving the receiver instance and x.  This test is meant to ascertain if the two instances represent the same type and have the same effective values.
*/
func (r *Applies) Equals(o interface{}) (equals bool) {
	z, ok := o.(*Applies)
	if !ok {
		return
	}

	if z.IsZero() && r.IsZero() {
		equals = true
		return
	} else if z.IsZero() || r.IsZero() {
		return
	}

	if len(*z) != len(*r) {
		return
	}

	R := *r
	for i := 0; i < len(R); i++ {
		if z2, ok := R[i].(*AttributeType); ok {
			A := *z
			if !z2.Equals(A[i].(*AttributeType)) {
				return
			}
		}
	}
	equals = true

	return
}

/*
Equals returns a boolean value indicative of the test result involving the receiver instance and x.  This test is meant to ascertain if the two instances represent the same type and have the same effective values.
*/
func (r *RequiredAttributeTypes) Equals(o interface{}) (equals bool) {
	z, ok := o.(*RequiredAttributeTypes)
	if !ok {
		return
	}

	if z.IsZero() && r.IsZero() {
		equals = true
		return
	} else if z.IsZero() || r.IsZero() {
		return
	}

	if len(*z) != len(*r) {
		return
	}

	R := *r
	for i := 0; i < len(R); i++ {
		if z2, ok := R[i].(*AttributeType); ok {
			A := *z
			if !z2.Equals(A[i].(*AttributeType)) {
				return
			}
		}
	}
	equals = true

	return
}

/*
Equals returns a boolean value indicative of the test result involving the receiver instance and x.  This test is meant to ascertain if the two instances represent the same type and have the same effective values.
*/
func (r *ProhibitedAttributeTypes) Equals(o interface{}) (equals bool) {
	z, ok := o.(*ProhibitedAttributeTypes)
	if !ok {
		return
	}

	if z.IsZero() && r.IsZero() {
		equals = true
		return
	} else if z.IsZero() || r.IsZero() {
		return
	}

	if len(*z) != len(*r) {
		return
	}

	R := *r
	for i := 0; i < len(R); i++ {
		if z2, ok := R[i].(*AttributeType); ok {
			A := *z
			if !z2.Equals(A[i].(*AttributeType)) {
				return
			}
		}
	}
	equals = true

	return
}

/*
Equals returns a boolean value indicative of the test result involving the receiver instance and x.  This test is meant to ascertain if the two instances represent the same type and have the same effective values.
*/
func (r *PermittedAttributeTypes) Equals(o interface{}) (equals bool) {
	z, ok := o.(*PermittedAttributeTypes)
	if !ok {
		return
	}

	if z.IsZero() && r.IsZero() {
		equals = true
		return
	} else if z.IsZero() || r.IsZero() {
		return
	}

	if len(*z) != len(*r) {
		return
	}

	R := *r
	for i := 0; i < len(R); i++ {
		if z2, ok := R[i].(*AttributeType); ok {
			A := *z
			if !z2.Equals(A[i].(*AttributeType)) {
				return
			}
		}
	}
	equals = true

	return
}

/*
Equals compares the receiver to the provided interface value (which must be of the same effective type).  A comparison is performed, and an equality-indicative boolean value is returned.
*/
func (r MatchingRuleUsesManifest) Equals(x interface{}) (equals bool) {
	z, ok := x.(MatchingRuleUsesManifest)
	if !ok {
		return
	}

	if z.IsZero() || r.IsZero() {
		return
	}

	if len(r) != len(z) {
		return
	}

	for k, v := range r {
		if v2, exists := z[k]; !exists {
			return
		} else {
			if !v.Equals(v2) {
				return
			}
		}
	}
	equals = true

	return
}

/*
Equals compares the receiver to the provided interface value (which must be of the same effective type).  A comparison is performed, and an equality-indicative boolean value is returned.
*/
func (r DITStructureRulesManifest) Equals(x interface{}) (equals bool) {
	z, ok := x.(DITStructureRulesManifest)
	if !ok {
		return
	}

	if z.IsZero() || r.IsZero() {
		return
	}

	if len(r) != len(z) {
		return
	}

	for k, v := range r {
		if v2, exists := z[k]; !exists {
			return
		} else {
			if !v.Equals(v2) {
				return
			}
		}
	}
	equals = true

	return
}

/*
Equals compares the receiver to the provided interface value (which must be of the same effective type).  A comparison is performed, and an equality-indicative boolean value is returned.
*/
func (r DITContentRulesManifest) Equals(x interface{}) (equals bool) {
	z, ok := x.(DITContentRulesManifest)
	if !ok {
		return
	}

	if z.IsZero() && r.IsZero() {
		equals = true
		return
	} else if z.IsZero() || r.IsZero() {
		return
	}

	if len(r) != len(z) {
		return
	}

	for k, v := range r {
		if v2, exists := z[k]; !exists {
			return
		} else {
			if !v.Equals(v2) {
				return
			}
		}
	}
	equals = true

	return
}

/*
Equals compares the receiver to the provided interface value (which must be of the same effective type).  A comparison is performed, and an equality-indicative boolean value is returned.
*/
func (r AttributeTypesManifest) Equals(x interface{}) (equals bool) {
	z, ok := x.(AttributeTypesManifest)
	if !ok {
		return
	}

	if z.IsZero() && r.IsZero() {
		equals = true
		return
	} else if z.IsZero() || r.IsZero() {
		return
	}

	if len(r) != len(z) {
		return
	}

	for k, v := range r {
		if v2, exists := z[k]; !exists {
			return
		} else {
			if !v.Equals(v2) {
				return
			}
		}
	}
	equals = true

	return
}

/*
Equals compares the receiver to the provided interface value (which must be of the same effective type).  A comparison is performed, and an equality-indicative boolean value is returned.
*/
func (r ObjectClassesManifest) Equals(x interface{}) (equals bool) {
	z, ok := x.(ObjectClassesManifest)
	if !ok {
		return
	}

	if z.IsZero() && r.IsZero() {
		equals = true
		return
	} else if z.IsZero() || r.IsZero() {
		return
	}

	if len(r) != len(z) {
		return
	}

	for k, v := range r {
		if v2, exists := z[k]; !exists {
			return
		} else {
			if !v.Equals(v2) {
				return
			}
		}
	}
	equals = true

	return
}

/*
Equals compares the receiver to the provided interface value (which must be of the same effective type).  A comparison is performed, and an equality-indicative boolean value is returned.
*/
func (r MatchingRulesManifest) Equals(x interface{}) (equals bool) {
	z, ok := x.(MatchingRulesManifest)
	if !ok {
		return
	}

	if z.IsZero() && r.IsZero() {
		equals = true
		return
	} else if z.IsZero() || r.IsZero() {
		return
	}

	if len(r) != len(z) {
		return
	}

	for k, v := range r {
		if v2, exists := z[k]; !exists {
			return
		} else {
			if !v.Equals(v2) {
				return
			}
		}
	}
	equals = true

	return
}

/*
Equals compares the receiver to the provided interface value (which must be of the same effective type).  A comparison is performed, and an equality-indicative boolean value is returned.
*/
func (r LDAPSyntaxesManifest) Equals(x interface{}) (equals bool) {
	z, ok := x.(LDAPSyntaxesManifest)
	if !ok {
		return
	}

	if z.IsZero() && r.IsZero() {
		equals = true
		return
	} else if z.IsZero() || r.IsZero() {
		return
	}

	if len(r) != len(z) {
		return
	}

	for k, v := range r {
		if v2, exists := z[k]; !exists {
			return
		} else {
			if !v.Equals(v2) {
				return
			}
		}
	}
	equals = true

	return
}

/*
Equals compares the receiver to the provided interface value (which must be of the same effective type).  A comparison is performed, and an equality-indicative boolean value is returned.
*/
func (r NameFormsManifest) Equals(x interface{}) (equals bool) {
	z, ok := x.(NameFormsManifest)
	if !ok {
		return
	}

	if z.IsZero() && r.IsZero() {
		equals = true
		return
	} else if z.IsZero() || r.IsZero() {
		return
	}

	if len(r) != len(z) {
		return
	}

	for k, v := range r {
		if v2, exists := z[k]; !exists {
			return
		} else {
			if !v.Equals(v2) {
				return
			}
		}
	}
	equals = true

	return
}

/*
Equals returns a boolean value indicative of the test result involving the receiver instance and x.  This test is meant to ascertain if the two instances represent the same type and have the same effective values.

Description text is ignored.
*/
func (r Name) Equals(x interface{}) (equals bool) {
	if r.IsZero() {
		return
	}

	switch tv := x.(type) {
	case nil:
		return
	case string:
		for _, z := range r {
			equals = equalFold(z.(string), tv)
			if equals {
				break
			}
		}
	case Name:
		for _, z := range r {
			for _, y := range tv {
				if equals = equalFold(z.(string), y.(string)); equals {
					return
				}
			}
		}
	}

	return
}

/*
Equals returns a boolean value indicative of the test result involving the receiver instance and x.  This test is meant to ascertain if the two instances represent the same type and have the same effective values.

Description text is ignored.
*/
func (r *AttributeType) Equals(x interface{}) (equals bool) {
	z, ok := x.(*AttributeType)
	if !ok {
		return
	}

	if z.IsZero() && r.IsZero() {
		equals = true
		return
	} else if z.IsZero() || r.IsZero() {
		return
	}

	if !z.Name.Equals(r.Name) {
		return
	}

	if r.OID != z.OID {
		return
	}

	if z.Usage != r.Usage {
		return
	}

	if z.bools != r.bools {
		return
	}

	if !z.SuperType.IsZero() && !r.SuperType.IsZero() {
		if z.SuperType.OID != r.SuperType.OID {
			return
		}
	}

	if !r.Syntax.Equals(z.Syntax) {
		return
	}

	if !r.Equality.Equals(z.Equality) {
		return
	}

	if !r.Ordering.Equals(z.Ordering) {
		return
	}
	equals = r.Extensions.Equals(z.Extensions)

	return
}

/*
Equals compares the receiver to the provided interface value (which must be of the same effective type).  A comparison is performed, and an equality-indicative boolean value is returned.
*/
func (r Extensions) Equals(x interface{}) (equals bool) {
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
Equals compares the receiver to the provided interface value (which must be of the same effective type).  A comparison is performed, and an equality-indicative boolean value is returned.

Description text is ignored.
*/
func (r *LDAPSyntax) Equals(x interface{}) bool {
	if assert, equals := x.(*LDAPSyntax); equals {
		return assert.OID.string() == r.OID.string()
	}

	return false
}

/*
Equals compares the receiver to the provided interface value (which must be of the same effective type).  A comparison is performed, and an equality-indicative boolean value is returned.

Description text is ignored.
*/
func (r *MatchingRule) Equals(x interface{}) bool {
	if assert, equals := x.(*MatchingRule); equals {
		return (assert.OID.string() == r.OID.string() &&
			assert.Name.Equals(r.Name))
	}

	return false
}

func (r Kind) is(x Kind) bool {
	return r == x
}

/*
is returns a boolean value indicative of whether the provided interface value is either a Kind or a Boolean AND is enabled within the receiver.
*/
func (r ObjectClass) is(b interface{}) bool {
	switch tv := b.(type) {
	case Boolean:
		return r.bools.is(tv)
	case Kind:
		return r.Kind.is(tv)
	}

	return false
}

/*
Equals compares the receiver to the provided interface value (which must be of the same effective type).  A comparison is performed, and an equality-indicative boolean value is returned.

Description text is ignored.
*/
func (r *ObjectClass) Equals(x interface{}) (equals bool) {

	z, ok := x.(*ObjectClass)
	if !ok {
		return
	}

	if z.IsZero() && r.IsZero() {
		equals = true
		return
	} else if z.IsZero() || r.IsZero() {
		return
	}

	if !z.Name.Equals(r.Name) {
		return
	}

	if r.Kind != z.Kind {
		return
	}

	if r.bools != z.bools {
		return
	}

	if !r.Must.Equals(z.Must) {
		return
	}

	if !r.May.Equals(z.May) {
		return
	}

	if !z.SuperClass.IsZero() && !r.SuperClass.IsZero() {
		if !r.SuperClass.Equals(z.SuperClass) {
			return
		}
	}

	equals = r.Extensions.Equals(z.Extensions)

	return
}

/*
Equals compares the receiver to the provided interface value (which must be of the same effective type).  A comparison is performed, and an equality-indicative boolean value is returned.

Description text is ignored.
*/
func (r *DITContentRule) Equals(x interface{}) (equals bool) {

	z, ok := x.(*DITContentRule)
	if !ok {
		return
	}

	if z.IsZero() && r.IsZero() {
		equals = true
		return
	} else if z.IsZero() || r.IsZero() {
		return
	}

	if r.OID != z.OID {
		return
	}

	if !r.Name.Equals(z.Name) {
		return
	}

	if !r.Aux.Equals(z.Aux) {
		return
	}

	if !r.Must.Equals(z.Must) {
		return
	}

	if !r.May.Equals(z.May) {
		return
	}

	if !r.Not.Equals(z.Not) {
		return
	}

	equals = r.Extensions.Equals(z.Extensions)

	return
}

/*
Equals compares the receiver to the provided interface value (which must be of the same effective type).  A comparison is performed, and an equality-indicative boolean value is returned.

Description text is ignored.
*/
func (r *DITStructureRule) Equals(x interface{}) (equals bool) {

	z, ok := x.(*DITStructureRule)
	if !ok {
		return
	}

	if z.IsZero() && r.IsZero() {
		equals = true
		return
	} else if z.IsZero() || r.IsZero() {
		return
	}

	if r.ID != z.ID {
		return
	}

	if !r.Name.Equals(z.Name) {
		return
	}

	if !r.Form.Equals(z.Form) {
		return
	}

	if !r.SuperiorRules.Equals(z.SuperiorRules) {
		return
	}

	equals = r.Extensions.Equals(z.Extensions)

	return
}

/*
Equals compares the receiver to the provided interface value (which must be of the same effective type).  A comparison is performed, and an equality-indicative boolean value is returned.

Description text is ignored.
*/
func (r *NameForm) Equals(x interface{}) (equals bool) {

	z, ok := x.(*NameForm)
	if !ok {
		return
	}

	if z.IsZero() && r.IsZero() {
		equals = true
		return
	} else if z.IsZero() || r.IsZero() {
		return
	}

	if r.OID != z.OID {
		return
	}

	if !r.Name.Equals(z.Name) {
		return
	}

	if !r.OC.Equals(z.OC) {
		return
	}

	if !r.Must.Equals(z.Must) {
		return
	}

	if !r.May.Equals(z.May) {
		return
	}

	equals = r.Extensions.Equals(z.Extensions)

	return
}

/*
IsHumanReadable returns a boolean value indicative of whether the receiver instance of *LDAPSyntax supports values that are human-readable.
*/
func (r LDAPSyntax) IsHumanReadable() bool {
	return r.bools.is(HumanReadable)
}

/*
is returns a boolean value indicative of whether the provided interface argument is an enabled Boolean option.
*/
func (r LDAPSyntax) is(b interface{}) bool {
	switch tv := b.(type) {
	case Boolean:
		return r.bools.is(tv)
	}

	return false
}

/*
is returns a boolean value indicative of whether the provided interface argument is an enabled Boolean option.
*/
func (r MatchingRule) is(b interface{}) bool {
	switch tv := b.(type) {
	case Boolean:
		return r.bools.is(tv)
	case *LDAPSyntax:
		return r.Syntax.OID == tv.OID
	}

	return false
}

/*
is returns a boolean value indicative of whether the provided interface argument is an enabled Boolean option.
*/
func (r MatchingRuleUse) is(b interface{}) bool {
	switch tv := b.(type) {
	case *AttributeType:
		if _, x := r.Applies.Index(tv); !x {
			return false
		}
	case Boolean:
		return r.bools.is(tv)
	}

	return false
}

/*
is returns a boolean value indicative of whether the provided interface argument is either an enabled Boolean value, or an associated *MatchingRule or *LDAPSyntax.

In the case of an *LDAPSyntax argument, if the receiver is in fact a sub type of another *AttributeType instance, a reference to that super type is chased and analyzed accordingly.
*/
func (r AttributeType) is(b interface{}) bool {
	switch tv := b.(type) {
	case Boolean:
		return r.bools.is(tv)
	case *MatchingRule:
		switch tv.OID {
		case r.Equality.OID, r.Ordering.OID, r.Substring.OID:
			return true
		}
	case *LDAPSyntax:
		if r.Syntax != nil {
			return r.Syntax.OID == tv.OID
		} else if !r.SuperType.IsZero() {
			return r.SuperType.is(tv)
		}
	}

	return false
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
Belongs returns a boolean value indicative of whether the provided AUXILIARY *ObjectClass belongs to the receiver instance of *DITContentRule.
*/
func (r DITContentRule) Belongs(aux *ObjectClass) (belongs bool) {
	if aux.IsZero() || !aux.Kind.is(Auxiliary) {
		return
	}
	_, belongs = r.Aux.Index(aux)

	return
}

/*
Requires returns a boolean value indicative of whether the provided value is required per the receiver.
*/
func (r DITContentRule) Requires(x interface{}) (required bool) {
	switch tv := x.(type) {
	case *AttributeType:
		_, required = r.Must.Index(tv)
	}

	return
}

/*
Permits returns a boolean value indicative of whether the provided value is allowed from use per the receiver.
*/
func (r DITContentRule) Permits(x interface{}) (permitted bool) {
	switch tv := x.(type) {
	case *AttributeType:
		_, permitted = r.May.Index(tv)
	}

	return
}

/*
Prohibits returns a boolean value indicative of whether the provided value is prohibited from use per the receiver.
*/
func (r DITContentRule) Prohibits(x interface{}) (prohibited bool) {
	switch tv := x.(type) {
	case *AttributeType:
		_, prohibited = r.Not.Index(tv)
	}

	return
}

/*
IsZero returns a boolean value indicating whether the receiver is unset.
*/
func (r *Boolean) IsZero() (zero bool) {
	if zero = r == nil; zero {
		return
	}
	zero = *r == 0

	return
}

/*
IsZero returns a boolean value indicating whether the receiver is unset.
*/
func (r OID) IsZero() bool {
	return len(r) == 0
}

/*
IsZero returns a boolean value indicating whether the receiver is unset.
*/
func (r Description) IsZero() bool {
	return len(r) == 0
}

/*
IsZero returns a boolean value indicating whether the receiver is unset.
*/
func (r Applies) IsZero() bool {
	return len(r) == 0
}

/*
IsZero returns a boolean value indicating whether the receiver is unset.
*/
func (r Usage) IsZero() bool {
	return r == 0 // implies userApplication
}

/*
IsZero returns a boolean value indicating whether the receiver is unset.
*/
func (r Kind) IsZero() bool {
	return r == 0
}

/*
IsZero returns a boolean value indicating whether the receiver is unset.
*/
func (r Extensions) IsZero() (zero bool) {
	if zero = r == nil; zero {
		return
	}
	zero = len(r) == 0

	return
}

/*
IsZero returns a boolean value indicating whether the receiver is unset.
*/
func (r AliasesManifest) IsZero() (zero bool) {
	if zero = r == nil; zero {
		return
	}
	zero = len(r) == 0

	return
}

/*
IsZero returns a boolean value indicating whether the receiver is unset.
*/
func (r *Name) IsZero() (zero bool) {
	if zero = r == nil; zero {
		return
	}

	for _, v := range *r {
		if len(v.(string)) != 0 {
			return
		}
	}
	zero = true

	return
}

/*
IsZero returns a boolean value indicating whether the receiver is unset.
*/
func (r *LDAPSyntax) IsZero() (zero bool) {
	if zero = r == nil; zero {
		return
	}
	zero = (r.OID.IsZero() && r.Description.IsZero())

	return
}

/*
IsZero returns a boolean value indicating whether the receiver is unset.
*/
func (r *MatchingRule) IsZero() (zero bool) {
	if zero = r == nil; zero {
		return
	}
	zero = (r.OID.IsZero() && r.Name.IsZero())

	return
}

/*
IsZero returns a boolean value indicating whether the receiver is unset.
*/
func (r *MatchingRuleUse) IsZero() (zero bool) {
	if zero = r == nil; zero {
		return
	}
	zero = (r.OID.IsZero() && r.Name.IsZero())

	return
}

/*
IsZero returns a boolean value indicating whether the receiver is unset.
*/
func (r *ObjectClass) IsZero() (zero bool) {
	if zero = r == nil; zero {
		return
	}
	zero = (r.OID.IsZero() && r.Name.IsZero() &&
		r.Kind.IsZero() && r.SuperClass.IsZero())

	return
}

/*
IsZero returns a boolean value indicating whether the receiver is unset.
*/
func (r *AttributeType) IsZero() bool {
	return r == nil
}

/*
IsZero returns a boolean value indicating whether the receiver is unset.
*/
func (r *DITStructureRule) IsZero() (zero bool) {
	if zero = r == nil; zero {
		return
	}
	zero = (r.Form.IsZero() && r.Name.IsZero())

	return
}

/*
IsZero returns a boolean value indicating whether the receiver is unset.
*/
func (r *DITContentRule) IsZero() (zero bool) {
	if zero = r == nil; zero {
		return
	}
	zero = (r.Aux.IsZero() && r.Name.IsZero())

	return
}

/*
IsZero returns a boolean value indicating whether the receiver is unset.
*/
func (r *NameForm) IsZero() (zero bool) {
	if zero = r == nil; zero {
		return
	}
	zero = (r.OID.IsZero() && r.OC.IsZero())

	return
}

/*
IsZero returns a boolean value indicating whether the receiver is unset.
*/
func (r *SuperiorObjectClasses) IsZero() (zero bool) {
	if zero = r == nil; zero {
		return
	}
	zero = len(*r) == 0

	return
}

/*
IsZero returns a boolean value indicating whether the receiver is unset.
*/
func (r *AuxiliaryObjectClasses) IsZero() (zero bool) {
	if zero = r == nil; zero {
		return
	}
	zero = len(*r) == 0

	return
}

/*
IsZero returns a boolean value indicating whether the receiver is unset.
*/
func (r *RequiredAttributeTypes) IsZero() (zero bool) {
	if zero = r == nil; zero {
		return
	}
	zero = len(*r) == 0

	return
}

/*
IsZero returns a boolean value indicating whether the receiver is unset.
*/
func (r *PermittedAttributeTypes) IsZero() (zero bool) {
	if zero = r == nil; zero {
		return
	}
	zero = len(*r) == 0

	return
}

/*
IsZero returns a boolean value indicating whether the receiver is unset.
*/
func (r *ProhibitedAttributeTypes) IsZero() (zero bool) {
	if zero = r == nil; zero {
		return
	}
	zero = len(*r) == 0

	return
}

/*
IsZero returns a boolean value indicating whether the receiver is unset.
*/
func (r *SuperiorDITStructureRules) IsZero() (zero bool) {
	if zero = r == nil; zero {
		return
	}
	zero = len(*r) == 0

	return
}

/*
IsZero returns a boolean value indicating whether the receiver is unset.
*/
func (r DITContentRulesManifest) IsZero() (zero bool) {
	if zero = r == nil; zero {
		return
	}

	zero = len(r) == 0

	return
}

/*
IsZero returns a boolean value indicating whether the receiver is unset.
*/
func (r DITStructureRulesManifest) IsZero() (zero bool) {
	if zero = r == nil; zero {
		return
	}

	zero = len(r) == 0

	return
}

/*
IsZero returns a boolean value indicating whether the receiver is unset.
*/
func (r MatchingRuleUsesManifest) IsZero() (zero bool) {
	if zero = r == nil; zero {
		return
	}

	zero = len(r) == 0

	return
}

/*
IsZero returns a boolean value indicating whether the receiver is unset.
*/
func (r MatchingRulesManifest) IsZero() (zero bool) {
	if zero = r == nil; zero {
		return
	}

	zero = len(r) == 0

	return
}

/*
IsZero returns a boolean value indicating whether the receiver is unset.
*/
func (r LDAPSyntaxesManifest) IsZero() (zero bool) {
	if zero = r == nil; zero {
		return
	}

	zero = len(r) == 0

	return
}

/*
IsZero returns a boolean value indicating whether the receiver is unset.
*/
func (r ObjectClassesManifest) IsZero() (zero bool) {
	if zero = r == nil; zero {
		return
	}

	zero = len(r) == 0

	return
}

/*
IsZero returns a boolean value indicating whether the receiver is unset.
*/
func (r AttributeTypesManifest) IsZero() (zero bool) {
	if zero = r == nil; zero {
		return
	}

	zero = len(r) == 0

	return
}

/*
IsZero returns a boolean value indicating whether the receiver is unset.
*/
func (r NameFormsManifest) IsZero() (zero bool) {
	if zero = r == nil; zero {
		return
	}

	zero = len(r) == 0

	return
}

package schemax

type collection []interface{}

/*
len returns the length of the receiver as an int.
*/
func (c *collection) len() int {
	return len(*c)
}

/*
index is a panic-proof slice indexer that returns an interface member based on the idx integer argument. This method is not thread-safe unto itself, and should only be called in situations where thread protection is provided at a higher level.

A negative index integer returns the element at index 'length+-idx'. A positive integer returns the nth slice member. If the provided integer is larger than the total length, the final element is returned.
*/
func (c collection) index(idx int) interface{} {
	L := c.len()
	if L == 0 {
		return nil
	}

	if idx < 0 {
		var x int = L + idx
		if x < 0 {
			return c[0]
		} else {
			return c[x]
		}
	} else if idx > L {
		return c[L-1]
	} else {
		return c[idx]
	}

	return nil
}

func (r collection) isZero() bool {
	if &r == nil {
		return true
	}
	return r.len() == 0
}

/*
Equal returns a boolean value indicative of the test result involving the receiver instance and x.  This test is meant to ascertain if the two instances represent the same type and have the same effective values.
*/
func (r collection) equal(x collection) (equals bool) {
	if r.len() != x.len() {
		return
	}

	if x.isZero() && r.isZero() {
		equals = true
		return
	} else if x.isZero() || r.isZero() {
		return
	}

	for i := 0; i < r.len(); i++ {
		switch tv := r.index(i).(type) {
		case *LDAPSyntax:
			assert, ok := x.index(i).(*LDAPSyntax)
			if !ok {
				return
			}
			if !tv.Equal(assert) {
				return
			}
		case *MatchingRule:
			assert, ok := x.index(i).(*MatchingRule)
			if !ok {
				return
			}
			if !tv.Equal(assert) {
				return
			}
		case *MatchingRuleUse:
			assert, ok := x.index(i).(*MatchingRuleUse)
			if !ok {
				return
			}
			if !tv.Equal(assert) {
				return
			}
		case *AttributeType:
			assert, ok := x.index(i).(*AttributeType)
			if !ok {
				return
			}
			if !tv.Equal(assert) {
				return
			}
		case *ObjectClass:
			assert, ok := x.index(i).(*ObjectClass)
			if !ok {
				return
			}
			if !tv.Equal(assert) {
				return
			}
		case *NameForm:
			assert, ok := x.index(i).(*NameForm)
			if !ok {
				return
			}
			if !tv.Equal(assert) {
				return
			}
		case *DITContentRule:
			assert, ok := x.index(i).(*DITContentRule)
			if !ok {
				return
			}
			if !tv.Equal(assert) {
				return
			}
		case *DITStructureRule:
			assert, ok := x.index(i).(*DITStructureRule)
			if !ok {
				return
			}
			if !tv.Equal(assert) {
				return
			}
		case string:
			assert, ok := x.index(i).(string)
			if !ok {
				return
			}
			if tv != assert {
				return
			}
		default:
			return
		}
	}

	equals = true
	return
}

/*
containsOID returns an element index number and a boolean value indicative of the result of a search by OID conducted upon the receiver. This method is not thread-safe unto itself, and should only be called in situations where thread protection is provided at a higher level.

dITStructureRules do not possess an OID, and thus collections of that definition type have no meaningful use of this method.
*/
func (c collection) containsOID(x OID) (index int, contains bool) {
	index = -1

	for i := 0; i < c.len(); i++ {
		el := c.index(i)
		if el == nil {
			continue
		}

		switch tv := el.(type) {
		case *LDAPSyntax:
			contains = tv.OID.Equal(x)
		case *MatchingRule:
			contains = tv.OID.Equal(x)
		case *AttributeType:
			contains = tv.OID.Equal(x)
		case *MatchingRuleUse:
			contains = tv.OID.Equal(x)
		case *ObjectClass:
			contains = tv.OID.Equal(x)
		case *DITContentRule:
			contains = tv.OID.Equal(x)
		case *NameForm:
			contains = tv.OID.Equal(x)
		default:
			return
		}

		if contains {
			index = i
			break
		}
	}

	return
}

/*
containsID returns an element index number and a boolean value indicative of the result of a search by Rule ID conducted upon the receiver. This method is not thread-safe unto itself, and should only be called in situations where thread protection is provided at a higher level.

collections of dITStructureRules are the only types that will have any meaningful use for this method.
*/
func (c collection) containsID(x interface{}) (index int, found bool) {
	index = -1

	for i := 0; i < c.len(); i++ {
		el := c.index(i)
		if el == nil {
			continue
		}

		switch tve := el.(type) {
		case *DITStructureRule:
			found = tve.ID.Equal(NewRuleID(x))
		default:
			return
		}

		if found {
			index = i
			break
		}
	}

	return
}

/*
containsName returns an element index number and a boolean value indicative of the result of a search by a string name conducted upon the receiver. This method is not thread-safe unto itself, and should only be called in situations where thread protection is provided at a higher level.

ldapSyntax definitions do not possess a Name, and thus collections of that definition type have no meaningful use of this method.
*/
func (c collection) containsName(x string) (index int, found bool) {
	index = -1

	for i := 0; i < c.len(); i++ {
		el := c.index(i)
		if el == nil {
			continue
		}

		switch tve := el.(type) {
		case *MatchingRule:
			found = tve.Name.Equal(x)
		case *AttributeType:
			found = tve.Name.Equal(x)
		case *MatchingRuleUse:
			found = tve.Name.Equal(x)
		case *ObjectClass:
			found = tve.Name.Equal(x)
		case *DITStructureRule:
			found = tve.Name.Equal(x)
		case *DITContentRule:
			found = tve.Name.Equal(x)
		case *NameForm:
			found = tve.Name.Equal(x)
		default:
			return
		}

		if found {
			index = i
			break
		}
	}

	return
}

func (c collection) containsDesc(x string) (index int, found bool) {
	index = -1

	for i := 0; i < c.len(); i++ {
		el := c.index(i)
		if el == nil {
			continue
		}

		switch tve := el.(type) {
		case *LDAPSyntax:
			// Syntaxes don't really have names like all other
			// definition types, so instead we have to think a
			// bit creatively and use the syntax DESC field.
			desc := toLower(replaceAll(string(tve.Description), ` `, ``))
			desc = replaceAll(desc, `syntax`, ``)
			x = toLower(replaceAll(x, ` `, ``))
			x = replaceAll(x, `syntax`, ``)
			if contains(desc, x) {
				found = true
				break
			}
		default:
			return
		}

		if found {
			index = i
			break
		}
	}

	return
}

/*
contains returns an element index number and a boolean value indicative of the result of a search conducted upon the receiver. This method is not thread-safe unto itself, and should only be called in situations where thread protection is provided at a higher level.

Possible search terms are:

  - string name (does not apply to collections of ldapSyntax definitions)
  - Actual OID, or string representation of OID (does not apply to collections of dITStructureRule definitions)
  - Rule ID, as an int or uint (only applies to collections of dITStructureRule definitions)
*/
func (c collection) contains(x interface{}) (index int, found bool) {
	index = -1
	if c.len() == 0 {
		return
	}

	switch tv := x.(type) {
	case OID:
		// Everything EXCEPT dITStructureRules
		index, found = c.containsOID(tv)
	case string:
		if isDigit(tv) {
			index, found = c.containsID(tv)
		} else if isNumericalOID(tv) {
			// Everything EXCEPT dITStructureRules
			index, found = c.containsOID(NewOID(tv))
		} else {
			// Everything EXCEPT ldapSyntaxes
			index, found = c.containsName(tv)

			// Try to satisfy LDAPSyntax searches based on
			// DESC field IF an OID is *NOT* being used.
			if !found {
				index, found = c.containsDesc(tv)
			}
		}
	case uint, int:
		// ONLY dITStructureRules
		index, found = c.containsID(tv)
	}

	return
}

/*
append assigns the provided interface value to the receiver. An error is returned if there is a type-mismatch, or if an unsupported type is provided. This method is not thread-safe unto itself, and should only be called in situations where thread protection is provided at a higher level.
*/
func (c *collection) append(x interface{}) error {
	switch tv := x.(type) {
	case *LDAPSyntax,
		*MatchingRule,
		*Equality,
		*Substring,
		*Ordering,
		*AttributeType,
		*SuperiorAttributeType,
		*MatchingRuleUse,
		*ObjectClass,
		*StructuralObjectClass,
		*DITContentRule,
		*NameForm,
		*DITStructureRule:
		// ok
	default:
		return raise(unexpectedType, "Unsupported type (%T) for collection append", tv)
	}

	// If there is at least one element, make sure we
	// only appending new elements of the same type.
	if c.len() > 0 {
		first := c.index(0)
		if sprintf("%T", first) != sprintf("%T", x) {
			return raise(unexpectedType, "Unsupported type mixture: cannot append %T to %T-based collection", x, first)
		}
	}

	*c = append(*c, x)
	return nil
}

func (r collection) attrs_oids_string() (str string) {
	switch len(r) {
	case 0:
		return
	case 1:
		assert, ok := r[0].(*AttributeType)
		if !ok {
			return
		}
		return assert.Name.Index(0)
	}

	str += `( `
	for i := 0; i < len(r); i++ {
		assert, ok := r[i].(*AttributeType)
		if !ok {
			return ``
		}
		str += assert.Name.Index(0) + ` $ `
	}
	if str[len(str)-3:] == ` $ ` {
		str = str[:len(str)-2]
	}
	str += `)`

	return
}

func (r collection) dsrs_ids_string() (str string) {
	switch len(r) {
	case 0:
		return
	case 1:
		assert, ok := r[0].(*DITStructureRule)
		if !ok {
			return
		}
		return assert.ID.String()
	}

	str += `( `
	for i := 0; i < len(r); i++ {
		assert, ok := r[i].(*DITStructureRule)
		if !ok {
			return ``
		}
		str += assert.ID.String() + ` $ `
	}
	if str[len(str)-3:] == ` $ ` {
		str = str[:len(str)-2]
	}
	str += `)`

	return
}

func (r collection) ocs_oids_string() (str string) {
	switch len(r) {
	case 0:
		return
	case 1:
		assert, ok := r[0].(*ObjectClass)
		if !ok {
			return
		}
		return assert.Name.Index(0)
	}

	str += `( `
	for i := 0; i < len(r); i++ {
		assert, ok := r[0].(*ObjectClass)
		if !ok {
			return ``
		}
		str += assert.Name.Index(0) + ` $ `
	}
	if str[len(str)-3:] == ` $ ` {
		str = str[:len(str)-2]
	}
	str += `)`

	return
}

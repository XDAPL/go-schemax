package schemax

func (r genericList) set(x interface{}) genericList {
	return append(r, x)
}

/*
Value returns the xth value present within the receiver, or a zero-length string if no value(s) were set, or if x is invalid in some fashion.

The returned value is NOT encapsulated in single quotes ("'").
*/
func (r Name) Value(x int) (value string) {
	if r.IsZero() || x < 0 || x > len(r)-1 {
		return
	}

	return r[x].(string)
}

/*
Value returns the xth value present within the receiver, or a zero-length string if no value(s) were set, or if x is invalid in some fashion.
*/
func (r *SuperiorObjectClasses) Value(x int) *ObjectClass {
	if r.IsZero() || x < 0 || x > len(*r)-1 {
		return nil
	}
	N := *r

	return N[x].(*ObjectClass)
}

/*
Value returns the xth value present within the receiver, or a zero-length string if no value(s) were set, or if x is invalid in some fashion.
*/
func (r *AuxiliaryObjectClasses) Value(x int) *ObjectClass {
	if r.IsZero() || x < 0 || x > len(*r)-1 {
		return nil
	}
	N := *r

	return N[x].(*ObjectClass)
}

/*
Value returns the xth value present within the receiver, or a zero-length string if no value(s) were set, or if x is invalid in some fashion.
*/
func (r *PermittedAttributeTypes) Value(x int) *AttributeType {
	if r.IsZero() || x < 0 || x > len(*r)-1 {
		return nil
	}
	N := *r

	return N[x].(*AttributeType)
}

/*
Value returns the xth value present within the receiver, or a zero-length string if no value(s) were set, or if x is invalid in some fashion.
*/
func (r *RequiredAttributeTypes) Value(x int) *AttributeType {
	if r.IsZero() || x < 0 || x > len(*r)-1 {
		return nil
	}
	N := *r

	return N[x].(*AttributeType)
}

/*
Value returns the xth value present within the receiver, or a zero-length string if no value(s) were set, or if x is invalid in some fashion.
*/
func (r *ProhibitedAttributeTypes) Value(x int) *AttributeType {
	if r.IsZero() || x < 0 || x > len(*r)-1 {
		return nil
	}
	N := *r

	return N[x].(*AttributeType)
}

/*
Value returns the xth value present within the receiver, or a zero-length string if no value(s) were set, or if x is invalid in some fashion.
*/
func (r *SuperiorDITStructureRules) Value(x int) *DITStructureRule {
	if r.IsZero() || x < 0 || x > len(*r)-1 {
		return nil
	}
	N := *r

	return N[x].(*DITStructureRule)
}

/*
Index returns the integer value indicative of the index number that matches the Name or RuleID of the provided interface value.

A boolean value indicative of a successful index location is returned.  Additionally, in the event of a failed index retrieval, the index is set to -1.
*/
func (r *SuperiorDITStructureRules) Index(x interface{}) (idx int, has bool) {
	var rid RuleID
	var term string

	switch tv := x.(type) {
	case int:
		rid = NewRuleID(tv)
	case uint:
		rid = NewRuleID(tv)
	case string:
		if !isDigit(tv) {
			term = stripTags(tv)
		} else {
			rid = NewRuleID(tv)
		}
	case *DITStructureRule:
		if tv.IsZero() {
			return
		}
		rid = tv.ID
	default:
		return
	}

	for i, d := range *r {
		if z, ok := d.(*DITStructureRule); ok {
			if z.ID == rid || z.Name.Equals(term) {
				has = true
				idx = i
				return
			}
		}
	}
	idx = -1

	return
}

/*
Index returns the integer value indicative of the index number that matches the Name or OID of the provided interface value.

A boolean value indicative of a successful index location is returned.  Additionally, in the event of a failed index retrieval, the index is set to -1.
*/
func (r *SuperiorObjectClasses) Index(x interface{}) (idx int, has bool) {
	var term string
	idx = -1

	switch tv := x.(type) {
	case string:
		term = stripTags(tv)
	case *ObjectClass:
		if tv.IsZero() {
			return
		}
		term = tv.OID.string()
	default:
		return
	}

	for i, d := range *r {
		if z, ok := d.(*ObjectClass); ok {
			if string(z.OID) == term || z.Name.Equals(term) {
				has = true
				idx = i
				break
			}
			d = nil
		}
	}

	return
}

/*
Index returns the integer value indicative of the index number that matches the Name or OID of the provided interface value.

A boolean value indicative of a successful index location is returned.  Additionally, in the event of a failed index retrieval, the index is set to -1.
*/
func (r *AuxiliaryObjectClasses) Index(x interface{}) (idx int, has bool) {
	var term string
	idx = -1

	switch tv := x.(type) {
	case string:
		term = stripTags(tv)
	case *ObjectClass:
		if tv.IsZero() || !tv.Kind.is(Auxiliary) {
			return
		}
		term = tv.OID.string()
	default:
		return
	}

	for i, d := range *r {
		if z, ok := d.(*ObjectClass); ok {
			if !z.Kind.is(Auxiliary) {
				continue
			}
			if string(z.OID) == term || z.Name.Equals(term) {
				has = true
				idx = i
				break
			}
			d = nil
		}
	}

	return
}

/*
Index returns the integer value indicative of the index number that matches the Name or OID of the provided interface value.

A boolean value indicative of a successful index location is returned.  Additionally, in the event of a failed index retrieval, the index is set to -1.
*/
func (r *RequiredAttributeTypes) Index(x interface{}) (idx int, has bool) {
	var term string
	idx = -1

	switch tv := x.(type) {
	case string:
		term = stripTags(tv)
	case *AttributeType:
		if tv.IsZero() {
			return
		}
		term = tv.OID.string()
	default:
		return
	}

	for i, d := range *r {
		if z, ok := d.(*AttributeType); ok {
			if string(z.OID) == term || z.Name.Equals(term) {
				has = true
				idx = i
				break
			}
			d = nil
		}
	}

	return
}

/*
Index returns the integer value indicative of the index number that matches the Name or OID of the provided interface value.

A boolean value indicative of a successful index location is returned.  Additionally, in the event of a failed index retrieval, the index is set to -1.
*/
func (r *PermittedAttributeTypes) Index(x interface{}) (idx int, has bool) {
	var term string
	idx = -1

	switch tv := x.(type) {
	case string:
		term = stripTags(tv)
	case *AttributeType:
		if tv.IsZero() {
			return
		}
		term = tv.OID.string()
	default:
		return
	}

	for i, d := range *r {
		if z, ok := d.(*AttributeType); ok {
			if string(z.OID) == term || z.Name.Equals(term) {
				has = true
				idx = i
				break
			}
			d = nil
		}
	}

	return
}

/*
Index returns the integer value indicative of the index number that matches the Name or OID of the provided interface value.

A boolean value indicative of a successful index location is returned.  Additionally, in the event of a failed index retrieval, the index is set to -1.
*/
func (r ProhibitedAttributeTypes) Index(x interface{}) (idx int, has bool) {
	var term string
	idx = -1

	switch tv := x.(type) {
	case string:
		term = stripTags(tv)
	case *AttributeType:
		if tv.IsZero() {
			return
		}
		term = tv.OID.string()
	default:
		return
	}

	for i, d := range r {
		if z, ok := d.(*AttributeType); ok {
			if string(z.OID) == term || z.Name.Equals(term) {
				has = true
				idx = i
				break
			}
			d = nil
		}
	}

	return
}

/*
Index returns the integer value indicative of the index number that matches the provided interface value (a string name).

A boolean value indicative of a successful index location is returned.  Additionally, in the event of a failed index retrieval, the index is set to -1.
*/
func (r Name) Index(x interface{}) (idx int, has bool) {
	var term string
	idx = -1

	switch tv := x.(type) {
	case string:
		term = stripTags(tv)
	default:
		return
	}

	if r.IsZero() || len(term) == 0 {
		return
	}

	var n interface{}
	for idx, n = range r {
		if has = term == n.(string); has {
			break
		}
	}

	return
}

/*
Index returns the integer value indicative of the index number that matches the Name or OID of the provided interface value.

A boolean value indicative of a successful index location is returned.  Additionally, in the event of a failed index retrieval, the index is set to -1.
*/
func (r Applies) Index(x interface{}) (idx int, has bool) {
	var term string
	idx = -1

	switch tv := x.(type) {
	case string:
		term = stripTags(tv)
	case *AttributeType:
		term = tv.OID.string()
	default:
		return
	}

	for i, d := range r {
		if z, ok := d.(*AttributeType); ok {
			if string(z.OID) == term || z.Name.Equals(term) {
				has = true
				idx = i
				break
			}
			d = nil
		}
	}

	return
}

/*
Set applies one or more instances of *ObjectClass to the receiver.  Uniqueness is enforced by OID.
*/
func (r *AuxiliaryObjectClasses) Set(x ...interface{}) {
	for _, n := range x {
		if z, ok := n.(*ObjectClass); ok {
			if z.IsZero() {
				continue
			}
			if _, exists := r.Index(z); !exists {
				*r = append(*r, z)
			}
		}
	}

	R := make(AuxiliaryObjectClasses, len(*r), len(*r))
	for i, n := range *r {
		R[i] = n
	}
	*r = R
}

/*
Set applies one or more instances of *ObjectClass to the receiver.  Uniqueness is enforced by OID.
*/
func (r *SuperiorObjectClasses) Set(x ...interface{}) {
	for _, n := range x {
		if z, ok := n.(*ObjectClass); ok {
			if z.IsZero() {
				continue
			}
			if _, exists := r.Index(z); !exists {
				*r = append(*r, z)
			}
		}
	}

	R := make(SuperiorObjectClasses, len(*r), len(*r))
	for i, n := range *r {
		R[i] = n
	}
	*r = R
}

/*
Set applies one or more instances of *AttributeType to the receiver.  Uniqueness is enforced by OID.
*/
func (r *PermittedAttributeTypes) Set(x ...interface{}) {
	for _, n := range x {
		if z, ok := n.(*AttributeType); ok {
			if z.IsZero() {
				continue
			}
			if _, exists := r.Index(z); !exists {
				*r = append(*r, z)
			}
		}
	}

	R := make(PermittedAttributeTypes, len(*r), len(*r))
	for i, n := range *r {
		R[i] = n
	}
	*r = R
}

/*
Set applies one or more instances of *AttributeType to the receiver.  Uniqueness is enforced by OID.
*/
func (r *RequiredAttributeTypes) Set(x ...interface{}) {
	for _, n := range x {
		if z, ok := n.(*AttributeType); ok {
			if z.IsZero() {
				continue
			}
			if _, exists := r.Index(z); !exists {
				*r = append(*r, z)
			}
		}
	}

	R := make(RequiredAttributeTypes, len(*r), len(*r))
	for i, n := range *r {
		R[i] = n
	}
	*r = R
}

/*
Set applies one or more instances of *AttributeType to the receiver.  Uniqueness is enforced by OID.
*/
func (r *ProhibitedAttributeTypes) Set(x ...interface{}) {
	for _, n := range x {
		if z, ok := n.(*AttributeType); ok {
			if z.IsZero() {
				continue
			}
			if _, exists := r.Index(z); !exists {
				*r = append(*r, z)
			}
		}
	}

	R := make(ProhibitedAttributeTypes, len(*r), len(*r))
	for i, n := range *r {
		R[i] = n
	}
	*r = R
}

/*
Set applies one or more instances of *DITStructureRule to the receiver.  Uniqueness is enforced by RuleID.
*/
func (r *SuperiorDITStructureRules) Set(x ...interface{}) {
	for _, n := range x {
		if z, ok := n.(*DITStructureRule); ok {
			if z.IsZero() {
				continue
			}
			if _, exists := r.Index(z); !exists {
				*r = append(*r, z)
			}
		}
	}

	R := make(SuperiorDITStructureRules, len(*r), len(*r))
	for i, n := range *r {
		R[i] = n
	}
	*r = R
}

/*
Set applies one or more instances of string to the receiver.  Uniqueness is enforced through literal comparison.
*/
func (r *Name) Set(x ...string) {
	if len(x) == 0 {
		return
	}

	for n := range x {
		N := stripTags(x[n])
		if isNumericalOID(N) {
			continue
		}

		if _, exists := r.Index(N); !exists {
			*r = append(*r, N)
		}
	}

	R := make(Name, len(*r), len(*r))
	for i, n := range *r {
		R[i] = n
	}
	*r = R
}

/*
Set applies one or more instances of *AttributeType to the receiver.  Uniqueness is enforced by OID.
*/
func (r *Applies) Set(x ...interface{}) {
	for _, n := range x {
		if z, ok := n.(*AttributeType); ok {
			if z.IsZero() {
				continue
			}
			if _, exists := r.Index(z); !exists {
				*r = append(*r, z)
			}
		}
	}

	R := make(Applies, len(*r), len(*r))
	for i, n := range *r {
		R[i] = n
	}
	*r = R
}

package schemax

/*
manifest.go deals with all map[X]X type derivatives.
*/

/*
Set assigns the provided key and oid to the receiver instance.  The arguments must both be strings.  The first argument must be the alias name, while the second must be its ASN.1 dotted object identifier.  Subsequent values are ignored.
*/
func (r *AliasesManifest) Set(als ...interface{}) {
	if r.IsZero() {
		*r = NewAliasesManifest()
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
		if v.string() == oid || equalFold(alias, i) {
			continue
		}
	}

	R := *r
	R[alias] = OID(oid)
	*r = R

	return
}

/*
Resolve returns a registered OID and a boolean value indicative of a successful lookup. A search is conducted using the provided alias key name, with or without the colon:number suffix included.
*/
func (r AliasesManifest) Resolve(key string) (oid OID, ok bool) {

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
			oid = OID(oid.string() + `.` + rest)
		}
	default:
		if len(rest) == 0 {
			return
		}
		oid = OID(rest)
	}
	ok = isNumericalOID(oid.string())

	return
}

/*
Get returns an instance of *NameForm if found within the receiver, else nil.
*/
func (r NameFormsManifest) Get(x string, alm AliasesManifest) (z *NameForm) {
	switch {
	case isNumericalOID(x):
		if z, _ = r[OID(stripTags(x))]; z != nil {
			return
		}
	default:
		n := stripTags(x)

		if !alm.IsZero() {
			if oid, resolved := alm.Resolve(n); resolved {
				return r.Get(string(oid), nil)
			}
		}

		for _, z = range r {
			if z.IsZero() {
				continue
			}
			if z.Name.Equals(n) {
				break
			}
			z = nil
		}
	}

	return
}

/*
Get returns an instance of *AttributeType if found within the receiver, else nil.
*/
func (r AttributeTypesManifest) Get(x string, alm AliasesManifest) (z *AttributeType) {
	switch {
	case isNumericalOID(x):
		if z, _ = r[OID(stripTags(x))]; z != nil {
			return
		}
	default:
		n := stripTags(x)

		if !alm.IsZero() {
			if oid, resolved := alm.Resolve(n); resolved {
				return r.Get(string(oid), nil)
			}
		}

		for _, z = range r {
			if z.IsZero() {
				continue
			}
			if z.Name.Equals(n) {
				break
			}
			z = nil
		}
	}

	return
}

/*
Get returns an instance of *ObjectClass if found within the receiver.
*/
func (r ObjectClassesManifest) Get(x string, alm AliasesManifest) (z *ObjectClass) {
	switch {
	case isNumericalOID(x):
		if z, _ = r[OID(stripTags(x))]; z != nil {
			return
		}
	default:
		n := stripTags(x)

		if !alm.IsZero() {
			if oid, resolved := alm.Resolve(n); resolved {
				return r.Get(string(oid), nil)
			}
		}

		for _, z = range r {
			if z.IsZero() {
				continue
			}
			if z.Name.Equals(n) {
				break
			}
			z = nil
		}
	}

	return
}

/*
Get returns an instance of *LDAPSyntax if found within the receiver.
*/
func (r LDAPSyntaxesManifest) Get(x string, alm AliasesManifest) (z *LDAPSyntax) {
	switch {
	case isNumericalOID(x):
		if z, _ = r[OID(stripTags(x))]; z != nil {
			return
		}
	default:
		if !alm.IsZero() {
			if oid, resolved := alm.Resolve(x); resolved {
				return r.Get(string(oid), nil)
			}
		}

		for _, z = range r {
			if z.IsZero() {
				continue
			}

			// Syntaxes don't really have names like all other
			// definition types, so instead we have to think a
			// bit creatively and use the syntax desc field.
			//
			// TODO - improve this ...
			desc := toLower(replaceAll(string(z.Description), ` `, ``))
			desc = replaceAll(desc, `syntax`, ``)
			x = toLower(replaceAll(x, ` `, ``))
			x = replaceAll(x, `syntax`, ``)
			if contains(desc, x) {
				break
			}
			z = nil
		}
	}

	return
}

/*
Get returns an instance of *MatchingRule if found within the receiver.
*/
func (r MatchingRulesManifest) Get(x string, alm AliasesManifest) (z *MatchingRule) {
	switch {
	case isNumericalOID(x):
		if z, _ = r[OID(stripTags(x))]; z != nil {
			return
		}
	default:
		n := stripTags(x)

		if !alm.IsZero() {
			if oid, resolved := alm.Resolve(n); resolved {
				return r.Get(string(oid), nil)
			}
		}

		var o OID
		for o, z = range r {
			if r[o].Name.Equals(n) {
				break
			}
			z = nil
		}
	}

	return
}

/*
Get returns an instance of *MatchingRuleUse if found within the receiver.
*/
func (r MatchingRuleUsesManifest) Get(x string, alm AliasesManifest) (z *MatchingRuleUse) {
	switch {
	case isNumericalOID(x):
		if z, _ = r[OID(stripTags(x))]; z != nil {
			return
		}
	default:
		n := stripTags(x)

		if !alm.IsZero() {
			if oid, resolved := alm.Resolve(n); resolved {
				return r.Get(string(oid), nil)
			}
		}

		var o OID
		for o, z = range r {
			if r[o].Name.Equals(n) {
				break
			}
			z = nil
		}
	}

	return
}

/*
Get returns an instance of *DITContentRule if found within the receiver.
*/
func (r DITContentRulesManifest) Get(x string, alm AliasesManifest) (z *DITContentRule) {
	switch {
	case isNumericalOID(x):
		if z, _ = r[OID(stripTags(x))]; z != nil {
			return
		}
	default:
		n := stripTags(x)

		if !alm.IsZero() {
			if oid, resolved := alm.Resolve(n); resolved {
				return r.Get(string(oid), nil)
			}
		}

		var o OID
		for o, z = range r {
			if r[o].Name.Equals(n) {
				break
			}
			z = nil
		}
	}

	return
}

/*
Get returns an instance of DITStructureRule if found within the receiver.
*/
func (r DITStructureRulesManifest) Get(x string) (z *DITStructureRule) {
	rid := NewRuleID(x)
	if _, got := r[rid]; got {
		z = r[rid]
		return
	}

	var o RuleID
	for o, z = range r {
		if r[o].Name.Equals(x) {
			break
		}
		z = nil
	}

	return
}

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

/*
Exists returns a boolean value indicative of whether the named label exists within the receiver instance.
*/
func (r LDAPSyntaxesManifest) Exists(x interface{}, alm AliasesManifest) (exists bool) {
	var term string
	switch tv := x.(type) {
	case *LDAPSyntax:
		term = tv.OID.string()
	case string:
		n := stripTags(tv)

		if !alm.IsZero() {
			if oid, resolved := alm.Resolve(n); resolved {
				term = string(oid) //string(r.Get(oid).OID)
			} else {
				term = n
			}
		} else {
			term = n
		}
	default:
		return
	}

	for k, _ := range r {
		if string(k) == term {
			exists = true
			break
		}
	}

	return
}

/*
Exists returns a boolean value indicative of whether the named label exists within the receiver instance.
*/
func (r AttributeTypesManifest) Exists(x interface{}, alm AliasesManifest) (exists bool) {
	var term string
	switch tv := x.(type) {
	case string:
		n := stripTags(tv)

		if !alm.IsZero() {
			if oid, resolved := alm.Resolve(n); resolved {
				term = string(oid) //string(r.Get(oid).OID)
			} else {
				term = n
			}
		} else {
			term = n
		}
	case *AttributeType:
		term = tv.OID.string()
	default:
		return
	}

	// See if its an OID ...
	switch {
	case isNumericalOID(term):
		_, exists = r[OID(term)]
		return
	}

	// Assume its a name
	for _, v := range r {
		if v.Name.Equals(term) {
			return true
		}
	}

	return
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
Set assigned the provided interface{} value(s) to the receiver if valid and unique.
*/
func (r AttributeTypesManifest) Set(ats ...interface{}) {
	for _, at := range ats {
		if z, ok := at.(*AttributeType); ok {
			if z.IsZero() {
				continue
			}
			if r.Exists(z.OID.string(), nil) {
				continue
			}

			z.seq = len(r)
			r[z.OID] = z
		}
	}
}

/*
Set assigned the provided interface{} value(s) to the receiver if valid and unique.
*/
func (r ObjectClassesManifest) Set(ocs ...interface{}) {
	for _, oc := range ocs {
		if z, ok := oc.(*ObjectClass); ok {
			if z.IsZero() {
				continue
			}
			if r.Exists(z.OID.string(), nil) {
				continue
			}

			z.seq = len(r)
			r[z.OID] = z
		}
	}
}

/*
Exists returns a boolean value indicative of whether the named label exists within the receiver instance.
*/
func (r ObjectClassesManifest) Exists(x interface{}, alm AliasesManifest) (exists bool) {
	var term string
	switch tv := x.(type) {
	case string:
		n := stripTags(tv)

		if !alm.IsZero() {
			if oid, resolved := alm.Resolve(n); resolved {
				term = string(oid) //string(r.Get(oid).OID)
			} else {
				term = n
			}
		} else {
			term = n
		}
	case *ObjectClass:
		term = tv.OID.string()
	default:
		return
	}

	switch {
	case isNumericalOID(term):
		_, exists = r[OID(term)]
	default:
		// Assume its a name
		for _, v := range r {
			if exists = v.Name.Equals(term); exists {
				break
			}
		}
	}

	return
}

/*
Set assigned the provided interface{} value(s) to the receiver if valid and unique.
*/
func (r NameFormsManifest) Set(nfs ...interface{}) {
	for _, nf := range nfs {
		if z, ok := nf.(*NameForm); ok {
			if z.IsZero() {
				continue
			}
			if r.Exists(z.OID.string(), nil) {
				continue
			}

			z.seq = len(r)
			r[z.OID] = z
		}
	}
}

/*
Exists returns a boolean value indicative of whether the named label exists within the receiver instance.
*/
func (r NameFormsManifest) Exists(x interface{}, alm AliasesManifest) (exists bool) {
	var term string
	switch tv := x.(type) {
	case string:
		n := stripTags(tv)

		if !alm.IsZero() {
			if oid, resolved := alm.Resolve(n); resolved {
				term = string(oid) //string(r.Get(oid).OID)
			} else {
				term = n
			}
		} else {
			term = n
		}
	case *NameForm:
		term = tv.OID.string()
	default:
		return
	}

	switch {
	case isNumericalOID(term):
		_, exists = r[OID(term)]
	default:
		// Assume its a name
		for _, v := range r {
			if exists = v.Name.Equals(term); exists {
				break
			}
		}
	}

	return
}

/*
Set assigned the provided interface{} value(s) to the receiver if valid and unique.
*/
func (r DITContentRulesManifest) Set(dcrs ...interface{}) {
	for _, dcr := range dcrs {
		if z, ok := dcr.(*DITContentRule); ok {
			if z.IsZero() {
				continue
			}
			if r.Exists(z.OID.string(), nil) {
				continue
			}

			z.seq = len(r)
			r[z.OID] = z
		}
	}
}

/*
Exists returns a boolean value indicative of whether the named label exists within the receiver instance.
*/
func (r DITContentRulesManifest) Exists(x interface{}, alm AliasesManifest) (exists bool) {
	var term string
	switch tv := x.(type) {
	case string:
		n := stripTags(tv)

		if !alm.IsZero() {
			if oid, resolved := alm.Resolve(n); resolved {
				term = string(oid) //string(r.Get(oid).OID)
			} else {
				term = n
			}
		} else {
			term = n
		}
	case *DITContentRule:
		term = tv.OID.string()
	default:
		return
	}

	// See if its an OID ...
	switch {
	case isNumericalOID(term):
		_, exists = r[OID(term)]
	default:
		// Assume its a name
		for _, v := range r {
			if exists = v.Name.Equals(term); exists {
				break
			}
		}
	}

	return
}

/*
Set assigned the provided interface{} value(s) to the receiver if valid and unique.
*/
func (r DITStructureRulesManifest) Set(dsrs ...interface{}) {
	for _, dsr := range dsrs {
		if z, ok := dsr.(*DITStructureRule); ok {
			if z.IsZero() {
				continue
			}
			if r.Exists(z.ID.string(), nil) {
				continue
			}

			z.seq = len(r)
			r[z.ID] = z
		}
	}
}

/*
Set assigned the provided interface{} value(s) to the receiver if valid and unique.
*/
func (r DITStructureRulesManifest) Exists(x interface{}, alm AliasesManifest) (exists bool) {
	var term string
	switch tv := x.(type) {
	case string:
		n := stripTags(tv)

		if !alm.IsZero() {
			if oid, resolved := alm.Resolve(n); resolved {
				term = string(oid) //string(r.Get(oid).OID)
			} else {
				term = n
			}
		} else {
			term = n
		}
	case *DITStructureRule:
		term = tv.ID.string()
	}
	// See if its an OID ...
	switch {
	case isDigit(term):
		_, exists = r[NewRuleID(term)]
	default:
		// Assume its a name
		for _, v := range r {
			if exists = v.Name.Equals(term); exists {
				break
			}
		}
	}

	return
}

/*
Set assigned the provided interface{} value(s) to the receiver if valid and unique.
*/
func (r LDAPSyntaxesManifest) Set(lss ...interface{}) {
	for _, ls := range lss {
		if z, ok := ls.(*LDAPSyntax); ok {
			if z.IsZero() {
				continue
			}
			if r.Exists(z.OID.string(), nil) {
				continue
			}

			z.seq = len(r)
			r[z.OID] = z
		}
	}
}

/*
Exists returns a boolean value indicative of whether the named label exists within the receiver instance.
*/
func (r MatchingRulesManifest) Exists(x interface{}, alm AliasesManifest) (exists bool) {
	var term string
	switch tv := x.(type) {
	case string:
		n := stripTags(tv)

		if !alm.IsZero() {
			if oid, resolved := alm.Resolve(n); resolved {
				term = string(oid) //string(r.Get(oid).OID)
			} else {
				term = n
			}
		} else {
			term = n
		}
	case *MatchingRule:
		term = tv.OID.string()
	}

	// See if its an OID ...
	switch {
	case isNumericalOID(term):
		_, exists = r[OID(term)]
	default:
		for k, v := range r {
			if exists = (string(k) == term || v.Name.Equals(term)); exists {
				break
			}
		}
	}

	return
}

/*
Exists returns a boolean value indicative of whether the named label exists within the receiver instance.
*/
func (r MatchingRuleUsesManifest) Exists(x interface{}, alm AliasesManifest) (exists bool) {
	var term string
	switch tv := x.(type) {
	case string:
		n := stripTags(tv)

		if !alm.IsZero() {
			if oid, resolved := alm.Resolve(n); resolved {
				term = string(oid) //string(r.Get(oid).OID)
			} else {
				term = n
			}
		} else {
			term = n
		}
	case *MatchingRuleUse:
		term = tv.OID.string()
	}

	// See if its an OID ...
	switch {
	case isNumericalOID(term):
		_, exists = r[OID(term)]
	default:
		for k, v := range r {
			if exists = (string(k) == term || v.Name.Equals(term)); exists {
				break
			}
		}
	}

	return
}

/*
Set assigned the provided interface{} value(s) to the receiver if valid and unique.
*/
func (r MatchingRulesManifest) Set(mrs ...interface{}) {
	for _, mr := range mrs {
		if z, ok := mr.(*MatchingRule); ok {
			if z.IsZero() {
				continue
			}
			if r.Exists(z.OID.string(), nil) {
				continue
			}

			z.seq = len(r)
			r[z.OID] = z
		}
	}
}

/*
Set assigned the provided interface{} value(s) to the receiver if valid and unique.
*/
func (r MatchingRuleUsesManifest) Set(mrus ...interface{}) {
	for _, mru := range mrus {
		if z, ok := mru.(*MatchingRuleUse); ok {
			if z.IsZero() {
				continue
			}
			if r.Exists(z.OID.string(), nil) {
				continue
			}

			z.seq = len(r)
			r[z.OID] = z
		}
	}
}

/*
Refresh accepts an AttributeTypeManifest which will be processed and used to create new, or update existing, *MatchingRuleUse instances within the receiver.
*/
func (r MatchingRuleUsesManifest) Refresh(atm AttributeTypesManifest) {
	if atm.IsZero() {
		return
	}

	for _, a := range atm {
		if a.IsZero() {
			continue
		}

		if !a.Equality.IsZero() {
			if !r.Exists(a.Equality.OID, nil) {
				r.Set(&MatchingRuleUse{
					OID:     a.Equality.OID,
					Name:    a.Equality.Name,
					Applies: NewApplies()})
			}
			r[a.Equality.OID].Applies.Set(a)
		}
		if !a.Substring.IsZero() {
			if !r.Exists(a.Substring.OID, nil) {
				r.Set(&MatchingRuleUse{
					OID:     a.Substring.OID,
					Name:    a.Substring.Name,
					Applies: NewApplies()})
			}
			r[a.Substring.OID].Applies.Set(a)
		}
		if !a.Ordering.IsZero() {
			if !r.Exists(a.Ordering.OID, nil) {
				r.Set(&MatchingRuleUse{
					OID:     a.Ordering.OID,
					Name:    a.Ordering.Name,
					Applies: NewApplies()})
			}
			r[a.Ordering.OID].Applies.Set(a)
		}
	}
}

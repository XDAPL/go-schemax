package schemax

/*
NewMatchingRuleUses initializes a new [Collection] instance and
casts it as an [MatchingRuleUses] instance.
*/
func NewMatchingRuleUses() MatchingRuleUses {
	r := MatchingRuleUses(newCollection(``))
	r.cast().SetPushPolicy(r.canPush)

	return r
}

/*
String is a stringer method that returns the string representation
of the receiver instance.
*/
func (r MatchingRuleUse) String() (mu string) {
	if !r.IsZero() {
		mu = r.matchingRuleUse.s
	}

	return
}

/*
IsObsolete returns a Boolean value indicative of definition obsolescence.
*/
func (r MatchingRuleUse) IsObsolete() (o bool) {
	if !r.IsZero() {
		o = r.matchingRuleUse.Obsolete
	}

	return
}

/*
Applies returns an [AttributeTypes] instance containing pointer references
to all [AttributeType] instances to which the receiver applies.
*/
func (r MatchingRuleUse) Applies() (aa AttributeTypes) {
	if !r.IsZero() {
		aa = r.matchingRuleUse.Applies
	}

	return
}

/*
IsIdentifiedAs returns a Boolean value indicative of whether id matches
either the numericOID or descriptor of the receiver instance.  Case is
not significant in the matching process.
*/
func (r MatchingRuleUse) IsIdentifiedAs(id string) (ident bool) {
	if !r.IsZero() {
		ident = id == r.NumericOID() || r.matchingRuleUse.Name.contains(id)
	}

	return
}

func (r *matchingRuleUse) prepareString() (err error) {
	buf := newBuf()
	r.t = newTemplate(`matchingRuleUse`).
		Funcs(funcMap(map[string]any{
			`ExtensionSet`: r.Extensions.tmplFunc,
		}))
	if r.t, err = r.t.Parse(matchingRuleUseTmpl); err == nil {
		if err = r.t.Execute(buf, r); err == nil {
			r.s = buf.String()
		}
	}

	return
}

/*
List returns a map[string][]string instance which represents the current
inventory of matching rule use instances within the receiver.  The keys
are numeric OIDs, while the values are zero (0) or more string slices,
each representing a name by which the definition is known.
*/
func (r MatchingRuleUses) List() (list map[string][]string) {
	list = make(map[string][]string, 0)
	for i := 0; i < r.len(); i++ {
		def := r.index(i)
		list[def.NumericOID()] = def.Names().List()

	}

	return
}

/*
Type returns the string literal "matchingRuleUse".
*/
func (r MatchingRuleUse) Type() string {
	return `matchingRuleUse`
}

/*
Type returns the string literal "matchingRuleUses".
*/
func (r MatchingRuleUses) Type() string {
	return `matchingRuleUses`
}

func (r MatchingRuleUses) prepareStrings() (err error) {
	for i := 0; i < r.cast().Len() && err == nil; i++ {
		mu := r.index(i)
		if mu.IsZero() {
			mu.matchingRuleUse = new(matchingRuleUse)
		}
		err = mu.matchingRuleUse.prepareString()
	}

	return
}

/*
OID returns the string representation of an OID -- which is either
a numeric OID or descriptor -- that refers to the [MatchingRule]
upon which the receiver instance is based.
*/
func (r MatchingRuleUse) OID() (oid string) {
	if !r.IsZero() {
		oid = r.NumericOID() // default
		if r.matchingRuleUse.Name.len() > 0 {
			oid = r.matchingRuleUse.Name.index(0)
		}
	}

	return
}

/*
Extensions returns the [Extensions] instance -- if set -- within
the receiver.
*/
func (r MatchingRuleUse) Extensions() (e Extensions) {
	if !r.IsZero() {
		e = r.matchingRuleUse.Extensions
	}

	return
}

/*
NumericOID returns the string representation of the numeric OID
held by the receiver instance.
*/
func (r MatchingRuleUse) NumericOID() (noid string) {
	if !r.IsZero() {
		noid = r.matchingRuleUse.OID
	}

	return
}

/*
Name returns the string form of the principal name of the receiver instance, if set.
*/
func (r MatchingRuleUse) Name() (id string) {
	if !r.IsZero() {
		id = r.matchingRuleUse.Name.index(0)
	}

	return
}

/*
Names returns the underlying instance of [DefinitionName] from
within the receiver.
*/
func (r MatchingRuleUse) Names() (names DefinitionName) {
	return r.matchingRuleUse.Name
}

/*
Description returns the underlying (optional) descriptive text
assigned to the receiver instance.
*/
func (r MatchingRuleUse) Description() (desc string) {
	if !r.IsZero() {
		desc = r.matchingRuleUse.Desc
	}
	return
}

/*
IsZero returns a Boolean value indicative of nilness of the
receiver instance.
*/
func (r MatchingRuleUse) IsZero() bool {
	return r.matchingRuleUse == nil
}

func (r MatchingRuleUses) canPush(x ...any) (err error) {
	if len(x) == 0 {
		return
	}

	for i := 0; i < len(x) && err == nil; i++ {
		instance := x[i]
		err = errorf("Type assertion for %T has failed", instance)
		if mu, ok := instance.(MatchingRuleUse); ok && !mu.IsZero() {
			err = errorf("%T %s not unique", mu, mu.NumericOID())
			if tst := r.get(mu.NumericOID()); tst.IsZero() {
				err = nil
			}
		}
	}

	return
}

/*
Len returns the current integer length of the receiver instance.
*/
func (r MatchingRuleUses) Len() int {
	return r.len()
}

func (r MatchingRuleUses) len() int {
	return r.cast().Len()
}

/*
String is a stringer method that returns the string representation
of the receiver instance.
*/
func (r MatchingRuleUses) String() string {
	return r.cast().String()
}

/*
IsZero returns a Boolean value indicative of nilness of the
receiver instance.
*/
func (r MatchingRuleUses) IsZero() bool {
	return r.cast().IsZero()
}

/*
Index returns the instance of [MatchingRuleUse] found within the
receiver stack instance at index N.  If no instance is found at
the index specified, a zero [MatchingRuleUse] instance is returned.
*/
func (r MatchingRuleUses) Index(idx int) MatchingRuleUse {
	return r.index(idx)
}

func (r MatchingRuleUses) index(idx int) (mu MatchingRuleUse) {
	slice, found := r.cast().Index(idx)
	if found {
		if _mu, ok := slice.(MatchingRuleUse); ok {
			mu = _mu
		}
	}

	return
}

/*
Push returns an error following an attempt to push an [AttributeType]
instance into the receiver instance.
*/
func (r MatchingRuleUse) Push(at any) error {
	return r.push(at)
}

func (r MatchingRuleUse) push(at any) (err error) {
	err = errorf("Cannot push %T instance into %T", at, r)
	if !r.IsZero() {
		if at != nil {
			r.Applies().cast().Push(at)
			err = nil
		}
	}

	return
}

/*
Push returns an error following an attempt to push a [MatchingRuleUse]
instance into the receiver instance.
*/
func (r MatchingRuleUses) Push(mu any) error {
	return r.push(mu)
}

func (r MatchingRuleUses) push(mu any) (err error) {
	err = errorf("Cannot push %T instance into %T", mu, r)
	if mu != nil {
		r.cast().Push(mu)
		err = nil
	}

	return
}

func (r MatchingRuleUses) contains(id string) bool {
	return !r.get(id).IsZero()
}

func (r MatchingRuleUses) get(id string) (mu MatchingRuleUse) {
	for i := 0; i < r.len() && mu.IsZero(); i++ {
		if _mu := r.index(i); !_mu.IsZero() {
			if _mu.IsIdentifiedAs(id) {
				mu = _mu
			}
		}
	}

	return
}

/*
makeMatchingRuleUse fashions and returns a new MatchingRuleUse instance
based on the contents of the receiver instance.  The returned instance,
assuming a nil error condition, may have its Applies field populated
with "users" (AttributeType instances) of the indicated matchingRule.
*/
func (r MatchingRule) makeMatchingRuleUse() (mu MatchingRuleUse, err error) {
	err = errorf("%T instance is nil, cannot make %T", r, mu)
	if !r.IsZero() {
		_mu := new(matchingRuleUse)
		_mu.Applies = NewAttributeTypeOIDList()
		_mu.Name = r.matchingRule.Name
		_mu.OID = r.matchingRule.OID
		_mu.Desc = r.matchingRule.Desc
		_mu.Extensions = r.matchingRule.Extensions
		mu = MatchingRuleUse{_mu}
		if !mu.IsZero() {
			err = nil
		}
	}

	return
}

/*
updateMatchingRuleUses returns an instance of error following an attempt
to refresh the collection of MatchingRuleUse instances within the
receiver to include input variable "at" wherever appropriate.
*/
func (r Schema) updateMatchingRuleUses(ats AttributeTypes) (err error) {
	err = errorf("%T instance is nil; cannot proceed", ats)
	if ats.len() > 0 {
		err = nil
		for i := 0; i < ats.len() && err == nil; i++ {
			if at := ats.index(i); !at.IsZero() {
				for _, funk := range []func(AttributeType) error{
					r.updateEqualityUses,
					r.updateSubstringUses,
					r.updateOrderingUses,
				} {
					if err = funk(at); err != nil {
						break
					}
				}
			}
		}
	}

	if err == nil {
		err = r.MatchingRuleUses().prepareStrings()
	}

	return
}

// updateEqualityUses is called by updateMatchingRuleUses and will
// extract any equality matchingRule details from the input instance
// of AttributeType and store information about this association in
// the receiver's MU stack field.  An error is returned should any
// part of this process fail.
func (r Schema) updateEqualityUses(at AttributeType) (err error) {
	if eqty := at.Equality(); !eqty.IsZero() {
		mu := r.MatchingRuleUses().get(eqty.NumericOID())

		// If the MatchingRuleUse instance does not exist,
		// create it now.
		if mu.IsZero() {
			err = errorf("equality matchingRule %s not found", eqty)
			if _eqy := r.MatchingRules().get(eqty.NumericOID()); !_eqy.IsZero() {
				if mu, err = _eqy.makeMatchingRuleUse(); err == nil {
					r.MatchingRuleUses().push(mu)
				}
			}
		}

		// Append the AttributeType instance to the
		// MatchingRuleUse.Applies stack, assuming
		// no errors occur.
		if err == nil && !mu.IsZero() {
			mu.push(at)
		}
	}

	return
}

// updateSubstringUses is called by updateMatchingRuleUses and will
// extract any substring matchingRule details from the input instance
// of AttributeType and store information about this association in
// the receiver's MU stack field.  An error is returned should any
// part of this process fail.
func (r Schema) updateSubstringUses(at AttributeType) (err error) {
	if substr := at.Substring(); !substr.IsZero() {
		mu := r.MatchingRuleUses().get(substr.NumericOID())

		// If the MatchingRuleUse instance does not exist,
		// create it now.
		if mu.IsZero() {
			err = errorf("substring matchingRule %s not found", substr)
			if _sub := r.MatchingRules().get(substr.NumericOID()); !_sub.IsZero() {
				if mu, err = _sub.makeMatchingRuleUse(); err == nil {
					r.MatchingRuleUses().push(mu)
				}
			}
		}

		// Append the AttributeType instance to the
		// MatchingRuleUse.Applies stack, assuming
		// no errors occur.
		if err == nil && !mu.IsZero() {
			mu.push(at)
		}
	}

	return
}

// updateOrderingUses is called by updateMatchingRuleUses and will
// extract any ordering matchingRule details from the input instance
// of AttributeType and store information about this association in
// the receiver's MU stack field.  An error is returned should any
// part of this process fail.
func (r Schema) updateOrderingUses(at AttributeType) (err error) {
	if order := at.Ordering(); !order.IsZero() {
		mu := r.MatchingRuleUses().get(order.NumericOID())

		// If the MatchingRuleUse instance does not exist,
		// create it now.
		if mu.IsZero() {
			err = errorf("ordering matchingRule %s not found", order)
			if _ord := r.MatchingRules().get(order.NumericOID()); !_ord.IsZero() {
				if mu, err = _ord.makeMatchingRuleUse(); err == nil {
					r.MatchingRuleUses().push(mu)
				}
			}
		}

		// Append the AttributeType instance to the
		// MatchingRuleUse.Applies stack, assuming
		// no errors occur.
		if err == nil && !mu.IsZero() {
			mu.push(at)
		}
	}

	return
}

func (r MatchingRuleUse) setOID(_ string) {}
func (r MatchingRuleUse) macro() []string { return []string{} }

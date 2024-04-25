package schemax

/*
NewDITContentRules initializes a new [Collection] instance and
casts it as a [DITContentRules] instance.
*/
func NewDITContentRules() DITContentRules {
	r := DITContentRules(newCollection(``))
	r.cast().SetPushPolicy(r.canPush)

	return r
}

/*
Description returns the underlying (optional) descriptive text
assigned to the receiver instance.
*/
func (r DITContentRule) Description() (desc string) {
	if !r.IsZero() {
		desc = r.dITContentRule.Desc
	}

	return
}

/*
DITContentRules returns the [DITContentRules] instance from within
the receiver instance.
*/
func (r Schema) DITContentRules() (dcs DITContentRules) {
	slice, _ := r.cast().Index(dITContentRulesIndex)
	dcs, _ = slice.(DITContentRules)
	return
}

func newDITContentRule() *dITContentRule {
	return &dITContentRule{
		Name:       NewName(),
		Aux:        NewObjectClassOIDList(),
		Must:       NewAttributeTypeOIDList(),
		May:        NewAttributeTypeOIDList(),
		Not:        NewAttributeTypeOIDList(),
		Extensions: NewExtensions(),
	}
}

/*
SetSchema assigns an instance of [Schema] to the receiver instance.  This allows
internal verification of certain actions without the need for user input of
an instance of [Schema] manually at each juncture.

Note that the underlying [Schema] instance is automatically set when creating
instances of this type by way of parsing.

This is a fluent method.
*/
func (r *DITContentRule) SetSchema(schema Schema) *DITContentRule {
	if r.IsZero() {
		r.dITContentRule = newDITContentRule()
	}

	r.dITContentRule.schema = schema

	return r
}

func (r DITContentRule) schema() (s Schema) {
	if !r.IsZero() {
		s = r.dITContentRule.schema
	}

	return
}

/*
Inventory returns an instance of [Inventory] which represents the current
inventory of [DITContentRule] instances within the receiver.

The keys are numeric OIDs, while the values are zero (0) or more string
slices, each representing a name by which the definition is known.
*/
func (r DITContentRules) Inventory() (inv Inventory) {
	inv = make(Inventory, 0)
	for i := 0; i < r.len(); i++ {
		def := r.index(i)
		inv[def.NumericOID()] = def.Names().List()
	}

	return
}

/*
Maps returns slices of [DefinitionMap] instances.
*/
func (r DITContentRules) Maps() (defs DefinitionMaps) {
	defs = make(DefinitionMaps, r.Len())
	for i := 0; i < r.Len(); i++ {
		defs[i] = r.Index(i).Map()
	}

	return
}

/*
Map marshals the receiver instance into an instance of
[DefinitionMap].
*/
func (r DITContentRule) Map() (def DefinitionMap) {
	if r.IsZero() {
		return
	}

	var auxs []string
	for i := 0; i < r.Aux().Len(); i++ {
		m := r.Aux().Index(i)
		auxs = append(auxs, m.OID())
	}

	var musts []string
	for i := 0; i < r.Must().Len(); i++ {
		m := r.Must().Index(i)
		musts = append(musts, m.OID())
	}

	var mays []string
	for i := 0; i < r.May().Len(); i++ {
		m := r.May().Index(i)
		mays = append(mays, m.OID())
	}

	var nots []string
	for i := 0; i < r.Not().Len(); i++ {
		m := r.Not().Index(i)
		nots = append(nots, m.OID())
	}

	def = make(DefinitionMap, 0)
	def[`NUMERICOID`] = []string{r.NumericOID()}
	def[`NAME`] = r.Names().List()
	def[`DESC`] = []string{r.Description()}
	def[`OBSOLETE`] = []string{bool2str(r.IsObsolete())}
	def[`AUX`] = auxs
	def[`MUST`] = musts
	def[`MAY`] = mays
	def[`NOT`] = nots
	def[`TYPE`] = []string{r.Type()}
	def[`RAW`] = []string{r.String()}

	// copy our extensions from receiver r
	// into destination def.
	def = mapTransferExtensions(r, def)

	// Clean up any empty fields
	def.clean()

	return
}

/*
Type returns the string literal "dITContentRule".
*/
func (r DITContentRule) Type() string {
	return `dITContentRule`
}

/*
Type returns the string literal "dITContentRules".
*/
func (r DITContentRules) Type() string {
	return `dITContentRules`
}

/*
Len returns the current integer length of the receiver instance.
*/
func (r DITContentRules) Len() int {
	return r.len()
}

func (r DITContentRules) len() int {
	return r.cast().Len()
}

/*
String is a stringer method that returns the string representation
of the receiver instance.
*/
func (r DITContentRules) String() string {
	return r.cast().String()
}

/*
IsZero returns a Boolean value indicative of nilness of the
receiver instance.
*/
func (r DITContentRules) IsZero() bool {
	return r.cast().IsZero()
}

/*
Index returns the instance of [DITContentRule] found within the
receiver stack instance at index N.  If no instance is found at
the index specified, a zero [DITContentRule] instance is returned.
*/
func (r DITContentRules) Index(idx int) DITContentRule {
	return r.index(idx)
}

func (r DITContentRules) index(idx int) (dc DITContentRule) {
	slice, found := r.cast().Index(idx)
	if found {
		if _dc, ok := slice.(DITContentRule); ok {
			dc = _dc
		}
	}

	return
}

func (r DITContentRule) macro() (m []string) {
	if !r.IsZero() {
		m = r.dITContentRule.Macro
	}

	return
}

func (r DITContentRules) oIDsStringer(_ ...any) (present string) {
	var _present []string
	for i := 0; i < r.len(); i++ {
		_present = append(_present, r.index(i).OID())
	}

	switch len(_present) {
	case 0:
		break
	case 1:
		present = _present[0]
	default:
		padchar := string(rune(32))
		if !r.cast().IsPadded() {
			padchar = ``
		}

		joined := join(_present, padchar+`$`+padchar)
		present = "(" + padchar + joined + padchar + ")"
	}

	return
}

/*
IsIdentifiedAs returns a Boolean value indicative of whether id matches
either the numericOID or descriptor of the receiver instance.  Case is
not significant in the matching process.
*/
func (r DITContentRule) IsIdentifiedAs(id string) (ident bool) {
	if !r.IsZero() {
		ident = id == r.NumericOID() || r.dITContentRule.Name.contains(id)
	}

	return
}

/*
IsObsolete returns a Boolean value indicative of definition obsolescence.
*/
func (r DITContentRule) IsObsolete() (o bool) {
	if !r.IsZero() {
		o = r.dITContentRule.Obsolete
	}

	return
}

/*
OID returns the string representation of an OID -- which is either a
numeric OID or descriptor -- that is held by the receiver instance.
*/
func (r DITContentRule) OID() (oid string) {
	if !r.IsZero() {
		oid = r.NumericOID() // default
		if r.dITContentRule.Name.len() > 0 {
			oid = r.dITContentRule.Name.index(0)
		}
	}

	return
}

/*
Aux returns an [ObjectClasses] containing zero (0) or more auxiliary
[ObjectClass] definitions for use within entries governed by this rule.
*/
func (r DITContentRule) Aux() (aux ObjectClasses) {
	if !r.IsZero() {
		aux = r.dITContentRule.Aux
	}

	return
}

/*
Must returns an [AttributeTypes] containing zero (0) or more required
[AttributeType] definitions for use within entries governed by this rule.
*/
func (r DITContentRule) Must() (must AttributeTypes) {
	if !r.IsZero() {
		must = r.dITContentRule.Must
	}

	return
}

/*
May returns an AttributeTypes containing zero (0) or more allowed
[AttributeType] definitions for use within entries governed by this rule.
*/
func (r DITContentRule) May() (may AttributeTypes) {
	if !r.IsZero() {
		may = r.dITContentRule.May
	}

	return
}

/*
Not returns an AttributeTypes containing zero (0) or more [AttributeType]
definitions disallowed for use within entries governed by this rule.
*/
func (r DITContentRule) Not() (not AttributeTypes) {
	if !r.IsZero() {
		not = r.dITContentRule.Not
	}

	return
}

/*
SetStringer allows the assignment of an individual "stringer" function
or method to the receiver instance.

A non-nil value will be executed for every call of the String method
for the receiver instance.

Should the input stringer value be nil, the [text/template.Template]
value will be used automatically going forward.

This is a fluent method.
*/
func (r *DITContentRule) SetStringer(stringer func() string) DITContentRule {
	if r.IsZero() {
		r.dITContentRule = newDITContentRule()
	}

	r.dITContentRule.stringer = stringer

	return *r
}

/*
String is a stringer method that returns the string representation
of the receiver instance.
*/
func (r DITContentRule) String() (dcr string) {
	if !r.IsZero() {
		if r.dITContentRule.stringer != nil {
			dcr = r.dITContentRule.stringer()
		} else {
			dcr = r.dITContentRule.s
		}
	}

	return
}

func (r *dITContentRule) prepareString() (err error) {
	buf := newBuf()
	r.t = newTemplate(`dITContentRule`).
		Funcs(funcMap(map[string]any{
			`ExtensionSet`:  r.Extensions.tmplFunc,
			`StructuralOID`: r.OID.NumericOID,
			`AuxLen`:        r.Aux.len,
			`MayLen`:        r.May.len,
			`MustLen`:       r.Must.len,
			`NotLen`:        r.Not.len,
			`IsObsolete`:    func() bool { return r.Obsolete },
		}))

	if r.t, err = r.t.Parse(dITContentRuleTmpl); err == nil {
		if err = r.t.Execute(buf, struct {
			Definition *dITContentRule
			HIndent    string
		}{
			Definition: r,
			HIndent:    hindent(),
		}); err == nil {
			r.s = buf.String()
		}
	}

	return
}

/*
Name returns the string form of the principal name of the receiver instance, if set.
*/
func (r DITContentRule) Name() (id string) {
	if !r.IsZero() {
		id = r.dITContentRule.Name.index(0)
	}

	return
}

/*
Names returns the underlying instance of [Name] from within
the receiver.
*/
func (r DITContentRule) Names() (names Name) {
	return r.dITContentRule.Name
}

/*
Extensions returns the [Extensions] instance -- if set -- within
the receiver.
*/
func (r DITContentRule) Extensions() (e Extensions) {
	if !r.IsZero() {
		e = r.dITContentRule.Extensions
	}

	return
}

/*
NumericOID returns the string representation of the numeric OID
held by the receiver instance.
*/
func (r DITContentRule) NumericOID() (noid string) {
	if !r.IsZero() {
		if soid := r.dITContentRule.OID; !soid.IsZero() {
			noid = soid.NumericOID()
		}
	}

	return
}

func (r DITContentRule) setOID(_ string) {}

/*
Push returns an error following an attempt to push a [DITContentRule]
into the receiver stack instance.
*/
func (r DITContentRules) Push(dc any) error {
	return r.push(dc)
}

func (r DITContentRules) push(dc any) (err error) {
	err = errorf("%T instance is nil; cannot append to %T", dc, r)
	if dc != nil {
		r.cast().Push(dc)
		err = nil
	}

	return
}

/*
IsZero returns a Boolean value indicative of nilness of the
receiver instance.
*/
func (r DITContentRule) IsZero() bool {
	return r.dITContentRule == nil
}

/*
Contains calls [DITStructureRules.Get] to return a Boolean value indicative of
a successful, non-zero retrieval of an [DITStructureRule] instance -- matching
the provided id -- from within the receiver stack instance.
*/
func (r DITContentRules) Contains(id string) bool {
	return r.contains(id)
}

func (r DITContentRules) contains(id string) bool {
	return !r.get(id).IsZero()
}

/*
Get returns an instance of [DITContentRule] based upon a search for id within
the receiver stack instance.

The return instance, if not nil, was retrieved based upon a textual match of
the principal identifier of an [DITContentRule] and the provided id.

The return instance is nil if no match was made.

Case is not significant in the matching process.
*/
func (r DITContentRules) Get(id string) DITContentRule {
	return r.get(id)
}

func (r DITContentRules) get(id string) (dc DITContentRule) {
	for i := 0; i < r.len() && dc.IsZero(); i++ {
		if _dc := r.index(i); !_dc.IsZero() {
			if _dc.NumericOID() == id {
				dc = _dc
			} else if _dc.dITContentRule.Name.contains(id) {
				dc = _dc
			}
		}
	}

	return
}

func (r *dITContentRule) check() (err error) {
	if r == nil {
		err = errorf("%T is nil")
		return
	}

	if len(r.OID.NumericOID()) == 0 {
		err = errorf("%T lacks an OID", r)
	}

	return
}

// stackage closure func - do not exec directly.
func (r DITContentRules) canPush(x ...any) (err error) {
	if len(x) == 0 {
		return
	}

	for i := 0; i < len(x) && err == nil; i++ {
		instance := x[i]
		err = errorf("Type assertion for %T has failed", instance)
		if dc, ok := instance.(DITContentRule); ok && !dc.IsZero() {
			err = errorf("%T %s not unique", dc, dc.NumericOID())
			if tst := r.get(dc.NumericOID()); tst.IsZero() {
				err = nil
			}
		}
	}

	return
}

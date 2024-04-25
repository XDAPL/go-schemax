package schemax

/*
NewDITStructureRules initializes a new [Collection] instance and casts
it as an [DITStructureRules] instance.
*/
func NewDITStructureRules() DITStructureRules {
	r := DITStructureRules(newCollection(``))
	r.cast().SetPushPolicy(r.canPush)

	return r
}

/*
NewDITStructureRuleIDList initializes a new [RuleIDList] instance and casts it
as a [DITStructureRules] instance.

This is mainly used to define a series of superior [DITStructureRule] instances
specified by a subordinate instance of [DITStructureRule].
*/
func NewDITStructureRuleIDList() DITStructureRules {
	r := DITStructureRules(newRuleIDList(``))
	r.cast().
		SetPushPolicy(r.canPush).
		SetPresentationPolicy(r.iDsStringer)

	return r
}

/*
DITStructureRules returns the [DITStructureRules] instance from
within the receiver instance.
*/
func (r Schema) DITStructureRules() (dss DITStructureRules) {
	slice, _ := r.cast().Index(dITStructureRulesIndex)
	dss, _ = slice.(DITStructureRules)
	return
}

/*
SetSchema assigns an instance of [Schema] to the receiver instance.  This allows
internal verification of certain actions without the need for user input of
an instance of [Schema] manually at each juncture.

Note that the underlying [Schema] instance is automatically set when creating
instances of this type by way of parsing.

This is a fluent method.
*/
func (r *DITStructureRule) SetSchema(schema Schema) *DITStructureRule {
	if r.IsZero() {
		r.dITStructureRule = newDITStructureRule()
	}

	r.dITStructureRule.schema = schema

	return r
}

func (r DITStructureRule) schema() (s Schema) {
	if !r.IsZero() {
		s = r.dITStructureRule.schema
	}

	return
}

/*
NumericOID returns an empty string, as [DITStructureRule] definitions
do not bear numeric OIDs.  This method exists only to satisfy Go
interface requirements.
*/
func (r DITStructureRule) NumericOID() string { return `` }

/*
IsObsolete returns a Boolean value indicative of definition obsolescence.
*/
func (r DITStructureRule) IsObsolete() (o bool) {
	if !r.IsZero() {
		o = r.dITStructureRule.Obsolete
	}

	return
}

/*
Type returns the string literal "dITStructureRule".
*/
func (r DITStructureRule) Type() string {
	return `dITStructureRule`
}

/*
Type returns the string literal "dITStructureRules".
*/
func (r DITStructureRules) Type() string {
	return `dITStructureRules`
}

/*
SuperRules returns a [DITStructureRules] containing zero (0) or more
superior [DITStructureRule] instances from which the receiver extends.
*/
func (r DITStructureRule) SuperRules() (sup DITStructureRules) {
	if !r.IsZero() {
		sup = r.dITStructureRule.SuperRules
	}

	return
}

/*
ID returns the string representation of the underlying rule ID held by
the receiver instance.
*/
func (r DITStructureRule) ID() (id string) {
	if !r.IsZero() {
		_id := sprintf("%d", r.RuleID()) // default
		if id = r.Name(); len(id) == 0 {
			id = _id
		}
	}

	return
}

/*
IsIdentifiedAs returns a Boolean value indicative of whether id matches
either the numericOID or descriptor of the receiver instance.  Case is
not significant in the matching process.
*/
func (r DITStructureRule) IsIdentifiedAs(id string) (ident bool) {
	if !r.IsZero() {
		ident = id == sprintf("%d", r.RuleID()) ||
			r.dITStructureRule.Name.contains(id)
	}

	return
}

/*
Form returns the underlying instance of [NameForm] set within the
receiver. If unset, a zero instance is returned.
*/
func (r DITStructureRule) Form() (nf NameForm) {
	if !r.IsZero() {
		nf = r.dITStructureRule.Form
	}

	return
}

/*
Extensions returns the [Extensions] instance -- if set -- within
the receiver.
*/
func (r DITStructureRule) Extensions() (e Extensions) {
	if !r.IsZero() {
		e = r.dITStructureRule.Extensions
	}

	return
}

/*
RuleID returns the unsigned integer identifier held by the
receiver instance.
*/
func (r DITStructureRule) RuleID() (id uint) {
	if !r.IsZero() {
		id = r.dITStructureRule.ID
	}

	return
}

/*
Name returns the string form of the principal name of the receiver instance, if set.
*/
func (r DITStructureRule) Name() (id string) {
	if !r.IsZero() {
		id = r.dITStructureRule.Name.index(0)
	}

	return
}

/*
Names returns the underlying instance of [Name] from
within the receiver.
*/
func (r DITStructureRule) Names() (names Name) {
	return r.dITStructureRule.Name
}

/*
Description returns the underlying (optional) descriptive text
assigned to the receiver instance.
*/
func (r DITStructureRule) Description() (desc string) {
	if !r.IsZero() {
		desc = r.dITStructureRule.Desc
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
func (r *DITStructureRule) SetStringer(stringer func() string) DITStructureRule {
	if r.IsZero() {
		r.dITStructureRule = newDITStructureRule()
	}

	r.dITStructureRule.stringer = stringer

	return *r
}

func newDITStructureRule() *dITStructureRule {
	return &dITStructureRule{
		Name:       NewName(),
		SuperRules: NewDITStructureRuleIDList(),
		Extensions: NewExtensions(),
	}
}

/*
String is a stringer method that returns the string representation
of the receiver instance.
*/
func (r DITStructureRule) String() (dsr string) {
	if !r.IsZero() {
		if r.dITStructureRule.stringer != nil {
			dsr = r.dITStructureRule.stringer()
		} else {
			dsr = r.dITStructureRule.s
		}
	}

	return
}

/*
Inventory returns an instance of [Inventory] which represents the current
inventory of [DITStructureRule] instances within the receiver.

The keys are numeric OIDs, while the values are zero (0) or more string
slices, each representing a name by which the definition is known.
*/
func (r DITStructureRules) Inventory() (inv Inventory) {
	inv = make(Inventory, 0)
	for i := 0; i < r.len(); i++ {
		def := r.index(i)
		inv[itoa(int(def.RuleID()))] = def.Names().List()
	}

	return
}

func (r DITStructureRule) setOID(_ string) {}
func (r DITStructureRule) macro() []string { return []string{} }

// stackage closure func - do not exec directly (use String method)
func (r DITStructureRules) iDsStringer(_ ...any) (present string) {
	var _present []string
	for i := 0; i < r.len(); i++ {
		_present = append(_present, itoa(int(r.index(i).RuleID())))
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

		joined := join(_present, sprintf("%s %s",
			padchar, padchar))

		present = sprintf("(%s%s%s)",
			padchar, joined, padchar)
	}

	return
}

// stackage closure func - do not exec directly.
func (r DITStructureRules) canPush(x ...any) (err error) {
	if len(x) == 0 {
		return
	}

	for i := 0; i < len(x) && err == nil; i++ {
		instance := x[i]
		if ds, ok := instance.(DITStructureRule); !ok || ds.IsZero() {
			err = errorf("Type assertion for %T has failed", instance)
		} else if tst := r.get(ds.RuleID()); !tst.IsZero() {
			err = errorf("%T %d not unique", ds, ds.RuleID())
		}
	}

	return
}

/*
Len returns the current integer length of the receiver instance.
*/
func (r DITStructureRules) Len() int {
	return r.len()
}

func (r DITStructureRules) len() int {
	return r.cast().Len()
}

/*
String is a stringer method that returns the string representation
of the receiver instance.
*/
func (r DITStructureRules) String() string {
	return r.cast().String()
}

/*
IsZero returns a Boolean value indicative of nilness of the
receiver instance.
*/
func (r DITStructureRules) IsZero() bool {
	return r.cast().IsZero()
}

/*
Index returns the instance of [DITStructureRule] found within the
receiver stack instance at index N.  If no instance is found at
the index specified, a zero [DITStructureRule] instance is returned.
*/
func (r DITStructureRules) Index(idx int) DITStructureRule {
	return r.index(idx)
}

func (r DITStructureRules) index(idx int) (ds DITStructureRule) {
	if slice, found := r.cast().Index(idx); found {
		if _ds, ok := slice.(DITStructureRule); ok {
			ds = _ds
		}
	}

	return
}

/*
Push returns an error following an attempt to push a [DITStructureRule]
into the receiver stack instance.
*/
func (r DITStructureRules) Push(ds any) error {
	return r.push(ds)
}

func (r DITStructureRules) push(ds any) (err error) {
	if ds == nil {
		err = errorf("%T instance is nil; cannot append to %T", ds, r)
		return
	}

	r.cast().Push(ds)

	return
}

/*
Contains calls [Contains.Get] to return a Boolean value indicative of
a successful, non-zero retrieval of an [DITStructureRules] instance --
matching the provided id -- from within the receiver stack instance.
*/
func (r DITStructureRules) Contains(id string) bool {
	return r.contains(id)
}

func (r DITStructureRules) contains(id string) bool {
	return !r.get(id).IsZero()
}

/*
Get returns an instance of [DITStructureRule] based upon a search for id within
the receiver stack instance.

The return instance, if not nil, was retrieved based upon a textual match of
the principal identifier of an [DITStructureRule] and the provided id.

The return instance is nil if no match was made.

Case is not significant in the matching process.
*/
func (r DITStructureRules) Get(id any) DITStructureRule {
	return r.get(id)
}

func (r DITStructureRules) get(id any) (ds DITStructureRule) {
	L := r.len()
	if L == 0 {
		return
	}
	var n uint
	var name string

	switch tv := id.(type) {
	case int:
		if tv < 0 {
			return
		}
		n = uint(tv)
	case uint:
		n = tv
	case string:
		// string may be a string
		// uint, or a name.
		if _n, err := atoi(tv); err != nil {
			name = tv
		} else {
			return r.get(_n)
		}
	}

	for i := 0; i < L && ds.IsZero(); i++ {
		if _ds := r.index(i); !_ds.IsZero() {
			if _ds.ID() == name && len(name) > 0 {
				ds = _ds
			} else if _ds.dITStructureRule.ID == n {
				ds = _ds
			}
		}
	}

	return
}

/*
Maps returns slices of [DefinitionMap] instances.
*/
func (r DITStructureRules) Maps() (defs DefinitionMaps) {
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
func (r DITStructureRule) Map() (def DefinitionMap) {
	if r.IsZero() {
		return
	}

	var sups []string
	for i := 0; i < r.SuperRules().Len(); i++ {
		m := r.SuperRules().Index(i)
		sups = append(sups, itoa(int(m.RuleID())))
	}

	if r.Form().IsZero() {
		return
	}

	def = make(DefinitionMap, 0)
	def[`RULEID`] = []string{itoa(int(r.RuleID()))}
	def[`NAME`] = r.Names().List()
	def[`DESC`] = []string{r.Description()}
	def[`OBSOLETE`] = []string{bool2str(r.IsObsolete())}
	def[`FORM`] = []string{r.Form().OID()}
	def[`SUP`] = sups
	def[`TYPE`] = []string{r.Type()}
	def[`RAW`] = []string{r.String()}

	// copy our extensions from receiver r
	// into destination def.
	def = mapTransferExtensions(r, def)

	// Clean up any empty fields
	def.clean()

	return
}

func (r *dITStructureRule) prepareString() (err error) {
	buf := newBuf()
	r.t = newTemplate(`dITStructureRule`).
		Funcs(funcMap(map[string]any{
			`ExtensionSet`: r.Extensions.tmplFunc,
			`SuperLen`:     r.SuperRules.len,
			`IsObsolete`:   func() bool { return r.Obsolete },
		}))

	if r.t, err = r.t.Parse(dITStructureRuleTmpl); err == nil {
		if err = r.t.Execute(buf, struct {
			Definition *dITStructureRule
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

func (r *dITStructureRule) check() (err error) {
	if r == nil {
		err = errorf("%T instance is nil", r)
		return
	}

	if r.Form.IsZero() {
		err = errorf("%T missing %T", r, r.Form)
	}

	return
}

/*
IsZero returns a Boolean value indicative of nilness of the
receiver instance.
*/
func (r DITStructureRule) IsZero() bool {
	return r.dITStructureRule == nil
}

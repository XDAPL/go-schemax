package schemax

/*
NewNameForms initializes a new [NameForms] instance.
*/
func NewNameForms() NameForms {
	r := NameForms(newCollection(``))
	r.cast().SetPushPolicy(r.canPush)

	return r
}

/*
Nameforms returns the [NameForms] instance from within the receiver
instance.
*/
func (r Schema) NameForms() (nfs NameForms) {
	slice, _ := r.cast().Index(nameFormsIndex)
	nfs, _ = slice.(NameForms)
	return
}

/*
Maps returns slices of [DefinitionMap] instances.
*/
func (r NameForms) Maps() (defs DefinitionMaps) {
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
func (r NameForm) Map() (def DefinitionMap) {
	if r.IsZero() {
		return
	}

	var musts []string
	for i := 0; i < r.Must().Len(); i++ {
		m := r.Must().Index(i)
		musts = append(musts, m.OID())
	}

	// if either of these evaluate as true,
	// we shall not pass.
	if len(musts) == 0 {
		return
	} else if r.OC().IsZero() {
		return
	}

	var mays []string
	for i := 0; i < r.May().Len(); i++ {
		m := r.May().Index(i)
		mays = append(mays, m.OID())
	}

	def = make(DefinitionMap, 0)
	def[`NUMERICOID`] = []string{r.NumericOID()}
	def[`NAME`] = r.Names().List()
	def[`DESC`] = []string{r.Description()}
	def[`OBSOLETE`] = []string{bool2str(r.IsObsolete())}
	def[`OC`] = []string{r.OC().OID()}
	def[`MUST`] = musts
	def[`MAY`] = mays
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
NewNameForm initializes and returns an unpopulated new instance
of [NameForm].  Use of this function is only needed when building
an instance of [NameForm] manually using the various Set methods,
and without involvement of the parser.
*/
func NewNameForm() NameForm {
	return NameForm{
		&nameForm{
			Name:       NewName(),
			Must:       NewAttributeTypeOIDList(),
			May:        NewAttributeTypeOIDList(),
			Extensions: NewExtensions(),
		},
	}
}

/*
IsIdentifiedAs returns a Boolean value indicative of whether id matches
either the numericOID or descriptor of the receiver instance.  Case is
not significant in the matching process.
*/
func (r NameForm) IsIdentifiedAs(id string) (ident bool) {
	if !r.IsZero() {
		ident = id == r.NumericOID() || r.nameForm.Name.contains(id)
	}

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
func (r *NameForm) SetSchema(schema Schema) *NameForm {
	if r.IsZero() {
		r.nameForm = newNameForm()
	}

	r.nameForm.schema = schema

	return r
}

func (r NameForm) schema() (s Schema) {
	if !r.IsZero() {
		s = r.nameForm.schema
	}

	return
}

/*
SetName allows the manual assignment of one (1) or more RFC 4512-compliant
descriptor values by which the receiver instance is to be known.  This will
append to -- not replace -- any preexisting names.

This is a fluent method.
*/
/*
func (r *NameForm) SetName(name ...string) NameForm {
	if r.nameForm == nil {
		r.nameForm = new(nameForm)
	}

	if !r.IsZero() {
		for _, n := range name {
			r.nameForm.Name.cast().Push(n)
		}
	}

	return *r
}
*/

/*
SetNumericOID allows the manual assignment of a numeric OID to the
receiver instance if the following are all true:

  - The input id value is a syntactically valid numeric OID
  - The receiver does not already possess a numeric OID

This is a fluent method.
*/
func (r *NameForm) SetNumericOID(id string) NameForm {
	if r.nameForm == nil {
		r.nameForm = new(nameForm)
	}

	if isNumericOID(id) {
		if len(r.nameForm.OID) == 0 {
			r.nameForm.OID = id
		}
	}

	return *r
}

/*
OID returns the string representation of the OID -- whether numeric or
descriptor -- held by the receiver instance.  Note that a principal name
is requested first and, if not found, the numeric OID is returned as a
fallback.  The return value should NEVER be zero for a valid receiver.
*/
func (r NameForm) OID() (oid string) {
	if !r.IsZero() {
		oid = r.NumericOID() // default
		if r.nameForm.Name.len() > 0 {
			oid = r.nameForm.Name.index(0)
		}
	}

	return
}

/*
SetObsolete will declare the receiver instance as OBSOLETE.  Calls to
this method will not operate in a toggling manner in that there is no
way to "unset" a state of obsolescence.

This is a fluent method.
*/
func (r *NameForm) SetObsolete() NameForm {
	if r.nameForm == nil {
		r.nameForm = new(nameForm)
	}

	if !r.IsZero() {
		if !r.nameForm.Obsolete {
			r.nameForm.Obsolete = true
		}
	}

	return *r
}

/*
IsObsolete returns a Boolean value indicative of definition obsolescence.
*/
func (r NameForm) IsObsolete() (o bool) {
	if !r.IsZero() {
		o = r.nameForm.Obsolete
	}

	return
}

/*
SetOC assigns the referenced Structural [ObjectClass] -- which can be input in literal
ObjectClass form, or through a numeric OID to be searched in the underlying [Schema] --

This is a fluent method.
*/
//func (r NameForm) SetOC(x any) NameForm {
//}

/*
Extensions returns the [Extensions] instance -- if set -- within
the receiver.

Additions to the [Extensions] instance returned may be effected using
the [Extensions.Set] method. Existing [Extensions] cannot be altered.
*/
func (r NameForm) Extensions() (e Extensions) {
	if !r.IsZero() {
		e = r.nameForm.Extensions
	}

	return
}

/*
May returns an [AttributeTypes] containing zero (0) or more allowed
[AttributeType] definitions for use with this class.

Additions to the [AttributeTypes] instance returned may be effected
using the [AttributeTypes.Push] method.  Existing [AttributeType]
instances cannot be removed.
*/
func (r NameForm) May() (may AttributeTypes) {
	if !r.IsZero() {
		may = r.nameForm.May
	}

	return
}

/*
Must returns an [AttributeTypes] containing zero (0) or more required
[AttributeType] definitions for use with this class.

Additions to the [AttributeTypes] instance returned may be effected
using the [AttributeTypes.Push] method.  Existing [AttributeType]
instances cannot be removed.
*/
func (r NameForm) Must() (must AttributeTypes) {
	if !r.IsZero() {
		must = r.nameForm.Must
	}

	return
}

/*
NumericOID returns the string representation of the numeric OID
held by the receiver instance.
*/
func (r NameForm) NumericOID() (noid string) {
	if !r.IsZero() {
		noid = r.nameForm.OID
	}

	return
}

/*
Name returns the string form of the principal name of the receiver instance, if set.
*/
func (r NameForm) Name() (id string) {
	if !r.IsZero() {
		id = r.nameForm.Name.index(0)
	}

	return
}

/*
Names returns the underlying instance of [Name] from within the
receiver instance.
*/
func (r NameForm) Names() (names Name) {
	return r.nameForm.Name
}

func newNameForm() *nameForm {
	return &nameForm{
		Name:       NewName(),
		Must:       NewAttributeTypeOIDList(),
		May:        NewAttributeTypeOIDList(),
		Extensions: NewExtensions(),
	}
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
func (r *NameForm) SetStringer(stringer func() string) NameForm {
	if r.IsZero() {
		r.nameForm = newNameForm()
	}

	r.nameForm.stringer = stringer

	return *r
}

/*
String is a stringer method that returns the string representation
of the receiver instance.
*/
func (r NameForm) String() (nf string) {
	if !r.IsZero() {
		if r.nameForm.stringer != nil {
			nf = r.nameForm.stringer()
		} else {
			nf = r.nameForm.s
		}
	}

	return
}

func (r NameForm) macro() (m []string) {
	if !r.IsZero() {
		m = r.nameForm.Macro
	}

	return
}

func (r NameForm) setOID(x string) {
	if !r.IsZero() {
		r.nameForm.OID = x
	}
}

/*
prepareString returns an error indicative of an attempt to represent
the receiver instance as a string using text/template.
*/
func (r *nameForm) prepareString() (err error) {
	buf := newBuf()
	r.t = newTemplate(`nameForm`).
		Funcs(funcMap(map[string]any{
			`ExtensionSet`: r.Extensions.tmplFunc,
			`MayLen`:       r.May.len,
			`IsObsolete`:   func() bool { return r.Obsolete },
		}))

	if r.t, err = r.t.Parse(nameFormTmpl); err == nil {
		if err = r.t.Execute(buf, struct {
			Definition *nameForm
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
Description returns the underlying (optional) descriptive text
assigned to the receiver instance.
*/
func (r NameForm) Description() (desc string) {
	if !r.IsZero() {
		desc = r.nameForm.Desc
	}

	return
}

/*
SetDescription parses desc into the underlying Desc field within the
receiver instance.  Although a RFC 4512-compliant QuotedString is
required, the outer single-quotes need not be specified literally.
*/
/*
func (r *NameForm) SetDescription(desc string) NameForm {
	if len(desc) < 3 {
		return *r
	}

	if r.nameForm == nil {
		r.nameForm = new(nameForm)
	}

	if !(rune(desc[0]) == rune(39) && rune(desc[len(desc)-1]) == rune(39)) {
		desc = `'` + desc + `'`
		if !r.IsZero() {
			r.nameForm.Desc = desc
		}
	}

	return *r
}
*/

/*
Inventory returns an instance of [Inventory] which represents the current
inventory of [NameForm] instances within the receiver.

The keys are numeric OIDs, while the values are zero (0) or more string
slices, each representing a name by which the definition is known.
*/
func (r NameForms) Inventory() (inv Inventory) {
	inv = make(Inventory, 0)
	for i := 0; i < r.len(); i++ {
		def := r.index(i)
		inv[def.NumericOID()] = def.Names().List()
	}

	return
}

/*
Type returns the string literal "nameForm".
*/
func (r NameForm) Type() string {
	return `nameForm`
}

/*
Type returns the string literal "nameForms".
*/
func (r NameForms) Type() string {
	return `nameForms`
}

// stackage closure func - do not exec directly.
func (r NameForms) canPush(x ...any) (err error) {
	if len(x) == 0 {
		return
	}

	for i := 0; i < len(x) && err == nil; i++ {
		nf, ok := x[i].(NameForm)
		if !ok || nf.IsZero() {
			err = errorf("Type assertion for %T has failed", x[i])
		} else if tst := r.get(nf.NumericOID()); !tst.IsZero() {
			err = errorf("%T %s not unique", nf, nf.NumericOID())
		}
	}

	return
}

/*
Len returns the current integer length of the receiver instance.
*/
func (r NameForms) Len() int {
	return r.len()
}

func (r NameForms) len() int {
	return r.cast().Len()
}

/*
String is a stringer method that returns the string representation
of the receiver instance.
*/
func (r NameForms) String() string {
	return r.cast().String()
}

/*
IsZero returns a Boolean value indicative of nilness of the
receiver instance.
*/
func (r NameForms) IsZero() bool {
	return r.cast().IsZero()
}

/*
Index returns the instance of [NameForm] found within the receiver
stack instance at index N.  If no instance is found at the index
specified, a zero [NameForm] instance is returned.
*/
func (r NameForms) Index(idx int) NameForm {
	return r.index(idx)
}

func (r NameForms) index(idx int) (nf NameForm) {
	slice, found := r.cast().Index(idx)
	if found {
		if _nf, ok := slice.(NameForm); ok {
			nf = _nf
		}
	}

	return
}

/*
Push returns an error following an attempt to push a [NameForm]
instance into the receiver stack instance.
*/
func (r NameForms) Push(nf any) error {
	return r.push(nf)
}

func (r NameForms) push(nf any) (err error) {
	if nf == nil {
		err = errorf("%T instance is nil; cannot append to %T", nf, r)
		return
	}
	r.cast().Push(nf)

	return
}

/*
Contains calls [NameForms.Get] to return a Boolean value indicative
of a successful, non-zero retrieval of a [NameForm] instance -- matching
the provided id -- from within the receiver stack instance.
*/
func (r NameForms) Contains(id string) bool {
	return r.contains(id)
}

func (r NameForms) contains(id string) bool {
	return !r.get(id).IsZero()
}

/*
Get returns an instance of [NameForm] based upon a search for id within
the receiver stack instance.

The return instance, if not nil, was retrieved based upon a textual
match of the principal identifier of a [NameForm] and the provided id.

The return instance is nil if no match was made.

Case is not significant in the matching process.
*/
func (r NameForms) Get(id string) NameForm {
	return r.get(id)
}

func (r NameForms) get(id string) (nf NameForm) {
	for i := 0; i < r.len() && nf.IsZero(); i++ {
		if _nf := r.index(i); !_nf.IsZero() {
			if _nf.nameForm.OID == id {
				nf = _nf
			} else if _nf.nameForm.Name.contains(id) {
				nf = _nf
			}
		}
	}

	return
}

/*
OC returns the STRUCTURAL [ObjectClass] specified within the receiver instance.
*/
func (r NameForm) OC() ObjectClass {
	return r.nameForm.Structural
}

/*
IsZero returns a Boolean value indicative of nilness of the
receiver instance.
*/
func (r NameForm) IsZero() bool {
	return r.nameForm == nil
}

func (r *nameForm) check() (err error) {
	if r == nil {
		err = errorf("%T is nil", r)
		return
	}

	if len(r.OID) == 0 {
		err = errorf("%T lacks an OID", r)
	}

	return
}

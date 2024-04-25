package schemax

/*
NewLDAPSyntaxes initializes a new [Collection] instance and casts it
as an [LDAPSyntaxes] instance.
*/
func NewLDAPSyntaxes() LDAPSyntaxes {
	r := LDAPSyntaxes(newCollection(`ldapSyntaxes`))
	r.cast().SetPushPolicy(r.canPush)

	return r
}

/*
IsIdentifiedAs returns a Boolean value indicative of whether id matches
either the numericOID or description of the receiver instance.  Case is
not significant in the matching process.
*/
func (r LDAPSyntax) IsIdentifiedAs(id string) (ident bool) {
	if !r.IsZero() {
		ident = id == r.NumericOID() || eq(r.lDAPSyntax.Desc, id)
	}

	return
}

/*
Len returns the current integer length of the receiver instance.
*/
func (r LDAPSyntaxes) Len() int {
	return r.len()
}

func (r LDAPSyntaxes) len() int {
	return r.cast().Len()
}

/*
SetSchema assigns an instance of [Schema] to the receiver instance.  This allows
internal verification of certain actions without the need for user input of
an instance of [Schema] manually at each juncture.

Note that the underlying [Schema] instance is automatically set when creating
instances of this type by way of parsing.

This is a fluent method.
*/
func (r *LDAPSyntax) SetSchema(schema Schema) *LDAPSyntax {
	if r.IsZero() {
		r.lDAPSyntax = newLDAPSyntax()
	}

	r.lDAPSyntax.schema = schema

	return r
}

func (r LDAPSyntax) schema() (s Schema) {
	if !r.IsZero() {
		s = r.lDAPSyntax.schema
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
func (r *LDAPSyntax) SetStringer(stringer func() string) *LDAPSyntax {
	if r.IsZero() {
		r.lDAPSyntax = newLDAPSyntax()
	}

	r.lDAPSyntax.stringer = stringer

	return r
}

/*
xOrigin returns an instance of LDAPSyntaxes containing only definitions
which bear the X-ORIGIN value of x. Case is not significant in the matching
process, nor is whitespace (e.g.: RFC 4517 vs. RFC4517).
*/
func (r LDAPSyntaxes) xOrigin(x string) (lss LDAPSyntaxes) {
	lss = NewLDAPSyntaxes()
	for i := 0; i < r.Len(); i++ {
		ls := r.Index(i)
		if xo, found := ls.Extensions().Get(`X-ORIGIN`); found {
			if xo.contains(x) {
				lss.push(ls)
			}
		}
	}

	return
}

/*
String is a stringer method that returns the string representation
of the receiver instance.
*/
func (r LDAPSyntax) String() (def string) {
	if !r.IsZero() {
		if r.stringer != nil {
			def = r.stringer()
		} else {
			if len(r.lDAPSyntax.s) == 0 {
				var err error
				if err = r.lDAPSyntax.prepareString(); err != nil {
					return
				}
			}

			def = r.lDAPSyntax.s
		}
	}

	return
}

/*
IsZero returns a Boolean value indicative of nilness of the
receiver instance.
*/
func (r LDAPSyntax) IsZero() bool {
	return r.lDAPSyntax == nil
}

/*
IsZero returns a Boolean value indicative of nilness of the
receiver instance.
*/
func (r LDAPSyntaxes) IsZero() bool {
	return r.cast().IsZero()
}

/*
OID is an alias for the [LDAPSyntax.NumericOID] method, as
[LDAPSyntax] instances do not allow for the assignment of
names.  This method exists solely to satisfy Go's interface
signature requirements.
*/
func (r LDAPSyntax) OID() string { return r.NumericOID() }

/*
Name returns an empty string, as [LDAPSyntax] definitions do not
bear names.  This method exists only to satisfy Go's interface
signature requirements.
*/
func (r LDAPSyntax) Name() string { return `` }

/*
Names returns an empty instance of [Name], as names do not
apply to [LDAPSyntax] definitions.  This method exists only to satisfy
Go's interface signature requirements.
*/
func (r LDAPSyntax) Names() Name { return Name{} }

/*
String returns the string representation of the receiver instance.
*/
func (r LDAPSyntaxes) String() string {
	return r.cast().String()
}

func (r LDAPSyntax) macro() (m []string) {
	if !r.IsZero() {
		m = r.lDAPSyntax.Macro
	}

	return
}

func (r LDAPSyntax) setOID(x string) {
	if !r.IsZero() {
		r.lDAPSyntax.OID = x
	}
}

func (r LDAPSyntaxes) canPush(x ...any) (err error) {
	if len(x) == 0 {
		return
	}

	for i := 0; i < len(x) && err == nil; i++ {
		instance := x[i]
		if ls, ok := instance.(LDAPSyntax); !ok || ls.IsZero() {
			err = errorf("Type assertion for %T has failed", instance)
		} else if tst := r.get(ls.NumericOID()); !tst.IsZero() {
			err = errorf("%T %s not unique", ls, ls.NumericOID())
		}
	}

	return
}

/*
SetNumericOID allows the manual assignment of a numeric OID to the
receiver instance if the following are all true:

  - The input id value is a syntactically valid numeric OID
  - The receiver does not already possess a numeric OID

This is a fluent method.
*/
func (r *LDAPSyntax) SetNumericOID(id string) *LDAPSyntax {
	if r.IsZero() {
		r.lDAPSyntax = newLDAPSyntax()
	}

	if isNumericOID(id) {
		if len(r.lDAPSyntax.OID) == 0 {
			r.lDAPSyntax.OID = id
		}
	}

	return r
}

/*
Extensions returns the [Extensions] instance -- if set -- within
the receiver.
*/
func (r LDAPSyntax) Extensions() (e Extensions) {
	if !r.IsZero() {
		e = r.lDAPSyntax.Extensions
	}

	return
}

/*
SetExtension assigns key x to value xstrs within the receiver's underlying
[Extensions] instance.

This is a fluent method.
*/
func (r *LDAPSyntax) SetExtension(x string, xstrs ...string) *LDAPSyntax {
	if r.IsZero() {
		r.lDAPSyntax = newLDAPSyntax()
	}

	r.Extensions().Set(x, xstrs...)

	return r
}

/*
Maps returns slices of [DefinitionMap] instances.
*/
func (r LDAPSyntaxes) Maps() (defs DefinitionMaps) {
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
func (r LDAPSyntax) Map() (def DefinitionMap) {
	if r.IsZero() {
		return
	}

	def = make(DefinitionMap, 0)
	def[`NUMERICOID`] = []string{r.NumericOID()}
	def[`DESC`] = []string{r.Description()}
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
Type returns the string literal "ldapSyntax".
*/
func (r LDAPSyntax) Type() string {
	return `ldapSyntax`
}

/*
Type returns the string literal "ldapSyntaxes".
*/
func (r LDAPSyntaxes) Type() string {
	return `ldapSyntaxes`
}

/*
Description returns the underlying (optional) descriptive text
assigned to the receiver instance.
*/
func (r LDAPSyntax) Description() (desc string) {
	if !r.IsZero() {
		desc = r.lDAPSyntax.Desc
	}

	return
}

/*
SetDescription parses desc into the underlying Desc field within the
receiver instance.  Although a RFC 4512-compliant QuotedString is
required, the outer single-quotes need not be specified literally.

This is a fluent method.
*/
func (r *LDAPSyntax) SetDescription(desc string) *LDAPSyntax {
	if len(desc) < 3 {
		return r
	}

	if r.lDAPSyntax == nil {
		r.lDAPSyntax = new(lDAPSyntax)
	}

	if !(rune(desc[0]) == rune(39) && rune(desc[len(desc)-1]) == rune(39)) {
		if !r.IsZero() {
			r.lDAPSyntax.Desc = desc
		}
	}

	return r
}

/*
IsObsolete only returns a false Boolean value, as definition obsolescence
does not apply to [LDAPSyntax] definitions.  This method exists only to
satisfy Go's interface signature requirements.
*/
func (r LDAPSyntax) IsObsolete() bool { return false }

/*
Get returns an instance of [LDAPSyntax] based upon a search for id within
the receiver stack instance.

The return instance, if not nil, was retrieved based upon a textual match of
either of the following:

  - the numeric OID of an [LDAPSyntax] and the provided id
  - the description text -- minus whitespace -- and the provided id

The return instance is nil if no match was made.

Case is not significant in the matching process.
*/
func (r LDAPSyntaxes) Get(id string) LDAPSyntax {
	return r.get(id)
}

func (r LDAPSyntaxes) get(id string) (ls LDAPSyntax) {
	for i := 0; i < r.len() && ls.IsZero(); i++ {
		if _ls := r.index(i); !_ls.IsZero() {
			if eq(_ls.lDAPSyntax.OID, id) {
				ls = _ls
			} else {
				if desc := repAll(_ls.lDAPSyntax.Desc, ` `, ``); eq(desc, id) {
					ls = _ls
				}
			}
		}
	}

	return
}

/*
NumericOID returns the string representation of the numeric OID
held by the receiver instance.
*/
func (r LDAPSyntax) NumericOID() (noid string) {
	if !r.IsZero() {
		noid = r.lDAPSyntax.OID
	}

	return
}

/*
Index returns the instance of [LDAPSyntax] found within the
receiver stack instance at index N.  If no instance is found at
the index specified, a zero [LDAPSyntax] instance is returned.
*/
func (r LDAPSyntaxes) Index(idx int) LDAPSyntax {
	return r.index(idx)
}

func (r LDAPSyntaxes) index(idx int) (ls LDAPSyntax) {
	slice, found := r.cast().Index(idx)
	if found {
		if _ls, ok := slice.(LDAPSyntax); ok {
			ls = _ls
		}
	}

	return
}

/*
Push returns an error following an attempt to push an [LDAPSyntax]
into the receiver stack instance.
*/
func (r LDAPSyntaxes) Push(ls any) error {
	return r.push(ls)
}

func (r LDAPSyntaxes) push(ls any) (err error) {
	if ls == nil {
		err = errorf("%T instance is nil; cannot append to %T", ls, r)
		return
	}

	r.cast().Push(ls)

	return
}

/*
NewLDAPSyntax initializes and returns a new instance of [LDAPSyntax].
*/
func NewLDAPSyntax() LDAPSyntax {
	return LDAPSyntax{newLDAPSyntax()}
}

func newLDAPSyntax() *lDAPSyntax {
	return &lDAPSyntax{
		Extensions: NewExtensions(),
	}
}

func (r *lDAPSyntax) prepareString() (err error) {
	buf := newBuf()
	r.t = newTemplate(`ldapSyntax`).
		Funcs(funcMap(map[string]any{
			`ExtensionSet`: r.Extensions.tmplFunc,
		}))

	if r.t, err = r.t.Parse(lDAPSyntaxTmpl); err == nil {
		if err = r.t.Execute(buf, struct {
			Definition *lDAPSyntax
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
LDAPSyntaxes returns the [LDAPSyntaxes] instance from within the
receiver instance.
*/
func (r Schema) LDAPSyntaxes() (lss LDAPSyntaxes) {
	slice, _ := r.cast().Index(ldapSyntaxesIndex)
	lss, _ = slice.(LDAPSyntaxes)
	return
}

/*
LoadLDAPSyntaxes will load all package-included [LDAPSyntax] definitions
into the receiver instance.
*/
func (r Schema) LoadLDAPSyntaxes() Schema {
	_ = r.loadSyntaxes()
	return r
}

/*
loadSyntaxes returns an error following an attempt to load all
built-in ldapSyntax definitions found within this package into
the receiver instance.
*/
func (r Schema) loadSyntaxes() (err error) {
	if !r.IsZero() {
		for _, funk := range []func() error{
			r.loadRFC4517Syntaxes,
			r.loadRFC4523Syntaxes,
			r.loadRFC4530Syntaxes,
			r.loadRFC2307Syntaxes,
		} {
			if err = funk(); err != nil {
				break
			}
		}
	}

	return
}

/*
LoadRFC2307Syntaxes returns an error following an attempt to load
all RFC 2307 [LDAPSyntax] slices into the receiver instance.
*/
func (r Schema) LoadRFC2307Syntaxes() Schema {
	_ = r.loadRFC2307Syntaxes()
	return r
}

func (r Schema) loadRFC2307Syntaxes() (err error) {
	for k, v := range rfc2307Macros {
		r.SetMacro(k, v)
	}

	for i := 0; i < len(rfc2307LDAPSyntaxes) && err == nil; i++ {
		ls := rfc2307LDAPSyntaxes[i]
		err = r.ParseLDAPSyntax(string(ls))
	}

	return
}

/*
LoadRFC4517Syntaxes returns an error following an attempt to
load all RFC 4517 [LDAPSyntax] slices into the receiver instance.
*/
func (r Schema) LoadRFC4517Syntaxes() Schema {
	_ = r.loadRFC4517Syntaxes()
	return r
}

func (r Schema) loadRFC4517Syntaxes() (err error) {
	for i := 0; i < len(rfc4517LDAPSyntaxes) && err == nil; i++ {
		ls := rfc4517LDAPSyntaxes[i]
		err = r.ParseLDAPSyntax(string(ls))
	}

	return
}

/*
LoadRFC4523Syntaxes returns an error following an attempt to
load all RFC 4523 [LDAPSyntax] slices into the receiver instance.
*/
func (r Schema) LoadRFC4523Syntaxes() Schema {
	_ = r.loadRFC4523Syntaxes()
	return r
}

func (r Schema) loadRFC4523Syntaxes() (err error) {
	for i := 0; i < len(rfc4523LDAPSyntaxes) && err == nil; i++ {
		ls := rfc4523LDAPSyntaxes[i]
		err = r.ParseLDAPSyntax(string(ls))
	}

	return
}

/*
LoadRFC4530Syntaxes returns an error following an attempt to
load all RFC 4530 [LDAPSyntax] slices into the receiver instance.
*/
func (r Schema) LoadRFC4530Syntaxes() Schema {
	_ = r.loadRFC4530Syntaxes()
	return r
}

func (r Schema) loadRFC4530Syntaxes() (err error) {
	for i := 0; i < len(rfc4530LDAPSyntaxes) && err == nil; i++ {
		ls := rfc4530LDAPSyntaxes[i]
		err = r.ParseLDAPSyntax(string(ls))
	}

	return
}

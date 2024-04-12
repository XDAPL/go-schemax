package schemax

import (
	antlr4512 "github.com/JesseCoretta/go-rfc4512-antlr"
)

/*
ParseNameForm returns an error following an attempt to parse raw into the
receiver instance's [NameForms] instance.
*/
func (r *Schema) ParseNameForm(raw string) (err error) {
	var i antlr4512.Instance
	if i, err = antlr4512.ParseInstance(raw); err == nil {
		var n NameForm
		n, err = r.processNameForm(i.P.NameFormDescription())
		if err == nil {
			err = r.NameForms().push(n)
		}
	}

	return
}

/*
NewNameForms initializes a new [NameForms] instance.
*/
func NewNameForms() NameForms {
	r := NameForms(newCollection(``))
	r.cast().SetPushPolicy(r.canPush)

	return r
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
SetName allows the manual assignment of one (1) or more RFC 4512-compliant
descriptor values by which the receiver instance is to be known.  This will
append to -- not replace -- any preexisting names.

This is a fluent method.
*/
func (r NameForm) SetName(name ...string) NameForm {
	if r.nameForm == nil {
		r.nameForm = new(nameForm)
	}

	if !r.IsZero() {
		for _, n := range name {
			r.nameForm.Name.cast().Push(n)
		}
	}

	return r
}

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
Names returns the underlying instance of [DefinitionName] from within the
receiver instance.
*/
func (r NameForm) Names() (names DefinitionName) {
	return r.nameForm.Name
}

/*
String is a stringer method that returns the string representation
of the receiver instance.
*/
func (r NameForm) String() (nf string) {
	if !r.IsZero() {
		nf = r.nameForm.s
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
		if err = r.t.Execute(buf, r); err == nil {
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
func (r *NameForm) SetDescription(desc string) NameForm {
	if len(desc) < 3 {
		return *r
	}

	if r.nameForm == nil {
		r.nameForm = new(nameForm)
	}

	if !(rune(desc[0]) == rune(39) && rune(desc[len(desc)-1]) == rune(39)) {
		desc = `'` + desc + `'`
		if !r.IsZero() && isValidDescription(desc) {
			r.nameForm.Desc = desc
		}
	}

	return *r
}

/*
List returns a map[string][]string instance which represents the current
inventory of name form instances within the receiver.  The keys are
numeric OIDs, while the values are zero (0) or more string slices, each
representing a name by which the definition is known.
*/
func (r NameForms) List() (list map[string][]string) {
	list = make(map[string][]string, 0)
	for i := 0; i < r.len(); i++ {
		def := r.index(i)
		list[def.NumericOID()] = def.Names().List()

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

func (r Schema) processNameForm(ctx antlr4512.INameFormDescriptionContext) (nf NameForm, err error) {

	_nf := new(nameForm)
	_nf.schema = r

	for k, ct := 0, ctx.GetChildCount(); k < ct && err == nil; k++ {
		switch tv := ctx.GetChild(k).(type) {
		case *antlr4512.OpenParenContext,
			*antlr4512.CloseParenContext,
			*antlr4512.NumericOIDOrMacroContext:
			err = _nf.setCritical(tv)

		case *antlr4512.DefinitionNameContext,
			*antlr4512.DefinitionExtensionsContext,
			*antlr4512.DefinitionDescriptionContext:
			err = _nf.setMisc(tv)

		case *antlr4512.DefinitionObsoleteContext:
			_nf.Obsolete = true

		case *antlr4512.NFStructuralOCContext:
			err = _nf.structuralOCContext(tv)

		case *antlr4512.DefinitionMustContext:
			err = _nf.mustContext(tv)

		case *antlr4512.DefinitionMayContext:
			err = _nf.mayContext(tv)
		default:
			env := NameForm{_nf}
			err = isErrImpl(env.Type(), env.OID(), tv)
		}
	}

	// check for errors in new instance
	if err == nil {
		r.resolveByMacro(NameForm{_nf})

		if err = _nf.check(); err == nil {
			if err = _nf.prepareString(); err == nil {
				_nf.t = nil
				nf = NameForm{_nf}
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

func (r *nameForm) setCritical(ctx any) (err error) {

	switch tv := ctx.(type) {
	case *antlr4512.NumericOIDOrMacroContext:
		r.OID, r.Macro, err = numOIDContext(tv)
	case *antlr4512.OpenParenContext,
		*antlr4512.CloseParenContext:
		err = parenContext(tv)
	default:
		err = errorf("Unknown critical context '%T'", ctx)
	}

	return
}

func (r *nameForm) setMisc(ctx any) (err error) {

	switch tv := ctx.(type) {
	case *antlr4512.DefinitionNameContext:
		r.Name, err = nameContext(tv)
	case *antlr4512.DefinitionDescriptionContext:
		r.Desc, err = descContext(tv)
	case *antlr4512.DefinitionExtensionsContext:
		r.Extensions, err = extContext(tv)
	default:
		err = errorf("Unknown miscellaneous context '%T'", ctx)
	}

	return
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

func (r *nameForm) mustContext(ctx *antlr4512.DefinitionMustContext) (err error) {
	var must []string
	if must, err = mustContext(ctx); err != nil {
		return
	}

	r.Must = NewAttributeTypeOIDList()
	for i := 0; i < len(must); i++ {
		var mustt AttributeType
		if mustt = r.schema.AttributeTypes().get(must[i]); mustt.IsZero() {
			err = errorf("required attr '%s' not found; cannot process %T",
				must[i], r)
			break
		}
		r.Must.push(mustt)
	}

	return
}

func (r *nameForm) mayContext(ctx *antlr4512.DefinitionMayContext) (err error) {
	var may []string
	if may, err = mayContext(ctx); err != nil {
		return
	}

	r.May = NewAttributeTypeOIDList()
	for i := 0; i < len(may); i++ {
		var mayy AttributeType
		if mayy = r.schema.AttributeTypes().get(may[i]); mayy.IsZero() {
			err = errorf("required attr '%s' not found; cannot process %T",
				may[i], r)
			break
		}
		r.May.push(mayy)
	}

	return
}

func (r *nameForm) structuralOCContext(ctx *antlr4512.NFStructuralOCContext) (err error) {
	if ctx == nil {
		err = errorf("%T instance is nil", ctx)
		return
	}

	// Obtain OIDContext pointer
	o := ctx.OID()
	if o == nil {
		err = errorf("%T instance is nil", o)
		return
	}

	// Type assert o to a bonafide
	// *OIDContext instance.
	var n, d []string
	assert, ok := o.(*antlr4512.OIDContext)
	if !ok {
		err = errorf("structural objectClass OID assertion failed for %T", r)
		return
	}

	// determine the nature of the OID and process it.
	// Fail if we find anything else.
	if n, d, err = oIDContext(assert); err != nil {
		return
	}

	var oc ObjectClass
	if len(n) > 0 {
		oc = r.schema.ObjectClasses().get(n[0])
	} else if len(d) > 0 {
		oc = r.schema.ObjectClasses().get(d[0])
	} else {
		err = errorf("No %T OID or Descriptor found", r)
		return
	}

	if oc.IsZero() {
		err = errorf("structural class '%s' not found; cannot process %T",
			assert.GetText(), r)
	} else {
		r.Structural = oc
	}

	return
}

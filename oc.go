package schemax

import (
	"internal/rfc2079"
	"internal/rfc2798"
	"internal/rfc3671"
	"internal/rfc3672"
	"internal/rfc4512"
	"internal/rfc4519"
	"internal/rfc4523"
	"internal/rfc4524"

	antlr4512 "github.com/JesseCoretta/go-rfc4512-antlr"
)

var (
	rfc2079ObjectClasses rfc2079.ObjectClassDefinitions = rfc2079.AllObjectClasses
	rfc2798ObjectClasses rfc2798.ObjectClassDefinitions = rfc2798.AllObjectClasses
	rfc3671ObjectClasses rfc3671.ObjectClassDefinitions = rfc3671.AllObjectClasses
	rfc3672ObjectClasses rfc3672.ObjectClassDefinitions = rfc3672.AllObjectClasses
	rfc4512ObjectClasses rfc4512.ObjectClassDefinitions = rfc4512.AllObjectClasses
	rfc4519ObjectClasses rfc4519.ObjectClassDefinitions = rfc4519.AllObjectClasses
	rfc4523ObjectClasses rfc4523.ObjectClassDefinitions = rfc4523.AllObjectClasses
	rfc4524ObjectClasses rfc4524.ObjectClassDefinitions = rfc4524.AllObjectClasses
)

/*
ParseObjectClass parses an individual textual object class (raw) and
returns an error instance.

When no error occurs, the newly formed [ObjectClass] instance -- based
on the parsed contents of raw -- is added to the receiver [ObjectClasses]
instance.
*/
func (r Schema) ParseObjectClass(raw string) (err error) {
	i, err := parseI(raw)
	if err == nil {
		var o ObjectClass
		o, err = r.processObjectClass(i.P.ObjectClassDescription())
		if err == nil {
			err = r.ObjectClasses().push(o)
		}
	}

	return
}

/*
NewObjectClasses initializes and returns a new [ObjectClasses] instance,
configured to allow the storage of all [ObjectClass] instances.
*/
func NewObjectClasses() ObjectClasses {
	r := ObjectClasses(newCollection(``))
	r.cast().SetPushPolicy(r.canPush)

	return r
}

/*
NewObjectClassOIDList initializes and returns a new [ObjectClasses] that
has been cast from an instance of [OIDList] and configured to allow the
storage of arbitrary [ObjectClass] instances.
*/
func NewObjectClassOIDList() ObjectClasses {
	r := ObjectClasses(newOIDList(``))
	r.cast().
		SetPushPolicy(r.canPush).
		SetPresentationPolicy(r.oIDsStringer)

	return r
}

/*
IsObsolete returns a Boolean value indicative of definition obsolescence.
*/
func (r ObjectClass) IsObsolete() (o bool) {
	if !r.IsZero() {
		o = r.objectClass.Obsolete
	}

	return
}

/*
Kind returns the uint-based kind value associated with the receiver
instance.
*/
func (r ObjectClass) Kind() (kind uint) {
	kind = StructuralKind
	if !r.IsZero() {
		switch k := r.objectClass.Kind; k {
		case AbstractKind:
			kind = k
		case AuxiliaryKind:
			kind = k
		}
	}

	return
}

/*
SuperClasses returns an [ObjectClasses] containing zero (0) or more
superior [ObjectClass] instances from which the receiver extends.
*/
func (r ObjectClass) SuperClasses() (sup ObjectClasses) {
	if !r.IsZero() {
		sup = r.objectClass.SuperClasses
	}

	return
}

/*
Type returns the string literal "objectClass".
*/
func (r ObjectClass) Type() string {
	return `objectClass`
}

/*
Type returns the string literal "objectClasses".
*/
func (r ObjectClasses) Type() string {
	return `objectClasses`
}

/*
AllMust returns an [AttributeTypes] containing zero (0) or more required
[AttributeType] definitions for use with this class as well as those specified
by any and all applicable super classes. Duplicate references are silently
discarded.
*/
func (r ObjectClass) AllMust() (must AttributeTypes) {
	must = NewAttributeTypeOIDList()

	if sup := r.SuperClasses(); !sup.IsZero() {
		for i := 0; i < sup.len(); i++ {
			sm := sup.index(i)
			if sc := sm.AllMust(); !sc.IsZero() {
				for j := 0; j < sc.len(); j++ {
					must.push(sc.index(i))
				}
			}
		}
	}

	if lm := r.Must(); !lm.IsZero() {
		for i := 0; i < lm.len(); i++ {
			must.push(lm.index(i))
		}
	}

	return
}

/*
Must returns an [AttributeTypes] containing zero (0) or more required
[AttributeType] definitions for use with this class.
*/
func (r ObjectClass) Must() (must AttributeTypes) {
	if !r.IsZero() {
		must = r.objectClass.Must
	}

	return
}

/*
AllMay returns an [AttributeTypes] containing zero (0) or more allowed
[AttributeType] definitions for use with this class as well as those
specified by any and all applicable super classes.  Duplicate references
are silently discarded.
*/
func (r ObjectClass) AllMay() (may AttributeTypes) {
	may = NewAttributeTypeOIDList()

	if sup := r.SuperClasses(); !sup.IsZero() {
		for i := 0; i < sup.len(); i++ {
			sm := sup.index(i)
			if sc := sm.AllMay(); !sc.IsZero() {
				for j := 0; j < sc.len(); j++ {
					may.push(sc.index(i))
				}
			}
		}
	}

	if lm := r.May(); !lm.IsZero() {
		for i := 0; i < lm.len(); i++ {
			may.push(lm.index(i))
		}
	}

	return
}

/*
May returns an [AttributeTypes] containing zero (0) or more allowed
[AttributeType] definitions for use with this class.
*/
func (r ObjectClass) May() (may AttributeTypes) {
	if !r.IsZero() {
		may = r.objectClass.May
	}

	return
}

/*
Attributes returns an instance of [AttributeTypes] containing references
to [AttributeType] instances which reside within the receiver's MUST and MAY
clauses.  This is useful for accessing a simple inventory of any and all
available [AttributeType] definitions.
*/
func (r ObjectClass) Attributes() (a AttributeTypes) {
	a = NewAttributeTypeOIDList()

	for _, list := range []AttributeTypes{
		r.objectClass.Must,
		r.objectClass.May,
	} {
		for i := 0; i < list.len(); i++ {
			if at := list.index(i); !at.IsZero() {
				a.push(at)
			}
		}
	}

	return
}

/*
OID returns the string representation of an OID -- which is either a
numeric OID or descriptor -- that is held by the receiver instance.
*/
func (r ObjectClass) OID() (oid string) {
	if !r.IsZero() {
		oid = r.NumericOID() // default
		if r.objectClass.Name.len() > 0 {
			oid = r.objectClass.Name.index(0)
		}
	}

	return
}

/*
IsIdentifiedAs returns a Boolean value indicative of whether id matches
either the numericOID or descriptor of the receiver instance.  Case is
not significant in the matching process.
*/
func (r ObjectClass) IsIdentifiedAs(id string) (ident bool) {
	if !r.IsZero() {
		ident = id == r.NumericOID() ||
			r.objectClass.Name.contains(id)
	}

	return
}

/*
Name returns the string form of the principal name of the receiver instance, if set.
*/
func (r ObjectClass) Name() (id string) {
	if !r.IsZero() {
		id = r.objectClass.Name.index(0)
	}

	return
}

/*
Names returns the underlying instance of [DefinitionName] from
within the receiver.
*/
func (r ObjectClass) Names() (names DefinitionName) {
	return r.objectClass.Name
}

/*
Extensions returns the [Extensions] instance -- if set -- within
the receiver.
*/
func (r ObjectClass) Extensions() (e Extensions) {
	if !r.IsZero() {
		e = r.objectClass.Extensions
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
func (r *ObjectClass) SetNumericOID(id string) ObjectClass {
	if r.objectClass == nil {
		r.objectClass = newObjectClass()
	}

	if isNumericOID(id) {
		if len(r.objectClass.OID) == 0 {
			r.objectClass.OID = id
		}
	}

	return *r
}

/*
NumericOID returns the string representation of the numeric OID
held by the receiver instance.
*/
func (r ObjectClass) NumericOID() (noid string) {
	if !r.IsZero() {
		noid = r.objectClass.OID
	}

	return
}

/*
String is a stringer method that returns the string representation
of the receiver instance.
*/
func (r ObjectClass) String() (oc string) {
	if !r.IsZero() {
		oc = r.objectClass.s
	}

	return
}

/*
SetSuperClass appends the input value(s) to the the super classes stack within the
the receiver. Valid input types are string, to represent an RFC 4512 OID residing
in the underlying [Schema] instance, or an actual [ObjectClass] instance already
obtained or crafted.

This is a fluent method.
*/
func (r *ObjectClass) SetSuperClass(x ...any) ObjectClass {
	if r.IsZero() {
		r.objectClass = newObjectClass()
	}

	var err error
	for i := 0; i < len(x) && err == nil; i++ {
		var sup ObjectClass
		switch tv := x[i].(type) {
		case string:
			sup = r.schema.ObjectClasses().get(tv)
		case ObjectClass:
			sup = tv
		default:
			err = errorf("Unsupported super class type %T", tv)
			continue
		}

		err = r.objectClass.verifySuperClass(sup.objectClass)
		if err == nil && !sup.IsZero() {
			r.objectClass.SuperClasses.Push(sup)
		}
	}

	return *r
}

/*
verifySuperClass returns an error following the execution of basic sanity
checks meant to assess the intended super type chain.
*/
func (r *objectClass) verifySuperClass(sup *objectClass) (err error) {
	// perform basic sanity check of super type
	if err = sup.check(); err != nil {
		return
	}

	// make sure super type and sub type aren't
	// one-in-the-same.
	if r.OID == sup.OID {
		err = errorf("cyclical super type loop detected (%s)", r.OID)
	}

	return
}

func newObjectClass() *objectClass {
	return &objectClass{
		Must:         NewAttributeTypeOIDList(),
		May:          NewAttributeTypeOIDList(),
		SuperClasses: NewObjectClassOIDList(),
		Extensions:   NewExtensions(),
	}
}

func (r ObjectClass) macro() (m []string) {
	if !r.IsZero() {
		m = r.objectClass.Macro
	}

	return
}

func (r ObjectClass) setOID(x string) {
	if !r.IsZero() {
		r.objectClass.OID = x
	}
}

func (r *objectClass) prepareString() (err error) {
	buf := newBuf()
	r.t = newTemplate(`objectClass`).
		Funcs(funcMap(map[string]any{
			`ExtensionSet`: r.Extensions.tmplFunc,
			`MustLen`:      r.Must.len,
			`MayLen`:       r.May.len,
			`SuperLen`:     r.SuperClasses.len,
			`IsObsolete`:   func() bool { return r.Obsolete },
		}))
	if r.t, err = r.t.Parse(objectClassTmpl); err == nil {
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
func (r ObjectClass) Description() (desc string) {
	if !r.IsZero() {
		desc = r.objectClass.Desc
	}

	return
}

/*
SetDescription parses desc into the underlying Desc field within the
receiver instance.  Although a RFC 4512-compliant QuotedString is
required, the outer single-quotes need not be specified literally.

This is a fluent method.
*/
func (r *ObjectClass) SetDescription(desc string) ObjectClass {
	if len(desc) == 0 {
		return *r
	}

	if r.objectClass == nil {
		r.objectClass = new(objectClass)
	}

	if !(rune(desc[0]) == rune(39) && rune(desc[len(desc)-1]) == rune(39)) {
		desc = `'` + desc + `'`
		if !r.IsZero() && isValidDescription(desc) {
			r.objectClass.Desc = desc
		}
	}

	return *r
}

/*
List returns a map[string][]string instance which represents the current
inventory of object class instances within the receiver.  The keys are
numeric OIDs, while the values are zero (0) or more string slices, each
representing a name by which the definition is known.
*/
func (r ObjectClasses) List() (list map[string][]string) {
	list = make(map[string][]string, 0)
	for i := 0; i < r.len(); i++ {
		def := r.index(i)
		list[def.NumericOID()] = def.Names().List()
	}

	return
}

/*
IsZero returns a Boolean value indicative of nilness of the
receiver instance.
*/
func (r ObjectClasses) IsZero() bool {
	return r.cast().IsZero()
}

// stackage closure func - do not exec directly (use String method)
func (r ObjectClasses) oIDsStringer(_ ...any) (present string) {
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

// stackage closure func - do not exec directly.
func (r ObjectClasses) canPush(x ...any) (err error) {
	if len(x) == 0 {
		return
	}

	instance := x[0]
	oc, ok := instance.(ObjectClass)
	err = errorf("Type assertion for %T has failed", oc)
	if ok {
		err = nil
		// scan receiver for push candidate
		for i := 0; i < r.cast().Len() && err == nil; i++ {
			slice, _ := r.cast().Index(i)
			switch tv := slice.(type) {
			case ObjectClass:
				if tv.objectClass.OID == oc.objectClass.OID {
					err = errorf("%T not unique", oc)
					break
				}
			default:
				err = errorf("Unsupported type %T", tv)
			}
		}
	}

	return
}

/*
Len returns the current integer length of the receiver instance.
*/
func (r ObjectClasses) Len() int {
	return r.len()
}

func (r ObjectClasses) len() int {
	return r.cast().Len()
}

/*
String is a stringer method that returns the string representation
of the receiver instance.
*/
func (r ObjectClasses) String() string {
	return r.cast().String()
}

/*
Index returns the instance of [ObjectClass] found within the
receiver stack instance at index N.  If no instance is found at
the index specified, a zero [ObjectClass] instance is returned.
*/
func (r ObjectClasses) Index(idx int) ObjectClass {
	return r.index(idx)
}

func (r ObjectClasses) index(idx int) (oc ObjectClass) {
	slice, found := r.cast().Index(idx)
	if found {
		if _oc, ok := slice.(ObjectClass); ok {
			oc = _oc
		}
	}

	return
}

/*
Push returns an error following an attempt to push an [ObjectClass]
into the receiver stack instance.
*/
func (r ObjectClasses) Push(oc any) error {
	return r.push(oc)
}

func (r ObjectClasses) push(oc any) (err error) {
	err = errorf("%T instance is nil; cannot append to %T", oc, r)
	if oc != nil {
		r.cast().Push(oc)
		err = nil
	}

	return
}

/*
Contains calls [ObjectClasses.Get] to return a Boolean value indicative
of a successful, non-zero retrieval of an [ObjectClass] instance -- matching
the provided id -- from within the receiver stack instance.
*/
func (r ObjectClasses) Contains(id string) bool {
	return r.contains(id)
}

func (r ObjectClasses) contains(id string) bool {
	return !r.get(id).IsZero()
}

/*
Get returns an instance of [ObjectClass] based upon a search for id within
the receiver stack instance.

The return instance, if not nil, was retrieved based upon a textual match of
the principal identifier of an [ObjectClass] and the provided id.

The return instance is nil if no match was made.

Case is not significant in the matching process.
*/
func (r ObjectClasses) Get(id string) ObjectClass {
	return r.get(id)
}

func (r ObjectClasses) get(id string) (oc ObjectClass) {
	for i := 0; i < r.len() && oc.IsZero(); i++ {
		if _oc := r.index(i); !_oc.IsZero() {
			if _oc.objectClass.OID == id {
				oc = _oc
			} else if _oc.objectClass.Name.contains(id) {
				oc = _oc
			}
		}
	}

	return
}

func (r Schema) processObjectClass(ctx antlr4512.IObjectClassDescriptionContext) (oc ObjectClass, err error) {

	_oc := new(objectClass)
	_oc.schema = r

	for k, ct := 0, ctx.GetChildCount(); k < ct && err == nil; k++ {
		switch tv := ctx.GetChild(k).(type) {
		case *antlr4512.NumericOIDOrMacroContext,
			*antlr4512.OpenParenContext, *antlr4512.CloseParenContext:
			err = _oc.setCritical(tv)
		case *antlr4512.DefinitionNameContext,
			*antlr4512.DefinitionExtensionsContext,
			*antlr4512.DefinitionDescriptionContext:
			err = _oc.setMisc(tv)
		case *antlr4512.DefinitionObsoleteContext:
			_oc.Obsolete = true
		case *antlr4512.OCKindContext:
			_oc.Kind = StructuralKind // default
			if strx := tv.AuxiliaryKind(); strx != nil {
				_oc.Kind = AuxiliaryKind
			} else if stra := tv.AbstractKind(); stra != nil {
				_oc.Kind = AbstractKind
			}
		case *antlr4512.OCSuperClassesContext:
			err = _oc.superClassesContext(r, tv)
		case *antlr4512.DefinitionMustContext:
			err = _oc.mustContext(r, tv)
		case *antlr4512.DefinitionMayContext:
			err = _oc.mayContext(r, tv)
		default:
			env := ObjectClass{_oc}
			err = isErrImpl(env.Type(), env.OID(), tv)
		}
	}

	// check for errors in new instance
	if err == nil {
		r.resolveByMacro(ObjectClass{_oc})

		if err = _oc.check(); err == nil {
			if err = _oc.prepareString(); err == nil {
				_oc.t = nil
				oc = ObjectClass{_oc}
			}
		}
	}

	return
}

/*
IsZero returns a Boolean value indicative of nilness of the
receiver instance.
*/
func (r ObjectClass) IsZero() bool {
	return r.objectClass == nil
}

func (r *objectClass) setCritical(ctx any) (err error) {
	err = errorf("Unknown critical context '%T'", ctx)

	switch tv := ctx.(type) {
	case *antlr4512.NumericOIDOrMacroContext:
		r.OID, r.Macro, err = numOIDContext(tv)
	case *antlr4512.OpenParenContext, *antlr4512.CloseParenContext:
		err = parenContext(tv)
	}

	return
}

func (r *objectClass) setMisc(ctx any) (err error) {
	err = errorf("Unknown miscellaneous context '%T'", ctx)

	switch tv := ctx.(type) {
	case *antlr4512.DefinitionNameContext:
		r.Name, err = nameContext(tv)
	case *antlr4512.DefinitionDescriptionContext:
		r.Desc, err = descContext(tv)
	case *antlr4512.DefinitionExtensionsContext:
		r.Extensions, err = extContext(tv)
	}

	return
}

func superClassContext(ctx *antlr4512.OCSuperClassesContext) (sup []string, err error) {
	if ctx != nil {
		var n, d []string

		if x := ctx.OID(); x != nil {
			n, d, err = oIDContext(x.(*antlr4512.OIDContext))
		} else if y := ctx.OIDs(); y != nil {
			n, d, err = oIDsContext(y.(*antlr4512.OIDsContext))
		}

		if err == nil {
			sup = append(sup, n...)
			sup = append(sup, d...)
		}

		if len(sup) == 0 {
			if err != nil {
				err = errorf("No super classes parsed from %T: %v", ctx, err)
			} else {
				err = errorf("No super classes parsed from %T", ctx)
			}
		}
	}

	return
}

func (r *objectClass) superClassesContext(s Schema, ctx *antlr4512.OCSuperClassesContext) (err error) {
	var sup []string
	if sup, err = superClassContext(ctx); len(sup) == 0 {
		err = errorf("failed to process %T", ctx)
		return
	}

	r.SuperClasses = NewObjectClassOIDList()
	for i := 0; i < len(sup); i++ {
		var soc ObjectClass
		if soc = s.ObjectClasses().get(sup[i]); soc.IsZero() {
			err = errorf("super class '%s' not found; cannot process %T",
				sup[i], r)
			break
		}
		r.SuperClasses.push(soc)
	}

	return
}

func (r *objectClass) mustContext(s Schema, ctx *antlr4512.DefinitionMustContext) (err error) {
	var must []string
	if must, err = mustContext(ctx); err != nil {
		return
	}

	r.Must = NewAttributeTypeOIDList()
	for i := 0; i < len(must); i++ {
		var mustt AttributeType
		if mustt = s.AttributeTypes().get(must[i]); mustt.IsZero() {
			err = errorf("required attr '%s' not found; cannot process %T",
				must[i], r)
			break
		}
		r.Must.push(mustt)
	}

	return
}

func (r *objectClass) mayContext(s Schema, ctx *antlr4512.DefinitionMayContext) (err error) {
	var may []string
	if may, err = mayContext(ctx); err != nil {
		return
	}

	r.May = NewAttributeTypeOIDList()
	for i := 0; i < len(may); i++ {
		var mayy AttributeType
		if mayy = s.AttributeTypes().get(may[i]); mayy.IsZero() {
			err = errorf("required attr '%s' not found; cannot process %T",
				may[i], r)
			break
		}
		r.May.push(mayy)
	}

	return
}

func (r *objectClass) check() (err error) {

	if r == nil {
		err = errorf("%T is nil", r)
		return
	}
	if len(r.OID) == 0 {
		err = errorf("%T lacks an OID", r)
	}

	return
}

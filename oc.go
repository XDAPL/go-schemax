package schemax

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
NewObjectClass initializes and returns a new instance of
[ObjectClass], ready for manual population.

This function need not be used when parsing is engaged.
*/
func NewObjectClass() ObjectClass {
	return ObjectClass{newObjectClass()}
}

func newObjectClass() *objectClass {
	return &objectClass{
		Name:         NewName(),
		Must:         NewAttributeTypeOIDList(),
		May:          NewAttributeTypeOIDList(),
		SuperClasses: NewObjectClassOIDList(),
		Extensions:   NewExtensions(),
	}
}

/*
SetName assigns the provided names to the receiver instance.

Name instances must conform to RFC 4512 descriptor format but
need not be quoted.

This is a fluent method.
*/
func (r *ObjectClass) SetName(x ...string) *ObjectClass {
	if len(x) == 0 {
		return r
	}

	if r.IsZero() {
		r.objectClass = newObjectClass()
	}

	for i := 0; i < len(x); i++ {
		r.objectClass.Name.Push(x[i])
	}

	return r
}

/*
SetKind assigns the abstraction of an objectClass "kind" to the receiver
instance. Valid input types are as follows:

  - Literal uint consts ([AbstractKind], [AuxiliaryKind], [StructuralKind])
  - int equivalents of uint consts
  - string names of known objectClass kinds

In the case of string names, case is not significant.

This is a fluent method.
*/
func (r *ObjectClass) SetKind(k any) *ObjectClass {
	if r.objectClass == nil {
		r.objectClass = new(objectClass)
	}

	switch tv := k.(type) {
	case string:
		switch lc(tv) {
		case `structural`, ``:
			r.objectClass.Kind = StructuralKind
		case `abstract`:
			r.objectClass.Kind = AbstractKind
		case `auxiliary`:
			r.objectClass.Kind = AuxiliaryKind
		}
	case int:
		if 0 <= tv && tv <= 2 {
			r.objectClass.Kind = uint(tv)
		}
	case uint:
		switch tv {
		case AbstractKind,
			AuxiliaryKind,
			StructuralKind:
			r.objectClass.Kind = tv
		}
	}

	return r
}

/*
SetMust assigns the provided input values as required [AttributeType]
instances advertised through the receiver.

This is a fluent method.
*/
func (r *ObjectClass) SetMust(m ...any) *ObjectClass {
	if r.IsZero() {
		r.objectClass = newObjectClass()
	}

	var err error
	for i := 0; i < len(m) && err == nil; i++ {
		var at AttributeType
		switch tv := m[i].(type) {
		case string:
			at = r.schema().AttributeTypes().get(tv)
		case AttributeType:
			at = tv
		default:
			err = errorf("Unsupported required attributeType %T", tv)
		}

		if err == nil && !at.IsZero() {
			r.objectClass.Must.Push(at)
		}
	}

	return r
}

/*
SetMay assigns the provided input values as permitted [AttributeType]
instances advertised through the receiver.

This is a fluent method.
*/
func (r *ObjectClass) SetMay(m ...any) *ObjectClass {
	if r.IsZero() {
		r.objectClass = newObjectClass()
	}

	var err error
	for i := 0; i < len(m) && err == nil; i++ {
		var at AttributeType
		switch tv := m[i].(type) {
		case string:
			at = r.schema().AttributeTypes().get(tv)
		case AttributeType:
			at = tv
		default:
			err = errorf("Unsupported permitted attributeType %T", tv)
		}

		if err == nil && !at.IsZero() {
			r.objectClass.May.Push(at)
		}
	}

	return r
}

/*
SetSuperClass appends the input value(s) to the the super classes stack within the
the receiver.  Valid input types are string, to represent an RFC 4512 OID residing
in the underlying [Schema] instance, or an bonafide [ObjectClass] instance already
obtained or crafted.

This is a fluent method.
*/
func (r *ObjectClass) SetSuperClass(x ...any) *ObjectClass {
	if r.IsZero() {
		r.objectClass = newObjectClass()
	}

	var err error
	for i := 0; i < len(x) && err == nil; i++ {
		var sup ObjectClass
		switch tv := x[i].(type) {
		case string:
			sup = r.schema().ObjectClasses().get(tv)
		case ObjectClass:
			sup = tv
		default:
			err = errorf("Unsupported super class type %T", tv)
		}

		err = r.objectClass.verifySuperClass(sup.objectClass)
		if err == nil && !sup.IsZero() {
			r.objectClass.SuperClasses.push(sup)
		}
	}

	return r
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

/*
IsObsolete returns a Boolean value indicative of definition
obsolescence.
*/
func (r ObjectClass) IsObsolete() (o bool) {
	if !r.IsZero() {
		o = r.objectClass.Obsolete
	}

	return
}

/*
SetObsolete sets the receiver instance to OBSOLETE if not already set.

Obsolescence cannot be unset.

This is a fluent method.
*/
func (r *ObjectClass) SetObsolete() *ObjectClass {
	if !r.IsZero() {
		if !r.IsObsolete() {
			r.objectClass.Obsolete = true
		}
	}

	return r
}

/*
Kind returns the uint-based kind value associated with the
receiver instance.
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
Inventory returns an instance of [Inventory] which represents the current
inventory of [ObjectClass] instances within the receiver.

The keys are numeric OIDs, while the values are zero (0) or more string
slices, each representing a name by which the definition is known.
*/
func (r ObjectClasses) Inventory() (inv Inventory) {
	inv = make(Inventory, 0)
	for i := 0; i < r.len(); i++ {
		def := r.index(i)
		inv[def.NumericOID()] = def.Names().List()

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
func (r *ObjectClass) SetSchema(schema Schema) *ObjectClass {
	if r.IsZero() {
		r.objectClass = newObjectClass()
	}

	r.objectClass.schema = schema

	return r
}

func (r ObjectClass) schema() (s Schema) {
	if !r.IsZero() {
		s = r.objectClass.schema
	}

	return
}

/*
Commit finalizes the manual build phase of the receiver instance. When executed,
this will allow proper string representation of the receiver instance.

This method need not be executed on parsed [Definition] instances, and is intended
solely for [Definition] instances built manually using the various Set<...> methods
available for instances of this type.

This is a fluent method.
*/
func (r ObjectClass) Commit() ObjectClass {
	if !r.IsZero() {
		if len(r.objectClass.s) == 0 {
			r.objectClass.prepareString()
		}
	}

	return r
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
Names returns the underlying instance of [Name] from within
the receiver.
*/
func (r ObjectClass) Names() (names Name) {
	return r.objectClass.Name
}

/*
Attributes returns an instance of [AttributeTypes] containing references
to [AttributeType] instances which reside within the receiver's MUST and MAY
clauses.

This is useful if a complete list of all immediate [AttributeType] instances
which are available for use and does not include those made available through
super classes.

Duplicate references are silently discarded.
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
AllAttributes returns an [AttributeTypes] instance containing zero (0) or
more permitted or required [AttributeType] definitions for use with this
class as well as those specified by any and all applicable super classes.

This is useful if a complete list of all [AttributeType] instances which
are available for use.

Duplicate references are silently discarded.
*/
func (r ObjectClass) AllAttributes() (all AttributeTypes) {
	all = NewAttributeTypeOIDList()

	for _, attrs := range []AttributeTypes{
		r.AllMay(),
		r.AllMust(),
	} {
		for i := 0; i < attrs.len(); i++ {
			all.push(attrs.index(i))
		}
	}

	return
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
func (r *ObjectClass) SetNumericOID(id string) *ObjectClass {
	if r.objectClass == nil {
		r.objectClass = newObjectClass()
	}

	if isNumericOID(id) {
		if len(r.objectClass.OID) == 0 {
			r.objectClass.OID = id
		}
	}

	return r
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
SetStringer allows the assignment of an individual "stringer" function
or method to the receiver instance.

A non-nil value will be executed for every call of the String method
for the receiver instance.

Should the input stringer value be nil, the [text/template.Template]
value will be used automatically going forward.

This is a fluent method.
*/
func (r *ObjectClass) SetStringer(stringer func() string) *ObjectClass {
	if r.IsZero() {
		r.objectClass = newObjectClass()
	}

	r.objectClass.stringer = stringer

	return r
}

/*
String is a stringer method that returns the string representation
of the receiver instance.
*/
func (r ObjectClass) String() (oc string) {
	if !r.IsZero() {
		if r.objectClass.stringer != nil {
			oc = r.objectClass.stringer()
		} else {
			oc = r.objectClass.s
		}
	}

	return
}

/*
ObjectClasses returns the [ObjectClasses] instance from within
the receiver instance.
*/
func (r Schema) ObjectClasses() (ocs ObjectClasses) {
	slice, _ := r.cast().Index(objectClassesIndex)
	ocs, _ = slice.(ObjectClasses)
	return
}

func (r *objectClass) prepareString() (err error) {
	buf := newBuf()
	var kind string = `STRUCTURAL`
	switch r.Kind {
	case AbstractKind:
		kind = `ABSTRACT`
	case AuxiliaryKind:
		kind = `AUXILIARY`
	}

	r.t = newTemplate(`objectClass`).
		Funcs(funcMap(map[string]any{
			`Kind`:         func() string { return kind },
			`ExtensionSet`: r.Extensions.tmplFunc,
			`MustLen`:      r.Must.len,
			`MayLen`:       r.May.len,
			`SuperLen`:     r.SuperClasses.len,
			`IsObsolete`:   func() bool { return r.Obsolete },
		}))
	if r.t, err = r.t.Parse(objectClassTmpl); err == nil {
		if err = r.t.Execute(buf, struct {
			Definition *objectClass
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
*/                                                                      
func (r *ObjectClass) SetDescription(desc string) *ObjectClass {    
        if len(desc) < 3 {                                              
                return r                                                
        }                                                               
                                                                        
        if r.objectClass == nil {                                     
                r.objectClass = newObjectClass()
        }                                                               
                                                                        
        if !(rune(desc[0]) == rune(39) && rune(desc[len(desc)-1]) == rune(39)) {
                if !r.IsZero() {                                        
                        r.objectClass.Desc = desc                     
                }                                                       
        }                                                               
                                                                        
        return r                                                        
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
	} else if !r.cast().IsInit() {
		err = errorf("Uninitialized %T; cannot push", r)
		return
	}

	// scan receiver for push candidate
	for i := 0; i < len(x); i++ {
		instance := x[i]
		oc, ok := instance.(ObjectClass)
		if !ok {
			err = errorf("Type assertion for %T has failed", oc)
			break
		} else if oc.IsZero() {
			err = errorf("%T instance cannot be pushed; is zero", oc)
			break
		}

		if r.contains(oc.objectClass.OID) {
			err = errorf("%T not unique", oc)
			break
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
	if oc == nil {
		err = errorf("%T instance is nil; cannot append to %T", oc, r)
		return
	}

	r.cast().Push(oc)

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

/*
IsZero returns a Boolean value indicative of nilness of the
receiver instance.
*/
func (r ObjectClass) IsZero() bool {
	return r.objectClass == nil
}

/*
LoadObjectClasses returns an error following an attempt to load all
package-included [ObjectClass] slices into the receiver instance.
*/
func (r Schema) LoadObjectClasses() Schema {
	_ = r.loadObjectClasses()
	return r
}

func (r Schema) loadObjectClasses() (err error) {
	if !r.IsZero() {
		for _, funk := range []func() error{
			r.loadRFC4512ObjectClasses,
			r.loadRFC4519ObjectClasses,
			r.loadRFC4523ObjectClasses,
			r.loadRFC4524ObjectClasses,
			r.loadRFC2307ObjectClasses,
			r.loadRFC2079ObjectClasses,
			r.loadRFC2798ObjectClasses,
			r.loadRFC3671ObjectClasses,
			r.loadRFC3672ObjectClasses,
		} {
			if err = funk(); err != nil {
				break
			}
		}
	}

	return
}

/*
LoadRFC2079ObjectClasses returns an error following an attempt to
load all RFC 2079 [ObjectClass] slices into the receiver instance.
*/
func (r Schema) LoadRFC2079ObjectClasses() Schema {
	_ = r.loadRFC2079ObjectClasses()
	return r
}

func (r Schema) loadRFC2079ObjectClasses() (err error) {
	for i := 0; i < len(rfc2079ObjectClasses) && err == nil; i++ {
		oc := rfc2079ObjectClasses[i]
		err = r.ParseObjectClass(string(oc))
	}

	return
}

/*
LoadRFC2798ObjectClasses returns an error following an attempt to
load all RFC 2798 [ObjectClass] slices into the receiver instance.
*/
func (r Schema) LoadRFC2798ObjectClasses() Schema {
	_ = r.loadRFC2798ObjectClasses()
	return r
}

func (r Schema) loadRFC2798ObjectClasses() (err error) {
	for i := 0; i < len(rfc2798ObjectClasses) && err == nil; i++ {
		oc := rfc2798ObjectClasses[i]
		err = r.ParseObjectClass(string(oc))
	}

	return
}

/*
LoadRFC2307ObjectClasses returns an error following an attempt to
load all RFC 2307 [ObjectClass] slices into the receiver instance.
*/
func (r Schema) LoadRFC2307ObjectClasses() Schema {
	_ = r.loadRFC2307ObjectClasses()
	return r
}

func (r Schema) loadRFC2307ObjectClasses() (err error) {
	for k, v := range rfc2307Macros {
		r.SetMacro(k, v)
	}

	for i := 0; i < len(rfc2307ObjectClasses) && err == nil; i++ {
		oc := rfc2307ObjectClasses[i]
		err = r.ParseObjectClass(string(oc))
	}

	return
}

/*
LoadRFC3671ObjectClasses returns an error following an attempt to
load all RFC 3671 [ObjectClass] slices into the receiver instance.
*/
func (r Schema) LoadRFC3671ObjectClasses() Schema {
	_ = r.loadRFC3671ObjectClasses()
	return r
}

func (r Schema) loadRFC3671ObjectClasses() (err error) {
	for i := 0; i < len(rfc3671ObjectClasses) && err == nil; i++ {
		oc := rfc3671ObjectClasses[i]
		err = r.ParseObjectClass(string(oc))
	}

	return
}

/*
LoadRFC3672ObjectClasses returns an error following an attempt to
load all RFC 3672 [ObjectClass] slices into the receiver instance.
*/
func (r Schema) LoadRFC3672ObjectClasses() Schema {
	_ = r.loadRFC3672ObjectClasses()
	return r
}

func (r Schema) loadRFC3672ObjectClasses() (err error) {
	for i := 0; i < len(rfc3672ObjectClasses) && err == nil; i++ {
		oc := rfc3672ObjectClasses[i]
		err = r.ParseObjectClass(string(oc))
	}

	return
}

/*
LoadRFC4512ObjectClasses returns an error following an attempt to
load all RFC 4512 [ObjectClass] slices into the receiver instance.
*/
func (r Schema) LoadRFC4512ObjectClasses() Schema {
	_ = r.loadRFC4512ObjectClasses()
	return r
}

/*
LoadRFC4530AttributeTypes returns an error following an attempt to
load all RFC 4530 [ObjectClass] slices into the receiver instance.
*/
func (r Schema) loadRFC4512ObjectClasses() (err error) {
	for i := 0; i < len(rfc4512ObjectClasses) && err == nil; i++ {
		oc := rfc4512ObjectClasses[i]
		err = r.ParseObjectClass(string(oc))
	}

	return
}

/*
LoadRFC4519ObjectClasses returns an error following an attempt to
load all RFC 4519 [ObjectClass] slices into the receiver instance.
*/
func (r Schema) LoadRFC4519ObjectClasses() Schema {
	_ = r.loadRFC4519ObjectClasses()
	return r
}

func (r Schema) loadRFC4519ObjectClasses() (err error) {
	for i := 0; i < len(rfc4519ObjectClasses) && err == nil; i++ {
		oc := rfc4519ObjectClasses[i]
		err = r.ParseObjectClass(string(oc))
	}

	return
}

/*
LoadRFC4523ObjectClasses returns an error following an attempt to
load all RFC 4523 [ObjectClass] slices into the receiver instance.
*/
func (r Schema) LoadRFC4523ObjectClasses() Schema {
	_ = r.loadRFC4523ObjectClasses()
	return r
}

func (r Schema) loadRFC4523ObjectClasses() (err error) {
	for i := 0; i < len(rfc4523ObjectClasses) && err == nil; i++ {
		oc := rfc4523ObjectClasses[i]
		err = r.ParseObjectClass(string(oc))
	}

	return
}

/*
LoadRFC4524ObjectClasses returns an error following an attempt to
load all RFC 4524 [ObjectClass] slices into the receiver instance.
*/
func (r Schema) LoadRFC4524ObjectClasses() Schema {
	_ = r.loadRFC4524ObjectClasses()
	return r
}

func (r Schema) loadRFC4524ObjectClasses() (err error) {
	for i := 0; i < len(rfc4524ObjectClasses) && err == nil; i++ {
		oc := rfc4524ObjectClasses[i]
		err = r.ParseObjectClass(string(oc))
	}

	return
}

package schemax

/*
NewAttributeType initializes and returns a new instance of
[AttributeType], ready for manual population.

This function need not be used when parsing is engaged.
*/
func NewAttributeType() AttributeType {
	return AttributeType{newAttributeType()}
}

/*
newAttributeType initializes a new instance of the private (embedded)
type *attributeType.
*/
func newAttributeType() *attributeType {
	return &attributeType{
		Name:       NewName(),
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
func (r *AttributeType) SetSchema(schema Schema) *AttributeType {
	if r.IsZero() {
		r.attributeType = newAttributeType()
	}

	r.attributeType.schema = schema

	return r
}

func (r AttributeType) schema() (s Schema) {
	if !r.IsZero() {
		s = r.attributeType.schema
	}

	return
}

/*
Inventory returns an instance of [Inventory] which represents the current
inventory of [AttributeType] instances within the receiver.

The keys are numeric OIDs, while the values are zero (0) or more string
slices, each representing a name by which the definition is known.
*/
func (r AttributeTypes) Inventory() (inv Inventory) {
	inv = make(Inventory, 0)
	for i := 0; i < r.len(); i++ {
		def := r.index(i)
		inv[def.NumericOID()] = def.Names().List()

	}

	return
}

// stackage closure func - do not exec directly (use String method)
func (r AttributeTypes) oIDsStringer(_ ...any) (present string) {
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
		present = `(` + padchar + joined + padchar + `)`
	}

	return
}

// stackage closure func - do not exec directly.
func (r AttributeTypes) canPush(x ...any) (err error) {
	if len(x) == 0 {
		return
	}

	for i := 0; i < len(x) && err == nil; i++ {
		instance := x[i]
		if at, ok := instance.(AttributeType); !ok || at.IsZero() {
			err = errorf("Type assertion for %T has failed", instance)
		} else {
			if tst := r.get(at.NumericOID()); !tst.IsZero() {
				err = errorf("%T %s not unique", at, at.NumericOID())
			}
		}
	}

	return
}

/*
AttributeTypes returns the [AttributeTypes] instance from within
the receiver instance.
*/
func (r Schema) AttributeTypes() (ats AttributeTypes) {
	slice, _ := r.cast().Index(attributeTypesIndex)
	ats, _ = slice.(AttributeTypes)
	return
}

/*
Len returns the current integer length of the receiver instance.
*/
func (r AttributeTypes) Len() int {
	return r.len()
}

func (r AttributeTypes) len() int {
	return r.cast().Len()
}

/*
String is a stringer method that returns the string representation
of the receiver instance.
*/
func (r AttributeTypes) String() string {
	return r.cast().String()
}

/*
Index returns the instance of [AttributeType] found within the
receiver stack instance at index N.  If no instance is found at
the index specified, a zero [AttributeType] instance is returned.
*/
func (r AttributeTypes) Index(idx int) AttributeType {
	return r.index(idx)
}

func (r AttributeTypes) index(idx int) (at AttributeType) {
	slice, found := r.cast().Index(idx)
	if found {
		if _at, ok := slice.(AttributeType); ok {
			at = _at
		}
	}

	return
}

/*
Push returns an error following an attempt to push an AttributeType
into the receiver stack instance.
*/
func (r AttributeTypes) Push(at any) error {
	return r.push(at)
}

func (r AttributeTypes) push(at any) (err error) {
	if at == nil {
		err = errorf("%T instance is nil; cannot append to %T", at, r)
		return
	}

	r.cast().Push(at)

	return
}

/*
Contains calls [AttributeTypes.Get] to return a Boolean value indicative
of a successful, non-zero retrieval of an [AttributeType] instance -- matching
the provided id -- from within the receiver stack instance.
*/
func (r AttributeTypes) Contains(id string) bool {
	return r.contains(id)
}

func (r AttributeTypes) contains(id string) bool {
	return !r.get(id).IsZero()
}

/*
Get returns an instance of [AttributeType] based upon a search for id within
the receiver stack instance.

The return instance, if not nil, was retrieved based upon a textual match of
the principal identifier of an [AttributeType] and the provided id.

The return instance is nil if no match was made.

Case is not significant in the matching process.
*/
func (r AttributeTypes) Get(id string) AttributeType {
	return r.get(id)
}

func (r AttributeTypes) get(id string) (at AttributeType) {
	for i := 0; i < r.len() && at.IsZero(); i++ {
		if _at := r.index(i); !_at.IsZero() {
			if _at.attributeType.OID == id {
				at = _at
			} else if _at.attributeType.Name.contains(id) {
				at = _at
			}
		}
	}

	return
}

/*
IsZero returns a Boolean value indicative of nilness of the
receiver instance.
*/
func (r AttributeType) IsZero() bool {
	return r.attributeType == nil
}

/*
IsIdentifiedAs returns a Boolean value indicative of whether id matches
either the numericOID or descriptor of the receiver instance.  Case is
not significant in the matching process.
*/
func (r AttributeType) IsIdentifiedAs(id string) (ident bool) {
	if !r.IsZero() {
		ident = id == r.NumericOID() || r.attributeType.Name.contains(id)
	}

	return
}

/*
IsImmutable returns a Boolean value indicative of whether the receiver
instance has its NO-USER-MODIFICATIONS option enabled. As such, only a
DSA may manage values of this type when a value of true is in effect.
*/
func (r AttributeType) IsImmutable() (o bool) {
	if !r.IsZero() {
		o = r.attributeType.Immutable
	}

	return
}

/*
IsCollective returns a Boolean value indicative of whether the receiver
is COLLECTIVE.  A value of true is mutually exclusive of SINGLE-VALUE'd
[AttributeType] instances.
*/
func (r AttributeType) IsCollective() (o bool) {
	if !r.IsZero() {
		o = r.attributeType.Collective
	}

	return
}

/*
IsSingleValued returns a Boolean value indicative of whether the receiver
is set to only allow one (1) value to be assigned to an entry using this
type.  A value of true is mutually exclusive of COLLECTIVE [AttributeType]
instances.
*/
func (r AttributeType) IsSingleValued() (o bool) {
	if !r.IsZero() {
		o = r.attributeType.Single
	}

	return
}

/*
IsObsolete returns a Boolean value indicative of definition obsolescence.
*/
func (r AttributeType) IsObsolete() (o bool) {
	if !r.IsZero() {
		o = r.attributeType.Obsolete
	}

	return
}

/*
Name returns the string form of the principal name of the receiver instance, if set.
*/
func (r AttributeType) Name() (id string) {
	if !r.IsZero() {
		id = r.attributeType.Name.index(0)
	}

	return
}

/*
Names returns the underlying instance of [Name] from within
the receiver.
*/
func (r AttributeType) Names() (names Name) {
	return r.attributeType.Name
}

/*
NewAttributeTypes initializes and returns a new [AttributeTypes] instance,
configured to allow the storage of all [AttributeType] instances.
*/
func NewAttributeTypes() AttributeTypes {
	r := AttributeTypes(newCollection(`attributeTypes`))
	r.cast().SetPushPolicy(r.canPush)

	return r
}

/*
NewAttributeTypeOIDList initializes and returns a new [AttributeTypes] that has
been cast from an instance of [OIDList] and configured to allow the storage of
arbitrary [AttributeType] instances.
*/
func NewAttributeTypeOIDList() AttributeTypes {
	r := AttributeTypes(newOIDList(`at_oidlist`))
	r.cast().
		SetPushPolicy(r.canPush).
		SetPresentationPolicy(r.oIDsStringer)

	return r
}

/*
OID returns the string representation of an OID -- which is either a
numeric OID or descriptor -- that is held by the receiver instance.
*/
func (r AttributeType) OID() (oid string) {
	if !r.IsZero() {
		oid = r.NumericOID() // default
		if r.attributeType.Name.len() > 0 {
			oid = r.attributeType.Name.index(0)
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
func (r *AttributeType) SetNumericOID(id string) *AttributeType {
	if r.IsZero() {
		r.attributeType = newAttributeType()
	}

	if isNumericOID(id) {
		if len(r.attributeType.OID) == 0 {
			r.attributeType.OID = id
		}
	}

	return r
}

/*
SetExtension assigns key x to value xstrs within the receiver's underlying
[Extensions] instance.

This is a fluent method.
*/
func (r *AttributeType) SetExtension(x string, xstrs ...string) *AttributeType {
	if r.IsZero() {
		r.attributeType = newAttributeType()
	}

	r.Extensions().Set(x, xstrs...)

	return r
}

/*
Extensions returns the [Extensions] instance -- if set -- within
the receiver.
*/
func (r AttributeType) Extensions() (e Extensions) {
	if !r.IsZero() {
		e = r.attributeType.Extensions
	}

	return
}

/*
SetName assigns the provided names to the receiver instance.

Name instances must conform to RFC 4512 descriptor format but
need not be quoted.
*/
func (r *AttributeType) SetName(x ...string) *AttributeType {
	if len(x) == 0 {
		return r
	}

	if r.IsZero() {
		r.attributeType = newAttributeType()
	}

	for i := 0; i < len(x); i++ {
		r.attributeType.Name.Push(x[i])
	}

	return r
}

/*
NumericOID returns the string representation of the numeric OID
held by the receiver instance.
*/
func (r AttributeType) NumericOID() (noid string) {
	if !r.IsZero() {
		noid = r.attributeType.OID
	}

	return
}

/*
Description returns the underlying (optional) descriptive text
assigned to the receiver instance.
*/
func (r AttributeType) Description() (desc string) {
	if !r.IsZero() {
		desc = r.attributeType.Desc
	}

	return
}

/*
SetDescription parses desc into the underlying Desc field within the
receiver instance.  Although a RFC 4512-compliant QuotedString is
required, the outer single-quotes need not be specified literally.
*/
func (r *AttributeType) SetDescription(desc string) *AttributeType {
	if len(desc) < 3 {
		return r
	}

	if r.attributeType == nil {
		r.attributeType = new(attributeType)
	}

	if !(rune(desc[0]) == rune(39) && rune(desc[len(desc)-1]) == rune(39)) {
		if !r.IsZero() {
			r.attributeType.Desc = desc
		}
	}

	return r
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
func (r *AttributeType) SetStringer(stringer func() string) *AttributeType {
	if r.IsZero() {
		r.attributeType = newAttributeType()
	}

	r.attributeType.stringer = stringer

	return r
}

func (r *attributeType) prepareString() (err error) {
	buf := newBuf()
	r.t = newTemplate(`attributeType`).
		Funcs(funcMap(map[string]any{
			`Substring`:    func() string { return r.Substring.OID() },
			`Ordering`:     func() string { return r.Ordering.OID() },
			`Equality`:     func() string { return r.Equality.OID() },
			`Syntax`:       func() string { return r.Syntax.NumericOID() },
			`SuperType`:    func() string { return r.SuperType.OID() },
			`ExtensionSet`: r.Extensions.tmplFunc,
			`IsObsolete`:   func() bool { return r.Obsolete },
			`IsSingleVal`:  func() bool { return r.Single },
			`IsCollective`: func() bool { return r.Collective },
			`IsNoUserMod`:  func() bool { return r.Immutable },
			`Usage`:        func() string { return AttributeType{r}.Usage() },
		}))

	if r.t, err = r.t.Parse(attributeTypeTmpl); err == nil {
		if err = r.t.Execute(buf, struct {
			Definition *attributeType
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
String is a stringer method that returns the string representation
of the receiver instance.
*/
func (r AttributeType) String() (at string) {
	if !r.IsZero() {
		if r.stringer != nil {
			at = r.stringer()
		} else {
			if len(r.attributeType.s) == 0 {
				var err error
				if err = r.attributeType.prepareString(); err != nil {
					return
				}
			}

			at = r.attributeType.s
		}
	}

	return
}

/*
Syntax returns the string representation of the [LDAPSyntax] numeric if
set within the receiver instance. If unset, a zero string is returned.
*/
func (r AttributeType) Syntax() (syn string) {
	if !r.IsZero() {
		if !r.attributeType.Syntax.IsZero() {
			syn = r.attributeType.Syntax.NumericOID()
		}
	}

	return
}

/*
SetSyntax assigns x to the receiver instance as an instance of [LDAPSyntax].

This is a fluent method.
*/
func (r *AttributeType) SetSyntax(x any) *AttributeType {
	if r.IsZero() {
		r.attributeType = newAttributeType()
	}

	var syn LDAPSyntax
	switch tv := x.(type) {
	case string:
		if sch := r.schema(); !sch.IsZero() {
			syn = sch.LDAPSyntaxes().get(tv)
		}
	case LDAPSyntax:
		syn = tv
	}

	if !syn.IsZero() {
		r.attributeType.Syntax = syn
	}

	return r
}

/*
Equality returns the underlying instance of [MatchingRule] if set within
the receiver instance. If unset, a zero instance is returned.
*/
func (r AttributeType) Equality() (eql MatchingRule) {
	if !r.IsZero() {
		eql = r.attributeType.Equality
	}

	return
}

/*
SetEquality sets the equality [MatchingRule] of the receiver instance.

Input x may be the string representation of the desired equality matchingRule,
by name or numeric OID.  This requires an underlying instance of [Schema] set
within the receiver instance. The [Schema] instance is expected to possess
the referenced [MatchingRule], else a zero instance will ultimately be returned
and discarded.

Input x may also be a bonafide, non-zero instance of an equality [MatchingRule].
This requires no underlying [Schema] instance, but may allow the introduction of
bogus [MatchingRule] instances.

This is a fluent method.
*/
func (r *AttributeType) SetEquality(x any) *AttributeType {
	if r.IsZero() {
		r.attributeType = newAttributeType()
	}

	var mr MatchingRule
	switch tv := x.(type) {
	case string:
		if sch := r.schema(); !sch.IsZero() {
			mr = sch.MatchingRules().get(tv)
		}
	case MatchingRule:
		mr = tv
	}

	if !mr.IsZero() {
		r.attributeType.Equality = mr
	}

	return r
}

/*
Substring returns the underlying instance of [MatchingRule] if set within
the receiver instance. If unset, a zero instance is returned.
*/
func (r AttributeType) Substring() (sub MatchingRule) {
	if !r.IsZero() {
		sub = r.attributeType.Substring
	}

	return
}

/*
SetSubstring sets the substring [MatchingRule] of the receiver instance.

Input x may be the string representation of the desired substring matchingRule,
by name or numeric OID.  This requires an underlying instance of [Schema] set
within the receiver instance. The [Schema] instance is expected to possess
the referenced [MatchingRule], else a zero instance will ultimately be returned
and discarded.

Input x may also be a bonafide, non-zero instance of a substring [MatchingRule].
This requires no underlying [Schema] instance, but may allow the introduction of
bogus [MatchingRule] instances.

This is a fluent method.
*/
func (r *AttributeType) SetSubstring(x any) *AttributeType {
	if r.IsZero() {
		r.attributeType = newAttributeType()
	}

	var mr MatchingRule
	switch tv := x.(type) {
	case string:
		if sch := r.schema(); !sch.IsZero() {
			mr = sch.MatchingRules().get(tv)
		}
	case MatchingRule:
		mr = tv
	}

	if !mr.IsZero() {
		r.attributeType.Substring = mr
	}

	return r
}

/*
Ordering returns the underlying instance of [MatchingRule] if set within
the receiver instance. If unset, a zero instance is returned.
*/
func (r AttributeType) Ordering() (ord MatchingRule) {
	if !r.IsZero() {
		ord = r.attributeType.Ordering
	}

	return
}

/*
SetOrdering sets the ordering [MatchingRule] of the receiver instance.

Input x may be the string representation of the desired ordering matchingRule,
by name or numeric OID.  This requires an underlying instance of [Schema] set
within the receiver instance. The [Schema] instance is expected to possess
the referenced [MatchingRule], else a zero instance will ultimately be returned
and discarded.

Input x may also be a bonafide, non-zero instance of an ordering [MatchingRule].
This requires no underlying [Schema] instance, but may allow the introduction of
bogus [MatchingRule] instances.

This is a fluent method.
*/
func (r *AttributeType) SetOrdering(x any) *AttributeType {
	if r.IsZero() {
		return r
	}

	var mr MatchingRule
	switch tv := x.(type) {
	case string:
		if sch := r.schema(); !sch.IsZero() {
			mr = sch.MatchingRules().get(tv)
		}
	case MatchingRule:
		mr = tv
	}

	if !mr.IsZero() {
		r.attributeType.Ordering = mr
	}

	return r
}

/*
SuperType returns the underlying instance of [AttributeType] if set within
the receiver instance as its super type. If unset, a zero instance is returned.
*/
func (r AttributeType) SuperType() (sup AttributeType) {
	if !r.IsZero() {
		sup = r.attributeType.SuperType
	}

	return
}

/*
SetSuperType sets the super type of the receiver to the value provided. Valid
input types are string, to represent an RFC 4512 OID residing in the underlying
[Schema] instance, or an actual [AttributeType] instance already obtained or crafted.

This is a fluent method.
*/
func (r *AttributeType) SetSuperType(x any) *AttributeType {
	if r.IsZero() {
		r.attributeType = newAttributeType()
	}

	var sup AttributeType
	switch tv := x.(type) {
	case string:
		if !r.schema().IsZero() {
			sup = r.schema().AttributeTypes().get(tv)
		}
	case AttributeType:
		sup = tv
	}

	var err error
	//if !r.schema().IsZero() {
	//        err = r.attributeType.verifySuperType(sup.attributeType)
	//}

	if err == nil && !sup.IsZero() {
		r.attributeType.SuperType = sup
	}

	return r
}

/*
SetSingleValue assigns the input value to the underlying SingleValue field
within the receiver.

Input types may be bool, or string representations of bool. When strings
are used, case is not significant.

Note that a value of true will be ignored if the receiver is a collective
[AttributeType].

This is a fluent method.
*/
func (r *AttributeType) SetSingleValue(x any) *AttributeType {
	r.setBoolean(`sv`, x)
	return r
}

/*
SetCollective assigns the input value to the underlying Collective field
within the receiver.

Input types may be bool, or string representations of bool. When strings
are used, case is not significant.

Note that a value of true will be ignored if the receiver is a single-valued
[AttributeType].

This is a fluent method.
*/
func (r *AttributeType) SetCollective(x any) *AttributeType {
	r.setBoolean(`c`, x)
	return r
}

/*
SetImmutable assigns the input value to the underlying Immutable field
within the receiver.

Input types may be bool, or string representations of bool. When strings
are used, case is not significant.

This is a fluent method.
*/
func (r *AttributeType) SetImmutable(x any) *AttributeType {
	r.setBoolean(`num`, x)
	return r
}

/*
SetObsolete assigns the input value to the underlying Obsolete field within
the receiver.

Input types may be bool, or string representations of bool. When strings
are used, case is not significant.

Obsolescence cannot be unset.

This is a fluent method.
*/
func (r *AttributeType) SetObsolete(x any) *AttributeType {
	r.setBoolean(`obs`, x)
	return r
}

func (r *AttributeType) setBoolean(t string, x any) *AttributeType {

	var Bool bool
	switch tv := x.(type) {
	case string:
		if eq(tv, `true`) {
			Bool = true
		}
	case bool:
		Bool = tv
	default:
		return r
	}

	switch t {
	case `sv`:
		if !r.attributeType.Collective {
			r.attributeType.Single = Bool
		}
	case `c`:
		if !r.attributeType.Single {
			r.attributeType.Collective = Bool
		}
	case `num`:
		if !r.attributeType.Immutable {
			r.attributeType.Immutable = Bool
		}
	case `obs`:
		if !r.attributeType.Obsolete {
			r.attributeType.Obsolete = Bool
		}
	}

	return r
}

/*
SetUsage assigns the specified USAGE to the receiver. Input types
may be string, int or uint.

	1, or directoryOperation
	2, or distributedOperation
	3, or dSAOperation

Any other value results in assignment of the userApplications USAGE.

This is a fluent method.
*/
func (r *AttributeType) SetUsage(u any) *AttributeType {
	if !r.IsZero() {
		switch tv := u.(type) {
		case string:
			switch lc(tv) {
			case `directoryoperation`:
				r.attributeType.Usage = DirectoryOperationUsage
			case `distributedoperation`:
				r.attributeType.Usage = DistributedOperationUsage
			case `dsaoperation`:
				r.attributeType.Usage = DSAOperationUsage
			default:
				r.attributeType.Usage = UserApplicationsUsage
			}
		case uint:
			r.SetUsage(int(tv))
		case int:
			switch tv {
			case 1:
				r.attributeType.Usage = DirectoryOperationUsage
			case 2:
				r.attributeType.Usage = DistributedOperationUsage
			case 3:
				r.attributeType.Usage = DSAOperationUsage
			default:
				r.attributeType.Usage = UserApplicationsUsage
			}
		}
	}

	return r
}

/*
Type returns the string literal "attributeType".
*/
func (r AttributeType) Type() string {
	return `attributeType`
}

/*
Type returns the string literal "attributeTypes".
*/
func (r AttributeTypes) Type() string {
	return `attributeTypes`
}

/*
IsZero returns a Boolean value indicative of nilness of the
receiver instance.
*/
func (r AttributeTypes) IsZero() bool {
	return r.cast().IsZero()
}

/*
Usage returns the string representation of the underlying USAGE if set
within the receiver instance. If unset, a zero string -- which implies
use of the "userApplication" [AttributeType] USAGE value by default --
is returned.
*/
func (r AttributeType) Usage() (usage string) {
	if !r.IsZero() {
		switch v := r.attributeType.Usage; int(v) {
		case 0:
			break // zero is default (userApplication)
		case 1:
			usage = `directoryOperation`
		case 2:
			usage = `distributedOperation`
		case 3:
			usage = `dSAOperation`
		}
	}

	return
}

/*
Maps returns slices of [DefinitionMap] instances.
*/
func (r AttributeTypes) Maps() (defs DefinitionMaps) {
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
func (r AttributeType) Map() (def DefinitionMap) {
	if r.IsZero() {
		return
	}

	def = make(DefinitionMap, 0)
	def[`NUMERICOID`] = []string{r.NumericOID()}
	def[`NAME`] = r.Names().List()
	def[`DESC`] = []string{r.Description()}
	def[`OBSOLETE`] = []string{bool2str(r.IsObsolete())}
	def[`SUP`] = []string{r.SuperType().OID()}
	def[`EQUALITY`] = []string{r.Equality().OID()}
	def[`SUBSTR`] = []string{r.Substring().OID()}
	def[`ORDERING`] = []string{r.Ordering().OID()}
	def[`SYNTAX`] = []string{r.Syntax()}
	def[`SINGLE-VALUE`] = []string{bool2str(r.IsSingleValued())}
	def[`COLLECTIVE`] = []string{bool2str(r.IsCollective())}
	def[`NO-USER-MODIFICATION`] = []string{bool2str(r.IsImmutable())}
	def[`USAGE`] = []string{r.Usage()}
	def[`TYPE`] = []string{r.Type()}
	def[`RAW`] = []string{r.String()}

	// copy our extensions from receiver r
	// into destination def.
	def = mapTransferExtensions(r, def)

	// Clean up any empty fields
	def.clean()

	return
}

func (r AttributeType) macro() (m []string) {
	if !r.IsZero() {
		m = r.attributeType.Macro
	}

	return
}

func (r AttributeType) setOID(x string) {
	if !r.IsZero() {
		r.attributeType.OID = x
	}
}

/*
LoadAttributeTypes returns an error following an attempt to load all
package-included [AttributeType] slices into the receiver instance.
*/
func (r Schema) LoadAttributeTypes() Schema {
	_ = r.loadAttributeTypes()
	return r
}

/*
loadAttributeTypes returns an error following an attempt to load
all package-included [AttributeType] slices into the receiver instance.
*/
func (r Schema) loadAttributeTypes() (err error) {
	if !r.IsZero() {
		for _, funk := range []func() error{
			r.loadRFC4512AttributeTypes,
			r.loadRFC2079AttributeTypes,
			r.loadRFC2798AttributeTypes,
			r.loadRFC3045AttributeTypes,
			r.loadRFC3672AttributeTypes,
			r.loadRFC4519AttributeTypes,
			r.loadRFC2307AttributeTypes,
			r.loadRFC3671AttributeTypes,
			r.loadRFC4523AttributeTypes,
			r.loadRFC4524AttributeTypes,
			r.loadRFC4530AttributeTypes,
		} {
			if err = funk(); err != nil {
				break
			}
		}
	}

	return
}

/*
LoadRFC2079AttributeTypes returns an error following an attempt to
load all RFC 2079 [AttributeType] slices into the receiver instance.
*/
func (r Schema) LoadRFC2079AttributeTypes() Schema {
	_ = r.loadRFC2079AttributeTypes()
	return r
}

func (r Schema) loadRFC2079AttributeTypes() (err error) {
	for i := 0; i < len(rfc2079AttributeTypes) && err == nil; i++ {
		at := rfc2079AttributeTypes[i]
		err = r.ParseAttributeType(string(at))
	}

	return
}

/*
LoadRFC2307AttributeTypes returns an error following an attempt to
load all RFC 2307 [AttributeType] slices into the receiver instance.
*/
func (r Schema) LoadRFC2307AttributeTypes() Schema {
	_ = r.loadRFC2307AttributeTypes()
	return r
}

func (r Schema) loadRFC2307AttributeTypes() (err error) {
	for k, v := range rfc2307Macros {
		r.SetMacro(k, v)
	}

	for i := 0; i < len(rfc2307AttributeTypes) && err == nil; i++ {
		at := rfc2307AttributeTypes[i]
		err = r.ParseAttributeType(string(at))
	}

	return
}

/*
LoadRFC2798AttributeTypes returns an error following an attempt to
load all RFC 2798 [AttributeType] slices into the receiver instance.
*/
func (r Schema) LoadRFC2798AttributeTypes() Schema {
	_ = r.loadRFC2798AttributeTypes()
	return r
}

func (r Schema) loadRFC2798AttributeTypes() (err error) {
	for i := 0; i < len(rfc2798AttributeTypes) && err == nil; i++ {
		at := rfc2798AttributeTypes[i]
		err = r.ParseAttributeType(string(at))
	}

	return
}

/*
LoadRFC3045AttributeTypes returns an error following an attempt to
load all RFC 3045 [AttributeType] slices into the receiver instance.
*/
func (r Schema) LoadRFC3045AttributeTypes() Schema {
	_ = r.loadRFC3045AttributeTypes()
	return r
}

func (r Schema) loadRFC3045AttributeTypes() (err error) {
	for i := 0; i < len(rfc3045AttributeTypes) && err == nil; i++ {
		at := rfc3045AttributeTypes[i]
		err = r.ParseAttributeType(string(at))
	}

	return
}

/*
LoadRFC3671AttributeTypes returns an error following an attempt to
load all RFC 3671 [AttributeType] slices into the receiver instance.
*/
func (r Schema) LoadRFC3671AttributeTypes() Schema {
	_ = r.loadRFC3671AttributeTypes()
	return r
}

func (r Schema) loadRFC3671AttributeTypes() (err error) {
	for i := 0; i < len(rfc3671AttributeTypes) && err == nil; i++ {
		at := rfc3671AttributeTypes[i]
		err = r.ParseAttributeType(string(at))
	}

	return
}

/*
LoadRFC3672AttributeTypes returns an error following an attempt to
load all RFC 3672 [AttributeType] slices into the receiver instance.
*/
func (r Schema) LoadRFC3672AttributeTypes() Schema {
	_ = r.loadRFC3672AttributeTypes()
	return r
}

func (r Schema) loadRFC3672AttributeTypes() (err error) {
	for i := 0; i < len(rfc3672AttributeTypes) && err == nil; i++ {
		at := rfc3672AttributeTypes[i]
		err = r.ParseAttributeType(string(at))
	}

	return
}

/*
LoadRFC4512AttributeTypes returns an error following an attempt to
load all RFC 4512 [AttributeType] slices into the receiver instance.
*/
func (r Schema) LoadRFC4512AttributeTypes() Schema {
	_ = r.loadRFC4512AttributeTypes()
	return r
}

func (r Schema) loadRFC4512AttributeTypes() (err error) {
	for i := 0; i < len(rfc4512AttributeTypes) && err == nil; i++ {
		at := rfc4512AttributeTypes[i]
		err = r.ParseAttributeType(string(at))
	}

	return
}

/*
LoadRFC4519AttributeTypes returns an error following an attempt to
load all RFC 4519 [AttributeType] slices into the receiver instance.
*/
func (r Schema) LoadRFC4519AttributeTypes() Schema {
	_ = r.loadRFC4519AttributeTypes()
	return r
}

func (r Schema) loadRFC4519AttributeTypes() (err error) {
	for i := 0; i < len(rfc4519AttributeTypes) && err == nil; i++ {
		at := rfc4519AttributeTypes[i]
		err = r.ParseAttributeType(string(at))
	}

	return
}

/*
LoadRFC4523AttributeTypes returns an error following an attempt to
load all RFC 4523 [AttributeType] slices into the receiver instance.
*/
func (r Schema) LoadRFC4523AttributeTypes() Schema {
	_ = r.loadRFC4523AttributeTypes()
	return r
}

func (r Schema) loadRFC4523AttributeTypes() (err error) {
	for i := 0; i < len(rfc4523AttributeTypes) && err == nil; i++ {
		at := rfc4523AttributeTypes[i]
		err = r.ParseAttributeType(string(at))
	}

	return
}

/*
LoadRFC4524AttributeTypes returns an error following an attempt to
load all RFC 4524 [AttributeType] slices into the receiver instance.
*/
func (r Schema) LoadRFC4524AttributeTypes() Schema {
	_ = r.loadRFC4524AttributeTypes()
	return r
}

func (r Schema) loadRFC4524AttributeTypes() (err error) {
	for i := 0; i < len(rfc4524AttributeTypes) && err == nil; i++ {
		at := rfc4524AttributeTypes[i]
		err = r.ParseAttributeType(string(at))
	}

	return
}

/*
LoadRFC4530AttributeTypes returns an error following an attempt to
load all RFC 4530 [AttributeType] slices into the receiver instance.
*/
func (r Schema) LoadRFC4530AttributeTypes() Schema {
	_ = r.loadRFC4530AttributeTypes()
	return r
}

func (r Schema) loadRFC4530AttributeTypes() (err error) {
	for i := 0; i < len(rfc4530AttributeTypes) && err == nil; i++ {
		at := rfc4530AttributeTypes[i]
		err = r.ParseAttributeType(string(at))
	}

	return
}

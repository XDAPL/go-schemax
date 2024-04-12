package schemax

import (
	"internal/rfc2079"
	"internal/rfc2798"
	"internal/rfc3045"
	"internal/rfc3671"
	"internal/rfc3672"
	"internal/rfc4512"
	"internal/rfc4519"
	"internal/rfc4523"
	"internal/rfc4524"
	"internal/rfc4530"

	antlr4512 "github.com/JesseCoretta/go-rfc4512-antlr"
)

var (
	rfc2079AttributeTypes rfc2079.AttributeTypeDefinitions = rfc2079.AllAttributeTypes
	rfc2798AttributeTypes rfc2798.AttributeTypeDefinitions = rfc2798.AllAttributeTypes
	rfc3045AttributeTypes rfc3045.AttributeTypeDefinitions = rfc3045.AllAttributeTypes
	rfc3671AttributeTypes rfc3671.AttributeTypeDefinitions = rfc3671.AllAttributeTypes
	rfc3672AttributeTypes rfc3672.AttributeTypeDefinitions = rfc3672.AllAttributeTypes
	rfc4512AttributeTypes rfc4512.AttributeTypeDefinitions = rfc4512.AllAttributeTypes
	rfc4519AttributeTypes rfc4519.AttributeTypeDefinitions = rfc4519.AllAttributeTypes
	rfc4523AttributeTypes rfc4523.AttributeTypeDefinitions = rfc4523.AllAttributeTypes
	rfc4524AttributeTypes rfc4524.AttributeTypeDefinitions = rfc4524.AllAttributeTypes
	rfc4530AttributeTypes rfc4530.AttributeTypeDefinitions = rfc4530.AllAttributeTypes
)

/*
ParseAttributeType parses an individual textual attribute type (raw) and
returns an error instance.

When no error occurs, the newly formed [AttributeType] instance -- based on
the parsed contents of raw -- is added to the receiver [AttributeTypes]
slice instance.
*/
func (r Schema) ParseAttributeType(raw string) error {
	i, err := parseI(raw)
	if err == nil {
		var a AttributeType
		a, err = r.processAttributeType(i.P.AttributeTypeDescription())
		if err == nil {
			err = r.AttributeTypes().push(a)
		}
	}

	return err
}

/*
newAttributeType initializes a new instance of the private (embedded)
type *attributeType.
*/
func newAttributeType() *attributeType {
	return &attributeType{
		Extensions: NewExtensions(),
	}
}

func (r AttributeType) Map() (def map[string][]string) {
	if r.IsZero() {
		return
	}

	def = make(map[string][]string, 0)
	def[`NUMERICOID`]           = []string{r.NumericOID()}
	def[`NAME`]	            = r.Names().List()
	def[`DESC`]	            = []string{r.Description()}
	def[`OBSOLETE`]             = []string{bool2str(r.IsObsolete())}
	def[`SUP`]	            = []string{r.SuperType().OID()}
	def[`EQUALITY`]             = []string{r.Equality().OID()}
	def[`SUBSTR`]	            = []string{r.Substring().OID()}
	def[`ORDERING`]             = []string{r.Ordering().OID()}
	def[`SYNTAX`]	            = []string{r.Syntax()}
	def[`SINGLE-VALUE`]         = []string{bool2str(r.IsSingleValued())}
	def[`COLLECTIVE`]	    = []string{bool2str(r.IsCollective())}
	def[`NO-USER-MODIFICATION`] = []string{bool2str(r.IsImmutable())}
	def[`USAGE`]	            = []string{r.Usage()} // is always numeric OID

	exts := r.Extensions()
	for _, k := range exts.Keys() {
		if ext, found := exts.get(k); found {
			def[k] = ext.List()
		}
	}

	// Clean up any empty fields
	for k, v := range def {
		if len(v) == 0 {
			delete(def, k)
		} else if len(v[0]) == 0 {
			delete(def, k)
		}
	}

	return
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
		o = r.attributeType.NoUserMod
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
		o = r.attributeType.SingleVal
	}

	return
}

/*
SetObsolete will declare the receiver instance as OBSOLETE.  Calls to
this method will not operate in a toggling manner in that there is no
way to "unset" a state of obsolescence.

This is a fluent method.
*/
func (r *AttributeType) SetObsolete() AttributeType {
	if r.attributeType == nil {
		r.attributeType = newAttributeType()
	}

	if !r.IsZero() {
		if !r.attributeType.Obsolete {
			r.attributeType.Obsolete = true
		}
	}

	return *r
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
Name returns the string form of the principal name of the receiver instance, if set.
*/
func (r AttributeType) Name() (id string) {
	if !r.IsZero() {
		id = r.attributeType.Name.index(0)
	}

	return
}

/*
Names returns the underlying instance of [DefinitionName] from within
the receiver.
*/
func (r AttributeType) Names() (names DefinitionName) {
	return r.attributeType.Name
}

/*
NewAttributeTypes initializes and returns a new [AttributeTypes] instance,
configured to allow the storage of all [AttributeType] instances.
*/
func NewAttributeTypes() AttributeTypes {
	r := AttributeTypes(newCollection(``))
	r.cast().SetPushPolicy(r.canPush)

	return r
}

/*
NewAttributeTypeOIDList initializes and returns a new [AttributeTypes] that has
been cast from an instance of [OIDList] and configured to allow the storage of
arbitrary [AttributeType] instances.
*/
func NewAttributeTypeOIDList() AttributeTypes {
	r := AttributeTypes(newOIDList(``))
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
func (r *AttributeType) SetNumericOID(id string) AttributeType {
	if r.attributeType == nil {
		r.attributeType = newAttributeType()
	}

	if isNumericOID(id) {
		if len(r.attributeType.OID) == 0 {
			r.attributeType.OID = id
		}
	}

	return *r
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
func (r *AttributeType) SetDescription(desc string) AttributeType {
	if len(desc) < 3 {
		return *r
	}

	if r.attributeType == nil {
		r.attributeType = new(attributeType)
	}

	if !(rune(desc[0]) == rune(39) && rune(desc[len(desc)-1]) == rune(39)) {
		desc = `'` + desc + `'`
		if !r.IsZero() && isValidDescription(desc) {
			r.attributeType.Desc = desc
		}
	}

	return *r
}

func (r AttributeType) schema() (s Schema) {
	if !r.IsZero() {
		s = r.attributeType.schema
	}

	return
}

/*
String is a stringer method that returns the string representation
of the receiver instance.
*/
func (r AttributeType) String() (at string) {
	if !r.IsZero() {
		at = r.attributeType.s
	}

	return
}

func (r *attributeType) prepareString() (err error) {
	buf := newBuf()
	r.t = newTemplate(`attributeType`).
		Funcs(funcMap(map[string]any{
			`ExtensionSet`: r.Extensions.tmplFunc,
			`IsObsolete`:   func() bool { return r.Obsolete },
			`Usage`:        func() string { return AttributeType{r}.Usage() },
		}))

	if r.t, err = r.t.Parse(attributeTypeTmpl); err == nil {
		if err = r.t.Execute(buf, r); err == nil {
			r.s = buf.String()
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
Schema instance, or an actual [AttributeType] instance already obtained or crafted.

This is a fluent method.
*/
func (r AttributeType) SetSuperType(x any) AttributeType {
	if r.IsZero() {
		return r
	} else if r.schema().IsZero() {
		return r
	}

	var sup AttributeType
	switch tv := x.(type) {
	case string:
		sup = r.schema().AttributeTypes().get(tv)
	case AttributeType:
		sup = tv
	}

	if err := r.attributeType.verifySuperType(sup.attributeType); err == nil && !sup.IsZero() {
		r.attributeType.SuperType = sup
	}

	return r
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
SetUsage assigns the specified USAGE to the receiver. Input types
may be string, int or uint.

	1, or directoryOperation
	2, or distributedOperation
	3, or dSAOperation

Any other value results in assignment of the userApplication USAGE.

This is a fluent method.
*/
func (r AttributeType) SetUsage(u any) AttributeType {
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
				r.attributeType.Usage = UserApplicationUsage
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
				r.attributeType.Usage = UserApplicationUsage
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
List returns a map[string][]string instance which represents the current
inventory of attribute type instances within the receiver.  The keys are
numeric OIDs, while the values are zero (0) or more string slices, each
representing a name by which the definition is known.

For example: "2.5.4.3" = []string{"cn","commonName"}
*/
func (r AttributeTypes) List() (list map[string][]string) {
	list = make(map[string][]string, 0)
	for i := 0; i < r.len(); i++ {
		def := r.index(i)
		list[def.NumericOID()] = def.Names().List()

	}

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
Push returns an error following an attempt to push an AttributeType
into the receiver stack instance.
*/
func (r AttributeTypes) Push(at any) error {
	return r.push(at)
}

func (r AttributeTypes) push(at any) (err error) {
	err = errorf("%T instance is nil; cannot append to %T", at, r)
	if at != nil {
		r.cast().Push(at)
		err = nil
	}

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

func (r Schema) processAttributeType(ctx antlr4512.IAttributeTypeDescriptionContext) (at AttributeType, err error) {
	_at := new(attributeType)
	_at.schema = r

	if ctx.OpenParen() == nil || ctx.CloseParen() == nil {
		err = errorf("Definition encapsulation error; missing parenthesis?")
		return
	}

	for k, ct := 0, ctx.GetChildCount(); k < ct && err == nil; k++ {
		switch tv := ctx.GetChild(k).(type) {
		case *antlr4512.NumericOIDOrMacroContext,
			*antlr4512.OpenParenContext,
			*antlr4512.CloseParenContext:
			err = _at.setCritical(tv)

		case *antlr4512.ATUsageContext,
			*antlr4512.DefinitionNameContext,
			*antlr4512.DefinitionExtensionsContext,
			*antlr4512.DefinitionDescriptionContext:
			err = _at.setMisc(tv)

		case *antlr4512.ATCollectiveContext,
			*antlr4512.ATSingleValueContext,
			*antlr4512.DefinitionObsoleteContext,
			*antlr4512.ATNoUserModificationContext:
			err = _at.setBooleanContexts(tv)

		case *antlr4512.DefinitionSyntaxContext,
			*antlr4512.MinimumUpperBoundsContext:
			err = _at.setSyntaxContexts(tv)

		case *antlr4512.ATEqualityContext,
			*antlr4512.ATOrderingContext,
			*antlr4512.ATSubstringContext:
			err = _at.setMatchingRuleContexts(tv)

		case *antlr4512.ATSuperTypeContext:
			err = _at.setSuperTypeContext(tv)

		default:
			env := AttributeType{_at}
			err = isErrImpl(env.Type(), env.OID(), tv)
		}
	}

	// check for errors in new instance, and (if needed)
	// resolve macro to numeric OID.
	if err == nil {
		r.resolveByMacro(AttributeType{_at})

		if err = _at.check(); err == nil {
			if err = _at.prepareString(); err == nil {
				_at.t = nil
				at = AttributeType{_at}
			}
		}
	}

	return
}

/*
MinUpperBounds returns the minimum upper bounds as a uint instance.
If zero (0), no minimum upper bounds was specified for the definition,
meaning no value length limit is imposed through its use.
*/
func (r AttributeType) MinUpperBounds() uint {
	return r.attributeType.MUB
}

/*
IsZero returns a Boolean value indicative of nilness of the
receiver instance.
*/
func (r AttributeType) IsZero() bool {
	return r.attributeType == nil
}

func (r *attributeType) setBooleanContexts(ctx any) (err error) {
	switch ctx.(type) {
	case *antlr4512.ATCollectiveContext:
		r.Collective = true
	case *antlr4512.ATSingleValueContext:
		r.SingleVal = true
	case *antlr4512.DefinitionObsoleteContext:
		r.Obsolete = true
	case *antlr4512.ATNoUserModificationContext:
		r.NoUserMod = true
	default:
		err = errorf("Unknown boolean context %T", ctx)
	}

	return
}

func (r *attributeType) check() (err error) {
	if r == nil {
		err = errorf("%T is nil", r)
		return
	}

	if len(r.OID) == 0 {
		err = errorf("%T lacks an OID", r)
		return
	}

	if r.Collective && r.SingleVal {
		err = errorf("%T is both collective AND single-valued", r)
	}

	return
}

func (r *attributeType) effectiveSyntax() (ls LDAPSyntax) {
	if r == nil {
		return ls
	}

	syn := r.Syntax
	sat := r.SuperType.attributeType
	var ssyn LDAPSyntax
	if sat != nil {
		ssyn = sat.effectiveSyntax()
	}

	if !ssyn.IsZero() {
		ls = ssyn
	} else if !syn.IsZero() {
		if ssyn.IsZero() {
			ls = syn
		}
	}

	return
}

func (r *attributeType) setMisc(ctx any) (err error) {

	switch tv := ctx.(type) {
	case *antlr4512.DefinitionNameContext:
		r.Name, err = nameContext(tv)
	case *antlr4512.ATUsageContext:
		err = r.setUsageContext(tv)
	case *antlr4512.DefinitionDescriptionContext:
		r.Desc, err = descContext(tv)
	case *antlr4512.DefinitionExtensionsContext:
		r.Extensions, err = extContext(tv)
	default:
		err = errorf("Unknown miscellaneous context '%T'", ctx)
	}

	return
}

func (r *attributeType) setCritical(ctx any) (err error) {

	switch tv := ctx.(type) {
	case *antlr4512.NumericOIDOrMacroContext:
		r.OID, r.Macro, err = numOIDContext(tv)
	case *antlr4512.OpenParenContext, *antlr4512.CloseParenContext:
		err = parenContext(tv)
	default:
		err = errorf("Unknown critical context '%T'", ctx)
	}

	return
}

func (r *attributeType) setMatchingRuleContexts(ctx any) (err error) {

	switch tv := ctx.(type) {
	case *antlr4512.ATEqualityContext:
		err = r.setMatchingRuleContext(1, tv.OID())
	case *antlr4512.ATOrderingContext:
		err = r.setMatchingRuleContext(2, tv.OID())
	case *antlr4512.ATSubstringContext:
		err = r.setMatchingRuleContext(3, tv.OID())
	default:
		err = errorf("Unknown matchingRule context %T", ctx)
	}

	return
}

func (r *attributeType) setMatchingRuleContext(typ int, ctx antlr4512.IOIDContext) (err error) {
	var n, d []string
	if n, d, err = oIDContext(ctx.(*antlr4512.OIDContext)); err != nil {
		return
	}

	var mr string
	if len(n) > 0 {
		mr = n[0]
	} else if len(d) > 0 {
		mr = d[0]
	} else {
		err = errorf("No %T value found", ctx)
		return
	}

	var mrl MatchingRule
	if mrl = r.schema.MatchingRules().get(mr); mrl.IsZero() {
		err = errorf("matching rule '%s' not found", mr)
		return
	}

	switch typ {
	case 1:
		r.Equality = mrl
	case 2:
		r.Ordering = mrl
	case 3:
		r.Substring = mrl
	default:
		err = errorf("Invalid matchingRule type '%d'", typ)
	}

	return
}

/*
setSuperTypeContext processes the intended super type of the receiver, returning
an error if invalid or incomplete.
*/
func (r *attributeType) setSuperTypeContext(ctx *antlr4512.ATSuperTypeContext) (err error) {
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
		err = errorf("super type OID assertion failed for %T", o)
		return
	}

	// determine the nature of the OID and process it.
	// Fail if we find anything else.
	if n, d, err = oIDContext(assert); err != nil {
		return
	}

	// Figure out what we got from oIDContext
	// processing (numeric OID or descriptor?)
	var _sup string
	if len(n) > 0 {
		_sup = n[0]
	} else if len(d) > 0 {
		_sup = d[0]
	} else {
		err = errorf("No %T OID or Descriptor found", ctx)
		return
	}

	// look up the intended super type of r, fail if not
	// already loaded into Schema.AttributeTypes.
	sup := r.schema.AttributeTypes().get(_sup)
	if sup.IsZero() {
		err = errorf("super type '%s' not found in Schema", _sup)
		return
	}

	// make sure our super type chain is valid. If
	// so, ship it out.
	if err = r.verifySuperType(sup.attributeType); err == nil {
		r.SuperType = sup
	}

	return
}

/*
verifySuperType returns an error following the execution of basic sanity
checks meant to assess the intended super type chain.
*/
func (r *attributeType) verifySuperType(sup *attributeType) (err error) {
	// perform basic sanity check of super type
	if err = sup.check(); err != nil {
		return
	}

	// make sure super type and sub type aren't
	// one-in-the-same.
	if r.OID == sup.OID {
		err = errorf("cyclical super type loop detected (%s)", r.OID)
		return
	}

	// make sure that, if our super type is a sub
	// type itself, that its own super type isn't
	// the intended sub type (the receiver).
	if !sup.SuperType.IsZero() {
		if r.OID == sup.SuperType.NumericOID() {
			err = errorf("cyclical super type recursion detected (%s)", r.OID)
			return
		}
	}

	return
}

func (r *attributeType) setUsageContext(ctx any) (err error) {
	switch tv := ctx.(type) {
	case int:
		err = errorf("Invalid %T.Usage value", r)
		if 0 <= tv && tv <= 3 {
			err = nil
			r.Usage = uint(tv)
		}
	case uint:
		err = r.setUsageContext(int(tv))
	case string:
		err = errorf("Invalid usage text '%s'", tv)
		switch tv {
		case `userApplication`, ``:
			err = nil
		case `directoryOperation`:
			err = nil
			r.Usage = uint(1)
		case `distributedOperation`:
			err = nil
			r.Usage = uint(2)
		case `dSAOperation`:
			err = nil
			r.Usage = uint(3)
		}
	case *antlr4512.ATUsageContext:
		err = r.setUsageContext(tv.GetText())
	}

	return
}

func (r *attributeType) setSyntaxContexts(ctx any) (err error) {
	err = errorf("Unknown attribute syntax context %T", ctx)

	switch tv := ctx.(type) {
	case *antlr4512.DefinitionSyntaxContext:
		var _syn string
		if _syn, err = syntaxContext(tv); err != nil {
			break
		}

		err = errorf("syntax '%s' not found", _syn)
		if syn := r.schema.LDAPSyntaxes().get(_syn); !syn.IsZero() {
			r.Syntax = syn
			err = nil
		}

	case *antlr4512.MinimumUpperBoundsContext:
		err = errorf("Min. Upper bounds %T instance is nil", tv)
		if mb := tv.MinUpperBounds(); mb != nil {
			var m int
			m, err = atoi(trimS(trim(mb.GetText(), `{}`)))
			r.MUB = uint(m)
		}
	}

	return
}

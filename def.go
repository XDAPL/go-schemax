package schemax

import (
	"reflect"
)

/*
Definition provides methods common to all discrete definition types:

- *AttributeType

- *ObjectClass

- *LDAPSyntax

- *MatchingRule

- *NameForm

- *DITContentRule

- *DITStructureRule
*/
type Definition interface {
	// SetObsolete assigns the OBSOLETE bit to the receiver field value,
	// indicating definition obsolescence within the subschema subentry.
	SetObsolete()

	// Obsolete returns a boolean value indicative of whether the definition
	// is set with the OBSOLETE bit value.
	Obsolete() bool

	// Equal performs a deep-equal between the receiver and the provided
	// interface type, if applicable. A boolean value indicative of value
	// equality is returned.
	Equal(interface{}) bool

	// IsZero returns a boolean value indicative of whether the receiver
	// is considered zero, or undefined.
	IsZero() bool

	// Validate returns an error during the validation process. If a non
	// nil error is returned, it will reveal the nature of the deficiency.
	// This method is automatically run by the Marshal process, but can be
	// useful in cases where definitions are crafted manually and not as
	// a result of raw parsing.
	Validate() error

	// String is a convenient, but unsafe, alternative to the recommended
	// Unmarshal procedure for definitions. This method will Unmarshal a
	// Definition-based type into a string value. However, if any errors
	// occur during this process, a zero length string is returned, thus
	// making this unsafe in situations where error handling is critical.
	String() string

	// SetInfo assigns the byte slice to the receiver. This is a user-leveraged
	// field intended to allow arbitrary information (documentation?) to be
	// assigned to the definition.
	SetInfo([]byte)

	// Info returns the assigned informational byte slice instance stored
	// within the receiver.
	Info() []byte

	// Map returns a map[string][]string instance containing the receiver's
	// components. A nil map is returned if validation checks fail.
	Map() map[string][]string

	// Type returns the formal string name of the type of definition held
	// by the receiver. The value will be one of:
	//
	//   - LDAPSyntax
	//   - MatchingRule
	//   - AttributeType
	//   - MatchingRuleUse
	//   - ObjectClass
	//   - DITContentRule
	//   - NameForm
	//   - DITStructureRule
	Type() string

	// SetSpecifier assigns a string value to the receiver, useful for placement
	// into configurations that require a type name (e.g.: attributetype). This
	// will be displayed at the beginning of the definition value during the 
	// unmarshal or unsafe stringification process.
	SetSpecifier(string)

	// SetUnmarshalFunc assigns the provided DefinitionUnmarshalFunc signature
	// value to the receiver. The provided function shall be executed during the
	// unmarshal or unsafe stringification process.
	SetUnmarshalFunc(DefinitionUnmarshalFunc)

	// UnmarshalFunc is a package-included function that honors the signature
	// of the first class (closure) DefinitionUnmarshalFunc type. The purpose
	// of this function, and similar user-devised ones, is to help unmarshal
	// a definition with specific formatting included, such as linebreaks,
	// leading specifier declarations and indenting.
	UnmarshalFunc() (string, error)
}

/*
OID is a type alias for []int that describes an ASN.1 Object Identifier.
*/
type OID []int

/*
NewOID returns an instance of OID based on the provided input variable x. Supported OID input types are []string, []int or string.
*/
func NewOID(x interface{}) (oid OID) {
	switch tv := x.(type) {
	case string:
		return NewOID(split(tv, `.`))
	case OID:
		return tv
	case []string:
		var o OID
		for _, arc := range tv {
			x, err := atoi(arc)
			if err != nil {
				return
			}
			if x < 0 {
				return
			}
			o = append(o, x)
		}
		oid = o
	case []int:
		var o OID
		for _, arc := range tv {
			if arc < 0 {
				return
			}
			o = append(o, arc)
		}
		oid = o
	}

	return
}

/*
IsZero returns a boolean value indicative of whether the receiver is undefined.
*/
func (r OID) IsZero() bool {
	return r.Len() == 0
}

/*
IsZero returns a boolean value indicative of whether the receiver is undefined.
*/
func (r Description) IsZero() bool {
	return len(r) == 0
}

/*
String is a stringer method that returns the string-form of the receiver instance.
*/
func (r OID) String() string {
	var o []string
	for _, arc := range r {
		o = append(o, itoa(arc))
	}
	return join(o, `.`)
}

/*
Len returns an integer that reflects the number of arcs within the receiver.
*/
func (r OID) Len() int {
	return len(r)
}

/*
Equal returns a boolean value indicative of whether the receiver and x are equal.
*/
func (r OID) Equal(x interface{}) bool {
	o := NewOID(x)
	if r.Len() != o.Len() {
		return false
	}

	for idx, arc := range r {
		if arc != o[idx] {
			return false
		}
	}

	return true
}

/*
Description manifests as a single text value intended to describe the nature and purpose of the definition bearing this type in human readable format.
*/
type Description string

/*
String is a stringer method that returns the string-form of the receiver instance.
*/
func (r Description) String() string {
	return `'` + string(r) + `'`
}

/*
definition can describe the nature and contents of any single specific schema definition:

 - an AttributeType instance, or ...
 - an ObjectClass instance, or ...
 - an LDAPSyntax instance, or ...
 - a MatchingRule instance, or ...
 - a MatchingRuleUse instance, or ...
 - a DITContentRule instance, or ...
 - a DITStructureRule instance, or ...
 - a NameForm instance

The "typ" field MUST only contain one (1) of the above.

Length/Cap of "fields" and "values" MUST always be equal, even if no actual field _values_ were perceived.

This type is only defined for operations inherent to this package.

The alm field stores the OPTIONAL user-provided Macros instance for the resolution of OID "aliases".
*/
type definition struct {
	alm    *Macros
	typ    reflect.Type
	fields []reflect.StructField
	values []reflect.Value
	labels map[int]string
	meths  map[int]parseMeth
}

func definitionType(def definition) (n string) {
	n = `unknown`
	t := def.typ.Name()
	switch t {
	case `ObjectClass`:
		n = `oc`
	case `StructuralObjectClass`:
		n = `soc`
	case `AttributeType`:
		n = `at`
	case `SuperiorAttributeType`:
		n = `sat`
	case `LDAPSyntax`:
		n = `ls`
	case `Equality`:
		n = `eq`
	case `Substring`:
		n = `ss`
	case `Ordering`:
		n = `ord`
	case `MatchingRule`:
		n = `mr`
	case `MatchingRuleUse`:
		n = `mru`
	case `DITContentRule`:
		n = `dcr`
	case `DITStructureRule`:
		n = `dsr`
	case `NameForm`:
		n = `nf`
	}

	return
}

/*
newDefinition will parse the provided interface (x) into an instance of *definition, which is returned along with a success-indicative boolean value.  If values are present within the provided interface, they are preserved.
*/
func newDefinition(z interface{}, alm *Macros) (def *definition, ok bool) {
	def = new(definition)
	def.alm = alm

	switch tv := z.(type) {
	case *AttributeType, *SuperiorAttributeType:
		if assert, is := tv.(*SuperiorAttributeType); is {
			def.typ = valueOf(assert).Elem().Type()
			def.meths = assert.methMap()
			def.labels = assert.labelMap()
		} else {
			def.typ = valueOf(tv.(*AttributeType)).Elem().Type()
			def.meths = tv.(*AttributeType).methMap()
			def.labels = tv.(*AttributeType).labelMap()
		}
	case *ObjectClass:
		def.typ = valueOf(tv).Elem().Type()
		def.meths = tv.methMap()
		def.labels = tv.labelMap()
	case *LDAPSyntax:
		def.typ = valueOf(tv).Elem().Type()
		def.meths = tv.methMap()
		def.labels = tv.labelMap()
	case *MatchingRule:
		def.typ = valueOf(tv).Elem().Type()
		def.meths = tv.methMap()
		def.labels = tv.labelMap()
	case *MatchingRuleUse:
		def.typ = valueOf(tv).Elem().Type()
		def.meths = tv.methMap()
		def.labels = tv.labelMap()
	case *DITContentRule:
		def.typ = valueOf(tv).Elem().Type()
		def.meths = tv.methMap()
		def.labels = tv.labelMap()
	case *DITStructureRule:
		def.typ = valueOf(tv).Elem().Type()
		def.meths = tv.methMap()
		def.labels = tv.labelMap()
	case *NameForm:
		def.typ = valueOf(tv).Elem().Type()
		def.meths = tv.methMap()
		def.labels = tv.labelMap()
	default:
		return
	}

	def.fields = make([]reflect.StructField, def.typ.NumField(), def.typ.NumField())
	def.values = make([]reflect.Value, len(def.fields), cap(def.fields))

	for i := 0; i < len(def.fields); i++ {
		def.fields[i] = def.typ.Field(i)
		def.values[i] = valueOf(z).Elem().Field(i)
	}
	ok = true

	return
}

/*
fvindex returns the integer index for given field or value based on the input struct field name (e.g.: `OID`)
*/
func (def definition) lfindex(term string) (idx int) {
	return lfindex(term, def)
}

/*
definitionType returns the acronym describing the nature of the definition:

 - `at`  (AttributeType)
 - `oc`  (ObjectClass)
 - `ls`  (LDAPSyntax)
 - `mr`  (MatchingRule)
 - `nf`  (NameForm)
 - `mru` (MatchingRuleUse)
 - `dcr` (DITContentRule)
 - `dsr` (DITStructureRule)
*/
func (def definition) definitionType() (n string) {
	return definitionType(def)
}

func (def *definition) setKind(value string, idx int) (err error) {
	if def.alreadySet(idx) {
		return // silent discard
	}
	def.values[idx].Set(valueOf(newKind(value)))

	return
}

func (def *definition) setdefinitionFlags(label string, x interface{}) (err error) {
	switch tv := x.(type) {
	case *AttributeType:
		switch label {
		case SingleValue.String():
			tv.setdefinitionFlags(SingleValue)
			return
		case Collective.String():
			tv.setdefinitionFlags(Collective)
			return
		case Obsolete.String():
			tv.setdefinitionFlags(Obsolete)
			return
		case NoUserModification.String():
			tv.setdefinitionFlags(NoUserModification)
			return
		case HumanReadable.String():
			tv.setdefinitionFlags(HumanReadable)
			return
		}
	case *ObjectClass:
		switch label {
		case Obsolete.String():
			tv.setdefinitionFlags(Obsolete)
			return
		}
	case *LDAPSyntax:
		switch label {
		case HumanReadable.String():
			tv.setdefinitionFlags(HumanReadable)
			return
		}
	case *MatchingRule:
		switch label {
		case Obsolete.String():
			tv.setdefinitionFlags(Obsolete)
			return
		}
	case *MatchingRuleUse:
		switch label {
		case Obsolete.String():
			tv.setdefinitionFlags(Obsolete)
			return
		}
	case *DITContentRule:
		switch label {
		case Obsolete.String():
			tv.setdefinitionFlags(Obsolete)
			return
		}
	case *DITStructureRule:
		switch label {
		case Obsolete.String():
			tv.setdefinitionFlags(Obsolete)
			return
		}
	case *NameForm:
		switch label {
		case Obsolete.String():
			tv.setdefinitionFlags(Obsolete)
			return
		}
	}
	err = raise(invalidFlag,
		"setdefinitionFlags: unable to resolve '%T' type (label:'%s')",
		x, label)

	return
}

func (def *definition) setExtensions(label string, value []string, idx int) (err error) {
	z, ok := def.values[idx].Interface().(Extensions)
	if !ok {
		return raise(unknownDefinition,
			"setExtensions: unexpected type '%T'",
			def.values[idx].Interface())
	} else if z.IsZero() {
		z = NewExtensions()
	}

	z.Set(label, value)
	def.values[idx].Set(valueOf(z))

	return nil
}

func (def *definition) setUsage(value string, idx int) (err error) {
	def.values[idx].Set(valueOf(newUsage(value)))

	return
}

func (def *definition) setDesc(idx int, value string) (err error) {
	if def.alreadySet(idx) {
		return // silent discard
	}
	def.values[idx].Set(valueOf(Description(value)))

	return
}

func (def *definition) setName(idx int, value ...string) (err error) {
	if def.alreadySet(idx) {
		return // silent discard
	}

	if name := NewName(value...); !name.IsZero() {
		def.values[idx].Set(valueOf(name))
	}

	return
}

func (def *definition) setStructuralObjectClass(
	occ ObjectClassCollection,
	x interface{},
	value string,
	idx int) (err error) {

	if def.alreadySet(idx) {
		return
	}

	oc := occ.Get(value)
	if oc.IsZero() {
		return raise(invalidObjectClass,
			"setStructuralObjectClass: no such %T was found in %T for value '%s' (type: %T)",
			oc, value)
	}
	if !oc.Kind.is(Structural) {
		return raise(invalidObjectClass,
			"setStructuralObjectClass: %T (%s) is not STRUCTURAL kind",
			oc, oc.OID.String())
	}
	def.values[idx].Set(valueOf(StructuralObjectClass{oc}))

	return
}

func (def *definition) setSyntax(
	lsc LDAPSyntaxCollection,
	x interface{},
	value string,
	idx int) (err error) {

	switch def.definitionType() {
	case `mr`:
		ls := lsc.Get(value)
		if ls.IsZero() {
			return raise(invalidSyntax,
				"setSyntax: no %T in %T for '%s' (type: %T)",
				ls, lsc, value, x)
		}
		def.values[idx].Set(valueOf(ls))
	case `at`:
		err = def.setAttrTypeSyntax(lsc, x, value, idx)
	}

	return
}

func (def *definition) setAttrTypeSyntax(
	lsc LDAPSyntaxCollection,
	x interface{},
	value string, idx int) (err error) {

	assert, ok := x.(*AttributeType)
	if !ok {
		return raise(unknownDefinition,
			"setAttrTypeSyntax: unexpected type '%T'", x)
	}

	// Check (and handle) a MUB in the event we're using one ...
	if mub, _, ok := parse_mub(value); ok {
		assert.setMUB(mub[1])
		// MUB detected
		if ls := lsc.Get(mub[0]); !ls.IsZero() {
			def.values[idx].Set(valueOf(ls))
		} else {
			err = raise(invalidSyntax,
				"setAttrTypeSyntax: no %T in %T for '%s' (type: %T, mub:true)",
				ls, lsc, mub[0], x)
		}
	} else {
		// No MUB detected
		if ls := lsc.Get(value); !ls.IsZero() {
			def.values[idx].Set(valueOf(ls))
		} else {
			err = raise(invalidSyntax,
				"setAttrTypeSyntax: no %T in %T for '%s' (type: %T)",
				ls, lsc, value, x)
		}
	}

	return
}

func (def *definition) setNameForm(
	nfc NameFormCollection,
	x interface{},
	value string, idx int) (err error) {

	nf := nfc.Get(value)
	if nf.IsZero() {
		return raise(unknownElement,
			"setNameForm: no such %T was found in %T for value '%s' (type: %T)",
			nf, nfc, value, x)
	}
	def.values[idx].Set(valueOf(nf))

	return
}

func (def *definition) setSuperiorObjectClasses(
	occ ObjectClassCollection,
	x interface{},
	value []string,
	idx int) (err error) {

	z, ok := def.values[idx].Interface().(ObjectClassCollection)
	if !ok {
		z = NewSuperiorObjectClasses()
	}

	for _, v := range value {
		oc := occ.Get(v)
		if oc.IsZero() {
			return raiseUnknownElement(`setSuperiorObjectClasses`,
				oc, occ, v, x)
		}

		z.Set(oc)
		def.values[idx].Set(valueOf(z))
	}

	return
}

func (def *definition) setAuxiliaryObjectClasses(
	occ ObjectClassCollection,
	x interface{},
	value []string,
	idx int) (err error) {

	z, ok := def.values[idx].Interface().(ObjectClassCollection)
	if !ok {
		z = NewAuxiliaryObjectClasses()
	}

	for _, v := range value {
		oc := occ.Get(v)
		if oc.IsZero() {
			return raiseUnknownElement(`setAuxiliaryObjectClasses`,
				oc, occ, v, x)
		}

		z.Set(oc)
		def.values[idx].Set(valueOf(z))
	}

	return
}

func (def *definition) setApplicableAttributeTypes(
	atc AttributeTypeCollection,
	x interface{},
	value []string,
	idx int) (err error) {

	z, ok := def.values[idx].Interface().(AttributeTypeCollection)
	if !ok {
		z = NewApplicableAttributeTypes()
	}

	for _, v := range value {
		at := atc.Get(v)
		if at.IsZero() {
			return raiseUnknownElement(`setApplicableAttributeTypes`,
				at, atc, v, x)
		}

		z.Set(at)
		def.values[idx].Set(valueOf(z))
	}

	return
}

func (def *definition) setRequiredAttributeTypes(
	atc AttributeTypeCollection,
	x interface{},
	value []string,
	idx int) (err error) {

	z, ok := def.values[idx].Interface().(AttributeTypeCollection)
	if !ok {
		z = NewRequiredAttributeTypes()
	}

	for _, v := range value {
		at := atc.Get(v)
		if at.IsZero() {
			return raiseUnknownElement(`setRequiredAttributeTypes`,
				at, atc, v, x)
		}

		z.Set(at)
		def.values[idx].Set(valueOf(z))
	}

	return
}

func (def *definition) setPermittedAttributeTypes(
	atc AttributeTypeCollection,
	x interface{},
	value []string,
	idx int) (err error) {

	z, ok := def.values[idx].Interface().(AttributeTypeCollection)
	if !ok {
		z = NewPermittedAttributeTypes()
	}

	for _, v := range value {
		at := atc.Get(v)
		if at.IsZero() {
			return raiseUnknownElement(`setPermittedAttributeTypes`,
				at, atc, v, x)
		}

		z.Set(at)
		def.values[idx].Set(valueOf(z))
	}

	return
}

func (def *definition) setProhibitedAttributeTypes(
	atc AttributeTypeCollection,
	x interface{},
	value []string,
	idx int) (err error) {

	z, ok := def.values[idx].Interface().(AttributeTypeCollection)
	if !ok {
		z = NewProhibitedAttributeTypes()
	}

	for _, v := range value {
		at := atc.Get(v)
		if at.IsZero() {
			return raiseUnknownElement(`setProhibitedAttributeTypes`,
				at, atc, v, x)
		}

		z.Set(at)
		def.values[idx].Set(valueOf(z))
	}

	return
}

func (def *definition) setSuperiorAttributeType(
	atc AttributeTypeCollection,
	x interface{},
	value string,
	idx int) (err error) {

	if def.alreadySet(idx) {
		return // silent discard
	}

	at := atc.Get(value)
	if at.IsZero() {
		return raiseUnknownElement(`setSuperiorAttributeType`,
			at, atc, value, x)
	}
	def.values[idx].Set(valueOf(SuperiorAttributeType{at}))

	return nil
}

/*
setSuperiorDITStructureRules sets the SUP value of argument x, given the string value argument. This process fails if the manifest lookup fails.
*/
func (def *definition) setSuperiorDITStructureRules(
	dsrc DITStructureRuleCollection,
	x interface{},
	value []string,
	idx int) (err error) {

	z, ok := def.values[idx].Interface().(DITStructureRuleCollection)
	if !ok {
		z = NewSuperiorDITStructureRules()
	}

	for _, v := range value {
		dsr := dsrc.Get(v)
		if dsr.IsZero() {
			return raiseUnknownElement(`setSuperiorDITStructureRules`,
				dsr, dsrc, v, x)
		}

		z.Set(dsr)
		def.values[idx].Set(valueOf(z))
	}

	return
}

/*
setEqSubOrd sets the SUBSTR, ORDERING or EQUALITY value of argument x, given the string value argument. This process fails if the manifest lookup fails.
*/
func (def *definition) setEqSubOrd(
	mrc MatchingRuleCollection,
	x interface{},
	value string,
	idx int) (err error) {

	if def.alreadySet(idx) {
		return
	}

	assert, ok := x.(*AttributeType)
	if !ok {
		return raise(unknownDefinition,
			"setEqSubOrd: unexpected type '%T'", x)
	} else if assert.IsZero() {
		return raise(emptyDefinition,
			"setEqSubOrd: received empty '%T'", assert)
	}

	mr := mrc.Get(value)
	if mr.IsZero() {
		return raiseUnknownElement(
			`setEqSubOrd`,
			mr, mrc, value, x)
	}

	if err = def.setMR(assert, mr, idx); err != nil {
		err = raise(err,
			"setMR: def.labels[%d] ('%s') not set (type:'%T',value:'%s')",
			idx, def.labels[idx], x, value)
	}

	return
}

func (def *definition) setMR(
	dest *AttributeType,
	mr *MatchingRule,
	idx int) (err error) {

	// !! It is crucial we recognize when a supertype is in effect !!
	switch dest.SuperType.IsZero() {
	case true:
		switch def.labels[idx] {
		case `EQUALITY`:
			def.values[idx].Set(valueOf(Equality{mr}))
		case `SUBSTR`:
			def.values[idx].Set(valueOf(Substring{mr}))
		case `ORDERING`:
			def.values[idx].Set(valueOf(Ordering{mr}))
		default:
			err = unknownElement
		}
	case false:
		switch def.labels[idx] {
		case `EQUALITY`:
			def.values[idx].Set(valueOf(dest.SuperType.Equality))
		case `SUBSTR`:
			def.values[idx].Set(valueOf(dest.SuperType.Substring))
		case `ORDERING`:
			def.values[idx].Set(valueOf(dest.SuperType.Ordering))
		default:
			err = unknownElement
		}
	}

	return
}

func (def *definition) alreadySet(idx int) (isSet bool) {
	if idx < 0 {
		return
	}

	switch tv := def.values[idx].Interface().(type) {
	case Kind:
		isSet = !tv.IsZero()
	case Name:
		isSet = !tv.IsZero()
	case definitionFlags:
		isSet = !tv.IsZero()
	case Description:
		isSet = !tv.IsZero()
	case *AttributeType:
		isSet = !tv.IsZero()
	case RequiredAttributeTypes:
		isSet = !tv.IsZero()
	case PermittedAttributeTypes:
		isSet = !tv.IsZero()
	case ProhibitedAttributeTypes:
		isSet = !tv.IsZero()
	case SuperiorAttributeType:
		isSet = !tv.IsZero()
	case SuperiorObjectClasses:
		isSet = !tv.IsZero()
	case *ObjectClass:
		isSet = !tv.IsZero()
	case StructuralObjectClass:
		isSet = !tv.IsZero()
	case AuxiliaryObjectClasses:
		isSet = !tv.IsZero()
	case *LDAPSyntax:
		isSet = !tv.IsZero()
	case *MatchingRule:
		isSet = !tv.IsZero()
	case *MatchingRuleUse:
		isSet = !tv.IsZero()
	case ApplicableAttributeTypes:
		isSet = !tv.IsZero()
	case Equality:
		isSet = !tv.IsZero()
	case Ordering:
		isSet = !tv.IsZero()
	case Substring:
		isSet = !tv.IsZero()
	case *NameForm:
		isSet = !tv.IsZero()
	case *DITStructureRule:
		isSet = !tv.IsZero()
	case *SuperiorDITStructureRules:
		isSet = !tv.IsZero()
	case *DITContentRule:
		isSet = !tv.IsZero()
	default:
		isSet = true // better safe than sorry ...
	}

	return
}

type Name collection

/*
Len returns an integer indicative of the number of NAME values present within the receiver.
*/
func (r Name) Len() int {
	return len(r)
}

/*
Equal returns a boolean indicative of whether the value(s) provided match the receiver.
*/
func (r Name) Equal(x interface{}) bool {
	//return collection(r).equal(collection(n))
	if r.IsZero() {
		return false
	}

	switch tv := x.(type) {
	case nil:
		return false
	case string:
		for _, z := range r {
			if eq := equalFold(z.(string), tv); eq {
				return eq
			}
		}
	case Name:
		for _, z := range r {
			for _, y := range tv {
				if eq := equalFold(z.(string), y.(string)); eq {
					return eq
				}
			}
		}
	}

	return false
}

/*
NewName returns an initialized instance of Name.
*/
func NewName(n ...string) (name Name) {
	name = make(Name, 0)
	if len(n) == 0 {
		return
	}
	name.Set(n...)

	return
}

/*
Set applies one or more instances of string to the receiver.  Uniqueness is enforced through literal comparison.
*/
func (r *Name) Set(x ...string) {
	if len(x) == 0 {
		return
	}

	for n := range x {
		N := stripTags(x[n])
		if isNumericalOID(N) {
			continue
		}

		if _, found := r.Contains(N); !found {
			*r = append(*r, N)
		}
	}

	R := make(Name, len(*r), len(*r))
	for i, n := range *r {
		R[i] = n
	}
	*r = R
}

/*
Contains returns the index number and presence boolean that reflects the result of a term search.
*/
func (r Name) Contains(x interface{}) (idx int, has bool) {
	var term string
	idx = -1

	switch tv := x.(type) {
	case string:
		term = stripTags(tv)
	default:
		return
	}

	if r.IsZero() || len(term) == 0 {
		return
	}

	var n interface{}
	for idx, n = range r {
		if has = term == n.(string); has {
			break
		}
	}

	return
}

/*
Index returns the nth Name value within the receiver, or a zero-length string.
*/
func (r Name) Index(i int) string {
	val := collection(r).index(i)
	str, ok := val.(string)
	if !ok {
		return ``
	}

	return str
}

/*
IsZero returns a boolean value indicative of whether the receiver is undefined.
*/
func (r *Name) IsZero() bool {
	if r == nil {
		return true
	}
	return len(*r) == 0
}

/*
String is a qdescrs-compliant stringer method suitable for presentation.
*/
func (r Name) String() (str string) {
	switch len(r) {
	case 0:
		return
	case 1:
		return `'` + r[0].(string) + `'`
	}

	str += `( `
	for i := 0; i < len(r); i++ {
		str += `'` + r[i].(string) + `' `
	}
	str += `)`

	return
}

func (r Name) strings() (strings []string) {
	strings = make([]string, len(r), len(r))
	for i, v := range r {
		if assert, ok := v.(string); ok {
			strings[i] = assert
		}
	}

	return
}

func validateNames(value ...string) (err error) {
	// check slice as a whole
	switch {
	case len(value) > nameListMaxLen:
		err = raise(invalidName,
			"validation error: length of slices exceeds limit (%d)",
			nameListMaxLen)
	case len(value) == 0:
		err = raise(invalidName,
			"validation error: zero-length slices detected")
	}

	if err != nil {
		return
	}

	// check slice members
	for _, n := range value {
		if err = validateName(n); err != nil {
			return
		}
	}

	return
}

func validateName(value string) (err error) {
	switch {
	case len(value) > nameMaxLen:
		err = raise(invalidName,
			"validation error: length of slice exceeds limit (%d)",
			nameMaxLen)
	case len(value) == 0:
		err = raise(invalidName, "validateName: zero-length name")
	}

	for _, c := range value {
		if err != nil {
			break
		}

		switch {
		case !runeIsLetter(c) && !runeIsDigit(c) && c != '-':
			err = raise(invalidName,
				"validation error: bad char '%c'", c)
		}
	}

	return
}

func validateDesc(x interface{}) (err error) {
	var value string
	switch tv := x.(type) {
	case Description:
		value = string(tv)
	case []byte:
		value = string(tv)
	case string:
		value = tv
	default:
		return raise(invalidDescription,
			"validateDesc: unexpected type %T", tv)
	}

	switch {
	case len(value) > descMaxLen:
		err = raise(invalidDescription,
			"validation error: length exceeds limit (%d)",
			descMaxLen)
	case len(value) > 0 && !isUTF8([]byte(value)):
		err = raise(invalidDescription,
			"validation error: not UTF-8")
	}

	return
}

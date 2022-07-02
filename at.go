package schemax

import "sync"

/*
AttributeTypeCollection describes all of the following types:

• AttributeTypes

• RequiredAttributeTypes

• PermittedAttributeTypes

• ProhibitedAttributeTypes

• ApplicableAttributeTypes
*/
type AttributeTypeCollection interface {
	// Get returns the *AttributeType instance retrieved as a result
	// of a term search, based on Name or OID. If no match is found,
	// nil is returned.
	Get(interface{}) *AttributeType

	// Index returns the *AttributeType instance stored at the nth
	// index within the receiver, or nil.
	Index(int) *AttributeType

	// Equal performs a deep-equal between the receiver and the
	// interface AttributeTypeCollection provided.
	Equal(AttributeTypeCollection) bool

	// Set returns an error instance based on an attempt to add
	// the provided *AttributeType instance to the receiver.
	Set(*AttributeType) error

	// Contains returns the index number and presence boolean that
	// reflects the result of a term search within the receiver.
	Contains(interface{}) (int, bool)

	// String returns a properly-delimited sequence of string
	// values, either as a Name or OID, for the receiver type.
	String() string

	// Label returns the field name associated with the interface
	// types, or a zero string if no label is appropriate.
	Label() string

	// IsZero returns a boolean value indicative of whether the
	// receiver is considered zero, or undefined.
	IsZero() bool

	// Len returns an integer value indicative of the current
	// number of elements stored within the receiver.
	Len() int

	// SetSpecifier assigns a string value to all definitions within
	// the receiver. This value is used in cases where a definition
	// type name (e.g.: attributetype, objectclass, etc.) is required.
	// This value will be displayed at the beginning of the definition
	// value during the unmarshal or unsafe stringification process.
	SetSpecifier(string)

	// SetUnmarshaler assigns the provided DefinitionUnmarshaler
	// signature to all definitions within the receiver. The provided
	// function shall be executed during the unmarshal or unsafe
	// stringification process.
	SetUnmarshaler(DefinitionUnmarshaler)
}

/*
Usage describes the intended usage of an AttributeType definition as a single text value.  This can be one of four constant values, the first of which (userApplication) is implied in the absence of any other value and is not necessary to reveal in such a case.
*/
type Usage uint8

const (
	UserApplication Usage = iota
	DirectoryOperation
	DistributedOperation
	DSAOperation
)

/*
AttributeType conforms to the specifications of RFC4512 Section 4.1.2.
*/
type AttributeType struct {
	OID         OID
	Name        Name
	Description Description
	Obsolete    bool
	SuperType   SuperiorAttributeType
	Equality    Equality
	Ordering    Ordering
	Substring   Substring
	Syntax      *LDAPSyntax
	Usage       Usage
	Extensions  *Extensions
	flags       atFlags
	mub         uint
	ufn         DefinitionUnmarshaler
	spec        string
	info        []byte
}

/*
Type returns the formal name of the receiver in order to satisfy signature requirements of the Definition interface type.
*/
func (r *AttributeType) Type() string {
	return `AttributeType`
}

/*
String is an unsafe convenience wrapper for Unmarshal(r). If an error is encountered, an empty string definition is returned. If reliability and error handling are important, use Unmarshal.
*/
func (r *AttributeType) String() (def string) {
	def, _ = r.unmarshal()
	return
}

/*
HumanReadable is a convenience wrapper for Extensions.HumanReadable(), which returns a boolean value indicative of human readability. If super typing is in effect, an attempt to determine human readability by recursive inheritance is made. Failing this, if the receiver is assigned an *LDAPSyntax value, it is evaluated similarly. A fallback of true is returned as a last recourse.
*/
func (r *AttributeType) HumanReadable() bool {
	if !r.SuperType.IsZero() {
		return r.SuperType.HumanReadable()
	}
	if !r.Syntax.IsZero() {
		return r.Syntax.HumanReadable()
	}

	return true
}

/*
SuperiorAttributeType contains an embedded instance of *AttributeType. This alias type reflects the SUP field of an attributeType definition.
*/
type SuperiorAttributeType struct {
	*AttributeType
}

/*
AttributeTypes is a thread-safe collection of *AttributeType slice instances.
*/
type AttributeTypes struct {
	mutex  *sync.Mutex
	slice  collection
	macros *Macros
}

/*
ApplicableAttributeTypes contains an embedded instance of *AttributeTypes. This alias type reflects the APPLIES field of a matchingRuleUse definition.
*/
type ApplicableAttributeTypes struct {
	*AttributeTypes
}

/*
String returns a properly-delimited sequence of string values, either as a Name or OID, for the receiver type.
*/
func (r ApplicableAttributeTypes) String() string {
	return r.slice.attrs_oids_string()
}

/*
RequiredAttributeTypes contains an embedded instance of *AttributeTypes. This alias type reflects the MUST fields of a dITContentRule or objectClass definitions.
*/
type RequiredAttributeTypes struct {
	*AttributeTypes
}

/*
String returns a properly-delimited sequence of string values, either as a Name or OID, for the receiver type.
*/
func (r RequiredAttributeTypes) String() string {
	return r.slice.attrs_oids_string()
}

/*
PermittedAttributeTypes contains an embedded instance of *AttributeTypes. This alias type reflects the MAY fields of a dITContentRule or objectClass definitions.
*/
type PermittedAttributeTypes struct {
	*AttributeTypes
}

/*
String returns a properly-delimited sequence of string values, either as a Name or OID, for the receiver type.
*/
func (r PermittedAttributeTypes) String() string {
	return r.slice.attrs_oids_string()
}

/*
ProhibitedAttributeTypes contains an embedded instance of *AttributeTypes. This alias type reflects the NOT field of a dITContentRule definition.
*/
type ProhibitedAttributeTypes struct {
	*AttributeTypes
}

/*
String returns a properly-delimited sequence of string values, either as a Name or OID, for the receiver type.
*/
func (r ProhibitedAttributeTypes) String() string {
	return r.slice.attrs_oids_string()
}

/*
SetMacros assigns the *Macros instance to the receiver, allowing subsequent OID resolution capabilities during the addition of new slice elements.
*/
func (r *AttributeTypes) SetMacros(macros *Macros) {
	r.macros = macros
}

/*
SetSpecifier is a convenience method that executes the SetSpecifier method in iterative fashion for all definitions within the receiver.
*/
func (r *AttributeTypes) SetSpecifier(spec string) {
	for i := 0; i < r.Len(); i++ {
		r.Index(i).SetSpecifier(spec)
	}
}

/*
SetUnmarshaler is a convenience method that executes the SetUnmarshaler method in iterative fashion for all definitions within the receiver.
*/
func (r *AttributeTypes) SetUnmarshaler(fn DefinitionUnmarshaler) {
	for i := 0; i < r.Len(); i++ {
		r.Index(i).SetUnmarshaler(fn)
	}
}

/*
String is a stringer method that returns the string-form of the receiver instance.
*/
func (r Usage) String() string {
	switch r {
	case DirectoryOperation:
		return `directoryOperation`
	case DistributedOperation:
		return `distributedOperation`
	case DSAOperation:
		return `dSAOperation`
	}

	return `` // default is userApplication, but it need not be stated literally
}

/*
Contains is a thread-safe method that returns a collection slice element index integer and a presence-indicative boolean value based on a term search conducted within the receiver.
*/
func (r *AttributeTypes) Contains(x interface{}) (int, bool) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if !r.macros.IsZero() {
		if oid, resolved := r.macros.Resolve(x); resolved {
			return r.slice.contains(oid)
		}
	}
	return r.slice.contains(x)
}

/*
Index is a thread-safe method that returns the nth collection slice element if defined, else nil. This method supports use of negative indices which should be used with special care.
*/
func (r *AttributeTypes) Index(idx int) *AttributeType {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	assert, _ := r.slice.index(idx).(*AttributeType)
	return assert
}

/*
Get combines Contains and Index method executions to return an entry based on a term search conducted within the receiver.
*/
func (r *AttributeTypes) Get(x interface{}) *AttributeType {
	idx, found := r.Contains(x)
	if !found {
		return nil
	}

	return r.Index(idx)
}

/*
Len is a thread-safe method that returns the effective length of the receiver slice collection.
*/
func (r AttributeTypes) Len() int {
	if &r == nil {
		return 0
	}

	r.mutex.Lock()
	defer r.mutex.Unlock()

	return r.slice.len()
}

/*
String is a non-functional stringer method needed to satisfy interface type requirements and should not be used. See the String() method for ApplicableAttributeTypes, RequiredAttributeTypes, PermittedAttributeTypes and ProhibitedAttributeTypes instead.
*/
func (r *AttributeTypes) String() string {
	return ``
}

/*
IsZero returns a boolean value indicative of whether the receiver is considered empty or uninitialized.
*/
func (r *AttributeTypes) IsZero() bool {
	if r != nil {
		return r.slice.isZero()
	}
	return r == nil
}

/*
IsZero returns a boolean value indicative of whether the receiver is considered empty or uninitialized.
*/
func (r *AttributeType) IsZero() bool {
	return r == nil
}

/*
IsZero returns a boolean value indicative of whether the receiver is considered empty or uninitialized.
*/
func (r SuperiorAttributeType) IsZero() bool {
	return r.AttributeType.IsZero()
}

/*
Set is a thread-safe append method that returns an error instance indicative of whether the append operation failed in some manner. Uniqueness is enforced for new elements based on Object Identifier and not the effective Name of the definition, if defined.
*/
func (r *AttributeTypes) Set(x *AttributeType) error {
	if _, exists := r.Contains(x.OID); exists {
		return nil //silent
	}

	r.mutex.Lock()
	defer r.mutex.Unlock()

	return r.slice.append(x)
}

/*
SetSpecifier assigns a string value to the receiver, useful for placement into configurations that require a type name (e.g.: attributetype). This will be displayed at the beginning of the definition value during the unmarshal or unsafe stringification process.
*/
func (r *AttributeType) SetSpecifier(spec string) {
	r.spec = spec
}

/*
Equal performs a deep-equal between the receiver and the provided collection type.
*/
func (r *AttributeTypes) Equal(x AttributeTypeCollection) bool {
	return r.slice.equal(x.(*AttributeTypes).slice)
}

/*
Equal performs a deep-equal between the receiver and the provided definition type.

Description text is ignored.
*/
func (r *AttributeType) Equal(x interface{}) (eq bool) {
	var z *AttributeType
	switch tv := x.(type) {
	case *AttributeType:
		z = tv
	case SuperiorAttributeType:
		z = tv.AttributeType
	default:
		return
	}

	if z.IsZero() && r.IsZero() {
		eq = true
		return
	} else if z.IsZero() || r.IsZero() {
		return
	}

	if !z.Name.Equal(r.Name) {
		return
	}

	if !r.OID.Equal(z.OID) {
		return
	}

	if z.Usage != r.Usage {
		return
	}

	if z.flags != r.flags {
		return
	}

	if !z.SuperType.IsZero() && !r.SuperType.IsZero() {
		if !z.SuperType.OID.Equal(r.SuperType.OID) {
			return
		}
	}

	if !r.Syntax.Equal(z.Syntax) {
		return
	}

	if !r.Equality.Equal(z.Equality) {
		return
	}

	if !r.Ordering.Equal(z.Ordering) {
		return
	}

	if !r.Substring.Equal(z.Substring) {
		return
	}

	noexts := z.Extensions.IsZero() && r.Extensions.IsZero()
	if !noexts {
		eq = r.Extensions.Equal(z.Extensions)
	} else {
		eq = true
	}

	return
}

/*
NewAttributeType returns a newly initialized, yet effectively nil, instance of *AttributeType.

Users generally do not need to execute this function unless an instance of the returned type will be manually populated (as opposed to parsing a raw text definition).
*/
func NewAttributeType() *AttributeType {
	at := new(AttributeType)
	at.Extensions = NewExtensions()
	return at
}

/*
NewAttributeTypes initializes and returns a new AttributeTypeCollection interface object.
*/
func NewAttributeTypes() AttributeTypeCollection {
	var x interface{} = &AttributeTypes{
		mutex: &sync.Mutex{},
		slice: make(collection, 0, 0),
	}
	return x.(AttributeTypeCollection)
}

/*
NewApplicableAttributeTypes initializes an embedded instance of *AttributeTypes within the return value.
*/
func NewApplicableAttributeTypes() AttributeTypeCollection {
	var z *AttributeTypes = &AttributeTypes{
		mutex: &sync.Mutex{},
		slice: make(collection, 0, 0),
	}
	var x interface{} = &ApplicableAttributeTypes{z}
	return x.(AttributeTypeCollection)
}

/*
NewRequiredAttributeTypes initializes an embedded instance of *AttributeTypes within the return value.
*/
func NewRequiredAttributeTypes() AttributeTypeCollection {
	var z *AttributeTypes = &AttributeTypes{
		mutex: &sync.Mutex{},
		slice: make(collection, 0, 0),
	}
	var x interface{} = &RequiredAttributeTypes{z}
	return x.(AttributeTypeCollection)
}

/*
NewPermittedAttributeTypes initializes an embedded instance of *AttributeTypes within the return value.
*/
func NewPermittedAttributeTypes() AttributeTypeCollection {
	var z *AttributeTypes = &AttributeTypes{
		mutex: &sync.Mutex{},
		slice: make(collection, 0, 0),
	}
	var x interface{} = &PermittedAttributeTypes{z}
	return x.(AttributeTypeCollection)
}

/*
NewProhibitedAttributeTypes initializes an embedded instance of *AttributeTypes within the return value.
*/
func NewProhibitedAttributeTypes() AttributeTypeCollection {
	var z *AttributeTypes = &AttributeTypes{
		mutex: &sync.Mutex{},
		slice: make(collection, 0, 0),
	}
	var x interface{} = &ProhibitedAttributeTypes{z}
	return x.(AttributeTypeCollection)
}

func newUsage(x interface{}) Usage {
	switch tv := x.(type) {
	case string:
		switch toLower(tv) {
		case toLower(DirectoryOperation.String()):
			return DirectoryOperation
		case toLower(DistributedOperation.String()):
			return DistributedOperation
		case toLower(DSAOperation.String()):
			return DSAOperation
		}
	case uint:
		switch tv {
		case 0x1:
			return DirectoryOperation
		case 0x2:
			return DistributedOperation
		case 0x3:
			return DSAOperation
		}
	case int:
		if tv >= 0 {
			return newUsage(uint(tv))
		}
	}

	return UserApplication
}

/*
MaxLength returns the integer value, if one was specified, that defines the maximum acceptable value size supported by this *AttributeType per its associated *LDAPSyntax.  If not applicable, a 0 is returned.
*/
func (r *AttributeType) MaxLength() int {
	return int(r.mub)
}

/*
SetMaxLength sets the minimum upper bounds, or maximum length, of the receiver instance. The argument must be a positive, non-zero integer.

This will only apply to *AttributeTypes that use a human-readable syntax.
*/
func (r *AttributeType) SetMaxLength(max int) {
	r.setMUB(max)
}

/*
setMUB assigns the number (or string) as the minimum upper bounds value for the receiver.
*/
func (r *AttributeType) setMUB(mub interface{}) {

	switch tv := mub.(type) {
	case string:
		n, err := atoi(tv)
		if err != nil || n < 0 {
			return
		}
		r.mub = uint(n)
	case int:
		if tv > 0 {
			r.mub = uint(tv)
		}
	case uint:
		r.mub = tv
	}

}

/*
is returns a boolean value indicative of whether the provided interface argument is either an enabled atFlags value, or an associated *MatchingRule or *LDAPSyntax.

In the case of an *LDAPSyntax argument, if the receiver is in fact a sub type of another *AttributeType instance, a reference to that super type is chased and analyzed accordingly.
*/
func (r *AttributeType) is(b interface{}) bool {
	switch tv := b.(type) {
	case atFlags:
		return r.flags.is(tv)
	case *MatchingRule:
		switch {
		case tv.Equal(r.Equality.OID):
			return true
		case tv.Equal(r.Ordering.OID):
			return true
		case tv.Equal(r.Substring.OID):
			return true
		}
	case *LDAPSyntax:
		if r.Syntax != nil {
			return r.Syntax.OID.Equal(tv.OID)
		} else if !r.SuperType.IsZero() {
			return r.SuperType.is(tv)
		}
	}

	return false
}

/*
getSyntax will traverse the supertype chain upwards until it finds an explicit SYNTAX definition
*/
func (r *AttributeType) getSyntax() *LDAPSyntax {
	if r.IsZero() {
		return nil
	}
	if r.Syntax.IsZero() {
		return r.SuperType.getSyntax()
	}

	return r.Syntax
}

/*
Validate returns an error that reflects any fatal condition observed regarding the receiver configuration.
*/
func (r *AttributeType) Validate() (err error) {
	return r.validate()
}

func (r *AttributeType) validate() (err error) {
	if r.IsZero() {
		return raise(isZero, "%T.validate", r)
	}

	if err = validateFlag(r.flags); err != nil {
		return
	}

	var ls *LDAPSyntax
	if ls, err = r.validateSyntax(); err != nil {
		return
	}

	if err = r.validateMatchingRules(ls); err != nil {
		return
	}

	if err = validateNames(r.Name.strings()...); err != nil {
		return
	}

	if err = validateDesc(r.Description); err != nil {
		return
	}

	if !r.SuperType.IsZero() {
		err = r.SuperType.Validate()
	} else {
		if r.Syntax.IsZero() {
			err = raise(invalidUnmarshal, "%T.unmarshal: %T.%T: %s (not sub-typed)",
				r, r, r.Syntax, isZero.Error())
		}
	}

	return
}

func (r *AttributeType) validateSyntax() (ls *LDAPSyntax, err error) {
	ls = r.getSyntax()
	if ls.IsZero() {
		err = raise(invalidSyntax,
			"checkMatchingRules: %T is missing a syntax", r)
	}

	return
}

func (r *AttributeType) validateMatchingRules(ls *LDAPSyntax) (err error) {
	if err = r.validateEquality(ls); err != nil {
		return err
	}

	if err = r.validateOrdering(ls); err != nil {
		return err
	}

	if err = r.validateSubstr(ls); err != nil {
		return err
	}

	return
}

func (r *AttributeType) validateEquality(ls *LDAPSyntax) error {
	if !r.Equality.IsZero() {
		if contains(toLower(r.Equality.Name.Index(0)), `ordering`) ||
			contains(toLower(r.Equality.Name.Index(0)), `substring`) {
			return raise(invalidMatchingRule,
				"validateEquality: %T.Equality uses non-equality %T syntax (%s)",
				r, r.Equality, r.Equality.Syntax.OID.String())
		}
	}

	return nil
}

func (r *AttributeType) validateSubstr(ls *LDAPSyntax) error {
	if !r.Substring.IsZero() {
		if !contains(toLower(r.Substring.Name.Index(0)), `substring`) {
			return raise(invalidMatchingRule,
				"validateSubstr: %T.Substring uses non-substring %T syntax (%s)",
				r, r.Substring, r.Substring.Syntax.OID.String())
		}
	}

	return nil
}

func (r *AttributeType) validateOrdering(ls *LDAPSyntax) error {
	if !r.Ordering.IsZero() {
		if !contains(toLower(r.Ordering.Name.Index(0)), `ordering`) {
			return raise(invalidMatchingRule,
				"validateOrdering: %T.Ordering uses non-substring %T syntax (%s)",
				r, r.Ordering, r.Ordering.Syntax.OID.String())
		}
	}

	return nil
}

/*
SetUnmarshaler assigns the provided DefinitionUnmarshaler signature value to the receiver. The provided function shall be executed during the unmarshal or unsafe stringification process.
*/
func (r *AttributeType) SetUnmarshaler(fn DefinitionUnmarshaler) {
	r.ufn = fn
}

/*
SetInfo assigns the byte slice to the receiver. This is a user-leveraged field intended to allow arbitrary information (documentation?) to be assigned to the definition.
*/
func (r *AttributeType) SetInfo(info []byte) {
	r.info = info
}

/*
Info returns the assigned informational byte slice instance stored within the receiver.
*/
func (r *AttributeType) Info() []byte {
	return r.info
}

/*
Map is a convenience method that returns a map[string][]string instance containing the effective contents of the receiver.
*/
func (r *AttributeType) Map() (def map[string][]string) {
	if err := r.Validate(); err != nil {
		return
	}

	def = make(map[string][]string, 14)
	def[`RAW`] = []string{r.String()}
	def[`OID`] = []string{r.OID.String()}
	def[`TYPE`] = []string{r.Type()}

	if len(r.info) > 0 {
		def[`INFO`] = []string{string(r.info)}
	}

	if !r.Name.IsZero() {
		def[`NAME`] = make([]string, 0)
		for i := 0; i < r.Name.Len(); i++ {
			def[`NAME`] = append(def[`NAME`], r.Name.Index(i))
		}
	}

	if r.Usage != UserApplication {
		def[`USAGE`] = []string{r.Usage.String()}
	}

	if len(r.Description) > 0 {
		def[`DESC`] = []string{r.Description.String()}
	}

	if !r.Syntax.IsZero() {
		syn := r.Syntax.OID.String()
		def[`SYNTAX`] = []string{syn}
		if r.MaxLength() > 0 {
			def[`MUB`] = []string{itoa(r.MaxLength())}
		}
	}

	if !r.Equality.IsZero() {
		term := r.Equality.Name.Index(0)
		if len(term) == 0 {
			term = r.Equality.OID.String()
		}
		def[`EQUALITY`] = []string{term}
	}

	if !r.Substring.IsZero() {
		term := r.Substring.Name.Index(0)
		if len(term) == 0 {
			term = r.Substring.OID.String()
		}
		def[`SUBSTR`] = []string{term}
	}

	if !r.Ordering.IsZero() {
		term := r.Ordering.Name.Index(0)
		if len(term) == 0 {
			term = r.Ordering.OID.String()
		}
		def[`ORDERING`] = []string{term}
	}

	if !r.SuperType.IsZero() {
		term := r.SuperType.Name.Index(0)
		if len(term) == 0 {
			term = r.SuperType.OID.String()
		}
		def[`SUP`] = []string{term}
	}

	if !r.Extensions.IsZero() {
		for i := 0; i < r.Extensions.Len(); i++ {
			ext := r.Extensions.Index(i)
			def[ext.Label] = ext.Value
		}
	}

	if r.Obsolete {
		def[`OBSOLETE`] = []string{`TRUE`}
	}

	if r.Collective() {
		def[`COLLECTIVE`] = []string{`TRUE`}
	}

	if r.NoUserModification() {
		def[`NO-USER-MODIFICATION`] = []string{`TRUE`}
	}

	if r.SingleValue() {
		def[`SINGLE-VALUE`] = []string{`TRUE`}
	}

	return def
}

/*
AttributeTypeUnmarshaler is a package-included function that honors the signature of the first class (closure) DefinitionUnmarshaler type.

The purpose of this function, and similar user-devised ones, is to unmarshal a definition with specific formatting included, such as linebreaks, leading specifier declarations and indenting.
*/
func AttributeTypeUnmarshaler(x interface{}) (def string, err error) {
	var r *AttributeType
	switch tv := x.(type) {
	case *AttributeType:
		if tv.IsZero() {
			err = raise(isZero, "%T is nil", tv)
			return
		}
		r = tv
	default:
		err = raise(unexpectedType,
			"Bad type for unmarshal (%T)", tv)
		return
	}

	var (
		WHSP string = ` `
		idnt string = "\n\t"
		head string = `(`
		tail string = `)`
	)

	if len(r.spec) > 0 {
		head = r.spec + WHSP + head
	}

	def += head + WHSP + r.OID.String()

	if !r.Name.IsZero() {
		def += idnt + r.Name.Label()
		def += WHSP + r.Name.String()
	}

	if !r.Description.IsZero() {
		def += idnt + r.Description.Label()
		def += WHSP + r.Description.String()
	}

	if r.Obsolete {
		def += idnt + `OBSOLETE`
	}

	if !r.SuperType.IsZero() {
		def += idnt + r.SuperType.Label()
		def += WHSP + r.SuperType.Name.Index(0)
	}

	if !r.Equality.IsZero() {
		def += idnt + r.Equality.Label()
		def += WHSP + r.Equality.Name.Index(0)
	}

	if !r.Ordering.IsZero() {
		def += idnt + r.Ordering.Label()
		def += WHSP + r.Ordering.Name.Index(0)
	}

	if !r.Substring.IsZero() {
		def += idnt + r.Substring.Label()
		def += WHSP + r.Substring.Name.Index(0)
	}

	if !r.Syntax.IsZero() {
		def += idnt + r.Syntax.Label()
		def += WHSP + r.Syntax.OID.String()
		if r.MaxLength() > 0 {
			def += `{` + itoa(r.MaxLength()) + `}`
		}
	}

	if r.SingleValue() {
		def += idnt + SingleValue.String()
	}

	if r.Collective() {
		def += idnt + Collective.String()
	}

	if r.NoUserModification() {
		def += idnt + NoUserModification.String()
	}

	if r.Usage != UserApplication {
		def += idnt + r.Usage.Label()
		def += WHSP + r.Usage.String()
	}

	for i := 0; i < r.Extensions.Len(); i++ {
		if ext := r.Extensions.Index(i); !ext.IsZero() {
			def += idnt + ext.String()
		}
	}

	def += WHSP + tail

	return
}

func (r *AttributeType) unmarshal() (string, error) {
	if err := r.validate(); err != nil {
		err = raise(invalidUnmarshal, err.Error())
		return ``, err
	}

	if r.ufn != nil {
		return r.ufn(r)
	}
	return r.unmarshalBasic()
}

func (r *AttributeType) unmarshalBasic() (def string, err error) {
	var (
		WHSP string = ` `
		head string = `(`
		tail string = `)`
	)

	if len(r.spec) > 0 {
		head = r.spec + WHSP + head
	}

	def += head + WHSP + r.OID.String()

	if !r.Name.IsZero() {
		def += WHSP + r.Name.Label()
		def += WHSP + r.Name.String()
	}

	if !r.Description.IsZero() {
		def += WHSP + r.Description.Label()
		def += WHSP + r.Description.String()
	}

	if r.Obsolete {
		def += WHSP + `OBSOLETE`
	}

	if !r.SuperType.IsZero() {
		def += WHSP + r.SuperType.Label()
		def += WHSP + r.SuperType.Name.Index(0)
	}

	if !r.Equality.IsZero() {
		def += WHSP + r.Equality.Label()
		def += WHSP + r.Equality.Name.Index(0)
	}

	if !r.Ordering.IsZero() {
		def += WHSP + r.Ordering.Label()
		def += WHSP + r.Ordering.Name.Index(0)
	}

	if !r.Substring.IsZero() {
		def += WHSP + r.Substring.Label()
		def += WHSP + r.Substring.Name.Index(0)
	}

	if !r.Syntax.IsZero() {
		def += WHSP + r.Syntax.Label()
		def += WHSP + r.Syntax.OID.String()
		if r.MaxLength() > 0 {
			def += `{` + itoa(r.MaxLength()) + `}`
		}
	}

	if r.SingleValue() {
		def += WHSP + SingleValue.String()
	}

	if r.Collective() {
		def += WHSP + Collective.String()
	}

	if r.NoUserModification() {
		def += WHSP + NoUserModification.String()
	}

	if r.Usage != UserApplication {
		def += WHSP + r.Usage.Label()
		def += WHSP + r.Usage.String()
	}

	if !r.Extensions.IsZero() {
		def += WHSP + r.Extensions.String()
	}

	def += WHSP + tail

	return
}

/*
atFlags is an unsigned 8-bit integer that describes zero or more perceived values that only appear visually when TRUE. Such verisimilitude is revealed by the presence of the indicated atFlags value's "label" name, such as `SINGLE-VALUE`.  The actual value "TRUE" is never actually seen in textual format.
*/
type atFlags uint8

const (
	SingleValue        atFlags = 1 << iota // 1
	Collective                             // 2
	NoUserModification                     // 4
)

func (r atFlags) IsZero() bool {
	return uint8(r) == 0
}

/*
is returns a boolean value indicative of whether the specified value is enabled within the receiver instance.
*/
func (r atFlags) is(o atFlags) bool {
	return r.enabled(o)
}

/*
String is a stringer method that returns the name(s) atFlags receiver in question, whether it represents multiple boolean flags or only one.

If only a specific atFlags string is desired (if enabled), use atFlags.<Name>() (e.g: atFlags.SingleValue()).
*/
func (r atFlags) String() (val string) {

	// Look for so-called "pure"
	// boolean values first ...
	switch r {
	case NoUserModification:
		return `NO-USER-MODIFICATION`
	case SingleValue:
		return `SINGLE-VALUE`
	case Collective:
		return `COLLECTIVE`
	}

	// Assume multiple boolean bits
	// are set concurrently ...
	strs := []atFlags{
		Collective,
		SingleValue,
		NoUserModification,
	}

	vals := make([]string, 0)
	for _, v := range strs {
		if r.enabled(v) {
			vals = append(vals, v.String())
		}
	}

	if len(vals) != 0 {
		val = join(vals, ` `)
	}

	return
}

/*
SingleValue returns the `SINGLE-VALUE` flag if the appropriate bits are set within the receiver instance.
*/
func (r atFlags) SingleValue() string {
	if r.enabled(SingleValue) {
		return r.String()
	}

	return ``
}

/*
Collective returns the `COLLECTIVE` flag if the appropriate bits are set within the receiver instance.
*/
func (r atFlags) Collective() string {
	if r.enabled(Collective) {
		return r.String()
	}

	return ``
}

/*
NoUserModification returns the `NO-USER-MODIFICATION` flag if the appropriate bits are set within the receiver instance.
*/
func (r atFlags) NoUserModification() string {
	if r.enabled(NoUserModification) {
		return r.String()
	}

	return ``
}

/*
Unset removes the specified atFlags bits from the receiver instance of atFlags, thereby "disabling" the provided option.
*/
func (r *atFlags) Unset(o atFlags) {
	r.unset(o)
}

/*
Set adds the specified atFlags bits to the receiver instance of atFlags, thereby "enabling" the provided option.
*/
func (r *atFlags) Set(o atFlags) {
	r.set(o)
}

func (r *atFlags) set(o atFlags) {
	*r = *r ^ o
}

func (b *atFlags) unset(o atFlags) {
	*b = *b &^ o
}

func (r atFlags) enabled(o atFlags) bool {
	return r&o != 0
}

/*
setatFlags is a private method used by reflect to set boolean values.
*/
func (r *AttributeType) setATFlags(b atFlags) {
	r.flags.set(b)
}

/*
SetCollective marks the receiver as COLLECTIVE.
*/
func (r *AttributeType) SetCollective() {
	r.flags.set(Collective)
}

/*
Collective returns a boolean value indicative of whether the receiver describes a COLLECTIVE attribute type.
*/
func (r *AttributeType) Collective() bool {
	return r.flags.is(Collective)
}

/*
SetCollective marks the receiver as NO-USER-MODIFICATION.
*/
func (r *AttributeType) SetNoUserModification() {
	r.flags.set(NoUserModification)
}

/*
NoUserModification returns a boolean value indicative of whether the receiver describes a NO-USER-MODIFICATION attribute type.
*/
func (r *AttributeType) NoUserModification() bool {
	return r.flags.is(NoUserModification)
}

/*
SetSingleValue marks the receiver as SINGLE-VALUE.
*/
func (r *AttributeType) SetSingleValue() {
	r.flags.set(SingleValue)
}

/*
SingleValue returns a boolean value indicative of whether the receiver describes a SINGLE-VALUE attribute type.
*/
func (r *AttributeType) SingleValue() bool {
	return r.flags.is(SingleValue)
}

func validateFlag(b atFlags) (err error) {
	if b.is(Collective) && b.is(SingleValue) {
		return raise(invalidFlag,
			"Cannot have single-valued collective attribute")
	}

	return
}

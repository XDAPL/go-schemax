package schemax

import "sync"

/*
LDAPSyntaxTypeCollection describes all LDAPSyntax-based types.
*/
type LDAPSyntaxCollection interface {
	// Get returns the *LDAPSyntax instance retrieved as a result
	// of a term search, based on Name or OID. If no match is found,
	// nil is returned.
	Get(interface{}) *LDAPSyntax

	// Index returns the *LDAPSyntax instance stored at the nth
	// index within the receiver, or nil.
	Index(int) *LDAPSyntax

	// Equal performs a deep-equal between the receiver and the
	// interface LDAPSyntaxCollection provided.
	Equal(LDAPSyntaxCollection) bool

	// Set returns an error instance based on an attempt to add
	// the provided *AttributeType instance to the receiver.
	Set(*LDAPSyntax) error

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
LDAPSyntax conforms to the specifications of RFC4512 Section 4.1.5.
*/
type LDAPSyntax struct {
	OID         OID
	Description Description
	Extensions  *Extensions
	ufn         DefinitionUnmarshaler
	spec        string
	info        []byte
}

/*
LDAPSyntaxes is a thread-safe collection of *LDAPSyntax slice instances.
*/
type LDAPSyntaxes struct {
	mutex  *sync.Mutex
	slice  collection
	macros *Macros
}

/*
Type returns the formal name of the receiver in order to satisfy signature requirements of the Definition interface type.
*/
func (r *LDAPSyntax) Type() string {
	return `LDAPSyntax`
}

/*
Equal performs a deep-equal between the receiver and the provided collection type.
*/
func (r LDAPSyntaxes) Equal(x LDAPSyntaxCollection) bool {
	return r.slice.equal(x.(*LDAPSyntaxes).slice)
}

/*
SetMacros assigns the *Macros instance to the receiver, allowing subsequent OID resolution capabilities during the addition of new slice elements.
*/
func (r *LDAPSyntaxes) SetMacros(macros *Macros) {
	r.macros = macros
}

/*
SetSpecifier is a convenience method that executes the SetSpecifier method in iterative fashion for all definitions within the receiver.
*/
func (r *LDAPSyntaxes) SetSpecifier(spec string) {
	for i := 0; i < r.Len(); i++ {
		r.Index(i).SetSpecifier(spec)
	}
}

/*
SetUnmarshaler is a convenience method that executes the SetUnmarshaler method in iterative fashion for all definitions within the receiver.
*/
func (r *LDAPSyntaxes) SetUnmarshaler(fn DefinitionUnmarshaler) {
	for i := 0; i < r.Len(); i++ {
		r.Index(i).SetUnmarshaler(fn)
	}
}

/*
Contains is a thread-safe method that returns a collection slice element index integer and a presence-indicative boolean value based on a term search conducted within the receiver.
*/
func (r LDAPSyntaxes) Contains(x interface{}) (int, bool) {
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
func (r LDAPSyntaxes) Index(idx int) *LDAPSyntax {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	assert, _ := r.slice.index(idx).(*LDAPSyntax)
	return assert
}

/*
Get combines Contains and Index method executions to return an entry based on a term search conducted within the receiver.
*/
func (r LDAPSyntaxes) Get(x interface{}) *LDAPSyntax {
	idx, found := r.Contains(x)
	if !found {
		return nil
	}

	return r.Index(idx)
}

/*
Len is a thread-safe method that returns the effective length of the receiver slice collection.
*/
func (r LDAPSyntaxes) Len() int {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	return r.slice.len()
}

/*
String is a non-functional stringer method needed to satisfy interface type requirements and should not be used. There is no practical application for a list of ldapSyntax object identifiers in this package.
*/
func (r LDAPSyntaxes) String() string { return `` }

/*
String is an unsafe convenience wrapper for Unmarshal(r). If an error is encountered, an empty string definition is returned. If reliability and error handling are important, use Unmarshal.
*/
func (r LDAPSyntax) String() (def string) {
	def, _ = r.unmarshal()
	return
}

/*
IsZero returns a boolean value indicative of whether the receiver is considered empty or uninitialized.
*/
func (r LDAPSyntaxes) IsZero() bool {
	return r.slice.len() == 0
}

/*
IsZero returns a boolean value indicative of whether the receiver is considered empty or uninitialized.
*/
func (r *LDAPSyntax) IsZero() bool {
	return r == nil
}

/*
Set is a thread-safe append method that returns an error instance indicative of whether the append operation failed in some manner. Uniqueness is enforced for new elements based on Object Identifier and not the effective Name of the definition, if defined.
*/
func (r *LDAPSyntaxes) Set(x *LDAPSyntax) error {
	if _, exists := r.Contains(x.OID); exists {
		return nil
	}

	r.mutex.Lock()
	defer r.mutex.Unlock()

	return r.slice.append(x)
}

/*
SetSpecifier assigns a string value to the receiver, useful for placement into configurations that require a type name (e.g.: attributetype). This will be displayed at the beginning of the definition value during the unmarshal or unsafe stringification process.
*/
func (r *LDAPSyntax) SetSpecifier(spec string) {
	r.spec = spec
}

/*
SetInfo assigns the byte slice to the receiver. This is a user-leveraged field intended to allow arbitrary information (documentation?) to be assigned to the definition.
*/
func (r *LDAPSyntax) SetInfo(info []byte) {
	r.info = info
}

/*
Info returns the assigned informational byte slice instance stored within the receiver.
*/
func (r *LDAPSyntax) Info() []byte {
	return r.info
}

/*
SetUnmarshaler assigns the provided DefinitionUnmarshaler signature value to the receiver. The provided function shall be executed during the unmarshal or unsafe stringification process.
*/
func (r *LDAPSyntax) SetUnmarshaler(fn DefinitionUnmarshaler) {
	r.ufn = fn
}

/*
NewLDAPSyntax returns a newly initialized, yet effectively nil, instance of *LDAPSyntax.

Users generally do not need to execute this function unless an instance of the returned type will be manually populated (as opposed to parsing a raw text definition).
*/
func NewLDAPSyntax() *LDAPSyntax {
	ls := new(LDAPSyntax)
	ls.Extensions = NewExtensions()
	return ls
}

/*
NewLDAPSyntaxes initializes and returns a new LDAPSyntaxesCollection interface object.
*/
func NewLDAPSyntaxes() LDAPSyntaxCollection {
	var x interface{} = &LDAPSyntaxes{
		mutex: &sync.Mutex{},
		slice: make(collection, 0, 0),
	}
	return x.(LDAPSyntaxCollection)
}

/*
HumanReadable is a convenience wrapper for Extensions.HumanReadable().
*/
func (r *LDAPSyntax) HumanReadable() bool {
	if !r.Extensions.IsZero() {
		return r.Extensions.HumanReadable()
	}
	return true
}

/*
SetHumanReadable sets the LDAPSyntax Extension field `X-NOT-HUMAN-READABLE` as `TRUE` when given a value of false, and deletes said field when given a value of true.
*/
func (r *LDAPSyntax) SetHumanReadable(x bool) {
	r.Extensions.SetHumanReadable(x)
}

/*
Equal performs a deep-equal between the receiver and the provided definition type.

Description text is ignored.
*/
func (r *LDAPSyntax) Equal(x interface{}) (eq bool) {
	z, ok := x.(*LDAPSyntax)
	if !ok {
		return
	}

	if z.OID.String() != r.OID.String() {
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
Validate returns an error that reflects any fatal condition observed regarding the receiver configuration.
*/
func (r *LDAPSyntax) Validate() (err error) {
	return r.validate()
}

func (r *LDAPSyntax) validate() (err error) {
	if r.IsZero() {
		return raise(isZero, "%T.validate", r)
	}

	if err = validateDesc(r.Description); err != nil {
		return
	}

	return
}

func (r *LDAPSyntax) unmarshal() (string, error) {
	if err := r.validate(); err != nil {
		err = raise(invalidUnmarshal, err.Error())
		return ``, err
	}

	if r.ufn != nil {
		return r.ufn(r)
	}
	return r.unmarshalBasic()
}

/*
LDAPSyntaxUnmarshaler is a package-included function that honors the signature of the first class (closure) DefinitionUnmarshaler type.

The purpose of this function, and similar user-devised ones, is to unmarshal a definition with specific formatting included, such as linebreaks, leading specifier declarations and indenting.
*/
func LDAPSyntaxUnmarshaler(x interface{}) (def string, err error) {
	var r *LDAPSyntax
	switch tv := x.(type) {
	case *LDAPSyntax:
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

	// Description will never be zero
	def += idnt + r.Description.Label()
	def += WHSP + r.Description.String()

	if !r.Extensions.IsZero() {
		for i := 0; i < r.Extensions.Len(); i++ {
			if ext := r.Extensions.Index(i); !ext.IsZero() {
				def += idnt + ext.String()
			}
		}
	}

	def += WHSP + tail

	return
}

/*
Map is a convenience method that returns a map[string][]string instance containing the effective contents of the receiver.
*/
func (r *LDAPSyntax) Map() (def map[string][]string) {
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

	if len(r.Description) > 0 {
		def[`DESC`] = []string{r.Description.String()}
	}

	for i := 0; i < r.Extensions.Len(); i++ {
		ext := r.Extensions.Index(i)
		def[ext.Label] = ext.Value
	}

	return def
}

func (r *LDAPSyntax) unmarshalBasic() (def string, err error) {

	var (
		WHSP string = ` `
		head string = `(`
		tail string = `)`
	)

	if len(r.spec) > 0 {
		head = r.spec + WHSP + head
	}

	def += head + WHSP + r.OID.String()

	// Description will never be zero
	def += WHSP + r.Description.Label()
	def += WHSP + r.Description.String()

	if !r.Extensions.IsZero() {
		def += WHSP + r.Extensions.String()
	}

	def += WHSP + tail

	return
}

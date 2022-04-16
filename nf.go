package schemax

import (
	"fmt"
	"sync"
)

/*
NameFormCollection describes all NameForms-based types.
*/
type NameFormCollection interface {
	// Contains returns the index number and presence boolean that
	// reflects the result of a term search within the receiver.
        Contains(interface{}) (int, bool)

	// // Get returns the *NameForm instance retrieved as a result
	// of a term search, based on Name or OID. If no match is found,
	// nil is returned.
        Get(interface{}) *NameForm

	// Index returns the *NameForm instance stored at the nth
	// index within the receiver, or nil.
        Index(int) *NameForm

	// Equal performs a deep-equal between the receiver and the
	// interface NameFormCollection provided.
        Equal(NameFormCollection) bool

	// Set returns an error instance based on an attempt to add
	// the provided *NameForm instance to the receiver.
        Set(*NameForm) error

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
}

/*
NameForm conforms to the specifications of RFC4512 Section 4.1.7.2. Boolean values, e.g: 'OBSOLETE', are supported internally and are not explicit fields.
*/
type NameForm struct {
	OID         OID
	Name        Name
	Description Description
	OC          StructuralObjectClass
	Must        AttributeTypeCollection
	May         AttributeTypeCollection
	Extensions  Extensions
	bools       Boolean
}

/*
NameForms is a thread-safe collection of *NameForm slice instances.
*/
type NameForms struct {
	mutex  *sync.Mutex
	slice  collection
	macros *Macros
}

/*
Equal performs a deep-equal between the receiver and the provided collection type.
*/
func (r NameForms) Equal(x NameFormCollection) bool {
        return r.slice.equal(x.(*NameForms).slice)
}

/*
SetMacros assigns the *Macros instance to the receiver, allowing subsequent OID resolution capabilities during the addition of new slice elements.
*/
func (r *NameForms) SetMacros(macros *Macros) {
	r.macros = macros
}

/*
Contains is a thread-safe method that returns a collection slice element index integer and a presence-indicative boolean value based on a term search conducted within the receiver.
*/
func (r NameForms) Contains(x interface{}) (int, bool) {
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
func (r NameForms) Index(idx int) *NameForm {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	assert, _ := r.slice.index(idx).(*NameForm)
	return assert
}

/*
Get combines Contains and Index method executions to return an entry based on a term search conducted within the receiver.
*/
func (r NameForms) Get(x interface{}) *NameForm {
	idx, found := r.Contains(x)
	if !found {
		return nil
	}

	return r.Index(idx)
}

/*
Len is a thread-safe method that returns the effective length of the receiver slice collection.
*/
func (r NameForms) Len() int {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	return r.slice.len()
}

/*
String is a stringer method used to return the properly-delimited and formatted series of attributeType name or OID definitions.
*/
func (r NameForms) String() string { return `` }

/*
String is an unsafe convenience wrapper for Unmarshal(r). If an error is encountered, an empty string definition is returned. If reliability and error handling are important, use Unmarshal.
*/
func (r NameForm) String() (def string) {
	def, _ = Unmarshal(r)
	return
}

/*
IsZero returns a boolean value indicative of whether the receiver is considered empty or uninitialized.
*/
func (r NameForms) IsZero() bool {
	return r.slice.len() == 0
}

/*
IsZero returns a boolean value indicative of whether the receiver is considered empty or uninitialized.
*/
func (r *NameForm) IsZero() bool {
	return r == nil
}

/*
Set is a thread-safe append method that returns an error instance indicative of whether the append operation failed in some manner. Uniqueness is enforced for new elements based on Object Identifier and not the effective Name of the definition, if defined.
*/
func (r *NameForms) Set(x *NameForm) error {
	if _, exists := r.Contains(x.OID); exists {
		return fmt.Errorf("%T already contains %T:%s", r, x, x.OID)
	}

	r.mutex.Lock()
	defer r.mutex.Unlock()

	return r.slice.append(x)
}

/*
Equal performs a deep-equal between the receiver and the provided definition type.

Description text is ignored.
*/
func (r *NameForm) Equal(x interface{}) (equals bool) {

	z, ok := x.(*NameForm)
	if !ok {
		return
	}

	if z.IsZero() && r.IsZero() {
		equals = true
		return
	} else if z.IsZero() || r.IsZero() {
		return
	}

	if !r.OID.Equal(z.OID) {
		return
	}

	if !r.Name.Equal(z.Name) {
		return
	}

	if !r.OC.Equal(z.OC) {
		return
	}

	if !r.Must.Equal(z.Must) {
		return
	}

	if !r.May.Equal(z.May) {
		return
	}

	equals = r.Extensions.Equal(z.Extensions)

	return
}

/*
NewNameForms initializes and returns a new NameFormCollection interface object.
*/
func NewNameForms() NameFormCollection {
	var x interface{} = &NameForms{
		mutex: &sync.Mutex{},
		slice: make(collection, 0, 0),
	}
	return x.(NameFormCollection)
}

/*
Validate returns an error that reflects any fatal condition observed regarding the receiver configuration.
*/
func (r *NameForm) Validate() (err error) {
	return r.validate()
}

func (r *NameForm) validate() (err error) {
	if r.IsZero() {
		return raise(isZero, "%T.validate", r)
	}

	if err = validateBool(r.bools); err != nil {
		return
	}

	if err = validateNames(r.Name.strings()...); err != nil {
		return
	}

	if err = r.validateMustMay(); err != nil {
		return
	}

	if err = r.validateStructuralObjectClass(); err != nil {
		return
	}

	if err = validateDesc(r.Description); err != nil {
		return
	}

	return
}

func (r *NameForm) validateMustMay() (err error) {
	if r.Must.IsZero() {
		err = raise(invalidNameForm,
			"validateMustMay: missing %T.Must (%T)",
			r, r.Must)
	}

	return
}

func (r *NameForm) validateStructuralObjectClass() (err error) {
	if r.OC.IsZero() {
		err = raise(invalidNameForm,
			"validateOC: Missing %T", r.OC)
	} else if !r.OC.Kind.is(Structural) {
		err = raise(invalidNameForm,
			"validateOC: %T lacks an embedded STRUCTURAL type (is '%s')",
			r.OC, r.OC.Kind.String())
	}

	if err == nil {
		err = r.OC.validate() // generalized validation
	}

	return
}

func (r *NameForm) unmarshal(namesok bool) (def string, err error) {
	if err = r.validate(); err != nil {
		err = raise(invalidUnmarshal, err.Error())
		return
	}

	WHSP := ` `

	def += `(` + WHSP + r.OID.String() // will never be zero length

	if !r.Name.IsZero() {
		def += WHSP + r.Name.Label()
		def += WHSP + r.Name.String()
	}

	if !r.Description.IsZero() {
		def += WHSP + r.Description.Label()
		def += WHSP + r.Description.String()
	}

	// OC will never be zero
	def += WHSP + r.OC.Label()
	def += WHSP + r.OC.String()

	// Must will never be zero
	def += WHSP + r.Must.Label()
	def += WHSP + r.Must.String()

	if !r.May.IsZero() {
		def += WHSP + r.May.Label()
		def += WHSP + r.May.String()
	}

	if r.bools.enabled(Obsolete) {
		def += WHSP + r.bools.Obsolete()
	}

	if !r.Extensions.IsZero() {
		def += WHSP + r.Extensions.String()
	}

	def += WHSP + `)`

	return
}

package schemax

import (
	"fmt"
	"sync"
)

/*
LDAPSyntaxTypeCollection describes all LDAPSyntax-based types.
*/
type LDAPSyntaxCollection interface {
	// Contains returns the index number and presence boolean that
	// reflects the result of a term search within the receiver.
        Contains(interface{}) (int, bool)

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
LDAPSyntax conforms to the specifications of RFC4512 Section 4.1.5. Internal Boolean support is managed only for a declaration of "human-readable" status, per RFC2252 Section 4.3.2.
*/
type LDAPSyntax struct {
	OID         OID
	Description Description
	Extensions  Extensions
	bools       Boolean
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
String is a stringer method used to return the properly-delimited and formatted series of attributeType name or OID definitions.
*/
func (r LDAPSyntaxes) String() string { return `` }

/*
String is an unsafe convenience wrapper for Unmarshal(r). If an error is encountered, an empty string definition is returned. If reliability and error handling are important, use Unmarshal.
*/
func (r LDAPSyntax) String() (def string) {
	def, _ = Unmarshal(r)
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
		return fmt.Errorf("%T already contains %T:%s", r, x, x.OID)
	}

	r.mutex.Lock()
	defer r.mutex.Unlock()

	return r.slice.append(x)
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
IsHumanReadable returns a boolean value indicative of whether the receiver instance of *LDAPSyntax supports values that are human-readable.
*/
func (r LDAPSyntax) IsHumanReadable() bool {
	return r.bools.is(HumanReadable)
}

/*
is returns a boolean value indicative of whether the provided interface argument is an enabled Boolean option.
*/
func (r LDAPSyntax) is(b interface{}) bool {
	switch tv := b.(type) {
	case Boolean:
		return r.bools.is(tv)
	}

	return false
}

/*
Equal performs a deep-equal between the receiver and the provided definition type.

Description text is ignored.
*/
func (r *LDAPSyntax) Equal(x interface{}) bool {
	if assert, equals := x.(*LDAPSyntax); equals {
		return assert.OID.String() == r.OID.String()
	}

	return false
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

func (r *LDAPSyntax) unmarshal(namesok bool) (def string, err error) {
	if err = r.validate(); err != nil {
		err = raise(invalidUnmarshal, err.Error())
		return
	}

	WHSP := ` `

	def += `(` + WHSP + r.OID.String() // will never be zero length

	// Description will never be zero
	def += WHSP + r.Description.Label()
	def += WHSP + r.Description.String()

	if r.bools.enabled(Obsolete) {
		def += WHSP + r.bools.Obsolete()
	}

	if !r.Extensions.IsZero() {
		def += WHSP + r.Extensions.String()
	}

	def += WHSP + `)`

	return
}

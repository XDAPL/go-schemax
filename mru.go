package schemax

import (
	"fmt"
	"sync"
)

/*
MatchingRuleUseCollection describes all MatchingRuleUses-based types.
*/
type MatchingRuleUseCollection interface {
	// Contains returns the index number and presence boolean that
	// reflects the result of a term search within the receiver.
        Contains(interface{}) (int, bool)

	// Get returns the *MatchingRuleUse instance retrieved as a result
	// of a term search, based on Name or OID. If no match is found,
	// nil is returned.
        Get(interface{}) *MatchingRuleUse

	// Index returns the *MatchingRuleUse instance stored at the nth
	// index within the receiver, or nil.
        Index(int) *MatchingRuleUse

	// Equal performs a deep-equal between the receiver and the
	// interface MatchingRuleUseCollection provided.
        Equal(MatchingRuleUseCollection) bool

	// Set returns an error instance based on an attempt to add
	// the provided *MatchingRuleUse instance to the receiver.
        Set(*MatchingRuleUse) error

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
MatchingRuleUse conforms to the specifications of RFC4512 Section 4.1.4. Boolean values, e.g: 'OBSOLETE', are supported internally and are not explicit fields.
*/
type MatchingRuleUse struct {
	OID         OID
	Name        Name
	Description Description
	Applies     ApplicableAttributeTypes
	Extensions  Extensions
	bools       Boolean
}

/*
MatchingRuleUses is a thread-safe collection of *MatchingRuleUse slice instances.
*/
type MatchingRuleUses struct {
	mutex *sync.Mutex
	slice collection
}

/*
Equal performs a deep-equal between the receiver and the provided collection type.
*/
func (r MatchingRuleUses) Equal(x MatchingRuleUseCollection) bool {
        return r.slice.equal(x.(*MatchingRuleUses).slice)
}

/*
Contains is a thread-safe method that returns a collection slice element index integer and a presence-indicative boolean value based on a term search conducted within the receiver.
*/
func (r MatchingRuleUses) Contains(x interface{}) (int, bool) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	return r.slice.contains(x)
}

/*
Index is a thread-safe method that returns the nth collection slice element if defined, else nil. This method supports use of negative indices which should be used with special care.
*/
func (r MatchingRuleUses) Index(idx int) *MatchingRuleUse {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	assert, _ := r.slice.index(idx).(*MatchingRuleUse)
	return assert
}

/*
Get combines Contains and Index method executions to return an entry based on a term search conducted within the receiver.
*/
func (r MatchingRuleUses) Get(x interface{}) *MatchingRuleUse {
	idx, found := r.Contains(x)
	if !found {
		return nil
	}

	return r.Index(idx)
}

/*
Len is a thread-safe method that returns the effective length of the receiver slice collection.
*/
func (r MatchingRuleUses) Len() int {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	return r.slice.len()
}

/*
String is a stringer method used to return the properly-delimited and formatted series of attributeType name or OID definitions.
*/
func (r MatchingRuleUses) String() string { return `` }

/*
String is an unsafe convenience wrapper for Unmarshal(r). If an error is encountered, an empty string definition is returned. If reliability and error handling are important, use Unmarshal.
*/
func (r MatchingRuleUse) String() (def string) {
	def, _ = Unmarshal(r)
	return
}

/*
IsZero returns a boolean value indicative of whether the receiver is considered empty or uninitialized.
*/
func (r MatchingRuleUses) IsZero() bool {
	return r.slice.len() == 0
}

/*
IsZero returns a boolean value indicative of whether the receiver is considered empty or uninitialized.
*/
func (r *MatchingRuleUse) IsZero() bool {
	return r == nil
}

/*
Set is a thread-safe append method that returns an error instance indicative of whether the append operation failed in some manner. Uniqueness is enforced for new elements based on Object Identifier and not the effective Name of the definition, if defined.
*/
func (r *MatchingRuleUses) Set(x *MatchingRuleUse) error {
	if _, exists := r.Contains(x.OID); exists {
		return fmt.Errorf("%T already contains %T:%s", r, x, x.OID)
	}

	r.mutex.Lock()
	defer r.mutex.Unlock()

	return r.slice.append(x)
}

/*
Equal performs a deep-equal between the receiver and the provided collection type.

Description text is ignored.
*/
func (r *MatchingRuleUse) Equal(x interface{}) (equals bool) {
	z, ok := x.(MatchingRuleUse)
	if !ok {
		return
	}

	if z.IsZero() && r.IsZero() {
		equals = true
		return
	} else if z.IsZero() || r.IsZero() {
		return
	}

	if !z.OID.Equal(r.OID) {
		return
	}

	if !z.Name.Equal(r.Name) {
		return
	}

	if z.bools != r.bools {
		return
	}

	if !z.Applies.Equal(r.Applies) {
		return
	}
	equals = r.Extensions.Equal(z.Extensions)

	return
}

/*
NewMatchingRuleUses initializes and returns a new MatchingRuleUsesCollection interface object.
*/
func NewMatchingRuleUses() MatchingRuleUseCollection {
	var x interface{} = &MatchingRuleUses{
		mutex: &sync.Mutex{},
		slice: make(collection, 0, 0),
	}
	return x.(MatchingRuleUseCollection)
}

/*
is returns a boolean value indicative of whether the provided interface argument is an enabled Boolean option.
*/
func (r MatchingRuleUse) is(b interface{}) bool {
	switch tv := b.(type) {
	case *AttributeType:
		if _, x := r.Applies.Contains(tv); !x {
			return false
		}
	case Boolean:
		return r.bools.is(tv)
	}

	return false
}

/*
Validate returns an error that reflects any fatal condition observed regarding the receiver configuration.
*/
func (r *MatchingRuleUse) Validate() (err error) {
	return r.validate()
}

func (r *MatchingRuleUse) validate() (err error) {
	if r.IsZero() {
		return raise(isZero, "%T.validate", r)
	}

	if r.OID.IsZero() {
		return raise(isZero, "%T.validate: no %T",
			r, r.OID)
	}

	if err = validateNames(r.Name.strings()...); err != nil {
		return
	}

	if err = validateDesc(r.Description); err != nil {
		return
	}

	if r.Applies.IsZero() {
		return raise(isZero, "%T.validate: no %T",
			r, r.Applies)
	}

	if err = validateBool(r.bools); err != nil {
		return
	}

	return
}

func (r *MatchingRuleUse) unmarshal(namesok bool) (def string, err error) {
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

	if r.is(Obsolete) {
		def += WHSP + r.bools.Obsolete()
	}

	// Applies will never be zero
	def += WHSP + r.Applies.Label()
	def += WHSP + r.Applies.String()

	if !r.Extensions.IsZero() {
		def += WHSP + r.Extensions.String()
	}

	def += WHSP + `)`

	return
}

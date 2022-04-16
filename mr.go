package schemax

import (
	"fmt"
	"sync"
)

/*
MatchingRuleCollection describes all MatchingRules-based types.
*/
type MatchingRuleCollection interface {
	// Contains returns the index number and presence boolean that
	// reflects the result of a term search within the receiver.
        Contains(interface{}) (int, bool)

	// Get returns the *MatchingRule instance retrieved as a result
	// of a term search, based on Name or OID. If no match is found,
	// nil is returned.
        Get(interface{}) *MatchingRule

	// Index returns the *MatchingRule instance stored at the nth
	// index within the receiver, or nil.
        Index(int) *MatchingRule

	// Equal performs a deep-equal between the receiver and the
	// interface MatchingRuleCollection provided.
        Equal(MatchingRuleCollection) bool

	// Set returns an error instance based on an attempt to add
	// the provided *MatchingRule instance to the receiver.
        Set(*MatchingRule) error

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
MatchingRule conforms to the specifications of RFC4512 Section 4.1.3. Boolean values, e.g: 'OBSOLETE', are supported internally and are not explicit fields.
*/
type MatchingRule struct {
	OID         OID
	Name        Name
	Description Description
	Syntax      *LDAPSyntax
	Extensions  Extensions
	bools       Boolean
}

/*
Equality circumscribes an embedded pointer to an instance of *MatchingRule.  This type alias is intended for use solely within instances of AttributeType via its "Equality" struct field.
*/
type Equality struct {
	*MatchingRule
}

/*
Equal performs a deep-equal between the receiver and the provided definition type.

Description text is ignored.
*/
func (r Equality) Equal(x interface{}) (equals bool) {
	z, ok := x.(Equality)
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

	if !z.Syntax.Equal(r.Syntax) {
		return
	}

	if z.bools != r.bools {
		return
	}
	equals = r.Extensions.Equal(z.Extensions)

	return
}

/*
Ordering circumscribes an embedded pointer to an instance of *MatchingRule.  This type alias is intended for use solely within instances of AttributeType via its "Ordering" struct field.
*/
type Ordering struct {
	*MatchingRule
}

/*
Equal performs a deep-equal between the receiver and the provided collection type.
*/
func (r Ordering) Equal(x interface{}) (equals bool) {
	z, ok := x.(Ordering)
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

	if !z.Syntax.Equal(r.Syntax) {
		return
	}

	if z.bools != r.bools {
		return
	}
	equals = r.Extensions.Equal(z.Extensions)

	return
}

/*
Substring circumscribes an embedded pointer to an instance of *MatchingRule.  This type alias is intended for use solely within instances of AttributeType via its "Substring" struct field.
*/
type Substring struct {
	*MatchingRule
}

/*
Equal performs a deep-equal between the receiver and the provided definition type.

Description text is ignored.
*/
func (r Substring) Equal(x interface{}) (equals bool) {
	z, ok := x.(Substring)
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

	if !z.Syntax.Equal(r.Syntax) {
		return
	}

	if z.bools != r.bools {
		return
	}
	equals = r.Extensions.Equal(z.Extensions)

	return
}

/*
MatchingRules is a thread-safe collection of *MatchingRule slice instances.
*/
type MatchingRules struct {
	mutex  *sync.Mutex
	slice  collection
	macros *Macros
}

/*
Equal performs a deep-equal between the receiver and the provided collection type.
*/
func (r MatchingRules) Equal(x MatchingRuleCollection) bool {
        return r.slice.equal(x.(*MatchingRules).slice)
}

/*
SetMacros assigns the *Macros instance to the receiver, allowing subsequent OID resolution capabilities during the addition of new slice elements.
*/
func (r *MatchingRules) SetMacros(macros *Macros) {
	r.macros = macros
}

/*
Contains is a thread-safe method that returns a collection slice element index integer and a presence-indicative boolean value based on a term search conducted within the receiver.
*/
func (r MatchingRules) Contains(x interface{}) (int, bool) {
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
func (r MatchingRules) Index(idx int) *MatchingRule {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	assert, _ := r.slice.index(idx).(*MatchingRule)
	return assert
}

/*
Get combines Contains and Index method executions to return an entry based on a term search conducted within the receiver.
*/
func (r MatchingRules) Get(x interface{}) *MatchingRule {
	idx, found := r.Contains(x)
	if !found {
		return nil
	}

	return r.Index(idx)
}

/*
Len is a thread-safe method that returns the effective length of the receiver slice collection.
*/
func (r MatchingRules) Len() int {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	return r.slice.len()
}

/*
String is a stringer method used to return the properly-delimited and formatted series of attributeType name or OID definitions.
*/
func (r MatchingRules) String() string { return `` }

/*
String is an unsafe convenience wrapper for Unmarshal(r). If an error is encountered, an empty string definition is returned. If reliability and error handling are important, use Unmarshal.
*/
func (r MatchingRule) String() (def string) {
	def, _ = Unmarshal(r)
	return
}

/*
IsZero returns a boolean value indicative of whether the receiver is considered empty or uninitialized.
*/
func (r MatchingRules) IsZero() bool {
	return r.slice.len() == 0
}

/*
IsZero returns a boolean value indicative of whether the receiver is considered empty or uninitialized.
*/
func (r *MatchingRule) IsZero() bool {
	return r == nil
}

/*
Set is a thread-safe append method that returns an error instance indicative of whether the append operation failed in some manner. Uniqueness is enforced for new elements based on Object Identifier and not the effective Name of the definition, if defined.
*/
func (r *MatchingRules) Set(x *MatchingRule) error {
	if _, exists := r.Contains(x.OID); exists {
		return fmt.Errorf("%T already contains %T:%s", r, x, x.OID)
	}

	r.mutex.Lock()
	defer r.mutex.Unlock()

	return r.slice.append(x)
}

/*
NewMatchingRules initializes and returns a new MatchingRulesCollection interface object.
*/
func NewMatchingRules() MatchingRuleCollection {
	var x interface{} = &MatchingRules{
		mutex: &sync.Mutex{},
		slice: make(collection, 0, 0),
	}
	return x.(MatchingRuleCollection)
}

/*
is returns a boolean value indicative of whether the provided interface argument is an enabled Boolean option.
*/
func (r MatchingRule) is(b interface{}) bool {
	switch tv := b.(type) {
	case Boolean:
		return r.bools.is(tv)
	case *LDAPSyntax:
		return r.OID.Equal(tv.OID)
	}

	return false
}

/*
Equal performs a deep-equal between the receiver and the provided collection type.

Description text is ignored.
*/
func (r *MatchingRule) Equal(x interface{}) bool {
	if assert, equals := x.(*MatchingRule); equals {
		return (assert.OID.String() == r.OID.String() &&
			assert.Name.Equal(r.Name))
	}

	return false
}

func (r *MatchingRule) validateSyntax() (err error) {
	if r.Syntax.IsZero() {
		err = raise(invalidSyntax,
			"%T.validateSyntax: zero syntax", r)
	}

	return
}

/*
Validate returns an error that reflects any fatal condition observed regarding the receiver configuration.
*/
func (r *MatchingRule) Validate() (err error) {
	return r.validate()
}

func (r *MatchingRule) validate() (err error) {
	if r.IsZero() {
		return raise(isZero, "%T.validate", r)
	}

	if err = r.validateSyntax(); err != nil {
		return
	}

	if err = r.validateSyntax(); err != nil {
		return
	}

	if err = validateBool(r.bools); err != nil {
		return
	}

	if err = validateNames(r.Name.strings()...); err != nil {
		return
	}

	if err = validateDesc(r.Description); err != nil {
		return
	}

	return
}

func (r *MatchingRule) unmarshal(namesok bool) (def string, err error) {
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

	// Syntax will never be zero
	def += WHSP + r.Syntax.Label()
	def += WHSP + r.Syntax.OID.String()

	if !r.Extensions.IsZero() {
		def += WHSP + r.Extensions.String()
	}

	def += WHSP + `)`

	return
}

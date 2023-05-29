package schemax

import "sync"

/*
MatchingRuleUseCollection describes all MatchingRuleUses-based types.
*/
type MatchingRuleUseCollection interface {
	// Get returns the *MatchingRuleUse instance retrieved as a result
	// of a term search, based on Name or OID. If no match is found,
	// nil is returned.
	Get(any) *MatchingRuleUse

	// Index returns the *MatchingRuleUse instance stored at the nth
	// index within the receiver, or nil.
	Index(int) *MatchingRuleUse

	// Equal performs a deep-equal between the receiver and the
	// interface MatchingRuleUseCollection provided.
	Equal(MatchingRuleUseCollection) bool

	// Set returns an error instance based on an attempt to add
	// the provided *MatchingRuleUse instance to the receiver.
	Set(*MatchingRuleUse) error

	// Refresh will update the collection of MatchingRuleUse
	// instances based on the contents of the provided instance
	// of AttributeTypeCollection.
	Refresh(AttributeTypeCollection) error

	// Contains returns the index number and presence boolean that
	// reflects the result of a term search within the receiver.
	Contains(any) (int, bool)

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
MatchingRuleUse conforms to the specifications of RFC4512 Section 4.1.4.
*/
type MatchingRuleUse struct {
	OID         OID
	Name        Name
	Description Description
	Obsolete    bool
	Applies     AttributeTypeCollection
	Extensions  *Extensions
	ufn         DefinitionUnmarshaler
	spec        string
	info        []byte
}

/*
MatchingRuleUses is a thread-safe collection of *MatchingRuleUse slice instances.
*/
type MatchingRuleUses struct {
	mutex *sync.Mutex
	slice collection
}

/*
Type returns the formal name of the receiver in order to satisfy signature requirements of the Definition interface type.
*/
func (r *MatchingRuleUse) Type() string {
	return `MatchingRuleUse`
}

/*
Equal performs a deep-equal between the receiver and the provided collection type.
*/
func (r MatchingRuleUses) Equal(x MatchingRuleUseCollection) bool {
	return r.slice.equal(x.(*MatchingRuleUses).slice)
}

/*
SetSpecifier is a convenience method that executes the SetSpecifier method in iterative fashion for all definitions within the receiver.
*/
func (r *MatchingRuleUses) SetSpecifier(spec string) {
	for i := 0; i < r.Len(); i++ {
		r.Index(i).SetSpecifier(spec)
	}
}

/*
SetUnmarshaler is a convenience method that executes the SetUnmarshaler method in iterative fashion for all definitions within the receiver.
*/
func (r *MatchingRuleUses) SetUnmarshaler(fn DefinitionUnmarshaler) {
	for i := 0; i < r.Len(); i++ {
		r.Index(i).SetUnmarshaler(fn)
	}
}

/*
Contains is a thread-safe method that returns a collection slice element index integer and a presence-indicative boolean value based on a term search conducted within the receiver.
*/
func (r MatchingRuleUses) Contains(x any) (int, bool) {
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
func (r MatchingRuleUses) Get(x any) *MatchingRuleUse {
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
	if &r == nil {
		return 0
	}

	r.mutex.Lock()
	defer r.mutex.Unlock()

	return r.slice.len()
}

/*
String is a non-functional stringer method needed to satisfy interface type requirements and should not be used. There is no practical application for a list of matchingRuleUse names or object identifiers in this package.
*/
func (r MatchingRuleUses) String() string { return `` }

/*
String is an unsafe convenience wrapper for Unmarshal(r). If an error is encountered, an empty string definition is returned. If reliability and error handling are important, use Unmarshal.
*/
func (r MatchingRuleUse) String() (def string) {
	def, _ = r.unmarshal()
	return
}

/*
SetSpecifier assigns a string value to the receiver, useful for placement into configurations that require a type name (e.g.: matchingruleuse). This will be displayed at the beginning of the definition value during the unmarshal or unsafe stringification process.
*/
func (r *MatchingRuleUse) SetSpecifier(spec string) {
	r.spec = spec
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
		return nil //silent
	}

	r.mutex.Lock()
	defer r.mutex.Unlock()

	return r.slice.append(x)
}

/*
Refresh accepts an AttributeTypeCollection which will be processed and used to create new, or update existing, *MatchingRuleUse instances within the receiver.
*/
func (r *MatchingRuleUses) Refresh(atc AttributeTypeCollection) (err error) {
	if atc.IsZero() {
		err = raise(noContent, "%T is nil", atc)
		return
	}

	createOrAppend := func(o OID, n Name, at *AttributeType) error {

		// If the definition exists already, save the index number.
		// If not found, handle it properly.
		idx, found := r.Contains(o)
		if !found {

			// Create an empty definition
			// to start with.
			r.Set(&MatchingRuleUse{
				OID:     o,
				Name:    n,
				Applies: NewApplicableAttributeTypes(),
			})

			// Make sure it actually got added
			idx, found = r.Contains(o)
			if !found {
				return raise(invalidMatchingRuleUse,
					"Attempt to register %s within %T failed for reasons unknown",
					o, r)
			}
		}

		// Append attributeType to receiver. Any definitions
		// already present as APPLIES values are silently
		// ignored.
		if e := r.Index(idx).Applies.Set(at); e != nil {
			return e
		}

		return nil
	}

	for i := 0; i < atc.Len(); i++ {
		a := atc.Index(i)
		if a.IsZero() {
			continue
		}

		if !a.Equality.IsZero() {
			if err = createOrAppend(a.Equality.OID, a.Equality.Name, a); err != nil {
				return
			}
		}

		if !a.Substring.IsZero() {
			if err = createOrAppend(a.Substring.OID, a.Substring.Name, a); err != nil {
				return
			}
		}

		if !a.Ordering.IsZero() {
			if err = createOrAppend(a.Ordering.OID, a.Ordering.Name, a); err != nil {
				return
			}
		}
	}

	return
}

/*
SetInfo assigns the byte slice to the receiver. This is a user-leveraged field intended to allow arbitrary information (documentation?) to be assigned to the definition.
*/
func (r *MatchingRuleUse) SetInfo(info []byte) {
	r.info = info
}

/*
Info returns the assigned informational byte slice instance stored within the receiver.
*/
func (r *MatchingRuleUse) Info() []byte {
	return r.info
}

/*
SetUnmarshaler assigns the provided DefinitionUnmarshaler signature value to the receiver. The provided function shall be executed during the unmarshal or unsafe stringification process.
*/
func (r *MatchingRuleUse) SetUnmarshaler(fn DefinitionUnmarshaler) {
	r.ufn = fn
}

/*
Equal performs a deep-equal between the receiver and the provided collection type.

Description text is ignored.
*/
func (r *MatchingRuleUse) Equal(x any) (eq bool) {
	z, ok := x.(MatchingRuleUse)
	if !ok {
		return
	}

	if z.IsZero() && r.IsZero() {
		eq = true
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

	if !z.Applies.Equal(r.Applies) {
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
NewMatchingRuleUse returns a newly initialized, yet effectively nil, instance of *MatchingRuleUse.

Users generally do not need to execute this function unless an instance of the returned type will be manually populated (as opposed to parsing a raw text definition).
*/
func NewMatchingRuleUse() *MatchingRuleUse {
	mru := new(MatchingRuleUse)
	mru.Applies = NewApplicableAttributeTypes()
	mru.Extensions = NewExtensions()
	return mru
}

/*
NewMatchingRuleUses initializes and returns a new MatchingRuleUsesCollection interface object.
*/
func NewMatchingRuleUses() MatchingRuleUseCollection {
	var x any = &MatchingRuleUses{
		mutex: &sync.Mutex{},
		slice: make(collection, 0, 0),
	}
	return x.(MatchingRuleUseCollection)
}

func (r *MatchingRuleUse) is(b any) bool {
	switch tv := b.(type) {
	case *AttributeType:
		if _, x := r.Applies.Contains(tv); !x {
			return false
		}
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

	return
}

func (r *MatchingRuleUse) unmarshal() (string, error) {
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
Map is a convenience method that returns a map[string][]string instance containing the effective contents of the receiver.
*/
func (r *MatchingRuleUse) Map() (def map[string][]string) {
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

	if len(r.Description) > 0 {
		def[`DESC`] = []string{r.Description.String()}
	}

	if !r.Applies.IsZero() {
		def[`APPLIES`] = make([]string, 0)
		for i := 0; i < r.Applies.Len(); i++ {
			appl := r.Applies.Index(i)
			term := appl.Name.Index(0)
			if len(term) == 0 {
				term = appl.OID.String()
			}
			def[`APPLIES`] = append(def[`APPLIES`], term)
		}
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

	return def
}

/*
MatchingRuleUseUnmarshaler is a package-included function that honors the signature of the first class (closure) DefinitionUnmarshaler type.

The purpose of this function, and similar user-devised ones, is to unmarshal a definition with specific formatting included, such as linebreaks, leading specifier declarations and indenting.
*/
func MatchingRuleUseUnmarshaler(x any) (def string, err error) {
	var r *MatchingRuleUse
	switch tv := x.(type) {
	case *MatchingRuleUse:
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

	// Applies will never be zero
	def += idnt + r.Applies.Label()
	def += WHSP + r.Applies.String()

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

func (r *MatchingRuleUse) unmarshalBasic() (def string, err error) {
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

	// Applies will never be zero
	def += WHSP + r.Applies.Label()
	def += WHSP + r.Applies.String()

	if !r.Extensions.IsZero() {
		def += WHSP + r.Extensions.String()
	}

	def += WHSP + tail

	return
}

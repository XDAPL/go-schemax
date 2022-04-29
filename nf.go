package schemax

import "sync"

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
	flags       definitionFlags
	ufn         DefinitionUnmarshalFunc
	spec        string
	info        []byte
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
Type returns the formal name of the receiver in order to satisfy signature requirements of the Definition interface type.
*/
func (r *NameForm) Type() string {
	return `NameForm`
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
SetInfo assigns the byte slice to the receiver. This is a user-leveraged field intended to allow arbitrary information (documentation?) to be assigned to the definition.
*/
func (r *NameForm) SetInfo(info []byte) {
	r.info = info
}

/*
Info returns the assigned informational byte slice instance stored within the receiver.
*/
func (r *NameForm) Info() []byte {
	return r.info
}

/*
SetUnmarshalFunc assigns the provided DefinitionUnmarshalFunc signature value to the receiver. The provided function shall be executed during the unmarshal or unsafe stringification process.
*/
func (r *NameForm) SetUnmarshalFunc(fn DefinitionUnmarshalFunc) {
	r.ufn = fn
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
	def, _ = r.unmarshal()
	return
}

/*
SetSpecifier assigns a string value to the receiver, useful for placement into configurations that require a type name (e.g.: nameform). This will be displayed at the beginning of the definition value during the unmarshal or unsafe stringification process.
*/
func (r *NameForm) SetSpecifier(spec string) {
	r.spec = spec
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
		return nil //silent
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

	if err = validateFlag(r.flags); err != nil {
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

func (r *NameForm) unmarshal() (string, error) {
	if err := r.validate(); err != nil {
		err = raise(invalidUnmarshal, err.Error())
		return ``, err
	}

	if r.ufn != nil {
		return r.ufn()
	}
	return r.unmarshalBasic()
}

/*
Map is a convenience method that returns a map[string][]string instance containing the effective contents of the receiver.
*/
func (r *NameForm) Map() (def map[string][]string) {
	if err := r.Validate(); err != nil {
		return
	}

	def = make(map[string][]string, 14)
	def[`RAW`] = []string{r.String()}
	def[`OID`] = []string{r.OID.String()}

	if !r.Name.IsZero() {
		def[`NAME`] = make([]string, 0)
		for i := 0; i < r.Name.Len(); i++ {
			def[`NAME`] = append(def[`NAME`], r.Name.Index(i))
		}
	}

	if len(r.Description) > 0 {
		def[`DESC`] = []string{r.Description.String()}
	}

	if !r.OC.IsZero() {
		def[`OC`] = []string{r.OC.OID.String(), r.OC.Name.Index(0)}
	}

	if !r.Must.IsZero() {
		def[`MUST`] = make([]string, 0)
		for i := 0; i < r.Must.Len(); i++ {
			must := r.Must.Index(i)
			term := must.Name.Index(0)
			if len(term) == 0 {
				term = must.OID.String()
			}
			def[`MUST`] = append(def[`MUST`], term)
		}
	}

	if !r.May.IsZero() {
		def[`MAY`] = make([]string, 0)
		for i := 0; i < r.May.Len(); i++ {
			must := r.May.Index(i)
			term := must.Name.Index(0)
			if len(term) == 0 {
				term = must.OID.String()
			}
			def[`MAY`] = append(def[`MAY`], term)
		}
	}

	if !r.Extensions.IsZero() {
		for k, v := range r.Extensions {
			def[k] = v
		}
	}

	if r.Obsolete() {
		def[`OBSOLETE`] = []string{`TRUE`}
	}

	return def
}

/*
UnmarshalFunc is a package-included function that honors the signature of the first class (closure) DefinitionUnmarshalFunc type.

The purpose of this function, and similar user-devised ones, is to unmarshal a definition with specific formatting included, such as linebreaks, leading specifier declarations and indenting.
*/
func (r *NameForm) UnmarshalFunc() (def string, err error) {
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

	if r.Obsolete() {
		def += idnt + Obsolete.String()
	}

	// OC will never be zero
	def += idnt + r.OC.Label()
	def += WHSP + r.OC.String()

	// Must will never be zero
	def += idnt + r.Must.Label()
	def += WHSP + r.Must.String()

	if !r.May.IsZero() {
		def += idnt + r.May.Label()
		def += WHSP + r.May.String()
	}

	if !r.Extensions.IsZero() {
		def += idnt + r.Extensions.String()
	}

	def += WHSP + tail

	return
}

func (r *NameForm) unmarshalBasic() (def string, err error) {
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

	if r.Obsolete() {
		def += WHSP + Obsolete.String()
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

	if !r.Extensions.IsZero() {
		def += WHSP + r.Extensions.String()
	}

	def += WHSP + tail

	return
}

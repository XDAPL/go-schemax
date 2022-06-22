package schemax

import "sync"

const nhr = `X-NOT-HUMAN-READABLE` // needed for certain LDAPSyntax definitions ...

/*
Extension represents any single definition extension such as "X-ORIGIN 'RFC4512'."
*/
type Extension struct {
	Label string
	Value []string
}

/*
len returns the number of members as an integer.
*/
func (r Extension) Len() int {
	return len(r.Value)
}

func (r *Extension) IsZero() bool {
	return r == nil
}

/*
String returns the qdstring form of the receiver. Single valued extensions return as "X-PARAM 'VALUE'", while multiple values are returned as "X-PARAM ( 'VALUE1' 'VALUE2' )". If no values, a zero string is returned.
*/
func (r Extension) String() (exts string) {
	l := r.Len()
	switch l {
	case 0:
		return
	case 1:
		exts += r.Label + ` '` + r.Value[0] + `'`
	default:
		vals := make([]string, l)
		for mi, mv := range r.Value {
			val := `'` + mv + `'`
			vals[mi] = val
		}
		exts += r.Label + ` ( ` + join(vals, ` `) + ` )`
	}

	return exts
}

/*
Equal performs a deep equal between the receiver and input instances, returning a boolean value indicative of equality.
*/
func (r *Extension) Equal(z *Extension) (eq bool) {
	if r.Label != z.Label {
		return
	}

	if r.Len() != z.Len() {
		return
	} else if r.Len()&z.Len() == 0 {
		eq = true
		return
	}

	for idx, val := range r.Value {
		if val != z.Value[idx] {
			return
		}
	}

	eq = true
	return
}

/*
Extensions is a map structure associating a string extension name (e.g.: X-ORIGIN) with a non-zero string value.
*/
type Extensions struct {
	mutex *sync.Mutex
	slice collection
}

/*
Get returns an instance of *Extension if its label matches the input value. Case is not important during the evaluation process.

A nil instance is returned if no match found.
*/
func (r Extensions) Get(label string) *Extension {
	for i := 0; i < r.Len(); i++ {
		ext := r.Index(i)
		if equalFold(label, ext.Label) {
			return ext
		}
	}

	return nil
}

/*
Index is a thread-safe method that returns the nth collection slice element if defined, else nil. This method supports use of negative indices which should be used with special care.
*/
func (r Extensions) Index(idx int) *Extension {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	assert, _ := r.slice.index(idx).(*Extension)
	return assert
}

/*
Contains is a thread-safe method that returns a collection slice element index integer and a presence-indicative boolean value based on a term search conducted within the receiver.
*/
func (r Extensions) Contains(x interface{}) (int, bool) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	return r.slice.contains(x)
}

/*
Exists returns a boolean value indicative of whether the named label exists within the receiver instance.

DEPRECATED: Use Extensions.Contains instead.
*/
func (r Extensions) Exists(label string) (exists bool) {
	for _, e := range r.slice {
		ext, ok := e.(*Extension)
		if !ok {
			continue
		}
		if ext.Label == label {
			exists = true
			break
		}
	}

	return
}

/*
labelIsValid returns a boolean value indicative of whether the field name is valid in that it begins with an `X-`.
*/
func (r Extension) labelIsValid() (valid bool) {
	if len(r.Label) < 2 {
		return
	}
	valid = r.Label[:2] == `X-`

	return
}

/*
IsZero returns a boolean value indicative of whether the receiver is considered empty or uninitialized.
*/
func (r *Extensions) IsZero() bool {
	if r != nil {
		return r.slice.isZero()
	}
	return r == nil
}

func (r *Extensions) Len() int {
	return r.slice.len()
}

/*
Set assigns the provided label and value(s) to the receiver instance. The interface{} input argument can be a pre-crafted *Extension instance, or two or more string slices. In the case of strings, the first argument must be a valid extension label (e.g.: X-ORIGIN). All subsequent values are considered values.

All subsequent values (strings or slices of strings) are interpreted as values to be assigned to said label.
*/
func (r *Extensions) Set(x ...interface{}) {
	var ext *Extension = new(Extension)
	switch tv := x[0].(type) {
	case string:
		if len(x) < 2 {
			return
		}
		vals := make([]string, 0)
		for i := 1; i < len(x); i++ {
			val, ok := x[i].(string)
			if !ok {
				continue
			}
			vals = append(vals, val)
		}
		ext.Label = tv
		ext.Value = vals
	case *Extension:
		if tv == nil {
			return
		}
		ext = tv
	default:
		return
	}

	if !ext.labelIsValid() || ext.Len() == 0 {
		return
	}

	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.slice.append(ext)

	return
}

func (r *Extensions) delete(x interface{}) {
	var idx int = -1
	var found bool
	switch tv := x.(type) {
	case string:
		idx, found = r.Contains(tv)
		if !found {
			return
		}
	case int:
		if r.Len() > tv && tv >= 0 {
			idx = tv
		}
	default:
		return
	}

	if idx == -1 {
		return
	}

	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.slice.delete(idx)
}

/*
Equal compares the receiver to the provided interface value (which must be of the same effective type).  A comparison is performed, and an equality-indicative boolean value is returned.
*/
func (r Extensions) Equal(x interface{}) (eq bool) {
	assert, ok := x.(*Extensions)
	if !ok {
		return
	}

	if assert.Len() != r.Len() {
		return
	}

	for k, v := range r.slice {
		e, ok := v.(*Extension)
		if !ok {
			return
		}

		if !e.Equal(assert.Index(k)) {
			return
		}

	}
	eq = true

	return
}

/*
String is a stringer method that returns the receiver data as a compliant schema definition component.
*/
func (r Extensions) String() (exts string) {
	return join(r.strings(), ` `)
}

/*
HumanReadable is a convenience method that searches the receiver for the 'X-NOT-HUMAN-READABLE' extension key and returns a boolean value indicative of whether the key's value is not TRUE. If the key is not found (or is explicitly set to FALSE) for some reason, a boolean value of true is returned as a fallback.
*/
func (r Extensions) HumanReadable() bool {
	if ext := r.Get(nhr); !ext.IsZero() {
		if ext.Len() > 0 {
			return ext.Value[0] != `TRUE`
		}
	}

	return true
}

/*
SetHumanReadable creates the X-NOT-HUMAN-READABLE extension with a value 'TRUE' if a boolean value of false is provided. If true, the extension is deleted if present.
*/
func (r *Extensions) SetHumanReadable(x bool) {
	r.delete(nhr)
	if !x {
		r.Set(nhr, `TRUE`)
	}
}

func (r Extensions) strings() (exts []string) {
	exts = make([]string, r.Len())
	for i := 0; i < r.Len(); i++ {
		if ext := r.Index(i); !ext.IsZero() {
			exts[i] = ext.String()
		}
	}

	return
}

/*
NewExtensions returns a new instance of Extensions, intended for assignment to any definition type.
*/
func NewExtensions() *Extensions {
	return &Extensions{
		mutex: &sync.Mutex{},
		slice: make(collection, 0),
	}
}

package schemax

/*
NewName initializes and returns a new instance of [Name].
*/
func NewName() (name Name) {
	name = Name(newQDescrList(``))
	name.cast().SetPresentationPolicy(name.smvStringer)
	return
}

/*
Len returns the current integer length of the receiver instance.
*/
func (r Name) Len() int {
	return r.len()
}

func (r Name) len() int {
	return r.cast().Len()
}

/*
Push returns an error following an attempt to push name into
the receiver instance.  The name value must be non-zero in
length and must be a valid RFC 4512 "descr" qualifier.
*/
func (r Name) Push(name string) error {
	return r.push(name)
}

func (r Name) push(name string) (err error) {
	if len(name) == 0 {
		err = errorf("%T descriptor is nil; cannot append to %T", name, r)
		return
	} else if !isDescriptor(name) {
		err = errorf("'%s' represents an invalid descriptor", name)
		return
	}

	r.cast().Push(name)

	return
}

// Contains returns a Boolean value indicative of whether the input
// name value was found within the receiver instance.  Case-folding
// is not significant in this operation.
func (r Name) Contains(name string) bool {
	return r.contains(name)
}

func (r Name) contains(name string) (found bool) {
	for i := 0; i < r.len(); i++ {
		if eq(name, r.index(i)) {
			found = true
			break
		}
	}

	return
}

/*
List returns slices of string names from within the receiver.
*/
func (r Name) List() (list []string) {
	for i := 0; i < r.len(); i++ {
		list = append(list, r.index(i))
	}

	return
}

/*
IsZero returns a Boolean value indicative of nilness of the
receiver instance.
*/
func (r Name) IsZero() bool {
	return r.cast().IsZero()
}

/*
String is a stringer method that returns a properly formatted
quoted descriptor list based on the number of string values
within the receiver instance.
*/
func (r Name) String() string {
	return r.cast().String()
}

// smvStringer (single-or-multivalue) is the underlying
// stringer mechanism called by r.String.
//
// This method qualifies for stackage.PresentationPolicy.
func (r Name) smvStringer(_ ...any) (smv string) {
	switch tv := r.len(); tv {
	case 0:
		break
	case 1:
		smv = `'` + r.index(0) + `'`
	default:
		var smvs []string
		for i := 0; i < tv; i++ {
			smvs = append(smvs, `'`+r.index(i)+`'`)
		}

		padchar := string(rune(32))
		if !r.cast().IsPadded() {
			padchar = ``
		}
		smv = `(` + padchar + join(smvs, ` `) + padchar + `)`
	}

	return
}

/*
Index returns the instance of string found within the receiver stack
instance at index N.  If no instance is found at the index specified,
a zero string instance is returned.
*/
func (r Name) Index(idx int) string {
	return r.index(idx)
}

func (r Name) index(idx int) (name string) {
	raw, _ := r.cast().Index(idx)
	name, _ = raw.(string)
	return
}

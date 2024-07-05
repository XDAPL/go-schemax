package schemax

/*
NewName initializes and returns a new instance of [QuotedDescriptorList].
*/
func NewName() (name QuotedDescriptorList) {
	name = QuotedDescriptorList(newQDescrList(``))
	name.cast().SetPresentationPolicy(name.smvStringer)
	return
}

/*
Len returns the current integer length of the receiver instance.
*/
func (r QuotedDescriptorList) Len() int {
	return r.len()
}

func (r QuotedDescriptorList) len() int {
	return r.cast().Len()
}

/*
Push returns an error following an attempt to push name into
the receiver instance.  The name value must be non-zero in
length and must be a valid RFC 4512 "descr" qualifier.
*/
func (r QuotedDescriptorList) Push(name string) error {
	return r.push(name)
}

func (r QuotedDescriptorList) push(name string) (err error) {
	if len(name) == 0 {
		err = ErrNilInput
		return
	} else if !isDescriptor(name) {
		err = mkerr("Invalid RFC 4512 descriptor '" + name + `'`)
		return
	}

	r.cast().Push(name)

	return
}

/*
Contains returns a Boolean value indicative of whether the input
name value was found within the receiver instance.  Case-folding
is not significant in this operation.
*/
func (r QuotedDescriptorList) Contains(name string) bool {
	return r.contains(name)
}

func (r QuotedDescriptorList) contains(name string) (found bool) {
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
func (r QuotedDescriptorList) List() (list []string) {
	for i := 0; i < r.len(); i++ {
		list = append(list, r.index(i))
	}

	return
}

/*
IsZero returns a Boolean value indicative of a nil receiver state.
*/
func (r QuotedDescriptorList) IsZero() bool {
	return r.cast().IsZero()
}

/*
String is a stringer method that returns a properly formatted
quoted descriptor list based on the number of string values
within the receiver instance.
*/
func (r QuotedDescriptorList) String() string {
	return r.cast().String()
}

// smvStringer (single-or-multivalue) is the underlying
// stringer mechanism called by r.String.
//
// This method qualifies for stackage.PresentationPolicy.
func (r QuotedDescriptorList) smvStringer(_ ...any) (smv string) {
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
func (r QuotedDescriptorList) Index(idx int) string {
	return r.index(idx)
}

func (r QuotedDescriptorList) index(idx int) (name string) {
	raw, _ := r.cast().Index(idx)
	name, _ = raw.(string)
	return
}

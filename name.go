package schemax

import (
	antlr4512 "github.com/JesseCoretta/go-rfc4512-antlr"
)

/*
NewName initializes and returns a new instance of [DefinitionName].
*/
func NewName() (name DefinitionName) {
	name = DefinitionName(newQDescrList(``))
	name.cast().SetPresentationPolicy(name.smvStringer)
	return
}

/*
Len returns the current integer length of the receiver instance.
*/
func (r DefinitionName) Len() int {
	return r.len()
}

func (r DefinitionName) len() int {
	return r.cast().Len()
}

/*
Push returns an error following an attempt to push name into
the receiver instance.  The name value must be non-zero in
length and must be a valid RFC 4512 "descr" qualifier.
*/
func (r DefinitionName) Push(name string) error {
	return r.push(name)
}

func (r DefinitionName) push(name string) (err error) {
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
func (r DefinitionName) Contains(name string) bool {
	return r.contains(name)
}

func (r DefinitionName) contains(name string) (found bool) {
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
func (r DefinitionName) List() (list []string) {
	for i := 0; i < r.len(); i++ {
		list = append(list, r.index(i))
	}

	return
}

/*
IsZero returns a Boolean value indicative of nilness of the
receiver instance.
*/
func (r DefinitionName) IsZero() bool {
	return r.cast().IsZero()
}

/*
String is a stringer method that returns a properly formatted
quoted descriptor list based on the number of string values
within the receiver instance.
*/
func (r DefinitionName) String() string {
	return r.cast().String()
}

// smvStringer (single-or-multivalue) is the underlying
// stringer mechanism called by r.String.
//
// This method qualifies for stackage.PresentationPolicy.
func (r DefinitionName) smvStringer(_ ...any) (smv string) {
	switch tv := r.len(); tv {
	case 0:
		break
	case 1:
		smv = sprintf("'%s'", r.index(0))
	default:
		var smvs []string
		for i := 0; i < tv; i++ {
			val := sprintf("'%s'", r.index(i))
			smvs = append(smvs, val)
		}

		padchar := string(rune(32))
		if !r.cast().IsPadded() {
			padchar = ``
		}
		smv = sprintf("(%s%s%s)",
			padchar,
			join(smvs, ` `),
			padchar)
	}

	return
}

/*
Index returns the instance of string found within the receiver stack
instance at index N.  If no instance is found at the index specified,
a zero string instance is returned.
*/
func (r DefinitionName) Index(idx int) string {
	return r.index(idx)
}

func (r DefinitionName) index(idx int) (name string) {
	raw, _ := r.cast().Index(idx)
	name, _ = raw.(string)
	return
}

// nameContext extracts zero (0) or more names from the ctx instance,
// and returns them via names ([]string).  Note that presence of a name
// value applies to all definitions EXCEPT LDAPSyntax.
func nameContext(ctx *antlr4512.DefinitionNameContext) (names DefinitionName, err error) {
	if ctx == nil {
		err = errorf("%T instance is nil", ctx)
		return
	}

	names = NewName()
	_names := ctx.AllQuotedDescriptor()
	for i := 0; i < len(_names); i++ {
		if sym := _names[i].GetText(); len(sym) > 0 {
			names.push(trim(sym, `''`))
		}
	}

	// Determine if parentheticals were needed, and
	// if so, that both L & R were specified.  Note
	// that parentheticals are required if more than
	// one name is present, though a single name MAY
	// be similarly encapsulated, but this is optional.
	switch l := names.Len(); l {
	case 0:
		err = errorf("No names parsed from %T", ctx)
	default:
		var ct int
		if ctx.OpenParen() != nil {
			ct++
		}
		if ctx.CloseParen() != nil {
			ct++
		}

		if ct%2 != 0 {
			// odd number of parentheticals is bad in any scenario.
			err = errorf("Missing open or close parenthesis in %T", ctx)
		} else if ct != 2 && l > 1 {
			// A multi-valued NAME MUST be parenthetical.
			err = errorf("Multi-valued NAME must be parenthetical")
		}
		// A single-valued NAME MAY be parenthetical.
	}

	return
}

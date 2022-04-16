package schemax

/*
Boolean is an unsigned 8-bit integer that describes zero or more perceived values that only manifest visually when TRUE. Such verisimilitude is revealed by the presence of the indicated Boolean value's "label" name, such as `SINGLE-VALUE`.  The actual value "TRUE" is never actually seen in textual format.
*/
type Boolean uint8

const (
	Obsolete           Boolean = 1 << iota // 1
	SingleValue                            // 2
	Collective                             // 4
	NoUserModification                     // 8
	HumanReadable                          // 16 (applies only to *LDAPSyntax instances)
)

func (r Boolean) IsZero() bool {
	return uint8(r) == 0
}

/*
is returns a boolean value indicative of whether the specified value is enabled within the receiver instance.
*/
func (r Boolean) is(o Boolean) bool {
	return r.enabled(o)
}

/*
String is a stringer method that returns the name(s) Boolean receiver in question, whether it represents multiple boolean flags or only one.

If only a specific Boolean string is desired (if enabled), use Boolean.<Name>() (e.g: Boolean.Obsolete()).
*/
func (r Boolean) String() (val string) {

	// Look for so-called "pure"
	// boolean values first ...
	switch r {
	case NoUserModification:
		return `NO-USER-MODIFICATION`
	case SingleValue:
		return `SINGLE-VALUE`
	case Obsolete:
		return `OBSOLETE`
	case Collective:
		return `COLLECTIVE`
	case HumanReadable:
		return `HUMAN-READABLE`
	}

	// Assume multiple boolean bits
	// are set concurrently ...
	strs := []Boolean{
		Obsolete,
		Collective,
		SingleValue,
		HumanReadable,
		NoUserModification,
	}

	vals := make([]string, 0)
	for _, v := range strs {
		if r.enabled(v) {
			vals = append(vals, v.String())
		}
	}

	if len(vals) != 0 {
		val = join(vals, ` `)
	}

	return
}

/*
Obsolete returns the `OBSOLETE` flag if the appropriate bits are set within the receiver instance.
*/
func (r Boolean) Obsolete() string {
	if r.enabled(Obsolete) {
		return r.String()
	}

	return ``
}

/*
SingleValue returns the `SINGLE-VALUE` flag if the appropriate bits are set within the receiver instance.
*/
func (r Boolean) SingleValue() string {
	if r.enabled(SingleValue) {
		return r.String()
	}

	return ``
}

/*
Collective returns the `COLLECTIVE` flag if the appropriate bits are set within the receiver instance.
*/
func (r Boolean) Collective() string {
	if r.enabled(Collective) {
		return r.String()
	}

	return ``
}

/*
HumanReadable returns the `HUMAN-READABLE` flag if the appropriate bits are set within the receiver instance.  Note this would only apply to *LDAPSyntax instances bearing this type.
*/
func (r Boolean) HumanReadable() string {
	if r.enabled(HumanReadable) {
		return r.String()
	}

	return ``
}

/*
NoUserModification returns the `NO-USER-MODIFICATION` flag if the appropriate bits are set within the receiver instance.
*/
func (r Boolean) NoUserModification() string {
	if r.enabled(NoUserModification) {
		return r.String()
	}

	return ``
}

/*
Unset removes the specified Boolean bits from the receiver instance of Boolean, thereby "disabling" the provided option.
*/
func (r *Boolean) Unset(o Boolean) {
	r.unset(o)
}

/*
Set adds the specified Boolean bits to the receiver instance of Boolean, thereby "enabling" the provided option.
*/
func (r *Boolean) Set(o Boolean) {
	r.set(o)
}

func (r *Boolean) set(o Boolean) {
	*r = *r ^ o
}

func (b *Boolean) unset(o Boolean) {
	*b = *b &^ o
}

func (r Boolean) enabled(o Boolean) bool {
	return r&o != 0
}

/*
setBoolean is a private method used by reflect to set boolean values.
*/
func (r *DITStructureRule) setBoolean(b Boolean) {
	r.bools.set(b)
}

/*
Obsolete returns a boolean value that reflects whether the receiver instance is marked OBSOLETE.
*/
func (r DITStructureRule) Obsolete() bool {
	return r.bools.is(Obsolete)
}

/*
setBoolean is a private method used by reflect to set boolean values.
*/
func (r *DITContentRule) setBoolean(b Boolean) {
	r.bools.set(b)
}

/*
Obsolete returns a boolean value that reflects whether the receiver instance is marked OBSOLETE.
*/
func (r DITContentRule) Obsolete() bool {
	return r.bools.is(Obsolete)
}

/*
setBoolean is a private method used by reflect to set boolean values.
*/
func (r *MatchingRuleUse) setBoolean(b Boolean) {
	r.bools.set(b)
}

/*
Obsolete returns a boolean value that reflects whether the receiver instance is marked OBSOLETE.
*/
func (r MatchingRuleUse) Obsolete() bool {
	return r.bools.is(Obsolete)
}

/*
setBoolean is a private method used by reflect to set boolean values.
*/
func (r *MatchingRule) setBoolean(b Boolean) {
	r.bools.set(b)
}

/*
Obsolete returns a boolean value that reflects whether the receiver instance is marked OBSOLETE.
*/
func (r MatchingRule) Obsolete() bool {
	return r.bools.is(Obsolete)
}

/*
setBoolean is a private method used by reflect to set boolean values.
*/
func (r *LDAPSyntax) setBoolean(b Boolean) {
	r.bools.set(b)
}

/*
Obsolete returns a boolean value that reflects whether the receiver instance is marked OBSOLETE.
*/
func (r LDAPSyntax) Obsolete() bool {
	return r.bools.is(Obsolete)
}

/*
setBoolean is a private method used by reflect to set boolean values.
*/
func (r *ObjectClass) setBoolean(b Boolean) {
	r.bools.set(b)
}

/*
Obsolete returns a boolean value that reflects whether the receiver instance is marked OBSOLETE.
*/
func (r ObjectClass) Obsolete() bool {
	return r.bools.is(Obsolete)
}

/*
setBoolean is a private method used by reflect to set boolean values.
*/
func (r *AttributeType) setBoolean(b Boolean) {
	r.bools.set(b)
}

/*
Obsolete returns a boolean value that reflects whether the receiver instance is marked OBSOLETE.
*/
func (r AttributeType) Obsolete() bool {
	return r.bools.is(Obsolete)
}

/*
setBoolean is a private method used by reflect to set boolean values.
*/
func (r *NameForm) setBoolean(b Boolean) {
	r.bools.set(b)
}

/*
Obsolete returns a boolean value that reflects whether the receiver instance is marked OBSOLETE.
*/
func (r NameForm) Obsolete() bool {
	return r.bools.is(Obsolete)
}

/*
SetObsolete marks the receiver as OBSOLETE.
*/
func (r *AttributeType) SetObsolete() {
	r.bools.set(Obsolete)
}

/*
SetObsolete marks the receiver as OBSOLETE.
*/
func (r *ObjectClass) SetObsolete() {
	r.bools.set(Obsolete)
}

/*
SetObsolete marks the receiver as OBSOLETE.
*/
func (r *MatchingRule) SetObsolete() {
	r.bools.set(Obsolete)
}

/*
SetObsolete marks the receiver as OBSOLETE.
*/
func (r *MatchingRuleUse) SetObsolete() {
	r.bools.set(Obsolete)
}

/*
SetObsolete marks the receiver as OBSOLETE.
*/
func (r *DITStructureRule) SetObsolete() {
	r.bools.set(Obsolete)
}

/*
SetObsolete marks the receiver as OBSOLETE.
*/
func (r *DITContentRule) SetObsolete() {
	r.bools.set(Obsolete)
}

/*
SetObsolete marks the receiver as OBSOLETE.
*/
func (r *NameForm) SetObsolete() {
	r.bools.set(Obsolete)
}

/*
SetObsolete marks the receiver as OBSOLETE.
*/
func (r *LDAPSyntax) SetObsolete() {
	r.bools.set(Obsolete)
}

/*
SetHumanReadable marks the LDAPSyntax receiver as HumanReadable.
*/
func (r *LDAPSyntax) SetHumanReadable() {
	r.bools.set(HumanReadable)
}

/*
HumanReadable returns a boolean value indicative of whether the receiver describes a HUMAN-READABLE attribute type, typically manifesting as the X-NOT-HUMAN-READABLE ldapSyntax extension.
*/
func (r AttributeType) HumanReadable() bool {
	return r.bools.is(HumanReadable)
}

/*
SetCollective marks the receiver as COLLECTIVE.
*/
func (r *AttributeType) SetCollective() {
	r.bools.set(Collective)
}

/*
Collective returns a boolean value indicative of whether the receiver describes a COLLECTIVE attribute type.
*/
func (r AttributeType) Collective() bool {
	return r.bools.is(Collective)
}

/*
SetCollective marks the receiver as NO-USER-MODIFICATION.
*/
func (r *AttributeType) SetNoUserModification() {
	r.bools.set(NoUserModification)
}

/*
NoUserModification returns a boolean value indicative of whether the receiver describes a NO-USER-MODIFICATION attribute type.
*/
func (r AttributeType) NoUserModification() bool {
	return r.bools.is(NoUserModification)
}

/*
SetSingleValue marks the receiver as SINGLE-VALUE.
*/
func (r *AttributeType) SetSingleValue() {
	r.bools.set(SingleValue)
}

/*
SingleValue returns a boolean value indicative of whether the receiver describes a SINGLE-VALUE attribute type.
*/
func (r AttributeType) SingleValue() bool {
	return r.bools.is(SingleValue)
}

func validateBool(b Boolean) (err error) {
	if b.is(Collective) && b.is(SingleValue) {
		return raise(invalidBoolean,
			"Cannot have single-valued collective attribute")
	}

	return
}

package schemax

/*
definitionFlags is an unsigned 8-bit integer that describes zero or more perceived values that only manifest visually when TRUE. Such verisimilitude is revealed by the presence of the indicated definitionFlags value's "label" name, such as `SINGLE-VALUE`.  The actual value "TRUE" is never actually seen in textual format.
*/
type definitionFlags uint8

const (
	Obsolete           definitionFlags = 1 << iota // 1
	SingleValue                                    // 2
	Collective                                     // 4
	NoUserModification                             // 8
	HumanReadable                                  // 16 (applies only to *LDAPSyntax instances)
)

func (r definitionFlags) IsZero() bool {
	return uint8(r) == 0
}

/*
is returns a boolean value indicative of whether the specified value is enabled within the receiver instance.
*/
func (r definitionFlags) is(o definitionFlags) bool {
	return r.enabled(o)
}

/*
String is a stringer method that returns the name(s) definitionFlags receiver in question, whether it represents multiple boolean flags or only one.

If only a specific definitionFlags string is desired (if enabled), use definitionFlags.<Name>() (e.g: definitionFlags.Obsolete()).
*/
func (r definitionFlags) String() (val string) {

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
	strs := []definitionFlags{
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
func (r definitionFlags) Obsolete() string {
	if r.enabled(Obsolete) {
		return r.String()
	}

	return ``
}

/*
SingleValue returns the `SINGLE-VALUE` flag if the appropriate bits are set within the receiver instance.
*/
func (r definitionFlags) SingleValue() string {
	if r.enabled(SingleValue) {
		return r.String()
	}

	return ``
}

/*
Collective returns the `COLLECTIVE` flag if the appropriate bits are set within the receiver instance.
*/
func (r definitionFlags) Collective() string {
	if r.enabled(Collective) {
		return r.String()
	}

	return ``
}

/*
HumanReadable returns the `HUMAN-READABLE` flag if the appropriate bits are set within the receiver instance.  Note this would only apply to *LDAPSyntax instances bearing this type.
*/
func (r definitionFlags) HumanReadable() string {
	if r.enabled(HumanReadable) {
		return r.String()
	}

	return ``
}

/*
NoUserModification returns the `NO-USER-MODIFICATION` flag if the appropriate bits are set within the receiver instance.
*/
func (r definitionFlags) NoUserModification() string {
	if r.enabled(NoUserModification) {
		return r.String()
	}

	return ``
}

/*
Unset removes the specified definitionFlags bits from the receiver instance of definitionFlags, thereby "disabling" the provided option.
*/
func (r *definitionFlags) Unset(o definitionFlags) {
	r.unset(o)
}

/*
Set adds the specified definitionFlags bits to the receiver instance of definitionFlags, thereby "enabling" the provided option.
*/
func (r *definitionFlags) Set(o definitionFlags) {
	r.set(o)
}

func (r *definitionFlags) set(o definitionFlags) {
	*r = *r ^ o
}

func (b *definitionFlags) unset(o definitionFlags) {
	*b = *b &^ o
}

func (r definitionFlags) enabled(o definitionFlags) bool {
	return r&o != 0
}

/*
setdefinitionFlags is a private method used by reflect to set boolean values.
*/
func (r *DITStructureRule) setdefinitionFlags(b definitionFlags) {
	r.flags.set(b)
}

/*
Obsolete returns a boolean value that reflects whether the receiver instance is marked OBSOLETE.
*/
func (r DITStructureRule) Obsolete() bool {
	return r.flags.is(Obsolete)
}

/*
setdefinitionFlags is a private method used by reflect to set boolean values.
*/
func (r *DITContentRule) setdefinitionFlags(b definitionFlags) {
	r.flags.set(b)
}

/*
Obsolete returns a boolean value that reflects whether the receiver instance is marked OBSOLETE.
*/
func (r DITContentRule) Obsolete() bool {
	return r.flags.is(Obsolete)
}

/*
setdefinitionFlags is a private method used by reflect to set boolean values.
*/
func (r *MatchingRuleUse) setdefinitionFlags(b definitionFlags) {
	r.flags.set(b)
}

/*
Obsolete returns a boolean value that reflects whether the receiver instance is marked OBSOLETE.
*/
func (r MatchingRuleUse) Obsolete() bool {
	return r.flags.is(Obsolete)
}

/*
setdefinitionFlags is a private method used by reflect to set boolean values.
*/
func (r *MatchingRule) setdefinitionFlags(b definitionFlags) {
	r.flags.set(b)
}

/*
Obsolete returns a boolean value that reflects whether the receiver instance is marked OBSOLETE.
*/
func (r MatchingRule) Obsolete() bool {
	return r.flags.is(Obsolete)
}

/*
setdefinitionFlags is a private method used by reflect to set boolean values.
*/
func (r *LDAPSyntax) setdefinitionFlags(b definitionFlags) {
	r.flags.set(b)
}

/*
Obsolete returns a boolean value that reflects whether the receiver instance is marked OBSOLETE.
*/
func (r LDAPSyntax) Obsolete() bool {
	return r.flags.is(Obsolete)
}

/*
setdefinitionFlags is a private method used by reflect to set boolean values.
*/
func (r *ObjectClass) setdefinitionFlags(b definitionFlags) {
	r.flags.set(b)
}

/*
Obsolete returns a boolean value that reflects whether the receiver instance is marked OBSOLETE.
*/
func (r ObjectClass) Obsolete() bool {
	return r.flags.is(Obsolete)
}

/*
setdefinitionFlags is a private method used by reflect to set boolean values.
*/
func (r *AttributeType) setdefinitionFlags(b definitionFlags) {
	r.flags.set(b)
}

/*
Obsolete returns a boolean value that reflects whether the receiver instance is marked OBSOLETE.
*/
func (r AttributeType) Obsolete() bool {
	return r.flags.is(Obsolete)
}

/*
setdefinitionFlags is a private method used by reflect to set boolean values.
*/
func (r *NameForm) setdefinitionFlags(b definitionFlags) {
	r.flags.set(b)
}

/*
Obsolete returns a boolean value that reflects whether the receiver instance is marked OBSOLETE.
*/
func (r NameForm) Obsolete() bool {
	return r.flags.is(Obsolete)
}

/*
SetObsolete marks the receiver as OBSOLETE.
*/
func (r *AttributeType) SetObsolete() {
	r.flags.set(Obsolete)
}

/*
SetObsolete marks the receiver as OBSOLETE.
*/
func (r *ObjectClass) SetObsolete() {
	r.flags.set(Obsolete)
}

/*
SetObsolete marks the receiver as OBSOLETE.
*/
func (r *MatchingRule) SetObsolete() {
	r.flags.set(Obsolete)
}

/*
SetObsolete marks the receiver as OBSOLETE.
*/
func (r *MatchingRuleUse) SetObsolete() {
	r.flags.set(Obsolete)
}

/*
SetObsolete marks the receiver as OBSOLETE.
*/
func (r *DITStructureRule) SetObsolete() {
	r.flags.set(Obsolete)
}

/*
SetObsolete marks the receiver as OBSOLETE.
*/
func (r *DITContentRule) SetObsolete() {
	r.flags.set(Obsolete)
}

/*
SetObsolete marks the receiver as OBSOLETE.
*/
func (r *NameForm) SetObsolete() {
	r.flags.set(Obsolete)
}

/*
SetObsolete marks the receiver as OBSOLETE.
*/
func (r *LDAPSyntax) SetObsolete() {
	r.flags.set(Obsolete)
}

/*
SetHumanReadable marks the LDAPSyntax receiver as HumanReadable.
*/
func (r *LDAPSyntax) SetHumanReadable() {
	r.flags.set(HumanReadable)
}

/*
HumanReadable returns a boolean value indicative of whether the receiver describes a HUMAN-READABLE attribute type, typically manifesting as the X-NOT-HUMAN-READABLE ldapSyntax extension.
*/
func (r *AttributeType) HumanReadable() bool {
	return r.flags.is(HumanReadable)
}

/*
SetCollective marks the receiver as COLLECTIVE.
*/
func (r *AttributeType) SetCollective() {
	r.flags.set(Collective)
}

/*
Collective returns a boolean value indicative of whether the receiver describes a COLLECTIVE attribute type.
*/
func (r *AttributeType) Collective() bool {
	return r.flags.is(Collective)
}

/*
SetCollective marks the receiver as NO-USER-MODIFICATION.
*/
func (r *AttributeType) SetNoUserModification() {
	r.flags.set(NoUserModification)
}

/*
NoUserModification returns a boolean value indicative of whether the receiver describes a NO-USER-MODIFICATION attribute type.
*/
func (r *AttributeType) NoUserModification() bool {
	return r.flags.is(NoUserModification)
}

/*
SetSingleValue marks the receiver as SINGLE-VALUE.
*/
func (r *AttributeType) SetSingleValue() {
	r.flags.set(SingleValue)
}

/*
SingleValue returns a boolean value indicative of whether the receiver describes a SINGLE-VALUE attribute type.
*/
func (r *AttributeType) SingleValue() bool {
	return r.flags.is(SingleValue)
}

func validateFlag(b definitionFlags) (err error) {
	if b.is(Collective) && b.is(SingleValue) {
		return raise(invalidFlag,
			"Cannot have single-valued collective attribute")
	}

	return
}

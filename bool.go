package schemax

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
func (r Boolean) string() (val string) {

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
			vals = append(vals, v.string())
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
		return r.string()
	}

	return ``
}

/*
SingleValue returns the `SINGLE-VALUE` flag if the appropriate bits are set within the receiver instance.
*/
func (r Boolean) SingleValue() string {
	if r.enabled(SingleValue) {
		return r.string()
	}

	return ``
}

/*
Collective returns the `COLLECTIVE` flag if the appropriate bits are set within the receiver instance.
*/
func (r Boolean) Collective() string {
	if r.enabled(Collective) {
		return r.string()
	}

	return ``
}

/*
HumanReadable returns the `HUMAN-READABLE` flag if the appropriate bits are set within the receiver instance.  Note this would only apply to *LDAPSyntax instances bearing this type.
*/
func (r Boolean) HumanReadable() string {
	if r.enabled(HumanReadable) {
		return r.string()
	}

	return ``
}

/*
NoUserModification returns the `NO-USER-MODIFICATION` flag if the appropriate bits are set within the receiver instance.
*/
func (r Boolean) NoUserModification() string {
	if r.enabled(NoUserModification) {
		return r.string()
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
setBoolean is a private method used by reflect to set boolean values.
*/
func (r *DITContentRule) setBoolean(b Boolean) {
	r.bools.set(b)
}

/*
setBoolean is a private method used by reflect to set boolean values.
*/
func (r *MatchingRuleUse) setBoolean(b Boolean) {
	r.bools.set(b)
}

/*
setBoolean is a private method used by reflect to set boolean values.
*/
func (r *MatchingRule) setBoolean(b Boolean) {
	r.bools.set(b)
}

/*
setBoolean is a private method used by reflect to set boolean values.
*/
func (r *LDAPSyntax) setBoolean(b Boolean) {
	r.bools.set(b)
}

/*
setBoolean is a private method used by reflect to set boolean values.
*/
func (r *ObjectClass) setBoolean(b Boolean) {
	r.bools.set(b)
}

/*
setBoolean is a private method used by reflect to set boolean values.
*/
func (r *AttributeType) setBoolean(b Boolean) {
	r.bools.set(b)
}

/*
setBoolean is a private method used by reflect to set boolean values.
*/
func (r *NameForm) setBoolean(b Boolean) {
	r.bools.set(b)
}

func (r *AttributeType) SetObsolete() {
	r.bools.set(Obsolete)
}

func (r *ObjectClass) SetObsolete() {
	r.bools.set(Obsolete)
}

func (r *MatchingRule) SetObsolete() {
	r.bools.set(Obsolete)
}

func (r *MatchingRuleUse) SetObsolete() {
	r.bools.set(Obsolete)
}

func (r *DITStructureRule) SetObsolete() {
	r.bools.set(Obsolete)
}

func (r *DITContentRule) SetObsolete() {
	r.bools.set(Obsolete)
}

func (r *NameForm) SetObsolete() {
	r.bools.set(Obsolete)
}

func (r *LDAPSyntax) SetHumanReadable() {
	r.bools.set(HumanReadable)
}

func (r *AttributeType) SetCollective() {
	r.bools.set(Collective)
}

func (r *AttributeType) SetNoUserModification() {
	r.bools.set(NoUserModification)
}

func (r *AttributeType) SetSingleValue() {
	r.bools.set(SingleValue)
}

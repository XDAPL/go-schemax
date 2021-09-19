package schemax

/*
String is a stringer method that returns the string-form of the receiver instance.
*/
func (r OID) string() string {
	return string(r)
}

/*
String is a stringer method that returns the string-form of the receiver instance.
*/
func (r Description) string() string {
	return `'` + string(r) + `'`
}

/*
String is a stringer method that returns the string-form of the receiver instance.
*/
func (r Usage) String() string {
	switch r {
	case DirectoryOperation:
		return `directoryOperation`
	case DistributedOperation:
		return `distributedOperation`
	case DSAOperation:
		return `dSAOperation`
	}

	return `` // default is userApplication, but it need not be stated literally
}

/*
String is a stringer method that returns the string-form of the receiver instance.
*/
func (r Kind) String() string {
	switch r {
	case Abstract:
		return `ABSTRACT`
	case Structural:
		return `STRUCTURAL`
	case Auxiliary:
		return `AUXILIARY`
	}

	return `` // no default
}

/*
String is a stringer method that returns the receiver data as a compliant schema definition component.
*/
func (r Extensions) String() (exts string) {
	for _, v := range r.strings() {
		exts += v + ` `
	}
	if len(exts) == 0 {
		return
	}
	if exts[len(exts)-1] == ' ' {
		exts = exts[:len(exts)-1]
	}

	return
}

func (r Extensions) strings() (exts []string) {
	exts = make([]string, len(r))
	ct := 0
	for k, v := range r {
		if len(v) == 1 {
			exts[ct] = k + ` '` + v[0] + `'`
		} else if len(v) > 1 {
			vals := make([]string, len(v))
			for mi, mv := range v {
				val := `'` + mv + `'`
				vals[mi] = val
			}
			exts[ct] = k + ` ( ` + join(vals, ` `) + ` )`
		}
		ct++
	}

	return
}

/*
String is a qdescrs-compliant stringer method suitable for presentation.
*/
func (r Name) String() (str string) {
	switch len(r) {
	case 0:
		return
	case 1:
		return `'` + r[0].(string) + `'`
	}

	str += `( `
	for i := 0; i < len(r); i++ {
		str += `'` + r[i].(string) + `' `
	}
	str += `)`

	return
}

func (r Name) strings() (strings []string) {
	strings = make([]string, len(r), len(r))
	for i, v := range r {
		if assert, ok := v.(string); ok {
			strings[i] = assert
		}
	}

	return
}

func (r NameForm) String() string {
	return r.Name.Value(0)
}

func (r SuperiorDITStructureRules) String() (str string) {
	switch len(r) {
	case 0:
		return
	case 1:
		return r[0].(*DITStructureRule).ID.string()
	}

	str += `( `
	for i := 0; i < len(r); i++ {
		str += r[i].(*DITStructureRule).ID.string() + ` $ `
	}
	if str[len(str)-3:] == ` $ ` {
		str = str[:len(str)-2]
	}
	str += `)`

	return
}

func (r SuperiorObjectClasses) String() (str string) {
	return genericList(r).ocs_oids_string()
}

func (r AuxiliaryObjectClasses) String() (str string) {
	return genericList(r).ocs_oids_string()
}

func (r PermittedAttributeTypes) String() (str string) {
	return genericList(r).attrs_oids_string()
}

func (r ProhibitedAttributeTypes) String() (str string) {
	return genericList(r).attrs_oids_string()
}

func (r RequiredAttributeTypes) String() (str string) {
	return genericList(r).attrs_oids_string()
}

func (r Applies) String() (str string) {
	return genericList(r).attrs_oids_string()
}

func (r genericList) attrs_oids_string() (str string) {
	switch len(r) {
	case 0:
		return
	case 1:
		return r[0].(*AttributeType).Name.Value(0)
	}

	str += `( `
	for i := 0; i < len(r); i++ {
		str += r[i].(*AttributeType).Name.Value(0) + ` $ `
	}
	if str[len(str)-3:] == ` $ ` {
		str = str[:len(str)-2]
	}
	str += `)`

	return
}

func (r genericList) ocs_oids_string() (str string) {
	switch len(r) {
	case 0:
		return
	case 1:
		return r[0].(*ObjectClass).Name.Value(0)
	}

	str += `( `
	for i := 0; i < len(r); i++ {
		str += r[i].(*ObjectClass).Name.Value(0) + ` $ `
	}
	if str[len(str)-3:] == ` $ ` {
		str = str[:len(str)-2]
	}
	str += `)`

	return
}

func (r StructuralObjectClass) String() (oc string) {
	if r.Name.IsZero() && r.OID.IsZero() {
		return
	}

	switch {
	case !r.Name.IsZero():
		return r.Name.Value(0)
	}

	return r.OID.string()
}

func (r RuleID) string() string {
	return itoa(int(r))
}

func (r *AttributeType) unmarshal(namesok bool) (def string, err error) {
	if err = r.validate(); err != nil {
		err = raise(invalidUnmarshal, err.Error())
		return
	}

	WHSP := ` `

	def += `(` + WHSP + r.OID.string() // will never be zero length

	if !r.Name.IsZero() {
		def += WHSP + r.Name.Label()
		def += WHSP + r.Name.String()
	}

	if !r.Description.IsZero() {
		def += WHSP + r.Description.Label()
		def += WHSP + r.Description.string()
	}

	if r.bools.is(Obsolete) {
		def += WHSP + Obsolete.string()
	}

	if !r.SuperType.IsZero() {
		def += WHSP + r.SuperType.Label()
		def += WHSP + r.SuperType.Name.Value(0)
	}

	if !r.Syntax.IsZero() {
		def += WHSP + r.Syntax.Label()
		def += WHSP + r.Syntax.OID.string()
	}

	if !r.Equality.IsZero() {
		def += WHSP + r.Equality.Label()
		def += WHSP + r.Equality.Name.Value(0)
	}

	if !r.Ordering.IsZero() {
		def += WHSP + r.Ordering.Label()
		def += WHSP + r.Ordering.Name.Value(0)
	}

	if !r.Substring.IsZero() {
		def += WHSP + r.Substring.Label()
		def += WHSP + r.Substring.Name.Value(0)
	}

	if r.Usage != UserApplication {
		def += WHSP + r.Usage.Label()
		def += WHSP + r.Usage.String()
	}

	if r.bools.is(NoUserModification) {
		def += WHSP + NoUserModification.string()
	}

	if !r.Extensions.IsZero() {
		def += WHSP + r.Extensions.String()
	}

	def += WHSP + `)`

	return
}

func (r *ObjectClass) unmarshal(namesok bool) (def string, err error) {
	if err = r.validate(); err != nil {
		err = raise(invalidUnmarshal, err.Error())
		return
	}

	WHSP := ` `

	def += `(` + WHSP + r.OID.string() // will never be zero length

	if !r.Name.IsZero() {
		def += WHSP + r.Name.Label()
		def += WHSP + r.Name.String()
	}

	if !r.Description.IsZero() {
		def += WHSP + r.Description.Label()
		def += WHSP + r.Description.string()
	}

	if r.is(Obsolete) {
		def += WHSP + r.bools.Obsolete()
	}

	if !r.SuperClass.IsZero() {
		def += WHSP + r.SuperClass.Label()
		def += WHSP + r.SuperClass.String()
	}

	// Kind will never be zero
	def += WHSP + r.Kind.String()

	if !r.Must.IsZero() {
		def += WHSP + r.Must.Label()
		def += WHSP + r.Must.String()
	}
	if !r.May.IsZero() {
		def += WHSP + r.May.Label()
		def += WHSP + r.May.String()
	}

	if !r.Extensions.IsZero() {
		def += WHSP + r.Extensions.String()
	}

	def += WHSP + `)`

	return
}

func (r *LDAPSyntax) unmarshal(namesok bool) (def string, err error) {
	if err = r.validate(); err != nil {
		err = raise(invalidUnmarshal, err.Error())
		return
	}

	WHSP := ` `

	def += `(` + WHSP + r.OID.string() // will never be zero length

	// Description will never be zero
	def += WHSP + r.Description.Label()
	def += WHSP + r.Description.string()

	if r.bools.enabled(Obsolete) {
		def += WHSP + r.bools.Obsolete()
	}

	if !r.Extensions.IsZero() {
		def += WHSP + r.Extensions.String()
	}

	def += WHSP + `)`

	return
}

func (r *MatchingRule) unmarshal(namesok bool) (def string, err error) {
	if err = r.validate(); err != nil {
		err = raise(invalidUnmarshal, err.Error())
		return
	}

	WHSP := ` `

	def += `(` + WHSP + r.OID.string() // will never be zero length

	if !r.Name.IsZero() {
		def += WHSP + r.Name.Label()
		def += WHSP + r.Name.String()
	}

	if !r.Description.IsZero() {
		def += WHSP + r.Description.Label()
		def += WHSP + r.Description.string()
	}

	if r.is(Obsolete) {
		def += WHSP + r.bools.Obsolete()
	}

	// Syntax will never be zero
	def += WHSP + r.Syntax.Label()
	def += WHSP + r.Syntax.OID.string()

	if !r.Extensions.IsZero() {
		def += WHSP + r.Extensions.String()
	}

	def += WHSP + `)`

	return
}

func (r *MatchingRuleUse) unmarshal(namesok bool) (def string, err error) {
	if err = r.validate(); err != nil {
		err = raise(invalidUnmarshal, err.Error())
		return
	}

	WHSP := ` `

	def += `(` + WHSP + r.OID.string() // will never be zero length

	if !r.Name.IsZero() {
		def += WHSP + r.Name.Label()
		def += WHSP + r.Name.String()
	}

	if !r.Description.IsZero() {
		def += WHSP + r.Description.Label()
		def += WHSP + r.Description.string()
	}

	if r.is(Obsolete) {
		def += WHSP + r.bools.Obsolete()
	}

	// Applies will never be zero
	def += WHSP + r.Applies.Label()
	def += WHSP + r.Applies.String()

	if !r.Extensions.IsZero() {
		def += WHSP + r.Extensions.String()
	}

	def += WHSP + `)`

	return
}

func (r *NameForm) unmarshal(namesok bool) (def string, err error) {
	if err = r.validate(); err != nil {
		err = raise(invalidUnmarshal, err.Error())
		return
	}

	WHSP := ` `

	def += `(` + WHSP + r.OID.string() // will never be zero length

	if !r.Name.IsZero() {
		def += WHSP + r.Name.Label()
		def += WHSP + r.Name.String()
	}

	if !r.Description.IsZero() {
		def += WHSP + r.Description.Label()
		def += WHSP + r.Description.string()
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

	if r.bools.enabled(Obsolete) {
		def += WHSP + r.bools.Obsolete()
	}

	if !r.Extensions.IsZero() {
		def += WHSP + r.Extensions.String()
	}

	def += WHSP + `)`

	return
}

func (r *DITContentRule) unmarshal(namesok bool) (def string, err error) {
	if err = r.validate(); err != nil {
		err = raise(invalidUnmarshal, err.Error())
		return
	}

	WHSP := ` `

	def += `(` + WHSP + r.OID.string() // will never be zero length

	if !r.Name.IsZero() {
		def += WHSP + r.Name.Label()
		def += WHSP + r.Name.String()
	}

	if !r.Description.IsZero() {
		def += WHSP + r.Description.Label()
		def += WHSP + r.Description.string()
	}

	if r.bools.enabled(Obsolete) {
		def += WHSP + r.bools.Obsolete()
	}

	if !r.Aux.IsZero() {
		def += WHSP + r.Aux.Label()
		def += WHSP + r.Aux.String()
	}

	if !r.Must.IsZero() {
		def += WHSP + r.Must.Label()
		def += WHSP + r.Must.String()
	}

	if !r.May.IsZero() {
		def += WHSP + r.May.Label()
		def += WHSP + r.May.String()
	}

	if !r.Not.IsZero() {
		def += WHSP + r.Not.Label()
		def += WHSP + r.Not.String()
	}

	if !r.Extensions.IsZero() {
		def += WHSP + r.Extensions.String()
	}

	def += WHSP + `)`

	return
}

func (r *DITStructureRule) unmarshal(namesok bool) (def string, err error) {
	if err = r.validate(); err != nil {
		err = raise(invalidUnmarshal, err.Error())
		return
	}

	WHSP := ` `

	def += `(` + WHSP + r.ID.string() // will never be zero length

	if !r.Name.IsZero() {
		def += WHSP + r.Name.Label()
		def += WHSP + r.Name.String()
	}

	if !r.Description.IsZero() {
		def += WHSP + r.Description.Label()
		def += WHSP + r.Description.string()
	}

	if r.bools.enabled(Obsolete) {
		def += WHSP + r.bools.Obsolete()
	}

	// Form will never be zero
	def += WHSP + r.Form.Label()
	def += WHSP + r.Form.String()

	if !r.SuperiorRules.IsZero() {
		def += WHSP + r.SuperiorRules.Label()
		def += WHSP + r.SuperiorRules.String()
	}

	if !r.Extensions.IsZero() {
		def += WHSP + r.Extensions.String()
	}

	def += WHSP + `)`

	return
}

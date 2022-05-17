package schemax

/*
NewSubschema returns a partially-initialized instance of *Subschema.
*/
func NewSubschema() (s *Subschema) {
	return &Subschema{
		LSC:  NewLDAPSyntaxes(),
		MRC:  NewMatchingRules(),
		ATC:  NewAttributeTypes(),
		MRUC: NewMatchingRuleUses(),
		OCC:  NewObjectClasses(),
		DCRC: NewDITContentRules(),
		NFC:  NewNameForms(),
		DSRC: NewDITStructureRules(),
	}
}

/*
Subschema provides high-level definition management, storage of all definitions of schemata and convenient method wrappers.

The first field of this type, "DN", contains a string representation of the appropriate subschema distinguished name.  Oftentimes, this will be something like "cn=schema", or "cn=subschema" but may vary between directory implementations.

All remaining fields within this type are based on the collection (slice) types defined in this package.  For example, Subschema "NFC" field is of the NameForms Collection type, and is intended to store all successfully-parsed "nameForm" definitions.

The user must initialize these map-based fields before use, whether they do so themselves manually, OR by use of a suitable "populator function", such as the PopulateDefaultLDAPSyntaxes() function to populate the "LSC" field.

The purpose of this type is to centralize all of the relevant slices that the user would otherwise have to manage individually.  This is useful for portability reasons, and allows an entire library of X.501 schema definitions to be stored within a single object.

This type also provides high-level oversight of what would otherwise be isolated low-level operations that may or may not be invalid.  For example, when operating in a purely low-level fashion (without the use of *Subschema), there is nothing stopping a user from adding a so-called "RequiredAttributeTypes" slice member to a "ProhibitedAttributeTypes" slice.  Each of those slices is throughly unaware of the other.   However, when conducting this operation via a convenient method extended by *Subschema, additional correlative checks can (and will!) be conducted to avoid such invalid actions.

Overall, use of *Subschema makes generalized use of this package slightly easier but is NOT necessarily required in all situations.
*/
type Subschema struct {
	DN   string                     // often "cn=schema" or "cn=subschema", varies per imp.
	LSC  LDAPSyntaxCollection       // LDAP Syntaxes
	MRC  MatchingRuleCollection     // Matching Rules
	ATC  AttributeTypeCollection    // Attribute Types
	MRUC MatchingRuleUseCollection  // Matching Rule "Uses"
	OCC  ObjectClassCollection      // Object Classes
	DCRC DITContentRuleCollection   // DIT Content Rules
	NFC  NameFormCollection         // Name Forms
	DSRC DITStructureRuleCollection // DIT Structure Rules
}

/*
schema.go deals with the high-level interactions within the subschema as a whole.  A higher degree of referential integrity is present, as individual list and map contents are verified with other such collections to avoid (or fail in the presence of) certain mutexes (e.g.: An *AttributeType present within both ProhibitedAttributeTypes and RequiredAttributeTypes lists associated with a *DITContentRule instance that resides within a single *Subschema instance).
*/

/*
Set will take the assigned interface{} x and attempt to store (register) it within the appropriate slice member within the receiver instance of the receiver.  A boolean value is returned indicative of success.

This method is only useful when NOT marshaling raw definitions into type instances (e.g.: if one is manually crafting a specific definition manually).
*/
func (r *Subschema) set(x interface{}) (ok bool) {
	switch tv := x.(type) {
	case *AttributeType:
		if z := r.ATC.Get(tv.OID.String()); z == nil {
			r.ATC.Set(tv)
			ok = tv != nil
		}
	case *ObjectClass:
		if z := r.OCC.Get(tv.OID.String()); z == nil {
			r.OCC.Set(tv)
			ok = tv != nil
		}
	case *LDAPSyntax:
		if z := r.LSC.Get(tv.OID.String()); z == nil {
			r.LSC.Set(tv)
			ok = tv != nil
		}
	case *MatchingRule:
		if z := r.MRC.Get(tv.OID.String()); z == nil {
			r.MRC.Set(tv)
			ok = tv != nil
		}
	case *MatchingRuleUse:
		if z := r.MRUC.Get(tv.OID.String()); z == nil {
			r.MRUC.Set(tv)
			ok = tv != nil
		}
	case *DITContentRule:
		if z := r.DCRC.Get(tv.OID.String()); z == nil {
			r.DCRC.Set(tv)
			ok = tv != nil
		}
	case *DITStructureRule:
		if z := r.DSRC.Get(tv.ID.String()); z == nil {
			r.DSRC.Set(tv)
			ok = tv != nil
		}
	case *NameForm:
		if z := r.NFC.Get(tv.OID.String()); z == nil {
			r.NFC.Set(tv)
			ok = tv != nil
		}
	}

	return
}

/*
PopulateDefaultLDAPSyntaxes is a Subschema-housed convenience method for PopulateDefaultLDAPSyntaxes, and includes internal indexing to maintain proper ordering of definitions.
*/
func (r *Subschema) PopulateDefaultLDAPSyntaxe() {
	r.LSC = PopulateDefaultLDAPSyntaxes()
}

/*
PopulateDefaultMatchingRules is a Subschema-housed convenience method for PopulateDefaultMatchingRules, and includes internal indexing to maintain proper ordering of definitions.
*/
func (r *Subschema) PopulateDefaultMatchingRules() {
	r.MRC = PopulateDefaultMatchingRules()
}

/*
PopulateDefaultObjectClasses is a Subschema-housed convenience method for PopulateDefaultObjectClasses, and includes internal indexing to maintain proper ordering of definitions.
*/
func (r *Subschema) PopulateDefaultObjectClasses() {
	r.OCC = PopulateDefaultObjectClasses()
}

/*
PopulateDefaultAttributeTypes is a Subschema-housed convenience method for PopulateDefaultAttributeTypes, and includes internal indexing to maintain proper ordering of definitions.
*/
func (r *Subschema) PopulateDefaultAttributeTypes() {
	r.ATC = PopulateDefaultAttributeTypes()
}

/*
GetAttributeType is a convenient wrapper method for AttributeType.Get.  Using this method, the user need not worry about providing an Aliases manually.
*/
func (r *Subschema) GetAttributeType(x string) *AttributeType {
	if len(x) == 0 {
		return nil
	}

	return r.ATC.Get(x)
}

/*
Set will assign x to the receiver, storing the content within the appropriate manifest.  A boolean value is returned indicative of success.
*/
func (r *Subschema) SetAttributeType(x *AttributeType) (ok bool) {
	return r.set(x)
}

/*
GetAttributeType is a convenient wrapper method for AttributeType.Get.  Using this method, the user need not worry about providing an Aliases manually.
*/
func (r Subschema) GetObjectClass(x string) *ObjectClass {
	if len(x) == 0 {
		return nil
	}

	return r.OCC.Get(x)
}

/*
Set will assign x to the receiver, storing the content within the appropriate manifest.  A boolean value is returned indicative of success.
*/
func (r *Subschema) SetObjectClass(x *ObjectClass) (ok bool) {
	return r.set(x)
}

/*
GetNameForm is a convenient wrapper method for NameForms.Get.  Using this method, the user need not worry about providing an Aliases manually.
*/
func (r Subschema) GetNameForm(x string) *NameForm {
	if len(x) == 0 {
		return nil
	}

	return r.NFC.Get(x)
}

/*
Set will assign x to the receiver, storing the content within the appropriate manifest.  A boolean value is returned indicative of success.
*/
func (r *Subschema) SetNameForm(x *NameForm) (ok bool) {
	return r.set(x)
}

/*
GetDITContentRule is a convenient wrapper method for DITContentRules.Get.  Using this method, the user need not worry about providing an Aliases manually.
*/
func (r Subschema) GetDITContentRule(x string) *DITContentRule {
	if len(x) == 0 {
		return nil
	}

	return r.DCRC.Get(x)
}

/*
Set will assign x to the receiver, storing the content within the appropriate manifest.  A boolean value is returned indicative of success.
*/
func (r *Subschema) SetDITContentRule(x *DITContentRule) (ok bool) {
	return r.set(x)
}

/*
GetDITStructureRule is a convenient wrapper method for DITStructureRules.Get.
*/
func (r Subschema) GetDITStructureRule(x string) *DITStructureRule {
	if len(x) == 0 {
		return nil
	}

	return r.DSRC.Get(x)
}

/*
Set will assign x to the receiver, storing the content within the appropriate manifest.  A boolean value is returned indicative of success.
*/
func (r *Subschema) SetDITStructureRule(x *DITStructureRule) (ok bool) {
	return r.set(x)
}

/*
GetLDAPSyntax is a convenient wrapper method for LDAPSyntaxes.Get.  Using this method, the user need not worry about providing an Aliases manually.
*/
func (r Subschema) GetLDAPSyntax(x string) *LDAPSyntax {
	if len(x) == 0 {
		return nil
	}

	return r.LSC.Get(x)
}

/*
Set will assign x to the receiver, storing the content within the appropriate manifest.  A boolean value is returned indicative of success.
*/
func (r *Subschema) SetLDAPSyntax(x *LDAPSyntax) (ok bool) {
	return r.set(x)
}

/*
GetMatchingRule is a convenient wrapper method for MatchingRule.Get.  Using this method, the user need not worry about providing an Aliases manually.
*/
func (r Subschema) GetMatchingRule(x string) *MatchingRule {
	if len(x) == 0 {
		return nil
	}

	return r.MRC.Get(x)
}

/*
Set will assign x to the receiver, storing the content within the appropriate manifest.  A boolean value is returned indicative of success.
*/
func (r *Subschema) SetMatchingRule(x *MatchingRule) (ok bool) {
	return r.set(x)
}

/*
GetMatchingRuleUse is a convenient wrapper method for MatchingRuleUses.Get.  Using this method, the user need not worry about providing an Aliases manually.
*/
func (r Subschema) GetMatchingRuleUse(x string) *MatchingRuleUse {
	if len(x) == 0 {
		return nil
	}

	return r.MRUC.Get(x)
}

/*
Set will assign x to the receiver, storing the content within the appropriate manifest.  A boolean value is returned indicative of success.
*/
func (r *Subschema) SetMatchingRuleUse(x *MatchingRuleUse) (ok bool) {
	return r.set(x)
}

/*
MarshalAttributeType will attempt to perform a marshal operation upon the provided string definition under the assumption that it represents an Attribute Type.  An error is returned in the event that a parsing failure occurs in any way. No *AttributeType instance is stored in such an event.

In the event a successful marshal, a new *AttributeType instance is catalogued within the Subschema.ATC (AttributeTypes) struct field.

This method requires the complete population of the following collections before-hand:

• LDAPSyntaxes (Subschema.LSC)

• MatchingRules (Subschema.MRC)

• AttributeTypes (Subschema.ATC, if a particular definition a sub type of another type)
*/
func (r *Subschema) MarshalAttributeType(def string) (err error) {

	// prelaunch checks and balances
	if len(def) == 0 {
		return raise(invalidMarshal,
			"MarshalAttributeType: definition is zero-length")
	}

	// marshal it
	var x AttributeType
	if err = Marshal(
		def, &x, r.ATC,
		nil, r.LSC, r.MRC,
		nil, nil, nil, nil); err != nil {

		return
	}

	// register attribute (and refresh our
	// MatchingRuleUses instance)
	r.set(&x)

	return
}

/*
MarshalObjectClass will attempt to perform a marshal operation upon the provided string definition under the assumption that it represents an Object Class.  An error is returned in the event that a parsing failure occurs in any way. No *ObjectClass instance is stored in such an event.

In the event a successful marshal, a new *ObjectClass instance is catalogued within the Subschema.OCC (ObjectClass) struct field.

This method requires the complete population of the following collections before-hand:

• AttributeTypes (Subschema.ATC) beforehand

• ObjectClasses (Subschema.OCC, if a particular definition a sub class one or more other classes)
*/
func (r *Subschema) MarshalObjectClass(def string) (err error) {

	// prelaunch checks and balances
	if len(def) == 0 {
		return raise(invalidMarshal,
			"MarshalObjectClass: definition is zero-length")
	}

	// marshal it
	var x ObjectClass
	x.SuperClass = NewSuperiorObjectClasses()
	x.Must = NewRequiredAttributeTypes()
	x.May = NewPermittedAttributeTypes()

	if err = Marshal(
		def, &x, r.ATC,
		r.OCC, nil, nil,
		nil, nil, nil, nil); err != nil {

		return
	}

	// register it
	r.set(&x)

	return
}

/*
MarshalLDAPSyntax will attempt to perform a marshal operation upon the provided string definition under the assumption that it represents an LDAP Syntax.  An error is returned in the event that a parsing failure occurs in any way. No *LDAPSyntax instance is stored in such an event.

In the event a successful marshal, a new *LDAPSyntax instance is catalogued within the Subschema.LSC (LDAPSyntaxes) struct field.

It is important to note that it is unusual for end users to "invent" new LDAP syntaxes. Most directory implementations do not allow the designation of such definitions arbitrarily.  As such, this method (and those that are similar) is intended only for the marshaling and/or registration of "official LDAP syntaxes" (i.e.: those that are defined in an accepted RFC or Internet-Draft).

Users are strongly advised to simply set the Subschema.LSC field using the DefaultLDAPSyntaxes global variable.

LDAP Syntaxes are wholly independent in the referential sense.  They are dependent upon no other schema definition type, thus they should ALWAYS be populated BEFORE all other definition types.
*/
func (r *Subschema) MarshalLDAPSyntax(def string) (err error) {

	// prelaunch checks and balances
	if len(def) == 0 {
		return raise(invalidMarshal,
			"MarshalLDAPSyntax: definition is zero-length")
	}

	// marshal it
	var x LDAPSyntax
	if err = Marshal(
		def, &x, nil,
		nil, r.LSC, nil,
		nil, nil, nil, nil); err != nil {

		return
	}

	// register it
	r.set(&x)

	return
}

/*
MarshalMatchingRule will attempt to perform a marshal operation upon the provided string definition under the assumption that it represents an Matching Rule.  An error is returned in the event that a parsing failure occurs in any way. No *MatchingRule instance is stored in such an event.

In the event a successful marshal, a new *MatchingRule instance is catalogued within the Subschema.MRC (MatchingRules) struct field.

It is important to note that it is unusual for end users to "invent" new matching rules. Most directory implementations do not allow the designation of such definitions arbitrarily.  As such, this method (and those that are similar) is intended only for the marshaling and/or registration of "official matching rules" (i.e.: those that are defined in an accepted RFC or Internet-Draft).

Users are strongly advised to simply set the Subschema.MRC field using the DefaultMatchingRules global variable.

This method requires the complete population of the LDAPSyntaxes (Subschema.LSC) beforehand.
*/
func (r *Subschema) MarshalMatchingRule(def string) (err error) {

	// prelaunch checks and balances
	if len(def) == 0 {
		return raise(invalidMarshal,
			"MarshalMatchingRule: definition is zero-length")
	}

	// marshal it
	var x MatchingRule
	if err = Marshal(
		def, &x, nil,
		nil, r.LSC, r.MRC,
		nil, nil, nil, nil); err != nil {

		return
	}

	// register it
	r.set(&x)

	return
}

/*
MarshalNameForm will attempt to perform a marshal operation upon the provided string definition under the assumption that it represents a Name Form.  An error is returned in the event that a parsing failure occurs in any way. No *NameForm instance is stored in such an event.

In the event a successful marshal, a new *NameForm instance is catalogued within the Subschema.NFC (NameForms) struct field.

This method requires the complete population of the following collections before-hand:

• AttributeTypes (Subschema.ATC)

• ObjectClasses (Subschema.OCC)
*/
func (r *Subschema) MarshalNameForm(def string) (err error) {

	// prelaunch checks and balances
	if len(def) == 0 {
		return raise(invalidMarshal,
			"MarshalNameForm: definition is zero-length")
	}

	// marshal it
	var x NameForm
	x.Must = NewRequiredAttributeTypes()
	x.May = NewPermittedAttributeTypes()

	if err = Marshal(
		def, &x, r.ATC,
		r.OCC, nil, nil,
		nil, nil, nil, r.NFC); err != nil {

		return
	}

	// register it
	r.set(&x)

	return
}

/*
MarshalDITStructureRule will attempt to perform a marshal operation upon the provided string definition under the assumption that it represents a DIT Structure Rule.  An error is returned in the event that a parsing failure occurs in any way. No *DITStructureRule instance is stored in such an event.

In the event a successful marshal, a new *DITStructureRule instance is catalogued within the Subschema.DSRC (DITStructureRules) struct field.

This method requires the complete population of the following collections before-hand:

• NameForms (Subschema.NFC)

• DITStructureRules (Subschema.DSRC, if superior DIT Structure Rules are referenced)
*/
func (r *Subschema) MarshalDITStructureRule(def string) (err error) {

	// prelaunch checks and balances
	if len(def) == 0 {
		return raise(invalidMarshal,
			"MarshalDITStructureRule: definition is zero-length")
	}

	// marshal it
	var x DITStructureRule
	x.SuperiorRules = NewSuperiorDITStructureRules()

	if err = Marshal(
		def, &x, nil,
		nil, nil, nil,
		nil, nil, r.DSRC, r.NFC); err != nil {

		return
	}

	// register it
	r.set(&x)

	return
}

/*
MarshalDITContentRule will attempt to perform a marshal operation upon the provided string definition under the assumption that it represents a DIT Content Rule.  An error is returned in the event that a parsing failure occurs in any way. No *DITContentRule instance is stored in such an event.

In the event a successful marshal, a new *DITContentRule instance is catalogued within the Subschema.DCRC (DITContentRules) struct field.

Note that the OID assigned to the raw DIT Content Rule definition MUST correlate to the registered OID of a STRUCTURAL *ObjectClass instance.

This method requires the complete population of the following collections before-hand:

• AttributeTypes (Subschema.ATC)

• ObjectClasses (Subschema.OCC)
*/
func (r *Subschema) MarshalDITContentRule(def string) (err error) {

	// prelaunch checks and balances
	if len(def) == 0 {
		return raise(invalidMarshal,
			"MarshalDITContentRule: definition is zero-length")
	}

	// marshal it
	var x DITContentRule
	x.Aux = NewAuxiliaryObjectClasses()
	x.Must = NewRequiredAttributeTypes()
	x.May = NewPermittedAttributeTypes()
	x.Not = NewProhibitedAttributeTypes()

	if err = Marshal(
		def, &x, r.ATC,
		r.OCC, nil, nil,
		nil, r.DCRC, nil, nil); err != nil {

		return
	}

	// register it
	r.set(&x)

	return
}

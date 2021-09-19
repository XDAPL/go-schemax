package schemax

/*
schema.go deals with the high-level interactions within the subschema as a whole.  A higher degree of referential integrity is present, as individual list and map contents are verified with other such collections to avoid (or fail in the presence of) certain mutexes (e.g.: An *AttributeType present within both ProhibitedAttributeTypes and RequiredAttributeTypes lists associated with a *DITContentRule instance that resides within a single *Subschema instance).
*/

/*
Set will take the assigned interface{} x and attempt to store (register) it within the appropriate Manifest within the receiver instance of the receiver.  A boolean value is returned indicative of success.

This method is only useful when NOT marshaling raw definitions into type instances (e.g.: if one is manually crafting a specific definition manually).
*/
func (r *Subschema) set(x interface{}) (ok bool) {
	switch tv := x.(type) {
	case *AttributeType:
		if z := r.ATM.Get(tv.OID.string(), nil); z == nil {
			r.ATM.Set(tv)
			ok = tv != nil
		}
	case *ObjectClass:
		if z := r.OCM.Get(tv.OID.string(), nil); z == nil {
			r.OCM.Set(tv)
			ok = tv != nil
		}
	case *LDAPSyntax:
		if z := r.LSM.Get(tv.OID.string(), nil); z == nil {
			r.LSM.Set(tv)
			ok = tv != nil
		}
	case *MatchingRule:
		if z := r.MRM.Get(tv.OID.string(), nil); z == nil {
			r.MRM.Set(tv)
			ok = tv != nil
		}
	case *MatchingRuleUse:
		if z := r.MRUM.Get(tv.OID.string(), nil); z == nil {
			r.MRUM.Set(tv)
			ok = tv != nil
		}
	case *DITContentRule:
		if z := r.DCRM.Get(tv.OID.string(), nil); z == nil {
			r.DCRM.Set(tv)
			ok = tv != nil
		}
	case *DITStructureRule:
		if z := r.DSRM.Get(tv.ID.string()); z == nil {
			r.DSRM.Set(tv)
			ok = tv != nil
		}
	case *NameForm:
		if z := r.NFM.Get(tv.OID.string(), nil); z == nil {
			r.NFM.Set(tv)
			ok = tv != nil
		}
	}

	return
}

/*
PopulateDefaultLDAPSyntaxesManifest is a Subschema-housed convenience method for PopulateDefaultLDAPSyntaxesManifest, and includes internal indexing to maintain proper ordering of definitions.
*/
func (r *Subschema) PopulateDefaultLDAPSyntaxesManifest() {
	r.LSM = PopulateDefaultLDAPSyntaxesManifest()
}

/*
PopulateDefaultMatchingRulesManifest is a Subschema-housed convenience method for PopulateDefaultMatchingRulesManifest, and includes internal indexing to maintain proper ordering of definitions.
*/
func (r *Subschema) PopulateDefaultMatchingRulesManifest() {
	r.MRM = PopulateDefaultMatchingRulesManifest()
}

/*
PopulateDefaultObjectClassesManifest is a Subschema-housed convenience method for PopulateDefaultObjectClassesManifest, and includes internal indexing to maintain proper ordering of definitions.
*/
func (r *Subschema) PopulateDefaultObjectClassesManifest() {
	r.OCM = PopulateDefaultObjectClassesManifest()
}

/*
PopulateDefaultAttributeTypesManifest is a Subschema-housed convenience method for PopulateDefaultAttributeTypesManifest, and includes internal indexing to maintain proper ordering of definitions.
*/
func (r *Subschema) PopulateDefaultAttributeTypesManifest() {
	r.ATM = PopulateDefaultAttributeTypesManifest()
}

/*
GetAttributeType is a convenient wrapper method for AttributeTypeManifest.Get.  Using this method, the user need not worry about providing an AliasesManifest manually.
*/
func (r *Subschema) GetAttributeType(x string) *AttributeType {
	if len(x) == 0 {
		return nil
	}

	return r.ATM.Get(x, r.ALM)
}

/*
SetAlias is a convenient wrapper method for AliasesManifest.Set, and will set the underlying manifest within the receiver with the provided name and OID.
*/
func (r *Subschema) SetAlias(a, o string) {
	r.ALM.Set(a, o)
}

/*
Set will assign x to the receiver, storing the content within the appropriate manifest.  A boolean value is returned indicative of success.
*/
func (r *Subschema) SetAttributeType(x *AttributeType) (ok bool) {
	return r.set(x)
}

/*
GetAttributeType is a convenient wrapper method for AttributeTypeManifest.Get.  Using this method, the user need not worry about providing an AliasesManifest manually.
*/
func (r *Subschema) GetObjectClass(x string) *ObjectClass {
	if len(x) == 0 {
		return nil
	}

	return r.OCM.Get(x, r.ALM)
}

/*
Set will assign x to the receiver, storing the content within the appropriate manifest.  A boolean value is returned indicative of success.
*/
func (r *Subschema) SetObjectClass(x *ObjectClass) (ok bool) {
	return r.set(x)
}

/*
GetNameForm is a convenient wrapper method for NameFormsManifest.Get.  Using this method, the user need not worry about providing an AliasesManifest manually.
*/
func (r *Subschema) GetNameForm(x string) *NameForm {
	if len(x) == 0 {
		return nil
	}

	return r.NFM.Get(x, r.ALM)
}

/*
Set will assign x to the receiver, storing the content within the appropriate manifest.  A boolean value is returned indicative of success.
*/
func (r *Subschema) SetNameForm(x *NameForm) (ok bool) {
	return r.set(x)
}

/*
GetDITContentRule is a convenient wrapper method for DITContentRulesManifest.Get.  Using this method, the user need not worry about providing an AliasesManifest manually.
*/
func (r *Subschema) GetDITContentRule(x string) *DITContentRule {
	if len(x) == 0 {
		return nil
	}

	return r.DCRM.Get(x, r.ALM)
}

/*
Set will assign x to the receiver, storing the content within the appropriate manifest.  A boolean value is returned indicative of success.
*/
func (r *Subschema) SetDITContentRule(x *DITContentRule) (ok bool) {
	return r.set(x)
}

/*
GetDITStructureRule is a convenient wrapper method for DITStructureRulesManifest.Get.
*/
func (r *Subschema) GetDITStructureRule(x string) *DITStructureRule {
	if len(x) == 0 {
		return nil
	}

	return r.DSRM.Get(x)
}

/*
Set will assign x to the receiver, storing the content within the appropriate manifest.  A boolean value is returned indicative of success.
*/
func (r *Subschema) SetDITStructureRule(x *DITStructureRule) (ok bool) {
	return r.set(x)
}

/*
GetLDAPSyntax is a convenient wrapper method for LDAPSyntaxesManifest.Get.  Using this method, the user need not worry about providing an AliasesManifest manually.
*/
func (r *Subschema) GetLDAPSyntax(x string) *LDAPSyntax {
	if len(x) == 0 {
		return nil
	}

	return r.LSM.Get(x, r.ALM)
}

/*
Set will assign x to the receiver, storing the content within the appropriate manifest.  A boolean value is returned indicative of success.
*/
func (r *Subschema) SetLDAPSyntax(x *LDAPSyntax) (ok bool) {
	return r.set(x)
}

/*
GetMatchingRule is a convenient wrapper method for MatchingRuleManifest.Get.  Using this method, the user need not worry about providing an AliasesManifest manually.
*/
func (r *Subschema) GetMatchingRule(x string) *MatchingRule {
	if len(x) == 0 {
		return nil
	}

	return r.MRM.Get(x, r.ALM)
}

/*
Set will assign x to the receiver, storing the content within the appropriate manifest.  A boolean value is returned indicative of success.
*/
func (r *Subschema) SetMatchingRule(x *MatchingRule) (ok bool) {
	return r.set(x)
}

/*
GetMatchingRuleUse is a convenient wrapper method for MatchingRuleUsesManifest.Get.  Using this method, the user need not worry about providing an AliasesManifest manually.
*/
func (r *Subschema) GetMatchingRuleUse(x string) *MatchingRuleUse {
	if len(x) == 0 {
		return nil
	}

	return r.MRUM.Get(x, r.ALM)
}

/*
Set will assign x to the receiver, storing the content within the appropriate manifest.  A boolean value is returned indicative of success.
*/
func (r *Subschema) SetMatchingRuleUse(x *MatchingRuleUse) (ok bool) {
	return r.set(x)
}

/*
MarshalAttributeType will attempt to perform a marshal operation upon the provided string definition under the assumption that it represents an Attribute Type.  An error is returned in the event that a parsing failure occurs in any way. No *AttributeType instance is stored in such an event.

In the event a successful marshal, a new *AttributeType instance is catalogued within the Subschema.ATM (AttributeTypesManifest) struct field.

This method requires the complete population of the following Manifests before-hand:

  - LDAPSyntaxesManifest (Subschema.LSM)
  - MatchingRulesManifest (Subschema.MRM)
  - AttributeTypesManifest (Subschema.ATM, if a particular definition a sub type of another type)
*/
func (r *Subschema) MarshalAttributeType(def string) (err error) {

	// prelaunch checks and balances
	if len(def) == 0 {
		return raise(invalidMarshal,
			"MarshalAttributeType: definition is zero-length")
	} else if r.ATM == nil {
		return raise(invalidMarshal,
			"MarshalAttributeType: %T.ATM (%T) is nil (uninitialized)",
			r, r.ATM)
	} else if r.LSM == nil {
		return raise(invalidMarshal,
			"MarshalAttributeType: %T.LSM (%T) is nil (uninitialized)",
			r, r.LSM)
	} else if r.MRM == nil {
		return raise(invalidMarshal,
			"MarshalAttributeType: %T.MRM (%T) is nil (uninitialized)",
			r, r.MRM)
	}

	// marshal it
	var x AttributeType
	if err = Marshal(
		def, &x, r.ALM,
		r.ATM, nil, r.LSM,
		r.MRM, nil, nil,
		nil, nil); err != nil {

		return
	}

	// register attribute (and refresh our
	// MatchingRuleUsesManifest instance)
	r.set(&x)

	return
}

/*
MarshalObjectClass will attempt to perform a marshal operation upon the provided string definition under the assumption that it represents an Object Class.  An error is returned in the event that a parsing failure occurs in any way. No *ObjectClass instance is stored in such an event.

In the event a successful marshal, a new *ObjectClass instance is catalogued within the Subschema.OCM (ObjectClassManifest) struct field.

This method requires the complete population of the following Manifests before-hand:

  - AttributeTypesManifest (Subschema.ATM) beforehand
  - ObjectClassesManifest (Subschema.OCM, if a particular definition a sub class one or more other classes)
*/
func (r *Subschema) MarshalObjectClass(def string) (err error) {

	// prelaunch checks and balances
	if len(def) == 0 {
		return raise(invalidMarshal,
			"MarshalObjectClass: definition is zero-length")
	} else if r.OCM == nil {
		return raise(invalidMarshal,
			"MarshalObjectClass: %T.OCM (%T) is nil (uninitialized)",
			r, r.OCM)
	} else if r.ATM == nil {
		return raise(invalidMarshal,
			"MarshalObjectClass: %T.ATM (%T) is nil (uninitialized)",
			r, r.ATM)
	}

	// marshal it
	var x ObjectClass
	if err = Marshal(
		def, &x, r.ALM,
		r.ATM, r.OCM, nil,
		nil, nil, nil,
		nil, nil); err != nil {

		return
	}

	// register it
	r.set(&x)

	return
}

/*
MarshalLDAPSyntax will attempt to perform a marshal operation upon the provided string definition under the assumption that it represents an LDAP Syntax.  An error is returned in the event that a parsing failure occurs in any way. No *LDAPSyntax instance is stored in such an event.

In the event a successful marshal, a new *LDAPSyntax instance is catalogued within the Subschema.LSM (LDAPSyntaxesManifest) struct field.

It is important to note that it is unusual for end users to "invent" new LDAP syntaxes. Most directory implementations do not allow the designation of such definitions arbitrarily.  As such, this method (and those that are similar) is intended only for the marshaling and/or registration of "official LDAP syntaxes" (i.e.: those that are defined in an accepted RFC or Internet-Draft).

Users are strongly advised to simply set the Subschema.LSM field using the DefaultLDAPSyntaxesManifest global variable.

LDAP Syntaxes are wholly independent in the referential sense.  They are dependent upon no other schema definition type, thus they should ALWAYS be populated BEFORE all other definition types.
*/
func (r *Subschema) MarshalLDAPSyntax(def string) (err error) {

	// prelaunch checks and balances
	if len(def) == 0 {
		return raise(invalidMarshal,
			"MarshalLDAPSyntax: definition is zero-length")
	} else if r.LSM == nil {
		return raise(invalidMarshal,
			"MarshalLDAPSyntax: %T.LSM (%T) is nil (uninitialized)",
			r, r.LSM)
	}

	// marshal it
	var x LDAPSyntax
	if err = Marshal(
		def, &x, r.ALM,
		nil, nil, r.LSM,
		nil, nil, nil,
		nil, nil); err != nil {

		return
	}

	// register it
	r.set(&x)

	return
}

/*
MarshalMatchingRule will attempt to perform a marshal operation upon the provided string definition under the assumption that it represents an Matching Rule.  An error is returned in the event that a parsing failure occurs in any way. No *MatchingRule instance is stored in such an event.

In the event a successful marshal, a new *MatchingRule instance is catalogued within the Subschema.MRM (MatchingRulesManifest) struct field.

It is important to note that it is unusual for end users to "invent" new matching rules. Most directory implementations do not allow the designation of such definitions arbitrarily.  As such, this method (and those that are similar) is intended only for the marshaling and/or registration of "official matching rules" (i.e.: those that are defined in an accepted RFC or Internet-Draft).

Users are strongly advised to simply set the Subschema.MRM field using the DefaultMatchingRulesManifest global variable.

This method requires the complete population of the LDAPSyntaxesManifest (Subschema.LSM) beforehand.
*/
func (r *Subschema) MarshalMatchingRule(def string) (err error) {

	// prelaunch checks and balances
	if len(def) == 0 {
		return raise(invalidMarshal,
			"MarshalMatchingRule: definition is zero-length")
	} else if r.LSM == nil {
		return raise(invalidMarshal,
			"MarshalMatchingRule: %T.LSM (%T) is nil (uninitialized)",
			r, r.LSM)
	} else if r.MRM == nil {
		return raise(invalidMarshal,
			"MarshalMatchingRule: %T.MRM (%T) is nil (uninitialized)",
			r, r.MRM)
	}

	// marshal it
	var x MatchingRule
	if err = Marshal(
		def, &x, r.ALM,
		nil, nil, r.LSM,
		r.MRM, nil, nil,
		nil, nil); err != nil {

		return
	}

	// register it
	r.set(&x)

	return
}

/*
MarshalNameForm will attempt to perform a marshal operation upon the provided string definition under the assumption that it represents a Name Form.  An error is returned in the event that a parsing failure occurs in any way. No *NameForm instance is stored in such an event.

In the event a successful marshal, a new *NameForm instance is catalogued within the Subschema.NFM (NameFormsManifest) struct field.

This method requires the complete population of the following Manifests before-hand:

  - AttributeTypesManifest (Subschema.ATM)
  - ObjectClassesManifest (Subschema.OCM)
*/
func (r *Subschema) MarshalNameForm(def string) (err error) {

	// prelaunch checks and balances
	if len(def) == 0 {
		return raise(invalidMarshal,
			"MarshalNameForm: definition is zero-length")
	} else if r.ATM == nil {
		return raise(invalidMarshal,
			"MarshalNameForm: %T.ATM (%T) is nil (uninitialized)",
			r, r.ATM)
	} else if r.OCM == nil {
		return raise(invalidMarshal,
			"MarshalNameForm: %T.OCM (%T) is nil (uninitialized)",
			r, r.OCM)
	} else if r.NFM == nil {
		return raise(invalidMarshal,
			"MarshalNameForm: %T.NFM (%T) is nil (uninitialized)",
			r, r.NFM)
	}

	// marshal it
	var x NameForm
	if err = Marshal(
		def, &x, r.ALM,
		r.ATM, r.OCM, nil,
		nil, nil, nil,
		nil, r.NFM); err != nil {

		return
	}

	// register it
	r.set(&x)

	return
}

/*
MarshalDITStructureRule will attempt to perform a marshal operation upon the provided string definition under the assumption that it represents a DIT Structure Rule.  An error is returned in the event that a parsing failure occurs in any way. No *DITStructureRule instance is stored in such an event.

In the event a successful marshal, a new *DITStructureRule instance is catalogued within the Subschema.DSRM (DITStructureRulesManifest) struct field.

This method requires the complete population of the following Manifests before-hand:

  - NameFormsManifest (Subschema.NFM)
  - DITStructureRulesManifest (Subschema.DSRM, if any when superior DIT Structure Rules are referenced)
*/
func (r *Subschema) MarshalDITStructureRule(def string) (err error) {

	// prelaunch checks and balances
	if len(def) == 0 {
		return raise(invalidMarshal,
			"MarshalDITStructureRule: definition is zero-length")
	} else if r.DSRM == nil {
		return raise(invalidMarshal,
			"MarshalDITStructureRule: %T.DSRM (%T) is nil (uninitialized)",
			r, r.DSRM)
	} else if r.NFM == nil {
		return raise(invalidMarshal,
			"MarshalDITStructureRule: %T.NFM (%T) is nil (uninitialized)",
			r, r.NFM)
	}

	// marshal it
	var x DITStructureRule
	if err = Marshal(
		def, &x, nil,
		nil, nil, nil,
		nil, nil, nil,
		r.DSRM, r.NFM); err != nil {

		return
	}

	// register it
	r.set(&x)

	return
}

/*
MarshalDITContentRule will attempt to perform a marshal operation upon the provided string definition under the assumption that it represents a DIT Content Rule.  An error is returned in the event that a parsing failure occurs in any way. No *DITContentRule instance is stored in such an event.

In the event a successful marshal, a new *DITContentRule instance is catalogued within the Subschema.DCRM (DITContentRulesManifest) struct field.

Note that the OID assigned to the raw DIT Content Rule definition MUST correlate to the registered OID of a STRUCTURAL *ObjectClass instance.

This method requires the complete population of the following Manifests before-hand:

  - AttributeTypesManifest (Subschema.ATM)
  - ObjectClassesManifest (Subschema.OCM)
*/
func (r *Subschema) MarshalDITContentRule(def string) (err error) {

	// prelaunch checks and balances
	if len(def) == 0 {
		return raise(invalidMarshal,
			"MarshalDITContentRule: definition is zero-length")
	} else if r.DCRM == nil {
		return raise(invalidMarshal,
			"MarshalDITContentRule: %T.DCRM (%T) is nil (uninitialized)",
			r, r.DCRM)
	} else if r.ATM == nil {
		return raise(invalidMarshal,
			"MarshalDITContentRule: %T.ATM (%T) is nil (uninitialized)",
			r, r.ATM)
	} else if r.OCM == nil {
		return raise(invalidMarshal,
			"MarshalDITContentRule: %T.OCM (%T) is nil (uninitialized)",
			r, r.OCM)
	}

	// marshal it
	var x DITContentRule
	if err = Marshal(
		def, &x, r.ALM,
		r.ATM, r.OCM, nil,
		nil, nil, r.DCRM,
		nil, nil); err != nil {

		return
	}

	// register it
	r.set(&x)

	return
}

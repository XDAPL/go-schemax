package schemax

/*
methlab.go deals with labels that are assigned to methods and stored in lookup maps. For example, `USAGE` for Usage.

What, exactly, were you expecting?
*/

func (r Extensions) labelIsValid(label string) (valid bool) {
	if len(label) < 2 {
		return
	}
	valid = label[:2] == `X-`

	return
}

/*
Label returns the known label for the receiver, if one exists.
*/
func (r Boolean) Label() string {
	return `` // no visible label
}

/*
Label returns the known label for the receiver, if one exists.
*/
func (r Extensions) Label() string {
	return `` // no visible label
}

/*
Label returns the known label for the receiver, if one exists.
*/
func (r Kind) Label() string {
	return `` // no visible label
}

/*
Label returns the known label for the receiver, if one exists.
*/
func (r OID) Label() string {
	return `` // no visible label
}

/*
Label returns the known label for the receiver, if one exists.
*/
func (r RuleID) Label() string {
	return `` // no visible label
}

/*
Label returns the known label for the receiver, if one exists.
*/
func (r NameForm) Label() string {
	return `FORM`
}

/*
Label returns the known label for the receiver, if one exists.
*/
func (r Applies) Label() string {
	return `APPLIES`
}

/*
Label returns the known label for the receiver, if one exists.
*/
func (r Usage) Label() string {
	return `USAGE`
}

/*
Label returns the known label for the receiver, if one exists.
*/
func (r SuperiorObjectClasses) Label() string {
	return `SUP`
}

/*
Label returns the known label for the receiver, if one exists.
*/
func (r SuperiorAttributeType) Label() string {
	return `SUP`
}

/*
Label returns the known label for the receiver, if one exists.
*/
func (r Equality) Label() string {
	return `EQUALITY`
}

/*
Label returns the known label for the receiver, if one exists.
*/
func (r Ordering) Label() string {
	return `ORDERING`
}

/*
Label returns the known label for the receiver, if one exists.
*/
func (r Substring) Label() string {
	return `SUBSTR`
}

/*
Label returns the known label for the receiver, if one exists.
*/
func (r SuperiorDITStructureRules) Label() string {
	return `SUP`
}

/*
Label returns the known label for the receiver, if one exists.
*/
func (r StructuralObjectClass) Label() string {
	return `OC`
}

/*
Label returns the known label for the receiver, if one exists.
*/
func (r AuxiliaryObjectClasses) Label() string {
	return `AUX`
}

/*
Label returns the known label for the receiver, if one exists.
*/
func (r RequiredAttributeTypes) Label() string {
	return `MUST`
}

/*
Label returns the known label for the receiver, if one exists.
*/
func (r PermittedAttributeTypes) Label() string {
	return `MAY`
}

/*
Label returns the known label for the receiver, if one exists.
*/
func (r ProhibitedAttributeTypes) Label() string {
	return `NOT`
}

/*
Label returns the known label for the receiver, if one exists.
*/
func (r Name) Label() string {
	return `NAME`
}

/*
Label returns the known label for the receiver, if one exists.
*/
func (r Description) Label() string {
	return `DESC`
}

/*
Label returns the known label for the receiver, if one exists.
*/
func (r LDAPSyntax) Label() string {
	return `SYNTAX`
}

/*
labelMap returns a map[int]string structure associating field names with labels.
*/
func (r AttributeType) labelMap() map[int]string {
	return map[int]string{
		0:  `OID`,                           // (OID)
		1:  Name{}.Label(),                  // NAME
		2:  Description(``).Label(),         // DESC
		3:  SuperiorAttributeType{}.Label(), // SUP
		4:  Equality{}.Label(),              // EQUALITY
		5:  Ordering{}.Label(),              // ORDERING
		6:  Substring{}.Label(),             // SUBSTR|SUBSTRING
		7:  LDAPSyntax{}.Label(),            // SYNTAX
		8:  Usage(0).Label(),                // USAGE
		9:  `EXT`,                           // (EXT)
		10: `BOOLS`,                         // [OBSOLETE,SINGLE-VALUE,NO-USER-MODIFICATION,COLLECTIVE]
		11: `MUB`,                           // unsigned non-zero integer setting min. length
		12: `SEQ`,                           // [SEQUENCE] (unsigned sequence number, *Subschema only)
	}
}

/*
labelMap returns a map[int]string structure associating field names with labels.
*/
func (r ObjectClass) labelMap() map[int]string {
	return map[int]string{
		0: `OID`,                             // (OID)
		1: Name{}.Label(),                    // NAME
		2: Description(``).Label(),           // DESC
		3: SuperiorAttributeType{}.Label(),   // SUP
		4: `KIND`,                            // [ABSTRACT|STRUCTURAL|AUXILIARY]
		5: RequiredAttributeTypes{}.Label(),  // MUST
		6: PermittedAttributeTypes{}.Label(), // MAY
		7: `EXT`,                             // (EXT)
		8: `BOOLS`,                           // [OBSOLETE]
		9: `SEQ`,                             // [SEQUENCE]
	}
}

/*
labelMap returns a map[int]string structure associating field names with labels.
*/
func (r MatchingRule) labelMap() map[int]string {
	return map[int]string{
		0: `OID`,                   // (OID)
		1: Name{}.Label(),          // NAME
		2: Description(``).Label(), // DESC
		3: LDAPSyntax{}.Label(),    // SYNTAX
		4: `EXT`,                   // (EXT)
		5: `BOOLS`,                 // [OBSOLETE]
		6: `SEQ`,                   // [SEQUENCE]
	}
}

/*
labelMap returns a map[int]string structure associating field names with labels.
*/
func (r LDAPSyntax) labelMap() map[int]string {
	return map[int]string{
		0: `OID`,                   // (OID)
		1: Description(``).Label(), // DESC
		2: `EXT`,                   // (EXT)
		3: `BOOLS`,                 // [HUMAN-READABLE]
		4: `SEQ`,                   // [SEQUENCE]
	}
}

/*
labelMap returns a map[int]string structure associating field names with labels.
*/
func (r MatchingRuleUse) labelMap() map[int]string {
	return map[int]string{
		0: `OID`,                   // (OID)
		1: Name{}.Label(),          // NAME
		2: Description(``).Label(), // DESC
		3: Applies{}.Label(),       // APPLIES
		4: `EXT`,                   // (EXT)
		5: `BOOLS`,                 // [OBSOLETE]
		6: `SEQ`,                   // [SEQUENCE]
	}
}

/*
labelMap returns a map[int]string structure associating field names with labels.
*/
func (r DITContentRule) labelMap() map[int]string {
	return map[int]string{
		0: `OID`,                              // (OID)
		1: Name{}.Label(),                     // NAME
		2: Description(``).Label(),            // DESC
		3: AuxiliaryObjectClasses{}.Label(),   // AUX
		4: RequiredAttributeTypes{}.Label(),   // MUST
		5: PermittedAttributeTypes{}.Label(),  // MAY
		6: ProhibitedAttributeTypes{}.Label(), // NOT
		7: `EXT`,                              // (EXT)
		8: `BOOLS`,                            // [OBSOLETE]
		9: `SEQ`,                              // [SEQUENCE]
	}
}

/*
labelMap returns a map[int]string structure associating field names with labels.
*/
func (r DITStructureRule) labelMap() map[int]string {
	return map[int]string{
		0: `ID`,                                // (ID)
		1: Name{}.Label(),                      // NAME
		2: Description(``).Label(),             // DESC
		3: NameForm{}.Label(),                  // FORM
		4: SuperiorDITStructureRules{}.Label(), // SUP
		5: `EXT`,                               // (EXT)
		6: `BOOLS`,                             // [OBSOLETE]
		7: `SEQ`,                               // [SEQUENCE]
	}
}

/*
labelMap returns a map[int]string structure associating field names with labels.
*/
func (r NameForm) labelMap() map[int]string {
	return map[int]string{
		0: `OID`,                             // (OID)
		1: Name{}.Label(),                    // NAME
		2: Description(``).Label(),           // DESC
		3: StructuralObjectClass{}.Label(),   // OC
		4: RequiredAttributeTypes{}.Label(),  // MUST
		5: PermittedAttributeTypes{}.Label(), // MAY
		6: `EXT`,                             // (EXT)
		7: `BOOLS`,                           // [OBSOLETE]
		8: `SEQ`,                             // [SEQUENCE]
	}
}

/*
methMap returns a map[int]parseMeth structure containing field int -> parseMeth signature references.
*/
func (r AttributeType) methMap() (mm map[int]parseMeth) {
	return map[int]parseMeth{
		0:  parse_numericoid, // (OID)
		1:  parse_qdescrs,    // NAME
		2:  parse_qdstring,   // DESC
		3:  parse_numericoid, // SUP
		4:  parse_numericoid, // EQUALITY
		5:  parse_numericoid, // ORDERING
		6:  parse_numericoid, // SUBSTR
		7:  parse_numericoid, // SYNTAX
		8:  parse_usage,      // USAGE
		9:  parse_extensions, // (EXT)
		10: parse_boolean,    // [SINGLE-VALUE, COLLECTIVE, NO-USER-MODIFICATION, OBSOLETE]
		11: parse_mub,        // {SYNTAX "len"}
		12: parse_boolean,    // [SEQUENCE]
	}
}

func (r ObjectClass) methMap() (mm map[int]parseMeth) {
	return map[int]parseMeth{
		0: parse_numericoid, // (OID)
		1: parse_qdescrs,    // NAME
		2: parse_qdstring,   // DESC
		3: parse_oids_ids,   // SUP
		4: parse_kind,       // [STRUCTURAL, AUXILIARY, ABSTRACT]
		5: parse_oids_ids,   // MUST
		6: parse_oids_ids,   // MAY
		7: parse_extensions, // (EXT)
		8: parse_boolean,    // [OBSOLETE]
		9: parse_boolean,    // [SEQUENCE]
	}
}

func (r LDAPSyntax) methMap() (mm map[int]parseMeth) {
	return map[int]parseMeth{
		0: parse_numericoid, // (OID)
		1: parse_qdstring,   // DESC
		2: parse_extensions, // (EXT)
		3: parse_boolean,    // [HUMAN-READABLE]
		4: parse_boolean,    // [SEQUENCE]
	}
}

func (r MatchingRule) methMap() (mm map[int]parseMeth) {
	return map[int]parseMeth{
		0: parse_numericoid, // (OID)
		1: parse_qdescrs,    // NAME
		2: parse_qdstring,   // DESC
		3: parse_numericoid, // SYNTAX
		4: parse_extensions, // (EXT)
		5: parse_boolean,    // [OBSOLETE]
		6: parse_boolean,    // [SEQUENCE]
	}
}

func (r MatchingRuleUse) methMap() (mm map[int]parseMeth) {
	return map[int]parseMeth{
		0: parse_numericoid, // (OID)
		1: parse_qdescrs,    // NAME
		2: parse_qdstring,   // DESC
		3: parse_oids_ids,   // APPLIES
		4: parse_extensions, // (EXT)
		5: parse_boolean,    // [OBSOLETE]
		6: parse_boolean,    // [SEQUENCE]
	}
}

func (r DITContentRule) methMap() (mm map[int]parseMeth) {
	return map[int]parseMeth{
		0: parse_numericoid, // (OID)
		1: parse_qdescrs,    // NAME
		2: parse_qdstring,   // DESC
		3: parse_oids_ids,   // AUX
		4: parse_oids_ids,   // MUST
		5: parse_oids_ids,   // MAY
		6: parse_oids_ids,   // NOT
		7: parse_extensions, // (EXT)
		8: parse_boolean,    // [OBSOLETE]
		9: parse_boolean,    // [SEQUENCE]
	}
}

func (r DITStructureRule) methMap() (mm map[int]parseMeth) {
	return map[int]parseMeth{
		0: parse_ruleid,     // (ID)
		1: parse_qdescrs,    // NAME
		2: parse_qdstring,   // DESC
		3: parse_numericoid, // FORM
		4: parse_oids_ids,   // SUP
		5: parse_extensions, // (EXT)
		6: parse_boolean,    // [OBSOLETE]
		7: parse_boolean,    // [SEQUENCE]
	}
}

func (r NameForm) methMap() (mm map[int]parseMeth) {
	return map[int]parseMeth{
		0: parse_numericoid, // (OID)
		1: parse_qdescrs,    // NAME
		2: parse_qdstring,   // DESC
		3: parse_oids_ids,   // OC
		4: parse_oids_ids,   // MUST
		5: parse_oids_ids,   // MAY
		6: parse_extensions, // (EXT)
		7: parse_boolean,    // [OBSOLETE]
		8: parse_boolean,    // [SEQUENCE]
	}
}

package schemax

/*
methlab.go deals with labels that are assigned to methods and stored in lookup maps. For example, `USAGE` for Usage.

What, exactly, were you expecting?
*/

/*
parseMeth is the first class function bearing a signature shared by all fundamental parser methods.
*/
type parseMeth func(string) ([]string, string, bool)

/*
Label returns the known label for the receiver, if one exists.
*/
func (r atFlags) Label() string {
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
func (r ApplicableAttributeTypes) Label() string {
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
func (r AttributeTypes) Label() string {
	return `` // no visible label
}

/*
Label returns the known label for the receiver, if one exists.
*/
func (r ObjectClasses) Label() string {
	return `` // no visible label
}

/*
Label returns the known label for the receiver, if one exists.
*/
func (r LDAPSyntaxes) Label() string {
	return `` // no visible label
}

/*
Label returns the known label for the receiver, if one exists.
*/
func (r MatchingRules) Label() string {
	return `` // no visible label
}

/*
Label returns the known label for the receiver, if one exists.
*/
func (r MatchingRuleUses) Label() string {
	return `` // no visible label
}

/*
Label returns the known label for the receiver, if one exists.
*/
func (r NameForms) Label() string {
	return `` // no visible label
}

/*
Label returns the known label for the receiver, if one exists.
*/
func (r DITStructureRules) Label() string {
	return `` // no visible label
}

/*
Label returns the known label for the receiver, if one exists.
*/
func (r DITContentRules) Label() string {
	return `` // no visible label
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
		3:  `OBSOLETE`,                      // [OBSOLETE]
		4:  SuperiorAttributeType{}.Label(), // SUP
		5:  Equality{}.Label(),              // EQUALITY
		6:  Ordering{}.Label(),              // ORDERING
		7:  Substring{}.Label(),             // SUBSTR|SUBSTRING
		8:  LDAPSyntax{}.Label(),            // SYNTAX
		9:  Usage(0).Label(),                // USAGE
		10: `EXT`,                           // (EXT)
		11: `FLAGS`,                         // [OBSOLETE,SINGLE-VALUE,NO-USER-MODIFICATION,COLLECTIVE]
		12: `MUB`,                           // unsigned non-zero integer
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
		3: `OBSOLETE`,                        // [OBSOLETE]
		4: SuperiorAttributeType{}.Label(),   // SUP
		5: `KIND`,                            // [ABSTRACT|STRUCTURAL|AUXILIARY]
		6: RequiredAttributeTypes{}.Label(),  // MUST
		7: PermittedAttributeTypes{}.Label(), // MAY
		8: `EXT`,                             // (EXT)
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
		3: `OBSOLETE`,              // [OBSOLETE]
		4: LDAPSyntax{}.Label(),    // SYNTAX
		5: `EXT`,                   // (EXT)
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
	}
}

/*
labelMap returns a map[int]string structure associating field names with labels.
*/
func (r MatchingRuleUse) labelMap() map[int]string {
	return map[int]string{
		0: `OID`,                              // (OID)
		1: Name{}.Label(),                     // NAME
		2: Description(``).Label(),            // DESC
		3: `OBSOLETE`,                         // [OBSOLETE]
		4: ApplicableAttributeTypes{}.Label(), // APPLIES
		5: `EXT`,                              // (EXT)
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
		3: `OBSOLETE`,                         // [OBSOLETE]
		4: AuxiliaryObjectClasses{}.Label(),   // AUX
		5: RequiredAttributeTypes{}.Label(),   // MUST
		6: PermittedAttributeTypes{}.Label(),  // MAY
		7: ProhibitedAttributeTypes{}.Label(), // NOT
		8: `EXT`,                              // (EXT)
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
		3: `OBSOLETE`,                          // [OBSOLETE]
		4: NameForm{}.Label(),                  // FORM
		5: SuperiorDITStructureRules{}.Label(), // SUP
		6: `EXT`,                               // (EXT)
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
		3: `OBSOLETE`,                        // [OBSOLETE]
		4: StructuralObjectClass{}.Label(),   // OC
		5: RequiredAttributeTypes{}.Label(),  // MUST
		6: PermittedAttributeTypes{}.Label(), // MAY
		7: `EXT`,                             // (EXT)
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
		3:  parse_obsolete,   // [OBSOLETE]
		4:  parse_numericoid, // SUP
		5:  parse_numericoid, // EQUALITY
		6:  parse_numericoid, // ORDERING
		7:  parse_numericoid, // SUBSTR
		8:  parse_numericoid, // SYNTAX
		9:  parse_usage,      // USAGE
		10: parse_extensions, // (EXT)
		11: parse_atflags,    // [SINGLE-VALUE, COLLECTIVE, NO-USER-MODIFICATION]
		12: parse_mub,        // {SYNTAX "len"}
	}
}

func (r ObjectClass) methMap() (mm map[int]parseMeth) {
	return map[int]parseMeth{
		0: parse_numericoid, // (OID)
		1: parse_qdescrs,    // NAME
		2: parse_qdstring,   // DESC
		3: parse_obsolete,   // [OBSOLETE]
		4: parse_oids_ids,   // SUP
		5: parse_kind,       // [STRUCTURAL, AUXILIARY, ABSTRACT]
		6: parse_oids_ids,   // MUST
		7: parse_oids_ids,   // MAY
		8: parse_extensions, // (EXT)
	}
}

func (r LDAPSyntax) methMap() (mm map[int]parseMeth) {
	return map[int]parseMeth{
		0: parse_numericoid, // (OID)
		1: parse_qdstring,   // DESC
		2: parse_extensions, // (EXT)
	}
}

func (r MatchingRule) methMap() (mm map[int]parseMeth) {
	return map[int]parseMeth{
		0: parse_numericoid, // (OID)
		1: parse_qdescrs,    // NAME
		2: parse_qdstring,   // DESC
		3: parse_obsolete,   // [OBSOLETE]
		4: parse_numericoid, // SYNTAX
		5: parse_extensions, // (EXT)
	}
}

func (r MatchingRuleUse) methMap() (mm map[int]parseMeth) {
	return map[int]parseMeth{
		0: parse_numericoid, // (OID)
		1: parse_qdescrs,    // NAME
		2: parse_qdstring,   // DESC
		3: parse_obsolete,   // [OBSOLETE]
		4: parse_oids_ids,   // APPLIES
		5: parse_extensions, // (EXT)
	}
}

func (r DITContentRule) methMap() (mm map[int]parseMeth) {
	return map[int]parseMeth{
		0: parse_numericoid, // (OID)
		1: parse_qdescrs,    // NAME
		2: parse_qdstring,   // DESC
		3: parse_obsolete,   // [OBSOLETE]
		4: parse_oids_ids,   // AUX
		5: parse_oids_ids,   // MUST
		6: parse_oids_ids,   // MAY
		7: parse_oids_ids,   // NOT
		8: parse_extensions, // (EXT)
	}
}

func (r DITStructureRule) methMap() (mm map[int]parseMeth) {
	return map[int]parseMeth{
		0: parse_ruleid,     // (ID)
		1: parse_qdescrs,    // NAME
		2: parse_qdstring,   // DESC
		3: parse_obsolete,   // [OBSOLETE]
		4: parse_numericoid, // FORM
		5: parse_oids_ids,   // SUP
		6: parse_extensions, // (EXT)
	}
}

func (r NameForm) methMap() (mm map[int]parseMeth) {
	return map[int]parseMeth{
		0: parse_numericoid, // (OID)
		1: parse_qdescrs,    // NAME
		2: parse_qdstring,   // DESC
		3: parse_obsolete,   // [OBSOLETE]
		4: parse_oids_ids,   // OC
		5: parse_oids_ids,   // MUST
		6: parse_oids_ids,   // MAY
		7: parse_extensions, // (EXT)
	}
}

package schemax

var noContent error = newErr("No content was read")
var emptyDefinition error = newErr("Empty definition")
var unexpectedChar error = newErr("Unexpected character found")
var unexpectedType error = newErr("Unexpected type")
var invalidOID error = newErr("Invalid or zero-length dot-delimited ASN.1 object identifier")
var invalidDelim error = newErr("Bad delimiter found in definition")
var invalidLabel error = newErr("Bad label found in definition")
var invalidValue error = newErr("Bad value found in definition")
var unaddressableType error = newErr("Non-addressable or nil type")
var invalidMarshal error = newErr("Failed marshal attempt")
var invalidUnmarshal error = newErr("Failed unmarshal attempt")
var unknownDefinition error = newErr("Unknown or unregistered definition")
var unknownElement error = newErr("Unknown definition element")
var isZero error = newErr("Element or definition is zero (unpopulated or incomplete)")
var cannotOverwrite error = newErr("Overwrites forbidden")

var invalidFlag error = newErr("Invalid or inappropriate Boolean flag")
var invalidObjectClassKind error = newErr("Invalid kind of objectClass")
var invalidName error = newErr("Invalid name value(s)")
var invalidDescription error = newErr("Invalid description value")
var invalidNameForm error = newErr("No nameForm derived from definition")
var invalidDITContentRule error = newErr("No dITContentRule derived from definition")
var invalidDITStructureRule error = newErr("No dITStructureRule derived from definition")
var invalidSyntax error = newErr("No LDAPSyntax derived from definition")
var invalidObjectClass error = newErr("No objectClass derived from definition")
var invalidAttributeType error = newErr("No attributeType derived from definition")
var invalidMatchingRule error = newErr("No matchingRule derived from definition")
var invalidMatchingRuleUse error = newErr("No matchingRuleUses derived from definition")
var invalidUsage error = newErr("Invalid Usage value")

func raise(err error, text string, m ...any) error {
	if err == nil {
		if len(text) == 0 {
			return newErr("unspecified/unhandled exception")
		}
		return newErr(sprintf(text, m...))
	}

	if len(text) == 0 {
		return err
	}
	return newErr(sprintf(err.Error()+`: `+text, m...))
}

func raiseUnknownElement(funcname string, fobj any, fman any, val string, dest any) error {
	return raise(unknownElement, "%s: no such %T was found in %T for value '%s' (type: %T)",
		funcname, fobj, fman, val, dest)
}

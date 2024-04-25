package rfc4512

type ObjectClassDefinitions []ObjectClassDefinition
type ObjectClassDefinition string

var (
	AllObjectClasses ObjectClassDefinitions
)

var (
	Top              ObjectClassDefinition
	Alias            ObjectClassDefinition
	Subschema        ObjectClassDefinition
	ExtensibleObject ObjectClassDefinition
)

func (r ObjectClassDefinition) String() string {
	return string(r)
}

// Other object classes
func init() {
	Top = ObjectClassDefinition(`( 2.5.6.0 NAME 'top' ABSTRACT MUST objectClass X-ORIGIN 'RFC4512' )`)
	Alias = ObjectClassDefinition(`( 2.5.6.1 NAME 'alias' SUP top STRUCTURAL MUST aliasedObjectName X-ORIGIN 'RFC4512' )`)
	Subschema = ObjectClassDefinition(`( 2.5.20.1 NAME 'subschema' AUXILIARY MAY ( dITStructureRules $ nameForms $ dITContentRules $ objectClasses $ attributeTypes $ matchingRules $ matchingRuleUse ) X-ORIGIN 'RFC4512' )`)
	ExtensibleObject = ObjectClassDefinition(`( 1.3.6.1.4.1.1466.101.120.111 NAME 'extensibleObject' SUP top AUXILIARY X-ORIGIN 'RFC4512' )`)

	AllObjectClasses = ObjectClassDefinitions{
		Top,
		Alias,
		Subschema,
		ExtensibleObject,
	}
}

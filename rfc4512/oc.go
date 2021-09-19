package rfc4512

type RFC4512ObjectClasses []RFC4512ObjectClass
type RFC4512ObjectClass string

var (
	AllObjectClasses RFC4512ObjectClasses
)

var (
	Top              RFC4512ObjectClass
	Alias            RFC4512ObjectClass
	Subschema        RFC4512ObjectClass
	ExtensibleObject RFC4512ObjectClass
)

// Other object classes
func init() {
	Top = RFC4512ObjectClass(`( 2.5.6.0 NAME 'top' ABSTRACT MUST objectClass X-ORIGIN 'RFC4512' )`)
	Alias = RFC4512ObjectClass(`( 2.5.6.1 NAME 'alias' SUP top STRUCTURAL MUST aliasedObjectName X-ORIGIN 'RFC4512' )`)
	Subschema = RFC4512ObjectClass(`( 2.5.20.1 NAME 'subschema' AUXILIARY MAY ( dITStructureRules $ nameForms $ ditContentRules $ objectClasses $ attributeTypes $ matchingRules $ matchingRuleUse ) X-ORIGIN 'RFC4512' )`)
	ExtensibleObject = RFC4512ObjectClass(`( 1.3.6.1.4.1.1466.101.120.111 NAME 'extensibleObject' SUP top AUXILIARY X-ORIGIN 'RFC4512' )`)

	AllObjectClasses = RFC4512ObjectClasses{
		Top,
		Alias,
		Subschema,
		ExtensibleObject,
	}
}

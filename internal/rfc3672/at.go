package rfc3672

type AttributeTypeDefinitions []AttributeTypeDefinition
type AttributeTypeDefinition string

func (r AttributeTypeDefinitions) Len() int {
	return len(r)
}

var (
	AllAttributeTypes    AttributeTypeDefinitions
	AdministrativeRole   AttributeTypeDefinition
	SubtreeSpecification AttributeTypeDefinition
)

func (r AttributeTypeDefinition) String() string {
	return string(r)
}

func init() {
	AdministrativeRole = AttributeTypeDefinition(`( 2.5.18.5 NAME 'administrativeRole' EQUALITY objectIdentifierMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.38 USAGE directoryOperation X-ORIGIN 'RFC3672' )`)
	SubtreeSpecification = AttributeTypeDefinition(`( 2.5.18.6 NAME 'subtreeSpecification' SYNTAX 1.3.6.1.4.1.1466.115.121.1.45 SINGLE-VALUE USAGE directoryOperation X-ORIGIN 'RFC3672' )`)

	AllAttributeTypes = AttributeTypeDefinitions{
		AdministrativeRole,
		SubtreeSpecification,
	}
}

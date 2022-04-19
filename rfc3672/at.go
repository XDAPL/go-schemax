package rfc3672

type RFC3672AttributeTypes []RFC3672AttributeType
type RFC3672AttributeType string

var (
	AllAttributeTypes    RFC3672AttributeTypes
	AdministrativeRole   RFC3672AttributeType
	SubtreeSpecification RFC3672AttributeType
)

func init() {
	AdministrativeRole = RFC3672AttributeType(`( 2.5.18.5 NAME 'administrativeRole' EQUALITY objectIdentifierMatch USAGE directoryOperation SYNTAX 1.3.6.1.4.1.1466.115.121.1.38 X-ORIGIN 'RFC3672' )`)
	SubtreeSpecification = RFC3672AttributeType(`( 2.5.18.6 NAME 'subtreeSpecification' SINGLE-VALUE USAGE directoryOperation SYNTAX 1.3.6.1.4.1.1466.115.121.1.45 X-ORIGIN 'RFC3672' )`)

	AllAttributeTypes = RFC3672AttributeTypes{
		AdministrativeRole,
		SubtreeSpecification,
	}
}

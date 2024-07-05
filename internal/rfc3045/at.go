package rfc3045

type AttributeTypeDefinitions []AttributeTypeDefinition
type AttributeTypeDefinition string

func (r AttributeTypeDefinitions) Len() int {
	return len(r)
}

var (
	AllAttributeTypes AttributeTypeDefinitions
)

// Operational AttributeTypes
var (
	VendorName    AttributeTypeDefinition
	VendorVersion AttributeTypeDefinition
)

func (r AttributeTypeDefinition) String() string {
	return string(r)
}

func init() {
	VendorName = AttributeTypeDefinition(`( 1.3.6.1.1.4 NAME 'vendorName' EQUALITY 1.3.6.1.4.1.1466.109.114.1 SYNTAX 1.3.6.1.4.1.1466.115.121.1.15 SINGLE-VALUE NO-USER-MODIFICATION USAGE dSAOperation )`)
	VendorVersion = AttributeTypeDefinition(`( 1.3.6.1.1.5 NAME 'vendorVersion' EQUALITY 1.3.6.1.4.1.1466.109.114.1 SYNTAX 1.3.6.1.4.1.1466.115.121.1.15 SINGLE-VALUE NO-USER-MODIFICATION USAGE dSAOperation )`)

	AllAttributeTypes = []AttributeTypeDefinition{
		VendorName,
		VendorVersion,
	}
}

package rfc3045

type RFC3045AttributeTypes []RFC3045AttributeType
type RFC3045AttributeType string

var (
	AllAttributeTypes RFC3045AttributeTypes
)

// Operational AttributeTypes
var (
	VendorName    RFC3045AttributeType
	VendorVersion RFC3045AttributeType
)

func init() {
	VendorName = RFC3045AttributeType(`( 1.3.6.1.1.4 NAME 'vendorName' EQUALITY 1.3.6.1.4.1.1466.109.114.1 SYNTAX 1.3.6.1.4.1.1466.115.121.1.15 SINGLE-VALUE NO-USER-MODIFICATION USAGE dSAOperation )`)
	VendorVersion = RFC3045AttributeType(`( 1.3.6.1.1.5 NAME 'vendorVersion' EQUALITY 1.3.6.1.4.1.1466.109.114.1 SYNTAX 1.3.6.1.4.1.1466.115.121.1.15 SINGLE-VALUE NO-USER-MODIFICATION USAGE dSAOperation )`)

	AllAttributeTypes = []RFC3045AttributeType{
		VendorName,
		VendorVersion,
	}
}

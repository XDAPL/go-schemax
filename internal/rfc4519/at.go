package rfc4519

/*
AttributeTypeDefinitions is a slice type designed to store instances of AttributeTypeDefinition.
*/
type AttributeTypeDefinitions []AttributeTypeDefinition

/*
AttributeTypeDefinition is a string type designed to store a raw AttributeType definition.
*/
type AttributeTypeDefinition string

/*
AttributeTypes contains slices of all instances of AttributeTypeDefinition defined in this package.
*/
var AllAttributeTypes AttributeTypeDefinitions

var (
	BusinessCategory           AttributeTypeDefinition
	C                          AttributeTypeDefinition
	CN                         AttributeTypeDefinition
	DC                         AttributeTypeDefinition
	Description                AttributeTypeDefinition
	DestinationIndicator       AttributeTypeDefinition
	DistinguishedName          AttributeTypeDefinition
	DNQualifier                AttributeTypeDefinition
	EnhancedSearchGuide        AttributeTypeDefinition
	FacsimileTelephoneNumber   AttributeTypeDefinition
	GenerationQualifier        AttributeTypeDefinition
	GivenName                  AttributeTypeDefinition
	HouseIdentifier            AttributeTypeDefinition
	Initials                   AttributeTypeDefinition
	InternationalISDNNumber    AttributeTypeDefinition
	L                          AttributeTypeDefinition
	Member                     AttributeTypeDefinition
	Name                       AttributeTypeDefinition
	O                          AttributeTypeDefinition
	OU                         AttributeTypeDefinition
	Owner                      AttributeTypeDefinition
	PhysicalDeliveryOfficeName AttributeTypeDefinition
	PostalAddress              AttributeTypeDefinition
	PostalCode                 AttributeTypeDefinition
	PostOfficeBox              AttributeTypeDefinition
	PreferredDeliveryMethod    AttributeTypeDefinition
	RegisteredAddress          AttributeTypeDefinition
	RoleOccupant               AttributeTypeDefinition
	SearchGuide                AttributeTypeDefinition
	SeeAlso                    AttributeTypeDefinition
	SerialNumber               AttributeTypeDefinition
	SN                         AttributeTypeDefinition
	ST                         AttributeTypeDefinition
	Street                     AttributeTypeDefinition
	TelephoneNumber            AttributeTypeDefinition
	TeletexTerminalIdentifier  AttributeTypeDefinition
	TelexNumber                AttributeTypeDefinition
	Title                      AttributeTypeDefinition
	UID                        AttributeTypeDefinition
	UniqueMember               AttributeTypeDefinition
	UserPassword               AttributeTypeDefinition
	X121Address                AttributeTypeDefinition
	X500UniqueIdentifier       AttributeTypeDefinition
)

func (r AttributeTypeDefinition) String() string {
	return string(r)
}

// User AttributeTypes
func init() {
	BusinessCategory = AttributeTypeDefinition(`( 2.5.4.15 NAME 'businessCategory' EQUALITY caseIgnoreMatch SUBSTR caseIgnoreSubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.15 X-ORIGIN 'RFC4519' )`)
	C = AttributeTypeDefinition(`( 2.5.4.6 NAME 'c' SUP name SYNTAX 1.3.6.1.4.1.1466.115.121.1.11 SINGLE-VALUE X-ORIGIN 'RFC4519' )`)
	CN = AttributeTypeDefinition(`( 2.5.4.3 NAME ( 'cn' 'commonName' ) DESC 'RFC4519: common name(s) for which the entity is known by' SUP name X-ORIGIN 'RFC4519' )`)
	DC = AttributeTypeDefinition(`( 0.9.2342.19200300.100.1.25 NAME ( 'dc' 'domainComponent' ) EQUALITY caseIgnoreIA5Match SUBSTR caseIgnoreIA5SubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.26 SINGLE-VALUE X-ORIGIN 'RFC4519' )`)
	Description = AttributeTypeDefinition(`( 2.5.4.13 NAME 'description' EQUALITY caseIgnoreMatch SUBSTR caseIgnoreSubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.15 X-ORIGIN 'RFC4519' )`)
	DestinationIndicator = AttributeTypeDefinition(`( 2.5.4.27 NAME 'destinationIndicator' EQUALITY caseIgnoreMatch SUBSTR caseIgnoreSubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.44 X-ORIGIN 'RFC4519' )`)
	DistinguishedName = AttributeTypeDefinition(`( 2.5.4.49 NAME 'distinguishedName' EQUALITY distinguishedNameMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.12 X-ORIGIN 'RFC4519' )`)
	DNQualifier = AttributeTypeDefinition(`( 2.5.4.46 NAME 'dnQualifier' EQUALITY caseIgnoreMatch ORDERING caseIgnoreOrderingMatch SUBSTR caseIgnoreSubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.44 X-ORIGIN 'RFC4519' )`)
	EnhancedSearchGuide = AttributeTypeDefinition(`( 2.5.4.47 NAME 'enhancedSearchGuide' SYNTAX 1.3.6.1.4.1.1466.115.121.1.21 X-ORIGIN 'RFC4519' )`)
	FacsimileTelephoneNumber = AttributeTypeDefinition(`( 2.5.4.23 NAME 'facsimileTelephoneNumber' SYNTAX 1.3.6.1.4.1.1466.115.121.1.22 X-ORIGIN 'RFC4519' )`)
	GenerationQualifier = AttributeTypeDefinition(`( 2.5.4.44 NAME 'generationQualifier' SUP name X-ORIGIN 'RFC4519' )`)
	GivenName = AttributeTypeDefinition(`( 2.5.4.42 NAME 'givenName' SUP name X-ORIGIN 'RFC4519' )`)
	HouseIdentifier = AttributeTypeDefinition(`( 2.5.4.51 NAME 'houseIdentifier' EQUALITY caseIgnoreMatch SUBSTR caseIgnoreSubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.15 X-ORIGIN 'RFC4519' )`)
	Initials = AttributeTypeDefinition(`( 2.5.4.43 NAME 'initials' SUP name X-ORIGIN 'RFC4519' )`)
	InternationalISDNNumber = AttributeTypeDefinition(`( 2.5.4.25 NAME 'internationalISDNNumber' EQUALITY numericStringMatch SUBSTR numericStringSubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.36 X-ORIGIN 'RFC4519' )`)
	L = AttributeTypeDefinition(`( 2.5.4.7 NAME ( 'l' 'localityName' ) SUP name X-ORIGIN 'RFC4519' )`)
	Member = AttributeTypeDefinition(`( 2.5.4.31 NAME 'member' SUP distinguishedName X-ORIGIN 'RFC4519' )`)
	Name = AttributeTypeDefinition(`( 2.5.4.41 NAME 'name' EQUALITY caseIgnoreMatch SUBSTR caseIgnoreSubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.15 X-ORIGIN 'RFC4519' )`)
	O = AttributeTypeDefinition(`( 2.5.4.10 NAME ( 'o' 'organizationName' ) SUP name X-ORIGIN 'RFC4519' )`)
	OU = AttributeTypeDefinition(`( 2.5.4.11 NAME ( 'ou' 'organizationalUnitName' ) SUP name X-ORIGIN 'RFC4519' )`)
	Owner = AttributeTypeDefinition(`( 2.5.4.32 NAME 'owner' SUP distinguishedName X-ORIGIN 'RFC4519' )`)
	PhysicalDeliveryOfficeName = AttributeTypeDefinition(`( 2.5.4.19 NAME 'physicalDeliveryOfficeName' EQUALITY caseIgnoreMatch SUBSTR caseIgnoreSubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.15 X-ORIGIN 'RFC4519' )`)
	PostalAddress = AttributeTypeDefinition(`( 2.5.4.16 NAME 'postalAddress' EQUALITY caseIgnoreListMatch SUBSTR caseIgnoreListSubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.41 X-ORIGIN 'RFC4519' )`)
	PostalCode = AttributeTypeDefinition(`( 2.5.4.17 NAME 'postalCode' EQUALITY caseIgnoreMatch SUBSTR caseIgnoreSubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.15 X-ORIGIN 'RFC4519' )`)
	PostOfficeBox = AttributeTypeDefinition(`( 2.5.4.18 NAME 'postOfficeBox' EQUALITY caseIgnoreMatch SUBSTR caseIgnoreSubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.15 X-ORIGIN 'RFC4519' )`)
	PreferredDeliveryMethod = AttributeTypeDefinition(`( 2.5.4.28 NAME 'preferredDeliveryMethod' SYNTAX 1.3.6.1.4.1.1466.115.121.1.14 SINGLE-VALUE X-ORIGIN 'RFC4519' )`)
	RegisteredAddress = AttributeTypeDefinition(`( 2.5.4.26 NAME 'registeredAddress' SUP postalAddress SYNTAX 1.3.6.1.4.1.1466.115.121.1.41 X-ORIGIN 'RFC4519' )`)
	RoleOccupant = AttributeTypeDefinition(`( 2.5.4.33 NAME 'roleOccupant' SUP distinguishedName X-ORIGIN 'RFC4519' )`)
	SearchGuide = AttributeTypeDefinition(`( 2.5.4.14 NAME 'searchGuide' SYNTAX 1.3.6.1.4.1.1466.115.121.1.25 X-ORIGIN 'RFC4519' )`)
	SeeAlso = AttributeTypeDefinition(`( 2.5.4.34 NAME 'seeAlso' SUP distinguishedName X-ORIGIN 'RFC4519' )`)
	SerialNumber = AttributeTypeDefinition(`( 2.5.4.5 NAME 'serialNumber' EQUALITY caseIgnoreMatch SUBSTR caseIgnoreSubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.44 X-ORIGIN 'RFC4519' )`)
	SN = AttributeTypeDefinition(`( 2.5.4.4 NAME ( 'sn' 'surname' ) SUP name X-ORIGIN 'RFC4519' )`)
	ST = AttributeTypeDefinition(`( 2.5.4.8 NAME ( 'st' 'stateOrProvinceName' ) SUP name X-ORIGIN 'RFC4519' )`)
	Street = AttributeTypeDefinition(`( 2.5.4.9 NAME ( 'street' 'streetAddress' ) EQUALITY caseIgnoreMatch SUBSTR caseIgnoreSubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.15 X-ORIGIN 'RFC4519' )`)
	TelephoneNumber = AttributeTypeDefinition(`( 2.5.4.20 NAME 'telephoneNumber' EQUALITY telephoneNumberMatch SUBSTR telephoneNumberSubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.50 X-ORIGIN 'RFC4519' )`)
	TeletexTerminalIdentifier = AttributeTypeDefinition(`( 2.5.4.22 NAME 'teletexTerminalIdentifier' SYNTAX 1.3.6.1.4.1.1466.115.121.1.51 X-ORIGIN 'RFC4519' )`)
	TelexNumber = AttributeTypeDefinition(`( 2.5.4.21 NAME 'telexNumber' SYNTAX 1.3.6.1.4.1.1466.115.121.1.52 X-ORIGIN 'RFC4519' )`)
	Title = AttributeTypeDefinition(`( 2.5.4.12 NAME 'title' SUP name X-ORIGIN 'RFC4519' )`)
	UID = AttributeTypeDefinition(`( 0.9.2342.19200300.100.1.1 NAME ( 'uid' 'userid' ) DESC 'RFC4519: user identifier' EQUALITY caseIgnoreMatch SUBSTR caseIgnoreSubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.15{256} X-ORIGIN 'RFC4519' )`)
	UniqueMember = AttributeTypeDefinition(`( 2.5.4.50 NAME 'uniqueMember' EQUALITY uniqueMemberMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.34 X-ORIGIN 'RFC4519' )`)
	UserPassword = AttributeTypeDefinition(`( 2.5.4.35 NAME 'userPassword' EQUALITY octetStringMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.40 X-ORIGIN 'RFC4519' )`)
	X121Address = AttributeTypeDefinition(`( 2.5.4.24 NAME 'x121Address' EQUALITY numericStringMatch SUBSTR numericStringSubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.36 X-ORIGIN 'RFC4519' )`)
	X500UniqueIdentifier = AttributeTypeDefinition(`( 2.5.4.45 NAME 'x500UniqueIdentifier' EQUALITY bitStringMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.6 X-ORIGIN 'RFC4519' )`)

	AllAttributeTypes = AttributeTypeDefinitions{
		Name,
		DistinguishedName,
		BusinessCategory,
		C,
		CN,
		DC,
		Description,
		DestinationIndicator,
		DNQualifier,
		EnhancedSearchGuide,
		FacsimileTelephoneNumber,
		GenerationQualifier,
		GivenName,
		HouseIdentifier,
		Initials,
		InternationalISDNNumber,
		L,
		Member,
		O,
		OU,
		Owner,
		PhysicalDeliveryOfficeName,
		PostalAddress,
		PostalCode,
		PostOfficeBox,
		PreferredDeliveryMethod,
		RegisteredAddress,
		RoleOccupant,
		SearchGuide,
		SeeAlso,
		SerialNumber,
		SN,
		ST,
		Street,
		TelephoneNumber,
		TeletexTerminalIdentifier,
		TelexNumber,
		Title,
		UID,
		UniqueMember,
		UserPassword,
		X121Address,
		X500UniqueIdentifier,
	}
}

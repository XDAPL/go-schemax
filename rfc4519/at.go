package rfc4519

/*
RFC4519AttributeTypes is a slice type designed to store instances of RFC4519AttributeType.
*/
type RFC4519AttributeTypes []RFC4519AttributeType

/*
RFC4519AttributeType is a string type designed to store a raw AttributeType definition.
*/
type RFC4519AttributeType string

/*
AttributeTypes contains slices of all instances of RFC4519AttributeType defined in this package.
*/
var AllAttributeTypes RFC4519AttributeTypes

var (
	BusinessCategory           RFC4519AttributeType
	C                          RFC4519AttributeType
	CN                         RFC4519AttributeType
	DC                         RFC4519AttributeType
	Description                RFC4519AttributeType
	DestinationIndicator       RFC4519AttributeType
	DistinguishedName          RFC4519AttributeType
	DNQualifier                RFC4519AttributeType
	EnhancedSearchGuide        RFC4519AttributeType
	FacsimileTelephoneNumber   RFC4519AttributeType
	GenerationQualifier        RFC4519AttributeType
	GivenName                  RFC4519AttributeType
	HouseIdentifier            RFC4519AttributeType
	Initials                   RFC4519AttributeType
	InternationalISDNNumber    RFC4519AttributeType
	L                          RFC4519AttributeType
	Member                     RFC4519AttributeType
	Name                       RFC4519AttributeType
	O                          RFC4519AttributeType
	OU                         RFC4519AttributeType
	Owner                      RFC4519AttributeType
	PhysicalDeliveryOfficeName RFC4519AttributeType
	PostalAddress              RFC4519AttributeType
	PostalCode                 RFC4519AttributeType
	PostOfficeBox              RFC4519AttributeType
	PreferredDeliveryMethod    RFC4519AttributeType
	RegisteredAddress          RFC4519AttributeType
	RoleOccupant               RFC4519AttributeType
	SearchGuide                RFC4519AttributeType
	SeeAlso                    RFC4519AttributeType
	SerialNumber               RFC4519AttributeType
	SN                         RFC4519AttributeType
	ST                         RFC4519AttributeType
	Street                     RFC4519AttributeType
	TelephoneNumber            RFC4519AttributeType
	TeletexTerminalIdentifier  RFC4519AttributeType
	TelexNumber                RFC4519AttributeType
	Title                      RFC4519AttributeType
	UID                        RFC4519AttributeType
	UniqueMember               RFC4519AttributeType
	UserPassword               RFC4519AttributeType
	X121Address                RFC4519AttributeType
	X500UniqueIdentifier       RFC4519AttributeType
)

// User AttributeTypes
func init() {
	BusinessCategory = RFC4519AttributeType(`( 2.5.4.15 NAME 'businessCategory' EQUALITY caseIgnoreMatch SUBSTR caseIgnoreSubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.15 X-ORIGIN 'RFC4519' )`)
	C = RFC4519AttributeType(`( 2.5.4.6 NAME 'c' SUP name SYNTAX 1.3.6.1.4.1.1466.115.121.1.11 SINGLE-VALUE X-ORIGIN 'RFC4519' )`)
	CN = RFC4519AttributeType(`( 2.5.4.3 NAME ( 'cn' 'commonName' ) DESC 'RFC4519: common name(s) for which the entity is known by' SUP name X-ORIGIN 'RFC4519' )`)
	DC = RFC4519AttributeType(`( 0.9.2342.19200300.100.1.25 NAME 'dc' EQUALITY caseIgnoreIA5Match SUBSTR caseIgnoreIA5SubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.26 SINGLE-VALUE X-ORIGIN 'RFC4519' )`)
	Description = RFC4519AttributeType(`( 2.5.4.13 NAME 'description' EQUALITY caseIgnoreMatch SUBSTR caseIgnoreSubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.15 X-ORIGIN 'RFC4519' )`)
	DestinationIndicator = RFC4519AttributeType(`( 2.5.4.27 NAME 'destinationIndicator' EQUALITY caseIgnoreMatch SUBSTR caseIgnoreSubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.44 X-ORIGIN 'RFC4519' )`)
	DistinguishedName = RFC4519AttributeType(`( 2.5.4.49 NAME 'distinguishedName' EQUALITY distinguishedNameMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.12 X-ORIGIN 'RFC4519' )`)
	DNQualifier = RFC4519AttributeType(`( 2.5.4.46 NAME 'dnQualifier' EQUALITY caseIgnoreMatch ORDERING caseIgnoreOrderingMatch SUBSTR caseIgnoreSubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.44 X-ORIGIN 'RFC4519' )`)
	EnhancedSearchGuide = RFC4519AttributeType(`( 2.5.4.47 NAME 'enhancedSearchGuide' SYNTAX 1.3.6.1.4.1.1466.115.121.1.21 X-ORIGIN 'RFC4519' )`)
	FacsimileTelephoneNumber = RFC4519AttributeType(`( 2.5.4.23 NAME 'facsimileTelephoneNumber' SYNTAX 1.3.6.1.4.1.1466.115.121.1.22 X-ORIGIN 'RFC4519' )`)
	GenerationQualifier = RFC4519AttributeType(`( 2.5.4.44 NAME 'generationQualifier' SUP name X-ORIGIN 'RFC4519' )`)
	GivenName = RFC4519AttributeType(`( 2.5.4.42 NAME 'givenName' SUP name X-ORIGIN 'RFC4519' )`)
	HouseIdentifier = RFC4519AttributeType(`( 2.5.4.51 NAME 'houseIdentifier' EQUALITY caseIgnoreMatch SUBSTR caseIgnoreSubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.15 X-ORIGIN 'RFC4519' )`)
	Initials = RFC4519AttributeType(`( 2.5.4.43 NAME 'initials' SUP name X-ORIGIN 'RFC4519' )`)
	InternationalISDNNumber = RFC4519AttributeType(`( 2.5.4.25 NAME 'internationalISDNNumber' EQUALITY numericStringMatch SUBSTR numericStringSubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.36 X-ORIGIN 'RFC4519' )`)
	L = RFC4519AttributeType(`( 2.5.4.7 NAME 'l' SUP name X-ORIGIN 'RFC4519' )`)
	Member = RFC4519AttributeType(`( 2.5.4.31 NAME 'member' SUP distinguishedName X-ORIGIN 'RFC4519' )`)
	Name = RFC4519AttributeType(`( 2.5.4.41 NAME 'name' EQUALITY caseIgnoreMatch SUBSTR caseIgnoreSubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.15 X-ORIGIN 'RFC4519' )`)
	O = RFC4519AttributeType(`( 2.5.4.10 NAME 'o' SUP name X-ORIGIN 'RFC4519' )`)
	OU = RFC4519AttributeType(`( 2.5.4.11 NAME 'ou' SUP name X-ORIGIN 'RFC4519' )`)
	Owner = RFC4519AttributeType(`( 2.5.4.32 NAME 'owner' SUP distinguishedName X-ORIGIN 'RFC4519' )`)
	PhysicalDeliveryOfficeName = RFC4519AttributeType(`( 2.5.4.19 NAME 'physicalDeliveryOfficeName' EQUALITY caseIgnoreMatch SUBSTR caseIgnoreSubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.15 X-ORIGIN 'RFC4519' )`)
	PostalAddress = RFC4519AttributeType(`( 2.5.4.16 NAME 'postalAddress' EQUALITY caseIgnoreListMatch SUBSTR caseIgnoreListSubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.41 X-ORIGIN 'RFC4519' )`)
	PostalCode = RFC4519AttributeType(`( 2.5.4.17 NAME 'postalCode' EQUALITY caseIgnoreMatch SUBSTR caseIgnoreSubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.15 X-ORIGIN 'RFC4519' )`)
	PostOfficeBox = RFC4519AttributeType(`( 2.5.4.18 NAME 'postOfficeBox' EQUALITY caseIgnoreMatch SUBSTR caseIgnoreSubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.15 X-ORIGIN 'RFC4519' )`)
	PreferredDeliveryMethod = RFC4519AttributeType(`( 2.5.4.28 NAME 'preferredDeliveryMethod' SYNTAX 1.3.6.1.4.1.1466.115.121.1.14 SINGLE-VALUE X-ORIGIN 'RFC4519' )`)
	RegisteredAddress = RFC4519AttributeType(`( 2.5.4.26 NAME 'registeredAddress' SUP postalAddress SYNTAX 1.3.6.1.4.1.1466.115.121.1.41 X-ORIGIN 'RFC4519' )`)
	RoleOccupant = RFC4519AttributeType(`( 2.5.4.33 NAME 'roleOccupant' SUP distinguishedName X-ORIGIN 'RFC4519' )`)
	SearchGuide = RFC4519AttributeType(`( 2.5.4.14 NAME 'searchGuide' SYNTAX 1.3.6.1.4.1.1466.115.121.1.25 X-ORIGIN 'RFC4519' )`)
	SeeAlso = RFC4519AttributeType(`( 2.5.4.34 NAME 'seeAlso' SUP distinguishedName X-ORIGIN 'RFC4519' )`)
	SerialNumber = RFC4519AttributeType(`( 2.5.4.5 NAME 'serialNumber' EQUALITY caseIgnoreMatch SUBSTR caseIgnoreSubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.44 X-ORIGIN 'RFC4519' )`)
	SN = RFC4519AttributeType(`( 2.5.4.4 NAME 'sn' SUP name X-ORIGIN 'RFC4519' )`)
	ST = RFC4519AttributeType(`( 2.5.4.8 NAME 'st' SUP name X-ORIGIN 'RFC4519' )`)
	Street = RFC4519AttributeType(`( 2.5.4.9 NAME 'street' EQUALITY caseIgnoreMatch SUBSTR caseIgnoreSubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.15 X-ORIGIN 'RFC4519' )`)
	TelephoneNumber = RFC4519AttributeType(`( 2.5.4.20 NAME 'telephoneNumber' EQUALITY telephoneNumberMatch SUBSTR telephoneNumberSubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.50 X-ORIGIN 'RFC4519' )`)
	TeletexTerminalIdentifier = RFC4519AttributeType(`( 2.5.4.22 NAME 'teletexTerminalIdentifier' SYNTAX 1.3.6.1.4.1.1466.115.121.1.51 X-ORIGIN 'RFC4519' )`)
	TelexNumber = RFC4519AttributeType(`( 2.5.4.21 NAME 'telexNumber' SYNTAX 1.3.6.1.4.1.1466.115.121.1.52 X-ORIGIN 'RFC4519' )`)
	Title = RFC4519AttributeType(`( 2.5.4.12 NAME 'title' SUP name X-ORIGIN 'RFC4519' )`)
	UID = RFC4519AttributeType(`( 0.9.2342.19200300.100.1.1 NAME ( 'uid' 'userid' ) DESC 'RFC4519: user identifier' EQUALITY caseIgnoreMatch SUBSTR caseIgnoreSubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.15{256} X-ORIGIN 'RFC4519' )`)
	UniqueMember = RFC4519AttributeType(`( 2.5.4.50 NAME 'uniqueMember' EQUALITY uniqueMemberMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.34 X-ORIGIN 'RFC4519' )`)
	UserPassword = RFC4519AttributeType(`( 2.5.4.35 NAME 'userPassword' EQUALITY octetStringMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.40 X-ORIGIN 'RFC4519' )`)
	X121Address = RFC4519AttributeType(`( 2.5.4.24 NAME 'x121Address' EQUALITY numericStringMatch SUBSTR numericStringSubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.36 X-ORIGIN 'RFC4519' )`)
	X500UniqueIdentifier = RFC4519AttributeType(`( 2.5.4.45 NAME 'x500UniqueIdentifier' EQUALITY bitStringMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.6 X-ORIGIN 'RFC4519' )`)

	AllAttributeTypes = RFC4519AttributeTypes{
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

package rfc4524

type RFC4524AttributeTypes []RFC4524AttributeType
type RFC4524AttributeType string

var (
	AllAttributeTypes RFC4524AttributeTypes
)

var (
	AssociatedDomain     RFC4524AttributeType
	AssociatedName       RFC4524AttributeType
	BuildingName         RFC4524AttributeType
	CO                   RFC4524AttributeType
	DocumentAuthor       RFC4524AttributeType
	DocumentIdentifier   RFC4524AttributeType
	DocumentLocation     RFC4524AttributeType
	DocumentPublisher    RFC4524AttributeType
	DocumentTitle        RFC4524AttributeType
	DocumentVersion      RFC4524AttributeType
	Drink                RFC4524AttributeType
	HomePhone            RFC4524AttributeType
	HomePostalAddress    RFC4524AttributeType
	Host                 RFC4524AttributeType
	Info                 RFC4524AttributeType
	Mail                 RFC4524AttributeType
	Manager              RFC4524AttributeType
	Mobile               RFC4524AttributeType
	OrganizationalStatus RFC4524AttributeType
	Pager                RFC4524AttributeType
	PersonalTitle        RFC4524AttributeType
	RoomNumber           RFC4524AttributeType
	Secretary            RFC4524AttributeType
	UniqueIdentifier     RFC4524AttributeType
	UserClass            RFC4524AttributeType
)

func init() {
	AssociatedDomain = RFC4524AttributeType(`( 0.9.2342.19200300.100.1.37 NAME 'associatedDomain' EQUALITY caseIgnoreIA5Match SUBSTR caseIgnoreIA5SubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.26 X-ORIGIN 'RFC4524' )`)
	AssociatedName = RFC4524AttributeType(`( 0.9.2342.19200300.100.1.38 NAME 'associatedName' EQUALITY distinguishedNameMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.12 X-ORIGIN 'RFC4524' )`)
	BuildingName = RFC4524AttributeType(`( 0.9.2342.19200300.100.1.48 NAME 'buildingName' EQUALITY caseIgnoreMatch SUBSTR caseIgnoreSubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.15{256} X-ORIGIN 'RFC4524' )`)
	CO = RFC4524AttributeType(`( 0.9.2342.19200300.100.1.43 NAME ( 'co' 'friendlyCountryName' ) EQUALITY caseIgnoreMatch SUBSTR caseIgnoreSubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.15 X-ORIGIN 'RFC4524' )`)
	DocumentAuthor = RFC4524AttributeType(`( 0.9.2342.19200300.100.1.14 NAME 'documentAuthor' EQUALITY distinguishedNameMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.12 X-ORIGIN 'RFC4524' )`)
	DocumentIdentifier = RFC4524AttributeType(`( 0.9.2342.19200300.100.1.11 NAME 'documentIdentifier' EQUALITY caseIgnoreMatch SUBSTR caseIgnoreSubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.15{256} X-ORIGIN 'RFC4524' )`)
	DocumentLocation = RFC4524AttributeType(`( 0.9.2342.19200300.100.1.15 NAME 'documentLocation' EQUALITY caseIgnoreMatch SUBSTR caseIgnoreSubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.15{256} X-ORIGIN 'RFC4524' )`)
	DocumentPublisher = RFC4524AttributeType(`( 0.9.2342.19200300.100.1.56 NAME 'documentPublisher' EQUALITY caseIgnoreMatch SUBSTR caseIgnoreSubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.15 X-ORIGIN 'RFC4524' )`)
	DocumentTitle = RFC4524AttributeType(`( 0.9.2342.19200300.100.1.12 NAME 'documentTitle' EQUALITY caseIgnoreMatch SUBSTR caseIgnoreSubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.15{256} X-ORIGIN 'RFC4524' )`)
	DocumentVersion = RFC4524AttributeType(`( 0.9.2342.19200300.100.1.13 NAME 'documentVersion' EQUALITY caseIgnoreMatch SUBSTR caseIgnoreSubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.15{256} X-ORIGIN 'RFC4524' )`)
	Drink = RFC4524AttributeType(`( 0.9.2342.19200300.100.1.5 NAME ( 'drink' 'favouriteDrink' ) EQUALITY caseIgnoreMatch SUBSTR caseIgnoreSubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.15{256} X-ORIGIN 'RFC4524' )`)
	HomePhone = RFC4524AttributeType(`( 0.9.2342.19200300.100.1.20 NAME ( 'homePhone' 'homeTelephoneNumber' ) EQUALITY telephoneNumberMatch SUBSTR telephoneNumberSubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.50 X-ORIGIN 'RFC4524' )`)
	HomePostalAddress = RFC4524AttributeType(`( 0.9.2342.19200300.100.1.39 NAME 'homePostalAddress' EQUALITY caseIgnoreListMatch SUBSTR caseIgnoreListSubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.41 X-ORIGIN 'RFC4524' )`)
	Host = RFC4524AttributeType(`( 0.9.2342.19200300.100.1.9 NAME 'host' EQUALITY caseIgnoreMatch SUBSTR caseIgnoreSubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.15{256} X-ORIGIN 'RFC4524' )`)
	Info = RFC4524AttributeType(`( 0.9.2342.19200300.100.1.4 NAME 'info' EQUALITY caseIgnoreMatch SUBSTR caseIgnoreSubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.15{2048} X-ORIGIN 'RFC4524' )`)
	Mail = RFC4524AttributeType(`( 0.9.2342.19200300.100.1.3 NAME ( 'mail' 'rfc822Mailbox' ) EQUALITY caseIgnoreIA5Match SUBSTR caseIgnoreIA5SubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.26{256} X-ORIGIN 'RFC4524' )`)
	Manager = RFC4524AttributeType(`( 0.9.2342.19200300.100.1.10 NAME 'manager' EQUALITY distinguishedNameMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.12 X-ORIGIN 'RFC4524' )`)
	Mobile = RFC4524AttributeType(`( 0.9.2342.19200300.100.1.41 NAME ( 'mobile' 'mobileTelephoneNumber' ) EQUALITY telephoneNumberMatch SUBSTR telephoneNumberSubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.50 X-ORIGIN 'RFC4524' )`)
	OrganizationalStatus = RFC4524AttributeType(`( 0.9.2342.19200300.100.1.45 NAME 'organizationalStatus' EQUALITY caseIgnoreMatch SUBSTR caseIgnoreSubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.15{256} X-ORIGIN 'RFC4524' )`)
	Pager = RFC4524AttributeType(`( 0.9.2342.19200300.100.1.42 NAME ( 'pager' 'pagerTelephoneNumber' ) EQUALITY telephoneNumberMatch SUBSTR telephoneNumberSubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.50 X-ORIGIN 'RFC4524' )`)
	PersonalTitle = RFC4524AttributeType(`( 0.9.2342.19200300.100.1.40 NAME 'personalTitle' EQUALITY caseIgnoreMatch SUBSTR caseIgnoreSubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.15{256} X-ORIGIN 'RFC4524' )`)
	RoomNumber = RFC4524AttributeType(`( 0.9.2342.19200300.100.1.6 NAME 'roomNumber' EQUALITY caseIgnoreMatch SUBSTR caseIgnoreSubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.15{256} X-ORIGIN 'RFC4524' )`)
	Secretary = RFC4524AttributeType(`( 0.9.2342.19200300.100.1.21 NAME 'secretary' EQUALITY distinguishedNameMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.12 X-ORIGIN 'RFC4524' )`)
	UniqueIdentifier = RFC4524AttributeType(`( 0.9.2342.19200300.100.1.44 NAME 'uniqueIdentifier' EQUALITY caseIgnoreMatch SUBSTR caseIgnoreSubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.15{256} X-ORIGIN 'RFC4524' )`)
	UserClass = RFC4524AttributeType(`( 0.9.2342.19200300.100.1.8 NAME 'userClass' EQUALITY caseIgnoreMatch SUBSTR caseIgnoreSubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.15{256} X-ORIGIN 'RFC4524' )`)

	AllAttributeTypes = RFC4524AttributeTypes{
		AssociatedDomain,
		AssociatedName,
		BuildingName,
		CO,
		DocumentAuthor,
		DocumentIdentifier,
		DocumentLocation,
		DocumentPublisher,
		DocumentTitle,
		DocumentVersion,
		Drink,
		HomePhone,
		HomePostalAddress,
		Host,
		Info,
		Mail,
		Manager,
		Mobile,
		OrganizationalStatus,
		Pager,
		PersonalTitle,
		RoomNumber,
		Secretary,
		UniqueIdentifier,
		UserClass,
	}
}

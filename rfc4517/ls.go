package rfc4517

/*
RFC4517LDAPSyntaxes is a slice type designed to store RFC4517LDAPSyntax instances.
*/
type RFC4517LDAPSyntaxes []RFC4517LDAPSyntax

/*
RFC4517Syntax is a struct type that stores the raw RFC4517 syntax definition, along with a boolean value indicative of whether the syntax is considered human-readable.
*/
type RFC4517LDAPSyntax string

/*
LDAPSyntaxes contains slices of all instances of RFC4517LDAPSyntax defined in this package.
*/
var (
	AllLDAPSyntaxes RFC4517LDAPSyntaxes
)

var (
	ACIItem                     RFC4517LDAPSyntax
	AccessPoint                 RFC4517LDAPSyntax
	AttributeTypeDescription    RFC4517LDAPSyntax
	Audio                       RFC4517LDAPSyntax
	Binary                      RFC4517LDAPSyntax
	BitString                   RFC4517LDAPSyntax
	Boolean                     RFC4517LDAPSyntax
	CountryString               RFC4517LDAPSyntax
	DN                          RFC4517LDAPSyntax
	DataQualitySyntax           RFC4517LDAPSyntax
	DeliveryMethod              RFC4517LDAPSyntax
	DirectoryString             RFC4517LDAPSyntax
	DITContentRuleDescription   RFC4517LDAPSyntax
	DITStructureRuleDescription RFC4517LDAPSyntax
	DLSubmitPermission          RFC4517LDAPSyntax
	DSAQualitySyntax            RFC4517LDAPSyntax
	DSEType                     RFC4517LDAPSyntax
	EnhancedGuide               RFC4517LDAPSyntax
	FacsimileTelephoneNumber    RFC4517LDAPSyntax
	Fax                         RFC4517LDAPSyntax
	GeneralizedTime             RFC4517LDAPSyntax
	Guide                       RFC4517LDAPSyntax
	IA5String                   RFC4517LDAPSyntax
	INTEGER                     RFC4517LDAPSyntax
	JPEG                        RFC4517LDAPSyntax
	MatchingRuleDescription     RFC4517LDAPSyntax
	MatchingRuleUseDescription  RFC4517LDAPSyntax
	MailPreference              RFC4517LDAPSyntax
	MHSORAddress                RFC4517LDAPSyntax
	NameAndOptionalUID          RFC4517LDAPSyntax
	NameFormDescription         RFC4517LDAPSyntax
	NumericString               RFC4517LDAPSyntax
	ObjectClassDescription      RFC4517LDAPSyntax
	OID                         RFC4517LDAPSyntax
	OtherMailbox                RFC4517LDAPSyntax
	OctetString                 RFC4517LDAPSyntax
	PostalAddress               RFC4517LDAPSyntax
	ProtocolInformation         RFC4517LDAPSyntax
	PresentationAddress         RFC4517LDAPSyntax
	PrintableString             RFC4517LDAPSyntax
	SubtreeSpecification        RFC4517LDAPSyntax
	SupplierInformation         RFC4517LDAPSyntax
	SupplierOrConsumer          RFC4517LDAPSyntax
	SupplierAndConsumer         RFC4517LDAPSyntax
	TelephoneNumber             RFC4517LDAPSyntax
	TeletextTerminalIdentifier  RFC4517LDAPSyntax
	TelexNumber                 RFC4517LDAPSyntax
	UTCTime                     RFC4517LDAPSyntax
	LDAPSyntaxDescription       RFC4517LDAPSyntax
	ModifyRights                RFC4517LDAPSyntax
	LDAPSchemaDefinition        RFC4517LDAPSyntax
	LDAPSchemaDescription       RFC4517LDAPSyntax
	SubstringAssertion          RFC4517LDAPSyntax
)

func init() {
	ACIItem = RFC4517LDAPSyntax(`( 1.3.6.1.4.1.1466.115.121.1.1 DESC 'ACI Item' X-ORIGIN 'RFC4517' X-NOT-HUMAN-READABLE 'TRUE' )`)
	AccessPoint = RFC4517LDAPSyntax(`( 1.3.6.1.4.1.1466.115.121.1.2 DESC 'Access Point' X-ORIGIN 'RFC4517' )`)
	AttributeTypeDescription = RFC4517LDAPSyntax(`( 1.3.6.1.4.1.1466.115.121.1.3 DESC 'Attribute Type Description' X-ORIGIN 'RFC4517' )`)
	Audio = RFC4517LDAPSyntax(`( 1.3.6.1.4.1.1466.115.121.1.4 DESC 'Audio' X-ORIGIN 'RFC4517' X-NOT-HUMAN-READABLE 'TRUE' )`)
	Binary = RFC4517LDAPSyntax(`( 1.3.6.1.4.1.1466.115.121.1.5 DESC 'Binary' X-ORIGIN 'RFC4517' X-NOT-HUMAN-READABLE 'TRUE' )`)
	BitString = RFC4517LDAPSyntax(`( 1.3.6.1.4.1.1466.115.121.1.6 DESC 'Bit String' X-ORIGIN 'RFC4517' )`)
	Boolean = RFC4517LDAPSyntax(`( 1.3.6.1.4.1.1466.115.121.1.7 DESC 'Boolean' X-ORIGIN 'RFC4517' )`)
	CountryString = RFC4517LDAPSyntax(`( 1.3.6.1.4.1.1466.115.121.1.11 DESC 'Country String' X-ORIGIN 'RFC4517' )`)
	DN = RFC4517LDAPSyntax(`( 1.3.6.1.4.1.1466.115.121.1.12 DESC 'DN' X-ORIGIN 'RFC4517' )`)
	DataQualitySyntax = RFC4517LDAPSyntax(`( 1.3.6.1.4.1.1466.115.121.1.13 DESC 'Data Quality Syntax' X-ORIGIN 'RFC4517' )`)
	DeliveryMethod = RFC4517LDAPSyntax(`( 1.3.6.1.4.1.1466.115.121.1.14 DESC 'Delivery Method' X-ORIGIN 'RFC4517' )`)
	DirectoryString = RFC4517LDAPSyntax(`( 1.3.6.1.4.1.1466.115.121.1.15 DESC 'Directory String' X-ORIGIN 'RFC4517' )`)
	DITContentRuleDescription = RFC4517LDAPSyntax(`( 1.3.6.1.4.1.1466.115.121.1.16 DESC 'DIT Content Rule Description' X-ORIGIN 'RFC4517' )`)
	DITStructureRuleDescription = RFC4517LDAPSyntax(`( 1.3.6.1.4.1.1466.115.121.1.17 DESC 'DIT Structure Rule Description' X-ORIGIN 'RFC4517' )`)
	DLSubmitPermission = RFC4517LDAPSyntax(`( 1.3.6.1.4.1.1466.115.121.1.18 DESC 'DL Submit Permission' X-ORIGIN 'RFC4517' )`)
	DSAQualitySyntax = RFC4517LDAPSyntax(`( 1.3.6.1.4.1.1466.115.121.1.19 DESC 'DSA Quality Syntax' X-ORIGIN 'RFC4517' )`)
	DSEType = RFC4517LDAPSyntax(`( 1.3.6.1.4.1.1466.115.121.1.20 DESC 'DSE Type' X-ORIGIN 'RFC4517' )`)
	EnhancedGuide = RFC4517LDAPSyntax(`( 1.3.6.1.4.1.1466.115.121.1.21 DESC 'Enhanced Guide' X-ORIGIN 'RFC4517' )`)
	FacsimileTelephoneNumber = RFC4517LDAPSyntax(`( 1.3.6.1.4.1.1466.115.121.1.22 DESC 'Facsimile Telephone Number' X-ORIGIN 'RFC4517' )`)
	Fax = RFC4517LDAPSyntax(`( 1.3.6.1.4.1.1466.115.121.1.23 DESC 'Fax' X-ORIGIN 'RFC4517' )`)
	GeneralizedTime = RFC4517LDAPSyntax(`( 1.3.6.1.4.1.1466.115.121.1.24 DESC 'Generalized Time' X-ORIGIN 'RFC4517' )`)
	Guide = RFC4517LDAPSyntax(`( 1.3.6.1.4.1.1466.115.121.1.25 DESC 'Guide' X-ORIGIN 'RFC4517' )`)
	IA5String = RFC4517LDAPSyntax(`( 1.3.6.1.4.1.1466.115.121.1.26 DESC 'IA5 String' X-ORIGIN 'RFC4517' )`)
	INTEGER = RFC4517LDAPSyntax(`( 1.3.6.1.4.1.1466.115.121.1.27 DESC 'INTEGER' X-ORIGIN 'RFC4517' )`)
	JPEG = RFC4517LDAPSyntax(`( 1.3.6.1.4.1.1466.115.121.1.28 DESC 'JPEG' X-ORIGIN 'RFC4517' X-NOT-HUMAN-READABLE 'TRUE' )`)
	MatchingRuleDescription = RFC4517LDAPSyntax(`( 1.3.6.1.4.1.1466.115.121.1.30 DESC 'Matching Rule Description' X-ORIGIN 'RFC4517' )`)
	MatchingRuleUseDescription = RFC4517LDAPSyntax(`( 1.3.6.1.4.1.1466.115.121.1.31 DESC 'Matching Rule Use Description' X-ORIGIN 'RFC4517' )`)
	MailPreference = RFC4517LDAPSyntax(`( 1.3.6.1.4.1.1466.115.121.1.32 DESC 'Mail Preference' X-ORIGIN 'RFC4517' )`)
	MHSORAddress = RFC4517LDAPSyntax(`( 1.3.6.1.4.1.1466.115.121.1.33 DESC 'MHS OR Address' X-ORIGIN 'RFC4517' )`)
	NameAndOptionalUID = RFC4517LDAPSyntax(`( 1.3.6.1.4.1.1466.115.121.1.34 DESC 'Name And Optional UID' X-ORIGIN 'RFC4517' )`)
	NameFormDescription = RFC4517LDAPSyntax(`( 1.3.6.1.4.1.1466.115.121.1.35 DESC 'Name Form Description' X-ORIGIN 'RFC4517' )`)
	NumericString = RFC4517LDAPSyntax(`( 1.3.6.1.4.1.1466.115.121.1.36 DESC 'Numeric String' X-ORIGIN 'RFC4517' )`)
	ObjectClassDescription = RFC4517LDAPSyntax(`( 1.3.6.1.4.1.1466.115.121.1.37 DESC 'Object Class Description' X-ORIGIN 'RFC4517' )`)
	OID = RFC4517LDAPSyntax(`( 1.3.6.1.4.1.1466.115.121.1.38 DESC 'OID' X-ORIGIN 'RFC4517' )`)
	OtherMailbox = RFC4517LDAPSyntax(`( 1.3.6.1.4.1.1466.115.121.1.39 DESC 'Other Mailbox' X-ORIGIN 'RFC4517' )`)
	OctetString = RFC4517LDAPSyntax(`( 1.3.6.1.4.1.1466.115.121.1.40 DESC 'Octet String' X-ORIGIN 'RFC4517' )`)
	PostalAddress = RFC4517LDAPSyntax(`( 1.3.6.1.4.1.1466.115.121.1.41 DESC 'Postal Address' X-ORIGIN 'RFC4517' )`)
	ProtocolInformation = RFC4517LDAPSyntax(`( 1.3.6.1.4.1.1466.115.121.1.42 DESC 'Protocol Information' X-ORIGIN 'RFC4517' )`)
	PresentationAddress = RFC4517LDAPSyntax(`( 1.3.6.1.4.1.1466.115.121.1.43 DESC 'Presentation Address' X-ORIGIN 'RFC4517' )`)
	PrintableString = RFC4517LDAPSyntax(`( 1.3.6.1.4.1.1466.115.121.1.44 DESC 'Printable String' X-ORIGIN 'RFC4517' )`)
	SubtreeSpecification = RFC4517LDAPSyntax(`( 1.3.6.1.4.1.1466.115.121.1.45 DESC 'Subtree Specification' X-ORIGIN 'RFC4517' )`)
	SupplierInformation = RFC4517LDAPSyntax(`( 1.3.6.1.4.1.1466.115.121.1.46 DESC 'Supplier Information' X-ORIGIN 'RFC4517' )`)
	SupplierOrConsumer = RFC4517LDAPSyntax(`( 1.3.6.1.4.1.1466.115.121.1.47 DESC 'Supplier Or Consumer' X-ORIGIN 'RFC4517' )`)
	SupplierAndConsumer = RFC4517LDAPSyntax(`( 1.3.6.1.4.1.1466.115.121.1.48 DESC 'Supplier And Consumer' X-ORIGIN 'RFC4517' )`)
	TelephoneNumber = RFC4517LDAPSyntax(`( 1.3.6.1.4.1.1466.115.121.1.50 DESC 'Telephone Number' X-ORIGIN 'RFC4517' )`)
	TeletextTerminalIdentifier = RFC4517LDAPSyntax(`( 1.3.6.1.4.1.1466.115.121.1.51 DESC 'Teletext Terminal Identifier' X-ORIGIN 'RFC4517' )`)
	TelexNumber = RFC4517LDAPSyntax(`( 1.3.6.1.4.1.1466.115.121.1.52 DESC 'Telex Number' X-ORIGIN 'RFC4517' )`)
	UTCTime = RFC4517LDAPSyntax(`( 1.3.6.1.4.1.1466.115.121.1.53 DESC 'UTC Time' X-ORIGIN 'RFC4517' )`)
	LDAPSyntaxDescription = RFC4517LDAPSyntax(`( 1.3.6.1.4.1.1466.115.121.1.54 DESC 'LDAP Syntax Description' X-ORIGIN 'RFC4517' )`)
	ModifyRights = RFC4517LDAPSyntax(`( 1.3.6.1.4.1.1466.115.121.1.55 DESC 'Modify Rights' X-ORIGIN 'RFC4517' )`)
	LDAPSchemaDefinition = RFC4517LDAPSyntax(`( 1.3.6.1.4.1.1466.115.121.1.56 DESC 'LDAP Schema Definition' X-ORIGIN 'RFC4517' )`)
	LDAPSchemaDescription = RFC4517LDAPSyntax(`( 1.3.6.1.4.1.1466.115.121.1.57 DESC 'LDAP Schema Description' X-ORIGIN 'RFC4517' )`)
	SubstringAssertion = RFC4517LDAPSyntax(`( 1.3.6.1.4.1.1466.115.121.1.58 DESC 'Substring Assertion' X-ORIGIN 'RFC4517' )`)

	AllLDAPSyntaxes = RFC4517LDAPSyntaxes{
		ACIItem,
		AccessPoint,
		AttributeTypeDescription,
		Audio,
		Binary,
		BitString,
		Boolean,
		CountryString,
		DN,
		DataQualitySyntax,
		DeliveryMethod,
		DirectoryString,
		DITContentRuleDescription,
		DITStructureRuleDescription,
		DLSubmitPermission,
		DSAQualitySyntax,
		DSEType,
		EnhancedGuide,
		FacsimileTelephoneNumber,
		Fax,
		GeneralizedTime,
		Guide,
		IA5String,
		INTEGER,
		JPEG,
		MatchingRuleDescription,
		MatchingRuleUseDescription,
		MailPreference,
		MHSORAddress,
		NameAndOptionalUID,
		NameFormDescription,
		NumericString,
		ObjectClassDescription,
		OID,
		OtherMailbox,
		OctetString,
		PostalAddress,
		ProtocolInformation,
		PresentationAddress,
		PrintableString,
		SubtreeSpecification,
		SupplierInformation,
		SupplierOrConsumer,
		SupplierAndConsumer,
		TelephoneNumber,
		TeletextTerminalIdentifier,
		TelexNumber,
		UTCTime,
		LDAPSyntaxDescription,
		ModifyRights,
		LDAPSchemaDefinition,
		LDAPSchemaDescription,
		SubstringAssertion,
	}
}

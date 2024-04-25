package rfc4517

/*
LDAPSyntaxDefinitiones is a slice type designed to store LDAPSyntaxDefinition instances.
*/
type LDAPSyntaxDefinitions []LDAPSyntaxDefinition

/*
RFC4517Syntax is a struct type that stores the raw RFC4517 syntax definition, along with a boolean value indicative of whether the syntax is considered human-readable.
*/
type LDAPSyntaxDefinition string

/*
LDAPSyntaxes contains slices of all instances of LDAPSyntaxDefinition defined in this package.
*/
var (
	AllLDAPSyntaxes LDAPSyntaxDefinitions
)

var (
	ACIItem                     LDAPSyntaxDefinition
	AccessPoint                 LDAPSyntaxDefinition
	AttributeTypeDescription    LDAPSyntaxDefinition
	Audio                       LDAPSyntaxDefinition
	Binary                      LDAPSyntaxDefinition
	BitString                   LDAPSyntaxDefinition
	Boolean                     LDAPSyntaxDefinition
	CountryString               LDAPSyntaxDefinition
	DN                          LDAPSyntaxDefinition
	DataQualitySyntax           LDAPSyntaxDefinition
	DeliveryMethod              LDAPSyntaxDefinition
	DirectoryString             LDAPSyntaxDefinition
	DITContentRuleDescription   LDAPSyntaxDefinition
	DITStructureRuleDescription LDAPSyntaxDefinition
	DLSubmitPermission          LDAPSyntaxDefinition
	DSAQualitySyntax            LDAPSyntaxDefinition
	DSEType                     LDAPSyntaxDefinition
	EnhancedGuide               LDAPSyntaxDefinition
	FacsimileTelephoneNumber    LDAPSyntaxDefinition
	Fax                         LDAPSyntaxDefinition
	GeneralizedTime             LDAPSyntaxDefinition
	Guide                       LDAPSyntaxDefinition
	IA5String                   LDAPSyntaxDefinition
	INTEGER                     LDAPSyntaxDefinition
	JPEG                        LDAPSyntaxDefinition
	MatchingRuleDescription     LDAPSyntaxDefinition
	MatchingRuleUseDescription  LDAPSyntaxDefinition
	MailPreference              LDAPSyntaxDefinition
	MHSORAddress                LDAPSyntaxDefinition
	NameAndOptionalUID          LDAPSyntaxDefinition
	NameFormDescription         LDAPSyntaxDefinition
	NumericString               LDAPSyntaxDefinition
	ObjectClassDescription      LDAPSyntaxDefinition
	OID                         LDAPSyntaxDefinition
	OtherMailbox                LDAPSyntaxDefinition
	OctetString                 LDAPSyntaxDefinition
	PostalAddress               LDAPSyntaxDefinition
	ProtocolInformation         LDAPSyntaxDefinition
	PresentationAddress         LDAPSyntaxDefinition
	PrintableString             LDAPSyntaxDefinition
	SubtreeSpecification        LDAPSyntaxDefinition
	SupplierInformation         LDAPSyntaxDefinition
	SupplierOrConsumer          LDAPSyntaxDefinition
	SupplierAndConsumer         LDAPSyntaxDefinition
	TelephoneNumber             LDAPSyntaxDefinition
	TeletextTerminalIdentifier  LDAPSyntaxDefinition
	TelexNumber                 LDAPSyntaxDefinition
	UTCTime                     LDAPSyntaxDefinition
	LDAPSyntaxDescription       LDAPSyntaxDefinition
	ModifyRights                LDAPSyntaxDefinition
	LDAPSchemaDefinition        LDAPSyntaxDefinition
	LDAPSchemaDescription       LDAPSyntaxDefinition
	SubstringAssertion          LDAPSyntaxDefinition
)

func (r LDAPSyntaxDefinition) String() string {
	return string(r)
}

func init() {
	ACIItem = LDAPSyntaxDefinition(`( 1.3.6.1.4.1.1466.115.121.1.1 DESC 'ACI Item' X-ORIGIN 'RFC4517' X-NOT-HUMAN-READABLE 'TRUE' )`)
	AccessPoint = LDAPSyntaxDefinition(`( 1.3.6.1.4.1.1466.115.121.1.2 DESC 'Access Point' X-ORIGIN 'RFC4517' )`)
	AttributeTypeDescription = LDAPSyntaxDefinition(`( 1.3.6.1.4.1.1466.115.121.1.3 DESC 'Attribute Type Description' X-ORIGIN 'RFC4517' )`)
	Audio = LDAPSyntaxDefinition(`( 1.3.6.1.4.1.1466.115.121.1.4 DESC 'Audio' X-ORIGIN 'RFC4517' X-NOT-HUMAN-READABLE 'TRUE' )`)
	Binary = LDAPSyntaxDefinition(`( 1.3.6.1.4.1.1466.115.121.1.5 DESC 'Binary' X-ORIGIN 'RFC4517' X-NOT-HUMAN-READABLE 'TRUE' )`)
	BitString = LDAPSyntaxDefinition(`( 1.3.6.1.4.1.1466.115.121.1.6 DESC 'Bit String' X-ORIGIN 'RFC4517' )`)
	Boolean = LDAPSyntaxDefinition(`( 1.3.6.1.4.1.1466.115.121.1.7 DESC 'Boolean' X-ORIGIN 'RFC4517' )`)
	CountryString = LDAPSyntaxDefinition(`( 1.3.6.1.4.1.1466.115.121.1.11 DESC 'Country String' X-ORIGIN 'RFC4517' )`)
	DN = LDAPSyntaxDefinition(`( 1.3.6.1.4.1.1466.115.121.1.12 DESC 'DN' X-ORIGIN 'RFC4517' )`)
	DataQualitySyntax = LDAPSyntaxDefinition(`( 1.3.6.1.4.1.1466.115.121.1.13 DESC 'Data Quality Syntax' X-ORIGIN 'RFC4517' )`)
	DeliveryMethod = LDAPSyntaxDefinition(`( 1.3.6.1.4.1.1466.115.121.1.14 DESC 'Delivery Method' X-ORIGIN 'RFC4517' )`)
	DirectoryString = LDAPSyntaxDefinition(`( 1.3.6.1.4.1.1466.115.121.1.15 DESC 'Directory String' X-ORIGIN 'RFC4517' )`)
	DITContentRuleDescription = LDAPSyntaxDefinition(`( 1.3.6.1.4.1.1466.115.121.1.16 DESC 'DIT Content Rule Description' X-ORIGIN 'RFC4517' )`)
	DITStructureRuleDescription = LDAPSyntaxDefinition(`( 1.3.6.1.4.1.1466.115.121.1.17 DESC 'DIT Structure Rule Description' X-ORIGIN 'RFC4517' )`)
	DLSubmitPermission = LDAPSyntaxDefinition(`( 1.3.6.1.4.1.1466.115.121.1.18 DESC 'DL Submit Permission' X-ORIGIN 'RFC4517' )`)
	DSAQualitySyntax = LDAPSyntaxDefinition(`( 1.3.6.1.4.1.1466.115.121.1.19 DESC 'DSA Quality Syntax' X-ORIGIN 'RFC4517' )`)
	DSEType = LDAPSyntaxDefinition(`( 1.3.6.1.4.1.1466.115.121.1.20 DESC 'DSE Type' X-ORIGIN 'RFC4517' )`)
	EnhancedGuide = LDAPSyntaxDefinition(`( 1.3.6.1.4.1.1466.115.121.1.21 DESC 'Enhanced Guide' X-ORIGIN 'RFC4517' )`)
	FacsimileTelephoneNumber = LDAPSyntaxDefinition(`( 1.3.6.1.4.1.1466.115.121.1.22 DESC 'Facsimile Telephone Number' X-ORIGIN 'RFC4517' )`)
	Fax = LDAPSyntaxDefinition(`( 1.3.6.1.4.1.1466.115.121.1.23 DESC 'Fax' X-ORIGIN 'RFC4517' )`)
	GeneralizedTime = LDAPSyntaxDefinition(`( 1.3.6.1.4.1.1466.115.121.1.24 DESC 'Generalized Time' X-ORIGIN 'RFC4517' )`)
	Guide = LDAPSyntaxDefinition(`( 1.3.6.1.4.1.1466.115.121.1.25 DESC 'Guide' X-ORIGIN 'RFC4517' )`)
	IA5String = LDAPSyntaxDefinition(`( 1.3.6.1.4.1.1466.115.121.1.26 DESC 'IA5 String' X-ORIGIN 'RFC4517' )`)
	INTEGER = LDAPSyntaxDefinition(`( 1.3.6.1.4.1.1466.115.121.1.27 DESC 'INTEGER' X-ORIGIN 'RFC4517' )`)
	JPEG = LDAPSyntaxDefinition(`( 1.3.6.1.4.1.1466.115.121.1.28 DESC 'JPEG' X-ORIGIN 'RFC4517' X-NOT-HUMAN-READABLE 'TRUE' )`)
	MatchingRuleDescription = LDAPSyntaxDefinition(`( 1.3.6.1.4.1.1466.115.121.1.30 DESC 'Matching Rule Description' X-ORIGIN 'RFC4517' )`)
	MatchingRuleUseDescription = LDAPSyntaxDefinition(`( 1.3.6.1.4.1.1466.115.121.1.31 DESC 'Matching Rule Use Description' X-ORIGIN 'RFC4517' )`)
	MailPreference = LDAPSyntaxDefinition(`( 1.3.6.1.4.1.1466.115.121.1.32 DESC 'Mail Preference' X-ORIGIN 'RFC4517' )`)
	MHSORAddress = LDAPSyntaxDefinition(`( 1.3.6.1.4.1.1466.115.121.1.33 DESC 'MHS OR Address' X-ORIGIN 'RFC4517' )`)
	NameAndOptionalUID = LDAPSyntaxDefinition(`( 1.3.6.1.4.1.1466.115.121.1.34 DESC 'Name And Optional UID' X-ORIGIN 'RFC4517' )`)
	NameFormDescription = LDAPSyntaxDefinition(`( 1.3.6.1.4.1.1466.115.121.1.35 DESC 'Name Form Description' X-ORIGIN 'RFC4517' )`)
	NumericString = LDAPSyntaxDefinition(`( 1.3.6.1.4.1.1466.115.121.1.36 DESC 'Numeric String' X-ORIGIN 'RFC4517' )`)
	ObjectClassDescription = LDAPSyntaxDefinition(`( 1.3.6.1.4.1.1466.115.121.1.37 DESC 'Object Class Description' X-ORIGIN 'RFC4517' )`)
	OID = LDAPSyntaxDefinition(`( 1.3.6.1.4.1.1466.115.121.1.38 DESC 'OID' X-ORIGIN 'RFC4517' )`)
	OtherMailbox = LDAPSyntaxDefinition(`( 1.3.6.1.4.1.1466.115.121.1.39 DESC 'Other Mailbox' X-ORIGIN 'RFC4517' )`)
	OctetString = LDAPSyntaxDefinition(`( 1.3.6.1.4.1.1466.115.121.1.40 DESC 'Octet String' X-ORIGIN 'RFC4517' )`)
	PostalAddress = LDAPSyntaxDefinition(`( 1.3.6.1.4.1.1466.115.121.1.41 DESC 'Postal Address' X-ORIGIN 'RFC4517' )`)
	ProtocolInformation = LDAPSyntaxDefinition(`( 1.3.6.1.4.1.1466.115.121.1.42 DESC 'Protocol Information' X-ORIGIN 'RFC4517' )`)
	PresentationAddress = LDAPSyntaxDefinition(`( 1.3.6.1.4.1.1466.115.121.1.43 DESC 'Presentation Address' X-ORIGIN 'RFC4517' )`)
	PrintableString = LDAPSyntaxDefinition(`( 1.3.6.1.4.1.1466.115.121.1.44 DESC 'Printable String' X-ORIGIN 'RFC4517' )`)
	SubtreeSpecification = LDAPSyntaxDefinition(`( 1.3.6.1.4.1.1466.115.121.1.45 DESC 'Subtree Specification' X-ORIGIN 'RFC4517' )`)
	SupplierInformation = LDAPSyntaxDefinition(`( 1.3.6.1.4.1.1466.115.121.1.46 DESC 'Supplier Information' X-ORIGIN 'RFC4517' )`)
	SupplierOrConsumer = LDAPSyntaxDefinition(`( 1.3.6.1.4.1.1466.115.121.1.47 DESC 'Supplier Or Consumer' X-ORIGIN 'RFC4517' )`)
	SupplierAndConsumer = LDAPSyntaxDefinition(`( 1.3.6.1.4.1.1466.115.121.1.48 DESC 'Supplier And Consumer' X-ORIGIN 'RFC4517' )`)
	TelephoneNumber = LDAPSyntaxDefinition(`( 1.3.6.1.4.1.1466.115.121.1.50 DESC 'Telephone Number' X-ORIGIN 'RFC4517' )`)
	TeletextTerminalIdentifier = LDAPSyntaxDefinition(`( 1.3.6.1.4.1.1466.115.121.1.51 DESC 'Teletext Terminal Identifier' X-ORIGIN 'RFC4517' )`)
	TelexNumber = LDAPSyntaxDefinition(`( 1.3.6.1.4.1.1466.115.121.1.52 DESC 'Telex Number' X-ORIGIN 'RFC4517' )`)
	UTCTime = LDAPSyntaxDefinition(`( 1.3.6.1.4.1.1466.115.121.1.53 DESC 'UTC Time' X-ORIGIN 'RFC4517' )`)
	LDAPSyntaxDescription = LDAPSyntaxDefinition(`( 1.3.6.1.4.1.1466.115.121.1.54 DESC 'LDAP Syntax Description' X-ORIGIN 'RFC4517' )`)
	ModifyRights = LDAPSyntaxDefinition(`( 1.3.6.1.4.1.1466.115.121.1.55 DESC 'Modify Rights' X-ORIGIN 'RFC4517' )`)
	LDAPSchemaDefinition = LDAPSyntaxDefinition(`( 1.3.6.1.4.1.1466.115.121.1.56 DESC 'LDAP Schema Definition' X-ORIGIN 'RFC4517' )`)
	LDAPSchemaDescription = LDAPSyntaxDefinition(`( 1.3.6.1.4.1.1466.115.121.1.57 DESC 'LDAP Schema Description' X-ORIGIN 'RFC4517' )`)
	SubstringAssertion = LDAPSyntaxDefinition(`( 1.3.6.1.4.1.1466.115.121.1.58 DESC 'Substring Assertion' X-ORIGIN 'RFC4517' )`)

	AllLDAPSyntaxes = LDAPSyntaxDefinitions{
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

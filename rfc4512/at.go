package rfc4512

type RFC4512AttributeTypes []RFC4512AttributeType
type RFC4512AttributeType string

var (
	DirectorySchemaAttributeTypes RFC4512AttributeTypes
	OperationalAttributeTypes     RFC4512AttributeTypes
	DSAAttributeTypes             RFC4512AttributeTypes
	AllAttributeTypes             RFC4512AttributeTypes
)

// Operational AttributeTypes
var (
	AliasedObjectName      RFC4512AttributeType
	ObjectClass            RFC4512AttributeType
	CreatorsName           RFC4512AttributeType
	CreateTimestamp        RFC4512AttributeType
	ModifiersName          RFC4512AttributeType
	ModifyTimestamp        RFC4512AttributeType
	StructuralObjectClass  RFC4512AttributeType
	GoverningStructureRule RFC4512AttributeType
)

// Directory Schema AttributeTypes
var (
	ObjectClasses     RFC4512AttributeType
	SubschemaSubentry RFC4512AttributeType
	AttributeTypes    RFC4512AttributeType
	MatchingRules     RFC4512AttributeType
	MatchingRuleUse   RFC4512AttributeType
	LDAPSyntaxes      RFC4512AttributeType
	DITContentRules   RFC4512AttributeType
	DITStructureRules RFC4512AttributeType
	NameForms         RFC4512AttributeType
)

// DSA AttributeTypes
var (
	AltServer               RFC4512AttributeType
	NamingContexts          RFC4512AttributeType
	SupportedControl        RFC4512AttributeType
	SupportedExtension      RFC4512AttributeType
	SupportedFeatures       RFC4512AttributeType
	SupportedLDAPVersion    RFC4512AttributeType
	SupportedSASLMechanisms RFC4512AttributeType
)

func init() {
	ObjectClasses = RFC4512AttributeType(`( 2.5.21.6 NAME 'objectClasses' EQUALITY objectIdentifierFirstComponentMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.37 USAGE directoryOperation X-ORIGIN 'RFC4512' )`)
	SubschemaSubentry = RFC4512AttributeType(`( 2.5.18.10 NAME 'subschemaSubentry' EQUALITY distinguishedNameMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.12 SINGLE-VALUE NO-USER-MODIFICATION USAGE directoryOperation X-ORIGIN 'RFC4512' )`)
	AttributeTypes = RFC4512AttributeType(`( 2.5.21.5 NAME 'attributeTypes' EQUALITY objectIdentifierFirstComponentMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.3 USAGE directoryOperation X-ORIGIN 'RFC4512' )`)
	MatchingRules = RFC4512AttributeType(`( 2.5.21.4 NAME 'matchingRules' EQUALITY objectIdentifierFirstComponentMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.30 USAGE directoryOperation X-ORIGIN 'RFC4512' )`)
	MatchingRuleUse = RFC4512AttributeType(`( 2.5.21.8 NAME 'matchingRuleUse' EQUALITY objectIdentifierFirstComponentMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.31 USAGE directoryOperation X-ORIGIN 'RFC4512' )`)
	LDAPSyntaxes = RFC4512AttributeType(`( 1.3.6.1.4.1.1466.101.120.16 NAME 'ldapSyntaxes' EQUALITY objectIdentifierFirstComponentMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.54 USAGE directoryOperation X-ORIGIN 'RFC4512' )`)
	DITContentRules = RFC4512AttributeType(`( 2.5.21.2 NAME 'dITContentRules' EQUALITY objectIdentifierFirstComponentMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.16 USAGE directoryOperation X-ORIGIN 'RFC4512' )`)
	DITStructureRules = RFC4512AttributeType(`( 2.5.21.1 NAME 'dITStructureRules' EQUALITY integerFirstComponentMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.17 USAGE directoryOperation X-ORIGIN 'RFC4512' )`)
	NameForms = RFC4512AttributeType(`( 2.5.21.7 NAME 'nameForms' EQUALITY objectIdentifierFirstComponentMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.35 USAGE directoryOperation X-ORIGIN 'RFC4512' )`)

	AltServer = RFC4512AttributeType(`( 1.3.6.1.4.1.1466.101.120.6 NAME 'altServer' SYNTAX 1.3.6.1.4.1.1466.115.121.1.26 USAGE dSAOperation X-ORIGIN 'RFC4512' )`)
	NamingContexts = RFC4512AttributeType(`( 1.3.6.1.4.1.1466.101.120.5 NAME 'namingContexts' SYNTAX 1.3.6.1.4.1.1466.115.121.1.12 USAGE dSAOperation X-ORIGIN 'RFC4512' )`)
	SupportedControl = RFC4512AttributeType(`( 1.3.6.1.4.1.1466.101.120.13 NAME 'supportedControl' SYNTAX 1.3.6.1.4.1.1466.115.121.1.38 USAGE dSAOperation X-ORIGIN 'RFC4512' )`)
	SupportedExtension = RFC4512AttributeType(`( 1.3.6.1.4.1.1466.101.120.7 NAME 'supportedExtension' SYNTAX 1.3.6.1.4.1.1466.115.121.1.38 USAGE dSAOperation X-ORIGIN 'RFC4512' )`)
	SupportedFeatures = RFC4512AttributeType(`( 1.3.6.1.4.1.4203.1.3.5 NAME 'supportedFeatures' EQUALITY objectIdentifierMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.38 USAGE dSAOperation X-ORIGIN 'RFC4512' )`)
	SupportedLDAPVersion = RFC4512AttributeType(`( 1.3.6.1.4.1.1466.101.120.15 NAME 'supportedLDAPVersion' SYNTAX 1.3.6.1.4.1.1466.115.121.1.27 USAGE dSAOperation X-ORIGIN 'RFC4512' )`)
	SupportedSASLMechanisms = RFC4512AttributeType(`( 1.3.6.1.4.1.1466.101.120.14 NAME 'supportedSASLMechanisms' SYNTAX 1.3.6.1.4.1.1466.115.121.1.15 USAGE dSAOperation X-ORIGIN 'RFC4512' )`)

	AliasedObjectName = RFC4512AttributeType(`( 2.5.4.1 NAME 'aliasedObjectName' EQUALITY distinguishedNameMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.12 SINGLE-VALUE X-ORIGIN 'RFC4512' )`)
	ObjectClass = RFC4512AttributeType(`( 2.5.4.0 NAME 'objectClass' EQUALITY objectIdentifierMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.38 X-ORIGIN 'RFC4512' )`)
	CreatorsName = RFC4512AttributeType(`( 2.5.18.3 NAME 'creatorsName' EQUALITY distinguishedNameMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.12 SINGLE-VALUE NO-USER-MODIFICATION USAGE directoryOperation X-ORIGIN 'RFC4512' )`)
	CreateTimestamp = RFC4512AttributeType(`( 2.5.18.1 NAME 'createTimestamp' EQUALITY generalizedTimeMatch ORDERING generalizedTimeOrderingMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.24 SINGLE-VALUE NO-USER-MODIFICATION USAGE directoryOperation X-ORIGIN 'RFC4512' )`)
	ModifiersName = RFC4512AttributeType(`( 2.5.18.4 NAME 'modifiersName' EQUALITY distinguishedNameMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.12 SINGLE-VALUE NO-USER-MODIFICATION USAGE directoryOperation X-ORIGIN 'RFC4512' )`)
	ModifyTimestamp = RFC4512AttributeType(`( 2.5.18.2 NAME 'modifyTimestamp' EQUALITY generalizedTimeMatch ORDERING generalizedTimeOrderingMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.24 SINGLE-VALUE NO-USER-MODIFICATION USAGE directoryOperation X-ORIGIN 'RFC4512' )`)
	StructuralObjectClass = RFC4512AttributeType(`( 2.5.21.9 NAME 'structuralObjectClass' EQUALITY objectIdentifierMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.38 SINGLE-VALUE NO-USER-MODIFICATION USAGE directoryOperation X-ORIGIN 'RFC4512' )`)
	GoverningStructureRule = RFC4512AttributeType(`( 2.5.21.10 NAME 'governingStructureRule' EQUALITY integerMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.27 SINGLE-VALUE NO-USER-MODIFICATION USAGE directoryOperation X-ORIGIN 'RFC4512' )`)
}

func init() {

	OperationalAttributeTypes = RFC4512AttributeTypes{
		AliasedObjectName,
		ObjectClass,
		CreatorsName,
		CreateTimestamp,
		ModifiersName,
		ModifyTimestamp,
		StructuralObjectClass,
		GoverningStructureRule,
	}
	opATlen := len(OperationalAttributeTypes)

	DirectorySchemaAttributeTypes = RFC4512AttributeTypes{
		ObjectClasses,
		SubschemaSubentry,
		AttributeTypes,
		MatchingRules,
		MatchingRuleUse,
		LDAPSyntaxes,
		DITContentRules,
		DITStructureRules,
		NameForms,
	}
	dsATlen := len(DirectorySchemaAttributeTypes)

	DSAAttributeTypes = RFC4512AttributeTypes{
		AltServer,
		NamingContexts,
		SupportedControl,
		SupportedExtension,
		SupportedFeatures,
		SupportedLDAPVersion,
		SupportedSASLMechanisms,
	}
	dsaATlen := len(DSAAttributeTypes)

	total := opATlen + dsATlen + dsaATlen
	cur := 0
	AllAttributeTypes = make(RFC4512AttributeTypes, total, total)

	for _, v := range []RFC4512AttributeTypes{
		OperationalAttributeTypes,
		DirectorySchemaAttributeTypes,
		DSAAttributeTypes} {
		for i := 0; i < len(v); i++ {
			cur++
			AllAttributeTypes[cur-1] = v[i]
		}
	}
}

package rfc4512

type AttributeTypeDefinitions []AttributeTypeDefinition
type AttributeTypeDefinition string

var (
	DirectorySchemaAttributeTypes AttributeTypeDefinitions
	OperationalAttributeTypes     AttributeTypeDefinitions
	DSAAttributeTypes             AttributeTypeDefinitions
	AllAttributeTypes             AttributeTypeDefinitions
)

// Operational AttributeTypes
var (
	AliasedObjectName      AttributeTypeDefinition
	ObjectClass            AttributeTypeDefinition
	CreatorsName           AttributeTypeDefinition
	CreateTimestamp        AttributeTypeDefinition
	ModifiersName          AttributeTypeDefinition
	ModifyTimestamp        AttributeTypeDefinition
	StructuralObjectClass  AttributeTypeDefinition
	GoverningStructureRule AttributeTypeDefinition
)

// Directory Schema AttributeTypes
var (
	ObjectClasses     AttributeTypeDefinition
	SubschemaSubentry AttributeTypeDefinition
	AttributeTypes    AttributeTypeDefinition
	MatchingRules     AttributeTypeDefinition
	MatchingRuleUse   AttributeTypeDefinition
	LDAPSyntaxes      AttributeTypeDefinition
	DITContentRules   AttributeTypeDefinition
	DITStructureRules AttributeTypeDefinition
	NameForms         AttributeTypeDefinition
)

// DSA AttributeTypes
var (
	AltServer               AttributeTypeDefinition
	NamingContexts          AttributeTypeDefinition
	SupportedControl        AttributeTypeDefinition
	SupportedExtension      AttributeTypeDefinition
	SupportedFeatures       AttributeTypeDefinition
	SupportedLDAPVersion    AttributeTypeDefinition
	SupportedSASLMechanisms AttributeTypeDefinition
)

func (r AttributeTypeDefinition) String() string {
	return string(r)
}

func init() {
	ObjectClasses = AttributeTypeDefinition(`( 2.5.21.6 NAME 'objectClasses' EQUALITY objectIdentifierFirstComponentMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.37 USAGE directoryOperation X-ORIGIN 'RFC4512' )`)
	SubschemaSubentry = AttributeTypeDefinition(`( 2.5.18.10 NAME 'subschemaSubentry' EQUALITY distinguishedNameMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.12 SINGLE-VALUE NO-USER-MODIFICATION USAGE directoryOperation X-ORIGIN 'RFC4512' )`)
	AttributeTypes = AttributeTypeDefinition(`( 2.5.21.5 NAME 'attributeTypes' EQUALITY objectIdentifierFirstComponentMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.3 USAGE directoryOperation X-ORIGIN 'RFC4512' )`)
	MatchingRules = AttributeTypeDefinition(`( 2.5.21.4 NAME 'matchingRules' EQUALITY objectIdentifierFirstComponentMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.30 USAGE directoryOperation X-ORIGIN 'RFC4512' )`)
	MatchingRuleUse = AttributeTypeDefinition(`( 2.5.21.8 NAME 'matchingRuleUse' EQUALITY objectIdentifierFirstComponentMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.31 USAGE directoryOperation X-ORIGIN 'RFC4512' )`)
	LDAPSyntaxes = AttributeTypeDefinition(`( 1.3.6.1.4.1.1466.101.120.16 NAME 'ldapSyntaxes' EQUALITY objectIdentifierFirstComponentMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.54 USAGE directoryOperation X-ORIGIN 'RFC4512' )`)
	DITContentRules = AttributeTypeDefinition(`( 2.5.21.2 NAME 'dITContentRules' EQUALITY objectIdentifierFirstComponentMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.16 USAGE directoryOperation X-ORIGIN 'RFC4512' )`)
	DITStructureRules = AttributeTypeDefinition(`( 2.5.21.1 NAME 'dITStructureRules' EQUALITY integerFirstComponentMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.17 USAGE directoryOperation X-ORIGIN 'RFC4512' )`)
	NameForms = AttributeTypeDefinition(`( 2.5.21.7 NAME 'nameForms' EQUALITY objectIdentifierFirstComponentMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.35 USAGE directoryOperation X-ORIGIN 'RFC4512' )`)

	AltServer = AttributeTypeDefinition(`( 1.3.6.1.4.1.1466.101.120.6 NAME 'altServer' SYNTAX 1.3.6.1.4.1.1466.115.121.1.26 USAGE dSAOperation X-ORIGIN 'RFC4512' )`)
	NamingContexts = AttributeTypeDefinition(`( 1.3.6.1.4.1.1466.101.120.5 NAME 'namingContexts' SYNTAX 1.3.6.1.4.1.1466.115.121.1.12 USAGE dSAOperation X-ORIGIN 'RFC4512' )`)
	SupportedControl = AttributeTypeDefinition(`( 1.3.6.1.4.1.1466.101.120.13 NAME 'supportedControl' SYNTAX 1.3.6.1.4.1.1466.115.121.1.38 USAGE dSAOperation X-ORIGIN 'RFC4512' )`)
	SupportedExtension = AttributeTypeDefinition(`( 1.3.6.1.4.1.1466.101.120.7 NAME 'supportedExtension' SYNTAX 1.3.6.1.4.1.1466.115.121.1.38 USAGE dSAOperation X-ORIGIN 'RFC4512' )`)
	SupportedFeatures = AttributeTypeDefinition(`( 1.3.6.1.4.1.4203.1.3.5 NAME 'supportedFeatures' EQUALITY objectIdentifierMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.38 USAGE dSAOperation X-ORIGIN 'RFC4512' )`)
	SupportedLDAPVersion = AttributeTypeDefinition(`( 1.3.6.1.4.1.1466.101.120.15 NAME 'supportedLDAPVersion' SYNTAX 1.3.6.1.4.1.1466.115.121.1.27 USAGE dSAOperation X-ORIGIN 'RFC4512' )`)
	SupportedSASLMechanisms = AttributeTypeDefinition(`( 1.3.6.1.4.1.1466.101.120.14 NAME 'supportedSASLMechanisms' SYNTAX 1.3.6.1.4.1.1466.115.121.1.15 USAGE dSAOperation X-ORIGIN 'RFC4512' )`)

	AliasedObjectName = AttributeTypeDefinition(`( 2.5.4.1 NAME 'aliasedObjectName' EQUALITY distinguishedNameMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.12 SINGLE-VALUE X-ORIGIN 'RFC4512' )`)
	ObjectClass = AttributeTypeDefinition(`( 2.5.4.0 NAME 'objectClass' EQUALITY objectIdentifierMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.38 X-ORIGIN 'RFC4512' )`)
	CreatorsName = AttributeTypeDefinition(`( 2.5.18.3 NAME 'creatorsName' EQUALITY distinguishedNameMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.12 SINGLE-VALUE NO-USER-MODIFICATION USAGE directoryOperation X-ORIGIN 'RFC4512' )`)
	CreateTimestamp = AttributeTypeDefinition(`( 2.5.18.1 NAME 'createTimestamp' EQUALITY generalizedTimeMatch ORDERING generalizedTimeOrderingMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.24 SINGLE-VALUE NO-USER-MODIFICATION USAGE directoryOperation X-ORIGIN 'RFC4512' )`)
	ModifiersName = AttributeTypeDefinition(`( 2.5.18.4 NAME 'modifiersName' EQUALITY distinguishedNameMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.12 SINGLE-VALUE NO-USER-MODIFICATION USAGE directoryOperation X-ORIGIN 'RFC4512' )`)
	ModifyTimestamp = AttributeTypeDefinition(`( 2.5.18.2 NAME 'modifyTimestamp' EQUALITY generalizedTimeMatch ORDERING generalizedTimeOrderingMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.24 SINGLE-VALUE NO-USER-MODIFICATION USAGE directoryOperation X-ORIGIN 'RFC4512' )`)
	StructuralObjectClass = AttributeTypeDefinition(`( 2.5.21.9 NAME 'structuralObjectClass' EQUALITY objectIdentifierMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.38 SINGLE-VALUE NO-USER-MODIFICATION USAGE directoryOperation X-ORIGIN 'RFC4512' )`)
	GoverningStructureRule = AttributeTypeDefinition(`( 2.5.21.10 NAME 'governingStructureRule' EQUALITY integerMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.27 SINGLE-VALUE NO-USER-MODIFICATION USAGE directoryOperation X-ORIGIN 'RFC4512' )`)

	OperationalAttributeTypes = AttributeTypeDefinitions{
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

	DirectorySchemaAttributeTypes = AttributeTypeDefinitions{
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

	DSAAttributeTypes = AttributeTypeDefinitions{
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
	AllAttributeTypes = make(AttributeTypeDefinitions, total, total)

	for _, v := range []AttributeTypeDefinitions{
		OperationalAttributeTypes,
		DirectorySchemaAttributeTypes,
		DSAAttributeTypes} {
		for i := 0; i < len(v); i++ {
			cur++
			AllAttributeTypes[cur-1] = v[i]
		}
	}
}

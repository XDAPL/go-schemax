package rfc4517

/*
MatchingRuleDefinitions is a slice type designed to store instances of MatchingRuleDefinition.
*/
type MatchingRuleDefinitions []MatchingRuleDefinition

func (r MatchingRuleDefinitions) Len() int {
	return len(r)
}

/*
MatchingRuleDefinition is a string type designed to store a raw MatchingRule definition.
*/
type MatchingRuleDefinition string

/*
MatchingRules contains slices of all instances of MatchingRuleDefinition defined in this package.
*/
var AllMatchingRules MatchingRuleDefinitions

var (
	ProtocolInformationMatch            MatchingRuleDefinition // Yes, technically this is RFC2252
	PresentationAddressMatch            MatchingRuleDefinition // Yes, technically this is RFC2256
	CaseIgnoreIA5SubstringsMatch        MatchingRuleDefinition
	CaseIgnoreIA5Match                  MatchingRuleDefinition
	CaseExactIA5Match                   MatchingRuleDefinition
	ObjectIdentifierMatch               MatchingRuleDefinition
	DistinguishedNameMatch              MatchingRuleDefinition
	CaseIgnoreMatch                     MatchingRuleDefinition
	CaseIgnoreOrderingMatch             MatchingRuleDefinition
	CaseIgnoreSubstringsMatch           MatchingRuleDefinition
	CaseExactMatch                      MatchingRuleDefinition
	CaseExactOrderingMatch              MatchingRuleDefinition
	CaseExactSubstringsMatch            MatchingRuleDefinition
	NumericStringMatch                  MatchingRuleDefinition
	NumericStringOrderingMatch          MatchingRuleDefinition
	NumericStringSubstringsMatch        MatchingRuleDefinition
	CaseIgnoreListMatch                 MatchingRuleDefinition
	CaseIgnoreListSubstringsMatch       MatchingRuleDefinition
	BooleanMatch                        MatchingRuleDefinition
	IntegerMatch                        MatchingRuleDefinition
	IntegerOrderingMatch                MatchingRuleDefinition
	BitStringMatch                      MatchingRuleDefinition
	OctetStringMatch                    MatchingRuleDefinition
	OctetStringOrderingMatch            MatchingRuleDefinition
	TelephoneNumberMatch                MatchingRuleDefinition
	TelephoneNumberSubstringsMatch      MatchingRuleDefinition
	UniqueMemberMatch                   MatchingRuleDefinition
	GeneralizedTimeMatch                MatchingRuleDefinition
	GeneralizedTimeOrderingMatch        MatchingRuleDefinition
	IntegerFirstComponentMatch          MatchingRuleDefinition
	ObjectIdentifierFirstComponentMatch MatchingRuleDefinition
	DirectoryStringFirstComponentMatch  MatchingRuleDefinition
	WordMatch                           MatchingRuleDefinition
	KeywordMatch                        MatchingRuleDefinition
)

func (r MatchingRuleDefinition) String() string {
	return string(r)
}

func init() {
	PresentationAddressMatch = MatchingRuleDefinition(`( 2.5.13.22 NAME 'presentationAddressMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.43 X-ORIGIN 'RFC2256' )`)
	ProtocolInformationMatch = MatchingRuleDefinition(`( 2.5.13.24 NAME 'protocolInformationMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.42 X-ORIGIN 'RFC2252' )`)
	CaseIgnoreIA5SubstringsMatch = MatchingRuleDefinition(`( 1.3.6.1.4.1.1466.109.114.3 NAME 'caseIgnoreIA5SubstringsMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.58 X-ORIGIN 'RFC4517' )`)
	CaseIgnoreIA5Match = MatchingRuleDefinition(`( 1.3.6.1.4.1.1466.109.114.2 NAME 'caseIgnoreIA5Match' SYNTAX 1.3.6.1.4.1.1466.115.121.1.26 X-ORIGIN 'RFC4517' )`)
	CaseExactIA5Match = MatchingRuleDefinition(`( 1.3.6.1.4.1.1466.109.114.1 NAME 'caseExactIA5Match' SYNTAX 1.3.6.1.4.1.1466.115.121.1.26 X-ORIGIN 'RFC4517' )`)
	ObjectIdentifierMatch = MatchingRuleDefinition(`( 2.5.13.0 NAME 'objectIdentifierMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.38 X-ORIGIN 'RFC4517' )`)
	DistinguishedNameMatch = MatchingRuleDefinition(`( 2.5.13.1 NAME 'distinguishedNameMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.12 X-ORIGIN 'RFC4517' )`)
	CaseIgnoreMatch = MatchingRuleDefinition(`( 2.5.13.2 NAME 'caseIgnoreMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.15 X-ORIGIN 'RFC4517' )`)
	CaseIgnoreOrderingMatch = MatchingRuleDefinition(`( 2.5.13.3 NAME 'caseIgnoreOrderingMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.15 X-ORIGIN 'RFC4517' )`)
	CaseIgnoreSubstringsMatch = MatchingRuleDefinition(`( 2.5.13.4 NAME 'caseIgnoreSubstringsMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.58 X-ORIGIN 'RFC4517' )`)
	CaseExactMatch = MatchingRuleDefinition(`( 2.5.13.5 NAME 'caseExactMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.15 X-ORIGIN 'RFC4517' )`)
	CaseExactOrderingMatch = MatchingRuleDefinition(`( 2.5.13.6 NAME 'caseExactOrderingMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.15 X-ORIGIN 'RFC4517' )`)
	CaseExactSubstringsMatch = MatchingRuleDefinition(`( 2.5.13.7 NAME 'caseExactSubstringsMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.58 X-ORIGIN 'RFC4517' )`)
	NumericStringMatch = MatchingRuleDefinition(`( 2.5.13.8 NAME 'numericStringMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.36 X-ORIGIN 'RFC4517' )`)
	NumericStringOrderingMatch = MatchingRuleDefinition(`( 2.5.13.9 NAME 'numericStringOrderingMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.36 X-ORIGIN 'RFC4517' )`)
	NumericStringSubstringsMatch = MatchingRuleDefinition(`( 2.5.13.10 NAME 'numericStringSubstringsMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.58 X-ORIGIN 'RFC4517' )`)
	CaseIgnoreListMatch = MatchingRuleDefinition(`( 2.5.13.11 NAME 'caseIgnoreListMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.41 X-ORIGIN 'RFC4517' )`)
	CaseIgnoreListSubstringsMatch = MatchingRuleDefinition(`( 2.5.13.12 NAME 'caseIgnoreListSubstringsMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.58 X-ORIGIN 'RFC4517' )`)
	BooleanMatch = MatchingRuleDefinition(`( 2.5.13.13 NAME 'booleanMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.7 X-ORIGIN 'RFC4517' )`)
	IntegerMatch = MatchingRuleDefinition(`( 2.5.13.14 NAME 'integerMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.27 X-ORIGIN 'RFC4517' )`)
	IntegerOrderingMatch = MatchingRuleDefinition(`( 2.5.13.15 NAME 'integerOrderingMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.27 X-ORIGIN 'RFC4517' )`)
	BitStringMatch = MatchingRuleDefinition(`( 2.5.13.16 NAME 'bitStringMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.6 X-ORIGIN 'RFC4517' )`)
	OctetStringMatch = MatchingRuleDefinition(`( 2.5.13.17 NAME 'octetStringMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.40 X-ORIGIN 'RFC4517' )`)
	OctetStringOrderingMatch = MatchingRuleDefinition(`( 2.5.13.18 NAME 'octetStringOrderingMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.40 X-ORIGIN 'RFC4517' )`)
	TelephoneNumberMatch = MatchingRuleDefinition(`( 2.5.13.20 NAME 'telephoneNumberMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.50 X-ORIGIN 'RFC4517' )`)
	TelephoneNumberSubstringsMatch = MatchingRuleDefinition(`( 2.5.13.21 NAME 'telephoneNumberSubstringsMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.58 X-ORIGIN 'RFC4517' )`)
	UniqueMemberMatch = MatchingRuleDefinition(`( 2.5.13.23 NAME 'uniqueMemberMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.34 X-ORIGIN 'RFC4517' )`)
	GeneralizedTimeMatch = MatchingRuleDefinition(`( 2.5.13.27 NAME 'generalizedTimeMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.24 X-ORIGIN 'RFC4517' )`)
	GeneralizedTimeOrderingMatch = MatchingRuleDefinition(`( 2.5.13.28 NAME 'generalizedTimeOrderingMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.24 X-ORIGIN 'RFC4517' )`)
	IntegerFirstComponentMatch = MatchingRuleDefinition(`( 2.5.13.29 NAME 'integerFirstComponentMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.27 X-ORIGIN 'RFC4517' )`)
	ObjectIdentifierFirstComponentMatch = MatchingRuleDefinition(`( 2.5.13.30 NAME 'objectIdentifierFirstComponentMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.38 X-ORIGIN 'RFC4517' )`)
	DirectoryStringFirstComponentMatch = MatchingRuleDefinition(`( 2.5.13.31 NAME 'directoryStringFirstComponentMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.15 X-ORIGIN 'RFC4517' )`)
	WordMatch = MatchingRuleDefinition(`( 2.5.13.32 NAME 'wordMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.15 X-ORIGIN 'RFC4517' )`)
	KeywordMatch = MatchingRuleDefinition(`( 2.5.13.33 NAME 'keywordMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.15 X-ORIGIN 'RFC4517' )`)

	AllMatchingRules = MatchingRuleDefinitions{
		ObjectIdentifierMatch,
		DistinguishedNameMatch,
		CaseIgnoreMatch,
		CaseIgnoreOrderingMatch,
		CaseIgnoreSubstringsMatch,
		CaseExactMatch,
		CaseExactOrderingMatch,
		CaseExactSubstringsMatch,
		NumericStringMatch,
		NumericStringOrderingMatch,
		NumericStringSubstringsMatch,
		CaseIgnoreIA5Match,
		CaseExactIA5Match,
		CaseIgnoreIA5SubstringsMatch,
		CaseIgnoreListMatch,
		CaseIgnoreListSubstringsMatch,
		BooleanMatch,
		IntegerMatch,
		IntegerOrderingMatch,
		BitStringMatch,
		OctetStringMatch,
		OctetStringOrderingMatch,
		PresentationAddressMatch,
		ProtocolInformationMatch,
		TelephoneNumberMatch,
		TelephoneNumberSubstringsMatch,
		UniqueMemberMatch,
		GeneralizedTimeMatch,
		GeneralizedTimeOrderingMatch,
		IntegerFirstComponentMatch,
		ObjectIdentifierFirstComponentMatch,
		DirectoryStringFirstComponentMatch,
		WordMatch,
		KeywordMatch,
	}
}

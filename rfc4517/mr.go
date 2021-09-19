package rfc4517

/*
RFC4517MatchingRules is a slice type designed to store instances of RFC4517MatchingRule.
*/
type RFC4517MatchingRules []RFC4517MatchingRule

/*
RFC4517MatchingRule is a string type designed to store a raw MatchingRule definition.
*/
type RFC4517MatchingRule string

/*
MatchingRules contains slices of all instances of RFC4517MatchingRule defined in this package.
*/
var AllMatchingRules RFC4517MatchingRules

var (
	ProtocolInformationMatch            RFC4517MatchingRule // Yes, technically this is RFC2252
	PresentationAddressMatch            RFC4517MatchingRule // Yes, technically this is RFC2256
	CaseIgnoreIA5SubstringsMatch        RFC4517MatchingRule
	CaseIgnoreIA5Match                  RFC4517MatchingRule
	CaseExactIA5Match                   RFC4517MatchingRule
	ObjectIdentifierMatch               RFC4517MatchingRule
	DistinguishedNameMatch              RFC4517MatchingRule
	CaseIgnoreMatch                     RFC4517MatchingRule
	CaseIgnoreOrderingMatch             RFC4517MatchingRule
	CaseIgnoreSubstringsMatch           RFC4517MatchingRule
	CaseExactMatch                      RFC4517MatchingRule
	CaseExactOrderingMatch              RFC4517MatchingRule
	CaseExactSubstringsMatch            RFC4517MatchingRule
	NumericStringMatch                  RFC4517MatchingRule
	NumericStringOrderingMatch          RFC4517MatchingRule
	NumericStringSubstringsMatch        RFC4517MatchingRule
	CaseIgnoreListMatch                 RFC4517MatchingRule
	CaseIgnoreListSubstringsMatch       RFC4517MatchingRule
	BooleanMatch                        RFC4517MatchingRule
	IntegerMatch                        RFC4517MatchingRule
	IntegerOrderingMatch                RFC4517MatchingRule
	BitStringMatch                      RFC4517MatchingRule
	OctetStringMatch                    RFC4517MatchingRule
	OctetStringOrderingMatch            RFC4517MatchingRule
	TelephoneNumberMatch                RFC4517MatchingRule
	TelephoneNumberSubstringsMatch      RFC4517MatchingRule
	UniqueMemberMatch                   RFC4517MatchingRule
	GeneralizedTimeMatch                RFC4517MatchingRule
	GeneralizedTimeOrderingMatch        RFC4517MatchingRule
	IntegerFirstComponentMatch          RFC4517MatchingRule
	ObjectIdentifierFirstComponentMatch RFC4517MatchingRule
	DirectoryStringFirstComponentMatch  RFC4517MatchingRule
	WordMatch                           RFC4517MatchingRule
	KeywordMatch                        RFC4517MatchingRule
)

func init() {
	PresentationAddressMatch = RFC4517MatchingRule(`( 2.5.13.22 NAME 'presentationAddressMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.43 X-ORIGIN 'RFC2256' )`)
	ProtocolInformationMatch = RFC4517MatchingRule(`( 2.5.13.24 NAME 'protocolInformationMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.42 X-ORIGIN 'RFC2252' )`)
	CaseIgnoreIA5SubstringsMatch = RFC4517MatchingRule(`( 1.3.6.1.4.1.1466.109.114.3 NAME 'caseIgnoreIA5SubstringsMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.58 X-ORIGIN 'RFC4517' )`)
	CaseIgnoreIA5Match = RFC4517MatchingRule(`( 1.3.6.1.4.1.1466.109.114.2 NAME 'caseIgnoreIA5Match' SYNTAX 1.3.6.1.4.1.1466.115.121.1.26 X-ORIGIN 'RFC4517' )`)
	CaseExactIA5Match = RFC4517MatchingRule(`( 1.3.6.1.4.1.1466.109.114.1 NAME 'caseExactIA5Match' SYNTAX 1.3.6.1.4.1.1466.115.121.1.26 X-ORIGIN 'RFC4517' )`)
	ObjectIdentifierMatch = RFC4517MatchingRule(`( 2.5.13.0 NAME 'objectIdentifierMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.38 X-ORIGIN 'RFC4517' )`)
	DistinguishedNameMatch = RFC4517MatchingRule(`( 2.5.13.1 NAME 'distinguishedNameMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.12 X-ORIGIN 'RFC4517' )`)
	CaseIgnoreMatch = RFC4517MatchingRule(`( 2.5.13.2 NAME 'caseIgnoreMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.15 X-ORIGIN 'RFC4517' )`)
	CaseIgnoreOrderingMatch = RFC4517MatchingRule(`( 2.5.13.3 NAME 'caseIgnoreOrderingMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.15 X-ORIGIN 'RFC4517' )`)
	CaseIgnoreSubstringsMatch = RFC4517MatchingRule(`( 2.5.13.4 NAME 'caseIgnoreSubstringsMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.58 X-ORIGIN 'RFC4517' )`)
	CaseExactMatch = RFC4517MatchingRule(`( 2.5.13.5 NAME 'caseExactMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.15 X-ORIGIN 'RFC4517' )`)
	CaseExactOrderingMatch = RFC4517MatchingRule(`( 2.5.13.6 NAME 'caseExactOrderingMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.15 X-ORIGIN 'RFC4517' )`)
	CaseExactSubstringsMatch = RFC4517MatchingRule(`( 2.5.13.7 NAME 'caseExactSubstringsMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.58 X-ORIGIN 'RFC4517' )`)
	NumericStringMatch = RFC4517MatchingRule(`( 2.5.13.8 NAME 'numericStringMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.36 X-ORIGIN 'RFC4517' )`)
	NumericStringOrderingMatch = RFC4517MatchingRule(`( 2.5.13.9 NAME 'numericStringOrderingMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.36 X-ORIGIN 'RFC4517' )`)
	NumericStringSubstringsMatch = RFC4517MatchingRule(`( 2.5.13.10 NAME 'numericStringSubstringsMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.58 X-ORIGIN 'RFC4517' )`)
	CaseIgnoreListMatch = RFC4517MatchingRule(`( 2.5.13.11 NAME 'caseIgnoreListMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.41 X-ORIGIN 'RFC4517' )`)
	CaseIgnoreListSubstringsMatch = RFC4517MatchingRule(`( 2.5.13.12 NAME 'caseIgnoreListSubstringsMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.58 X-ORIGIN 'RFC4517' )`)
	BooleanMatch = RFC4517MatchingRule(`( 2.5.13.13 NAME 'booleanMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.7 X-ORIGIN 'RFC4517' )`)
	IntegerMatch = RFC4517MatchingRule(`( 2.5.13.14 NAME 'integerMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.27 X-ORIGIN 'RFC4517' )`)
	IntegerOrderingMatch = RFC4517MatchingRule(`( 2.5.13.15 NAME 'integerOrderingMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.27 X-ORIGIN 'RFC4517' )`)
	BitStringMatch = RFC4517MatchingRule(`( 2.5.13.16 NAME 'bitStringMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.6 X-ORIGIN 'RFC4517' )`)
	OctetStringMatch = RFC4517MatchingRule(`( 2.5.13.17 NAME 'octetStringMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.40 X-ORIGIN 'RFC4517' )`)
	OctetStringOrderingMatch = RFC4517MatchingRule(`( 2.5.13.18 NAME 'octetStringOrderingMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.40 X-ORIGIN 'RFC4517' )`)
	TelephoneNumberMatch = RFC4517MatchingRule(`( 2.5.13.20 NAME 'telephoneNumberMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.50 X-ORIGIN 'RFC4517' )`)
	TelephoneNumberSubstringsMatch = RFC4517MatchingRule(`( 2.5.13.21 NAME 'telephoneNumberSubstringsMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.58 X-ORIGIN 'RFC4517' )`)
	UniqueMemberMatch = RFC4517MatchingRule(`( 2.5.13.23 NAME 'uniqueMemberMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.34 X-ORIGIN 'RFC4517' )`)
	GeneralizedTimeMatch = RFC4517MatchingRule(`( 2.5.13.27 NAME 'generalizedTimeMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.24 X-ORIGIN 'RFC4517' )`)
	GeneralizedTimeOrderingMatch = RFC4517MatchingRule(`( 2.5.13.28 NAME 'generalizedTimeOrderingMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.24 X-ORIGIN 'RFC4517' )`)
	IntegerFirstComponentMatch = RFC4517MatchingRule(`( 2.5.13.29 NAME 'integerFirstComponentMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.27 X-ORIGIN 'RFC4517' )`)
	ObjectIdentifierFirstComponentMatch = RFC4517MatchingRule(`( 2.5.13.30 NAME 'objectIdentifierFirstComponentMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.38 X-ORIGIN 'RFC4517' )`)
	DirectoryStringFirstComponentMatch = RFC4517MatchingRule(`( 2.5.13.31 NAME 'directoryStringFirstComponentMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.15 X-ORIGIN 'RFC4517' )`)
	WordMatch = RFC4517MatchingRule(`( 2.5.13.32 NAME 'wordMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.15 X-ORIGIN 'RFC4517' )`)
	KeywordMatch = RFC4517MatchingRule(`( 2.5.13.33 NAME 'keywordMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.15 X-ORIGIN 'RFC4517' )`)

	AllMatchingRules = []RFC4517MatchingRule{
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

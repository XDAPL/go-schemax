package rfc4523

/*
RFC4523MatchingRules is a slice type designed to store instances of RFC4523MatchingRule.
*/
type RFC4523MatchingRules []RFC4523MatchingRule

/*
RFC4523MatchingRule is a string type designed to store a raw MatchingRule definition.
*/
type RFC4523MatchingRule string

/*
MatchingRules contains slices of all instances of RFC4523MatchingRule defined in this package.
*/
var AllMatchingRules RFC4523MatchingRules

var (
	CertificateExactMatch     RFC4523MatchingRule
	CertificateMatch          RFC4523MatchingRule
	CertificatePairExactMatch RFC4523MatchingRule
	CertificatePairMatch      RFC4523MatchingRule
	CertificateListExactMatch RFC4523MatchingRule
	CertificateListMatch      RFC4523MatchingRule
	AlgorithmIdentifierMatch  RFC4523MatchingRule
)

func init() {

	CertificateExactMatch = RFC4523MatchingRule(`( 2.5.13.34 NAME 'certificateExactMatch' SYNTAX 1.3.6.1.1.15.1 X-ORIGIN 'RFC4523' )`)
	CertificateMatch = RFC4523MatchingRule(`( 2.5.13.35 NAME 'certificateMatch' SYNTAX 1.3.6.1.1.15.2 X-ORIGIN 'RFC4523' )`)
	CertificatePairExactMatch = RFC4523MatchingRule(`( 2.5.13.36 NAME 'certificatePairExactMatch' SYNTAX 1.3.6.1.1.15.3 X-ORIGIN 'RFC4523' )`)
	CertificatePairMatch = RFC4523MatchingRule(`( 2.5.13.37 NAME 'certificatePairMatch' SYNTAX 1.3.6.1.1.15.4 X-ORIGIN 'RFC4523' )`)
	CertificateListExactMatch = RFC4523MatchingRule(`( 2.5.13.38 NAME 'certificateListExactMatch' SYNTAX 1.3.6.1.1.15.5 X-ORIGIN 'RFC4523' )`)
	CertificateListMatch = RFC4523MatchingRule(`( 2.5.13.39 NAME 'certificateListMatch' SYNTAX 1.3.6.1.1.15.6 X-ORIGIN 'RFC4523' )`)
	AlgorithmIdentifierMatch = RFC4523MatchingRule(`( 2.5.13.40 NAME 'algorithmIdentifierMatch' SYNTAX 1.3.6.1.1.15.7 X-ORIGIN 'RFC4523' )`)
	// AlgorithmIdentifierMatch is incorrectly labeled 'algorithmIdentifier'
	// as the NAME, contrary to the section header in RFC4523 at the top of
	// page 7, and contrary to literally every other matching rule in this
	// RFC, so ...

	AllMatchingRules = []RFC4523MatchingRule{
		CertificateExactMatch,
		CertificateMatch,
		CertificatePairExactMatch,
		CertificatePairMatch,
		CertificateListExactMatch,
		CertificateListMatch,
		AlgorithmIdentifierMatch,
	}
}

package rfc4523

/*
MatchingRuleDefinitions is a slice type designed to store instances of MatchingRuleDefinition.
*/
type MatchingRuleDefinitions []MatchingRuleDefinition

/*
MatchingRuleDefinition is a string type designed to store a raw MatchingRule definition.
*/
type MatchingRuleDefinition string

/*
MatchingRules contains slices of all instances of MatchingRuleDefinition defined in this package.
*/
var AllMatchingRules MatchingRuleDefinitions

var (
	CertificateExactMatch     MatchingRuleDefinition
	CertificateMatch          MatchingRuleDefinition
	CertificatePairExactMatch MatchingRuleDefinition
	CertificatePairMatch      MatchingRuleDefinition
	CertificateListExactMatch MatchingRuleDefinition
	CertificateListMatch      MatchingRuleDefinition
	AlgorithmIdentifierMatch  MatchingRuleDefinition
)

func (r MatchingRuleDefinition) String() string {
	return string(r)
}

func init() {

	CertificateExactMatch = MatchingRuleDefinition(`( 2.5.13.34 NAME 'certificateExactMatch' SYNTAX 1.3.6.1.1.15.1 X-ORIGIN 'RFC4523' )`)
	CertificateMatch = MatchingRuleDefinition(`( 2.5.13.35 NAME 'certificateMatch' SYNTAX 1.3.6.1.1.15.2 X-ORIGIN 'RFC4523' )`)
	CertificatePairExactMatch = MatchingRuleDefinition(`( 2.5.13.36 NAME 'certificatePairExactMatch' SYNTAX 1.3.6.1.1.15.3 X-ORIGIN 'RFC4523' )`)
	CertificatePairMatch = MatchingRuleDefinition(`( 2.5.13.37 NAME 'certificatePairMatch' SYNTAX 1.3.6.1.1.15.4 X-ORIGIN 'RFC4523' )`)
	CertificateListExactMatch = MatchingRuleDefinition(`( 2.5.13.38 NAME 'certificateListExactMatch' SYNTAX 1.3.6.1.1.15.5 X-ORIGIN 'RFC4523' )`)
	CertificateListMatch = MatchingRuleDefinition(`( 2.5.13.39 NAME 'certificateListMatch' SYNTAX 1.3.6.1.1.15.6 X-ORIGIN 'RFC4523' )`)
	AlgorithmIdentifierMatch = MatchingRuleDefinition(`( 2.5.13.40 NAME 'algorithmIdentifierMatch' SYNTAX 1.3.6.1.1.15.7 X-ORIGIN 'RFC4523' )`)
	// AlgorithmIdentifierMatch is incorrectly labeled 'algorithmIdentifier'
	// as the NAME, contrary to the section header in RFC4523 at the top of
	// page 7, and contrary to literally every other matching rule in this
	// RFC, so ...

	AllMatchingRules = MatchingRuleDefinitions{
		CertificateExactMatch,
		CertificateMatch,
		CertificatePairExactMatch,
		CertificatePairMatch,
		CertificateListExactMatch,
		CertificateListMatch,
		AlgorithmIdentifierMatch,
	}
}

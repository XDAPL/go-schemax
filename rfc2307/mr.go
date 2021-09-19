package rfc2307

/*
RFC2307MatchingRules is a slice type designed to store instances of RFC2307MatchingRule.
*/
type RFC2307MatchingRules []RFC2307MatchingRule

/*
RFC2307MatchingRule is a string type designed to store a raw MatchingRule definition.
*/
type RFC2307MatchingRule string

/*
MatchingRules contains slices of all instances of RFC2307MatchingRule defined in this package.
*/
var AllMatchingRules RFC2307MatchingRules

var (
	CaseExactIA5SubstringsMatch RFC2307MatchingRule // No, technically this didn't come from RFC2307
	// but there's nowhere else to put it, but the
	// RFC in question does reference it :/
)

func init() {
	CaseExactIA5SubstringsMatch = RFC2307MatchingRule(`( 1.3.6.1.4.1.4203.1.2.1 NAME 'caseExactIA5SubstringsMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.26 X-ORIGIN 'RFC2307' )`)

	AllMatchingRules = []RFC2307MatchingRule{
		CaseExactIA5SubstringsMatch,
	}
}

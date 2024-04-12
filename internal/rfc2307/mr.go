package rfc2307

/*
MatchingRuleDefinitions is a slice type designed to store instances of MatchingRuleDefinition.
*/
type MatchingRuleDefinitions []MatchingRuleDefinition

/*
MatchingRuleDefinition is a string type designed to store a raw MatchingRule definition.
*/
type MatchingRuleDefinition string

/*
AllMatchingRules contains slices of all instances of MatchingRuleDefinition defined in this package.
*/
var AllMatchingRules MatchingRuleDefinitions

var (
	CaseExactIA5SubstringsMatch MatchingRuleDefinition // No, technically this didn't come from RFC2307
	// but there's nowhere else to put it, and the
	// RFC in question does reference it :/
)

func init() {
	CaseExactIA5SubstringsMatch = MatchingRuleDefinition(`( 1.3.6.1.4.1.4203.1.2.1 NAME 'caseExactIA5SubstringsMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.26 X-ORIGIN 'RFC2307' )`)

	AllMatchingRules = []MatchingRuleDefinition{
		CaseExactIA5SubstringsMatch,
	}
}

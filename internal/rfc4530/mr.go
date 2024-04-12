package rfc4530

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
	UUIDMatch,
	UUIDOrderingMatch MatchingRuleDefinition
)

func init() {

	UUIDMatch = MatchingRuleDefinition(`( 1.3.6.1.1.16.2 NAME 'uuidMatch' SYNTAX 1.3.6.1.1.16.1 X-ORIGIN 'RFC4530' )`)
	UUIDOrderingMatch = MatchingRuleDefinition(`( 1.3.6.1.1.16.3 NAME 'uuidOrderingMatch' SYNTAX 1.3.6.1.1.16.1 X-ORIGIN 'RFC4530' )`)

	AllMatchingRules = []MatchingRuleDefinition{
		UUIDMatch,
		UUIDOrderingMatch,
	}
}

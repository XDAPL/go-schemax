package rfc4530

/*
RFC4530MatchingRules is a slice type designed to store instances of RFC4530MatchingRule.
*/
type RFC4530MatchingRules []RFC4530MatchingRule

/*
RFC4530MatchingRule is a string type designed to store a raw MatchingRule definition.
*/
type RFC4530MatchingRule string

/*
MatchingRules contains slices of all instances of RFC4530MatchingRule defined in this package.
*/
var AllMatchingRules RFC4530MatchingRules

var (
	UUIDMatch,
	UUIDOrderingMatch RFC4530MatchingRule
)

func init() {

	UUIDMatch = RFC4530MatchingRule(`( 1.3.6.1.1.16.2 NAME 'uuidMatch' SYNTAX 1.3.6.1.1.16.1 X-ORIGIN 'RFC4530' )`)
	UUIDOrderingMatch = RFC4530MatchingRule(`( 1.3.6.1.1.16.3 NAME 'uuidOrderingMatch' SYNTAX 1.3.6.1.1.16.1 X-ORIGIN 'RFC4530' )`)

	AllMatchingRules = []RFC4530MatchingRule{
		UUIDMatch,
		UUIDOrderingMatch,
	}
}

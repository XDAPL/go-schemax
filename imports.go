package schemax

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"internal/rfc2079"
	"internal/rfc2307"
	"internal/rfc2798"
	"internal/rfc3045"
	"internal/rfc3671"
	"internal/rfc3672"
	"internal/rfc4512"
	"internal/rfc4517"
	"internal/rfc4519"
	"internal/rfc4523"
	"internal/rfc4524"
	"internal/rfc4530"
)

var (
	atoi    func(string) (int, error)           = strconv.Atoi
	itoa    func(int) string                    = strconv.Itoa
	sprintf func(string, ...any) string         = fmt.Sprintf
	printf  func(string, ...any) (int, error)   = fmt.Printf
	repAll  func(string, string, string) string = strings.ReplaceAll
	eq      func(string, string) bool           = strings.EqualFold
	split   func(string, string) []string       = strings.Split
	join    func([]string, string) string       = strings.Join
	hasPfx  func(string, string) bool           = strings.HasPrefix
	hasSfx  func(string, string) bool           = strings.HasSuffix
	lc      func(string) string                 = strings.ToLower
	uc      func(string) string                 = strings.ToUpper
	trim    func(string, string) string         = strings.Trim
	trimL   func(string, string) string         = strings.TrimLeft
	trimR   func(string, string) string         = strings.TrimRight
	trimS   func(string) string                 = strings.TrimSpace
	isDigit func(rune) bool                     = unicode.IsDigit
	isAlpha func(rune) bool                     = unicode.IsLetter
)

var (
	rfc2307Macros map[string]string = rfc2307.Macros

	rfc2079AttributeTypes rfc2079.AttributeTypeDefinitions = rfc2079.AllAttributeTypes
	rfc2307AttributeTypes rfc2307.AttributeTypeDefinitions = rfc2307.AllAttributeTypes
	rfc2798AttributeTypes rfc2798.AttributeTypeDefinitions = rfc2798.AllAttributeTypes
	rfc3045AttributeTypes rfc3045.AttributeTypeDefinitions = rfc3045.AllAttributeTypes
	rfc3671AttributeTypes rfc3671.AttributeTypeDefinitions = rfc3671.AllAttributeTypes
	rfc3672AttributeTypes rfc3672.AttributeTypeDefinitions = rfc3672.AllAttributeTypes
	rfc4512AttributeTypes rfc4512.AttributeTypeDefinitions = rfc4512.AllAttributeTypes
	rfc4519AttributeTypes rfc4519.AttributeTypeDefinitions = rfc4519.AllAttributeTypes
	rfc4523AttributeTypes rfc4523.AttributeTypeDefinitions = rfc4523.AllAttributeTypes
	rfc4524AttributeTypes rfc4524.AttributeTypeDefinitions = rfc4524.AllAttributeTypes
	rfc4530AttributeTypes rfc4530.AttributeTypeDefinitions = rfc4530.AllAttributeTypes

	rfc2079ObjectClasses rfc2079.ObjectClassDefinitions = rfc2079.AllObjectClasses
	rfc2307ObjectClasses rfc2307.ObjectClassDefinitions = rfc2307.AllObjectClasses
	rfc2798ObjectClasses rfc2798.ObjectClassDefinitions = rfc2798.AllObjectClasses
	rfc3671ObjectClasses rfc3671.ObjectClassDefinitions = rfc3671.AllObjectClasses
	rfc3672ObjectClasses rfc3672.ObjectClassDefinitions = rfc3672.AllObjectClasses
	rfc4512ObjectClasses rfc4512.ObjectClassDefinitions = rfc4512.AllObjectClasses
	rfc4519ObjectClasses rfc4519.ObjectClassDefinitions = rfc4519.AllObjectClasses
	rfc4523ObjectClasses rfc4523.ObjectClassDefinitions = rfc4523.AllObjectClasses
	rfc4524ObjectClasses rfc4524.ObjectClassDefinitions = rfc4524.AllObjectClasses

	rfc4517LDAPSyntaxes rfc4517.LDAPSyntaxDefinitions = rfc4517.AllLDAPSyntaxes
	rfc2307LDAPSyntaxes rfc2307.LDAPSyntaxDefinitions = rfc2307.AllLDAPSyntaxes
	rfc4523LDAPSyntaxes rfc4523.LDAPSyntaxDefinitions = rfc4523.AllLDAPSyntaxes
	rfc4530LDAPSyntaxes rfc4530.LDAPSyntaxDefinitions = rfc4530.AllLDAPSyntaxes

	rfc2307MatchingRules rfc2307.MatchingRuleDefinitions = rfc2307.AllMatchingRules
	rfc4517MatchingRules rfc4517.MatchingRuleDefinitions = rfc4517.AllMatchingRules
	rfc4523MatchingRules rfc4523.MatchingRuleDefinitions = rfc4523.AllMatchingRules
	rfc4530MatchingRules rfc4530.MatchingRuleDefinitions = rfc4530.AllMatchingRules
)

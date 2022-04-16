package schemax

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

/*
common import wrappers
*/
var (
        itoa         func(int) string                          = strconv.Itoa
        atoi         func(string) (int, error)                 = strconv.Atoi
        toLower      func(string) string                       = strings.ToLower
        join         func([]string, string) string             = strings.Join
        split        func(string, string) []string             = strings.Split
        contains     func(string, string) bool                 = strings.Contains
        replaceAll   func(string, string, string) string       = strings.ReplaceAll
        equalFold    func(string, string) bool                 = strings.EqualFold
        indexRune    func(string, rune) int                    = strings.IndexRune
        index        func(string, string) int                  = strings.Index
        trimSpace    func(string) string                       = strings.TrimSpace
	trim         func(string, string) string               = strings.Trim
        runeIsUpper  func(rune) bool                           = unicode.IsUpper
        runeIsLetter func(rune) bool                           = unicode.IsLetter
        runeIsDigit  func(rune) bool                           = unicode.IsDigit
        isUTF8       func([]byte) bool                         = utf8.Valid
        valueOf      func(interface{}) reflect.Value           = reflect.ValueOf
        printf       func(string, ...interface{}) (int, error) = fmt.Printf
        sprintf      func(string, ...interface{}) string       = fmt.Sprintf
        newErr       func(string) error                        = errors.New
)

// sanity limits
var (
        descMaxLen     = 4096 // bytes
        nameListMaxLen = 10   // per def
        nameMaxLen     = 128  // single name length
)

// Default Definition Collections
var (
        DefaultAttributeTypes AttributeTypeCollection
        DefaultObjectClasses  ObjectClassCollection
        DefaultLDAPSyntaxes   LDAPSyntaxCollection
        DefaultMatchingRules  MatchingRuleCollection
)

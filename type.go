package schemax

/*
type.go deals with type definition and instantiation, as well as global variable definition:
*/

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
	genListMaxLen  = 1024 // max genericList len
)

// Default Definition Manifests
var (
	DefaultAttributeTypesManifest AttributeTypesManifest
	DefaultObjectClassesManifest  ObjectClassesManifest
	DefaultLDAPSyntaxesManifest   LDAPSyntaxesManifest
	DefaultMatchingRulesManifest  MatchingRulesManifest
)

/*
parseMeth is the first class function bearing a signature shared by all fundamental parser methods.
*/
type parseMeth func(string) ([]string, string, bool)

/*
definition can describe the nature and contents of any single specific schema definition:

 - an AttributeType instance, or ...
 - an ObjectClass instance, or ...
 - an LDAPSyntax instance, or ...
 - a MatchingRule instance, or ...
 - a MatchingRuleUse instance, or ...
 - a DITContentRule instance, or ...
 - a DITStructureRule instance, or ...
 - a NameForm instance

The "typ" field MUST only contain one (1) of the above.

Length/Cap of "fields" and "values" MUST always be equal, even if no actual field _values_ were perceived.

This type is only defined for operations inherent to this package.

The alm field stores the OPTIONAL user-provided AliasesManifest instance for the resolution of OID aliases (a.k.a.: macros).
*/
type definition struct {
	alm    AliasesManifest
	typ    reflect.Type
	fields []reflect.StructField
	values []reflect.Value
	labels map[int]string
	meths  map[int]parseMeth
}

/*
Boolean is an unsigned 8-bit integer that describes zero or more perceived values that only manifest visually when TRUE. Such verisimilitude is revealed by the presence of the indicated Boolean value's "label" name, such as `SINGLE-VALUE`.  The actual value "TRUE" is never actually seen in textual format.
*/
type Boolean uint8

const (
	Obsolete           Boolean = 1 << iota // 1
	SingleValue                            // 2
	Collective                             // 4
	NoUserModification                     // 8
	HumanReadable                          // 16 (applies only to *LDAPSyntax instances)
)

/*
RuleID describes a numerical identifier for an instance of DITStructureRule.
*/
type RuleID uint

/*
OID describes any object identifier assigned to any schema type.
*/
type OID string

/*
Description manifests as a single text value intended to describe the nature and purpose of the definition bearing this type in human readable format.
*/
type Description string

/*
Usage describes the intended usage of an AttributeType definition as a single text value.  This can be one of four constant values, the first of which (userApplication) is implied in the absence of any other value and is not necessary to reveal in such a case.
*/
type Usage uint8

const (
	UserApplication Usage = iota
	DirectoryOperation
	DistributedOperation
	DSAOperation
)

/*
Kind is an unsigned 8-bit integer that describes the "kind" of ObjectClass definition bearing this type.  Only one distinct Kind value may be set for any given ObjectClass definition, and must be set explicitly (no default is implied).
*/
type Kind uint8

const (
	badKind Kind = iota
	Abstract
	Structural
	Auxiliary
)

/*
Subschema provides high-level definition management, storage of all definitions of schemata and convenient method wrappers.

The first field of this type, "DN", contains a string representation of the appropriate subschema distinguished name.  Oftentimes, this will be something like "cn=schema", or "cn=subschema" but may vary between directory implementations.

All remaining fields within this type are based on the Manifest (map) types defined in this package.  For example, Subschema "NFM" field is of the NameFormsManifest type, and is intended to store all successfully-parsed "nameForm" definitions.

The user must initialize these map-based fields before use, whether they do so themselves manually, OR by use of a suitable "populator function", such as the PopulateDefaultLDAPSyntaxes() function to populate the "LSM" field.

The purpose of this type is to centralize all of the relevant Manifests that the user would otherwise have to manage individually.  This is useful for portability reasons, and allows an entire library of X.501 schema definitions to be stored within a single object.

This type also provides high-level oversight of what would otherwise be isolated low-level operations that may or may not be invalid.  For example, when operating in a purely low-level fashion (without the use of *Subschema), there is nothing stopping a user from adding a so-called "RequiredAttributeTypes" slice member to a "ProhibitedAttributeTypes" slice.  Each of those slices is throughly unaware of the other.   However, when conducting this operation via a convenient method extended by *Subschema, additional correlative checks can (and will!) be conducted to avoid such invalid actions.

Overall, use of *Subschema makes generalized use of this package slightly easier but is NOT necessarily required in all situations.
*/
type Subschema struct {
	DN   string                    // often "cn=schema" or "cn=subschema", varies per imp.
	ALM  AliasesManifest           // OPTIONAL aliases      OID->string (is used for ALL other OID parses)
	LSM  LDAPSyntaxesManifest      // LDAP Syntaxes         OID->*LDAPSyntax
	MRM  MatchingRulesManifest     // Matching Rules        OID->*MatchingRule
	ATM  AttributeTypesManifest    // Attribute Types       OID->*AttributeType
	OCM  ObjectClassesManifest     // Object Classes        OID->*ObjectClass
	NFM  NameFormsManifest         // Name Forms            OID->*NameForm
	MRUM MatchingRuleUsesManifest  // Matching Rule "Uses"  OID->*MatchingRuleUse
	DCRM DITContentRulesManifest   // DIT Content Rules     OID->*DITContentRule
	DSRM DITStructureRulesManifest // DIT Structure Rules   ID->*DITStructureRule
}

/*
ObjectClass conforms to the specifications of RFC4512 Section 4.1.1. Boolean values, e.g: 'OBSOLETE', are supported internally and are not explicit fields.
*/
type ObjectClass struct {
	OID         OID
	Name        Name
	Description Description
	SuperClass  SuperiorObjectClasses
	Kind        Kind
	Must        RequiredAttributeTypes
	May         PermittedAttributeTypes
	Extensions  Extensions
	bools       Boolean
	seq         int
}

/*
StructuralObjectClass is a type alias of *ObjectClass intended for use solely within instances of NameForm within its "OC" field.
*/
type StructuralObjectClass struct {
	*ObjectClass
}

/*
AttributeType conforms to the specifications of RFC4512 Section 4.1.2. Boolean values, e.g: 'OBSOLETE', are supported internally and are not explicit fields.
*/
type AttributeType struct {
	OID         OID
	Name        Name
	Description Description
	SuperType   SuperiorAttributeType
	Equality    Equality
	Ordering    Ordering
	Substring   Substring
	Syntax      *LDAPSyntax
	Usage       Usage
	Extensions  Extensions
	bools       Boolean
	mub         uint
	seq         int
}

/*
SuperiorAttributeType is a type alias of AttributeType meant solely for use within an instance of *AttributeType (via its "SUP" struct field) to define its "super type" when applicable.
*/
type SuperiorAttributeType struct {
	*AttributeType
}

/*
MatchingRule conforms to the specifications of RFC4512 Section 4.1.3. Boolean values, e.g: 'OBSOLETE', are supported internally and are not explicit fields.
*/
type MatchingRule struct {
	OID         OID
	Name        Name
	Description Description
	Syntax      *LDAPSyntax
	Extensions  Extensions
	bools       Boolean
	seq         int
}

/*
Equality circumscribes an embedded pointer to an instance of *MatchingRule.  This type alias is intended for use solely within instances of AttributeType via its "Equality" struct field.
*/
type Equality struct {
	*MatchingRule
}

/*
Ordering circumscribes an embedded pointer to an instance of *MatchingRule.  This type alias is intended for use solely within instances of AttributeType via its "Ordering" struct field.
*/
type Ordering struct {
	*MatchingRule
}

/*
Substring circumscribes an embedded pointer to an instance of *MatchingRule.  This type alias is intended for use solely within instances of AttributeType via its "Substring" struct field.
*/
type Substring struct {
	*MatchingRule
}

/*
MatchingRuleUse conforms to the specifications of RFC4512 Section 4.1.4. Boolean values, e.g: 'OBSOLETE', are supported internally and are not explicit fields.
*/
type MatchingRuleUse struct {
	OID         OID
	Name        Name
	Description Description
	Applies     Applies
	Extensions  Extensions
	bools       Boolean
	seq         int
}

/*
LDAPSyntax conforms to the specifications of RFC4512 Section 4.1.5. Internal Boolean support is managed only for a declaration of "human-readable" status, per RFC2252 Section 4.3.2.
*/
type LDAPSyntax struct {
	OID         OID
	Description Description
	Extensions  Extensions
	bools       Boolean
	seq         int
}

/*
DITContentRule conforms to the specifications of RFC4512 Section 4.1.6. Boolean values, e.g: 'OBSOLETE', are supported internally and are not explicit fields.

The OID value of this type MUST match the OID of a known (and catalogued) STRUCTURAL *ObjectClass instance.
*/
type DITContentRule struct {
	OID         OID
	Name        Name
	Description Description
	Aux         AuxiliaryObjectClasses
	Must        RequiredAttributeTypes
	May         PermittedAttributeTypes
	Not         ProhibitedAttributeTypes
	Extensions  Extensions
	bools       Boolean
	seq         int
}

/*
DITStructureRule conforms to the specifications of RFC4512 Section 4.1.7.1. Boolean values, e.g: 'OBSOLETE', are supported internally and are not explicit fields.
*/
type DITStructureRule struct {
	ID            RuleID
	Name          Name
	Description   Description
	Form          *NameForm
	SuperiorRules SuperiorDITStructureRules
	Extensions    Extensions
	bools         Boolean
	seq           int
}

/*
NameForm conforms to the specifications of RFC4512 Section 4.1.7.2. Boolean values, e.g: 'OBSOLETE', are supported internally and are not explicit fields.
*/
type NameForm struct {
	OID         OID
	Name        Name
	Description Description
	OC          StructuralObjectClass
	Must        RequiredAttributeTypes
	May         PermittedAttributeTypes
	Extensions  Extensions
	bools       Boolean
	seq         int
}

type genericList []interface{}

/*
Name contains slices of string names that may be assigned to any definition except instances of *LDAPSyntax.
*/
type Name genericList

/*
AuxiliaryObjectClasses represents slices of AUXILIARY instances of *ObjectClass that entries controlled by a given *DITContentRule may belong.
*/
type AuxiliaryObjectClasses genericList

/*
SuperiorObjectClasses represents slices of instances of *ObjectClass to which the bearer *ObjectClass is structurally subordinate (or from which it is derived).
*/
type SuperiorObjectClasses genericList

/*
RequiredAttributeTypes represents slices of instances of *AttributeType that are required by the bearing instance of *ObjectClass or *DITContentRule.
*/
type RequiredAttributeTypes genericList

/*
PermittedAttributeTypes represents slices of instances of *AttributeType that are allowed by the bearing instance of *ObjectClass or *DITContentRule.
*/
type PermittedAttributeTypes genericList

/*
ProhibitedAttributeTypes represents slices of instances of *AttributeType that are disallowed by the bearing instance of *DITContentRule.
*/
type ProhibitedAttributeTypes genericList

/*
Applies represents slices of instances of *MatchingRuleUse, each of which represent an *AttributeType instance that is associated with a particular *MatchingRule instance.
*/
type Applies genericList

/*
AssociatedDITStructureRules represents slices of instances of *DITStructureRule. Instances of this type will vary depending on the nature of the bearer.
*/
type SuperiorDITStructureRules genericList

/*
AttributeTypesManifest is a map structure associating a dot-delimited ASN.1 object identifier with an instance of *AttributeType.
*/
type AttributeTypesManifest map[OID]*AttributeType

/*
ObjectClassesManifest is a map structure associating a dot-delimited ASN.1 object identifier with an instance of *ObjectClass.
*/
type ObjectClassesManifest map[OID]*ObjectClass

/*
LDAPSyntaxesManifest is a map structure associating a dot-delimited ASN.1 object identifier with an instance of *LDAPSyntax.
*/
type LDAPSyntaxesManifest map[OID]*LDAPSyntax

/*
MatchingRulesManifest is a map structure associating a dot-delimited ASN.1 object identifier with an instance of *MatchingRule.
*/
type MatchingRulesManifest map[OID]*MatchingRule

/*
MatchingRuleUsesManifest is a map structure associating a dot-delimited ASN.1 object identifier with an instance of *MatchingRuleUse.
*/
type MatchingRuleUsesManifest map[OID]*MatchingRuleUse

/*
DITContentRulesManifest is a map structure associating a dot-delimited ASN.1 object identifier with an instance of *DITContentRule.
*/
type DITContentRulesManifest map[OID]*DITContentRule

/*
DITStructureRulesManifest is a map structure associating an unsigned integer (a rule number) with an instance of *DITStructureRule.
*/
type DITStructureRulesManifest map[RuleID]*DITStructureRule

/*
NameFormsManifest is a map structure associating a dot-delimited ASN.1 object identifier with an instance of *NameForm.  Each ASN.1 object identifier associated with an instance of *NameForm MUST be the OID of the referenced STRUCTURAL *ObjectClass instance which said *NameForm is intended to govern.
*/
type NameFormsManifest map[OID]*NameForm

/*
Extensions is a map structure associating a string extension name (e.g.: X-ORIGIN) with a non-zero string value.
*/
type Extensions map[string][]string

/*
Aliases is a map structure associating a string alias (or "macro") name with a dot-delimited ASN.1 object identifier.  Use of this is limited to scenarios involving LDAP implementations that support the aliasing of OIDs for certain schema definition elements.
*/
type AliasesManifest map[string]OID

/*
new.go deals with the instantiation of new objects.
*/

func newGenericList() genericList {
	return make(genericList, 0)
}

/*
newDefinition will parse the provided interface (x) into an instance of *definition, which is returned along with a success-indicative boolean value.  If values are present within the provided interface, they are preserved.
*/
func newDefinition(z interface{}, alm AliasesManifest) (def *definition, ok bool) {
	def = new(definition)
	def.alm = alm

	switch tv := z.(type) {
	case *AttributeType, *SuperiorAttributeType:
		if assert, is := tv.(*SuperiorAttributeType); is {
			def.typ = valueOf(assert).Elem().Type()
			def.meths = assert.methMap()
			def.labels = assert.labelMap()
		} else {
			def.typ = valueOf(tv.(*AttributeType)).Elem().Type()
			def.meths = tv.(*AttributeType).methMap()
			def.labels = tv.(*AttributeType).labelMap()
		}
	case *ObjectClass:
		def.typ = valueOf(tv).Elem().Type()
		def.meths = tv.methMap()
		def.labels = tv.labelMap()
	case *LDAPSyntax:
		def.typ = valueOf(tv).Elem().Type()
		def.meths = tv.methMap()
		def.labels = tv.labelMap()
	case *MatchingRule:
		def.typ = valueOf(tv).Elem().Type()
		def.meths = tv.methMap()
		def.labels = tv.labelMap()
	case *MatchingRuleUse:
		def.typ = valueOf(tv).Elem().Type()
		def.meths = tv.methMap()
		def.labels = tv.labelMap()
	case *DITContentRule:
		def.typ = valueOf(tv).Elem().Type()
		def.meths = tv.methMap()
		def.labels = tv.labelMap()
	case *DITStructureRule:
		def.typ = valueOf(tv).Elem().Type()
		def.meths = tv.methMap()
		def.labels = tv.labelMap()
	case *NameForm:
		def.typ = valueOf(tv).Elem().Type()
		def.meths = tv.methMap()
		def.labels = tv.labelMap()
	default:
		return
	}

	def.fields = make([]reflect.StructField, def.typ.NumField(), def.typ.NumField())
	def.values = make([]reflect.Value, len(def.fields), cap(def.fields))

	for i := 0; i < len(def.fields); i++ {
		def.fields[i] = def.typ.Field(i)
		def.values[i] = valueOf(z).Elem().Field(i)
	}
	ok = true

	return
}

/*
NewName returns an initialized instance of Name.
*/
func NewName(n ...string) (name Name) {
	name = make(Name, 0)
	if len(n) == 0 {
		return
	}
	name.Set(n...)

	return
}

/*
NewApplies returns a new instance of Applies, intended for assignment to an instance of *MatchingRuleUse.
*/
func NewApplies() Applies {
	return make(Applies, 0)
}

/*
NewRequiredAttributeTypes returns a new instance of RequiredAttributeTypes, intended for assignment to an instance of *ObjectClass, *NameForm, or *DITContentRule.
*/
func NewRequiredAttributeTypes() RequiredAttributeTypes {
	return make(RequiredAttributeTypes, 0)
}

/*
NewPermittedAttributeTypes returns a new instance of PermittedAttributeTypes, intended for assignment to an instance of *ObjectClass, *NameForm or *DITContentRule.
*/
func NewPermittedAttributeTypes() PermittedAttributeTypes {
	return make(PermittedAttributeTypes, 0)
}

/*
NewProhibitedAttributeTypes returns a new instance of ProhibitedAttributeTypes, intended for assignment to an instance of *NameForm or *DITContentRule.
*/
func NewProhibitedAttributeTypes() ProhibitedAttributeTypes {
	return make(ProhibitedAttributeTypes, 0)
}

/*
NewSuperiorDITStructureRules returns a new instance of SuperiorDITStructureRules, intended for assignment to an instance of *DITStructureRule.
*/
func NewSuperiorDITStructureRules() SuperiorDITStructureRules {
	return make(SuperiorDITStructureRules, 0)
}

/*
NewSuperiorObjectClasses returns a new instance of SuperiorObjectClasses, intended for assignment to an instance of *ObjectClass.
*/
func NewSuperiorObjectClasses() SuperiorObjectClasses {
	return make(SuperiorObjectClasses, 0)
}

/*
NewAuxiliaryObjectClasses returns a new instance of AuxiliaryObjectClasses, intended for assignment to an instance of *DITContentRule.
*/
func NewAuxiliaryObjectClasses() AuxiliaryObjectClasses {
	return make(AuxiliaryObjectClasses, 0)
}

/*
NewAttributeTypesManifest returns a new instance of AttributeTypesManifest, intended for population by the user and use as a lookup table when needed.
*/
func NewAttributeTypesManifest() AttributeTypesManifest {
	return make(AttributeTypesManifest, 0)
}

/*
NewObjectClassesManifest returns a new instance of ObjectClassesManifest, intended for population by the user and use as a lookup table when needed.
*/
func NewObjectClassesManifest() ObjectClassesManifest {
	return make(ObjectClassesManifest, 0)
}

/*
NewNameFormsManifest returns a new instance of NameFormsManifest, intended for population by the user and use as a lookup table when needed.
*/
func NewNameFormsManifest() NameFormsManifest {
	return make(NameFormsManifest, 0)
}

/*
NewDITContentRulesManifest returns a new instance of DITContentRulesManifest, intended for population by the user and use as a lookup table when needed.
*/
func NewDITContentRulesManifest() DITContentRulesManifest {
	return make(DITContentRulesManifest, 0)
}

/*
NewDITStructureRulesManifest returns a new instance of DITStructureRulesManifest, intended for population by the user and use as a lookup table when needed.
*/
func NewDITStructureRulesManifest() DITStructureRulesManifest {
	return make(DITStructureRulesManifest, 0)
}

/*
NewExtensions returns a new instance of Extensions, intended for assignment to any definition type.
*/
func NewExtensions() Extensions {
	return make(Extensions, 0)
}

/*
NewExtensions returns a new instance of AliasesManifest, intended for use in resolving OID aliases (or "macros").
*/
func NewAliasesManifest() AliasesManifest {
	return make(AliasesManifest, 0)
}

/*
NewMatchingRulesManifest returns a new instance of MatchingRulesManifest.
*/
func NewMatchingRulesManifest() MatchingRulesManifest {
	return make(MatchingRulesManifest, 0)
}

/*
NewMatchingRuleUsesManifest returns a new instance of MatchingRuleUsesManifest.
*/
func NewMatchingRuleUsesManifest() MatchingRuleUsesManifest {
	return make(MatchingRuleUsesManifest, 0)
}

/*
NewLDAPSyntaxesManifest returns a new instance of LDAPSyntaxesManifest.
*/
func NewLDAPSyntaxesManifest() LDAPSyntaxesManifest {
	return make(LDAPSyntaxesManifest, 0)
}

/*
NewSubschema returns a partially-initialized instance of *Subschema.
*/
func NewSubschema() *Subschema {
	s := new(Subschema)
	s.MRUM = NewMatchingRuleUsesManifest()

	return s
}

/*
NewRuleID returns a new instance of *RuleID, intended for assignment to an instance of *DITStructureRule.
*/
func NewRuleID(x interface{}) (rid RuleID) {
	switch tv := x.(type) {
	case int:
		if tv < 0 {
			return
		}
		x := uint(tv)
		rid = RuleID(x)
	case uint:
		rid = RuleID(tv)
	case string:
		if isDigit(tv) {
			if n, err := atoi(tv); err == nil && n > 0 {
				rid = RuleID(uint(n))
			}
		}
	}

	return
}

func newUsage(x interface{}) Usage {
	switch tv := x.(type) {
	case string:
		switch toLower(tv) {
		case toLower(DirectoryOperation.String()):
			return DirectoryOperation
		case toLower(DistributedOperation.String()):
			return DistributedOperation
		case toLower(DSAOperation.String()):
			return DSAOperation
		}
	case uint:
		switch tv {
		case 0x1:
			return DirectoryOperation
		case 0x2:
			return DistributedOperation
		case 0x3:
			return DSAOperation
		}
	case int:
		if tv >= 0 {
			return newUsage(uint(tv))
		}
	}

	return UserApplication
}

func newKind(x interface{}) Kind {
	switch tv := x.(type) {
	case Kind:
		return newKind(tv.String())
	case string:
		switch toLower(tv) {
		case toLower(Abstract.String()):
			return Abstract
		case toLower(Structural.String()):
			return Structural
		case toLower(Auxiliary.String()):
			return Auxiliary
		}
	case uint:
		switch tv {
		case 0x1:
			return Abstract
		case 0x2:
			return Structural
		case 0x3:
			return Auxiliary
		}
	case int:
		if tv >= 0 {
			return newKind(uint(tv))
		}
	}

	return badKind
}

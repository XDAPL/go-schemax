package schemax

import (
	"text/template"

	"github.com/JesseCoretta/go-stackage"
)

const (
	UserApplicationUsage      uint = iota // RFC 4512 § 4.1.2
	DirectoryOperationUsage               // RFC 4512 § 4.1.2
	DistributedOperationUsage             // RFC 4512 § 4.1.2
	DSAOperationUsage                     // RFC 4512 § 4.1.2
)

const (
	StructuralKind uint = iota // RFC 4512 § 2.4.2, 4.1.1
	AbstractKind               // RFC 4512 § 2.4.1, 4.1.1
	AuxiliaryKind              // RFC 4512 § 2.4.3, 4.1.1
)

// OIDList implements oidlist per § 4.1 of RFC 4512.  Instances
// of this type need not be handled by users directly.
type OIDList stackage.Stack

// RuleIDList implements ruleidlist per § 4.1.7.1 of RFC 4512.
// Instances of this type need not be handled by users directly.
type RuleIDList stackage.Stack

// QuotedDescriptorList implements qdescrlist per § 4.1 of RFC 4512.
// Instances of this type need not be handled by users directly.
type QuotedDescriptorList stackage.Stack

// QuotedStringList implements qdstringlist per § 4.1 of RFC 4512.
// Instances of this type need not be handled by users directly.
type QuotedStringList stackage.Stack

// Collection implements a common stackage.Stack type alias and is
// not an RFC 4512 construct.  Instances of this type need not be
// handled by users directly.
type Collection stackage.Stack

// Schema is a practical implementation of a 'subschemaSubentry'
// in that individual definitions are accessible and collectively
// define the elements available for use in populating a directory.
//
// See the following methods to access any desired [Definition]
// qualifier types:
//
//   - [Schema.LDAPSyntaxes]
//   - [Schema.MatchingRules]
//   - [Schema.AttributeTypes]
//   - [Schema.MatchingRuleUses]
//   - [Schema.ObjectClasses]
//   - [Schema.DITContentRules]
//   - [Schema.NameForms]
//   - [Schema.DITStructureRules]
type Schema Collection

type (
	ObjectClasses     Collection // RFC 4512 § 4.2.1
	AttributeTypes    Collection // RFC 4512 § 4.2.2
	MatchingRules     Collection // RFC 4512 § 4.2.3
	MatchingRuleUses  Collection // RFC 4512 § 4.2.4
	LDAPSyntaxes      Collection // RFC 4512 § 4.2.5
	DITContentRules   Collection // RFC 4512 § 4.2.6
	DITStructureRules Collection // RFC 4512 § 4.2.7
	NameForms         Collection // RFC 4512 § 4.2.8
)

/*
DefinitionName aliases the QuotedDescriptorList type.
*/
type DefinitionName QuotedDescriptorList

/*
AttributeType implements § 4.1.2 of RFC 4512.
*/
type AttributeType struct {
	*attributeType
}

type attributeType struct {
	OID        string
	Macro      []string
	Desc       string
	Name       DefinitionName
	Obsolete   bool
	SingleVal  bool
	Collective bool
	NoUserMod  bool
	SuperType  AttributeType
	Equality   MatchingRule
	Ordering   MatchingRule
	Substring  MatchingRule
	Syntax     LDAPSyntax
	MUB        uint
	Usage      uint
	Extensions Extensions

	schema Schema

	t *template.Template
	s string
}

/*
DITContentRule implements § 4.1.6 of RFC 4512.
*/
type DITContentRule struct {
	*dITContentRule
}

type dITContentRule struct {
	Desc       string
	Name       DefinitionName
	Obsolete   bool
	OID        ObjectClass
	Macro      []string
	Aux        ObjectClasses
	Must       AttributeTypes
	May        AttributeTypes
	Not        AttributeTypes
	Extensions Extensions

	schema Schema

	t *template.Template
	s string
}

/*
DITStructureRule implements § 4.1.7.1 of RFC 4512.
*/
type DITStructureRule struct {
	*dITStructureRule
}

type dITStructureRule struct {
	ID         uint
	Desc       string
	Name       DefinitionName
	Obsolete   bool
	Form       NameForm
	SuperRules DITStructureRules
	Extensions Extensions

	schema Schema

	t *template.Template
	s string
}

/*
Extensions implements extensions as defined in § 4.1 of RFC 4512.
*/
type Extensions struct {
	extensions
}

type extensions map[string]QuotedStringList

/*
LDAPSyntax implements § 4.1.5 of RFC 4512.
*/
type LDAPSyntax struct {
	*lDAPSyntax
}

type lDAPSyntax struct {
	OID        string
	Macro      []string
	Desc       string
	Extensions Extensions

	schema Schema

	t *template.Template
	s string
}

/*
MatchingRule implements § 4.1.3 of RFC 4512.
*/
type MatchingRule struct {
	*matchingRule
}

type matchingRule struct {
	OID        string
	Macro      []string
	Name       DefinitionName
	Desc       string
	Obsolete   bool
	Syntax     LDAPSyntax
	Extensions Extensions

	schema Schema

	t *template.Template
	s string
}

/*
MatchingRuleUse implements § 4.1.4 of RFC 4512.
*/
type MatchingRuleUse struct {
	*matchingRuleUse
}

type matchingRuleUse struct {
	OID        string
	Name       DefinitionName
	Desc       string
	Obsolete   bool
	Applies    AttributeTypes
	Extensions Extensions

	schema Schema

	t *template.Template
	s string
}

/*
NameForm implements § 4.1.7.2 of RFC 4512.
*/
type NameForm struct {
	*nameForm
}

type nameForm struct {
	OID        string
	Macro      []string
	Desc       string
	Name       DefinitionName
	Obsolete   bool
	Structural ObjectClass
	Must       AttributeTypes
	May        AttributeTypes
	Extensions Extensions

	schema Schema

	t *template.Template
	s string
}

/*
ObjectClass implements § 4.1.1 of RFC 4512.
*/
type ObjectClass struct {
	*objectClass
}

type objectClass struct {
	OID          string
	Macro        []string
	Desc         string
	Name         DefinitionName
	Obsolete     bool
	SuperClasses ObjectClasses
	Kind         uint
	Must         AttributeTypes
	May          AttributeTypes
	Extensions   Extensions

	schema Schema

	t *template.Template
	s string
}

/*
Counters is a simple struct type defined to store current number
of definition instances within an instance of [Schema].

Instances of this type are not thread-safe. Users should implement
another metrics system if thread-safety for counters is required.
*/
type Counters struct {
	LS int
	MR int
	AT int
	MU int
	OC int
	DC int
	NF int
	DS int
}

/*
Definition is an interface type used to allow basic interaction with
any of the following instance types:

  - [LDAPSyntax]
  - [MatchingRule]
  - [AttributeType]
  - [MatchingRuleUse]
  - [ObjectClass]
  - [DITContentRule]
  - [NameForm]
  - [DITStructureRule]

Not all methods extended through instances of these types are made
available through this interface.
*/
type Definition interface {
	// NumericOID returns the string representation of the ASN.1
	// numeric OID value of the instance.  If the underlying type
	// is a DITStructureRule, no value is returned as this type
	// of definition does not bear an OID.
	NumericOID() string

	// Name returns the first string NAME value present within
	// the underlying DefinitionName stack instance.  A zero
	// string is returned if no names were set, or if the given
	// type instance is LDAPSyntax, which does not bear a name.
	Name() string

	// Names returns the underlying instance of DefinitionName.
	// If executed upon an instance of LDAPSyntax, an empty
	// instance is returned, as LDAPSyntaxes do not bear names.
	Names() DefinitionName

	// IsZero returns a Boolean value indicative of nilness
	// with respect to the embedded type instance.
	IsZero() bool

	// String returns the complete string representation of the
	// underlying definition type per § 4.1.x of RFC 4512.
	String() string

	// Description returns the DESC field of the underlying
	// definition type, else a zero string if undefined.
	Description() string

	// IsIdentifiedAs returns a Boolean value indicative of
	// whether the input string represents an identifying
	// value for the underlying definition.
	//
	// Acceptable input values vary based on the underlying
	// type.  See the individual method notes for details.
	//
	// Case-folding of input values is not significant in
	// the matching process, regardless of underlying type.
	IsIdentifiedAs(string) bool

	// IsObsolete returns a Boolean value indicative of the
	// condition of definition obsolescence. Executing this
	// method upon an LDAPSyntax receiver will always return
	// false, as the condition of obsolescence does not
	// apply to this definition type.
	IsObsolete() bool

	// Type returns the string literal name for the receiver
	// instance. For example, if the receiver is an LDAPSyntax
	// instance, the return value is "ldapSyntax".
	Type() string

	// Extensions returns the underlying Extension instance value.
	// An empty Extensions instance is returned if no extensions
	// were set.
	Extensions() Extensions

	// setOID empties the definition Macro slice following completion
	// of the resolution process of a numeric OID during the parsing
	// process.
	setOID(string)

	// macro allows private access to the underlying Macro field
	// present within (nearly) all Definition qualifier types. This
	// is used for a low-cyclo means of resolving macros to actual
	// numeric OIDs during the parsing phase.
	macro() []string
}

/*
Definitions is an interface type used to allow basic interaction with
any of the following stack types:

  - [LDAPSyntaxes]
  - [MatchingRules]
  - [AttributeTypes]
  - [MatchingRuleUses]
  - [ObjectClasses]
  - [DITContentRules]
  - [NameForms]
  - [DITStructureRules]

Not all methods extended through instances of these types are made
available through this interface.
*/
type Definitions interface {
	// Len returns the integer length of the receiver instance,
	// indicating the number of definitions residing within.
	Len() int

	// IsZero returns a Boolean value indicative of nilness.
	IsZero() bool

	// Contains returns a Boolean value indicative of whether the
	// specified string value represents the RFC 4512 OID of a
	// Definition qualifier found within the receiver instance.
	Contains(string) bool

	// String returns the string representation of the underlying
	// stack instance.  Note that the manner in which the instance
	// was initialized will influence the resulting output.
	String() string

	// Type returns the string literal name for the receiver instance.
	// For example, if the receiver is LDAPSyntaxes, "ldapSyntaxes"
	// is the return value.  This is useful in case switching scenarios.
	Type() string

	// Push returns an error instance following an attempt to push
	// an instance of any into the receiver stack instance.  The
	// appropriate instance type depends on the nature of the
	// underlying stack instance.
	Push(any) error

	// cast is a private method used to reduce the cyclomatic penalties
	// normally incurred during the handling the eight (8) distinct RFC
	// 4512 definition stack types. cast is meant to allow this package
	// to access certain stackage.Stack methods that we have not wrapped
	// explicitly.
	cast() stackage.Stack
}

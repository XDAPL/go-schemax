package schemax

import (
	"text/template"

	"github.com/JesseCoretta/go-stackage"
)

const (
	UserApplicationsUsage     uint = iota // RFC 4512 § 4.1.2
	DirectoryOperationUsage               // RFC 4512 § 4.1.2
	DistributedOperationUsage             // RFC 4512 § 4.1.2
	DSAOperationUsage                     // RFC 4512 § 4.1.2
)

const (
	StructuralKind uint = iota // RFC 4512 § 2.4.2, 4.1.1
	AbstractKind               // RFC 4512 § 2.4.1, 4.1.1
	AuxiliaryKind              // RFC 4512 § 2.4.3, 4.1.1
)

/*
UseHangingIndents, when true, will result in a newline character
(ASCII #10) being inserted prior to each field of a [Definition]
in string form followed by four space (ASCII #32) characters.

Not all directory products are flexible regarding use of newlines
and hanging indents within schema definitions. Should difficulties
related to parsing occur, set this variable to false prior to the
initialization of an instance of [Schema] to keep [Definition]
strings confined to a single line.
*/
var UseHangingIndents bool

/*
OIDList implements oidlist per § 4.1 of RFC 4512.  Instances
of this type need not be handled by users directly.

	oidlist = oid *( WSP DOLLAR WSP oid )
*/
type OIDList stackage.Stack

/*
RuleIDList implements ruleidlist per § 4.1.7.1 of RFC 4512.
Instances of this type need not be handled by users directly.

	ruleidlist = ruleid *( SP ruleid )
*/
type RuleIDList stackage.Stack

/*
QuotedDescriptorList implements qdescrlist per § 4.1 of RFC 4512.
Instances of this type need not be handled by users directly.

	qdescrlist = [ qdescr *( SP qdescr ) ]
*/
type QuotedDescriptorList stackage.Stack

/*
QuotedStringList implements qdstringlist per § 4.1 of RFC 4512.
Instances of this type need not be handled by users directly.

	qdstringlist = [ qdstring *( SP qdstring ) ]
*/
type QuotedStringList stackage.Stack

/*
Collection implements a common stackage.Stack type alias and is
not an RFC 4512 construct.  Instances of this type need not be
handled by users directly.
*/
type Collection stackage.Stack

/*
Extensions implements extensions as defined in § 4.1 of RFC 4512.

	extensions = *( SP xstring SP qdstrings )
	xstring	   = "X" HYPHEN 1*( ALPHA / HYPHEN / USCORE )
*/
type Extensions stackage.Stack

/*
Extension is the singular form of [Extensions], and contains a
single string key (XString) and a QuotedStringList (Values).
*/
type Extension struct {
	XString string
	Values  QuotedStringList
}

/*
Schema is a practical implementation of a 'subschemaSubentry'
in that individual definitions are accessible and collectively
define the elements available for use in populating a directory.

See the following methods to access any desired [Definition]
qualifier types:

  - [Schema.LDAPSyntaxes]
  - [Schema.MatchingRules]
  - [Schema.AttributeTypes]
  - [Schema.MatchingRuleUses]
  - [Schema.ObjectClasses]
  - [Schema.DITContentRules]
  - [Schema.NameForms]
  - [Schema.DITStructureRules]
*/
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
DefinitionMap implements a convenient map-based Definition type.  Use
of this type is normally indicated in external processing scenarios,
such as templating.

Note that, due to the underlying map instance from which this type
extends, ordering of fields (e.g.: NAME vs. DESC) cannot be guaranteed.
*/
type DefinitionMap map[string][]string

/*
DefinitionMaps implements slices of DefinitionMap instances, collectively
representing an entire type-specific stack (e.g.: AttributeTypes).
*/
type DefinitionMaps []DefinitionMap

/*
Name aliases the QuotedDescriptorList type.
*/
type Name QuotedDescriptorList

/*
AttributeType implements § 4.1.2 of RFC 4512.

	AttributeTypeDescription = LPAREN WSP
	    numericoid                    ; object identifier
	    [ SP "NAME" SP qdescrs ]      ; short names (descriptors)
	    [ SP "DESC" SP qdstring ]     ; description
	    [ SP "OBSOLETE" ]             ; not active
	    [ SP "SUP" SP oid ]           ; supertype
	    [ SP "EQUALITY" SP oid ]      ; equality matching rule
	    [ SP "ORDERING" SP oid ]      ; ordering matching rule
	    [ SP "SUBSTR" SP oid ]        ; substrings matching rule
	    [ SP "SYNTAX" SP noidlen ]    ; value syntax
	    [ SP "SINGLE-VALUE" ]         ; single-value
	    [ SP "COLLECTIVE" ]           ; collective
	    [ SP "NO-USER-MODIFICATION" ] ; not user modifiable
	    [ SP "USAGE" SP usage ]       ; usage
	    extensions WSP RPAREN         ; extensions

	usage = "userApplications"     /  ; user
	        "directoryOperation"   /  ; directory operational
	        "distributedOperation" /  ; DSA-shared operational
	        "dSAOperation"            ; DSA-specific operational
*/
type AttributeType struct {
	*attributeType
}

type attributeType struct {
	OID        string
	Macro      []string
	Desc       string
	Name       Name
	Obsolete   bool
	Single     bool
	Collective bool
	Immutable  bool
	SuperType  AttributeType
	Equality   MatchingRule
	Ordering   MatchingRule
	Substring  MatchingRule
	Syntax     LDAPSyntax
	MUB        uint
	Usage      uint
	Extensions Extensions

	schema Schema

	t        *template.Template
	s        string
	stringer func() string
}

/*
DITContentRule implements § 4.1.6 of RFC 4512.

	DITContentRuleDescription = LPAREN WSP
	    numericoid                 ; object identifier
	    [ SP "NAME" SP qdescrs ]   ; short names (descriptors)
	    [ SP "DESC" SP qdstring ]  ; description
	    [ SP "OBSOLETE" ]          ; not active
	    [ SP "AUX" SP oids ]       ; auxiliary object classes
	    [ SP "MUST" SP oids ]      ; attribute types
	    [ SP "MAY" SP oids ]       ; attribute types
	    [ SP "NOT" SP oids ]       ; attribute types
	    extensions WSP RPAREN      ; extensions
*/
type DITContentRule struct {
	*dITContentRule
}

type dITContentRule struct {
	Desc       string
	Name       Name
	Obsolete   bool
	OID        ObjectClass
	Macro      []string
	Aux        ObjectClasses
	Must       AttributeTypes
	May        AttributeTypes
	Not        AttributeTypes
	Extensions Extensions

	schema Schema

	t        *template.Template
	s        string
	stringer func() string
}

/*
DITStructureRule implements § 4.1.7.1 of RFC 4512.

	DITStructureRuleDescription = LPAREN WSP
	    ruleid                     ; rule identifier
	    [ SP "NAME" SP qdescrs ]   ; short names (descriptors)
	    [ SP "DESC" SP qdstring ]  ; description
	    [ SP "OBSOLETE" ]          ; not active
	    SP "FORM" SP oid           ; NameForm
	    [ SP "SUP" SP ruleids ]    ; superior rules <NOTE: SEE ERRATA # 7896>
	    extensions WSP RPAREN      ; extensions

	ruleids = ruleid / ( LPAREN WSP ruleidlist WSP RPAREN )
	ruleidlist = ruleid *( SP ruleid )
	ruleid = number
*/
type DITStructureRule struct {
	*dITStructureRule
}

type dITStructureRule struct {
	ID         uint
	Desc       string
	Name       Name
	Obsolete   bool
	Form       NameForm
	SuperRules DITStructureRules
	Extensions Extensions

	schema Schema

	t        *template.Template
	s        string
	stringer func() string
}

//	type Extensions struct {
//		extensions
//	}
type extensions map[string]QuotedStringList

/*
LDAPSyntax implements § 4.1.5 of RFC 4512.

	SyntaxDescription = LPAREN WSP
	    numericoid                 ; object identifier
	    [ SP "DESC" SP qdstring ]  ; description
	    extensions WSP RPAREN      ; extensions
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

	t        *template.Template
	s        string
	stringer func() string
}

/*
MatchingRule implements § 4.1.3 of RFC 4512.

	MatchingRuleDescription = LPAREN WSP
	    numericoid                 ; object identifier
	    [ SP "NAME" SP qdescrs ]   ; short names (descriptors)
	    [ SP "DESC" SP qdstring ]  ; description
	    [ SP "OBSOLETE" ]          ; not active
	    SP "SYNTAX" SP numericoid  ; assertion syntax
	    extensions WSP RPAREN      ; extensions
*/
type MatchingRule struct {
	*matchingRule
}

type matchingRule struct {
	OID        string
	Macro      []string
	Name       Name
	Desc       string
	Obsolete   bool
	Syntax     LDAPSyntax
	Extensions Extensions

	schema Schema

	t        *template.Template
	s        string
	stringer func() string
}

/*
MatchingRuleUse implements § 4.1.4 of RFC 4512.

	MatchingRuleUseDescription = LPAREN WSP
	    numericoid                 ; object identifier
	    [ SP "NAME" SP qdescrs ]   ; short names (descriptors)
	    [ SP "DESC" SP qdstring ]  ; description
	    [ SP "OBSOLETE" ]          ; not active
	    SP "APPLIES" SP oids       ; attribute types
	    extensions WSP RPAREN      ; extensions
*/
type MatchingRuleUse struct {
	*matchingRuleUse
}

type matchingRuleUse struct {
	OID        string
	Name       Name
	Desc       string
	Obsolete   bool
	Applies    AttributeTypes
	Extensions Extensions

	schema Schema

	t        *template.Template
	s        string
	stringer func() string
}

/*
NameForm implements § 4.1.7.2 of RFC 4512.

	NameFormDescription = LPAREN WSP
	    numericoid                 ; object identifier
	    [ SP "NAME" SP qdescrs ]   ; short names (descriptors)
	    [ SP "DESC" SP qdstring ]  ; description
	    [ SP "OBSOLETE" ]          ; not active
	    SP "OC" SP oid             ; structural object class
	    SP "MUST" SP oids          ; attribute types
	    [ SP "MAY" SP oids ]       ; attribute types
	    extensions WSP RPAREN      ; extensions
*/
type NameForm struct {
	*nameForm
}

type nameForm struct {
	OID        string
	Macro      []string
	Desc       string
	Name       Name
	Obsolete   bool
	Structural ObjectClass
	Must       AttributeTypes
	May        AttributeTypes
	Extensions Extensions

	schema Schema

	t        *template.Template
	s        string
	stringer func() string
}

/*
ObjectClass implements § 4.1.1 of RFC 4512.

	ObjectClassDescription = LPAREN WSP
	    numericoid                 ; object identifier
	    [ SP "NAME" SP qdescrs ]   ; short names (descriptors)
	    [ SP "DESC" SP qdstring ]  ; description
	    [ SP "OBSOLETE" ]          ; not active
	    [ SP "SUP" SP oids ]       ; superior object classes
	    [ SP kind ]                ; kind of class
	    [ SP "MUST" SP oids ]      ; attribute types
	    [ SP "MAY" SP oids ]       ; attribute types
	    extensions WSP RPAREN

	kind = "ABSTRACT" / "STRUCTURAL" / "AUXILIARY"
*/
type ObjectClass struct {
	*objectClass
}

type objectClass struct {
	OID          string
	Macro        []string
	Desc         string
	Name         Name
	Obsolete     bool
	SuperClasses ObjectClasses
	Kind         uint
	Must         AttributeTypes
	May          AttributeTypes
	Extensions   Extensions

	schema Schema

	t        *template.Template
	s        string
	stringer func() string
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
Inventory is a type alias of map[string][]string, and is used to
provide a simple manifest of all members of a Collective derivative,
such as LDAPSyntaxes.  This can be useful during activities such as
templating.

Unlike the DefinitionMap type, this type is only used to manifest
the most basic details of a collection of definitions, namely the
numerical and textual identifiers present.

Keys represent the numerical identifier for a definition, whether a
numeric OID or integer rule ID.  In the case of dITStructureRule
definitions, the rule ID -- an unsigned integer -- is used. In all
other cases, a numeric OID is used.

Values represent the NAME or DESC by which the definition is known.

Note that DESC is only used in the case of LDAPSyntax instances, and
NAME is used for all definition types *except* LDAPSyntax.
*/
type Inventory map[string][]string

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
	// the underlying Name stack instance.  A zero
	// string is returned if no names were set, or if the given
	// type instance is LDAPSyntax, which does not bear a name.
	Name() string

	// Names returns the underlying instance of Name.
	// If executed upon an instance of LDAPSyntax, an empty
	// instance is returned, as LDAPSyntaxes do not bear names.
	Names() Name

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

	// Map returns a DefinitionMap instance based upon the
	// the contents and state of the receiver instance.
	Map() DefinitionMap

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

	// Maps returns slices of DefinitionMap instances, each of which
	// are expressions of actual Definition qualifiers found within
	// the receiver instance.
	Maps() DefinitionMaps

	// Inventory returns an instance of Inventory containing
	// the fundamental numerical and textual identifiers for
	// the contents of any given definition collection. Note
	// that the semantics for representation within instances
	// of this type vary among definition types.
	Inventory() Inventory

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

func hindent() (x string) {
	x = string(rune(10)) + `    `
	if !UseHangingIndents {
		x = ` `
	}

	return
}

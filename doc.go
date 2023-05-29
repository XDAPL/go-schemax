/*
Package schemax provides methods, types and bidirectional marshaling functionality intended for use in the area of X.501/LDAP directory schema abstraction.

Types provided by this package are created based on the precepts of RFC2252 and RFC4512, as are the associated parsing techniques (e.g.: 'qdescrs', etc.).  A variety of methods and functions exist to make handling this data simpler.

Abstract

LDAP directories contain and serve data in a hierarchical manner. The syntax and evaluation of this data is governed by a schema, which itself is hierarchical in nature.

The nature of this package's operation is highly referential.  Objects are referenced via pointers, and these objects can inhabit multiple other multi-valued types. For example, *AttributeType instances are stored within a type-specific slices called Collections.  Each *AttributeType that exists can be referenced by other *AttributeType instances (in a scenario where "super typing" is in effect), or by other *ObjectClass instances via their own PermittedAttributeTypes (May) and RequiredAttributeTypes (Must) list types.  Literal "copies" of single objects are never made.  References are always through pointers.

Intended Audience

This package is primarily intended for any architect, developer or analyst looking to do one or more of the following:

• Parse (Marshal) textual LDAP schema definitions into objects

• Unmarshal objects into textual LDAP schema definitions

• Work with data models and structures that may represent or contain some facet of a schema

Parsing

This package aims to provide a fast, reliable, standards-compliant parsing routine for LDAP schema definitions into discrete, useful objects.  Parsing of raw values during the Marshal operation is conducted without the use of the regexp package, nor any facility based on regular expressions. Instead, precise byte-for-byte parsing/truncation of raw definitions is conducted, during which time known flags (or labels) are identified for specialized handling (e.g.: NAME).

For successful parsing, each definition must be limited to a single line when an entire schema file is streamed line-by-line. Future releases of this package will relax this requirement.  However, when marshaling single definitions, multi-line definitions are supported. Line feeds are removed outright, and recurring WHSP characters (spaces/tabs) are reduced to single spaces between fields.

Population Order

When POPULATING various collection types (e.g.: slices of definitions), the following "order of operations" MUST be honored:

1. LDAPSyntaxes

2. MatchingRules

3. AttributeTypes

4. MatchingRuleUses

5. ObjectClasses

6. DITContentRules

7. NameForms

8. DITStructureRules

Individually, collections containing the above elements should also be populated in order of referential superiority.  For example, all independent instances should be populated before those that depend upon them, such as in the cases of sub-types, sub-classes and sub-rules.

An obvious real-world example of this is for the 'name' attribute (per RFC4519), which is a super-type of several other key attribute types.  In such a case, those definitions that depend (or are based) upon the 'name' attribute WILL NOT MARSHAL until 'name' has been marshaled itself.

This logic applies no matter how the definitions are being received (or "read"), and applies whether or not the all-encompassing Subschema type is used.  In short: MIND YOUR ORDERING.

Iterating Collections

Iterating a collection type, such as AttributeTypeCollection instances, MUST be done using the Len() collection method, e.g.:

  for i := 0; i < attrs.Len(); i++ {
	attr := attrs.Index(i)
        ... do stuff ...
  }

Iteration will always returns collection members in FIFO ordering (First-In/First-Out).

Standard Definitions

Within subdirectories of this package are popular Go implementations of standard Attribute Types, Object Classes, Matching Rules and LDAP Syntaxes from RFCs that are recognized (almost) universally.  This includes (but is not limited to) content from RFC4512, RFC4519 and RFC2307.  Included in each subdirectory is an unmodified text copy of the Internet-Draft from which the relevant definitions originate.

The user is advised that some LDAP implementations have certain attribute types and object classes "built-in" and are not sourced from a schema file like others (rather they likely are compiled-in to the product).

This varies between implementations and, as such, inconsistencies may arise for someone using this product across various directory products. One size absolutely does not fit all.  In such a case, an attempt to marshal a schema file may fail due to unsatisfied super-type or super-class(es) dependencies.  To mitigate this, the user must somehow provide the lacking definitions, either by themselves or using one of the subdirectory packages.

OID Macros

Also known as OID "aliases", macros allow a succinct expression of an OID prefix by way of a text identifier.  As a real-world example, RFC2307 uses the alias "nisSchema" to describe the OID prefix of "1.3.6.1.1.1".

This package supports the registration and use of such aliases within the Macros map type.  Note that this is an all-or-nothing mechanism. Understand that if a non-nil Macros instance is detected, and unregistered aliases are encountered during a parsing run, normal operations will be impacted. As such, users are advised to anticipate any aliases needed in advance or to abolish their use altogether.

OID aliasing supports both dot (.) and colon (:) runes for delimitation, thus 'nisSchema.1.1' and 'nisSchema:1.1' are acceptable.

Custom Unmarshalers

By default, definitions are unmarshaled as single-line string values. This package supports the creation of user-authored closure functions to allow customization of this behavior. As a practical example, users may opt to write a function to convert definitions to individual CSV rows, or perhaps JSON objects -- literally anything, so long as it is a string.

Each definition type comes with a default (optional) UnmarshalFunc instance. This, or any custom function, can be invoked by-name as follows:

  myfunc := UnmarshalerFuncName		// Must conform to the DefinitionUnmarshaler signature!
  def.SetUnmarshaler(myfunc)		// assign the function by name to your definition
  raw, err := schemax.Unmarshal(def)	// unmarshal
  ...

In the above example, the definition would unmarshal as a standard RFC4512 schema definition, but with linebreaks and indenting added. This must be done for EVERY definition being unmarshaled.

User-authored functions MUST honor the following function signature (see the DefinitionUnmarshaler type for details).

  func(any) (string, error)

The structure of a given user-authored unmarshaler function will vary, but should generally reflect the example below.

 // Example attributeType unmarshaler
 func myCustomAttributeUnmarshaler(x any) (def string, err error) {

	// We'll use this variable to store the
	// value once we verify it's copacetic.

        var r *AttributeType

	// As you can clearly see in the signature, the user
	// will provide a value as an interface. We'll need
	// to type-assert whatever they provide to ensure
	// its the exact type we are prepared to handle.

        switch tv := x.(type) {
        case *AttributeType:
                if tv.IsZero() {
                        err = fmt.Errorf("%T is nil", tv)
                        return
                }
                r = tv
        default:
                err = fmt.Errorf("Bad type for unmarshal (%T)", tv)
                return
        }

	//
	// <<Your custom attributeType-handling code would go here>>
	//

	return
 }

 // ... later in your code ...
 //
 // This can be applied to a single definition
 // -OR- a collection of definitions!
 obj.SetUnmarshaler(myCustomAttributeUnmarshaler)

Map Unmarshaler

For added convenience, each definition includes a Map() method that returns a map[string][]string instance containing the effective contents of said definition. This is useful in situations where the user is more interested in simple access to string values in fields, as opposed to the complicated traversal of pointer instances that may exist within a definition.

  defmap := def.Map()
  if value, exists := defmap[`SYNTAX`]; exists {
	val := value[0]
	fmt.Printf("Syntax OID is: %s\n", val)
  }

Naturally, ordering of fields is lost due to use of a map in this fashion. It is up to the consuming application to ensure correct ordering of fields as described in RFC4512 section 4.1, wherever applicable.

Extended Information

Each definition supports the optional assignment of a []byte value containing "extended information", which can literally be anything (pre-rendered HTML, or even Graphviz content to name a few potential use cases). This package imposes no restrictions of any kind regarding the nature or length of the assigned byte slice.

The main purpose of this functionality is to allow the user to annotate information that goes well beyond the terse DESC field value normally present within definitions.

  info := []byte(`<html>...</html>`)
  def.SetInfo(info)
  ...
  fmt.Printf("%s\n", string(def.Info()))
*/
package schemax

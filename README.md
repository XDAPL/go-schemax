# go-schemax

[![RFC 4512](https://img.shields.io/badge/RFC-4512-blue?cacheSeconds=500000)](https://datatracker.ietf.org/doc/html/rfc4512) [![RFC4517](https://img.shields.io/badge/RFC-4517-blue?cacheSeconds=500000)](https://datatracker.ietf.org/doc/html/rfc4517) [![RFC4519](https://img.shields.io/badge/RFC-4519-blue?cacheSeconds=500000)](https://datatracker.ietf.org/doc/html/rfc4519) [![RFC4523](https://img.shields.io/badge/RFC-4523-blue?cacheSeconds=500000)](https://datatracker.ietf.org/doc/html/rfc4523) [![RFC4524](https://img.shields.io/badge/RFC-4524-blue?cacheSeconds=500000)](https://datatracker.ietf.org/doc/html/rfc4524) [![RFC4530](https://img.shields.io/badge/RFC-4530-blue?cacheSeconds=500000)](https://datatracker.ietf.org/doc/html/rfc4530) [![RFC2307](https://img.shields.io/badge/RFC-2307-blue?cacheSeconds=500000)](https://datatracker.ietf.org/doc/html/rfc2307) [![RFC2798](https://img.shields.io/badge/RFC-2798-blue?cacheSeconds=500000)](https://datatracker.ietf.org/doc/html/rfc2798) [![RFC3045](https://img.shields.io/badge/RFC-3045-blue?cacheSeconds=500000)](https://datatracker.ietf.org/doc/html/rfc3045) [![RFC3671](https://img.shields.io/badge/RFC-3671-blue?cacheSeconds=500000)](https://datatracker.ietf.org/doc/html/rfc3671) [![RFC3672](https://img.shields.io/badge/RFC-3672-blue?cacheSeconds=500000)](https://datatracker.ietf.org/doc/html/rfc3672)

Package schemax implements a powerful [RFC4512](https://www.rfc-editor.org/rfc/rfc4512.txt) parser.

Requires Go version 1.21 or higher.

[![Go Report Card](https://goreportcard.com/badge/JesseCoretta/go-schemax)](https://goreportcard.com/report/github.com/JesseCoretta/go-schemax) [![Reference](https://pkg.go.dev/badge/github.com/JesseCoretta/go-schemax.svg)](https://pkg.go.dev/github.com/JesseCoretta/go-schemax) [![License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat)](https://github.com/JesseCoretta/go-schemax/blob/main/LICENSE) [![Help Animals](https://img.shields.io/badge/help_animals-gray?label=%F0%9F%90%BE%20%F0%9F%98%BC%20%F0%9F%90%B6&labelColor=yellow)](https://github.com/JesseCoretta/JesseCoretta/blob/main/DONATIONS.md)

## License

The schemax package is available under the terms of the MIT license.  For further details, see the LICENSE file within the root of the repository.

## Schema Manifest

Definitions from select RFCs -- denoted by the RFC badges above -- are available for initialization within instances of `Schema`.  These are known as "built-in" definitions within the context of this package.

Users may introduce definitions from other sources -- official or not -- using the various `Parse<...>` methods extended through various types within this package.

## Releases

Two (2) releases are available for end-users:

| Version | Notes |
| :----- | :--- |
| 1.1.6 | Legacy, custom parser |
| >= 2.0.0 | Current, ANTLR parser |

## History of schemax

The goal of schemax has always been to provide a reliable parsing subsystem for schemata and `attributeDescription` values that allows transformation to usable Go objects. This goal will not change.

The original design of schemax (version < 2.0.0) involved a custom-made parser. While this design performed remarkably well for years, it was not without its shortcomings. 

The newly released build of schemax involves the import of an ANTLR4-based [RFC 4512](https://www.rfc-editor.org/rfc/rfc4512.txt) lexer/parser solution. This is made possible using a newly released "sister" package -- [`go-antlr4512`](https://github.com/JesseCoretta/go-antlr4512) -- which handles all of the low-level ANTLR actions such as tokenization.

Therefore, the new build of schemax is of a simpler fundamental design thanks to offloading the bulk of the parser to another package. This also keeps all code-grading penalties (due to ANTLR's characteristically high cyclomatic factors) confined elsewhere, and allows schemax to focus on extending the slick features users have come to expect.

## The Parser

The (ANTLR) parsing subsystem imported by the aforementioned sister package is flexible in terms of the following:

  - Presence of header, footer and right-pane Bash comments surrounding a given definition is acceptable
    - Note that comments are entirely discarded by ANTLR and are unavailable through the imported sister package
  - Definition prefixing allows variations of the standard [RFC 4512](https://www.rfc-editor.org/rfc/rfc4512.txt) "labels" during file and directory parsing
    - "`attributeTypes`", "`attributeType`" and other variations are permitted for `AttributeType` definitions
  - Definition delimitation -- using colon (`:`), equals (`=`) or whitespace (` `, `\t`) of any sensible combination -- are permitted for the purpose of separating a definition prefix (label) from its definition statement.
    - "attributeTypes: ...", "attributeType=...", "attributeType ..." are valid expressions
  - Multiple files are joined using ASCII #10 (newline) during directory parsing
    - Users need not worry about adding a trailing newline to each file to be read; schemax will do this for you if needed

## File and Directory Readers

The (legacy) "v1" release branches of schemax did not offer a robust file and directory parsing solution, rather it focused on the byte representations of a given definition and the tokens derived therein, leaving it to the end-user to devise a delivery method.

The "v2" release branches introduce proper `ParseFile` and `ParseDirectory` methods that greatly simplify use of this package in the midst of an established schema "library".  For example:

```
func main() {
	r := NewSchema()

	// Let's parse a directory into our
	// receiver instance of Schema (r).
	if err := r.ParseDirectory(`/ds/etc/schema`); err != nil {
		fmt.Println(err)
		return
	}

	// Check our definition counters
	fmt.Printf("%s", r.Counters())
	// Output:
	// LS: 67
	// MR: 44
	// AT: 131
	// MU: 29
	// OC: 39
	// DC: 1
	// NF: 1
	// DS: 1
}
```

Though the `ParseFile` function operates identically to the above-demonstrated `ParseDirectory` function, it is important to order the respective files and directories according to any applicable dependencies.  In other words, if "fileB.schema" requires definitions from "fileA.schema", "fileA.schema" must be parsed first.

Sub-directories encountered shall be traversed indefinitely. The effective name of a given directory is not significant.

Files encountered through directory traversal shall only be read and parsed IF the extension is ".schema".  This prevents other files -- such as text or `README.md` files -- from interfering with the parsing process needlessly.

An eligible schema file may contain one definition, or many. The effective name of an eligible schema file **is significant**, unlike directories.  Each schema file must be named in a manner that fosters the correct ordering of dependent definitions -- **_whether or not subdirectories are involved_**. To offer a real-world example, the 389DS/Netscape schema directory deployed during a typical installation is defined and governed in a similar manner.

The general rule-of-thumb is suggests that if the `ls -l` Bash command _consistently_ lists the indicated schema files in correct order, and assuming those files contain properly ordered and well-formed definitions, the parsing process should work nicely.

## The Schema Itself

The Schema type defined within this package is a [`stackage.Stack`](https://pkg.go.dev/github.com/JesseCoretta/go-stackage#Stack) derivative type.  An instance of a Schema can manifest in any of the following manners:

  - As an empty (unpopulated) Schema, initialized by way of the `NewEmptySchema` function
  - As a basic (minimally populated) Schema, initialized by way of the `NewBasicSchema` function
  - As a complete (fully populated) Schema, initialized by way of the `NewSchema` function

There are certain scenarios which call for one of the above initialization procedures:

  - An empty schema is ideal for LDAP professionals, and allows for the creation of a Schema of particularly narrow focus for R&D, testing or product development
  - A basic schema resembles the foundational (starting) Schema context observed in most directory server products, in that it comes "pre-loaded" with official LDAPSyntax and MatchingRule definitions -- but few to no AttributeTypes -- making it a most suitable empty canvas upon which a new Schema may be devised from scratch
  - A full schema is the most obvious choice for "Quick Start" scenarios, in that a Schema is produced containing a very large portion of the standard AttributeType and ObjectClass definitions used in the wild by most (if not all) directory products

Regardless of the content present, a given Schema is capable of storing definitions from all eight (8) RFC 4512 "categories".  These are known as "collections", and are stored in nested [`stackage.Stack`](https://pkg.go.dev/github.com/JesseCoretta/go-stackage#Stack) derivative types, accessed using any of the following methods:

  - `Schema.LDAPSyntaxes`
  - `Schema.MatchingRules`
  - `Schema.AttributeTypes`
  - `Schema.MatchingRuleUses`
  - `Schema.ObjectClasses`
  - `Schema.DITContentRules`
  - `Schema.NameForms`
  - `Schema.DITStructureRules`

Definition instances produced by way of parsing -- namely using one of the `Parse<Type>` functions -- will automatically gain internal access to the Schema instance in which it is stored.

However, definitions produced manually by way of the various `Set<Item>` methods extended through types defined within this package will require manual execution of the `SetSchema` method, using the intended Schema instance as the input argument.  Ideally this should occur early in the definition composition.

In either case, this internal reference is used for seamless verification of any reference, such as an LDAPSyntax, when introduced to a given type instance.  This ensures only well-formed definitions are created.

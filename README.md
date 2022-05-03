# go-schemax

[![GoDoc](https://godoc.org/github.com/JesseCoretta/go-schemax?status.svg)](https://godoc.org/github.com/JesseCoretta/go-schemax)

Abstract directory schema constructs and methods based on RFC4512 Section 4.1.

This is an early release, and should not be used for mission critical applications at this time.

Contributions and bug reports are most welcome.

## Overview

Package schemax provides methods, types and bidirectional marshaling functionality intended for use in the area of X.501/LDAP directory schema abstraction.

Types provided by this package are created based on the precepts of RFC2252 and RFC4512, as are the associated parsing techniques (e.g.: 'qdescrs', etc.). A variety of methods and functions exist to make handling this data simpler.

## Abstract

LDAP directories contain and serve data in a hierarchical manner. The syntax and evaluation of this data is governed by a schema, which itself is hierarchical in nature.

The nature of this package's operation is highly referential. Objects are referenced via pointers, and these objects can inhabit multiple other multi-valued types. For example, `*AttributeType` instances are stored within a type-specific map type called a manifest. Each `*AttributeType` that exists can be referenced by other `*AttributeType` instances (in a scenario where "super typing" is in effect), or by other `*ObjectClass` instances via their own PermittedAttributeTypes (May) and RequiredAttributeTypes (Must) list types. Literal "copies" of single objects are never made. References are always through pointers.

## Intended Audience

This package is primarily intended for any architect, developer or analyst looking to do one or more of the following:

 - Parse (Marshal) textual LDAP schema definitions into objects
 - Unmarshal objects into textual LDAP schema definitions
 - Work with data models and structures that may represent or contain some facet of a schema

## Parsing

This package aims to provide a fast, reliable, standards-compliant parsing routine for LDAP schema definitions into discrete, useful objects. Parsing of raw values during the Marshal operation is conducted without the use of the regexp package, nor any facility based on regular expressions. Instead, precise byte-for-byte parsing/truncation of raw definitions is conducted, during which time known flags (or labels) are identified for specialized handling (e.g.: NAME).

For successful parsing, each definition must be limited to a single line when an entire schema file is streamed line-by-line. Future releases of this package will relax this requirement. However, when marshaling single definitions, multi-line definitions are supported. Line feeds are removed outright, and recurring WHSP characters (spaces/tabs) are reduced to single spaces between fields.

## Population Order

When POPULATING various collection types (e.g.: slices of definitions), the following "order of operations" MUST be honored:

 1. LDAPSyntaxes
 2. MatchingRules
 3. AttributeTypes
 4. MatchingRuleUses
 5. ObjectClasses
 6. DITContentRules
 7. NameForms
 8. DITStructureRules

Individually, collections containing the above elements should also be populated in order of referential superiority. For example, all independent instances should be populated before those that depend upon them, such as in the cases of sub-types, sub-classes and sub-rules.

An obvious real-world example of this is for the 'name' attribute (per RFC4519), which is a super-type of several other key attribute types. In such a case, those definitions that depend (or are based) upon the 'name' attribute WILL NOT MARSHAL until 'name' has been marshaled itself.

This logic applies no matter how the definitions are being received (or "read"), and applies whether or not the all-encompassing Subschema type is used. In short: MIND YOUR ORDERING.

## Iterating Collections

Iterating a collection type, such as AttributeTypeCollection instances, MUST be done using the Len() collection method, e.g.:

  for i := 0; i < attrs.Len(); i++ {
	attr := attrs.Index(i)
        ... do stuff ...
  }
Iteration will always returns collection members in FIFO ordering (First-In/First-Out).

## Standard Definitions

Within subdirectories of this package are popular Go implementations of standard Attribute Types, Object Classes, Matching Rules and LDAP Syntaxes from RFCs that are recognized (almost) universally. This includes (but is not limited to) content from RFC4512, RFC4519 and RFC2307. Included in each subdirectory is an unmodified text copy of the Internet-Draft from which the relevant definitions originate.

The user is advised that some LDAP implementations have certain attribute types and object classes "built-in" and are not sourced from a schema file like others (rather they likely are compiled-in to the product).

This varies between implementations and, as such, inconsistencies may arise for someone using this product across various directory products. One size absolutely does not fit all. In such a case, an attempt to marshal a schema file may fail due to unsatisfied super-type or super-class(es) dependencies. To mitigate this, the user must somehow provide the lacking definitions, either by themselves or using one of the subdirectory packages.

## OID Macros

Also known as OID "aliases", macros allow a succinct expression of an OID prefix by way of a text identifier. As a real-world example, RFC2307 uses the alias "nisSchema" to describe the OID prefix of "1.3.6.1.1.1".

This package supports the registration and use of such aliases within the Macros map type. Note that this is an all-or-nothing mechanism. Understand that if a non-nil Macros instance is detected, and unregistered aliases are encountered during a parsing run, normal operations will be impacted. As such, users are advised to anticipate any aliases needed in advance or to abolish their use altogether.

OID aliasing supports both dot (.) and colon (:) runes for delimitation, thus 'nisSchema.1.1' and 'nisSchema:1.1' are acceptable.

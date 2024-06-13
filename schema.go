package schemax

/*
schema.go centralizes all schema operations within a single construct.
*/

const (
	ldapSyntaxesIndex      int = iota // 0
	matchingRulesIndex                // 1
	attributeTypesIndex               // 2
	matchingRuleUsesIndex             // 3
	objectClassesIndex                // 4
	dITContentRulesIndex              // 5
	nameFormsIndex                    // 6
	dITStructureRulesIndex            // 7
)

/*
NewSchema returns a new instance of [Schema] containing ALL
package-included definitions. See the internal directory
contents for a complete manifest.
*/
func NewSchema() (r Schema) {
	r = initSchema()
	var err error

	for _, funk := range []func() error{
		r.loadSyntaxes,
		r.loadMatchingRules,
		r.loadAttributeTypes,
		r.loadObjectClasses,
	} {
		if err = funk(); err != nil {
			break
		}
	}

	if err == nil {
		err = r.updateMatchingRuleUses(r.AttributeTypes())
	}

	// panic if ANY errors
	if err != nil {
		panic(err)
	}

	return
}

/*
NewBasicSchema initializes and returns an instance of [Schema].

The Schema instance shall only contain the [LDAPSyntax] and
[MatchingRule] definitions from the following RFCs:

  - RFC 2307
  - RFC 4517
  - RFC 4523
  - RFC 4530

This function produces a [Schema] that best resembles the schema
subsystem found in most DSA products today, in that [LDAPSyntax]
and [MatchingRule] definitions generally are not loaded by the
end user, however they are pre-loaded to allow immediate creation
of other (dependent) definition types, namely [AttributeType]
instances.
*/
func NewBasicSchema() (r Schema) {
	r = initSchema()
	var err error

	for _, funk := range []func() error{
		r.loadSyntaxes,
		r.loadMatchingRules,
	} {
		if err = funk(); err != nil {
			break
		}
	}

	// panic if ANY errors
	if err != nil {
		panic(err)
	}

	return
}

/*
NewEmptySchema initializes and returns an instance of [Schema] completely
initialized but devoid of any definitions whatsoever.

This function is intended for advanced users building a very specialized
[Schema] instance.
*/
func NewEmptySchema() (s Schema) {
	s = initSchema()
	return
}

/*
initSchema returns an initialized instance of Schema.
*/
func initSchema() Schema {
	return Schema(stackageList().
		SetID(`cn=schema`).
		SetCategory(`subschemaSubentry`).
		SetDelimiter(rune(10)).
		SetAuxiliary(map[string]any{
			`macros`:  make(map[string]string, 0),
			`options`: newOpts(),
		}).
		Mutex().
		Push(NewLDAPSyntaxes(), // 0
			NewMatchingRules(),      // 1
			NewAttributeTypes(),     // 2
			NewMatchingRuleUses(),   // 3
			NewObjectClasses(),      // 4
			NewDITContentRules(),    // 5
			NewNameForms(),          // 6
			NewDITStructureRules())) // 7
}

/*
SetMacro returns an error following an attempt to associate x with y.

x must be an RFC 4512-compliant descriptor, and y must be a legal numeric
OID.
*/
func (r Schema) SetMacro(x, y string) (err error) {
	if !isDescriptor(x) || !isNumericOID(y) {
		err = mkerr("Descriptor and/or numeric OID are zero length")
		return
	}

	_m := r.cast().Auxiliary()[`macros`]
	m, _ := _m.(map[string]string)
	m[x] = y

	return
}

/*
Options returns the underlying [Options] instance found within the
receiver instance.
*/
func (r Schema) Options() Options {
	_m := r.cast().Auxiliary()[`options`]
	m, _ := _m.(Options)
	return m
}

/*
GetMacro returns value y if associated with x.  A Boolean value, found,
is returned indicative of a match.

Case is not significant in the matching process.
*/
func (r Schema) GetMacro(x string) (y string, found bool) {
	_m := r.cast().Auxiliary()[`macros`]
	m, _ := _m.(map[string]string)

	for k, v := range m {
		if eq(x, k) {
			y = v
			found = true
		}
	}

	return
}

/*
GetMacroName returns value x if associated with numeric OID y. A
Boolean value, found, is returned indicative of a match.

Case is not applicable in the numeric OID matching process.
*/
func (r Schema) GetMacroName(y string) (x string, found bool) {
	_m := r.cast().Auxiliary()[`macros`]
	m, _ := _m.(map[string]string)

	for k, v := range m {
		if eq(y, v) {
			found = true
			x = k
			break
		}
	}

	return
}

/*
Override will attempt to override a separate incarnation of itself using
the [Definition] instance provided.

This is specifically to allow support for overriding certain [Definition]
instances, such as an [ObjectClass] to overcome inherent flaws in its
design.

The most common use case for this method is to allow users to override the
"groupOfNames" [ObjectClass] to remove the "member" [AttributeType] from the
MUST clause and, instead, place it in the MAY clause thereby allowing use of
memberless groups within a DIT.

This method SHOULD NOT be used in a cavalier manner; modifying official
[Definition] instances can wreck havoc on a directory and should only be
performed by skilled directory professionals and only when absolutely
necessary.

When overriding a [DITStructureRule] instance, a match is performed against
the respective [DITStructureRule.RuleID] values.  All other [Definition]
types are  matched using their respective numeric OIDs.  All replacement
[Definition] instances are subject to compliancy checks.

This is a fluent method.
*/
func (r Schema) Override(x Definition) Schema {
	if !r.Options().Positive(AllowOverride) {
		return r
	}

	switch x.Type() {
	case `ldapSyntax`:
		orig := r.LDAPSyntaxes().Get(x.NumericOID())
		orig.replace(x.(LDAPSyntax))
	case `matchingRule`:
		orig := r.MatchingRules().Get(x.NumericOID())
		orig.replace(x.(MatchingRule))
	case `matchingRuleUse`:
		orig := r.MatchingRuleUses().Get(x.NumericOID())
		orig.replace(x.(MatchingRuleUse))
	case `attributeType`:
		orig := r.AttributeTypes().Get(x.NumericOID())
		orig.replace(x.(AttributeType))
	case `objectClass`:
		orig := r.ObjectClasses().Get(x.NumericOID())
		orig.replace(x.(ObjectClass))
	case `nameForm`:
		orig := r.NameForms().Get(x.NumericOID())
		orig.replace(x.(NameForm))
	case `dITContentRule`:
		orig := r.DITContentRules().Get(x.NumericOID())
		orig.replace(x.(DITContentRule))
	case `dITStructureRule`:
		orig := r.DITStructureRules().Get(x.(DITStructureRule).ID())
		orig.replace(x.(DITStructureRule))
	}

	return r
}

/*
IsZero returns a Boolean value indicative of a nil receiver instance.
*/
func (r Schema) IsZero() bool {
	return r.cast().IsZero()
}

/*
ParseFile returns an error following an attempt to parse file. Only
files ending in ".schema" will be considered, however submission of
non-qualifying files shall not produce an error.
*/
func (r Schema) ParseFile(file string) (err error) {
	s := new4512Schema()
	if err = s.ParseFile(file); err != nil {
		return
	}

	// begin second phase
	err = r.incorporate(s)

	return
}

/*
ParseDirectory returns an error following an attempt to parse the
directory found at dir. Sub-directories encountered are traversed
indefinitely.  Files encountered will only be read if their name
ends in ".schema", at which point their contents are read into
bytes, processed using ANTLR and written to the receiver instance.
*/
func (r Schema) ParseDirectory(dir string) (err error) {
	s := new4512Schema()
	if err = s.ParseDirectory(dir); err != nil {
		return
	}

	// begin second phase
	err = r.incorporate(s)

	return
}

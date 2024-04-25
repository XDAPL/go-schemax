package schemax

/*
NewMatchingRules initializes a new [Collection] instance and
casts it as an [MatchingRules] instance.
*/
func NewMatchingRules() MatchingRules {
	r := MatchingRules(newCollection(`matchingRules`))
	r.cast().SetPushPolicy(r.canPush)

	return r
}

/*
NewMatchingRule initializes and returns a new instance of [MatchingRule].
*/
func NewMatchingRule() MatchingRule {
	return MatchingRule{newMatchingRule()}
}

func newMatchingRule() *matchingRule {
	return &matchingRule{
		Name:       NewName(),
		Extensions: NewExtensions(),
	}
}

/*
NumericOID returns the string representation of the numeric OID
held by the receiver instance.
*/
func (r MatchingRule) NumericOID() (noid string) {
	if !r.IsZero() {
		noid = r.matchingRule.OID
	}

	return
}

/*
SetNumericOID allows the manual assignment of a numeric OID to the
receiver instance if the following are all true:

  - The input id value is a syntactically valid numeric OID
  - The receiver does not already possess a numeric OID

This is a fluent method.
*/
func (r *MatchingRule) SetNumericOID(id string) *MatchingRule {
        if r.IsZero() {
                r.matchingRule = newMatchingRule()
        }

        if isNumericOID(id) {
                if len(r.matchingRule.OID) == 0 {
                        r.matchingRule.OID = id
                }
        }

        return r
}

/*
OID returns the string representation of an OID -- which is either a
numeric OID or descriptor -- that is held by the receiver instance.
*/
func (r MatchingRule) OID() (oid string) {
	if !r.IsZero() {
		oid = r.NumericOID() // default
		if r.matchingRule.Name.len() > 0 {
			oid = r.matchingRule.Name.index(0)
		}
	}

	return
}

/*
Syntax returns the string numeric OID value associated with
the underlying [LDAPSyntax] instance.
*/
func (r MatchingRule) Syntax() (desc string) {
	if !r.IsZero() {
		if !r.matchingRule.Syntax.IsZero() {
			desc = r.matchingRule.Syntax.NumericOID()
		}
	}

	return
}

/*
SetSyntax assigns x to the receiver instance as an instance of [LDAPSyntax].

This is a fluent method.
*/
func (r *MatchingRule) SetSyntax(x any) *MatchingRule {
	if r.IsZero() {
		r.matchingRule = newMatchingRule()
	}

	var syn LDAPSyntax
	switch tv := x.(type) {
	case string:
		if sch := r.schema(); !sch.IsZero() {
			syn = sch.LDAPSyntaxes().get(tv)
		}
	case LDAPSyntax:
		syn = tv
	}

	if !syn.IsZero() {
		r.matchingRule.Syntax = syn
	}

	return r
}

/*
SetSchema assigns an instance of [Schema] to the receiver instance.  This allows
internal verification of certain actions without the need for user input of
an instance of [Schema] manually at each juncture.

Note that the underlying [Schema] instance is automatically set when creating
instances of this type by way of parsing.

This is a fluent method.
*/
func (r *MatchingRule) SetSchema(schema Schema) *MatchingRule {
	if r.IsZero() {
		r.matchingRule = newMatchingRule()
	}

	r.matchingRule.schema = schema

	return r
}

func (r MatchingRule) schema() (s Schema) {
	if !r.IsZero() {
		s = r.matchingRule.schema
	}

	return
}

/*
SetDescription parses desc into the underlying Desc field within the
receiver instance.  Although a RFC 4512-compliant QuotedString is
required, the outer single-quotes need not be specified literally.

This is a fluent method.
*/
func (r *MatchingRule) SetDescription(desc string) *MatchingRule {
	if len(desc) < 3 {
		return r
	}

	if r.matchingRule == nil {
		r.matchingRule = new(matchingRule)
	}

	if !(rune(desc[0]) == rune(39) && rune(desc[len(desc)-1]) == rune(39)) {
		if !r.IsZero() {
			r.matchingRule.Desc = desc
		}
	}

	return r
}

/*
Description returns the underlying (optional) descriptive text
assigned to the receiver instance.
*/
func (r MatchingRule) Description() (desc string) {
	if !r.IsZero() {
		desc = r.matchingRule.Desc
	}
	return
}

/*
IsIdentifiedAs returns a Boolean value indicative of whether id matches
either the numericOID or descriptor of the receiver instance.  Case is
not significant in the matching process.
*/
func (r MatchingRule) IsIdentifiedAs(id string) (ident bool) {
	if !r.IsZero() {
		ident = id == r.NumericOID() || r.matchingRule.Name.contains(id)
	}

	return
}

/*
IsZero returns a Boolean value indicative of nilness of the
receiver instance.
*/
func (r MatchingRule) IsZero() bool {
	return r.matchingRule == nil
}

/*
MatchingRules returns the [MatchingRules] instance from within
the receiver instance.
*/
func (r Schema) MatchingRules() (mrs MatchingRules) {
	slice, _ := r.cast().Index(matchingRulesIndex)
	mrs, _ = slice.(MatchingRules)
	return
}

/*
Maps returns slices of [DefinitionMap] instances.
*/
func (r MatchingRules) Maps() (defs DefinitionMaps) {
	defs = make(DefinitionMaps, r.Len())
	for i := 0; i < r.Len(); i++ {
		defs[i] = r.Index(i).Map()
	}

	return
}

/*
Map marshals the receiver instance into an instance of
DefinitionMap.
*/
func (r MatchingRule) Map() (def DefinitionMap) {
	if r.IsZero() {
		return
	}

	def = make(DefinitionMap, 0)
	def[`NUMERICOID`] = []string{r.NumericOID()}
	def[`NAME`] = r.Names().List()
	def[`DESC`] = []string{r.Description()}
	def[`OBSOLETE`] = []string{bool2str(r.IsObsolete())}
	def[`SYNTAX`] = []string{r.Syntax()}
	def[`RAW`] = []string{r.String()}

	// copy our extensions from receiver r
	// into destination def.
	def = mapTransferExtensions(r, def)

	// Clean up any empty fields
	def.clean()

	return
}

/*
IsObsolete returns a Boolean value indicative of definition obsolescence.
*/
func (r MatchingRule) IsObsolete() (o bool) {
	if !r.IsZero() {
		o = r.matchingRule.Obsolete
	}

	return
}

/*
SetObsolete sets the receiver instance to OBSOLETE if not already set.

Obsolescence cannot be unset.

This is a fluent method.
*/
func (r *MatchingRule) SetObsolete() *MatchingRule {
	if !r.IsZero() {
		if !r.IsObsolete() {
			r.matchingRule.Obsolete = true
		}
	}

	return r
}

/*
SetName assigns the provided names to the receiver instance.

Name instances must conform to RFC 4512 descriptor format but
need not be quoted.
*/
func (r *MatchingRule) SetName(x ...string) *MatchingRule {
	if len(x) == 0 {
		return r
	}

	if r.IsZero() {
		r.matchingRule = newMatchingRule()
	}

	for i := 0; i < len(x); i++ {
		r.matchingRule.Name.Push(x[i])
	}

	return r
}

/*
Names returns the underlying instance of [Name] from within
the receiver.
*/
func (r MatchingRule) Names() (names Name) {
	return r.matchingRule.Name
}

/*
Name returns the string form of the principal name of the receiver instance, if set.
*/
func (r MatchingRule) Name() (id string) {
	if !r.IsZero() {
		id = r.matchingRule.Name.index(0)
	}

	return
}

/*
SetExtension assigns key x to value xstrs within the receiver's underlying
[Extensions] instance.

This is a fluent method.
*/
func (r *MatchingRule) SetExtension(x string, xstrs ...string) *MatchingRule {
	if r.IsZero() {
		r.matchingRule = newMatchingRule()
	}

	r.Extensions().Set(x, xstrs...)

	return r
}

/*
Extensions returns the [Extensions] instance -- if set -- within
the receiver.
*/
func (r MatchingRule) Extensions() (e Extensions) {
	if !r.IsZero() {
		e = r.matchingRule.Extensions
	}

	return
}

func (r MatchingRule) macro() (m []string) {
	if !r.IsZero() {
		m = r.matchingRule.Macro
	}

	return
}

func (r MatchingRule) setOID(x string) {
	if !r.IsZero() {
		r.matchingRule.OID = x
	}
}

/*
LoadMatchingRules will load all package-included [MatchingRule] definitions
into the receiver instance.
*/
func (r Schema) LoadMatchingRules() Schema {
	_ = r.loadMatchingRules()
	return r
}

/*
loadMatchingRules returns an error following an attempt to load
all matchingRule definitions found within this package into the
receiver instance.
*/
func (r Schema) loadMatchingRules() (err error) {
	if !r.IsZero() {
		for _, funk := range []func() error{
			r.loadRFC2307MatchingRules,
			r.loadRFC4517MatchingRules,
			r.loadRFC4523MatchingRules,
			r.loadRFC4530MatchingRules,
		} {
			if err = funk(); err != nil {
				break
			}
		}
	}

	return
}

/*
LoadRFC2307MatchingRules returns an error following an attempt to load
all RFC 2307 [MatchingRule] slices into the receiver instance.
*/
func (r Schema) LoadRFC2307MatchingRules() Schema {
	_ = r.loadRFC2307MatchingRules()
	return r
}

func (r Schema) loadRFC2307MatchingRules() (err error) {
	for k, v := range rfc2307Macros {
		r.SetMacro(k, v)
	}

	for i := 0; i < len(rfc2307MatchingRules) && err == nil; i++ {
		mr := rfc2307MatchingRules[i]
		err = r.ParseMatchingRule(string(mr))
	}

	return
}

/*
LoadRFC4517MatchingRules returns an error following an attempt to load
all RFC 4517 [MatchingRule] slices into the receiver instance.
*/
func (r Schema) LoadRFC4517MatchingRules() Schema {
	_ = r.loadRFC4517MatchingRules()
	return r
}

func (r Schema) loadRFC4517MatchingRules() (err error) {
	for i := 0; i < len(rfc4517MatchingRules) && err == nil; i++ {
		mr := rfc4517MatchingRules[i]
		err = r.ParseMatchingRule(string(mr))
	}

	return
}

/*
LoadRFC4523MatchingRules returns an error following an attempt to load
all RFC 4523 [MatchingRule] slices into the receiver instance.
*/
func (r Schema) LoadRFC4523MatchingRules() Schema {
	_ = r.loadRFC4523MatchingRules()
	return r
}

func (r Schema) loadRFC4523MatchingRules() (err error) {
	for i := 0; i < len(rfc4523MatchingRules) && err == nil; i++ {
		mr := rfc4523MatchingRules[i]
		err = r.ParseMatchingRule(string(mr))
	}

	return
}

/*
LoadRFC4530MatchingRules returns an error following an attempt to load
all RFC 4530 [MatchingRule] slices into the receiver instance.
*/
func (r Schema) LoadRFC4530MatchingRules() Schema {
	_ = r.loadRFC4530MatchingRules()
	return r
}

func (r Schema) loadRFC4530MatchingRules() (err error) {
	for i := 0; i < len(rfc4530MatchingRules) && err == nil; i++ {
		mr := rfc4530MatchingRules[i]
		err = r.ParseMatchingRule(string(mr))
	}

	return
}

func (r *matchingRule) prepareString() (err error) {
	buf := newBuf()
	r.t = newTemplate(`matchingRule`).
		Funcs(funcMap(map[string]any{
			`Syntax`:       func() string { return r.Syntax.NumericOID() },
			`ExtensionSet`: r.Extensions.tmplFunc,
		}))
	if r.t, err = r.t.Parse(matchingRuleTmpl); err == nil {
		if err = r.t.Execute(buf, struct {
			Definition *matchingRule
			HIndent    string
		}{
			Definition: r,
			HIndent:    hindent(),
		}); err == nil {
			r.s = buf.String()
		}
	}

	return
}

// stackage closure func - do not exec directly.
// cyclo=6
func (r MatchingRules) canPush(x ...any) (err error) {
	if len(x) == 0 {
		return
	}

	for i := 0; i < len(x) && err == nil; i++ {
		instance := x[i]
		if mr, ok := instance.(MatchingRule); !ok || mr.IsZero() {
			err = errorf("Type assertion for %T has failed", instance)
		} else if tst := r.get(mr.NumericOID()); !tst.IsZero() {
			err = errorf("%T %s not unique", mr, mr.NumericOID())
		}
	}

	return
}

/*
Len returns the current integer length of the receiver instance.
*/
func (r MatchingRules) Len() int {
	return r.len()
}

func (r MatchingRules) len() int {
	return r.cast().Len()
}

// cyclo=0
func (r MatchingRules) String() string {
	return r.cast().String()
}

/*                                                                      
String is a stringer method that returns the string representation      
of the receiver instance.                                               
*/                                                                      
func (r MatchingRule) String() (mr string) {
        if !r.IsZero() {                                                
                if r.stringer != nil {                                  
                        mr = r.stringer()                               
                } else {                                                
                        if len(r.matchingRule.s) == 0 {                
                                var err error                           
                                if err = r.matchingRule.prepareString(); err != nil {
                                        return                          
                                }                                       
                        }                                               
                                                                        
                        mr = r.matchingRule.s                          
                }                                                       
        }                                                               
                                                                        
        return                                                          
}

/*
IsZero returns a Boolean value indicative of nilness of the
receiver instance.
*/
func (r MatchingRules) IsZero() bool {
	return r.cast().IsZero()
}

/*
Index returns the instance of [MatchingRule] found within the
receiver stack instance at index N.  If no instance is found at
the index specified, a zero [MatchingRule] instance is returned.
*/
func (r MatchingRules) Index(idx int) MatchingRule {
	return r.index(idx)
}

func (r MatchingRules) index(idx int) (mr MatchingRule) {
	slice, found := r.cast().Index(idx)
	if found {
		if _mr, ok := slice.(MatchingRule); ok {
			mr = _mr
		}
	}

	return
}

/*
Push returns an error following an attempt to push a MatchingRule
into the receiver stack instance.
*/
func (r MatchingRules) Push(mr any) error {
	return r.push(mr)
}

func (r MatchingRules) push(mr any) (err error) {
	if mr == nil {
		err = errorf("%T instance is nil; cannot append to %T", mr, r)
		return
	}

	r.cast().Push(mr)

	return
}

/*
Contains calls [MatchingRules.Get] to return a Boolean value indicative of
a successful, non-zero retrieval of an [MatchingRule] instance -- matching
the provided id -- from within the receiver stack instance.
*/
func (r MatchingRules) Contains(id string) bool {
	return r.contains(id)
}

func (r MatchingRules) contains(id string) bool {
	return !r.get(id).IsZero()
}

/*
Type returns the string literal "matchingRule".
*/
func (r MatchingRule) Type() string {
	return `matchingRule`
}

/*
Type returns the string literal "matchingRules".
*/
func (r MatchingRules) Type() string {
	return `matchingRules`
}

/*
Get returns an instance of [MatchingRule] based upon a search for id within
the receiver stack instance.

The return instance, if not nil, was retrieved based upon a textual match of
the principal identifier of an [MatchingRule] and the provided id.

The return instance is nil if no match was made.

Case is not significant in the matching process.
*/
func (r MatchingRules) Get(id string) MatchingRule {
	return r.get(id)
}

func (r MatchingRules) get(id string) (mr MatchingRule) {
	for i := 0; i < r.len() && mr.IsZero(); i++ {
		if _mr := r.index(i); !_mr.IsZero() {
			if _mr.IsIdentifiedAs(id) {
				mr = _mr
			}
		}
	}

	return
}

package schemax

import (
	"internal/rfc2307"
	"internal/rfc4517"
	"internal/rfc4523"
	"internal/rfc4530"

	antlr4512 "github.com/JesseCoretta/go-rfc4512-antlr"
)

var (
	rfc2307MatchingRules rfc2307.MatchingRuleDefinitions = rfc2307.AllMatchingRules
	rfc4517MatchingRules rfc4517.MatchingRuleDefinitions = rfc4517.AllMatchingRules
	rfc4523MatchingRules rfc4523.MatchingRuleDefinitions = rfc4523.AllMatchingRules
	rfc4530MatchingRules rfc4530.MatchingRuleDefinitions = rfc4530.AllMatchingRules
)

func (r *Schema) ParseMatchingRule(raw string) (err error) {
	var i antlr4512.Instance
	if i, err = antlr4512.ParseInstance(raw); err == nil {
		var m MatchingRule
		if m, err = r.processMatchingRule(i.P.MatchingRuleDescription()); err != nil {
			return
		}

		err = r.MatchingRules().push(m)
	}

	return
}

/*
NewMatchingRules initializes a new Collection instance and
casts it as an MatchingRules instance.
*/
func NewMatchingRules() MatchingRules {
	r := MatchingRules(newCollection(``))
	r.cast().SetPushPolicy(r.canPush)

	return r
}

func (r MatchingRule) Syntax() (desc string) {
	if !r.IsZero() {
		if !r.matchingRule.Syntax.IsZero() {
			desc = r.matchingRule.Syntax.NumericOID()
		}
	}

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
Name returns the string form of the principal name of the receiver instance, if set.
*/
func (r MatchingRule) Name() (id string) {
	if !r.IsZero() {
		id = r.matchingRule.Name.index(0)
	}

	return
}

/*
Names returns the underlying instance of DefinitionName from within
the receiver.
*/
func (r MatchingRule) Names() (names DefinitionName) {
	return r.matchingRule.Name
}

/*
Extensions returns the Extensions instance -- if set -- within
the receiver.
*/
func (r MatchingRule) Extensions() (e Extensions) {
	if !r.IsZero() {
		e = r.matchingRule.Extensions
	}

	return
}

/*
String is a stringer method that returns the string representation
of the receiver instance.
*/
func (r MatchingRule) String() (mr string) {
	if !r.IsZero() {
		mr = r.matchingRule.s
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

func (r *matchingRule) prepareString() (err error) {
	buf := newBuf()
	r.t = newTemplate(`matchingRule`).
		Funcs(funcMap(map[string]any{
			`ExtensionSet`: r.Extensions.tmplFunc,
		}))
	if r.t, err = r.t.Parse(matchingRuleTmpl); err == nil {
		if err = r.t.Execute(buf, r); err == nil {
			r.s = buf.String()
		}
	}

	return
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
IsZero returns a Boolean value indicative of nilness of the
receiver instance.
*/
func (r MatchingRule) IsZero() bool {
	return r.matchingRule == nil
}

/*
List returns a map[string][]string instance which represents the current
inventory of matching rule instances within the receiver.  The keys are
numeric OIDs, while the values are zero (0) or more string slices, each
representing a name by which the definition is known.
*/
func (r MatchingRules) List() (list map[string][]string) {
	list = make(map[string][]string, 0)
	for i := 0; i < r.len(); i++ {
		def := r.index(i)
		list[def.NumericOID()] = def.Names().List()

	}

	return
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

// cyclo=0
func (r MatchingRules) contains(id string) bool {
	return !r.get(id).IsZero()
}

// cyclo=6
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

func (r Schema) processMatchingRule(ctx antlr4512.IMatchingRuleDescriptionContext) (mr MatchingRule, err error) {

	_mr := new(matchingRule)
	_mr.schema = r

	for k, ct := 0, ctx.GetChildCount(); k < ct && err == nil; k++ {
		switch tv := ctx.GetChild(k).(type) {
		case *antlr4512.OpenParenContext, *antlr4512.CloseParenContext:
			err = parenContext(tv)
		case *antlr4512.NumericOIDOrMacroContext:
			_mr.OID, _mr.Macro, err = numOIDContext(tv)
		case *antlr4512.DefinitionNameContext:
			_mr.Name, err = nameContext(tv)
		case *antlr4512.DefinitionDescriptionContext:
			_mr.Desc, err = descContext(tv)
		case *antlr4512.DefinitionObsoleteContext:
			_mr.Obsolete = true
		case *antlr4512.DefinitionSyntaxContext:
			_mr.Syntax, err = _mr.syntaxContext(tv)
		case *antlr4512.DefinitionExtensionsContext:
			_mr.Extensions, err = extContext(tv)
		default:
			env := MatchingRule{_mr}
			err = isErrImpl(env.Type(), env.OID(), tv)
		}
	}

	if err == nil {
		r.resolveByMacro(MatchingRule{_mr})

		if err = _mr.check(); err == nil {
			if err = _mr.prepareString(); err == nil {
				_mr.t = nil
				mr = MatchingRule{_mr}
			}
		}
	}

	return
}

func (r *matchingRule) syntaxContext(ctx *antlr4512.DefinitionSyntaxContext) (s LDAPSyntax, err error) {
	var syn string
	if syn, err = syntaxContext(ctx); err == nil {
		if s = r.schema.LDAPSyntaxes().get(syn); s.IsZero() {
			err = errorf("%T.Syntax (%s) not found; cannot process", r, syn)
		}
	}

	return
}

func (r *matchingRule) check() (err error) {
	if r == nil {
		err = errorf("%T is nil", r)
		return
	}

	if len(r.OID) == 0 {
		err = errorf("%T lacks an OID", r)
		return
	}

	if r.Syntax.IsZero() {
		err = errorf("%T.Syntax is nil", r)
	}

	return
}

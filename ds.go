package schemax

import (
	antlr4512 "github.com/JesseCoretta/go-rfc4512-antlr"
)

/*
ParseDITStructureRule parses an individual textual structure rule (raw)
and returns an error instance.

When no error occurs, the newly formed [DITStructureRule] instance -- based
on the parsed contents of raw -- is added to the receiver [DITStructureRules]
slice instance.
*/
func (r *Schema) ParseDITStructureRule(raw string) (err error) {
	i, err := parseI(raw)
	if err == nil {
		var dsr DITStructureRule
		dsr, err = r.processDITStructureRule(
			i.P.DITStructureRuleDescription())
		if err == nil {
			err = r.DITStructureRules().push(dsr)
		}
	}

	return
}

/*
NewDITStructureRules initializes a new [Collection] instance and casts
it as an [DITStructureRules] instance.
*/
func NewDITStructureRules() DITStructureRules {
	r := DITStructureRules(newCollection(``))
	r.cast().SetPushPolicy(r.canPush)

	return r
}

/*
NewDITStructureRuleIDList initializes a new [RuleIDList] instance and casts it
as a [DITStructureRules] instance.

This is mainly used to define a series of superior [DITStructureRule] instances
specified by a subordinate instance of [DITStructureRule].
*/
func NewDITStructureRuleIDList() DITStructureRules {
	r := DITStructureRules(newRuleIDList(``))
	r.cast().
		SetPushPolicy(r.canPush).
		SetPresentationPolicy(r.iDsStringer)

	return r
}

/*
NumericOID returns an empty string, as [DITStructureRule] definitions
do not bear numeric OIDs.  This method exists only to satisfy Go
interface requirements.
*/
func (r DITStructureRule) NumericOID() string { return `` }

/*
IsObsolete returns a Boolean value indicative of definition obsolescence.
*/
func (r DITStructureRule) IsObsolete() (o bool) {
	if !r.IsZero() {
		o = r.dITStructureRule.Obsolete
	}

	return
}

/*
Type returns the string literal "dITStructureRule".
*/
func (r DITStructureRule) Type() string {
	return `dITStructureRule`
}

/*
Type returns the string literal "dITStructureRules".
*/
func (r DITStructureRules) Type() string {
	return `dITStructureRules`
}

/*
SuperRules returns a [DITStructureRules] containing zero (0) or more
superior [DITStructureRule] instances from which the receiver extends.
*/
func (r DITStructureRule) SuperRules() (sup DITStructureRules) {
	if !r.IsZero() {
		sup = r.dITStructureRule.SuperRules
	}

	return
}

/*
ID returns the string representation of the underlying rule ID held by
the receiver instance.
*/
func (r DITStructureRule) ID() (id string) {
	if !r.IsZero() {
		_id := sprintf("%d", r.RuleID()) // default
		if id = r.Name(); len(id) == 0 {
			id = _id
		}
	}

	return
}

/*
IsIdentifiedAs returns a Boolean value indicative of whether id matches
either the numericOID or descriptor of the receiver instance.  Case is
not significant in the matching process.
*/
func (r DITStructureRule) IsIdentifiedAs(id string) (ident bool) {
	if !r.IsZero() {
		ident = id == sprintf("%d", r.RuleID()) ||
			r.dITStructureRule.Name.contains(id)
	}

	return
}

/*
Form returns the underlying instance of [NameForm] set within the
receiver. If unset, a zero instance is returned.
*/
func (r DITStructureRule) Form() (nf NameForm) {
	if !r.IsZero() {
		nf = r.dITStructureRule.Form
	}

	return
}

/*
Extensions returns the [Extensions] instance -- if set -- within
the receiver.
*/
func (r DITStructureRule) Extensions() (e Extensions) {
	if !r.IsZero() {
		e = r.dITStructureRule.Extensions
	}

	return
}

/*
RuleID returns the unsigned integer identifier held by the
receiver instance.
*/
func (r DITStructureRule) RuleID() (id uint) {
	if !r.IsZero() {
		id = r.dITStructureRule.ID
	}

	return
}

/*
Name returns the string form of the principal name of the receiver instance, if set.
*/
func (r DITStructureRule) Name() (id string) {
	if !r.IsZero() {
		id = r.dITStructureRule.Name.index(0)
	}

	return
}

/*
Names returns the underlying instance of [DefinitionName] from
within the receiver.
*/
func (r DITStructureRule) Names() (names DefinitionName) {
	return r.dITStructureRule.Name
}

/*
Description returns the underlying (optional) descriptive text
assigned to the receiver instance.
*/
func (r DITStructureRule) Description() (desc string) {
	if !r.IsZero() {
		desc = r.dITStructureRule.Desc
	}

	return
}

/*
String is a stringer method that returns the string representation
of the receiver instance.
*/
func (r DITStructureRule) String() (dsr string) {
	if !r.IsZero() {
		dsr = r.dITStructureRule.s
	}

	return
}

/*
List returns a map[uint][]string instance which represents the current
inventory of [DITStructureRule] instances within the receiver.  The keys
are numeric OIDs, while the values are zero (0) or more string slices,
each representing a name by which the definition is known.
*/
func (r DITStructureRules) List() (list map[uint][]string) {
	list = make(map[uint][]string, 0)
	for i := 0; i < r.len(); i++ {
		def := r.index(i)
		list[def.RuleID()] = def.Names().List()

	}

	return
}

func (r DITStructureRule) setOID(_ string) {}
func (r DITStructureRule) macro() []string { return []string{} }

// stackage closure func - do not exec directly (use String method)
func (r DITStructureRules) iDsStringer(_ ...any) (present string) {
	var _present []string
	for i := 0; i < r.len(); i++ {
		_present = append(_present, r.index(i).ID())
	}

	switch len(_present) {
	case 0:
		break
	case 1:
		present = _present[0]
	default:
		padchar := string(rune(32))
		if !r.cast().IsPadded() {
			padchar = ``
		}

		joined := join(_present, sprintf("%s %s",
			padchar, padchar))

		present = sprintf("(%s%s%s)",
			padchar, joined, padchar)
	}

	return
}

// stackage closure func - do not exec directly.
func (r DITStructureRules) canPush(x ...any) (err error) {
	if len(x) == 0 {
		return
	}

	for i := 0; i < len(x) && err == nil; i++ {
		instance := x[i]
		err = errorf("Type assertion for %T has failed", instance)
		if ds, ok := instance.(DITStructureRule); ok && !ds.IsZero() {
			err = errorf("%T %s not unique", ds, ds.NumericOID())
			if tst := r.get(ds.NumericOID()); tst.IsZero() {
				err = nil
			}
		}
	}

	return
}

/*
Len returns the current integer length of the receiver instance.
*/
func (r DITStructureRules) Len() int {
	return r.len()
}

func (r DITStructureRules) len() int {
	return r.cast().Len()
}

/*
String is a stringer method that returns the string representation
of the receiver instance.
*/
func (r DITStructureRules) String() string {
	return r.cast().String()
}

/*
IsZero returns a Boolean value indicative of nilness of the
receiver instance.
*/
func (r DITStructureRules) IsZero() bool {
	return r.cast().IsZero()
}

/*
Index returns the instance of [DITStructureRule] found within the
receiver stack instance at index N.  If no instance is found at
the index specified, a zero [DITStructureRule] instance is returned.
*/
func (r DITStructureRules) Index(idx int) DITStructureRule {
	return r.index(idx)
}

func (r DITStructureRules) index(idx int) (ds DITStructureRule) {
	if slice, found := r.cast().Index(idx); found {
		if _ds, ok := slice.(DITStructureRule); ok {
			ds = _ds
		}
	}

	return
}

/*
Push returns an error following an attempt to push a [DITStructureRule]
into the receiver stack instance.
*/
func (r DITStructureRules) Push(ds any) error {
	return r.push(ds)
}

func (r DITStructureRules) push(ds any) (err error) {
	err = errorf("%T instance is nil; cannot append to %T", ds, r)
	if ds != nil {
		r.cast().Push(ds)
		err = nil
	}

	return
}

func (r DITStructureRules) contains(id string) bool {
	return !r.get(id).IsZero()
}

func (r DITStructureRules) get(id any) (ds DITStructureRule) {
	L := r.len()
	if L == 0 {
		return
	}
	var n uint
	var name string

	switch tv := id.(type) {
	case int:
		if tv < 0 {
			return
		}
		n = uint(tv)
	case uint:
		n = tv
	case string:
		// string may be a string
		// uint, or a name.
		if _n, err := atoi(tv); err != nil {
			name = tv
		} else {
			return r.get(_n)
		}
	}

	for i := 0; i < L && ds.IsZero(); i++ {
		if _ds := r.index(i); !_ds.IsZero() {
			if _ds.ID() == name && len(name) > 0 {
				ds = _ds
			} else if _ds.dITStructureRule.ID == n {
				ds = _ds
			}
		}
	}

	return
}

func (r Schema) processDITStructureRule(ctx antlr4512.IDITStructureRuleDescriptionContext) (ds DITStructureRule, err error) {

	_ds := new(dITStructureRule)
	_ds.schema = r

	for k, ct := 0, ctx.GetChildCount(); k < ct && err == nil; k++ {
		switch tv := ctx.GetChild(k).(type) {
		case *antlr4512.StructureRuleContext,
			*antlr4512.OpenParenContext, *antlr4512.CloseParenContext:
			err = _ds.setCritical(tv)

		case *antlr4512.DefinitionNameContext,
			*antlr4512.DefinitionObsoleteContext,
			*antlr4512.DefinitionExtensionsContext,
			*antlr4512.DefinitionDescriptionContext:
			err = _ds.setMisc(tv)

		case *antlr4512.DSRFormContext:
			err = _ds.dSRNameFormContext(r, tv)

		case *antlr4512.DSRSuperRulesContext:
			err = _ds.superRulesContext(r, tv)
		default:
			env := DITStructureRule{_ds}
			err = isErrImpl(env.Type(), env.ID(), tv)
		}
	}

	if err == nil {
		// check for non-parser-related
		// errors in new instance.
		if err = _ds.check(); err == nil {
			if err = _ds.prepareString(); err == nil {
				_ds.t = nil
				ds = DITStructureRule{_ds}
			}
		}
	}

	return
}

func (r *dITStructureRule) prepareString() (err error) {
	buf := newBuf()
	r.t = newTemplate(`dITStructureRule`).
		Funcs(funcMap(map[string]any{
			`ExtensionSet`: r.Extensions.tmplFunc,
			`SuperLen`:     r.SuperRules.len,
			`IsObsolete`:   func() bool { return r.Obsolete },
		}))

	if r.t, err = r.t.Parse(dITStructureRuleTmpl); err == nil {
		if err = r.t.Execute(buf, r); err == nil {
			r.s = buf.String()
		}
	}

	return
}

func (r *dITStructureRule) check() (err error) { return }

func (r *dITStructureRule) superRulesContext(s Schema, ctx *antlr4512.DSRSuperRulesContext) (err error) {
	if ctx != nil {
		var sids []uint
		if x := ctx.StructureRule(); x != nil {
			var sid uint
			sid, err = ruleIDContext(x.(*antlr4512.StructureRuleContext))
			if err == nil {
				sids = append(sids, sid)
			}

		} else if y := ctx.StructureRules(); y != nil {
			sids, err = ruleIDsContext(y.(*antlr4512.StructureRulesContext))
		}

		if err == nil {
			err = errorf("No super rules parsed from %T", ctx)
			r.SuperRules = NewDITStructureRuleIDList()
			for i := 0; i < len(sids) && err == nil; i++ {
				err = errorf("super rule '%s' not found", sids[i])
				if sds := s.DITStructureRules().get(sids[i]); !sds.IsZero() {
					r.SuperRules.push(sds)
					err = nil
				}
			}
		}
	}

	return
}

func (r *dITStructureRule) dSRNameFormContext(s Schema, ctx *antlr4512.DSRFormContext) (err error) {
	err = errorf("%T is nil", ctx)
	if ctx != nil {
		o := ctx.OID()
		err = errorf("%T.%T is nil", ctx, o)
		if o != nil {
			var n, d []string
			assert, ok := o.(*antlr4512.OIDContext)
			if !ok {
				err = errorf("name form OID assertion failed for %T", o)
				return
			}

			if n, d, err = oIDContext(assert); err != nil {
				return
			}

			var nf string
			if len(n) > 0 {
				nf = n[0]
			} else if len(d) > 0 {
				nf = d[0]
			}

			err = errorf("name form '%s' not found", nf)
			if nff := s.NameForms().get(nf); !nff.IsZero() {
				err = nil
				r.Form = nff
			}
		}
	}

	return
}

// Use of single-valued OIDs applies to attributeType, dITStructureRule
// and nameForm definitions.
func ruleIDContext(ctx *antlr4512.StructureRuleContext) (id uint, err error) {

	dig := ctx.Number()
	if dig == nil {
		err = errorf("%T.%T is nil", ctx, dig)
		return
	}

	var _id int
	if _id, err = atoi(trimS(trim(dig.GetText(), `''`))); err != nil {
		return
	} else if _id < 0 {
		err = errorf("%T cannot be negative", ctx)
		return
	}
	id = uint(_id)

	return
}

// Use of multi-valued IDs applies solely to subordinate dITStructureRules
// in reference to their respective superior rules.
func ruleIDsContext(ctx *antlr4512.StructureRulesContext) (sids []uint, err error) {
	ch := ctx.AllStructureRule()

	for i := 0; i < len(ch) && err == nil; i++ {
		assert, ok := ch[i].(*antlr4512.StructureRuleContext)
		if !ok {
			err = errorf("%T type assertion failed for %T", ch[i], assert)
			break
		}

		var sid uint
		if sid, err = ruleIDContext(assert); err != nil {
			break
		}
		sids = append(sids, sid)
	}

	if len(sids) == 0 {
		err = errorf("No rule IDs parsed")
	}

	return
}

/*
IsZero returns a Boolean value indicative of nilness of the
receiver instance.
*/
func (r DITStructureRule) IsZero() bool {
	return r.dITStructureRule == nil
}

func (r *dITStructureRule) setCritical(ctx any) (err error) {
	err = errorf("Unknown critical context '%T'", ctx)

	switch tv := ctx.(type) {
	case *antlr4512.StructureRuleContext:
		r.ID, err = ruleIDContext(tv)
	case *antlr4512.OpenParenContext, *antlr4512.CloseParenContext:
		err = parenContext(tv)
	}

	return
}

func (r *dITStructureRule) setMisc(ctx any) (err error) {
	err = errorf("Unknown miscellaneous context '%T'", ctx)

	switch tv := ctx.(type) {
	case *antlr4512.DefinitionNameContext:
		r.Name, err = nameContext(tv)
	case *antlr4512.DefinitionDescriptionContext:
		r.Desc, err = descContext(tv)
	case *antlr4512.DefinitionExtensionsContext:
		r.Extensions, err = extContext(tv)
	case *antlr4512.DefinitionObsoleteContext:
		r.Obsolete = true
	}

	return
}

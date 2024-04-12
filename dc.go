package schemax

import (
	antlr4512 "github.com/JesseCoretta/go-rfc4512-antlr"
)

/*
ParseDITContentRule parses an individual textual attribute type (raw) and
returns an error instance.

When no error occurs, the newly formed [DITContentRule] instance -- based
on the parsed contents of raw -- is added to the receiver [DITContentRules]
slice instance.
*/
func (r *Schema) ParseDITContentRule(raw string) (err error) {
	var i antlr4512.Instance
	if i, err = antlr4512.ParseInstance(raw); err == nil {
		var d DITContentRule
		d, err = r.processDITContentRule(i.P.DITContentRuleDescription())
		if err == nil {
			err = r.DITContentRules().push(d)
		}
	}

	return
}

/*
NewDITContentRules initializes a new [Collection] instance and
casts it as a [DITContentRules] instance.
*/
func NewDITContentRules() DITContentRules {
	r := DITContentRules(newCollection(``))
	r.cast().SetPushPolicy(r.canPush)

	return r
}

/*
Description returns the underlying (optional) descriptive text
assigned to the receiver instance.
*/
func (r DITContentRule) Description() (desc string) {
	if !r.IsZero() {
		desc = r.dITContentRule.Desc
	}

	return
}

/*
List returns a map[string][]string instance which represents the current
inventory of [DITContentRule] instances within the receiver.  The keys
are numeric OIDs, while the values are zero (0) or more string slices,
each representing a name by which the definition is known.
*/
func (r DITContentRules) List() (list map[string][]string) {
	list = make(map[string][]string, 0)
	for i := 0; i < r.len(); i++ {
		def := r.index(i)
		list[def.NumericOID()] = def.Names().List()

	}

	return
}

/*
Type returns the string literal "dITContentRule".
*/
func (r DITContentRule) Type() string {
	return `dITContentRule`
}

/*
Type returns the string literal "dITContentRules".
*/
func (r DITContentRules) Type() string {
	return `dITContentRules`
}

/*
Len returns the current integer length of the receiver instance.
*/
func (r DITContentRules) Len() int {
	return r.len()
}

func (r DITContentRules) len() int {
	return r.cast().Len()
}

/*
String is a stringer method that returns the string representation
of the receiver instance.
*/
func (r DITContentRules) String() string {
	return r.cast().String()
}

/*
IsZero returns a Boolean value indicative of nilness of the
receiver instance.
*/
func (r DITContentRules) IsZero() bool {
	return r.cast().IsZero()
}

/*
Index returns the instance of [DITContentRule] found within the
receiver stack instance at index N.  If no instance is found at
the index specified, a zero [DITContentRule] instance is returned.
*/
func (r DITContentRules) Index(idx int) DITContentRule {
	return r.index(idx)
}

func (r DITContentRules) index(idx int) (dc DITContentRule) {
	slice, found := r.cast().Index(idx)
	if found {
		if _dc, ok := slice.(DITContentRule); ok {
			dc = _dc
		}
	}

	return
}

func (r DITContentRule) macro() (m []string) {
	if !r.IsZero() {
		m = r.dITContentRule.Macro
	}

	return
}

func (r DITContentRules) oIDsStringer(_ ...any) (present string) {
	var _present []string
	for i := 0; i < r.len(); i++ {
		_present = append(_present, r.index(i).OID())
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

		joined := join(_present, padchar+`$`+padchar)
		present = "(" + padchar + joined + padchar + ")"
	}

	return
}

/*
IsIdentifiedAs returns a Boolean value indicative of whether id matches
either the numericOID or descriptor of the receiver instance.  Case is
not significant in the matching process.
*/
func (r DITContentRule) IsIdentifiedAs(id string) (ident bool) {
	if !r.IsZero() {
		ident = id == r.NumericOID() || r.dITContentRule.Name.contains(id)
	}

	return
}

/*
IsObsolete returns a Boolean value indicative of definition obsolescence.
*/
func (r DITContentRule) IsObsolete() (o bool) {
	if !r.IsZero() {
		o = r.dITContentRule.Obsolete
	}

	return
}

/*
OID returns the string representation of an OID -- which is either a
numeric OID or descriptor -- that is held by the receiver instance.
*/
func (r DITContentRule) OID() (oid string) {
	if !r.IsZero() {
		oid = r.NumericOID() // default
		if r.dITContentRule.Name.len() > 0 {
			oid = r.dITContentRule.Name.index(0)
		}
	}

	return
}

/*
Aux returns an [ObjectClasses] containing zero (0) or more auxiliary
[ObjectClass] definitions for use within entries governed by this rule.
*/
func (r DITContentRule) Aux() (aux ObjectClasses) {
	if !r.IsZero() {
		aux = r.dITContentRule.Aux
	}

	return
}

/*
Must returns an [AttributeTypes] containing zero (0) or more required
[AttributeType] definitions for use within entries governed by this rule.
*/
func (r DITContentRule) Must() (must AttributeTypes) {
	if !r.IsZero() {
		must = r.dITContentRule.Must
	}

	return
}

/*
May returns an AttributeTypes containing zero (0) or more allowed
[AttributeType] definitions for use within entries governed by this rule.
*/
func (r DITContentRule) May() (may AttributeTypes) {
	if !r.IsZero() {
		may = r.dITContentRule.May
	}

	return
}

/*
Not returns an AttributeTypes containing zero (0) or more [AttributeType]
definitions disallowed for use within entries governed by this rule.
*/
func (r DITContentRule) Not() (not AttributeTypes) {
	if !r.IsZero() {
		not = r.dITContentRule.Not
	}

	return
}

/*
String is a stringer method that returns the string representation
of the receiver instance.
*/
func (r DITContentRule) String() (dsr string) {
	if !r.IsZero() {
		dsr = r.dITContentRule.s
	}

	return
}

func (r *dITContentRule) prepareString() (err error) {
	buf := newBuf()
	r.t = newTemplate(`dITContentRule`).
		Funcs(funcMap(map[string]any{
			`ExtensionSet`:  r.Extensions.tmplFunc,
			`StructuralOID`: r.OID.NumericOID,
			`AuxLen`:        r.Aux.len,
			`MayLen`:        r.May.len,
			`MustLen`:       r.Must.len,
			`NotLen`:        r.Not.len,
			`IsObsolete`:    func() bool { return r.Obsolete },
		}))
	if r.t, err = r.t.Parse(dITContentRuleTmpl); err == nil {
		if err = r.t.Execute(buf, r); err == nil {
			r.s = buf.String()
		}
	}

	return
}

/*
Name returns the string form of the principal name of the receiver instance, if set.
*/
func (r DITContentRule) Name() (id string) {
	if !r.IsZero() {
		id = r.dITContentRule.Name.index(0)
	}

	return
}

/*
Names returns the underlying instance of [DefinitionName] from within
the receiver.
*/
func (r DITContentRule) Names() (names DefinitionName) {
	return r.dITContentRule.Name
}

/*
Extensions returns the [Extensions] instance -- if set -- within
the receiver.
*/
func (r DITContentRule) Extensions() (e Extensions) {
	if !r.IsZero() {
		e = r.dITContentRule.Extensions
	}

	return
}

/*
NumericOID returns the string representation of the numeric OID
held by the receiver instance.
*/
func (r DITContentRule) NumericOID() (noid string) {
	if !r.IsZero() {
		if soid := r.dITContentRule.OID; !soid.IsZero() {
			noid = soid.NumericOID()
		}
	}

	return
}

func (r DITContentRule) setOID(_ string) {}

/*
Push returns an error following an attempt to push a [DITContentRule]
into the receiver stack instance.
*/
func (r DITContentRules) Push(dc any) error {
	return r.push(dc)
}

func (r DITContentRules) push(dc any) (err error) {
	err = errorf("%T instance is nil; cannot append to %T", dc, r)
	if dc != nil {
		r.cast().Push(dc)
		err = nil
	}

	return
}

/*
IsZero returns a Boolean value indicative of nilness of the
receiver instance.
*/
func (r DITContentRule) IsZero() bool {
	return r.dITContentRule == nil
}

func (r *dITContentRule) setMisc(ctx any) (err error) {
	err = errorf("Unknown miscellaneous context '%T'", ctx)

	switch tv := ctx.(type) {
	case *antlr4512.DefinitionNameContext:
		r.Name, err = nameContext(tv)
	case *antlr4512.DefinitionDescriptionContext:
		r.Desc, err = descContext(tv)
	case *antlr4512.DefinitionExtensionsContext:
		r.Extensions, err = extContext(tv)
	}

	return
}

func (r DITContentRules) contains(id string) bool {
	return !r.get(id).IsZero()
}

func (r DITContentRules) get(id string) (dc DITContentRule) {
	for i := 0; i < r.len() && dc.IsZero(); i++ {
		if _dc := r.index(i); !_dc.IsZero() {
			if _dc.OID() == id {
				dc = _dc
			} else if _dc.dITContentRule.Name.contains(id) {
				dc = _dc
			}
		}
	}

	return
}

func (r Schema) processDITContentRule(ctx antlr4512.IDITContentRuleDescriptionContext) (dc DITContentRule, err error) {

	_dc := new(dITContentRule)
	_dc.schema = r

	for k, ct := 0, ctx.GetChildCount(); k < ct && err == nil; k++ {
		switch tv := ctx.GetChild(k).(type) {
		case *antlr4512.OpenParenContext,
			*antlr4512.CloseParenContext:
			err = parenContext(tv)

		case *antlr4512.DefinitionNameContext,
			*antlr4512.DefinitionExtensionsContext,
			*antlr4512.DefinitionDescriptionContext:
			err = _dc.setMisc(tv)

		case *antlr4512.NumericOIDOrMacroContext:
			err = _dc.structuralOID(r, tv)

		case *antlr4512.DefinitionObsoleteContext:
			_dc.Obsolete = true

		case *antlr4512.DefinitionMustContext:
			err = _dc.mustContext(r, tv)

		case *antlr4512.DefinitionMayContext:
			err = _dc.mayContext(r, tv)

		case *antlr4512.DCRNotContext:
			err = _dc.notContext(r, tv)

		case *antlr4512.DCRAuxContext:
			err = _dc.auxContext(r, tv)

		default:
			env := DITContentRule{_dc}
			err = isErrImpl(env.Type(), env.OID(), tv)
		}
	}

	// check for errors in new instance
	if err == nil {
		if err = _dc.check(); err == nil {
			if err = _dc.prepareString(); err == nil {
				_dc.t = nil
				dc = DITContentRule{_dc}
			}
		}
	}

	return
}

func (r *dITContentRule) structuralOID(s Schema, ctx *antlr4512.NumericOIDOrMacroContext) (err error) {
	var oc string
	if oc, r.Macro, err = numOIDContext(ctx); err != nil {
		return
	}

	var soc ObjectClass
	if soc = s.ObjectClasses().get(oc); soc.IsZero() {
		err = errorf("structural class '%s' not found; cannot process %T",
			oc, r)
		return
	}
	r.OID = soc

	return
}

func (r *dITContentRule) mustContext(s Schema, ctx *antlr4512.DefinitionMustContext) (err error) {
	var must []string
	if must, err = mustContext(ctx); err != nil {
		return
	}

	r.Must = NewAttributeTypeOIDList()
	for i := 0; i < len(must); i++ {
		var mustt AttributeType
		if mustt = s.AttributeTypes().get(must[i]); mustt.IsZero() {
			err = errorf("required attr '%s' not found; cannot process %T",
				must[i], r)
			break
		}
		r.Must.push(mustt)
	}

	return
}

func (r *dITContentRule) mayContext(s Schema, ctx *antlr4512.DefinitionMayContext) (err error) {
	var may []string
	if may, err = mayContext(ctx); err != nil {
		return
	}

	r.May = NewAttributeTypeOIDList()
	for i := 0; i < len(may); i++ {
		var mayy AttributeType
		if mayy = s.AttributeTypes().get(may[i]); mayy.IsZero() {
			err = errorf("required attr '%s' not found; cannot process %T",
				may[i], r)
			break
		}
		r.May.push(mayy)
	}

	return
}

func (r *dITContentRule) notContext(s Schema, ctx *antlr4512.DCRNotContext) (err error) {
	var not []string
	if not, err = notContext(ctx); err != nil {
		return
	}

	r.Not = NewAttributeTypeOIDList()
	for i := 0; i < len(not); i++ {
		var nott AttributeType
		if nott = s.AttributeTypes().get(not[i]); nott.IsZero() {
			err = errorf("required attr '%s' not found; cannot process %T",
				not[i], r)
			break
		}
		r.Not.push(nott)
	}

	return
}

func notContext(ctx *antlr4512.DCRNotContext) (not []string, err error) {
	if ctx != nil {
		var n, d []string

		if x := ctx.OID(); x != nil {
			n, d, err = oIDContext(x.(*antlr4512.OIDContext))
		} else if y := ctx.OIDs(); y != nil {
			n, d, err = oIDsContext(y.(*antlr4512.OIDsContext))
		}

		if err == nil {
			not = append(not, n...)
			not = append(not, d...)
		}

		if len(not) == 0 {
			if err != nil {
				err = errorf("No forbidden attributes parsed from %T: %v", ctx, err)
			} else {
				err = errorf("No forbidden attributes parsed from %T", ctx)
			}
		}
	}

	return
}

func (r *dITContentRule) check() (err error) {
	if r == nil {
		err = errorf("%T is nil")
		return
	}

	if len(r.OID.NumericOID()) == 0 {
		err = errorf("%T lacks an OID", r)
	}

	return
}

// stackage closure func - do not exec directly.
func (r DITContentRules) canPush(x ...any) (err error) {
	if len(x) == 0 {
		return
	}

	for i := 0; i < len(x) && err == nil; i++ {
		instance := x[i]
		err = errorf("Type assertion for %T has failed", instance)
		if dc, ok := instance.(DITContentRule); ok && !dc.IsZero() {
			err = errorf("%T %s not unique", dc, dc.NumericOID())
			if tst := r.get(dc.NumericOID()); tst.IsZero() {
				err = nil
			}
		}
	}

	return
}

func (r *dITContentRule) auxContext(s Schema, ctx *antlr4512.DCRAuxContext) (err error) {
	var aux []string
	if aux, err = auxContext(ctx); err != nil {
		return
	}

	r.Aux = NewObjectClassOIDList()
	for i := 0; i < len(aux); i++ {
		var auxx ObjectClass
		if auxx = s.ObjectClasses().get(aux[i]); auxx.IsZero() {
			err = errorf("auxiliary class '%s' not found; cannot process %T",
				aux[i], r)
			break
		}
		r.Aux.push(auxx)
	}

	return
}

func auxContext(ctx *antlr4512.DCRAuxContext) (aux []string, err error) {
	if ctx != nil {
		var n, d []string

		if x := ctx.OID(); x != nil {
			n, d, err = oIDContext(x.(*antlr4512.OIDContext))
		} else if y := ctx.OIDs(); y != nil {
			n, d, err = oIDsContext(y.(*antlr4512.OIDsContext))
		}

		if err == nil {
			aux = append(aux, n...)
			aux = append(aux, d...)
		}

		if len(aux) == 0 {
			if err != nil {
				err = errorf("No auxiliary classes parsed from %T: %v", ctx, err)
			} else {
				err = errorf("No auxiliary classes parsed from %T", ctx)
			}
		}
	}

	return
}

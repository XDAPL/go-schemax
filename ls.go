package schemax

import (
	"internal/rfc2307"
	"internal/rfc4517"
	"internal/rfc4523"
	"internal/rfc4530"

	antlr4512 "github.com/JesseCoretta/go-rfc4512-antlr"
)

var (
	rfc2307Macros   map[string]string             = rfc2307.Macros
	rfc2307Syntaxes rfc2307.LDAPSyntaxDefinitions = rfc2307.AllLDAPSyntaxes
	rfc4517Syntaxes rfc4517.LDAPSyntaxDefinitions = rfc4517.AllLDAPSyntaxes
	rfc4523Syntaxes rfc4523.LDAPSyntaxDefinitions = rfc4523.AllLDAPSyntaxes
	rfc4530Syntaxes rfc4530.LDAPSyntaxDefinitions = rfc4530.AllLDAPSyntaxes
)

/*
ParseLDAPSyntax parses raw into an instance of [LDAPSyntax], which is
appended to the receiver's LS stack.
*/
func (r Schema) ParseLDAPSyntax(raw string) (err error) {
	var i antlr4512.Instance
	if i, err = antlr4512.ParseInstance(raw); err == nil {
		var l LDAPSyntax
		if l, err = r.processLDAPSyntax(i.P.LDAPSyntaxDescription()); err != nil {
			return
		}

		err = r.LDAPSyntaxes().push(l)
	}

	return
}

/*
NewLDAPSyntaxes initializes a new [Collection] instance and casts it
as an [LDAPSyntaxes] instance.
*/
func NewLDAPSyntaxes() LDAPSyntaxes {
	r := LDAPSyntaxes(newCollection(`ldapSyntaxes`))
	r.cast().SetPushPolicy(r.canPush)

	return r
}

/*
Type returns the string literal "ldapSyntax".
*/
func (r LDAPSyntax) Type() string {
	return `ldapSyntax`
}

/*
Type returns the string literal "ldapSyntaxes".
*/
func (r LDAPSyntaxes) Type() string {
	return `ldapSyntaxes`
}

/*
xOrigin returns an instance of LDAPSyntaxes containing only definitions
which bear the X-ORIGIN value of x. Case is not significant in the matching
process, nor is whitespace (e.g.: RFC 4517 vs. RFC4517).
*/
func (r LDAPSyntaxes) xOrigin(x string) (lss LDAPSyntaxes) {
	lss = NewLDAPSyntaxes()
	for i := 0; i < r.Len(); i++ {
		ls := r.Index(i)
		if xo, found := ls.Extensions().Get(`X-ORIGIN`); found {
			if xo.contains(x) {
				lss.push(ls)
			}
		}
	}

	return
}

/*
Extensions returns the [Extensions] instance -- if set -- within
the receiver.
*/
func (r LDAPSyntax) Extensions() (e Extensions) {
	if !r.IsZero() {
		e = r.lDAPSyntax.Extensions
	}

	return
}

/*
String is a stringer method that returns the string representation
of the receiver instance.
*/
func (r LDAPSyntax) String() (ls string) {
	if !r.IsZero() {
		ls = r.lDAPSyntax.s
	}

	return
}

/*
IsObsolete only returns a false Boolean value, as definition obsolescence
does not apply to [LDAPSyntax] definitions.  This method exists only to
satisfy Go's interface signature requirements.
*/
func (r LDAPSyntax) IsObsolete() bool { return false }

/*
IsIdentifiedAs returns a Boolean value indicative of whether id matches
either the numericOID or description of the receiver instance.  Case is
not significant in the matching process.
*/
func (r LDAPSyntax) IsIdentifiedAs(id string) (ident bool) {
	if !r.IsZero() {
		ident = id == r.NumericOID() || eq(r.lDAPSyntax.Desc, id)
	}

	return
}

func (r *lDAPSyntax) prepareString() (err error) {
	buf := newBuf()
	r.t = newTemplate(`ldapSyntax`).
		Funcs(funcMap(map[string]any{
			`ExtensionSet`: r.Extensions.tmplFunc,
		}))
	if r.t, err = r.t.Parse(lDAPSyntaxTmpl); err == nil {
		if err = r.t.Execute(buf, r); err == nil {
			r.s = buf.String()
		}
	}

	return
}

func (r LDAPSyntax) macro() (m []string) {
	if !r.IsZero() {
		m = r.lDAPSyntax.Macro
	}

	return
}

func (r LDAPSyntax) setOID(x string) {
	if !r.IsZero() {
		r.lDAPSyntax.OID = x
	}
}

/*
OID is an alias for the [LDAPSyntax.NumericOID] method, as
[LDAPSyntax] instances do not allow for the assignment of
names.  This method exists solely to satisfy Go's interface
signature requirements.
*/
func (r LDAPSyntax) OID() string { return r.NumericOID() }

/*
Name returns an empty string, as [LDAPSyntax] definitions do not
bear names.  This method exists only to satisfy Go's interface
signature requirements.
*/
func (r LDAPSyntax) Name() string { return `` }

/*
Names returns an empty instance of [DefinitionName], as names do not
apply to [LDAPSyntax] definitions.  This method exists only to satisfy
Go's interface signature requirements.
*/
func (r LDAPSyntax) Names() DefinitionName { return DefinitionName{} }

/*
NumericOID returns the string representation of the numeric OID
held by the receiver instance.
*/
func (r LDAPSyntax) NumericOID() (noid string) {
	if !r.IsZero() {
		noid = r.lDAPSyntax.OID
	}

	return
}

/*
Description returns the underlying (optional) descriptive text
assigned to the receiver instance.
*/
func (r LDAPSyntax) Description() (desc string) {
	if !r.IsZero() {
		desc = r.lDAPSyntax.Desc
	}
	return
}

/*
IsZero returns a Boolean value indicative of nilness of the
receiver instance.
*/
func (r LDAPSyntax) IsZero() bool {
	return r.lDAPSyntax == nil
}

/*
List returns a map[string]string instance which represents the current
inventory of [LDAPSyntax] instances within the receiver.  The keys are
numeric OIDs, while the value is the description of the [LDAPSyntax]
definition is known.
*/
func (r LDAPSyntaxes) List() (list map[string]string) {
	list = make(map[string]string, 0)
	for i := 0; i < r.len(); i++ {
		def := r.index(i)
		list[def.NumericOID()] = def.Description()
	}

	return
}

func (r LDAPSyntaxes) canPush(x ...any) (err error) {
	if len(x) == 0 {
		return
	}

	for i := 0; i < len(x) && err == nil; i++ {
		instance := x[i]
		if ls, ok := instance.(LDAPSyntax); !ok || ls.IsZero() {
			err = errorf("Type assertion for %T has failed", instance)
		} else if tst := r.get(ls.NumericOID()); !tst.IsZero() {
			err = errorf("%T %s not unique", ls, ls.NumericOID())
		}
	}

	return
}

/*
Len returns the current integer length of the receiver instance.
*/
func (r LDAPSyntaxes) Len() int {
	return r.len()
}

func (r LDAPSyntaxes) len() int {
	return r.cast().Len()
}

/*
IsZero returns a Boolean value indicative of nilness of the
receiver instance.
*/
func (r LDAPSyntaxes) IsZero() bool {
	return r.cast().IsZero()
}

func (r LDAPSyntaxes) String() string {
	return r.cast().String()
}

/*
Index returns the instance of [LDAPSyntax] found within the
receiver stack instance at index N.  If no instance is found at
the index specified, a zero [LDAPSyntax] instance is returned.
*/
func (r LDAPSyntaxes) Index(idx int) LDAPSyntax {
	return r.index(idx)
}

func (r LDAPSyntaxes) index(idx int) (ls LDAPSyntax) {
	slice, found := r.cast().Index(idx)
	if found {
		if _ls, ok := slice.(LDAPSyntax); ok {
			ls = _ls
		}
	}

	return
}

/*
Push returns an error following an attempt to push an [LDAPSyntax]
into the receiver stack instance.
*/
func (r LDAPSyntaxes) Push(ls any) error {
	return r.push(ls)
}

func (r LDAPSyntaxes) push(ls any) (err error) {
	if ls == nil {
		err = errorf("%T instance is nil; cannot append to %T", ls, r)
		return
	}

	r.cast().Push(ls)

	return
}

// cyclo=0
func (r LDAPSyntaxes) contains(id string) bool {
	return !r.get(id).IsZero()
}

// cyclo=6
func (r LDAPSyntaxes) get(id string) (ls LDAPSyntax) {
	for i := 0; i < r.len() && ls.IsZero(); i++ {
		if _ls := r.index(i); !_ls.IsZero() {
			if eq(_ls.lDAPSyntax.OID, id) {
				ls = _ls
			}
		}
	}

	return
}

func (r Schema) processLDAPSyntax(ctx antlr4512.ILDAPSyntaxDescriptionContext) (ls LDAPSyntax, err error) {
	_ls := new(lDAPSyntax)
	_ls.schema = r

	for k, ct := 0, ctx.GetChildCount(); k < ct && err == nil; k++ {
		switch tv := ctx.GetChild(k).(type) {
		case *antlr4512.OpenParenContext,
			*antlr4512.CloseParenContext:
			err = parenContext(tv)
		case *antlr4512.NumericOIDOrMacroContext:
			_ls.OID, _ls.Macro, err = numOIDContext(tv)
		case *antlr4512.DefinitionDescriptionContext:
			_ls.Desc, err = descContext(tv)
		case *antlr4512.DefinitionExtensionsContext:
			_ls.Extensions, err = extContext(tv)
		default:
			env := LDAPSyntax{_ls}
			err = isErrImpl(env.Type(), env.OID(), tv)
		}
	}

	if err == nil {
		r.resolveByMacro(LDAPSyntax{_ls})

		if err = _ls.check(); err == nil {
			if err = _ls.prepareString(); err == nil {
				_ls.t = nil
				ls = LDAPSyntax{_ls}
			}
		}
	}

	return
}

func (r *lDAPSyntax) check() (err error) {
	if r == nil {
		err = errorf("%T is nil", r)
		return
	}

	if len(r.OID) == 0 {
		err = errorf("%T lacks an OID", r)
	}

	return
}

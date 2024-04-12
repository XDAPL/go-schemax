package schemax

import (
	"bytes"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"
	"unicode"

	"github.com/JesseCoretta/go-rfc4512-antlr"
	"github.com/antlr4-go/antlr/v4"
)

var (
	atoi    func(string) (int, error)                         = strconv.Atoi
	itoa    func(int) string                                  = strconv.Itoa
	sprintf func(string, ...any) string                       = fmt.Sprintf
	printf  func(string, ...any) (int, error)                 = fmt.Printf
	repAll  func(string, string, string) string               = strings.ReplaceAll
	eq      func(string, string) bool                         = strings.EqualFold
	split   func(string, string) []string                     = strings.Split
	join    func([]string, string) string                     = strings.Join
	hasPfx  func(string, string) bool                         = strings.HasPrefix
	hasSfx  func(string, string) bool                         = strings.HasSuffix
	lc      func(string) string                               = strings.ToLower
	uc      func(string) string                               = strings.ToUpper
	trim    func(string, string) string                       = strings.Trim
	trimL   func(string, string) string                       = strings.TrimLeft
	trimR   func(string, string) string                       = strings.TrimRight
	trimS   func(string) string                               = strings.TrimSpace
	isDigit func(rune) bool                                   = unicode.IsDigit
	isAlpha func(rune) bool                                   = unicode.IsLetter
	parseI  func(string, ...bool) (antlr4512.Instance, error) = antlr4512.ParseInstance
)

func isAlnum(x rune) bool {
	return isDigit(x) || isAlpha(x)
}

func bool2str(x bool) string {
	if x {
		return `true`
	}

	return `false`
}

/*
errorf wraps errors.New and returns a non-nil instance of error
based upon a non-nil/non-zero msg input value with optional args.
*/
func errorf(msg any, x ...any) error {
	switch tv := msg.(type) {
	case string:
		if len(tv) > 0 {
			return errors.New(sprintf(tv, x...))
		}
	case error:
		if tv != nil {
			return errors.New(sprintf(tv.Error(), x...))
		}
	}

	return nil
}

/*
isErrImpl returns an error following an attempt to assert
ctx to *antlr.ErrorNodeImpl. If asserted, determine if the
underlying instance describes an erroneous token.
*/
func isErrImpl(typ, id string, ctx any) (err error) {
	if ern, ok := ctx.(*antlr.ErrorNodeImpl); ok {
		// Don't count stray newlines or
		// whitespaces as problem tokens
		txt := trimS(repAll(ern.GetText(), string(rune(10)), ``))
		if len(txt) > 0 {
			err = errorf("Unexpected or bogus token '%s' in %s '%s'",
				txt, typ, id)
		}
	}

	return
}

/*
parenContext returns an error based on a lack of opening and/or closing
parenthetical encapsulation.  This applies to any multi-valued attribute
sequence value, such as MUST ( cn $ sn $ o ) and names ( 'cn' 'commonName' ).
*/
func parenContext(ctx any) (err error) {
	if ctx != nil {
		switch tv := ctx.(type) {
		case *antlr4512.OpenParenContext:
			if hasPfx(tv.GetText(), `<missing`) {
				err = errorf("Missing open parenthesis for definition", tv)
			}
		case *antlr4512.CloseParenContext:
			if hasPfx(tv.GetText(), `<missing`) {
				err = errorf("Missing close parenthesis for definition", tv)
			}
		}
	}

	return
}

/*
descContext extracts descriptive quoted-string or quoted-descriptor
text from the ctx instance, returning it via desc.
*/
func descContext(ctx *antlr4512.DefinitionDescriptionContext) (desc string, err error) {
	if ctx != nil {
		if qs := ctx.QuotedString(); qs != nil {
			desc = trimS(trim(qs.GetText(), `''`))
		} else if qd := ctx.QuotedDescriptor(); qd != nil {
			desc = trimS(trim(qd.GetText(), `''`))
		}
	}

	return
}

/*
isValidDescription uses ANTLR to parse the input desc value in order to
determine whether it is an RFC 4512-compliant qdstr (Quoted Descriptor).
*/
func isValidDescription(desc string) (ok bool) {
	if p, err := antlr4512.ParseInstance(desc); err == nil {
		txt := p.P.DefinitionDescription().GetText()
		ok = !hasPfx(txt, `<missing`) && txt == desc
	}
	return
}

/*
extVContext extracts the extension values from the ctx instance, and
returns them via extv ([]string).
*/
func extVContext(ctx []antlr4512.IExtensionValueContext) (extv []string) {
	if ctx != nil {
		for i := 0; i < len(ctx); i++ {
			var ev string
			if qs := ctx[i].QuotedString(); qs != nil {
				ev = qs.GetText()
			} else if qd := ctx[i].QuotedDescriptor(); qd != nil {
				ev = qd.GetText()
			}
			if ev = trimS(trim(ev, `''`)); len(ev) > 0 {
				extv = append(extv, ev)
			}
		}
	}

	return
}

/*
numOIDContext extracts the numeric OID from the definition, returning
it via noid.  This applies to all definitions EXCEPT dITStructureRule.
*/
func numOIDContext(ctx *antlr4512.NumericOIDOrMacroContext) (noid string, macro []string, err error) {
	if ctx != nil {
		if n := ctx.NumericOID(); n != nil {
			if noid = n.GetText(); hasPfx(noid, `<missing`) {
				err = errorf("Missing numeric OID for definition")
			}
		} else if mc := ctx.Macro(); mc != nil {
			var ok bool
			if macro, ok = macroContext(mc); !ok {
				err = errorf("Invalid macro value(s)")
			}
		}
	}

	return
}

/*
macroContext extracts the macro data from the definition, returning
slices of macro values (name,oid_suffix) alongside a Boolean value.
The macro slice values are used to resolve and acquire the complete
numeric OID.  The Boolean value indicates that a complete macro was
acquired.
*/
func macroContext(ctx antlr4512.IMacroContext) (macro []string, ok bool) {
	var m, s string

	if m = ctx.Descriptor().GetText(); hasPfx(m, `<missing`) {
		return
	} else if s = ctx.MacroSuffix().GetText(); hasPfx(s, `<missing`) {
		return
	}

	if ok = len(m) > 0 && len(s) > 0; ok {
		if hasPfx(s, `:`) || hasPfx(s, `.`) {
			s = s[1:]
		}
		macro = []string{m, s}
	}
	return
}

/*
extContext extracts the X-<NAME> extension key/value pair from the
ctx instance, and returns it via e.  This applies to all definition
types.
*/
func extContext(ctx *antlr4512.DefinitionExtensionsContext) (e Extensions, err error) {
	if ctx != nil {
		exts := ctx.AllDefinitionExtension()
		e = NewExtensions()
		for i := 0; i < len(exts); i++ {
			if en := exts[i].ExtensionName(); en != nil {
				qdesc := exts[i].AllExtensionValue()
				if ev := extVContext(qdesc); len(ev) > 0 {
					e.Set(trimS(en.GetText()), ev...)
				}
			}
		}
	}

	return
}

/*
syntaxContext extracts the syntax numeric OID from the ctx instance,
and returns it through soid.  This applies to attributeType and
matchingRule definitions.
*/
func syntaxContext(ctx *antlr4512.DefinitionSyntaxContext) (soid string, err error) {
	if ctx != nil {
		_soid := ctx.NumericOID()
		if pfx := _soid.GetText(); hasPfx(pfx, `<missing`) {
			err = errorf("Missing numeric OID for definition syntax")
		} else {
			soid = _soid.GetText()
		}
	}

	return
}

/*
descrNumOIDContext returns the definition descriptor OR numeric OID
within x and identifies which source was positive via noid.  In no
scenario do both sources apply.
*/
func descrNumOIDContext(ctx *antlr4512.OIDContext) (x string, noid bool, err error) {
	if ctx != nil {
		if _x := ctx.Descriptor(); _x != nil {
			x = _x.GetText()
		} else if _y := ctx.NumericOID(); _y != nil {
			x = _y.GetText()
			noid = true
		}

		if len(x) == 0 {
			err = errorf("No %T parsed", ctx)
		}
	}

	return
}

/*
Use of single-valued OIDs applies to attributeType, dITStructureRule
and nameForm definitions.
*/
func oIDContext(ctx *antlr4512.OIDContext) (n, d []string, err error) {
	var (
		x    string
		noid bool
	)

	if x, noid, err = descrNumOIDContext(ctx); err != nil {
		return
	}

	if x = trimS(trim(x, `''`)); len(x) > 0 {
		if noid {
			n = append(n, x)
		} else {
			d = append(d, x)
		}
	}

	return
}

/*
Use of multi-valued OIDs applies to objectClass, dITContentRule,
nameForm, and matchingRuleUse definitions.
*/
func oIDsContext(ctx *antlr4512.OIDsContext) (n, d []string, err error) {
	ch := ctx.AllOID()

	for i := 0; i < len(ch); i++ {
		var _n, _d []string
		assert, ok := ch[i].(*antlr4512.OIDContext)
		if !ok {
			err = errorf("%T type assertion failed for %T", ch[i], assert)
			break
		}

		if _n, _d, err = oIDContext(assert); err != nil {
			break
		}

		if len(_n) > 0 {
			n = append(n, _n...)
		}

		if len(_d) > 0 {
			d = append(d, _d...)
		}
	}

	return
}

/*
Name applies to all definitions EXCEPT ldapSyntax.
*/
func scanNames(id string, names []string) (idx int) {
	idx = -1

	for i := 0; i < len(names); i++ {
		if eq(names[i], id) {
			idx = i
			break
		}
	}

	return
}

/*
must applies to OC and DCR.
*/
func mustContext(ctx *antlr4512.DefinitionMustContext) (must []string, err error) {
	if ctx != nil {
		var n, d []string

		if x := ctx.OID(); x != nil {
			n, d, err = oIDContext(x.(*antlr4512.OIDContext))
		} else if y := ctx.OIDs(); y != nil {
			n, d, err = oIDsContext(y.(*antlr4512.OIDsContext))
		}

		if err == nil {
			must = append(must, n...)
			must = append(must, d...)
		}

		if len(must) == 0 {
			if err != nil {
				err = errorf("No required attributes parsed from %T: %v", ctx, err)
			} else {
				err = errorf("No required attributes parsed from %T", ctx)
			}
		}
	}

	return
}

/*
may applies to OC and DCR.
*/
func mayContext(ctx *antlr4512.DefinitionMayContext) (may []string, err error) {
	if ctx != nil {
		var n, d []string

		if x := ctx.OID(); x != nil {
			n, d, err = oIDContext(x.(*antlr4512.OIDContext))
		} else if y := ctx.OIDs(); y != nil {
			n, d, err = oIDsContext(y.(*antlr4512.OIDsContext))
		}

		if err == nil {
			may = append(may, n...)
			may = append(may, d...)
		}

		if len(may) == 0 {
			if err != nil {
				err = errorf("No permitted attributes parsed from %T: %v", ctx, err)
			} else {
				err = errorf("No permitted attributes parsed from %T", ctx)
			}
		}
	}

	return
}

func removeNL(b string) string {
	return repAll(b,
		string(rune(10)),
		string(rune(32)))
}

/*
condenseWHSP returns input string b with all contiguous
WHSP characters condensed into single space characters.

WHSP is qualified through space or TAB chars (ASCII #32
and #9 respectively).
*/
func condenseWHSP(b string) (a string) {
	// remove leading and trailing
	// WHSP characters ...
	b = trimS(b)
	b = repAll(b, string(rune(10)), string(rune(32)))

	var last bool
	for i := 0; i < len(b); i++ {
		c := rune(b[i])
		switch c {
		// match space (32) or tab (9)
		case rune(9), rune(32):
			if !last {
				last = true
				a += string(rune(32))
			}

		default:
			if last {
				last = false
			}
			a += string(c)
		}
	}

	a = trimS(a) //once more
	return
}

/*
readFile uses os.ReadFile to attempt to read f into raw, which is
returned alongside an error.

Only files with a suffix of ".schema" will be eligible for processing.
*/
func readFile(f string) (raw []byte, err error) {
	if hasSfx(f, `.schema`) {
		raw, err = os.ReadFile(f)
	}

	return
}

/*
readDirectory recurses all files and folders specified at 'dir',
returning parsed schema bytes (content) alongside an error.

Only files with an extension of ".schema" will be parsed, but all
subdirectories will be traversed in search of these files.

File and directory naming schemes MUST guarantee the appropriate
ordering of any and all sub types, sub rules and sub classes which
would rely on the presence of dependency definitions (e.g.: 'cn'
cannot exist without 'name').
*/
func readDirectory(dir string) (content []byte, err error) {
	// remove any number of trailing
	// slashes from dir.
	dir = trimR(dir, `/`)

	// interim content storage
	var _content []byte

	// recurse dir path
	err = filepath.Walk(dir, func(p string, d fs.FileInfo, err error) error {
		if !d.IsDir() {
			var subcnt []byte
			// to avoid accidental splicing of two files
			// with regards to a missing newline (and
			// thus avoiding an error), insert a newline
			// char (rune(10)) if not already present at
			// the end of the byte sequence.
			if subcnt, err = readFile(p); err == nil && len(subcnt) > 0 {
				if string(subcnt[len(subcnt)-1]) != string(rune(10)) {
					_content = append(_content, []byte(string(rune(10)))...)
				}
				_content = append(_content, subcnt...)
			}
		}

		return err
	})

	if len(_content) > 0 && err == nil {
		content = _content
	}

	return
}

/*
is 'val' an unsigned number?
*/
func isNumber(val string) bool {
	if len(val) == 0 {
		return false
	}

	for i := 0; i < len(val); i++ {
		if !isDigit(rune(val[i])) {
			return false
		}
	}
	return true
}

func newTemplate(name string) *template.Template {
	return template.New(name)
}

func newBuf() *bytes.Buffer {
	return new(bytes.Buffer)
}

func funcMap(fm map[string]any) template.FuncMap {
	return template.FuncMap(fm)
}

func (r Counters) String() string {
	return "\nLS:" + itoa(r.LS) +
		"\nMR:" + itoa(r.MR) +
		"\nAT:" + itoa(r.AT) +
		"\nMU:" + itoa(r.MU) +
		"\nOC:" + itoa(r.OC) +
		"\nDC:" + itoa(r.DC) +
		"\nNF:" + itoa(r.NF) +
		"\nDS:" + itoa(r.DS)
}

func isNumericOID(id string) bool {
	if len(id) == 0 {
		return false
	}

	if !('0' <= rune(id[0]) && rune(id[0]) <= '2') || id[len(id)-1] == '.' {
		return false
	}

	var last rune
	for _, c := range id {
		switch {
		case c == '.':
			if last == c {
				return false
			}
			last = '.'
		case isDigit(c):
			last = c
			continue
		}
	}

	return true
}

/*
IsDescriptor scans the input string val and judges whether
it qualifies a descriptor, in that all of the following
evaluate as true:

  - non-zero in length
  - begins with an alphabetical character
  - ends in an alphanumeric character
  - contains only alphanumeric characters or hyphens
  - no contiguous hyphens

This function is an alternative to engaging the antlr4512
parser.
*/
func IsDescriptor(val string) bool {
	return isDescriptor(val)
}

func isDescriptor(val string) bool {
	if len(val) == 0 {
		return false
	}

	// must begin with an alpha.
	if !isAlpha(rune(val[0])) {
		return false
	}

	// can only end in alnum.
	if !isAlnum(rune(val[len(val)-1])) {
		return false
	}

	// watch hyphens to avoid contiguous use
	var lastHyphen bool

	// iterate all characters in val, checking
	// each one for "descr" validity.
	for i := 0; i < len(val); i++ {
		ch := rune(val[i])
		switch {
		case isAlnum(ch):
			lastHyphen = false
		case ch == '-':
			if lastHyphen {
				// cannot use consecutive hyphens
				return false
			}
			lastHyphen = true
		default:
			// invalid character (none of [a-zA-Z0-9\-])
			return false
		}
	}

	return true
}

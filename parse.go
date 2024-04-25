package schemax

import (
	"github.com/JesseCoretta/go-antlr4512"
)

var (
	new4512Schema func() antlr4512.Schema                          = antlr4512.NewSchema
	parseLS       func(string) (antlr4512.LDAPSyntax, error)       = antlr4512.ParseLDAPSyntax
	parseMR       func(string) (antlr4512.MatchingRule, error)     = antlr4512.ParseMatchingRule
	parseAT       func(string) (antlr4512.AttributeType, error)    = antlr4512.ParseAttributeType
	parseMU       func(string) (antlr4512.MatchingRuleUse, error)  = antlr4512.ParseMatchingRuleUse
	parseOC       func(string) (antlr4512.ObjectClass, error)      = antlr4512.ParseObjectClass
	parseDC       func(string) (antlr4512.DITContentRule, error)   = antlr4512.ParseDITContentRule
	parseNF       func(string) (antlr4512.NameForm, error)         = antlr4512.ParseNameForm
	parseDS       func(string) (antlr4512.DITStructureRule, error) = antlr4512.ParseDITStructureRule
)

/*
ParseLDAPSyntax parses raw into an instance of [LDAPSyntax], which is
appended to the receiver's LS stack.
*/
func (r Schema) ParseLDAPSyntax(raw string) error {
	def, err := parseLS(raw)
	if err == nil {
		var _def LDAPSyntax
		if _def, err = r.marshalLS(def); err == nil {
			r.LDAPSyntaxes().push(_def)
		}
	}

	return err
}

/*
ParseMatchingRule returns an error following an attempt to parse raw into the
receiver instance's [MatchingRules] instance.
*/
func (r Schema) ParseMatchingRule(raw string) error {
	def, err := parseMR(raw)
	if err == nil {
		var _def MatchingRule
		if _def, err = r.marshalMR(def); err == nil {
			r.MatchingRules().push(_def)
		}
	}

	return err
}

/*
ParseMatchingRuleUse returns an error following an attempt to parse raw
into the receiver instance's [MatchingRuleUses] instance.
*/
func (r Schema) ParseMatchingRuleUse(raw string) error {
        def, err := parseMU(raw)
        if err == nil {
                var _def MatchingRuleUse
                if _def, err = r.marshalMU(def); err == nil {
                        r.MatchingRuleUses().push(_def)
                }
        }

        return err
}

/*
ParseAttributeType parses an individual textual attribute
type (raw) and returns an error instance.

When no error occurs, the newly formed [AttributeType] instance
-- based on the parsed contents of raw -- is added to the receiver
[AttributeTypes] slice instance.
*/
func (r Schema) ParseAttributeType(raw string) error {
	def, err := parseAT(raw)
	if err == nil {
		var _def AttributeType
		if _def, err = r.marshalAT(def); err == nil {
			r.AttributeTypes().push(_def)
		}
	}

	return err
}

/*
ParseObjectClass parses an individual textual object class (raw) and
returns an error instance.

When no error occurs, the newly formed [ObjectClass] instance -- based
on the parsed contents of raw -- is added to the receiver [ObjectClasses]
instance.
*/
func (r Schema) ParseObjectClass(raw string) error {
	def, err := parseOC(raw)
	if err == nil {
		var _def ObjectClass
		if _def, err = r.marshalOC(def); err == nil {
			r.ObjectClasses().push(_def)
		}
	}

	return err
}

/*
ParseDITContentRule parses an individual textual attribute type (raw) and
returns an error instance.

When no error occurs, the newly formed [DITContentRule] instance -- based
on the parsed contents of raw -- is added to the receiver [DITContentRules]
slice instance.
*/
func (r Schema) ParseDITContentRule(raw string) error {
	def, err := parseDC(raw)
	if err == nil {
		var _def DITContentRule
		if _def, err = r.marshalDC(def); err == nil {
			r.DITContentRules().push(_def)
		}
	}

	return err
}

/*
ParseNameForm parses an individual textual attribute type (raw) and
returns an error instance.

When no error occurs, the newly formed [NameForm] instance -- based
on the parsed contents of raw -- is added to the receiver [NameForms]
slice instance.
*/
func (r Schema) ParseNameForm(raw string) error {
	def, err := parseNF(raw)
	if err == nil {
		var _def NameForm
		if _def, err = r.marshalNF(def); err == nil {
			r.NameForms().push(_def)
		}
	}

	return err
}

/*
ParseDITStructureRule parses an individual textual attribute type (raw) and
returns an error instance.

When no error occurs, the newly formed [DITStructureRule] instance -- based
on the parsed contents of raw -- is added to the receiver [DITStructureRules]
slice instance.
*/
func (r Schema) ParseDITStructureRule(raw string) error {
	def, err := parseDS(raw)
	if err == nil {
		var _def DITStructureRule
		if _def, err = r.marshalDS(def); err == nil {
			r.DITStructureRules().push(_def)
		}
	}

	return err
}

/*
incorporate returns an error following an attempt to marshal the contents
of s into r. The concept of "incorporation" is another term for post-parsing
"hand-off" from antlr4512 to schemax, and is the last of two (2) phases
involving the antlr4512 package.

Empty slice types within s shall not result in an error.
*/
func (r Schema) incorporate(s antlr4512.Schema) (err error) {
	for _, err := range []error{
		r.incorporateLS(s.LS),
		r.incorporateMR(s.MR),
		r.incorporateAT(s.AT),
		r.incorporateMU(s.MU),
		r.incorporateOC(s.OC),
		r.incorporateDC(s.DC),
		r.incorporateNF(s.NF),
		r.incorporateDS(s.DS),
	} {
		if err != nil {
			break
		}
	}

	return
}

func (r Schema) incorporateLS(s antlr4512.LDAPSyntaxes) (err error) {
	for i := 0; i < len(s); i++ {
		var def LDAPSyntax
		if def, err = r.marshalLS(s[i]); err != nil {
			break
		}

		r.LDAPSyntaxes().push(def)
	}

	return
}

func (r Schema) marshalLS(s antlr4512.LDAPSyntax) (def LDAPSyntax, err error) {
	// try to resolve a macro only if no
	// numeric OID is present.
	if len(s.Macro) == 2 && len(s.OID) == 0 {
		mc, found := r.GetMacro(s.Macro[0])
		if found {
			s.OID = mc + `.` + s.Macro[1]
		}
	}

	if !isNumericOID(s.OID) {
		err = errorf("Missing or invalid numeric OID for %T", def)
		return
	}

	// silently ignore attempts to marshal a duplicate definition.
	if lup := r.LDAPSyntaxes().get(s.OID); !lup.IsZero() {
		return
	}

	_def := newLDAPSyntax()
	_def.OID = s.OID
	_def.Desc = s.Desc
	_def.schema = r

	for k, v := range s.Extensions {
		_def.Extensions.Set(k, v...)
	}

	if err = _def.prepareString(); err == nil {
		def = LDAPSyntax{_def}
	}

	return
}

func (r Schema) incorporateMR(s antlr4512.MatchingRules) (err error) {

	for i := 0; i < len(s); i++ {
		var def MatchingRule
		if def, err = r.marshalMR(s[i]); err != nil {
			break
		}

		r.MatchingRules().push(def)
	}

	return
}

func (r Schema) marshalMR(s antlr4512.MatchingRule) (def MatchingRule, err error) {
	// try to resolve a macro only if no
	// numeric OID is present.
	if len(s.Macro) == 2 && len(s.OID) == 0 {
		mc, found := r.GetMacro(s.Macro[0])
		if found {
			s.OID = mc + `.` + s.Macro[1]
		}
	}

	if !isNumericOID(s.OID) {
		err = errorf("Missing or invalid numeric OID for %T", def)
		return
	} else if !isNumericOID(s.Syntax) {
		// try to resolve a quotedDescriptor to an OID
		// by scanning the DESC values of each syntax
		// within the receiver. Use of Whitespace and
		// case folding are not significant in the
		// matching process.
		if syn := r.LDAPSyntaxes().get(s.Syntax); !syn.IsZero() {
			// overwrite the DESC reference of syntax
			// with the intended syntax's numeric OID.
			s.Syntax = syn.NumericOID()
		}
	}

	if lup := r.MatchingRules().get(s.OID); !lup.IsZero() {
		// silently ignore attempts to marshal a duplicate definition.
		return
	}

	syn := r.LDAPSyntaxes().get(s.Syntax)
	if syn.IsZero() {
		// throw an error due to bad syntax ref
		err = errorf("Unknown %T OID '%s' for %T", syn, s.Syntax, def)
		return
	}

	_def := newMatchingRule()
	_def.OID = s.OID
	_def.Desc = s.Desc
	_def.Obsolete = s.Obsolete
	_def.schema = r
	_def.Syntax = syn

	for k, v := range s.Extensions {
		_def.Extensions.Set(k, v...)
	}

	for _, name := range s.Name {
		_def.Name.push(name)
	}

	if err = _def.prepareString(); err == nil {
		def = MatchingRule{_def}
	}

	return
}

func (r Schema) incorporateMU(s antlr4512.MatchingRuleUses) (err error) {

        for i := 0; i < len(s); i++ {
                var def MatchingRuleUse
                if def, err = r.marshalMU(s[i]); err != nil {
                        break
                }

                r.MatchingRuleUses().push(def)
        }

        return
}

func (r Schema) marshalMU(s antlr4512.MatchingRuleUse) (def MatchingRuleUse, err error) {
        if !isNumericOID(s.OID) {
                err = errorf("Missing or invalid numeric OID for %T", def)
                return
        }

        if lup := r.MatchingRules().get(s.OID); lup.IsZero() {
                // silently ignore attempts to reference a bogus matchingRule OID
                return
        }

        _def := newMatchingRuleUse()
        _def.OID = s.OID
        _def.Desc = s.Desc
        _def.Obsolete = s.Obsolete
        _def.schema = r

	for i := 0; i < len(s.Applies); i++ {
		_at := s.Applies[i]
		at := r.AttributeTypes().get(_at)
		if at.IsZero() {
			err = errorf("Applied attributeType '%s' not found", _at)
			return
		}
		_def.Applies.push(at)
	}

        for k, v := range s.Extensions {
                _def.Extensions.Set(k, v...)
        }

        for _, name := range s.Name {
                _def.Name.push(name)
        }

        if err = _def.prepareString(); err == nil {
                def = MatchingRuleUse{_def}
        }

        return
}

func (r Schema) incorporateAT(s antlr4512.AttributeTypes) (err error) {
	for i := 0; i < len(s); i++ {
		var def AttributeType
		if def, err = r.marshalAT(s[i]); err != nil {
			break
		}

		r.AttributeTypes().push(def)
	}

	return
}

func (r Schema) marshalAT(s antlr4512.AttributeType) (def AttributeType, err error) {
	// try to resolve a macro only if no
	// numeric OID is present.
	if len(s.Macro) == 2 && len(s.OID) == 0 {
		mc, found := r.GetMacro(s.Macro[0])
		if found {
			s.OID = mc + `.` + s.Macro[1]
		}
	}

	if !isNumericOID(s.OID) {
		err = errorf("Missing or invalid numeric OID for %T", def)
		return
	} else if lup := r.AttributeTypes().get(s.OID); !lup.IsZero() {
		// silently ignore attempts to marshal a duplicate definition.
		return
	}

	// alloc
	_def := newAttributeType()

	// Take any primitive values we can
	_def.OID = s.OID
	_def.Desc = s.Desc
	_def.Obsolete = s.Obsolete
	_def.Immutable = s.Immutable
	_def.Collective = s.Collective
	_def.Single = s.Single
	_def.MUB = s.MUB
	_def.schema = r

	// avoid bogus Boolean states
	if _def.Single && _def.Collective {
		err = errorf("%T '%s' is both single-valued and collective; aborting",
			s.OID, def)
		return
	}

	// verify and set an LDAPSyntax if specified
	if syn := s.Syntax; len(syn) > 0 {
		llup := r.LDAPSyntaxes().get(syn)
		if llup.IsZero() {
			// throw an error due to bad syntax ref
			err = errorf("Unknown %T OID '%s' for %T", llup, syn, def)
			return
		}
		_def.Syntax = llup
	}

	// verify any non-zero matchingrule references
	for idx, mrl := range []string{
		s.Equality,  // EQUALITY
		s.Substring, // SUBSTR
		s.Ordering,  // ORDERING
	} {
		if len(mrl) > 0 {
			// If mrl was non-zero, lookup using its
			// text value (name or OID). Fail if not
			// found, assign to appropriate field
			// otherwise.
			llup := r.MatchingRules().get(mrl)
			if llup.IsZero() {
				err = errorf("Unknown %T OID '%s' for %T", llup, mrl, def)
				return
			}

			switch idx {
			case 1:
				_def.Substring = llup
			case 2:
				_def.Ordering = llup
			default:
				_def.Equality = llup
			}
		}
	}

	// verify and set supertype, if specified.
	if _sup := s.SuperType; len(_sup) > 0 {
		sup := r.AttributeTypes().get(_sup)
		if sup.IsZero() {
			err = errorf("Unknown %T OID '%s' for %T", sup, _sup, def)
			return
		}
		_def.SuperType = sup
	}

	// verify and set usage if not default
	// "userApplications"
	if len(s.Usage) > 0 {
		switch lc(s.Usage) {
		case `directoryoperation`:
			_def.Usage = DirectoryOperationUsage
		case `distributedoperation`:
			_def.Usage = DistributedOperationUsage
		case `dsaoperation`:
			_def.Usage = DSAOperationUsage
		}
	}

	// Process and set any extensions
	for k, v := range s.Extensions {
		_def.Extensions.Set(k, v...)
	}

	// Process and set any names
	for _, name := range s.Name {
		_def.Name.push(name)
	}

	// save string representation and,
	// if no errors, release our newly
	// constructed instance.
	if err = _def.prepareString(); err == nil {
		def = AttributeType{_def}
	}

	return
}

func (r *Schema) incorporateOC(s antlr4512.ObjectClasses) (err error) {
	for i := 0; i < len(s); i++ {
		var def ObjectClass
		if def, err = r.marshalOC(s[i]); err != nil {
			break
		}

		r.ObjectClasses().push(def)
	}

	return
}

func (r *Schema) marshalOC(s antlr4512.ObjectClass) (oc ObjectClass, err error) {
	// try to resolve a macro only if no
	// numeric OID is present.
	if len(s.Macro) == 2 && len(s.OID) == 0 {
		mc, found := r.GetMacro(s.Macro[0])
		if found {
			s.OID = mc + `.` + s.Macro[1]
		}
	}

	if !isNumericOID(s.OID) {
		err = errorf("Missing or invalid numeric OID for %T", oc)
		return
	}

	// silently ignore attempts to marshal a duplicate definition.
	if lup := r.ObjectClasses().get(s.OID); !lup.IsZero() {
		return
	}

	_oc := newObjectClass()
	_oc.OID = s.OID
	_oc.Desc = s.Desc
	_oc.Obsolete = s.Obsolete

	// verify and set kind if not default of STRUCTURAL
	switch k := lc(s.Kind); k {
	case `structural`, ``:
		_oc.Kind = StructuralKind
	case `auxiliary`:
		_oc.Kind = AuxiliaryKind
	case `abstract`:
		_oc.Kind = AbstractKind
	default:
		err = errorf("Invalid kind '%s' for %T", k, oc)
		return
	}

	for _, must := range s.Must {
		m := r.AttributeTypes().get(must)
		if m.IsZero() {
			err = errorf("Unknown %T '%s' for %T.Must clause", m, must, oc)
			return
		}
		_oc.Must.push(m)
	}

	for _, may := range s.May {
		m := r.AttributeTypes().get(may)
		if m.IsZero() {
			err = errorf("Unknown %T '%s' for %T.May clause", m, may, oc)
			return
		}
		_oc.May.push(m)
	}

	for _, sup := range s.SuperClasses {
		m := r.ObjectClasses().get(sup)
		if m.IsZero() {
			err = errorf("Unknown %T '%s' for %T.SuperClasses clause", m, sup, oc)
			return
		}
		_oc.SuperClasses.push(m)
	}

	for k, v := range s.Extensions {
		_oc.Extensions.Set(k, v...)
	}

	// Process and set any names
	for _, name := range s.Name {
		_oc.Name.push(name)
	}

	if err = _oc.prepareString(); err == nil {
		oc = ObjectClass{_oc}
	}

	return
}

func (r *Schema) incorporateDC(s antlr4512.DITContentRules) (err error) {
	for i := 0; i < len(s); i++ {
		var def DITContentRule
		if def, err = r.marshalDC(s[i]); err != nil {
			break
		}

		r.DITContentRules().push(def)
	}

	return
}

func (r *Schema) marshalDC(s antlr4512.DITContentRule) (dc DITContentRule, err error) {
	if !isNumericOID(s.OID) {
		err = errorf("Missing or invalid numeric OID for %T", dc)
		return
	}

	_dc := newDITContentRule()

	soc := r.ObjectClasses().Get(s.OID)
	if soc.IsZero() {
		err = errorf("Unknown structural %T '%s' for %T.OID", soc, s.OID, dc)
		return
	}
	_dc.OID = soc

	_dc.Desc = s.Desc
	_dc.Obsolete = s.Obsolete

	for _, must := range s.Must {
		m := r.AttributeTypes().get(must)
		if m.IsZero() {
			err = errorf("Unknown %T '%s' for %T.Must clause", m, must, dc)
			return
		}
		_dc.Must.push(m)
	}

	for _, may := range s.May {
		m := r.AttributeTypes().get(may)
		if m.IsZero() {
			err = errorf("Unknown %T '%s' for %T.May clause", m, may, dc)
			return
		}
		_dc.May.push(m)
	}

	for _, not := range s.Not {
		m := r.AttributeTypes().get(not)
		if m.IsZero() {
			err = errorf("Unknown %T '%s' for %T.Not clause", m, not, dc)
			return
		}
		_dc.Not.push(m)
	}

	for _, aux := range s.Aux {
		m := r.ObjectClasses().get(aux)
		if m.IsZero() {
			err = errorf("Unknown %T '%s' for %T.Aux clause", m, aux, dc)
			return
		}
		_dc.Aux.push(m)
	}

	for k, v := range s.Extensions {
		_dc.Extensions.Set(k, v...)
	}

	for _, name := range s.Name {
		_dc.Name.push(name)
	}

	if err = _dc.prepareString(); err == nil {
		dc = DITContentRule{_dc}
	}

	return
}

func (r *Schema) incorporateNF(s antlr4512.NameForms) (err error) {
	for i := 0; i < len(s); i++ {
		var def NameForm
		if def, err = r.marshalNF(s[i]); err != nil {
			break
		}

		r.NameForms().push(def)
	}

	return
}

func (r *Schema) marshalNF(s antlr4512.NameForm) (nf NameForm, err error) {
	if !isNumericOID(s.OID) {
		err = errorf("Missing or invalid numeric OID for %T", nf)
		return
	}

	_nf := newNameForm()
	_nf.OID = s.OID
	_nf.Desc = s.Desc
	_nf.Obsolete = s.Obsolete

	oc := r.ObjectClasses().get(s.OC)
	if oc.IsZero() {
		err = errorf("Unknown structural %T '%s' for %T", oc, s.OC, nf)
		return
	}
	_nf.Structural = oc

	for _, must := range s.Must {
		m := r.AttributeTypes().get(must)
		if m.IsZero() {
			err = errorf("Unknown %T '%s' for %T.Must clause", m, must, nf)
			return
		}
		_nf.Must.push(m)
	}

	for _, may := range s.May {
		m := r.AttributeTypes().get(may)
		if m.IsZero() {
			err = errorf("Unknown %T '%s' for %T.May clause", m, may, nf)
			return
		}
		_nf.May.push(m)
	}

	for k, v := range s.Extensions {
		_nf.Extensions.Set(k, v...)
	}

	for _, name := range s.Name {
		_nf.Name.push(name)
	}

	if err = _nf.prepareString(); err == nil {
		nf = NameForm{_nf}
	}

	return
}

func (r *Schema) incorporateDS(s antlr4512.DITStructureRules) (err error) {
	for i := 0; i < len(s); i++ {
		var def DITStructureRule
		if def, err = r.marshalDS(s[i]); err != nil {
			break
		}

		r.DITStructureRules().push(def)
	}

	return
}

func (r *Schema) marshalDS(s antlr4512.DITStructureRule) (ds DITStructureRule, err error) {
	var ruleid int
	if ruleid, err = atoi(s.ID); err != nil {
		err = errorf("Invalid %T ruleid '%s'", ds, s.ID)
		return
	} else if ruleid < 0 {
		err = errorf("Invalid %T ruleid '%s'", ds, s.ID)
		return
	}

	_ds := newDITStructureRule()
	_ds.ID = uint(ruleid)
	_ds.Desc = s.Desc
	_ds.Obsolete = s.Obsolete

	nf := r.NameForms().get(s.Form)
	if nf.IsZero() {
		err = errorf("Unknown %T '%s' for %T", nf, s.Form, ds)
		return
	}
	_ds.Form = nf

	for _, sup := range s.SuperRules {
		m := r.DITStructureRules().get(sup)
		if m.IsZero() {
			err = errorf("Unknown %T '%s' for %T.SuperRules clause", m, sup, ds)
			return
		}
		_ds.SuperRules.push(m)
	}

	for k, v := range s.Extensions {
		_ds.Extensions.Set(k, v...)
	}

	for _, name := range s.Name {
		_ds.Name.push(name)
	}

	if err = _ds.prepareString(); err == nil {
		ds = DITStructureRule{_ds}
	}

	return
}

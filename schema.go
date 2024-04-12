package schemax

/*
schema.go centralizes all schema operations within a single construct.
*/

import (
	antlr4512 "github.com/JesseCoretta/go-rfc4512-antlr"
)

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

func (r Schema) resolveByMacro(def Definition) {
	if def.IsZero() {
		return
	}

	if mc := def.macro(); len(mc) == 2 {
		if o, found := r.GetMacro(mc[0]); found {
			ref := o + `.` + mc[1]
			if def.Type() == `dITContentRule` {
				if dcr := r.ObjectClasses().get(ref); !dcr.IsZero() {
					def.(DITContentRule).dITContentRule.OID = dcr
				}
			} else {
				def.setOID(ref)
			}
		}
	}
}

/*
SetMacro returns an error following an attempt to associate x with y.

x must be an RFC 4512-compliant descriptor, and y must be a legal numeric
OID.
*/
func (r Schema) SetMacro(x, y string) (err error) {
	if len(x) == 0 || len(y) == 0 {
		err = errorf("Descriptor and/or numeric OID are zero length")
		return
	}

	_m := r.cast().Auxiliary()[`macros`]
	m, _ := _m.(map[string]string)
	m[x] = y

	return
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
NewBasicSchema initializes and returns an instance of [Schema] completely
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
			`macros`: make(map[string]string, 0),
		}).
		Mutex().
		Push(
			NewLDAPSyntaxes(),       // 0
			NewMatchingRules(),      // 1
			NewAttributeTypes(),     // 2
			NewMatchingRuleUses(),   // 3
			NewObjectClasses(),      // 4
			NewDITContentRules(),    // 5
			NewNameForms(),          // 6
			NewDITStructureRules())) // 7
}

/*
String returns the string representation of the receiver instance.
*/
func (r Schema) String() string {
	return r.cast().String()
}

/*
LDAPSyntaxes returns the [LDAPSyntaxes] instance from within the
receiver instance.
*/
func (r Schema) LDAPSyntaxes() (lss LDAPSyntaxes) {
	slice, _ := r.cast().Index(ldapSyntaxesIndex)
	lss, _ = slice.(LDAPSyntaxes)
	return
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
AttributeTypes returns the [AttributeTypes] instance from within
the receiver instance.
*/
func (r Schema) AttributeTypes() (ats AttributeTypes) {
	slice, _ := r.cast().Index(attributeTypesIndex)
	ats, _ = slice.(AttributeTypes)
	return
}

/*
MatchingRuleUses returns the [MatchingRuleUses] instance from
within the receiver instance.
*/
func (r Schema) MatchingRuleUses() (mus MatchingRuleUses) {
	slice, _ := r.cast().Index(matchingRuleUsesIndex)
	mus, _ = slice.(MatchingRuleUses)
	return
}

/*
ObjectClasses returns the [ObjectClasses] instance from within
the receiver instance.
*/
func (r Schema) ObjectClasses() (ocs ObjectClasses) {
	slice, _ := r.cast().Index(objectClassesIndex)
	ocs, _ = slice.(ObjectClasses)
	return
}

/*
DITContentRules returns the [DITContentRules] instance from within
the receiver instance.
*/
func (r Schema) DITContentRules() (dcs DITContentRules) {
	slice, _ := r.cast().Index(dITContentRulesIndex)
	dcs, _ = slice.(DITContentRules)
	return
}

/*
Nameforms returns the [NameForms] instance from within the receiver
instance.
*/
func (r Schema) NameForms() (nfs NameForms) {
	slice, _ := r.cast().Index(nameFormsIndex)
	nfs, _ = slice.(NameForms)
	return
}

/*
DITStructureRules returns the [DITStructureRules] instance from
within the receiver instance.
*/
func (r Schema) DITStructureRules() (dss DITStructureRules) {
	slice, _ := r.cast().Index(dITStructureRulesIndex)
	dss, _ = slice.(DITStructureRules)
	return
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
	if r.IsZero() {
		err = errorf("%T instance is not initialized", r)
		return
	}

	var content []byte
	if content, err = readFile(file); err != nil {
		return
	}

	var i antlr4512.Instance
	if i, err = antlr4512.ParseInstance(string(content)); err != nil {
		return
	}

	defs := i.P.Definitions()
	if defs == nil {
		err = errorf("No definitions found in %T", defs)
		return
	}

	err = r.processSchemaDefinitions(defs)

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
	if r.IsZero() {
		err = errorf("%T instance is not initialized", r)
		return
	}

	var content []byte
	if content, err = readDirectory(dir); err != nil {
		return
	}

	var i antlr4512.Instance
	if i, err = antlr4512.ParseInstance(string(content)); err != nil {
		return
	}

	defs := i.P.Definitions()
	if defs == nil {
		err = errorf("No definitions found in %T", defs)
		return
	}

	err = r.processSchemaDefinitions(defs)

	return
}

/*
processSchemaDefinitions returns an error instance following an
attempt to parse all schema definitions found within the defs
IDefinitionsContext instance.
*/
func (r Schema) processSchemaDefinitions(defs antlr4512.IDefinitionsContext) (err error) {

	ct := defs.GetChildCount()
	for i := 0; i < ct; i++ {
		switch tv := defs.GetChild(i).(type) {
		case *antlr4512.LDAPSyntaxDescriptionsContext:
			err = r.processLSs(tv)
		case *antlr4512.MatchingRuleDescriptionsContext:
			err = r.processMRs(tv)
		case *antlr4512.AttributeTypeDescriptionsContext:
			err = r.processATs(tv)
		case *antlr4512.ObjectClassDescriptionsContext:
			err = r.processOCs(tv)
		case *antlr4512.DITContentRuleDescriptionsContext:
			err = r.processDCRs(tv)
		case *antlr4512.NameFormDescriptionsContext:
			err = r.processNFs(tv)
		case *antlr4512.DITStructureRuleDescriptionsContext:
			err = r.processDSRs(tv)
		}
	}

	if err == nil {
		err = r.updateMatchingRuleUses(r.AttributeTypes())
	}

	return
}

/*
processLSs returns an error instance following an attempt to parse all LDAPSyntax
definitions found in the input ls *LDAPSyntaxDescriptionsContext.
*/
func (r Schema) processLSs(ls *antlr4512.LDAPSyntaxDescriptionsContext) (err error) {
	_z := ls.AllLDAPSyntaxDescription()
	for j := 0; j < len(_z) && err == nil; j++ {
		var def LDAPSyntax
		if def, err = r.processLDAPSyntax(_z[j]); err != nil {
			break
		}
		err = r.LDAPSyntaxes().push(def)
	}

	return
}

/*
processMRs returns an error instance following an attempt to parse all MatchingRule
definitions found in the input mr *MatchingRuleDescriptionsContext.
*/
func (r Schema) processMRs(mr *antlr4512.MatchingRuleDescriptionsContext) (err error) {
	_z := mr.AllMatchingRuleDescription()
	for j := 0; j < len(_z); j++ {
		var def MatchingRule
		if def, err = r.processMatchingRule(_z[j]); err != nil {
			break
		}
		if err = r.MatchingRules().push(def); err != nil {
			break
		}
	}

	return
}

/*
processATs returns an error instance following an attempt to parse all AttributeType
definitions found within the input at *AttributeTypeDescriptionsContext.
*/
func (r Schema) processATs(at *antlr4512.AttributeTypeDescriptionsContext) (err error) {
	_z := at.AllAttributeTypeDescription()
	for j := 0; j < len(_z); j++ {
		var def AttributeType
		if def, err = r.processAttributeType(_z[j]); err != nil {
			break
		}
		if err = r.AttributeTypes().push(def); err != nil {
			break
		}
	}

	return
}

/*
processOCs returns an error instance following an attempt to parse all ObjectClass
definitions found within the input at *ObjectClassDescriptionsContext.
*/
func (r Schema) processOCs(oc *antlr4512.ObjectClassDescriptionsContext) (err error) {
	_z := oc.AllObjectClassDescription()
	for j := 0; j < len(_z); j++ {
		var def ObjectClass
		if def, err = r.processObjectClass(_z[j]); err != nil {
			break
		}
		if err = r.ObjectClasses().push(def); err != nil {
			break
		}
	}

	return
}

/*
processNFs returns an error instance following an attempt to parse all NameForm
definitions found within the input at *NameFormDescriptionsContext.
*/
func (r Schema) processNFs(nf *antlr4512.NameFormDescriptionsContext) (err error) {
	_z := nf.AllNameFormDescription()
	for j := 0; j < len(_z); j++ {
		var def NameForm
		if def, err = r.processNameForm(_z[j]); err != nil {
			break
		}
		if err = r.NameForms().push(def); err != nil {
			break
		}
	}

	return
}

/*
processDCRs returns an error instance following an attempt to parse all DITContentRule
definitions found within the input at *DITContentRuleDescriptionsContext.
*/
func (r Schema) processDCRs(dcr *antlr4512.DITContentRuleDescriptionsContext) (err error) {
	_z := dcr.AllDITContentRuleDescription()
	for j := 0; j < len(_z); j++ {
		var def DITContentRule
		def, err = r.processDITContentRule(_z[j])
		if err != nil {
			break
		}
		if err = r.DITContentRules().push(def); err != nil {
			break
		}
	}

	return
}

/*
processDSRs returns an error instance following an attempt to parse all DITStructureRule
definitions found within the input at *DITStructureRuleDescriptionsContext.
*/
func (r Schema) processDSRs(dsr *antlr4512.DITStructureRuleDescriptionsContext) (err error) {
	_z := dsr.AllDITStructureRuleDescription()
	for j := 0; j < len(_z); j++ {
		var def DITStructureRule
		def, err = r.processDITStructureRule(_z[j])
		if err != nil {
			break
		}
		if err = r.DITStructureRules().push(def); err != nil {
			break
		}
	}

	return
}

/*
Counters returns an instance of Counters containing the current number
of [Definition] instances currently within the receiver.
*/
func (r Schema) Counters() (c Counters) {
	if !r.IsZero() {
		c = Counters{
			LS: r.LDAPSyntaxes().len(),
			MR: r.MatchingRules().len(),
			AT: r.AttributeTypes().len(),
			MU: r.MatchingRuleUses().len(),
			OC: r.ObjectClasses().len(),
			DC: r.DITContentRules().len(),
			NF: r.NameForms().len(),
			DS: r.DITStructureRules().len(),
		}
	}

	return
}

/*
LoadAttributeTypes returns an error following an attempt to load all
package-included [AttributeType] slices into the receiver instance.
*/
func (r Schema) LoadAttributeTypes() Schema {
	_ = r.loadAttributeTypes()
	return r
}

/*
loadAttributeTypes returns an error following an attempt to load
all package-included [AttributeType] slices into the receiver instance.
*/
func (r Schema) loadAttributeTypes() (err error) {
	if !r.IsZero() {
		for _, funk := range []func() error{
			r.loadRFC2079AttributeTypes,
			r.loadRFC2798AttributeTypes,
			r.loadRFC3045AttributeTypes,
			r.loadRFC3672AttributeTypes,
			r.loadRFC4512AttributeTypes,
			r.loadRFC4519AttributeTypes,
			r.loadRFC3671AttributeTypes,
			r.loadRFC4523AttributeTypes,
			r.loadRFC4524AttributeTypes,
			r.loadRFC4530AttributeTypes,
		} {
			if err = funk(); err != nil {
				break
			}
		}
	}

	return
}

/*
LoadRFC2079AttributeTypes returns an error following an attempt to
load all RFC 2079 [AttributeType] slices into the receiver instance.
*/
func (r Schema) LoadRFC2079AttributeTypes() Schema {
	_ = r.loadRFC2079AttributeTypes()
	return r
}

func (r Schema) loadRFC2079AttributeTypes() (err error) {
	for i := 0; i < len(rfc2079AttributeTypes) && err == nil; i++ {
		at := rfc2079AttributeTypes[i]
		err = r.ParseAttributeType(string(at))
	}

	return
}

/*
LoadRFC2798AttributeTypes returns an error following an attempt to
load all RFC 2798 [AttributeType] slices into the receiver instance.
*/
func (r Schema) LoadRFC2798AttributeTypes() Schema {
	_ = r.loadRFC2798AttributeTypes()
	return r
}

func (r Schema) loadRFC2798AttributeTypes() (err error) {
	for i := 0; i < len(rfc2798AttributeTypes) && err == nil; i++ {
		at := rfc2798AttributeTypes[i]
		err = r.ParseAttributeType(string(at))
	}

	return
}

/*
LoadRFC3045AttributeTypes returns an error following an attempt to
load all RFC 3045 [AttributeType] slices into the receiver instance.
*/
func (r Schema) LoadRFC3045AttributeTypes() Schema {
	_ = r.loadRFC3045AttributeTypes()
	return r
}

func (r Schema) loadRFC3045AttributeTypes() (err error) {
	for i := 0; i < len(rfc3045AttributeTypes) && err == nil; i++ {
		at := rfc3045AttributeTypes[i]
		err = r.ParseAttributeType(string(at))
	}

	return
}

/*
LoadRFC3671AttributeTypes returns an error following an attempt to
load all RFC 3671 [AttributeType] slices into the receiver instance.
*/
func (r Schema) LoadRFC3671AttributeTypes() Schema {
	_ = r.loadRFC3671AttributeTypes()
	return r
}

func (r Schema) loadRFC3671AttributeTypes() (err error) {
	for i := 0; i < len(rfc3671AttributeTypes) && err == nil; i++ {
		at := rfc3671AttributeTypes[i]
		err = r.ParseAttributeType(string(at))
	}

	return
}

/*
LoadRFC3672AttributeTypes returns an error following an attempt to
load all RFC 3672 [AttributeType] slices into the receiver instance.
*/
func (r Schema) LoadRFC3672AttributeTypes() Schema {
	_ = r.loadRFC3672AttributeTypes()
	return r
}

func (r Schema) loadRFC3672AttributeTypes() (err error) {
	for i := 0; i < len(rfc3672AttributeTypes) && err == nil; i++ {
		at := rfc3672AttributeTypes[i]
		err = r.ParseAttributeType(string(at))
	}

	return
}

/*
LoadRFC4512AttributeTypes returns an error following an attempt to
load all RFC 4512 [AttributeType] slices into the receiver instance.
*/
func (r Schema) LoadRFC4512AttributeTypes() Schema {
	_ = r.loadRFC4512AttributeTypes()
	return r
}

func (r Schema) loadRFC4512AttributeTypes() (err error) {
	for i := 0; i < len(rfc4512AttributeTypes) && err == nil; i++ {
		at := rfc4512AttributeTypes[i]
		err = r.ParseAttributeType(string(at))
	}

	return
}

/*
LoadRFC4519AttributeTypes returns an error following an attempt to
load all RFC 4519 [AttributeType] slices into the receiver instance.
*/
func (r Schema) LoadRFC4519AttributeTypes() Schema {
	_ = r.loadRFC4519AttributeTypes()
	return r
}

func (r Schema) loadRFC4519AttributeTypes() (err error) {
	for i := 0; i < len(rfc4519AttributeTypes) && err == nil; i++ {
		at := rfc4519AttributeTypes[i]
		err = r.ParseAttributeType(string(at))
	}

	return
}

/*
LoadRFC4523AttributeTypes returns an error following an attempt to
load all RFC 4523 [AttributeType] slices into the receiver instance.
*/
func (r Schema) LoadRFC4523AttributeTypes() Schema {
	_ = r.loadRFC4523AttributeTypes()
	return r
}

func (r Schema) loadRFC4523AttributeTypes() (err error) {
	for i := 0; i < len(rfc4523AttributeTypes) && err == nil; i++ {
		at := rfc4523AttributeTypes[i]
		err = r.ParseAttributeType(string(at))
	}

	return
}

/*
LoadRFC4524AttributeTypes returns an error following an attempt to
load all RFC 4524 [AttributeType] slices into the receiver instance.
*/
func (r Schema) LoadRFC4524AttributeTypes() Schema {
	_ = r.loadRFC4524AttributeTypes()
	return r
}

func (r Schema) loadRFC4524AttributeTypes() (err error) {
	for i := 0; i < len(rfc4524AttributeTypes) && err == nil; i++ {
		at := rfc4524AttributeTypes[i]
		err = r.ParseAttributeType(string(at))
	}

	return
}

/*
LoadRFC4530AttributeTypes returns an error following an attempt to
load all RFC 4530 [AttributeType] slices into the receiver instance.
*/
func (r Schema) LoadRFC4530AttributeTypes() Schema {
	_ = r.loadRFC4530AttributeTypes()
	return r
}

func (r Schema) loadRFC4530AttributeTypes() (err error) {
	for i := 0; i < len(rfc4530AttributeTypes) && err == nil; i++ {
		at := rfc4530AttributeTypes[i]
		err = r.ParseAttributeType(string(at))
	}

	return
}

/*
LoadObjectClasses returns an error following an attempt to load all
package-included [ObjectClass] slices into the receiver instance.
*/
func (r Schema) LoadObjectClasses() Schema {
	_ = r.loadObjectClasses()
	return r
}

func (r Schema) loadObjectClasses() (err error) {
	if !r.IsZero() {
		for _, funk := range []func() error{
			r.loadRFC4512ObjectClasses,
			r.loadRFC4519ObjectClasses,
			r.loadRFC4523ObjectClasses,
			r.loadRFC4524ObjectClasses,
			r.loadRFC2079ObjectClasses,
			r.loadRFC2798ObjectClasses,
			r.loadRFC3671ObjectClasses,
			r.loadRFC3672ObjectClasses,
		} {
			if err = funk(); err != nil {
				break
			}
		}
	}

	return
}

/*
LoadRFC2079ObjectClasses returns an error following an attempt to
load all RFC 2079 [ObjectClass] slices into the receiver instance.
*/
func (r Schema) LoadRFC2079ObjectClasses() Schema {
	_ = r.loadRFC2079ObjectClasses()
	return r
}

func (r Schema) loadRFC2079ObjectClasses() (err error) {
	for i := 0; i < len(rfc2079ObjectClasses) && err == nil; i++ {
		oc := rfc2079ObjectClasses[i]
		err = r.ParseObjectClass(string(oc))
	}

	return
}

/*
LoadRFC2798ObjectClasses returns an error following an attempt to
load all RFC 2798 [ObjectClass] slices into the receiver instance.
*/
func (r Schema) LoadRFC2798ObjectClasses() Schema {
	_ = r.loadRFC2798ObjectClasses()
	return r
}

func (r Schema) loadRFC2798ObjectClasses() (err error) {
	for i := 0; i < len(rfc2798ObjectClasses) && err == nil; i++ {
		oc := rfc2798ObjectClasses[i]
		err = r.ParseObjectClass(string(oc))
	}

	return
}

/*
LoadRFC3671ObjectClasses returns an error following an attempt to
load all RFC 3671 [ObjectClass] slices into the receiver instance.
*/
func (r Schema) LoadRFC3671ObjectClasses() Schema {
	_ = r.loadRFC3671ObjectClasses()
	return r
}

func (r Schema) loadRFC3671ObjectClasses() (err error) {
	for i := 0; i < len(rfc3671ObjectClasses) && err == nil; i++ {
		oc := rfc3671ObjectClasses[i]
		err = r.ParseObjectClass(string(oc))
	}

	return
}

/*
LoadRFC3672ObjectClasses returns an error following an attempt to
load all RFC 3672 [ObjectClass] slices into the receiver instance.
*/
func (r Schema) LoadRFC3672ObjectClasses() Schema {
	_ = r.loadRFC3672ObjectClasses()
	return r
}

func (r Schema) loadRFC3672ObjectClasses() (err error) {
	for i := 0; i < len(rfc3672ObjectClasses) && err == nil; i++ {
		oc := rfc3672ObjectClasses[i]
		err = r.ParseObjectClass(string(oc))
	}

	return
}

/*
LoadRFC4512ObjectClasses returns an error following an attempt to
load all RFC 4512 [ObjectClass] slices into the receiver instance.
*/
func (r Schema) LoadRFC4512ObjectClasses() Schema {
	_ = r.loadRFC4512ObjectClasses()
	return r
}

/*
LoadRFC4530AttributeTypes returns an error following an attempt to
load all RFC 4530 [ObjectClass] slices into the receiver instance.
*/
func (r Schema) loadRFC4512ObjectClasses() (err error) {
	for i := 0; i < len(rfc4512ObjectClasses) && err == nil; i++ {
		oc := rfc4512ObjectClasses[i]
		err = r.ParseObjectClass(string(oc))
	}

	return
}

/*
LoadRFC4519ObjectClasses returns an error following an attempt to
load all RFC 4519 [ObjectClass] slices into the receiver instance.
*/
func (r Schema) LoadRFC4519ObjectClasses() Schema {
	_ = r.loadRFC4519ObjectClasses()
	return r
}

func (r Schema) loadRFC4519ObjectClasses() (err error) {
	for i := 0; i < len(rfc4519ObjectClasses) && err == nil; i++ {
		oc := rfc4519ObjectClasses[i]
		err = r.ParseObjectClass(string(oc))
	}

	return
}

/*
LoadRFC4523ObjectClasses returns an error following an attempt to
load all RFC 4523 [ObjectClass] slices into the receiver instance.
*/
func (r Schema) LoadRFC4523ObjectClasses() Schema {
	_ = r.loadRFC4523ObjectClasses()
	return r
}

func (r Schema) loadRFC4523ObjectClasses() (err error) {
	for i := 0; i < len(rfc4523ObjectClasses) && err == nil; i++ {
		oc := rfc4523ObjectClasses[i]
		err = r.ParseObjectClass(string(oc))
	}

	return
}

/*
LoadRFC4524ObjectClasses returns an error following an attempt to
load all RFC 4524 [ObjectClass] slices into the receiver instance.
*/
func (r Schema) LoadRFC4524ObjectClasses() Schema {
	_ = r.loadRFC4524ObjectClasses()
	return r
}

func (r Schema) loadRFC4524ObjectClasses() (err error) {
	for i := 0; i < len(rfc4524ObjectClasses) && err == nil; i++ {
		oc := rfc4524ObjectClasses[i]
		err = r.ParseObjectClass(string(oc))
	}

	return
}

/*
LoadLDAPSyntaxes will load all package-included [LDAPSyntax] definitions
into the receiver instance.
*/
func (r Schema) LoadLDAPSyntaxes() Schema {
	_ = r.loadSyntaxes()
	return r
}

/*
loadSyntaxes returns an error following an attempt to load all
ldapSyntax definitions found within this package into the receiver
instance.
*/
func (r Schema) loadSyntaxes() (err error) {
	if !r.IsZero() {
		for _, funk := range []func() error{
			r.loadRFC2307Syntaxes,
			r.loadRFC4517Syntaxes,
			r.loadRFC4523Syntaxes,
			r.loadRFC4530Syntaxes,
		} {
			if err = funk(); err != nil {
				break
			}
		}
	}

	return
}

/*
LoadRFC2307Syntaxes returns an error following an attempt to
load all RFC 2307 [LDAPSyntax] slices into the receiver instance.
*/
func (r Schema) LoadRFC2307Syntaxes() Schema {
	_ = r.loadRFC2307Syntaxes()
	return r
}

func (r Schema) loadRFC2307Syntaxes() (err error) {
	for k, v := range rfc2307Macros {
		r.SetMacro(k, v)
	}

	for i := 0; i < len(rfc2307Syntaxes) && err == nil; i++ {
		ls := rfc2307Syntaxes[i]
		err = r.ParseLDAPSyntax(string(ls))
	}

	return
}

/*
LoadRFC4517Syntaxes returns an error following an attempt to
load all RFC 4517 [LDAPSyntax] slices into the receiver instance.
*/
func (r Schema) LoadRFC4517Syntaxes() Schema {
	_ = r.loadRFC4517Syntaxes()
	return r
}

func (r Schema) loadRFC4517Syntaxes() (err error) {
	for i := 0; i < len(rfc4517Syntaxes) && err == nil; i++ {
		ls := rfc4517Syntaxes[i]
		err = r.ParseLDAPSyntax(string(ls))
	}

	return
}

/*
LoadRFC4523Syntaxes returns an error following an attempt to
load all RFC 4523 [LDAPSyntax] slices into the receiver instance.
*/
func (r Schema) LoadRFC4523Syntaxes() Schema {
	_ = r.loadRFC4523Syntaxes()
	return r
}

func (r Schema) loadRFC4523Syntaxes() (err error) {
	for i := 0; i < len(rfc4523Syntaxes) && err == nil; i++ {
		ls := rfc4523Syntaxes[i]
		err = r.ParseLDAPSyntax(string(ls))
	}

	return
}

/*
LoadRFC4530Syntaxes returns an error following an attempt to
load all RFC 4530 [LDAPSyntax] slices into the receiver instance.
*/
func (r Schema) LoadRFC4530Syntaxes() Schema {
	_ = r.loadRFC4530Syntaxes()
	return r
}

func (r Schema) loadRFC4530Syntaxes() (err error) {
	for i := 0; i < len(rfc4530Syntaxes) && err == nil; i++ {
		ls := rfc4530Syntaxes[i]
		err = r.ParseLDAPSyntax(string(ls))
	}

	return
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

package schemax

/*
fvindex returns the integer index for given field or value based on the input struct field name (e.g.: `OID`)
*/
func (def definition) lfindex(term string) (idx int) {
	return lfindex(term, def)
}

/*
definitionType returns the acronym describing the nature of the definition:

 - `at`  (AttributeType)
 - `oc`  (ObjectClass)
 - `ls`  (LDAPSyntax)
 - `mr`  (MatchingRule)
 - `nf`  (NameForm)
 - `mru` (MatchingRuleUse)
 - `dcr` (DITContentRule)
 - `dsr` (DITStructureRule)
*/
func (def definition) definitionType() (n string) {
	return definitionType(def)
}

func (def *definition) setKind(value string, idx int) (err error) {
	if def.alreadySet(idx) {
		return // silent discard
	}
	def.values[idx].Set(valueOf(newKind(value)))

	return
}

func (def *definition) setBoolean(label string, x interface{}) (err error) {
	switch tv := x.(type) {
	case *AttributeType:
		switch label {
		case SingleValue.string():
			tv.setBoolean(SingleValue)
			return
		case Collective.string():
			tv.setBoolean(Collective)
			return
		case Obsolete.string():
			tv.setBoolean(Obsolete)
			return
		case NoUserModification.string():
			tv.setBoolean(NoUserModification)
			return
		}
	case *ObjectClass:
		switch label {
		case Obsolete.string():
			tv.setBoolean(Obsolete)
			return
		}
	case *LDAPSyntax:
		switch label {
		case HumanReadable.string():
			tv.setBoolean(HumanReadable)
			return
		}
	case *MatchingRule:
		switch label {
		case Obsolete.string():
			tv.setBoolean(Obsolete)
			return
		}
	case *MatchingRuleUse:
		switch label {
		case Obsolete.string():
			tv.setBoolean(Obsolete)
			return
		}
	case *DITContentRule:
		switch label {
		case Obsolete.string():
			tv.setBoolean(Obsolete)
			return
		}
	case *DITStructureRule:
		switch label {
		case Obsolete.string():
			tv.setBoolean(Obsolete)
			return
		}
	case *NameForm:
		switch label {
		case Obsolete.string():
			tv.setBoolean(Obsolete)
			return
		}
	}
	err = raise(invalidBoolean,
		"setBoolean: unable to resolve '%T' type (label:'%s')",
		x, label)

	return
}

func (def *definition) setExtensions(label string, value []string, idx int) (err error) {
	z, ok := def.values[idx].Interface().(Extensions)
	if !ok {
		return raise(unknownDefinition,
			"setExtensions: unexpected type '%T'",
			def.values[idx].Interface())
	} else if z.IsZero() {
		z = NewExtensions()
	}

	z.Set(label, value)
	def.values[idx].Set(valueOf(z))

	return nil
}

func (def *definition) setUsage(value string, idx int) (err error) {
	def.values[idx].Set(valueOf(newUsage(value)))

	return
}

func (def *definition) setDesc(idx int, value string) (err error) {
	if def.alreadySet(idx) {
		return // silent discard
	}
	def.values[idx].Set(valueOf(Description(value)))

	return
}

func (def *definition) setName(idx int, value ...string) (err error) {
	if def.alreadySet(idx) {
		return // silent discard
	}

	if name := NewName(value...); !name.IsZero() {
		def.values[idx].Set(valueOf(name))
	}

	return
}

func (def *definition) setStructuralObjectClass(
	ocm ObjectClassesManifest,
	x interface{},
	value string,
	idx int) (err error) {

	if def.alreadySet(idx) {
		return
	}

	oc := ocm.Get(value, def.alm) // TODO deref alias / name
	if oc.IsZero() {
		return raise(invalidObjectClass,
			"setStructuralObjectClass: no such %T was found in %T for value '%s' (type: %T)",
			oc, value)
	} else if !oc.Kind.is(Structural) {
		//return raise(invalidObjectClass,
		//      "setStructuralObjectClass: %T (%s) is not STRUCTURAL kind",
		//	oc, oc.OID.string())
	}
	def.values[idx].Set(valueOf(StructuralObjectClass{oc}))

	return
}

func (def *definition) setSyntax(
	lsm LDAPSyntaxesManifest,
	x interface{},
	value string,
	idx int) (err error) {

	switch def.definitionType() {
	case `mr`:
		ls := lsm.Get(value, def.alm)
		if ls.IsZero() {
			return raise(invalidSyntax,
				"setSyntax: no %T in %T for '%s' (type: %T)",
				ls, lsm, value, x)
		}
		def.values[idx].Set(valueOf(ls))
	case `at`:
		err = def.setAttrTypeSyntax(lsm, x, value, idx)
	}

	return
}

func (def *definition) setAttrTypeSyntax(
	lsm LDAPSyntaxesManifest,
	x interface{},
	value string, idx int) (err error) {

	assert, ok := x.(*AttributeType)
	if !ok {
		return raise(unknownDefinition,
			"setAttrTypeSyntax: unexpected type '%T'", x)
	}

	// Check (and handle) a MUB in the event we're using one ...
	if mub, _, ok := parse_mub(value); ok {
		// MUB detected
		assert.setMUB(mub[1])
		if ls := lsm.Get(mub[0], def.alm); !ls.IsZero() {
			def.values[idx].Set(valueOf(ls))
		} else {
			err = raise(invalidSyntax,
				"setAttrTypeSyntax: no %T in %T for '%s' (type: %T, mub:true)",
				ls, lsm, mub[0], x)
		}
	} else {
		// No MUB detected
		if ls := lsm.Get(value, def.alm); !ls.IsZero() {
			def.values[idx].Set(valueOf(ls))
		} else {
			err = raise(invalidSyntax,
				"setAttrTypeSyntax: no %T in %T for '%s' (type: %T)",
				ls, lsm, value, x)
		}
	}

	return
}

func (def *definition) setNameForm(nfm NameFormsManifest, x interface{},
	value string, idx int) (err error) {

	nf := nfm.Get(value, def.alm) // TODO deref alias / name
	if nf.IsZero() {
		return raise(unknownElement,
			"setNameForm: no such %T was found in %T for value '%s' (type: %T)",
			nf, nfm, value, x)
	}
	def.values[idx].Set(valueOf(nf))

	return
}

func (def *definition) setSuperiorObjectClasses(
	ocm ObjectClassesManifest,
	x interface{},
	value []string,
	idx int) (err error) {

	z, ok := def.values[idx].Interface().(SuperiorObjectClasses)
	if !ok {
		return raise(unknownDefinition,
			"setSuperiorObjectClasses: unexpected type '%T'",
			def.values[idx].Interface())
	} else if z.IsZero() {
		z = NewSuperiorObjectClasses()
	}

	for _, v := range value {
		oc := ocm.Get(v, def.alm)
		if oc.IsZero() {
			return raiseUnknownElement(`setSuperiorObjectClasses`,
				oc, ocm, v, x)
		}

		z.Set(oc)
		def.values[idx].Set(valueOf(z))
	}

	return
}

func (def *definition) setAuxiliaryObjectClasses(
	ocm ObjectClassesManifest,
	x interface{},
	value []string,
	idx int) (err error) {

	z, ok := def.values[idx].Interface().(AuxiliaryObjectClasses)
	if !ok {
		return raise(unknownDefinition,
			"setAuxiliaryObjectClasses: unexpected type '%T'",
			def.values[idx].Interface())
	} else if z.IsZero() {
		z = NewAuxiliaryObjectClasses()
	}

	for _, v := range value {
		oc := ocm.Get(v, def.alm)
		if oc.IsZero() {
			return raiseUnknownElement(`setAuxiliaryObjectClasses`,
				oc, ocm, v, x)
		}

		z.Set(oc)
		def.values[idx].Set(valueOf(z))
	}

	return
}

func (def *definition) setApplies(
	atm AttributeTypesManifest,
	x interface{},
	value []string,
	idx int) (err error) {

	z, ok := def.values[idx].Interface().(Applies)
	if !ok {
		return raise(unknownDefinition,
			"setApplies: unexpected type '%T'",
			def.values[idx].Interface())
	} else if z.IsZero() {
		z = NewApplies()
	}

	for _, v := range value {
		at := atm.Get(v, def.alm)
		if at.IsZero() {
			return raiseUnknownElement(`setApplies`,
				at, atm, v, x)
		}

		z.Set(at)
		def.values[idx].Set(valueOf(z))
	}

	return
}

func (def *definition) setRequiredAttributeTypes(
	atm AttributeTypesManifest,
	x interface{},
	value []string,
	idx int) (err error) {

	z, ok := def.values[idx].Interface().(RequiredAttributeTypes)
	if !ok {
		return raise(unknownDefinition,
			"setRequiredAttributeTypes: unexpected type '%T'",
			def.values[idx].Interface())
	} else if z.IsZero() {
		z = NewRequiredAttributeTypes()
	}

	for _, v := range value {
		at := atm.Get(v, def.alm)
		if at.IsZero() {
			return raiseUnknownElement(`setRequiredAttributeTypes`,
				at, atm, v, x)
		}

		z.Set(at)
		def.values[idx].Set(valueOf(z))
	}

	return
}

func (def *definition) setPermittedAttributeTypes(
	atm AttributeTypesManifest,
	x interface{},
	value []string,
	idx int) (err error) {

	z, ok := def.values[idx].Interface().(PermittedAttributeTypes)
	if !ok {
		return raise(unknownDefinition,
			"setPermittedAttributeTypes: unexpected type '%T'",
			def.values[idx].Interface())
	} else if z.IsZero() {
		z = NewPermittedAttributeTypes()
	}

	for _, v := range value {
		at := atm.Get(v, def.alm)
		if at.IsZero() {
			return raiseUnknownElement(`setPermittedAttributeTypes`,
				at, atm, v, x)
		}

		z.Set(at)
		def.values[idx].Set(valueOf(z))
	}

	return
}

func (def *definition) setProhibitedAttributeTypes(
	atm AttributeTypesManifest,
	x interface{},
	value []string,
	idx int) (err error) {

	z, ok := def.values[idx].Interface().(ProhibitedAttributeTypes)
	if !ok {
		return raise(unknownDefinition,
			"setProhibitedAttributeTypes: unexpected type '%T'",
			def.values[idx].Interface())
	} else if z.IsZero() {
		z = NewProhibitedAttributeTypes()
	}

	for _, v := range value {
		at := atm.Get(v, def.alm)
		if at.IsZero() {
			return raiseUnknownElement(`setProhibitedAttributeTypes`,
				at, atm, v, x)
		}

		z.Set(at)
		def.values[idx].Set(valueOf(z))
	}

	return
}

func (def *definition) setSuperiorAttributeType(
	atm AttributeTypesManifest,
	x interface{},
	value string,
	idx int) (err error) {

	if def.alreadySet(idx) {
		return // silent discard
	}

	at := atm.Get(value, def.alm)
	if at.IsZero() {
		return raiseUnknownElement(`setSuperiorAttributeType`,
			at, atm, value, x)
	}
	def.values[idx].Set(valueOf(SuperiorAttributeType{at}))

	return nil
}

/*
setSuperiorDITStructureRules sets the SUP value of argument x, given the string value argument. This process fails if the manifest lookup fails.
*/
func (def *definition) setSuperiorDITStructureRules(
	dsrm DITStructureRulesManifest,
	x interface{},
	value []string,
	idx int) (err error) {

	z, ok := def.values[idx].Interface().(SuperiorDITStructureRules)
	if !ok {
		return raise(unknownDefinition,
			"setSuperiorDITStructureRules: unexpected type '%T'",
			def.values[idx].Interface())
	} else if z.IsZero() {
		z = NewSuperiorDITStructureRules()
	}

	for _, v := range value {
		if dsr := dsrm.Get(v); !dsr.IsZero() {
			z.Set(dsr)
			def.values[idx].Set(valueOf(z))
		} else {
			err = raiseUnknownElement(
				`setSuperiorDITStructureRules`,
				dsr, dsrm, v, x)
			break
		}
	}

	return
}

/*
setEqSubOrd sets the SUBSTR, ORDERING or EQUALITY value of argument x, given the string value argument. This process fails if the manifest lookup fails.
*/
func (def *definition) setEqSubOrd(
	mrm MatchingRulesManifest,
	x interface{},
	value string,
	idx int) (err error) {

	if def.alreadySet(idx) {
		return // silent discard
	}

	assert, ok := x.(*AttributeType)
	if !ok {
		return raise(unknownDefinition,
			"setEqSubOrd: unexpected type '%T'", x)
	} else if assert.IsZero() {
		return raise(emptyDefinition,
			"setEqSubOrd: received empty '%T'", assert)
	}

	mr := mrm.Get(value, def.alm)
	if mr.IsZero() {
		return raiseUnknownElement(
			`setEqSubOrd`,
			mr, mrm, value, x)
	}

	if err = def.setMR(assert, mr, idx); err != nil {
		err = raise(err,
			"setMR: def.labels[%d] ('%s') not set (type:'%T',value:'%s')",
			idx, def.labels[idx], x, value)
	}

	return
}

func (def *definition) setMR(
	dest *AttributeType,
	mr *MatchingRule,
	idx int) (err error) {

	// !! It is crucial we recognize when a supertype is in effect !!
	switch dest.SuperType.IsZero() {
	case true:
		switch def.labels[idx] {
		case `EQUALITY`:
			def.values[idx].Set(valueOf(Equality{mr}))
		case `SUBSTR`:
			def.values[idx].Set(valueOf(Substring{mr}))
		case `ORDERING`:
			def.values[idx].Set(valueOf(Ordering{mr}))
		default:
			err = unknownElement
		}
	case false:
		switch def.labels[idx] {
		case `EQUALITY`:
			def.values[idx].Set(valueOf(dest.SuperType.Equality))
		case `SUBSTR`:
			def.values[idx].Set(valueOf(dest.SuperType.Substring))
		case `ORDERING`:
			def.values[idx].Set(valueOf(dest.SuperType.Ordering))
		default:
			err = unknownElement
		}
	}

	return
}

func (def *definition) alreadySet(idx int) (isSet bool) {
	if idx < 0 {
		return
	}

	switch tv := def.values[idx].Interface().(type) {
	case Kind:
		isSet = !tv.IsZero()
	case Name:
		isSet = !tv.IsZero()
	case Boolean:
		isSet = !tv.IsZero()
	case Description:
		isSet = !tv.IsZero()
	case *AttributeType:
		isSet = !tv.IsZero()
	case RequiredAttributeTypes:
		isSet = !tv.IsZero()
	case PermittedAttributeTypes:
		isSet = !tv.IsZero()
	case ProhibitedAttributeTypes:
		isSet = !tv.IsZero()
	case SuperiorAttributeType:
		isSet = !tv.IsZero()
	case SuperiorObjectClasses:
		isSet = !tv.IsZero()
	case *ObjectClass:
		isSet = !tv.IsZero()
	case StructuralObjectClass:
		isSet = !tv.IsZero()
	case AuxiliaryObjectClasses:
		isSet = !tv.IsZero()
	case *LDAPSyntax:
		isSet = !tv.IsZero()
	case *MatchingRule:
		isSet = !tv.IsZero()
	case *MatchingRuleUse:
		isSet = !tv.IsZero()
	case Applies:
		isSet = !tv.IsZero()
	case Equality:
		isSet = !tv.IsZero()
	case Ordering:
		isSet = !tv.IsZero()
	case Substring:
		isSet = !tv.IsZero()
	case *NameForm:
		isSet = !tv.IsZero()
	case *DITStructureRule:
		isSet = !tv.IsZero()
	case *SuperiorDITStructureRules:
		isSet = !tv.IsZero()
	case *DITContentRule:
		isSet = !tv.IsZero()
	default:
		isSet = true // better safe than sorry ...
	}

	return
}

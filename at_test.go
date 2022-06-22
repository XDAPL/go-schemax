package schemax

import (
	"testing"
)

/*
Assemble an AttributeType manually (NO parsing) and evaluate collective misuse.
*/
func TestNoUserModificationAttributeType001(t *testing.T) {
	schema := testSubschema()

	mr := schema.GetMatchingRule(`generalizedTimeMatch`)
	ord := schema.GetMatchingRule(`generalizedTimeOrderingMatch`)
	syn := schema.GetLDAPSyntax(`1.3.6.1.4.1.1466.115.121.1.24`)

	if mr.IsZero() || syn.IsZero() {
		t.Errorf("%s failed: lookup error for supporting LDAPSyntax and/or MatchingRule definition values", t.Name())
		return
	}

	def := NewAttributeType()
	def.OID = NewOID(`2.5.18.2`)
	def.Name = NewName(`modifyTimestamp`)
	def.Syntax = syn
	def.Equality = Equality{mr}
	def.Ordering = Ordering{ord}
	def.SetSingleValue()
	def.SetNoUserModification()
	def.Usage = DirectoryOperation
	def.Extensions.Set(`X-ORIGIN`, `RFC4512`)

	um, err := Unmarshal(def)
	if err != nil {
		t.Errorf("%s failed: %s", t.Name(), err.Error())
		return
	}

	want := 218
	got := len(um)

	if want != got {
		t.Errorf("%s failed: unexpected raw length (want %d, got %d)", t.Name(), want, got)
	}
}

/*
Assemble an AttributeType manually (NO parsing) and evaluate collective misuse.
*/
func TestCollectiveAttributeType001(t *testing.T) {
	schema := testSubschema()

	mr := schema.GetMatchingRule(`caseIgnoreMatch`)
	syn := schema.GetLDAPSyntax(`1.3.6.1.4.1.1466.115.121.1.15`)

	if mr.IsZero() || syn.IsZero() {
		t.Errorf("%s failed: lookup error for supporting LDAPSyntax and/or MatchingRule definition values", t.Name())
		return
	}

	def := NewAttributeType()
	def.OID = NewOID(`1.3.6.1.4.1.56521.999.1.2.3.1`)
	def.Name = NewName(`genericTestAttribute`, `attrAltName`)
	def.Syntax = syn
	def.Equality = Equality{mr}
	def.Description = Description(`Generic test attribute`)
	def.SetCollective()

	um, err := Unmarshal(def)
	if err != nil {
		t.Errorf("%s failed: %s", t.Name(), err.Error())
		return
	}

	want := 182
	got := len(um)

	if want != got {
		t.Errorf("%s failed: unexpected raw length (want %d, got %d)", t.Name(), want, got)
	}
}

func TestAttributeTypeUsage001(t *testing.T) {
	schema := testSubschema()
	schema.PopulateDefaultAttributeTypes()
	name := schema.GetAttributeType(`name`)

	name.Usage = UserApplication

	if len(name.Usage.String()) != 0 {
		t.Errorf("%s error: USAGE value should be zero length (for userApplication), got `%s`", t.Name(), name.Usage)
	}
}

/*
Assemble an AttributeType manually and evaluate unmarshal reliability.
*/
func TestCompositeAttributeType001(t *testing.T) {
	schema := testSubschema()

	mr := schema.GetMatchingRule(`caseIgnoreMatch`)
	syn := schema.GetLDAPSyntax(`1.3.6.1.4.1.1466.115.121.1.15`)

	if mr.IsZero() || syn.IsZero() {
		t.Errorf("%s failed: lookup error for supporting LDAPSyntax and/or MatchingRule definition values", t.Name())
		return
	}

	def := new(AttributeType)
	def.OID = NewOID(`1.3.6.1.4.1.56521.999.1.2.3.1`)
	def.Name = NewName(`genericTestAttribute`, `attrAltName`)
	def.Syntax = syn
	def.Equality = Equality{mr}
	def.Description = Description(`Generic test attribute`)

	if err := def.Validate(); err != nil {
		t.Errorf("%s validation failed: %s", t.Name(), err.Error())
		return
	}

	raw, err := Unmarshal(def)
	if err != nil {
		t.Errorf("%s unmarshal failed: %s", t.Name(), err.Error())
		return
	}

	want := 171
	got := len(raw)

	if want != got {
		t.Errorf("%s failed: unexpected raw length (want %d, got %d)", t.Name(), want, got)
	}
}

func TestAttributeTypeMUB001(t *testing.T) {
	schema := testSubschema()

	def := `( 1.3.6.1.4.1.56521.999.100.2.1.13 NAME 'minUpperBoundTestAttr' DESC 'Generic attribute of limited value length' SYNTAX 1.3.6.1.4.1.1466.115.121.1.15{32} EQUALITY caseIgnoreMatch )`

	if err := schema.MarshalAttributeType(def); err != nil {
		t.Errorf("TestAttributeMUB marshal failed: %s", err.Error())
		return
	}

	um := schema.GetAttributeType(`minUpperBoundTestAttr`)
	if um.IsZero() {
		t.Errorf("TestAttributeMUB failed: test attribute not successfully registered in collection")
		return
	}

	mub, ok := um.Map()[`MUB`]
	if !ok {
		t.Errorf("TestAttributeMUB failed: length specifier not successfully retained or unmarshaled")
		return
	}

	switch len(mub) {
	case 1:
		got := mub[0]
		want := `32`
		if got != want {
			t.Errorf("TestAttributeMUB failed: MUB mismatch (want %s, got %s)", want, got)
		}
	default:
		want := 1
		got := len(mub)
		t.Errorf("TestAttributeMUB failed: unexpected number of MUB slices (want %d, got %d)", want, got)
	}
}

func TestParseAttributeType001(t *testing.T) {
	def := `( 1.3.6.1.4.1.56521.999.100.2.1.13 NAME 'testAttr' DESC 'Generic test attribute' OBSOLETE SYNTAX 1.3.6.1.4.1.1466.115.121.1.15 EQUALITY caseExactMatch )`

	var lsc LDAPSyntaxCollection = PopulateDefaultLDAPSyntaxes()
	var mrc MatchingRuleCollection = PopulateDefaultMatchingRules()
	var atc AttributeTypeCollection = NewAttributeTypes()
	var x AttributeType

	err := Marshal(def, &x, atc, nil, lsc, mrc, nil, nil, nil, nil)
	if err != nil {
		t.Errorf("%s failed: %s\n", t.Name(), err.Error())
		return
	}

	var um string
	if um, err = Unmarshal(&x); err != nil {
		t.Errorf("%s failed: %s", t.Name(), err.Error())
		return
	}

	// What went in should match
	// what comes out.
	want := len(def)
	got := len(um)

	if want != got {
		t.Errorf("%s failed: unexpected raw length (want %d, got %d)", t.Name(), want, got)
	}
}

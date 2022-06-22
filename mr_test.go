package schemax

import (
	"testing"
)

/*
Assemble a MatchingRule manually and evaluate unmarshal reliability.
*/
func TestCompositeMatchingRule001(t *testing.T) {
	dstr := NewLDAPSyntax()
	dstr.OID = NewOID(`1.3.6.1.4.1.1466.115.121.1.15`)
	dstr.Description = Description(`Directory String`)
	dstr.Extensions.Set(`X-ORIGIN`, `RFC4517`)

	def := NewMatchingRule()
	def.Name = NewName(`caseIgnoreMatch`)
	def.OID = NewOID(`2.5.13.2`)
	def.Syntax = dstr
	def.Extensions.Set(`X-ORIGIN`, `RFC4517`)
	def.Obsolete = true

	// Make sure validation checks return NO
	// errors before using your definition!
	um, err := Unmarshal(def)
	if err != nil {
		t.Errorf("%s failed: %s", t.Name(), err.Error())
		return
	}

	want := 100
	got := len(um)

	if want != got {
		t.Errorf("%s failed: unexpected raw length (want %d, got %d)", t.Name(), want, got)
	}
}

func TestParseMatchingRule001(t *testing.T) {
	def := `( 2.5.13.2 NAME 'caseIgnoreMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.15 X-ORIGIN 'RFC4517' )`

	var lsc LDAPSyntaxCollection = PopulateDefaultLDAPSyntaxes()
	var x MatchingRule

	err := Marshal(def, &x, nil, nil, lsc, nil, nil, nil, nil, nil)
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

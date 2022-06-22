package schemax

import (
	"testing"
)

/*
Assemble an LDAPSyntax manually and evaluate unmarshal reliability.
*/
func TestCompositeLDAPSyntax001(t *testing.T) {
	def := NewLDAPSyntax()

	// Let's recreate a well-known
	// syntax from RFC4517 piece by
	// piece.
	def.OID = NewOID(`1.3.6.1.4.1.1466.115.121.1.4`)
	def.Description = Description(`Audio`)
	def.Extensions.Set(`X-ORIGIN`, `RFC4517`)
	def.SetHumanReadable(false)

	// Make sure validation checks return NO
	// errors before using your definition!
	um, err := Unmarshal(def)
	if err != nil {
		t.Errorf("%s failed: %s", t.Name(), err.Error())
		return
	}

	want := 92
	got := len(um)

	if want != got {
		t.Errorf("%s failed: unexpected raw length (want %d, got %d)", t.Name(), want, got)
	}
}

func TestParseLDAPSyntax001(t *testing.T) {
	def := `( 1.3.6.1.4.1.1455.115.121.1.4 DESC 'Audio' X-ORIGIN 'RFC4517' X-NOT-HUMAN-READABLE 'TRUE' )`

	var x LDAPSyntax
	err := Marshal(def, &x, nil, nil, nil, nil, nil, nil, nil, nil)
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

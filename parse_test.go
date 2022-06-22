package schemax

import "testing"

/*
parse_test.go is for internal use
*/

func TestQdescrs(t *testing.T) {
	sv1 := "'name' REST-OF-DEF )"
	sv2 := `( 'RFC4517' ) REST-OF-DEF )`
	mv1 := `( 'RFC1234' 'RFC5678' ) REST-OF-DEF )`

	for _, x := range []string{sv1, sv2, mv1} {
		_, _, ok := parse_qdescrs(x)
		if !ok {
			t.Errorf("%s failed", t.Name())
			return
		}
	}
}

func TestOIDsIDs(t *testing.T) {
	sv1 := "name ) REST-OF-DEF )"
	sv2 := "0 REST-OF-DEF )" // dsr
	sv3 := `( name ) REST-OF-DEF )`
	mv1 := `( name $ l $ c ) REST-OF-DEF )`
	mv2 := `( 0 $ 1 ) REST-OF-DEF )` // dsr

	for _, x := range []string{sv1, sv2, sv3, mv1, mv2} {
		_, _, ok := parse_oids_ids(x)
		if !ok {
			t.Errorf("%s failed", t.Name())
			return
		}
	}
}

func TestStripTags(t *testing.T) {
	tagged := `name;lang-fr`
	stripped := stripTags(tagged)

	if stripped != `name` {
		t.Errorf("%s failed: tag strip unsuccessful", t.Name())
	}
}

func TestParse(t *testing.T) {
	raw := `( 1.3.6.1.4.1.56521.999.1.2 NAME 'testAttr' SYNTAX 1.3.6.1.4.1.1466.1.2.3.4 X-ORIGIN 'RFC9999' )`

	_, rest, ok := parse(raw)
	if !ok {
		t.Errorf("%s failed", t.Name())
		return
	}

	want := 67
	got := len(rest)

	if want != got {
		t.Errorf("%s failed: unexpected raw length (want %d, got %d)", t.Name(), want, got)
	}
}

func TestIsNumericalOID(t *testing.T) {
	val := `1,3,6.1.4.1`
	if isNumericalOID(val) {
		t.Errorf("%s failed: bad OID accepted as valid", t.Name())
		return
	}

	val = `3.3.6.1.4.1`
	if isNumericalOID(val) {
		t.Errorf("%s failed: bad OID accepted as valid", t.Name())
		return
	}

	val = `2`
	if isNumericalOID(val) {
		t.Errorf("%s failed: bad OID accepted as valid", t.Name())
		return
	}

	val = `2.1`
	if !isNumericalOID(val) {
		t.Errorf("%s failed: valid OID denied", t.Name())
	}
}

func TestParseMub(t *testing.T) {
	val := `1.3.6.1.4.1.1466.1.2.3.4{128} X-ORIGIN 'RFC9999' )`
	mub, _, ok := parse_mub(val)
	if !ok {
		t.Errorf("%s failed", t.Name())
		return
	}

	want := `128`
	got := mub[1]

	if want != got {
		t.Errorf("%s failed: MUB invalid (want %s, got %s)", t.Name(), want, got)
	}
}

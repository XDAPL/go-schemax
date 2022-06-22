package schemax

import "testing"

/*
collection_test.go is for internal use
*/

func TestCollectionIndex001(t *testing.T) {
	// this test focuses on panic-free
	// index calls. Some of these tests
	// use negative integers

	// Just use *Extension types for
	// test simplicity. Any supported
	// type would do ...
	c := collection{
		&Extension{
			Label: `X-TEST1`,
			Value: []string{`VALUE1`, `VALUE2`},
		},
		&Extension{
			Label: `X-TEST2`,
			Value: []string{`VALUE1`, `VALUE2`},
		},
		&Extension{
			Label: `X-TEST3`,
			Value: []string{`VALUE1`, `VALUE2`},
		},
	}

	// note some of these tests are
	// intentionally absurd. Don't
	// read into them too much ...
	tests := map[string]string{
		c.index(0).(*Extension).Label:   `X-TEST1`,
		c.index(1).(*Extension).Label:   `X-TEST2`,
		c.index(-1).(*Extension).Label:  `X-TEST3`,
		c.index(-2).(*Extension).Label:  `X-TEST2`,
		c.index(-3).(*Extension).Label:  `X-TEST1`,
		c.index(-15).(*Extension).Label: `X-TEST1`,
		c.index(101).(*Extension).Label: `X-TEST3`,
	}

	for want, got := range tests {
		if want != got {
			t.Errorf("%s validation failed: key mismatch (want %s, got %s)", t.Name(), want, got)
			return
		}
	}
}

func TestCollectionAppend001(t *testing.T) {
	c := make(collection, 0)
	bogus := struct {
		Key   string
		Value []string
	}{
		Key:   `bogusInstance`,
		Value: []string{`VALUE1`, `VALUE2`},
	}

	c.append(bogus)
	if c.len() != 0 {
		t.Errorf("%s validation failed: an instance of an unsupported type was accepted in append", t.Name())
		return
	}

	c.append(&Extension{
		Label: `X-TEST-PARAM`,
		Value: []string{`PARAM VALUE`},
	})

	if c.len() != 1 {
		t.Errorf("%s validation failed: an instance of a supported type was not accepted in append", t.Name())
		return
	}

	ext := c.index(0).(*Extension)

	want := `X-TEST-PARAM 'PARAM VALUE'`
	got := ext.String()

	if want != got {
		t.Errorf("%s validation failed: want %s, got %s", t.Name(), want, got)
	}
}

func TestCollectionDelete001(t *testing.T) {
	c := make(collection, 0)

	c.append(&Extension{
		Label: `X-TEST-PARAM`,
		Value: []string{`PARAM VALUE`},
	})

	if c.len() != 1 {
		t.Errorf("%s validation failed: valid type instance not accepted in append", t.Name())
		return
	}

	c.delete(0)
	if c.len() != 0 {
		t.Errorf("%s validation failed: index-based deletion unsuccessful", t.Name())
	}
}

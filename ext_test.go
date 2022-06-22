package schemax

import (
	"testing"
)

func TestExtension_String(t *testing.T) {
	var ext *Extension = new(Extension)
	ext.Label = `X-TEST-PARAM`
	ext.Value = []string{`TEST VALUE`}

	want := `X-TEST-PARAM 'TEST VALUE'`
	got := ext.String()
	if got != want {
		t.Errorf("%s failed: want %s, got %s", t.Name(), want, got)
	}
}

func TestExtensions_Set(t *testing.T) {
	var ext *Extension = new(Extension)
	ext.Label = `X-TEST-PARAM`
	ext.Value = []string{`TEST VALUE`}

	exts := NewExtensions()
	exts.Set(ext)
	exts.Set(`X-ALT-TEST-PARAM`, `VALUE`)
	exts.Set(`X-OTHER-TEST-PARAM`, `VALUE1`, `VALUE2`)

	want := 3
	got := exts.Len()

	if want != got {
		t.Errorf("%s failed: want len:%d, got len:%d", t.Name(), want, got)
	}
}

func TestExtensions_Get(t *testing.T) {
        var ext *Extension = new(Extension)
        ext.Label = `X-TEST-PARAM`
        ext.Value = []string{`TEST VALUE`}

        exts := NewExtensions()
        exts.Set(ext)
        exts.Set(`X-ALT-TEST-PARAM`, `VALUE`)
        exts.Set(`X-OTHER-TEST-PARAM`, `VALUE1`, `VALUE2`)

        want := 3
        got := exts.Len()

        if want != got {
                t.Errorf("%s failed: want len:%d, got len:%d", t.Name(), want, got)
		return
        }

	ex := exts.Get(`X-ALT-TEST-PARAM`)
	if ex.IsZero() {
                t.Errorf("%s failed: %T.Get failed", t.Name(), ex)
		return
	}

	if ex.Len() == 0 {
                t.Errorf("%s failed: insufficient %T value len", t.Name(), ex)
		return
	}

	wantv := `VALUE`
	gotv := ex.Value[0]

	if want != got {
		t.Errorf("%s failed: want %s, got %s", t.Name(), wantv, gotv)
	}
}

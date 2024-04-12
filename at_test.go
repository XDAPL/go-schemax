package schemax

import (
	"fmt"
	"testing"
)

func ExampleAttributeType_Map() {
	// initialize and populate a complete
	// Schema instance.
	var schema Schema = NewSchema()

	// Find the 'name' attributeType
	name := schema.AttributeTypes().Get(`name`)

	// marshal AttributeType into map[string][]string
	m := name.Map()

	fmt.Printf("%s", m[`NUMERICOID`][0])
	// Output: 2.5.4.41
}

/*
func TestNewAttributeType(t *testing.T) {
	var at AttributeType
	at.SetObsolete()
	t.Logf("Obsolete:%t\n", at.IsObsolete())

	at.SetNumericOID(`1.3.6.1.4.1.56521.999.5`)
	t.Logf("OID:%s\n", at.NumericOID())
}
*/

func TestParseAttributeType_bogus(t *testing.T) {
	raw := `( 1.2.3.4.5.6.7.8
		NAME 'bogusAttribute'
		DESC 'this is complete trash'
		SYNTAX 1.3.6.1.4.1.1466.115.121.1.58
		X-ORIGIN 'NOWHERE'` // missing closing paren

	if err := testSchema.ParseAttributeType(raw); err == nil {
		t.Errorf("%s error: invalid input triggered no error", t.Name())
		return
	}
}


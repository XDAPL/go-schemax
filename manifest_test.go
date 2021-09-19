package schemax

import (
	"testing"
)

func TestAliasesManifest001(t *testing.T) {
	schema := testSubschema()

	al := `jesse`
	jc := `1.3.6.1.4.1.56521`
	schema.SetAlias(al, jc)
	if _, exists := schema.ALM[al]; !exists {
		t.Errorf("Failed to set alias; did not resolve '%s' to OID:%s (%#v)", al, jc, schema.ALM)
	}
}

func TestAttributeTypesManifest001(t *testing.T) {

	schema := testSubschema()
	schema.ATM = PopulateDefaultAttributeTypesManifest()

	if at := schema.GetAttributeType(`userCertificate;binary`); at.IsZero() {
		t.Errorf("Unable to retrieve 'userCertificate' (%T) with tag", at)
	}

	schema.SetAttributeType(&AttributeType{
		Name:        NewName(`testAttr`, `otherTestName`),
		OID:         OID(`1.3.6.1.4.1.56521.999.18.100.1`),
		Syntax:      schema.GetLDAPSyntax(`integer`),
		Description: Description(`Test attribute`),
		Extensions: Extensions{
			`X-ORIGIN`: []string{`Jesse Coretta`},
		},
	})

	if at := schema.GetAttributeType(`testAttr`); at.IsZero() {
		t.Errorf("Unable to retrieve 'testAttr' %T", at)
	} else {
		if _, err := Unmarshal(at); err != nil {
			t.Errorf("Unable to unmarshal 'testAttr' (%T): %s", at, err.Error())
		}
	}

}

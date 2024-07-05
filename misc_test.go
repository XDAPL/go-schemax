package schemax

import (
	"fmt"
	"testing"
)

func ExampleIsNumericOID() {
	fmt.Printf("%t", IsNumericOID(`1.3.6.1.4.1`))
	// Output: true
}

func ExampleIsDescriptor() {
	fmt.Printf("%t", IsDescriptor(`telephoneNumber`))
	// Output: true
}

func TestMisc_codecov(t *testing.T) {
	_ = IsDescriptor(`test_`)
	_ = IsDescriptor(`tes--t`)
	_ = IsDescriptor(`te?st`)
	_ = IsDescriptor(``)
	_ = IsDescriptor(`_`)
	_ = IsNumericOID(`1.3.6.1.4.1`)
	_ = IsNumericOID(`^73`)
	_ = IsNumericOID(`l`)
	_ = condenseWHSP(`Lorem     ipsum 
dolor sit amet, 
           consectetur adipiscing elit,
  sed do eiusmod tempor

incididunt ut labore et dolore magna aliqua.`)
}

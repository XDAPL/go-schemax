package schemax

import (
	"fmt"
	"testing"
)

func ExampleNewName() {
	n := NewName()
	n.Push(`thisIsAName`)
	fmt.Println(n)
	// Output: 'thisIsAName'
}

func ExampleQuotedDescriptorList_Len() {
	n := NewName()
	n.Push(`thisIsAName`)
	fmt.Println(n.Len())
	// Output: 1
}

func ExampleQuotedDescriptorList_Contains() {
	n := NewName()
	n.Push(`thisIsAName`)
	fmt.Println(n.Contains(`otherName`))
	// Output: false
}

func ExampleQuotedDescriptorList_IsZero() {
	var n QuotedDescriptorList
	fmt.Println(n.IsZero())
	// Output: true
}

func ExampleQuotedDescriptorList_Index() {
	var n QuotedDescriptorList
	fmt.Println(n.Index(0))
	// Output:
}

func TestName_codecov(t *testing.T) {
	var n QuotedDescriptorList
	n.Push(``)
	n.Push(`^^^`)
	n.Push(`test`)
	n.Push(`test2`)
	n.Push(`test5`)
	n.Push(`test8`)
	n.Contains(`test7`)
	n.Contains(`test8`)
	n.cast().Push(`test3`)
	n.cast().NoPadding(true)
	n.smvStringer()
}

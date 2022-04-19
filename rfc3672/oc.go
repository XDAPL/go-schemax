package rfc3672

type RFC3672ObjectClasses []RFC3672ObjectClass
type RFC3672ObjectClass string

var (
	AllObjectClasses RFC3672ObjectClasses
	Subentry         RFC3672ObjectClass
)

func init() {
	Subentry = RFC3672ObjectClass(`( 2.5.17.0 NAME 'subentry' SUP top STRUCTURAL MUST ( cn $ subtreeSpecification ) X-ORIGIN 'RFC3672' )`)

	AllObjectClasses = RFC3672ObjectClasses{
		Subentry,
	}

}

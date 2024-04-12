package rfc3672

type ObjectClassDefinitions []ObjectClassDefinition
type ObjectClassDefinition string

var (
	AllObjectClasses ObjectClassDefinitions
	Subentry         ObjectClassDefinition
)

func init() {
	Subentry = ObjectClassDefinition(`( 2.5.17.0 NAME 'subentry' SUP top STRUCTURAL MUST ( cn $ subtreeSpecification ) X-ORIGIN 'RFC3672' )`)

	AllObjectClasses = ObjectClassDefinitions{
		Subentry,
	}

}

package rfc2079

type ObjectClassDefinitions []ObjectClassDefinition
type ObjectClassDefinition string

var (
	AllObjectClasses ObjectClassDefinitions
)

var (
	LabeledURIObject ObjectClassDefinition
)

func init() {

	LabeledURIObject = ObjectClassDefinition(`( 1.3.6.1.4.1.250.3.15 NAME 'labeledURIObject' DESC 'RFC2079: object that contains the URI attribute type' SUP top AUXILIARY MAY labeledURI X-ORIGIN 'RFC2079' )`)

	AllObjectClasses = ObjectClassDefinitions{
		LabeledURIObject,
	}

}

package rfc2079

type ObjectClassDefinitions []ObjectClassDefinition
type ObjectClassDefinition string

func (r ObjectClassDefinitions) Len() int {
	return len(r)
}

var (
	AllObjectClasses ObjectClassDefinitions
)

var (
	LabeledURIObject ObjectClassDefinition
)

func (r ObjectClassDefinition) String() string {
	return string(r)
}

func init() {

	LabeledURIObject = ObjectClassDefinition(`( 1.3.6.1.4.1.250.3.15 NAME 'labeledURIObject' DESC 'RFC2079: object that contains the URI attribute type' SUP top AUXILIARY MAY labeledURI X-ORIGIN 'RFC2079' )`)

	AllObjectClasses = ObjectClassDefinitions{
		LabeledURIObject,
	}

}

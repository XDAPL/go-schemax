package rfc2079

type AttributeTypeDefinitions []AttributeTypeDefinition
type AttributeTypeDefinition string

var (
	AllAttributeTypes AttributeTypeDefinitions
)

var (
	LabeledURI AttributeTypeDefinition
)

func (r AttributeTypeDefinition) String() string {
	return string(r)
}

func init() {

	LabeledURI = AttributeTypeDefinition(`( 1.3.6.1.4.1.250.1.57 NAME 'labeledURI' DESC 'RFC2079: Uniform Resource Identifier with optional label' EQUALITY caseExactMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.15 X-ORIGIN 'RFC2079' )`)

	AllAttributeTypes = AttributeTypeDefinitions{
		LabeledURI,
	}
}

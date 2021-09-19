package rfc2079

type RFC2079AttributeTypes []RFC2079AttributeType
type RFC2079AttributeType string

var (
	AllAttributeTypes RFC2079AttributeTypes
)

var (
	LabeledURI RFC2079AttributeType
)

func init() {

	LabeledURI = RFC2079AttributeType(`( 1.3.6.1.4.1.250.1.57 NAME 'labeledURI' DESC 'RFC2079: Uniform Resource Identifier with optional label' EQUALITY caseExactMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.15 X-ORIGIN 'RFC2079' )`)

	AllAttributeTypes = RFC2079AttributeTypes{
		LabeledURI,
	}
}

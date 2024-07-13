package rfc5020

type AttributeTypeDefinitions []AttributeTypeDefinition
type AttributeTypeDefinition string

func (r AttributeTypeDefinitions) Len() int {
	return len(r)
}

var (
	AllAttributeTypes AttributeTypeDefinitions
)

var (
	EntryDN AttributeTypeDefinition
)

func (r AttributeTypeDefinition) String() string {
	return string(r)
}

func init() {

	EntryDN = AttributeTypeDefinition(`( 1.3.6.1.1.20 NAME 'entryDN' DESC 'DN of the entry' EQUALITY distinguishedNameMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.12 SINGLE-VALUE NO-USER-MODIFICATION USAGE directoryOperation X-ORIGIN 'RFC5020' )`)

	AllAttributeTypes = AttributeTypeDefinitions{
		EntryDN,
	}
}

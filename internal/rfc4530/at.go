package rfc4530

type AttributeTypeDefinitions []AttributeTypeDefinition
type AttributeTypeDefinition string

var (
	AllAttributeTypes AttributeTypeDefinitions
)

var (
	EntryUUID AttributeTypeDefinition
)

func (r AttributeTypeDefinition) String() string {
	return string(r)
}

func init() {

	EntryUUID = AttributeTypeDefinition(`( 1.3.6.1.1.16.4 NAME 'entryUUID' DESC 'UUID of the entry' EQUALITY uuidMatch ORDERING uuidOrderingMatch SYNTAX 1.3.6.1.1.16.1 SINGLE-VALUE NO-USER-MODIFICATION USAGE directoryOperation X-ORIGIN 'RFC4530' )`)

	AllAttributeTypes = AttributeTypeDefinitions{
		EntryUUID,
	}
}

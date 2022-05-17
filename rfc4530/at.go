package rfc4530

type RFC4530AttributeTypes []RFC4530AttributeType
type RFC4530AttributeType string

var (
	AllAttributeTypes RFC4530AttributeTypes
)

var (
	EntryUUID		  RFC4530AttributeType
)

func init() {

	EntryUUID = RFC4530AttributeType(`( 1.3.6.1.1.16.4 NAME 'entryUUID' DESC 'UUID of the entry' EQUALITY uuidMatch ORDERING uuidOrderingMatch SYNTAX 1.3.6.1.1.16.1 SINGLE-VALUE NO-USER-MODIFICATION USAGE directoryOperation X-ORIGIN 'RFC4530' )`)

	AllAttributeTypes = RFC4530AttributeTypes{
		EntryUUID,
	}
}

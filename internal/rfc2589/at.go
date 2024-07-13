package rfc2589

type AttributeTypeDefinitions []AttributeTypeDefinition
type AttributeTypeDefinition string

func (r AttributeTypeDefinitions) Len() int {
	return len(r)
}

var (
	AllAttributeTypes AttributeTypeDefinitions
)

var (
	EntryTTL,
	DynamicSubtrees AttributeTypeDefinition
)

func (r AttributeTypeDefinition) String() string {
	return string(r)
}

func init() {

	EntryTTL = AttributeTypeDefinition(`( 1.3.6.1.4.1.1466.101.119.3 NAME 'entryTtl' DESC 'RFC2589: entry time-to-live' SYNTAX 1.3.6.1.4.1.1466.115.121.1.27 SINGLE-VALUE NO-USER-MODIFICATION USAGE dSAOperation X-ORIGIN 'RFC2589' )`)
	DynamicSubtrees = AttributeTypeDefinition(`( 1.3.6.1.4.1.1466.101.119.4 NAME 'dynamicSubtrees' DESC 'RFC2589: dynamic subtrees' SYNTAX 1.3.6.1.4.1.1466.115.121.1.12 NO-USER-MODIFICATION USAGE dSAOperation X-ORIGIN 'RFC2589' )`)

	AllAttributeTypes = AttributeTypeDefinitions{
		EntryTTL,
		DynamicSubtrees,
	}
}

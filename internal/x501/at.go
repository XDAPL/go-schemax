package x501

type AttributeTypeDefinitions []AttributeTypeDefinition
type AttributeTypeDefinition string

func (r AttributeTypeDefinitions) Len() int {
	return len(r)
}

var (
	AllAttributeTypes AttributeTypeDefinitions
)

var (
	SubschemaTimestamp,
	HasSubordinates,
	HierarchyLevel,
	HierarchyBelow,
	HierarchyParent,
	HierarchyTop AttributeTypeDefinition
)

func (r AttributeTypeDefinition) String() string {
	return string(r)
}

func init() {
	SubschemaTimestamp = AttributeTypeDefinition(`( 2.5.18.8 NAME 'subschemaTimestamp' DESC 'Indicates the time that the subschema subentry for the entry was created or last modified' EQUALITY generalizedTimeMatch ORDERING generalizedTimeOrderingMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.24 SINGLE-VALUE NO-USER-MODIFICATION USAGE directoryOperation X-ORIGIN 'X.501' )`)
	HasSubordinates = AttributeTypeDefinition(`( 2.5.18.9 NAME 'hasSubordinates' DESC 'Whether any subordinate entries exist below the entry holding this attribute' EQUALITY 2.5.13.13 SYNTAX 1.3.6.1.4.1.1466.115.121.1.7 SINGLE-VALUE NO-USER-MODIFICATION USAGE directoryOperation X-ORIGIN 'X.501' )`)

	HierarchyLevel = AttributeTypeDefinition(`( 2.5.18.17 NAME 'hierarchyLevel' EQUALITY integerMatch ORDERING integerOrderingMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.27 SINGLE-VALUE NO-USER-MODIFICATION USAGE directoryOperation X-ORIGIN 'X.501' )`)
	HierarchyBelow = AttributeTypeDefinition(`( 2.5.18.18 NAME 'hierarchyBelow' EQUALITY booleanMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.7 SINGLE-VALUE NO-USER-MODIFICATION USAGE directoryOperation X-ORIGIN 'X.501' )`)
	HierarchyParent = AttributeTypeDefinition(`( 2.5.18.19 NAME 'hierarchyParent' EQUALITY distinguishedNameMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.12 SINGLE-VALUE NO-USER-MODIFICATION USAGE directoryOperation X-ORIGIN 'X.501' )`)
	HierarchyTop = AttributeTypeDefinition(`( 2.5.18.20 NAME 'hierarchyTop' EQUALITY distinguishedNameMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.12 SINGLE-VALUE NO-USER-MODIFICATION USAGE directoryOperation X-ORIGIN 'X.501' )`)

	AllAttributeTypes = AttributeTypeDefinitions{
		SubschemaTimestamp,
		HasSubordinates,
		HierarchyLevel,
		HierarchyBelow,
		HierarchyParent,
		HierarchyTop,
	}
}

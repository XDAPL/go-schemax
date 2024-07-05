package rfc3671

type AttributeTypeDefinitions []AttributeTypeDefinition
type AttributeTypeDefinition string

func (r AttributeTypeDefinitions) Len() int {
	return len(r)
}

var AllAttributeTypes AttributeTypeDefinitions

var (
	CollectiveAttributeSubentries,
	CollectiveExclusions,
	C_L,
	C_St,
	C_Street,
	C_O,
	C_OU,
	C_PostalAddress,
	C_PostalCode,
	C_PostOfficeBox,
	C_PhysicalDeliveryOfficeName,
	C_TelephoneNumber,
	C_TelexNumber,
	C_FacsimileTelephoneNumber,
	C_IISDN AttributeTypeDefinition
)

func (r AttributeTypeDefinition) String() string {
	return string(r)
}

func init() {

	CollectiveAttributeSubentries = AttributeTypeDefinition(`( 2.5.18.12
		NAME 'collectiveAttributeSubentries'
	        EQUALITY distinguishedNameMatch
	        SYNTAX 1.3.6.1.4.1.1466.115.121.1.12
		NO-USER-MODIFICATION
	        USAGE directoryOperation )`)

	CollectiveExclusions = AttributeTypeDefinition(`( 2.5.18.7
		NAME 'collectiveExclusions'
	        EQUALITY objectIdentifierMatch
	        SYNTAX 1.3.6.1.4.1.1466.115.121.1.38
	        USAGE directoryOperation )`)

	C_L = AttributeTypeDefinition(`( 2.5.4.7.1
		NAME 'c-l'
		SUP l
		COLLECTIVE )`)

	C_St = AttributeTypeDefinition(`( 2.5.4.8.1
		NAME 'c-st'
		SUP st
		COLLECTIVE )`)

	C_Street = AttributeTypeDefinition(`( 2.5.4.9.1
		NAME 'c-street'
		SUP street
		COLLECTIVE )`)

	C_O = AttributeTypeDefinition(`( 2.5.4.10.1
		NAME 'c-o'
		SUP o
		COLLECTIVE )`)

	C_OU = AttributeTypeDefinition(`( 2.5.4.11.1
		NAME 'c-ou'
		SUP ou
		COLLECTIVE )`)

	C_PostalAddress = AttributeTypeDefinition(`( 2.5.4.16.1
		NAME 'c-PostalAddress'
		SUP postalAddress
		COLLECTIVE )`)

	C_PostalCode = AttributeTypeDefinition(`( 2.5.4.17.1
		NAME 'c-PostalCode'
		SUP postalCode
		COLLECTIVE )`)

	C_PostOfficeBox = AttributeTypeDefinition(`( 2.5.4.18.1
		NAME 'c-PostOfficeBox'
		SUP postOfficeBox
		COLLECTIVE )`)

	C_PhysicalDeliveryOfficeName = AttributeTypeDefinition(`( 2.5.4.19.1
		NAME 'c-PhysicalDeliveryOfficeName'
		SUP physicalDeliveryOfficeName
		COLLECTIVE )`)

	C_TelephoneNumber = AttributeTypeDefinition(`( 2.5.4.20.1
		NAME 'c-TelephoneNumber'
		SUP telephoneNumber
		COLLECTIVE )`)

	C_TelexNumber = AttributeTypeDefinition(`( 2.5.4.21.1
		NAME 'c-TelexNumber'
		SUP telexNumber
		COLLECTIVE )`)

	C_FacsimileTelephoneNumber = AttributeTypeDefinition(`( 2.5.4.23.1
		NAME 'c-FacsimileTelephoneNumber'
		SUP facsimileTelephoneNumber
		COLLECTIVE )`)

	C_IISDN = AttributeTypeDefinition(`( 2.5.4.25.1
		NAME 'c-InternationalISDNNumber'
		SUP internationalISDNNumber
		COLLECTIVE )`)

	AllAttributeTypes = AttributeTypeDefinitions{
		CollectiveAttributeSubentries,
		CollectiveExclusions,
		C_L,
		C_St,
		C_Street,
		C_O,
		C_OU,
		C_PostalAddress,
		C_PostalCode,
		C_PostOfficeBox,
		C_PhysicalDeliveryOfficeName,
		C_TelephoneNumber,
		C_TelexNumber,
		C_FacsimileTelephoneNumber,
		C_IISDN,
	}
}

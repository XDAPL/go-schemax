package rfc4519

type ObjectClassDefinitions []ObjectClassDefinition
type ObjectClassDefinition string

/*
AllObjectClasses contains slices of all instances of ObjectClassDefinition defined in this package.
*/
var AllObjectClasses ObjectClassDefinitions

var (
	ApplicationProcess   ObjectClassDefinition
	Country              ObjectClassDefinition
	DCObject             ObjectClassDefinition
	Device               ObjectClassDefinition
	GroupOfNames         ObjectClassDefinition
	GroupOfUniqueNames   ObjectClassDefinition
	Locality             ObjectClassDefinition
	Organization         ObjectClassDefinition
	OrganizationalPerson ObjectClassDefinition
	OrganizationalRole   ObjectClassDefinition
	OrganizationalUnit   ObjectClassDefinition
	Person               ObjectClassDefinition
	ResidentialPerson    ObjectClassDefinition
	UIDObject            ObjectClassDefinition
)

func (r ObjectClassDefinition) String() string {
	return string(r)
}

// User ObjectClasses
func init() {
	ApplicationProcess = ObjectClassDefinition(`( 2.5.6.11 NAME 'applicationProcess' SUP top STRUCTURAL MUST cn MAY ( seeAlso $ ou $ l $ description ) X-ORIGIN 'RFC4519' )`)
	Country = ObjectClassDefinition(`( 2.5.6.2 NAME 'country' SUP top STRUCTURAL MUST c MAY ( searchGuide $ description ) X-ORIGIN 'RFC4519' )`)
	DCObject = ObjectClassDefinition(`( 1.3.6.1.4.1.1466.344 NAME 'dcObject' SUP top AUXILIARY MUST dc X-ORIGIN 'RFC4519' )`)
	Device = ObjectClassDefinition(`( 2.5.6.14 NAME 'device' SUP top STRUCTURAL MUST cn MAY ( serialNumber $ seeAlso $ owner $ ou $ o $ l $ description ) X-ORIGIN 'RFC4519' )`)
	GroupOfNames = ObjectClassDefinition(`( 2.5.6.9 NAME 'groupOfNames' SUP top STRUCTURAL MUST ( member $ cn ) MAY ( businessCategory $ seeAlso $ owner $ ou $ o $ description ) X-ORIGIN 'RFC4519' )`)
	GroupOfUniqueNames = ObjectClassDefinition(`( 2.5.6.17 NAME 'groupOfUniqueNames' SUP top STRUCTURAL MUST ( uniqueMember $ cn ) MAY ( businessCategory $ seeAlso $ owner $ ou $ o $ description ) X-ORIGIN 'RFC4519' )`)
	Locality = ObjectClassDefinition(`( 2.5.6.3 NAME 'locality' SUP top STRUCTURAL MAY ( street $ seeAlso $ searchGuide $ st $ l $ description ) X-ORIGIN 'RFC4519' )`)
	Organization = ObjectClassDefinition(`( 2.5.6.4 NAME 'organization' SUP top STRUCTURAL MUST o MAY ( userPassword $ searchGuide $ seeAlso $ businessCategory $ x121Address $ registeredAddress $ destinationIndicator $ preferredDeliveryMethod $ telexNumber $ teletexTerminalIdentifier $ telephoneNumber $ internationalISDNNumber $ facsimileTelephoneNumber $ street $ postOfficeBox $ postalCode $ postalAddress $ physicalDeliveryOfficeName $ st $ l $ description ) X-ORIGIN 'RFC4519' )`)
	OrganizationalPerson = ObjectClassDefinition(`( 2.5.6.7 NAME 'organizationalPerson' SUP person STRUCTURAL MAY ( title $ x121Address $ registeredAddress $ destinationIndicator $ preferredDeliveryMethod $ telexNumber $ teletexTerminalIdentifier $ telephoneNumber $ internationalISDNNumber $ facsimileTelephoneNumber $ street $ postOfficeBox $ postalCode $ postalAddress $ physicalDeliveryOfficeName $ ou $ st $ l ) X-ORIGIN 'RFC4519' )`)
	OrganizationalRole = ObjectClassDefinition(`( 2.5.6.8 NAME 'organizationalRole' SUP top STRUCTURAL MUST cn MAY ( x121Address $ registeredAddress $ destinationIndicator $ preferredDeliveryMethod $ telexNumber $ teletexTerminalIdentifier $ telephoneNumber $ internationalISDNNumber $ facsimileTelephoneNumber $ seeAlso $ roleOccupant $ preferredDeliveryMethod $ street $ postOfficeBox $ postalCode $ postalAddress $ physicalDeliveryOfficeName $ ou $ st $ l $ description ) X-ORIGIN 'RFC4519' )`)
	OrganizationalUnit = ObjectClassDefinition(`( 2.5.6.5 NAME 'organizationalUnit' SUP top STRUCTURAL MUST ou MAY ( businessCategory $ description $ destinationIndicator $ facsimileTelephoneNumber $ internationalISDNNumber $ l $ physicalDeliveryOfficeName $ postalAddress $ postalCode $ postOfficeBox $ preferredDeliveryMethod $ registeredAddress $ searchGuide $ seeAlso $ st $ street $ telephoneNumber $ teletexTerminalIdentifier $ telexNumber $ userPassword $ x121Address ) X-ORIGIN 'RFC4519' )`)
	Person = ObjectClassDefinition(`( 2.5.6.6 NAME 'person' SUP top STRUCTURAL MUST ( sn $ cn ) MAY ( userPassword $ telephoneNumber $ seeAlso $ description ) X-ORIGIN 'RFC4519' )`)
	ResidentialPerson = ObjectClassDefinition(`( 2.5.6.10 NAME 'residentialPerson' SUP person STRUCTURAL MUST l MAY ( businessCategory $ x121Address $ registeredAddress $ destinationIndicator $ preferredDeliveryMethod $ telexNumber $ teletexTerminalIdentifier $ telephoneNumber $ internationalISDNNumber $ facsimileTelephoneNumber $ preferredDeliveryMethod $ street $ postOfficeBox $ postalCode $ postalAddress $ physicalDeliveryOfficeName $ st $ l ) X-ORIGIN 'RFC4519' )`)
	UIDObject = ObjectClassDefinition(`( 1.3.6.1.1.3.1 NAME 'uidObject' SUP top AUXILIARY MUST uid X-ORIGIN 'RFC4519' )`)

	AllObjectClasses = ObjectClassDefinitions{
		Person,
		ApplicationProcess,
		Country,
		DCObject,
		Device,
		GroupOfNames,
		GroupOfUniqueNames,
		Locality,
		Organization,
		OrganizationalPerson,
		OrganizationalRole,
		OrganizationalUnit,
		ResidentialPerson,
		UIDObject,
	}
}

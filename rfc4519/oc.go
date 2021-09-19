package rfc4519

type RFC4519ObjectClasses []RFC4519ObjectClass
type RFC4519ObjectClass string

/*
ObjectClasses contains slices of all instances of RFC4519ObjectClass defined in this package.
*/
var AllObjectClasses RFC4519ObjectClasses

var (
	ApplicationProcess   RFC4519ObjectClass
	Country              RFC4519ObjectClass
	DCObject             RFC4519ObjectClass
	Device               RFC4519ObjectClass
	GroupOfNames         RFC4519ObjectClass
	GroupOfUniqueNames   RFC4519ObjectClass
	Locality             RFC4519ObjectClass
	Organization         RFC4519ObjectClass
	OrganizationalPerson RFC4519ObjectClass
	OrganizationalRole   RFC4519ObjectClass
	OrganizationalUnit   RFC4519ObjectClass
	Person               RFC4519ObjectClass
	ResidentialPerson    RFC4519ObjectClass
	UIDObject            RFC4519ObjectClass
)

// User ObjectClasses
func init() {
	ApplicationProcess = RFC4519ObjectClass(`( 2.5.6.11 NAME 'applicationProcess' SUP top STRUCTURAL MUST cn MAY ( seeAlso $ ou $ l $ description ) X-ORIGIN 'RFC4519' )`)
	Country = RFC4519ObjectClass(`( 2.5.6.2 NAME 'country' SUP top STRUCTURAL MUST c MAY ( searchGuide $ description ) X-ORIGIN 'RFC4519' )`)
	DCObject = RFC4519ObjectClass(`( 1.3.6.1.4.1.1466.344 NAME 'dcObject' SUP top AUXILIARY MUST dc X-ORIGIN 'RFC4519' )`)
	Device = RFC4519ObjectClass(`( 2.5.6.14 NAME 'device' SUP top STRUCTURAL MUST cn MAY ( serialNumber $ seeAlso $ owner $ ou $ o $ l $ description ) X-ORIGIN 'RFC4519' )`)
	GroupOfNames = RFC4519ObjectClass(`( 2.5.6.9 NAME 'groupOfNames' SUP top STRUCTURAL MUST ( member $ cn ) MAY ( businessCategory $ seeAlso $ owner $ ou $ o $ description ) X-ORIGIN 'RFC4519' )`)
	GroupOfUniqueNames = RFC4519ObjectClass(`( 2.5.6.17 NAME 'groupOfUniqueNames' SUP top STRUCTURAL MUST ( uniqueMember $ cn ) MAY ( businessCategory $ seeAlso $ owner $ ou $ o $ description ) X-ORIGIN 'RFC4519' )`)
	Locality = RFC4519ObjectClass(`( 2.5.6.3 NAME 'locality' SUP top STRUCTURAL MAY ( street $ seeAlso $ searchGuide $ st $ l $ description ) X-ORIGIN 'RFC4519' )`)
	Organization = RFC4519ObjectClass(`( 2.5.6.4 NAME 'organization' SUP top STRUCTURAL MUST o MAY ( userPassword $ searchGuide $ seeAlso $ businessCategory $ x121Address $ registeredAddress $ destinationIndicator $ preferredDeliveryMethod $ telexNumber $ teletexTerminalIdentifier $ telephoneNumber $ internationalISDNNumber $ facsimileTelephoneNumber $ street $ postOfficeBox $ postalCode $ postalAddress $ physicalDeliveryOfficeName $ st $ l $ description ) X-ORIGIN 'RFC4519' )`)
	OrganizationalPerson = RFC4519ObjectClass(`( 2.5.6.7 NAME 'organizationalPerson' SUP person STRUCTURAL MAY ( title $ x121Address $ registeredAddress $ destinationIndicator $ preferredDeliveryMethod $ telexNumber $ teletexTerminalIdentifier $ telephoneNumber $ internationalISDNNumber $ facsimileTelephoneNumber $ street $ postOfficeBox $ postalCode $ postalAddress $ physicalDeliveryOfficeName $ ou $ st $ l ) X-ORIGIN 'RFC4519' )`)
	OrganizationalRole = RFC4519ObjectClass(`( 2.5.6.8 NAME 'organizationalRole' SUP top STRUCTURAL MUST cn MAY ( x121Address $ registeredAddress $ destinationIndicator $ preferredDeliveryMethod $ telexNumber $ teletexTerminalIdentifier $ telephoneNumber $ internationalISDNNumber $ facsimileTelephoneNumber $ seeAlso $ roleOccupant $ preferredDeliveryMethod $ street $ postOfficeBox $ postalCode $ postalAddress $ physicalDeliveryOfficeName $ ou $ st $ l $ description ) X-ORIGIN 'RFC4519' )`)
	OrganizationalUnit = RFC4519ObjectClass(`( 2.5.6.5 NAME 'organizationalUnit' SUP top STRUCTURAL MUST ou MAY ( businessCategory $ description $ destinationIndicator $ facsimileTelephoneNumber $ internationalISDNNumber $ l $ physicalDeliveryOfficeName $ postalAddress $ postalCode $ postOfficeBox $ preferredDeliveryMethod $ registeredAddress $ searchGuide $ seeAlso $ st $ street $ telephoneNumber $ teletexTerminalIdentifier $ telexNumber $ userPassword $ x121Address ) X-ORIGIN 'RFC4519' )`)
	Person = RFC4519ObjectClass(`( 2.5.6.6 NAME 'person' SUP top STRUCTURAL MUST ( sn $ cn ) MAY ( userPassword $ telephoneNumber $ seeAlso $ description ) X-ORIGIN 'RFC4519' )`)
	ResidentialPerson = RFC4519ObjectClass(`( 2.5.6.10 NAME 'residentialPerson' SUP person STRUCTURAL MUST l MAY ( businessCategory $ x121Address $ registeredAddress $ destinationIndicator $ preferredDeliveryMethod $ telexNumber $ teletexTerminalIdentifier $ telephoneNumber $ internationalISDNNumber $ facsimileTelephoneNumber $ preferredDeliveryMethod $ street $ postOfficeBox $ postalCode $ postalAddress $ physicalDeliveryOfficeName $ st $ l ) X-ORIGIN 'RFC4519' )`)
	UIDObject = RFC4519ObjectClass(`( 1.3.6.1.1.3.1 NAME 'uidObject' SUP top AUXILIARY MUST uid X-ORIGIN 'RFC4519' )`)

	AllObjectClasses = RFC4519ObjectClasses{
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

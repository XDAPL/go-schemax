package rfc4524

type ObjectClassDefinitions []ObjectClassDefinition
type ObjectClassDefinition string

func (r ObjectClassDefinitions) Len() int {
	return len(r)
}

var (
	AllObjectClasses ObjectClassDefinitions
)

var (
	Account              ObjectClassDefinition
	Document             ObjectClassDefinition
	DocumentSeries       ObjectClassDefinition
	Domain               ObjectClassDefinition
	DomainRelatedObject  ObjectClassDefinition
	FriendlyCountry      ObjectClassDefinition
	RFC822LocalPart      ObjectClassDefinition
	Room                 ObjectClassDefinition
	SimpleSecurityObject ObjectClassDefinition
)

func (r ObjectClassDefinition) String() string {
	return string(r)
}

func init() {

	Account = ObjectClassDefinition(`( 0.9.2342.19200300.100.4.5 NAME 'account' SUP top STRUCTURAL MUST uid MAY ( description $ seeAlso $ l $ o $ ou $ host ) X-ORIGIN 'RFC4524' )`)
	Document = ObjectClassDefinition(`( 0.9.2342.19200300.100.4.6 NAME 'document' SUP top STRUCTURAL MUST documentIdentifier MAY ( cn $ description $ seeAlso $ l $ o $ ou $ documentTitle $ documentVersion $ documentAuthor $ documentLocation $ documentPublisher ) X-ORIGIN 'RFC4524' )`)
	DocumentSeries = ObjectClassDefinition(`( 0.9.2342.19200300.100.4.9 NAME 'documentSeries' SUP top STRUCTURAL MUST cn MAY ( description $ l $ o $ ou $ seeAlso $ telephoneNumber ) X-ORIGIN 'RFC4524' )`)
	Domain = ObjectClassDefinition(`( 0.9.2342.19200300.100.4.13 NAME 'domain' SUP top STRUCTURAL MUST dc MAY ( userPassword $ searchGuide $ seeAlso $ businessCategory $ x121Address $ registeredAddress $ destinationIndicator $ preferredDeliveryMethod $ telexNumber $ teletexTerminalIdentifier $ telephoneNumber $ internationaliSDNNumber $ facsimileTelephoneNumber $ street $ postOfficeBox $ postalCode $ postalAddress $ physicalDeliveryOfficeName $ st $ l $ description $ o $ associatedName ) X-ORIGIN 'RFC4524' )`)
	DomainRelatedObject = ObjectClassDefinition(`( 0.9.2342.19200300.100.4.17 NAME 'domainRelatedObject' SUP top AUXILIARY MUST associatedDomain X-ORIGIN 'RFC4524' )`)
	FriendlyCountry = ObjectClassDefinition(`( 0.9.2342.19200300.100.4.18 NAME 'friendlyCountry' SUP country STRUCTURAL MUST co X-ORIGIN 'RFC4524' )`)
	RFC822LocalPart = ObjectClassDefinition(`( 0.9.2342.19200300.100.4.14 NAME 'rFC822localPart' SUP domain STRUCTURAL MAY ( cn $ description $ destinationIndicator $ facsimileTelephoneNumber $ internationaliSDNNumber $ physicalDeliveryOfficeName $ postalAddress $ postalCode $ postOfficeBox $ preferredDeliveryMethod $ registeredAddress $ seeAlso $ sn $ street $ telephoneNumber $ teletexTerminalIdentifier $ telexNumber $ x121Address ) X-ORIGIN 'RFC4524' )`)
	Room = ObjectClassDefinition(`( 0.9.2342.19200300.100.4.7 NAME 'room' SUP top STRUCTURAL MUST cn MAY ( roomNumber $ description $ seeAlso $ telephoneNumber ) X-ORIGIN 'RFC4524' )`)
	SimpleSecurityObject = ObjectClassDefinition(`( 0.9.2342.19200300.100.4.19 NAME 'simpleSecurityObject' SUP top AUXILIARY MUST userPassword X-ORIGIN 'RFC4524' )`)

	AllObjectClasses = ObjectClassDefinitions{
		Account,
		Document,
		DocumentSeries,
		Domain,
		DomainRelatedObject,
		FriendlyCountry,
		RFC822LocalPart,
		Room,
		SimpleSecurityObject,
	}

}

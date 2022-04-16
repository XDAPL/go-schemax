package rfc4524

type RFC4524ObjectClasses []RFC4524ObjectClass
type RFC4524ObjectClass string

var (
	AllObjectClasses RFC4524ObjectClasses
)

var (
	Account              RFC4524ObjectClass
	Document             RFC4524ObjectClass
	DocumentSeries       RFC4524ObjectClass
	Domain               RFC4524ObjectClass
	DomainRelatedObject  RFC4524ObjectClass
	FriendlyCountry      RFC4524ObjectClass
	RFC822LocalPart      RFC4524ObjectClass
	Room                 RFC4524ObjectClass
	SimpleSecurityObject RFC4524ObjectClass
)

func init() {

	Account = RFC4524ObjectClass(`( 0.9.2342.19200300.100.4.5 NAME 'account' SUP top STRUCTURAL MUST uid MAY ( description $ seeAlso $ l $ o $ ou $ host ) X-ORIGIN 'RFC4524' )`)
	Document = RFC4524ObjectClass(`( 0.9.2342.19200300.100.4.6 NAME 'document' SUP top STRUCTURAL MUST documentIdentifier MAY ( cn $ description $ seeAlso $ l $ o $ ou $ documentTitle $ documentVersion $ documentAuthor $ documentLocation $ documentPublisher ) X-ORIGIN 'RFC4524' )`)
	DocumentSeries = RFC4524ObjectClass(`( 0.9.2342.19200300.100.4.9 NAME 'documentSeries' SUP top STRUCTURAL MUST cn MAY ( description $ l $ o $ ou $ seeAlso $ telephoneNumber ) X-ORIGIN 'RFC4524' )`)
	Domain = RFC4524ObjectClass(`( 0.9.2342.19200300.100.4.13 NAME 'domain' SUP top STRUCTURAL MUST dc MAY ( userPassword $ searchGuide $ seeAlso $ businessCategory $ x121Address $ registeredAddress $ destinationIndicator $ preferredDeliveryMethod $ telexNumber $ teletexTerminalIdentifier $ telephoneNumber $ internationaliSDNNumber $ facsimileTelephoneNumber $ street $ postOfficeBox $ postalCode $ postalAddress $ physicalDeliveryOfficeName $ st $ l $ description $ o $ associatedName ) X-ORIGIN 'RFC4524' )`)
	DomainRelatedObject = RFC4524ObjectClass(`( 0.9.2342.19200300.100.4.17 NAME 'domainRelatedObject' SUP top AUXILIARY MUST associatedDomain X-ORIGIN 'RFC4524' )`)
	FriendlyCountry = RFC4524ObjectClass(`( 0.9.2342.19200300.100.4.18 NAME 'friendlyCountry' SUP country STRUCTURAL MUST co X-ORIGIN 'RFC4524' )`)
	RFC822LocalPart = RFC4524ObjectClass(`( 0.9.2342.19200300.100.4.14 NAME 'rFC822localPart' SUP domain STRUCTURAL MAY ( cn $ description $ destinationIndicator $ facsimileTelephoneNumber $ internationaliSDNNumber $ physicalDeliveryOfficeName $ postalAddress $ postalCode $ postOfficeBox $ preferredDeliveryMethod $ registeredAddress $ seeAlso $ sn $ street $ telephoneNumber $ teletexTerminalIdentifier $ telexNumber $ x121Address ) X-ORIGIN 'RFC4524' )`)
	Room = RFC4524ObjectClass(`( 0.9.2342.19200300.100.4.7 NAME 'room' SUP top STRUCTURAL MUST cn MAY ( roomNumber $ description $ seeAlso $ telephoneNumber ) X-ORIGIN 'RFC4524' )`)
	SimpleSecurityObject = RFC4524ObjectClass(`( 0.9.2342.19200300.100.4.19 NAME 'simpleSecurityObject' SUP top AUXILIARY MUST userPassword X-ORIGIN 'RFC4524' )`)

	AllObjectClasses = RFC4524ObjectClasses{
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

package rfc2798

type RFC2798ObjectClasses []RFC2798ObjectClass
type RFC2798ObjectClass string

var (
	AllObjectClasses RFC2798ObjectClasses
)

var (
	INetOrgPerson RFC2798ObjectClass
)

func init() {

	INetOrgPerson = RFC2798ObjectClass(`( 2.16.840.1.113730.3.2.2 NAME 'inetOrgPerson' SUP organizationalPerson STRUCTURAL MAY ( audio $ businessCategory $ carLicense $ departmentNumber $ displayName $ employeeNumber $ employeeType $ givenName $ homePhone $ homePostalAddress $ initials $ jpegPhoto $ labeledURI $ mail $ manager $ mobile $ o $ pager $ photo $ roomNumber $ secretary $ uid $ userCertificate $ x500uniqueIdentifier $ preferredLanguage $ userSMIMECertificate $ userPKCS12 ) X-ORIGIN 'RFC2798' )`)

	AllObjectClasses = RFC2798ObjectClasses{
		INetOrgPerson,
	}

}

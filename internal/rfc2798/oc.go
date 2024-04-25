package rfc2798

type ObjectClassDefinitions []ObjectClassDefinition
type ObjectClassDefinition string

var (
	AllObjectClasses ObjectClassDefinitions
)

var (
	INetOrgPerson ObjectClassDefinition
)

func (r ObjectClassDefinition) String() string {
	return string(r)
}

func init() {

	INetOrgPerson = ObjectClassDefinition(`( 2.16.840.1.113730.3.2.2 NAME 'inetOrgPerson' SUP organizationalPerson STRUCTURAL MAY ( audio $ businessCategory $ carLicense $ departmentNumber $ displayName $ employeeNumber $ employeeType $ givenName $ homePhone $ homePostalAddress $ initials $ jpegPhoto $ labeledURI $ mail $ manager $ mobile $ o $ pager $ photo $ roomNumber $ secretary $ uid $ userCertificate $ x500uniqueIdentifier $ preferredLanguage $ userSMIMECertificate $ userPKCS12 ) X-ORIGIN 'RFC2798' )`)

	AllObjectClasses = ObjectClassDefinitions{
		INetOrgPerson,
	}

}

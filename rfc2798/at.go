package rfc2798

type RFC2798AttributeTypes []RFC2798AttributeType
type RFC2798AttributeType string

var (
	AllAttributeTypes RFC2798AttributeTypes
)

var (
	Audio                RFC2798AttributeType
	Photo                RFC2798AttributeType
	CarLicense           RFC2798AttributeType
	DepartmentNumber     RFC2798AttributeType
	DisplayName          RFC2798AttributeType
	EmployeeNumber       RFC2798AttributeType
	EmployeeType         RFC2798AttributeType
	JPEGPhoto            RFC2798AttributeType
	PreferredLanguage    RFC2798AttributeType
	UserSMIMECertificate RFC2798AttributeType
	UserPKCS12           RFC2798AttributeType
)

func init() {

	Audio = RFC2798AttributeType(`( 0.9.2342.19200300.100.1.55 NAME 'audio' EQUALITY octetStringMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.40{250000} X-ORIGIN 'RFC2798' )`)
	CarLicense = RFC2798AttributeType(`( 2.16.840.1.113730.3.1.1 NAME 'carLicense' DESC 'vehicle license or registration plate' EQUALITY caseIgnoreMatch SUBSTR caseIgnoreSubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.15 X-ORIGIN 'RFC2798' )`)
	DepartmentNumber = RFC2798AttributeType(`( 2.16.840.1.113730.3.1.2 NAME 'departmentNumber' DESC 'identifies a department within an organization' EQUALITY caseIgnoreMatch SUBSTR caseIgnoreSubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.15 X-ORIGIN 'RFC2798' )`)
	DisplayName = RFC2798AttributeType(`( 2.16.840.1.113730.3.1.241 NAME 'displayName' DESC 'preferred name of a person to be used when displaying entries' EQUALITY caseIgnoreMatch SUBSTR caseIgnoreSubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.15 SINGLE-VALUE X-ORIGIN 'RFC2798' )`)
	EmployeeNumber = RFC2798AttributeType(`( 2.16.840.1.113730.3.1.3 NAME 'employeeNumber' DESC 'numerically identifies an employee within an organization' EQUALITY caseIgnoreMatch SUBSTR caseIgnoreSubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.15 SINGLE-VALUE X-ORIGIN 'RFC2798' )`)
	EmployeeType = RFC2798AttributeType(`( 2.16.840.1.113730.3.1.4 NAME 'employeeType' DESC 'type of employment for a person' EQUALITY caseIgnoreMatch SUBSTR caseIgnoreSubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.15 X-ORIGIN 'RFC2798' )`)
	JPEGPhoto = RFC2798AttributeType(`( 0.9.2342.19200300.100.1.60 NAME 'jpegPhoto' DESC 'a JPEG image' SYNTAX 1.3.6.1.4.1.1466.115.121.1.28 X-ORIGIN 'RFC2798' )`)
	Photo = RFC2798AttributeType(`( 0.9.2342.19200300.100.1.7 NAME 'photo' DESC 'RFC1274: photo (G3 fax)' SYNTAX 1.3.6.1.4.1.1466.115.121.1.23{25000} X-ORIGIN 'RFC2798' )`)
	PreferredLanguage = RFC2798AttributeType(`( 2.16.840.1.113730.3.1.39 NAME 'preferredLanguage' DESC 'preferred written or spoken language for a person' EQUALITY caseIgnoreMatch SUBSTR caseIgnoreSubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.15 SINGLE-VALUE X-ORIGIN 'RFC2798' )`)
	UserSMIMECertificate = RFC2798AttributeType(`( 2.16.840.1.113730.3.1.40 NAME 'userSMIMECertificate' DESC 'PKCS#7 SignedData used to support S/MIME' SYNTAX 1.3.6.1.4.1.1466.115.121.1.5 X-ORIGIN 'RFC2798' )`)
	UserPKCS12 = RFC2798AttributeType(`( 2.16.840.1.113730.3.1.216 NAME 'userPKCS12' DESC 'PKCS #12 PFX PDU for exchange of personal identity information' SYNTAX 1.3.6.1.4.1.1466.115.121.1.5 X-ORIGIN 'RFC2798' )`)

	AllAttributeTypes = RFC2798AttributeTypes{
		Audio,
		CarLicense,
		DepartmentNumber,
		DisplayName,
		EmployeeNumber,
		EmployeeType,
		JPEGPhoto,
		Photo,
		PreferredLanguage,
		UserSMIMECertificate,
		UserPKCS12,
	}
}

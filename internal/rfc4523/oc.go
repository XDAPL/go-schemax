package rfc4523

type ObjectClassDefinitions []ObjectClassDefinition
type ObjectClassDefinition string

var AllObjectClasses ObjectClassDefinitions

var (
	PKIUser                  ObjectClassDefinition
	PKICA                    ObjectClassDefinition
	CRLDistributionPoint     ObjectClassDefinition
	DeltaCRL                 ObjectClassDefinition
	StrongAuthenticationUser ObjectClassDefinition
	UserSecurityInformation  ObjectClassDefinition
	CertificationAuthority   ObjectClassDefinition
	CertificationAuthorityV2 ObjectClassDefinition
)

func (r ObjectClassDefinition) String() string {
	return string(r)
}

func init() {

	PKIUser = ObjectClassDefinition(`( 2.5.6.21 NAME 'pkiUser' DESC 'X.509 PKI User' SUP top AUXILIARY MAY userCertificate X-ORIGIN 'RFC4523' )`)
	PKICA = ObjectClassDefinition(`( 2.5.6.22 NAME 'pkiCA' DESC 'X.509 PKI Certificate Authority' SUP top AUXILIARY MAY ( cACertificate $ certificateRevocationList $ authorityRevocationList $ crossCertificatePair ) X-ORIGIN 'RFC4523' )`)
	CRLDistributionPoint = ObjectClassDefinition(`( 2.5.6.19 NAME 'cRLDistributionPoint' DESC 'X.509 CRL distribution point' SUP top STRUCTURAL MUST cn MAY ( certificateRevocationList $ authorityRevocationList $ deltaRevocationList ) X-ORIGIN 'RFC4523' )`)
	DeltaCRL = ObjectClassDefinition(`( 2.5.6.23 NAME 'deltaCRL' DESC 'X.509 delta CRL' SUP top AUXILIARY MAY deltaRevocationList X-ORIGIN 'RFC4523' )`)
	StrongAuthenticationUser = ObjectClassDefinition(`( 2.5.6.15 NAME 'strongAuthenticationUser' DESC 'X.521 strong authentication user' SUP top AUXILIARY MUST userCertificate X-ORIGIN 'RFC4523' )`)
	UserSecurityInformation = ObjectClassDefinition(`( 2.5.6.18 NAME 'userSecurityInformation' DESC 'X.521 user security information' SUP top AUXILIARY MAY ( supportedAlgorithms ) X-ORIGIN 'RFC4523' )`)
	CertificationAuthority = ObjectClassDefinition(`( 2.5.6.16 NAME 'certificationAuthority' DESC 'X.509 certificate authority' SUP top AUXILIARY MUST ( authorityRevocationList $ certificateRevocationList $ cACertificate ) MAY crossCertificatePair X-ORIGIN 'RFC4523' )`)
	CertificationAuthorityV2 = ObjectClassDefinition(`( 2.5.6.16.2 NAME 'certificationAuthority-V2' DESC 'X.509 certificate authority, version 2' SUP certificationAuthority AUXILIARY MAY deltaRevocationList X-ORIGIN 'RFC4523' )`)

	AllObjectClasses = ObjectClassDefinitions{
		PKIUser,
		PKICA,
		CRLDistributionPoint,
		DeltaCRL,
		StrongAuthenticationUser,
		UserSecurityInformation,
		CertificationAuthority,
		CertificationAuthorityV2,
	}
}

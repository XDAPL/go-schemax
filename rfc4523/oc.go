package rfc4523

type RFC4523ObjectClasses []RFC4523ObjectClass
type RFC4523ObjectClass string

var (
	AllObjectClasses RFC4523ObjectClasses
)

var (
	PKIUser                  RFC4523ObjectClass
	PKICA                    RFC4523ObjectClass
	CRLDistributionPoint     RFC4523ObjectClass
	DeltaCRL                 RFC4523ObjectClass
	StrongAuthenticationUser RFC4523ObjectClass
	UserSecurityInformation  RFC4523ObjectClass
	CertificationAuthority   RFC4523ObjectClass
	CertificationAuthorityV2 RFC4523ObjectClass
)

func init() {

	PKIUser = RFC4523ObjectClass(`( 2.5.6.21 NAME 'pkiUser' DESC 'X.509 PKI User' SUP top AUXILIARY MAY userCertificate X-ORIGIN 'RFC4523' )`)
	PKICA = RFC4523ObjectClass(`( 2.5.6.22 NAME 'pkiCA' DESC 'X.509 PKI Certificate Authority' SUP top AUXILIARY MAY ( cACertificate $ certificateRevocationList $ authorityRevocationList $ crossCertificatePair ) X-ORIGIN 'RFC4523' )`)
	CRLDistributionPoint = RFC4523ObjectClass(`( 2.5.6.19 NAME 'cRLDistributionPoint' DESC 'X.509 CRL distribution point' SUP top STRUCTURAL MUST cn MAY ( certificateRevocationList $ authorityRevocationList $ deltaRevocationList ) X-ORIGIN 'RFC4523' )`)
	DeltaCRL = RFC4523ObjectClass(`( 2.5.6.23 NAME 'deltaCRL' DESC 'X.509 delta CRL' SUP top AUXILIARY MAY deltaRevocationList X-ORIGIN 'RFC4523' )`)
	StrongAuthenticationUser = RFC4523ObjectClass(`( 2.5.6.15 NAME 'strongAuthenticationUser' DESC 'X.521 strong authentication user' SUP top AUXILIARY MUST userCertificate X-ORIGIN 'RFC4523' )`)
	UserSecurityInformation = RFC4523ObjectClass(`( 2.5.6.18 NAME 'userSecurityInformation' DESC 'X.521 user security information' SUP top AUXILIARY MAY ( supportedAlgorithms ) X-ORIGIN 'RFC4523' )`)
	CertificationAuthority = RFC4523ObjectClass(`( 2.5.6.16 NAME 'certificationAuthority' DESC 'X.509 certificate authority' SUP top AUXILIARY MUST ( authorityRevocationList $ certificateRevocationList $ cACertificate ) MAY crossCertificatePair X-ORIGIN 'RFC4523' )`)
	CertificationAuthorityV2 = RFC4523ObjectClass(`( 2.5.6.16.2 NAME 'certificationAuthority-V2' DESC 'X.509 certificate authority, version 2' SUP certificationAuthority AUXILIARY MAY deltaRevocationList X-ORIGIN 'RFC4523' )`)

	AllObjectClasses = RFC4523ObjectClasses{
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

package rfc4523

type RFC4523AttributeTypes []RFC4523AttributeType
type RFC4523AttributeType string

var (
	AllAttributeTypes RFC4523AttributeTypes
)

var (
	UserCertificate           RFC4523AttributeType
	CACertificate             RFC4523AttributeType
	CrossCertificatePair      RFC4523AttributeType
	CertificateRevocationList RFC4523AttributeType
	AuthorityRevocationList   RFC4523AttributeType
	DeltaRevocationList       RFC4523AttributeType
	SupportedAlgorithms       RFC4523AttributeType
)

func init() {

	UserCertificate = RFC4523AttributeType(`( 2.5.4.36 NAME 'userCertificate' DESC 'X.509 user certificate' EQUALITY certificateExactMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.8 X-ORIGIN 'RFC4523' )`)
	CACertificate = RFC4523AttributeType(`( 2.5.4.37 NAME 'cACertificate' DESC 'X.509 CA certificate' EQUALITY certificateExactMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.8 X-ORIGIN 'RFC4523' )`)
	CrossCertificatePair = RFC4523AttributeType(`( 2.5.4.40 NAME 'crossCertificatePair' DESC 'X.509 cross certificate pair' EQUALITY certificatePairExactMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.10 X-ORIGIN 'RFC4523' )`)
	CertificateRevocationList = RFC4523AttributeType(`( 2.5.4.39 NAME 'certificateRevocationList' DESC 'X.509 certificate revocation list' EQUALITY certificateListExactMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.9 X-ORIGIN 'RFC4523' )`)
	AuthorityRevocationList = RFC4523AttributeType(`( 2.5.4.38 NAME 'authorityRevocationList' DESC 'X.509 authority revocation list' EQUALITY certificateListExactMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.9 X-ORIGIN 'RFC4523' )`)
	DeltaRevocationList = RFC4523AttributeType(`( 2.5.4.53 NAME 'deltaRevocationList' DESC 'X.509 delta revocation list' EQUALITY certificateListExactMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.9 X-ORIGIN 'RFC4523' )`)
	SupportedAlgorithms = RFC4523AttributeType(`( 2.5.4.52 NAME 'supportedAlgorithms' DESC 'X.509 supported algorithms' EQUALITY algorithmIdentifierMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.49 X-ORIGIN 'RFC4523' )`)

	AllAttributeTypes = RFC4523AttributeTypes{
		UserCertificate,
		CACertificate,
		CrossCertificatePair,
		CertificateRevocationList,
		AuthorityRevocationList,
		DeltaRevocationList,
		SupportedAlgorithms,
	}
}

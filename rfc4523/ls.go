package rfc4523

/*
RFC4523LDAPSyntaxes is a slice type designed to store RFC4523LDAPSyntax instances.
*/
type RFC4523LDAPSyntaxes []RFC4523LDAPSyntax

/*
RFC4523Syntax is a struct type that stores the raw RFC4523 syntax definition, along with a boolean value indicative of whether the syntax is considered human-readable.
*/
type RFC4523LDAPSyntax string

/*
LDAPSyntaxes contains slices of all instances of RFC4523LDAPSyntax defined in this package.
*/
var AllLDAPSyntaxes RFC4523LDAPSyntaxes

var (
	Certificate                   RFC4523LDAPSyntax
	CertificateList               RFC4523LDAPSyntax
	CertificatePair               RFC4523LDAPSyntax
	SupportedAlgorithm            RFC4523LDAPSyntax
	CertificateExactAssertion     RFC4523LDAPSyntax
	CertificateAssertion          RFC4523LDAPSyntax
	CertificatePairExactAssertion RFC4523LDAPSyntax
	CertificatePairAssertion      RFC4523LDAPSyntax
	CertificateListExactAssertion RFC4523LDAPSyntax
	CertificateListAssertion      RFC4523LDAPSyntax
	AlgorithmIdentifier           RFC4523LDAPSyntax
)

func init() {

	Certificate = RFC4523LDAPSyntax(`( 1.3.6.1.4.1.1466.115.121.1.8 DESC 'Certificate' X-ORIGIN 'RFC4523' X-NOT-HUMAN-READABLE 'TRUE' )`)
	CertificateList = RFC4523LDAPSyntax(`( 1.3.6.1.4.1.1466.115.121.1.9 DESC 'Certificate List' X-ORIGIN 'RFC4523' X-NOT-HUMAN-READABLE 'TRUE' )`)
	CertificatePair = RFC4523LDAPSyntax(`( 1.3.6.1.4.1.1466.115.121.1.10 DESC 'Certificate Pair' X-ORIGIN 'RFC4523' X-NOT-HUMAN-READABLE 'TRUE' )`)
	SupportedAlgorithm = RFC4523LDAPSyntax(`( 1.3.6.1.4.1.1466.115.121.1.49 DESC 'Supported Algorithm' X-ORIGIN 'RFC4523' X-NOT-HUMAN-READABLE 'TRUE' )`)
	CertificateExactAssertion = RFC4523LDAPSyntax(`( 1.3.6.1.1.15.1 DESC 'X.509 Certificate Exact Assertion' X-ORIGIN 'RFC4523' X-NOT-HUMAN-READABLE 'TRUE' )`)
	CertificateAssertion = RFC4523LDAPSyntax(`( 1.3.6.1.1.15.2 DESC 'X.509 Certificate Assertion' X-ORIGIN 'RFC4523' )`)
	CertificatePairExactAssertion = RFC4523LDAPSyntax(`( 1.3.6.1.1.15.3 DESC 'X.509 Certificate Pair Exact Assertion' X-ORIGIN 'RFC4523' X-NOT-HUMAN-READABLE 'TRUE' )`)
	CertificatePairAssertion = RFC4523LDAPSyntax(`( 1.3.6.1.1.15.4 DESC 'X.509 Certificate Pair Assertion' X-ORIGIN 'RFC4523' X-NOT-HUMAN-READABLE 'TRUE' )`)
	CertificateListExactAssertion = RFC4523LDAPSyntax(`( 1.3.6.1.1.15.5 DESC 'X.509 Certificate List Exact Assertion' X-ORIGIN 'RFC4523' X-NOT-HUMAN-READABLE 'TRUE' )`)
	CertificateListAssertion = RFC4523LDAPSyntax(`( 1.3.6.1.1.15.6 DESC 'X.509 Certificate List Assertion' X-ORIGIN 'RFC4523' X-NOT-HUMAN-READABLE 'TRUE' )`)
	AlgorithmIdentifier = RFC4523LDAPSyntax(`( 1.3.6.1.1.15.7 DESC 'X.509 Algorithm Identifier' X-ORIGIN 'RFC4523' X-NOT-HUMAN-READABLE 'TRUE' )`)

	AllLDAPSyntaxes = RFC4523LDAPSyntaxes{
		Certificate,
		CertificateList,
		CertificatePair,
		SupportedAlgorithm,
		CertificateExactAssertion,
		CertificateAssertion,
		CertificatePairExactAssertion,
		CertificatePairAssertion,
		CertificateListExactAssertion,
		CertificateListAssertion,
		AlgorithmIdentifier,
	}
}

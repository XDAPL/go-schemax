package rfc4523

/*
LDAPSyntaxDefinitiones is a slice type designed to store LDAPSyntaxDefinition instances.
*/
type LDAPSyntaxDefinitions []LDAPSyntaxDefinition

/*
RFC4523Syntax is a struct type that stores the raw RFC4523 syntax definition, along with a boolean value indicative of whether the syntax is considered human-readable.
*/
type LDAPSyntaxDefinition string

/*
LDAPSyntaxes contains slices of all instances of LDAPSyntaxDefinition defined in this package.
*/
var AllLDAPSyntaxes LDAPSyntaxDefinitions

var (
	Certificate                   LDAPSyntaxDefinition
	CertificateList               LDAPSyntaxDefinition
	CertificatePair               LDAPSyntaxDefinition
	SupportedAlgorithm            LDAPSyntaxDefinition
	CertificateExactAssertion     LDAPSyntaxDefinition
	CertificateAssertion          LDAPSyntaxDefinition
	CertificatePairExactAssertion LDAPSyntaxDefinition
	CertificatePairAssertion      LDAPSyntaxDefinition
	CertificateListExactAssertion LDAPSyntaxDefinition
	CertificateListAssertion      LDAPSyntaxDefinition
	AlgorithmIdentifier           LDAPSyntaxDefinition
)

func init() {

	Certificate = LDAPSyntaxDefinition(`( 1.3.6.1.4.1.1466.115.121.1.8 DESC 'Certificate' X-ORIGIN 'RFC4523' X-NOT-HUMAN-READABLE 'TRUE' )`)
	CertificateList = LDAPSyntaxDefinition(`( 1.3.6.1.4.1.1466.115.121.1.9 DESC 'Certificate List' X-ORIGIN 'RFC4523' X-NOT-HUMAN-READABLE 'TRUE' )`)
	CertificatePair = LDAPSyntaxDefinition(`( 1.3.6.1.4.1.1466.115.121.1.10 DESC 'Certificate Pair' X-ORIGIN 'RFC4523' X-NOT-HUMAN-READABLE 'TRUE' )`)
	SupportedAlgorithm = LDAPSyntaxDefinition(`( 1.3.6.1.4.1.1466.115.121.1.49 DESC 'Supported Algorithm' X-ORIGIN 'RFC4523' X-NOT-HUMAN-READABLE 'TRUE' )`)
	CertificateExactAssertion = LDAPSyntaxDefinition(`( 1.3.6.1.1.15.1 DESC 'X.509 Certificate Exact Assertion' X-ORIGIN 'RFC4523' X-NOT-HUMAN-READABLE 'TRUE' )`)
	CertificateAssertion = LDAPSyntaxDefinition(`( 1.3.6.1.1.15.2 DESC 'X.509 Certificate Assertion' X-ORIGIN 'RFC4523' )`)
	CertificatePairExactAssertion = LDAPSyntaxDefinition(`( 1.3.6.1.1.15.3 DESC 'X.509 Certificate Pair Exact Assertion' X-ORIGIN 'RFC4523' X-NOT-HUMAN-READABLE 'TRUE' )`)
	CertificatePairAssertion = LDAPSyntaxDefinition(`( 1.3.6.1.1.15.4 DESC 'X.509 Certificate Pair Assertion' X-ORIGIN 'RFC4523' X-NOT-HUMAN-READABLE 'TRUE' )`)
	CertificateListExactAssertion = LDAPSyntaxDefinition(`( 1.3.6.1.1.15.5 DESC 'X.509 Certificate List Exact Assertion' X-ORIGIN 'RFC4523' X-NOT-HUMAN-READABLE 'TRUE' )`)
	CertificateListAssertion = LDAPSyntaxDefinition(`( 1.3.6.1.1.15.6 DESC 'X.509 Certificate List Assertion' X-ORIGIN 'RFC4523' X-NOT-HUMAN-READABLE 'TRUE' )`)
	AlgorithmIdentifier = LDAPSyntaxDefinition(`( 1.3.6.1.1.15.7 DESC 'X.509 Algorithm Identifier' X-ORIGIN 'RFC4523' X-NOT-HUMAN-READABLE 'TRUE' )`)

	AllLDAPSyntaxes = LDAPSyntaxDefinitions{
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

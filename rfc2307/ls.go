package rfc2307

/*
RFC2307LDAPSyntaxes is a slice type designed to store RFC2307LDAPSyntax instances.
*/
type RFC2307LDAPSyntaxes []RFC2307LDAPSyntax

/*
RFC2307Syntax is a struct type that stores the raw RFC2307 syntax definition, along with a boolean value indicative of whether the syntax is considered human-readable.
*/
type RFC2307LDAPSyntax string

/*
LDAPSyntaxes contains slices of all instances of RFC2307LDAPSyntax defined in this package.
*/
var AllLDAPSyntaxes RFC2307LDAPSyntaxes

var (
	NISNetgroupTripleSyntax RFC2307LDAPSyntax
	BootParameterSyntax     RFC2307LDAPSyntax
)

func init() {
	NISNetgroupTripleSyntax = RFC2307LDAPSyntax(`( 1.3.6.1.1.1.0.0 DESC 'RFC2307 NIS Netgroup Triple' X-ORIGIN 'RFC2307' )`)
	BootParameterSyntax = RFC2307LDAPSyntax(`( 1.3.6.1.1.1.0.1 DESC 'RFC2307 Boot Parameter' X-ORIGIN 'RFC2307' )`)

	AllLDAPSyntaxes = RFC2307LDAPSyntaxes{
		NISNetgroupTripleSyntax,
		BootParameterSyntax,
	}
}

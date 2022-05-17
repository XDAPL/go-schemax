package rfc4530

/*
RFC4530LDAPSyntaxes is a slice type designed to store RFC4530LDAPSyntax instances.
*/
type RFC4530LDAPSyntaxes []RFC4530LDAPSyntax

/*
RFC4530Syntax is a struct type that stores the raw RFC4530 syntax definition, along with a boolean value indicative of whether the syntax is considered human-readable.
*/
type RFC4530LDAPSyntax string

/*
LDAPSyntaxes contains slices of all instances of RFC4530LDAPSyntax defined in this package.
*/
var AllLDAPSyntaxes RFC4530LDAPSyntaxes

var (
	UUID			      RFC4530LDAPSyntax
)

func init() {

	UUID = RFC4530LDAPSyntax(`( 1.3.6.1.1.16.1 DESC 'UUID' X-ORIGIN 'RFC4530' )`)

	AllLDAPSyntaxes = RFC4530LDAPSyntaxes{
		UUID,
	}
}

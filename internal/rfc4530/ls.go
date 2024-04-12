package rfc4530

/*
LDAPSyntaxDefinitiones is a slice type designed to store LDAPSyntaxDefinition instances.
*/
type LDAPSyntaxDefinitions []LDAPSyntaxDefinition

/*
RFC4530Syntax is a struct type that stores the raw RFC4530 syntax definition, along with a boolean value indicative of whether the syntax is considered human-readable.
*/
type LDAPSyntaxDefinition string

/*
LDAPSyntaxes contains slices of all instances of LDAPSyntaxDefinition defined in this package.
*/
var AllLDAPSyntaxes LDAPSyntaxDefinitions

var (
	UUID LDAPSyntaxDefinition
)

func init() {

	UUID = LDAPSyntaxDefinition(`( 1.3.6.1.1.16.1 DESC 'UUID' X-ORIGIN 'RFC4530' )`)

	AllLDAPSyntaxes = LDAPSyntaxDefinitions{
		UUID,
	}
}

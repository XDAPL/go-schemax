package rfc2307

/*
LDAPSyntaxDefinitiones is a slice type designed to store LDAPSyntaxDefinition instances.
*/
type LDAPSyntaxDefinitions []LDAPSyntaxDefinition

func (r LDAPSyntaxDefinitions) Len() int {
	return len(r)
}

/*
LDAPSyntaxDefinition is a struct type that stores the raw RFC2307 syntax definition, along with a boolean value indicative of whether the syntax is considered human-readable.
*/
type LDAPSyntaxDefinition string

/*
AllLDAPSyntaxes contains slices of all instances of LDAPSyntaxDefinition defined in this package.
*/
var AllLDAPSyntaxes LDAPSyntaxDefinitions

var (
	NISNetgroupTripleSyntax LDAPSyntaxDefinition
	BootParameterSyntax     LDAPSyntaxDefinition
)

func init() {
	NISNetgroupTripleSyntax = LDAPSyntaxDefinition(`( nisSchema.0.0 DESC 'nisNetgroupTripleSyntax' X-ORIGIN 'RFC2307' )`)
	BootParameterSyntax = LDAPSyntaxDefinition(`( nisSchema.0.1 DESC 'bootParameterSyntax' X-ORIGIN 'RFC2307' )`)

	AllLDAPSyntaxes = LDAPSyntaxDefinitions{
		NISNetgroupTripleSyntax,
		BootParameterSyntax,
	}
}

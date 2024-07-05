package rfc2307

type AttributeTypeDefinitions []AttributeTypeDefinition
type AttributeTypeDefinition string

func (r AttributeTypeDefinitions) Len() int {
	return len(r)
}

var (
	AllAttributeTypes AttributeTypeDefinitions
)

var (
	UIDNumber         AttributeTypeDefinition
	GIDNumber         AttributeTypeDefinition
	Gecos             AttributeTypeDefinition
	HomeDirectory     AttributeTypeDefinition
	LoginShell        AttributeTypeDefinition
	ShadowLastChange  AttributeTypeDefinition
	ShadowMin         AttributeTypeDefinition
	ShadowMax         AttributeTypeDefinition
	ShadowWarning     AttributeTypeDefinition
	ShadowInactive    AttributeTypeDefinition
	ShadowExpire      AttributeTypeDefinition
	ShadowFlag        AttributeTypeDefinition
	MemberUID         AttributeTypeDefinition
	MemberNISNetgroup AttributeTypeDefinition
	NISNetgroupTriple AttributeTypeDefinition
	IPServicePort     AttributeTypeDefinition
	IPServiceProtocol AttributeTypeDefinition
	IPProtocolNumber  AttributeTypeDefinition
	ONCRPCNumber      AttributeTypeDefinition
	IPHostNumber      AttributeTypeDefinition
	IPNetworkNumber   AttributeTypeDefinition
	IPNetmaskNumber   AttributeTypeDefinition
	MACAddress        AttributeTypeDefinition
	BootParameter     AttributeTypeDefinition
	BootFile          AttributeTypeDefinition
	NISMapName        AttributeTypeDefinition
	NISMapEntry       AttributeTypeDefinition
)

func init() {

	UIDNumber = AttributeTypeDefinition(`( nisSchema.1.0 NAME 'uidNumber' DESC 'An integer uniquely identifying a user in an administrative domain' EQUALITY integerMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.27 SINGLE-VALUE X-ORIGIN 'RFC2307' )`)
	GIDNumber = AttributeTypeDefinition(`( nisSchema.1.1 NAME 'gidNumber' DESC 'An integer uniquely identifying a group in an administrative domain' EQUALITY integerMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.27 SINGLE-VALUE X-ORIGIN 'RFC2307' )`)
	Gecos = AttributeTypeDefinition(`( nisSchema.1.2 NAME 'gecos' DESC 'The GECOS field; the common name' EQUALITY caseIgnoreIA5Match SUBSTR caseIgnoreIA5SubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.26 SINGLE-VALUE X-ORIGIN 'RFC2307' )`)
	HomeDirectory = AttributeTypeDefinition(`( nisSchema.1.3 NAME 'homeDirectory' DESC 'The absolute path to the home directory' EQUALITY caseExactIA5Match SYNTAX 1.3.6.1.4.1.1466.115.121.1.26 SINGLE-VALUE X-ORIGIN 'RFC2307' )`)
	LoginShell = AttributeTypeDefinition(`( nisSchema.1.4 NAME 'loginShell' DESC 'The path to the login shell' EQUALITY caseExactIA5Match SYNTAX 1.3.6.1.4.1.1466.115.121.1.26 SINGLE-VALUE X-ORIGIN 'RFC2307' )`)
	ShadowLastChange = AttributeTypeDefinition(`( nisSchema.1.5 NAME 'shadowLastChange' EQUALITY integerMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.27 SINGLE-VALUE X-ORIGIN 'RFC2307' )`)
	ShadowMin = AttributeTypeDefinition(`( nisSchema.1.6 NAME 'shadowMin' EQUALITY integerMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.27 SINGLE-VALUE X-ORIGIN 'RFC2307' )`)
	ShadowMax = AttributeTypeDefinition(`( nisSchema.1.7 NAME 'shadowMax' EQUALITY integerMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.27 SINGLE-VALUE X-ORIGIN 'RFC2307' )`)
	ShadowWarning = AttributeTypeDefinition(`( nisSchema.1.8 NAME 'shadowWarning' EQUALITY integerMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.27 SINGLE-VALUE X-ORIGIN 'RFC2307' )`)
	ShadowInactive = AttributeTypeDefinition(`( nisSchema.1.9 NAME 'shadowInactive' EQUALITY integerMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.27 SINGLE-VALUE X-ORIGIN 'RFC2307' )`)
	ShadowExpire = AttributeTypeDefinition(`( nisSchema.1.10 NAME 'shadowExpire' EQUALITY integerMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.27 SINGLE-VALUE X-ORIGIN 'RFC2307' )`)
	ShadowFlag = AttributeTypeDefinition(`( nisSchema.1.11 NAME 'shadowFlag' EQUALITY integerMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.27 SINGLE-VALUE X-ORIGIN 'RFC2307' )`)
	MemberUID = AttributeTypeDefinition(`( nisSchema.1.12 NAME 'memberUid' EQUALITY caseExactIA5Match SUBSTR caseExactIA5SubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.26 X-ORIGIN 'RFC2307' )`)
	MemberNISNetgroup = AttributeTypeDefinition(`( nisSchema.1.13 NAME 'memberNisNetgroup' EQUALITY caseExactIA5Match SUBSTR caseExactIA5SubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.26 X-ORIGIN 'RFC2307' )`)
	NISNetgroupTriple = AttributeTypeDefinition(`( nisSchema.1.14 NAME 'nisNetgroupTriple' DESC 'Netgroup triple' SYNTAX 'nisNetgroupTripleSyntax' X-ORIGIN 'RFC2307' )`)
	IPServicePort = AttributeTypeDefinition(`( nisSchema.1.15 NAME 'ipServicePort' EQUALITY integerMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.27 SINGLE-VALUE X-ORIGIN 'RFC2307' )`)
	IPServiceProtocol = AttributeTypeDefinition(`( nisSchema.1.16 NAME 'ipServiceProtocol' SUP name X-ORIGIN 'RFC2307' )`)
	IPProtocolNumber = AttributeTypeDefinition(`( nisSchema.1.17 NAME 'ipProtocolNumber' EQUALITY integerMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.27 SINGLE-VALUE X-ORIGIN 'RFC2307' )`)
	ONCRPCNumber = AttributeTypeDefinition(`( nisSchema.1.18 NAME 'oncRpcNumber' EQUALITY integerMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.27 SINGLE-VALUE X-ORIGIN 'RFC2307' )`)
	IPHostNumber = AttributeTypeDefinition(`( nisSchema.1.19 NAME 'ipHostNumber' DESC 'IP address as a dotted decimal, eg. 192.168.1.1, omitting leading zeros' EQUALITY caseIgnoreIA5Match SYNTAX 1.3.6.1.4.1.1466.115.121.1.26{128} X-ORIGIN 'RFC2307' )`)
	IPNetworkNumber = AttributeTypeDefinition(`( nisSchema.1.20 NAME 'ipNetworkNumber' DESC 'IP network as a dotted decimal, eg. 192.168, omitting leading zeros' EQUALITY caseIgnoreIA5Match SYNTAX 1.3.6.1.4.1.1466.115.121.1.26{128} SINGLE-VALUE X-ORIGIN 'RFC2307' )`)
	IPNetmaskNumber = AttributeTypeDefinition(`( nisSchema.1.21 NAME 'ipNetmaskNumber' DESC 'IP netmask as a dotted decimal, eg. 255.255.255.0, omitting leading zeros' EQUALITY caseIgnoreIA5Match SYNTAX 1.3.6.1.4.1.1466.115.121.1.26{128} SINGLE-VALUE X-ORIGIN 'RFC2307' )`)
	MACAddress = AttributeTypeDefinition(`( nisSchema.1.22 NAME 'macAddress' DESC 'MAC address in maximal, colon separated hex notation, eg. 00:00:92:90:ee:e2' EQUALITY caseIgnoreIA5Match SYNTAX 1.3.6.1.4.1.1466.115.121.1.26{128} X-ORIGIN 'RFC2307' )`)
	BootParameter = AttributeTypeDefinition(`( nisSchema.1.23 NAME 'bootParameter' DESC 'rpc.bootparamd parameter' SYNTAX 'bootParameterSyntax' X-ORIGIN 'RFC2307' )`)
	BootFile = AttributeTypeDefinition(`( nisSchema.1.24 NAME 'bootFile' DESC 'Boot image name' EQUALITY caseExactIA5Match SYNTAX 1.3.6.1.4.1.1466.115.121.1.26 X-ORIGIN 'RFC2307' )`)
	NISMapName = AttributeTypeDefinition(`( nisSchema.1.26 NAME 'nisMapName' SUP name X-ORIGIN 'RFC2307' )`)
	NISMapEntry = AttributeTypeDefinition(`( nisSchema.1.27 NAME 'nisMapEntry' EQUALITY caseExactIA5Match SUBSTR caseExactIA5SubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.26{1024} SINGLE-VALUE X-ORIGIN 'RFC2307' )`)

	AllAttributeTypes = AttributeTypeDefinitions{
		UIDNumber,
		GIDNumber,
		Gecos,
		HomeDirectory,
		LoginShell,
		ShadowLastChange,
		ShadowMin,
		ShadowMax,
		ShadowWarning,
		ShadowInactive,
		ShadowExpire,
		ShadowFlag,
		MemberUID,
		MemberNISNetgroup,
		NISNetgroupTriple,
		IPServicePort,
		IPServiceProtocol,
		IPProtocolNumber,
		ONCRPCNumber,
		IPHostNumber,
		IPNetworkNumber,
		IPNetmaskNumber,
		MACAddress,
		BootParameter,
		BootFile,
		NISMapName,
		NISMapEntry,
	}
}

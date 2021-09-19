package rfc2307

type RFC2307AttributeTypes []RFC2307AttributeType
type RFC2307AttributeType string

var (
	AllAttributeTypes RFC2307AttributeTypes
)

var (
	UIDNumber         RFC2307AttributeType
	GIDNumber         RFC2307AttributeType
	Gecos             RFC2307AttributeType
	HomeDirectory     RFC2307AttributeType
	LoginShell        RFC2307AttributeType
	ShadowLastChange  RFC2307AttributeType
	ShadowMin         RFC2307AttributeType
	ShadowMax         RFC2307AttributeType
	ShadowWarning     RFC2307AttributeType
	ShadowInactive    RFC2307AttributeType
	ShadowExpire      RFC2307AttributeType
	ShadowFlag        RFC2307AttributeType
	MemberUID         RFC2307AttributeType
	MemberNISNetgroup RFC2307AttributeType
	NISNetgroupTriple RFC2307AttributeType
	IPServicePort     RFC2307AttributeType
	IPServiceProtocol RFC2307AttributeType
	IPProtocolNumber  RFC2307AttributeType
	ONCRPCNumber      RFC2307AttributeType
	IPHostNumber      RFC2307AttributeType
	IPNetworkNumber   RFC2307AttributeType
	IPNetmaskNumber   RFC2307AttributeType
	MACAddress        RFC2307AttributeType
	BootParameter     RFC2307AttributeType
	BootFile          RFC2307AttributeType
	NISMapName        RFC2307AttributeType
	NISMapEntry       RFC2307AttributeType
)

func init() {

	UIDNumber = RFC2307AttributeType(`( 1.3.6.1.1.1.1.0 NAME 'uidNumber' DESC 'An integer uniquely identifying a user in an administrative domain' EQUALITY integerMatch SYNTAX 'INTEGER' SINGLE-VALUE X-ORIGIN 'RFC2307' )`)
	GIDNumber = RFC2307AttributeType(`( 1.3.6.1.1.1.1.1 NAME 'gidNumber' DESC 'An integer uniquely identifying a group in an administrative domain' EQUALITY integerMatch SYNTAX 'INTEGER' SINGLE-VALUE X-ORIGIN 'RFC2307' )`)
	Gecos = RFC2307AttributeType(`( 1.3.6.1.1.1.1.2 NAME 'gecos' DESC 'The GECOS field; the common name' EQUALITY caseIgnoreIA5Match SUBSTRINGS caseIgnoreIA5SubstringsMatch SYNTAX 'IA5String' SINGLE-VALUE X-ORIGIN 'RFC2307' )`)
	HomeDirectory = RFC2307AttributeType(`( 1.3.6.1.1.1.1.3 NAME 'homeDirectory' DESC 'The absolute path to the home directory' EQUALITY caseExactIA5Match SYNTAX 'IA5String' SINGLE-VALUE X-ORIGIN 'RFC2307' )`)
	LoginShell = RFC2307AttributeType(`( 1.3.6.1.1.1.1.4 NAME 'loginShell' DESC 'The path to the login shell' EQUALITY caseExactIA5Match SYNTAX 'IA5String' SINGLE-VALUE X-ORIGIN 'RFC2307' )`)
	ShadowLastChange = RFC2307AttributeType(`( 1.3.6.1.1.1.1.5 NAME 'shadowLastChange' EQUALITY integerMatch SYNTAX 'INTEGER' SINGLE-VALUE X-ORIGIN 'RFC2307' )`)
	ShadowMin = RFC2307AttributeType(`( 1.3.6.1.1.1.1.6 NAME 'shadowMin' EQUALITY integerMatch SYNTAX 'INTEGER' SINGLE-VALUE X-ORIGIN 'RFC2307' )`)
	ShadowMax = RFC2307AttributeType(`( 1.3.6.1.1.1.1.7 NAME 'shadowMax' EQUALITY integerMatch SYNTAX 'INTEGER' SINGLE-VALUE X-ORIGIN 'RFC2307' )`)
	ShadowWarning = RFC2307AttributeType(`( 1.3.6.1.1.1.1.8 NAME 'shadowWarning' EQUALITY integerMatch SYNTAX 'INTEGER' SINGLE-VALUE X-ORIGIN 'RFC2307' )`)
	ShadowInactive = RFC2307AttributeType(`( 1.3.6.1.1.1.1.9 NAME 'shadowInactive' EQUALITY integerMatch SYNTAX 'INTEGER' SINGLE-VALUE X-ORIGIN 'RFC2307' )`)
	ShadowExpire = RFC2307AttributeType(`( 1.3.6.1.1.1.1.10 NAME 'shadowExpire' EQUALITY integerMatch SYNTAX 'INTEGER' SINGLE-VALUE X-ORIGIN 'RFC2307' )`)
	ShadowFlag = RFC2307AttributeType(`( 1.3.6.1.1.1.1.11 NAME 'shadowFlag' EQUALITY integerMatch SYNTAX 'INTEGER' SINGLE-VALUE X-ORIGIN 'RFC2307' )`)
	MemberUID = RFC2307AttributeType(`( 1.3.6.1.1.1.1.12 NAME 'memberUid' EQUALITY caseExactIA5Match SUBSTRINGS caseExactIA5SubstringsMatch SYNTAX 'IA5String' X-ORIGIN 'RFC2307' )`)
	MemberNISNetgroup = RFC2307AttributeType(`( 1.3.6.1.1.1.1.13 NAME 'memberNisNetgroup' EQUALITY caseExactIA5Match SUBSTRINGS caseExactIA5SubstringsMatch SYNTAX 'IA5String' X-ORIGIN 'RFC2307' )`)
	NISNetgroupTriple = RFC2307AttributeType(`( 1.3.6.1.1.1.1.14 NAME 'nisNetgroupTriple' DESC 'Netgroup triple' SYNTAX 'nisNetgroupTripleSyntax' X-ORIGIN 'RFC2307' )`)
	IPServicePort = RFC2307AttributeType(`( 1.3.6.1.1.1.1.15 NAME 'ipServicePort' EQUALITY integerMatch SYNTAX 'INTEGER' SINGLE-VALUE X-ORIGIN 'RFC2307' )`)
	IPServiceProtocol = RFC2307AttributeType(`( 1.3.6.1.1.1.1.16 NAME 'ipServiceProtocol' SUP name X-ORIGIN 'RFC2307' )`)
	IPProtocolNumber = RFC2307AttributeType(`( 1.3.6.1.1.1.1.17 NAME 'ipProtocolNumber' EQUALITY integerMatch SYNTAX 'INTEGER' SINGLE-VALUE X-ORIGIN 'RFC2307' )`)
	ONCRPCNumber = RFC2307AttributeType(`( 1.3.6.1.1.1.1.18 NAME 'oncRpcNumber' EQUALITY integerMatch SYNTAX 'INTEGER' SINGLE-VALUE X-ORIGIN 'RFC2307' )`)
	IPHostNumber = RFC2307AttributeType(`( 1.3.6.1.1.1.1.19 NAME 'ipHostNumber' DESC 'IP address as a dotted decimal, eg. 192.168.1.1, omitting leading zeros' EQUALITY caseIgnoreIA5Match SYNTAX 'IA5String{128}' X-ORIGIN 'RFC2307' )`)
	IPNetworkNumber = RFC2307AttributeType(`( 1.3.6.1.1.1.1.20 NAME 'ipNetworkNumber' DESC 'IP network as a dotted decimal, eg. 192.168, omitting leading zeros' EQUALITY caseIgnoreIA5Match SYNTAX 'IA5String{128}' SINGLE-VALUE X-ORIGIN 'RFC2307' )`)
	IPNetmaskNumber = RFC2307AttributeType(`( 1.3.6.1.1.1.1.21 NAME 'ipNetmaskNumber' DESC 'IP netmask as a dotted decimal, eg. 255.255.255.0, omitting leading zeros' EQUALITY caseIgnoreIA5Match SYNTAX 'IA5String{128}' SINGLE-VALUE X-ORIGIN 'RFC2307' )`)
	MACAddress = RFC2307AttributeType(`( 1.3.6.1.1.1.1.22 NAME 'macAddress' DESC 'MAC address in maximal, colon separated hex notation, eg. 00:00:92:90:ee:e2' EQUALITY caseIgnoreIA5Match SYNTAX 'IA5String{128}' X-ORIGIN 'RFC2307' )`)
	BootParameter = RFC2307AttributeType(`( 1.3.6.1.1.1.1.23 NAME 'bootParameter' DESC 'rpc.bootparamd parameter' SYNTAX 'bootParameterSyntax' X-ORIGIN 'RFC2307' )`)
	BootFile = RFC2307AttributeType(`( 1.3.6.1.1.1.1.24 NAME 'bootFile' DESC 'Boot image name' EQUALITY caseExactIA5Match SYNTAX 'IA5String' X-ORIGIN 'RFC2307' )`)
	NISMapName = RFC2307AttributeType(`( 1.3.6.1.1.1.1.26 NAME 'nisMapName' SUP name X-ORIGIN 'RFC2307' )`)
	NISMapEntry = RFC2307AttributeType(`( 1.3.6.1.1.1.1.27 NAME 'nisMapEntry' EQUALITY caseExactIA5Match SUBSTRINGS caseExactIA5SubstringsMatch SYNTAX 'IA5String{1024}' SINGLE-VALUE X-ORIGIN 'RFC2307' )`)

	AllAttributeTypes = RFC2307AttributeTypes{
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

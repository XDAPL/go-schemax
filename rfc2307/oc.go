package rfc2307

type RFC2307ObjectClasses []RFC2307ObjectClass
type RFC2307ObjectClass string

var (
	AllObjectClasses RFC2307ObjectClasses
)

var (
	POSIXAccount   RFC2307ObjectClass
	ShadowAccount  RFC2307ObjectClass
	POSIXGroup     RFC2307ObjectClass
	IPService      RFC2307ObjectClass
	IPProtocol     RFC2307ObjectClass
	ONCRPC         RFC2307ObjectClass
	IPHost         RFC2307ObjectClass
	IPNetwork      RFC2307ObjectClass
	NISNetgroup    RFC2307ObjectClass
	NISMap         RFC2307ObjectClass
	NISObject      RFC2307ObjectClass
	IEEE802Device  RFC2307ObjectClass
	BootableDevice RFC2307ObjectClass
)

func init() {

	POSIXAccount = RFC2307ObjectClass(`( 1.3.6.1.1.1.2.0 NAME 'posixAccount' SUP top AUXILIARY DESC 'Abstraction of an account with POSIX attributes' MUST ( cn $ uid $ uidNumber $ gidNumber $ homeDirectory ) MAY ( userPassword $ loginShell $ gecos $ description ) X-ORIGIN 'RFC2307' )`)
	ShadowAccount = RFC2307ObjectClass(`( 1.3.6.1.1.1.2.1 NAME 'shadowAccount' SUP top AUXILIARY DESC 'Additional attributes for shadow passwords' MUST uid MAY ( userPassword $ shadowLastChange $ shadowMin $ shadowMax $ shadowWarning $ shadowInactive $ shadowExpire $ shadowFlag $ description ) X-ORIGIN 'RFC2307' )`)
	POSIXGroup = RFC2307ObjectClass(`( 1.3.6.1.1.1.2.2 NAME 'posixGroup' SUP top STRUCTURAL DESC 'Abstraction of a group of accounts' MUST ( cn $ gidNumber ) MAY ( userPassword $ memberUid $ description ) X-ORIGIN 'RFC2307' )`)
	IPService = RFC2307ObjectClass(`( 1.3.6.1.1.1.2.3 NAME 'ipService' SUP top STRUCTURAL DESC 'Abstraction an Internet Protocol service. Maps an IP port and protocol (such as tcp or udp) to one or more names; the distinguished value of the cn attribute denotes the service's canonical name' MUST ( cn $ ipServicePort $ ipServiceProtocol ) MAY ( description ) X-ORIGIN 'RFC2307' )`)
	IPProtocol = RFC2307ObjectClass(`( 1.3.6.1.1.1.2.4 NAME 'ipProtocol' SUP top STRUCTURAL DESC 'Abstraction of an IP protocol. Maps a protocol number to one or more names. The distinguished value of the cn attribute denotes the protocol's canonical name' MUST ( cn $ ipProtocolNumber $ description ) MAY description X-ORIGIN 'RFC2307' )`)
	ONCRPC = RFC2307ObjectClass(`( 1.3.6.1.1.1.2.5 NAME 'oncRpc' SUP top STRUCTURAL DESC 'Abstraction of an Open Network Computing (ONC) [RFC1057] Remote Procedure Call (RPC) binding. This class maps an ONC RPC number to a name. The distinguished value of the cn attribute denotes the RPC service's canonical name' MUST ( cn $ oncRpcNumber $ description ) MAY description X-ORIGIN 'RFC2307' )`)
	IPHost = RFC2307ObjectClass(`( 1.3.6.1.1.1.2.6 NAME 'ipHost' SUP top AUXILIARY DESC 'Abstraction of a host, an IP device. The distinguished value of the cn attribute denotes the host's canonical name. Device SHOULD be used as a structural class' MUST ( cn $ ipHostNumber ) MAY ( l $ description $ manager ) X-ORIGIN 'RFC2307' )`)
	IPNetwork = RFC2307ObjectClass(`( 1.3.6.1.1.1.2.7 NAME 'ipNetwork' SUP top STRUCTURAL DESC 'Abstraction of a network. The distinguished value of the cn attribute denotes the network's canonical name' MUST ( cn $ ipNetworkNumber ) MAY ( ipNetmaskNumber $ l $ description $ manager ) X-ORIGIN 'RFC2307' )`)
	NISNetgroup = RFC2307ObjectClass(`( 1.3.6.1.1.1.2.8 NAME 'nisNetgroup' SUP top STRUCTURAL DESC 'Abstraction of a netgroup. May refer to other netgroups' MUST cn MAY ( nisNetgroupTriple $ memberNisNetgroup $ description ) X-ORIGIN 'RFC2307' )`)
	NISMap = RFC2307ObjectClass(`( 1.3.6.1.1.1.2.09 NAME 'nisMap' SUP top STRUCTURAL DESC 'A generic abstraction of a NIS map' MUST nisMapName MAY description X-ORIGIN 'RFC2307' )`)
	NISObject = RFC2307ObjectClass(`( 1.3.6.1.1.1.2.10 NAME 'nisObject' SUP top STRUCTURAL DESC 'An entry in a NIS map' MUST ( cn $ nisMapEntry $ nisMapName ) MAY description X-ORIGIN 'RFC2307' )`)
	IEEE802Device = RFC2307ObjectClass(`( 1.3.6.1.1.1.2.11 NAME 'ieee802Device' SUP top AUXILIARY DESC 'A device with a MAC address; device SHOULD be used as a structural class' MAY macAddress X-ORIGIN 'RFC2307' )`)
	BootableDevice = RFC2307ObjectClass(`( 1.3.6.1.1.1.2.12 NAME 'bootableDevice' SUP top AUXILIARY DESC 'A device with boot parameters; device SHOULD be used as a structural class' MAY ( bootFile $ bootParameter ) X-ORIGIN 'RFC2307' )`)

	AllObjectClasses = RFC2307ObjectClasses{
		POSIXAccount,
		ShadowAccount,
		POSIXGroup,
		IPService,
		IPProtocol,
		ONCRPC,
		IPHost,
		IPNetwork,
		NISNetgroup,
		NISMap,
		NISObject,
		IEEE802Device,
		BootableDevice,
	}

}

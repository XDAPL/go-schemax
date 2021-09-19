package rfc2079

type RFC2079ObjectClasses []RFC2079ObjectClass
type RFC2079ObjectClass string

var (
	AllObjectClasses RFC2079ObjectClasses
)

var (
	LabeledURIObject RFC2079ObjectClass
)

func init() {

	LabeledURIObject = RFC2079ObjectClass(`( 1.3.6.1.4.1.250.3.15 NAME 'labeledURIObject' DESC 'RFC2079: object that contains the URI attribute type' SUP top AUXILIARY MAY labeledURI X-ORIGIN 'RFC2079' )`)

	AllObjectClasses = RFC2079ObjectClasses{
		LabeledURIObject,
	}

}

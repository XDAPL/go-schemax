package rfc3671

type ObjectClassDefinitions []ObjectClassDefinition
type ObjectClassDefinition string

var CollectiveAttributeSubentry ObjectClassDefinition
var AllObjectClasses ObjectClassDefinitions

func (r ObjectClassDefinition) String() string {
	return string(r)
}

func init() {
	CollectiveAttributeSubentry = ObjectClassDefinition(`( 2.5.17.2
                NAME 'collectiveAttributeSubentry'
                AUXILIARY )`)

	AllObjectClasses = ObjectClassDefinitions{
		CollectiveAttributeSubentry,
	}
}

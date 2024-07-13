package rfc2589

type ObjectClassDefinitions []ObjectClassDefinition
type ObjectClassDefinition string

func (r ObjectClassDefinitions) Len() int {
	return len(r)
}

var (
	AllObjectClasses ObjectClassDefinitions
)

var (
	DynamicObject ObjectClassDefinition
)

func (r ObjectClassDefinition) String() string {
	return string(r)
}

func init() {

	DynamicObject = ObjectClassDefinition(`( 1.3.6.1.4.1.1466.101.119.2 NAME 'dynamicObject' DESC 'RFC2589: Dynamic Object' SUP top AUXILIARY X-ORIGIN 'RFC2589' )`)

	AllObjectClasses = ObjectClassDefinitions{
		DynamicObject,
	}

}

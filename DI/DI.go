package DI

const (
	ResourceSingleton = iota
	ResourceInstance
)

var container map[string]interface{}
var resourceTypes map[string]int

func init() {
	container = make(map[string]interface{})
	resourceTypes = make(map[string]int)
}

func Make(name string) interface{} {
	if resource, ok := container[name]; ok {
		if resourceTypes[name] == ResourceSingleton {
			return resource
		} else if resourceTypes[name] == ResourceInstance {
			return resource.(func() interface{})()
		}
	}

	return nil
}

func Singleton(name string, resource interface{}) {
	resourceTypes[name] = ResourceSingleton
	container[name] = resource
}

func Instance(name string, factory func() interface{}) {
	resourceTypes[name] = ResourceInstance
	container[name] = factory
}

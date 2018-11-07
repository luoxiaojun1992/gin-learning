package DI

var container map[string]interface{}

// Init container
func init() {
	container = make(map[string]interface{})
}

// Resolving dependency by resource name
func Resolve(name string) interface{} {
	if resource, ok := container[name]; ok {
		return resource.(func() interface{})()
	}

	return nil
}

// Injecting singleton resource
func Singleton(name string, resource interface{}) {
	Instance(name, func() interface{} {return resource})
}

// Injecting instance resource
func Instance(name string, factory func() interface{}) {
	container[name] = factory
}

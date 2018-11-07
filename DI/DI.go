package DI

var container map[string]interface{}

func init() {
	container = make(map[string]interface{})
}

func Make(name string) interface{} {
	if resource, ok := container[name]; ok {
		return resource.(func() interface{})()
	}

	return nil
}

func Singleton(name string, resource interface{}) {
	Instance(name, func() interface{} {return resource})
}

func Instance(name string, factory func() interface{}) {
	container[name] = factory
}

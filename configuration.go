package configuration

type Configuration = map[interface{}]interface{}

type Modifier func(Configuration) error

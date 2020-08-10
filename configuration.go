package configuration

type Configuration = map[string]interface{}

type Modifier func(Configuration) error

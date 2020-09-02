package configuration

import (
	"os"
	"strings"
	"sync"
)

const (
	Prefix = "SERVICE"
	Sep    = "="
)

type Configuration = map[string]interface{}

type Modifier func(Configuration) error

var (
	mutable Configuration
	once    = new(sync.Once)
)

func fromEnv() {
	mutable = Configuration{}
	for _, eq := range os.Environ() {
		tmp := strings.SplitN(eq, Sep, 2)
		if len(tmp) != 2 {
			continue
		}
		if strings.HasPrefix(tmp[0], Prefix) {
            mutable[strings.ReplaceAll(strings.ToLower(tmp[0]),"_",".")] = tmp[1]
		}
	}
}

func init() {
	once.Do(func() {
		fromEnv()
	})
}

func FromEnv() Configuration {
	return mutable
}

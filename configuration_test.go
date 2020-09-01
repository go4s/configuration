package configuration

import (
	"log"
	"os"
	"strings"
	"testing"
)

func TestFromEnv(t *testing.T) {
	for _, v := range os.Environ() {
		log.Println(strings.SplitN(v, "=", 2))
	}
}

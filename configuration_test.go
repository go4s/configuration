package configuration

import (
    "github.com/go-playground/assert/v2"
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

func TestFromFile(t *testing.T) {
    assert.Equal(t, "10=10", mutable["service.app.addr"])
}

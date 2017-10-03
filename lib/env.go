package lib

import (
    "os"
    "strings"
)

// UnmarshalEnv takes the Environment and puts it into a struct
func UnmarshalEnv() map[string]string {
	env := make(map[string]string)
	for _, i := range os.Environ() {
		sep := strings.Index(i, "=")
		env[i[0:sep]] = i[sep+1:]
	}
	return env
}

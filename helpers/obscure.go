package helpers

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/minskylab/meta-core/services/structures"
)

func ObscureMiddleString(val string, edgeSize int, maxMiddle int, obsChar string) string {
	if edgeSize*2 > len(val) {
		var builder strings.Builder
		for i := 0; i < len(val); i++ {
			builder.WriteString(obsChar)
		}
		return builder.String()
	}
	middleSize := len(val[edgeSize : len(val)-edgeSize])
	if middleSize > maxMiddle {
		middleSize = maxMiddle
	}

	var builder strings.Builder
	for i := 0; i < middleSize; i++ {
		builder.WriteString(obsChar)
	}

	return fmt.Sprint(val[:edgeSize], builder.String(), val[len(val)-edgeSize:])
}

func ObscureProcess(process *structures.Process) (*structures.Process, error) {
	process.Token = ObscureMiddleString(process.Token, 2, 5, "*")

	if process.Stack != nil {
		process.Stack.KeyPair.PrivateKey = ""
	}

	for _, task := range process.Definition.Tasks {
		for _, k := range reflect.ValueOf(task.Environment).MapKeys() {
			k := k.String()
			tVal := task.Environment[k]

			if val, ok := tVal.(string); ok {
				task.Environment[k] = ObscureMiddleString(val, 2, 5, "*")
			}
		}
	}

	if process.Definition.Credentials != nil {
		for _, cred := range process.Definition.Credentials {
			cred.Username = ObscureMiddleString(cred.Username, 2, 5, "*")
			cred.Password = ObscureMiddleString(cred.Password, 2, 5, "*")
			cred.Registry = ObscureMiddleString(cred.Registry, 2, 5, "*")
		}
	}

	return process, nil
}

package helpers

import (
	"encoding/base64"
	"encoding/json"

	"github.com/minskylab/meta-core/services/structures"
)

func InputContextToB64(context structures.InputContext) (string, error) {
	contextBytes, err := json.Marshal(context)
	if err != nil {
		return "", err
	}
	b64Bytes := base64.StdEncoding.EncodeToString(contextBytes)
	return b64Bytes, nil
}

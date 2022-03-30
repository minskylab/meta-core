package helpers

import (
	"encoding/base64"
	"encoding/json"

	log "github.com/sirupsen/logrus"

	"github.com/minskylab/meta-core/services/structures"
)

func InputContextToB64(context structures.InputContext) (string, error) {
	contextBytes, err := json.Marshal(context)
	log.Debugln(contextBytes)
	if err != nil {
		return "", err
	}
	b64Bytes := base64.StdEncoding.EncodeToString(contextBytes)
	return b64Bytes, nil
}

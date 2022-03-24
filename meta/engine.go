package meta

import (
	"net/http"

	"github.com/minskylab/meta-core/ent"
)

type Engine struct {
	entClient  *ent.Client
	httpClient *http.Client
}

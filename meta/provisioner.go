package meta

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/minskylab/meta-core/ent"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

// Provisioner is a standard for provider communication protocol.

type Provisioner interface {
	LaunchStackRequest(ctx context.Context, stack *ent.Stack) (*http.Request, error)
	PurgeStackRequest(ctx context.Context, stack *ent.Stack) (*http.Request, error)
	GetStacksRequest(ctx context.Context) (*http.Request, error)
	GetStackRequest(ctx context.Context, stack *ent.Stack) (*http.Request, error)
}

// TODO: Implement the meta-provider-standard based on http requests.

// POST /stack
// DELETE /stack
// GET /stacks
// GET /stack/:id

func (e *Engine) LaunchStackRequest(ctx context.Context, stack *ent.Stack) (*http.Request, error) {
	provider, err := e.entClient.Stack.QueryDeployment(stack).QueryProvider().Only(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	url := strings.TrimRight(provider.Hostname, "/") + "/stack"

	buff := new(bytes.Buffer)

	if err := json.NewEncoder(buff).Encode(stack); err != nil {
		return nil, errors.WithStack(err)
	}

	req, err := http.NewRequest(http.MethodPost, url, buff)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+provider.Token)

	return req, nil
}

func (e *Engine) PurgeStackRequest(ctx context.Context, id uuid.UUID) (*http.Request, error) {
	panic("implement me")
}

func (e *Engine) GetStacksRequest(ctx context.Context) (*http.Request, error) {
	panic("implement me")
}

func (e *Engine) GetStackRequest(ctx context.Context, id uuid.UUID) (*http.Request, error) {
	panic("implement me")
}

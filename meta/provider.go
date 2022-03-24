package meta

import (
	"context"

	"github.com/minskylab/meta-core/ent"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

// type Provider interface{}

func (e *Engine) LaunchStack(ctx context.Context, stack *ent.Stack) (*ent.Stack, error) {
	req, err := e.LaunchStackRequest(ctx, stack)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	res, err := e.httpClient.Do(req)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if res.StatusCode != 200 {
		// ...
	}

	panic("implement me")
}

func (e *Engine) PurgeStack(ctx context.Context, id uuid.UUID) (*ent.Stack, error) {
	panic("implement me")
}

func (e *Engine) GetStacks(ctx context.Context) ([]*ent.Stack, error) {
	panic("implement me")
}

func (e *Engine) GetStack(ctx context.Context, id uuid.UUID) (*ent.Stack, error) {
	panic("implement me")
}

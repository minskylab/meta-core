package routes

import (
	"context"

	"github.com/minskylab/meta-core/helpers"

	"github.com/minskylab/meta-core/services/structures"
)

func (S service) StartProcess(ctx context.Context, definition structures.ProcessDefinition) (*structures.Process, error) {
	process, err := helpers.StartNewProcess(ctx, S.DbClient, S.Conf, S.Ec2Client, &definition, false)
	if err != nil {
		return nil, err
	}

	return helpers.ObscureProcess(process)
}

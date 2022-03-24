package routes

import (
	"context"

	"github.com/minskylab/meta-core/helpers"

	"github.com/minskylab/meta-core/services/structures"
)

func (S service) DoneProcess(ctx context.Context, processIdentity structures.ProcessIdentity) (*structures.Process, error) {
	return helpers.StopAndDestroyProcess(ctx, S.DbClient, S.Conf, S.Ec2Client, processIdentity.Id)
}

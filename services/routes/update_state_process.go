package routes

import (
	"context"

	"github.com/minskylab/meta-core/helpers"

	"github.com/minskylab/meta-core/services/structures/processstate"

	"github.com/minskylab/meta-core/database"

	"github.com/minskylab/meta-core/services/structures"
)

func (S service) UpdateStateProcess(ctx context.Context, updater structures.ProcessUpdater) (*structures.Process, error) {
	process, err := database.UpdateProcessState(ctx, S.DbClient, updater.Id, updater.State)
	if err != nil {
		return nil, err
	}
	if updater.State == processstate.EXIT_TIMEOUT || updater.State == processstate.EXIT_SUCCESS {
		_, err := helpers.StopAndDestroyProcess(ctx, S.DbClient, S.Conf, S.Ec2Client, updater.Id)
		if err != nil {
			return nil, err
		}
	}
	return process, nil
}

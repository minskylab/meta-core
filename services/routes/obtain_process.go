package routes

import (
	"context"

	"github.com/minskylab/meta-core/helpers"

	"github.com/minskylab/meta-core/database"

	"github.com/minskylab/meta-core/services/structures"
)

func (S service) ObtainProcess(ctx context.Context, processId string) (*structures.Process, error) {
	process, err := database.ObtainProcessById(ctx, S.DbClient, processId)
	if err != nil {
		return nil, err
	}

	return helpers.ObscureProcess(process)
}

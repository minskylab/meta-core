package server

import (
	"context"

	"github.com/minskylab/meta-core/services/structures"
)

type Service interface {
	//kun:op POST /stack/deploy
	//kun:body definition
	//kun:success statusCode=200 body=stack
	PlaceStack(ctx context.Context, definition structures.StackDefinition) (*structures.Stack, error)

	//kun:op POST /stack/purge
	//kun:body s
	//kun:success statusCode=200 body=stack
	PurgeStack(ctx context.Context, s structures.StackIdentity) (*structures.Stack, error)

	//kun:op GET /process/{processId}
	//kun:success statusCode=200 body=process
	ObtainProcess(ctx context.Context, processId string) (*structures.Process, error)

	//kun:op POST /process/start
	//kun:body definition
	//kun:success statusCode=200 body=process
	StartProcess(ctx context.Context, definition structures.ProcessDefinition) (*structures.Process, error)

	//kun:op POST /process/update
	//kun:body updater
	//kun:success statusCode=200 body=process
	UpdateStateProcess(ctx context.Context, updater structures.ProcessUpdater) (*structures.Process, error)

	//kun:op POST /process/done
	//kun:body processIdentity
	//kun:success statusCode=200 body=process
	DoneProcess(ctx context.Context, processIdentity structures.ProcessIdentity) (*structures.Process, error)
}

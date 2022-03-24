package routes

import (
	"context"
	"fmt"

	"github.com/minskylab/meta-core/database"

	log "github.com/sirupsen/logrus"

	"github.com/minskylab/meta-core/helpers"

	"github.com/minskylab/meta-core/services/structures"
)

func (S service) PlaceStack(ctx context.Context, definition structures.StackDefinition) (*structures.Stack, error) {
	stack, err := helpers.DeployResourceStack(S.Conf, S.Ec2Client, definition)
	if err != nil {
		return nil, err
	}

	log.Infoln("Connect to your instance with:")

	cmd := fmt.Sprintf("ssh -i \"%s\" ubuntu@%s", *stack.KeyPair.FilePath, stack.Instance.PublicIp)
	log.Infoln(cmd)

	return database.SaveStack(S.DbClient, ctx, stack)
}

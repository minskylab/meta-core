package routes

import (
	"context"
	"fmt"

	"github.com/minskylab/meta-core/helpers"

	"github.com/minskylab/meta-core/database"

	"github.com/minskylab/meta-core/services/structures"
)

func (S service) PurgeStack(ctx context.Context, s structures.StackIdentity) (*structures.Stack, error) {
	stack, err := database.ObtainStack(S.DbClient, ctx, s.Id)
	if err != nil {
		return nil, err
	}
	if stack == nil {
		return nil, fmt.Errorf("error: stack id (%s) not found", s.Id)
	}

	err = helpers.TerminateInstance(S.Conf, S.Ec2Client, stack.Instance.Id)
	if err != nil {
		return nil, err
	}
	err = helpers.DeleteSecurityGroup(S.Conf, S.Ec2Client, stack.SecurityGroup.Id)
	if err != nil {
		return nil, err
	}
	err = helpers.DeleteKeyPair(S.Conf, S.Ec2Client, stack.KeyPair.Name, stack.KeyPair.Id)

	return stack, nil
}

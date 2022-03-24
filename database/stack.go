package database

import (
	"context"

	"github.com/minskylab/meta-core/ent"
	"github.com/minskylab/meta-core/services/structures"
)

func ObtainStack(dbClient *ent.Client, ctx context.Context, id string) (*structures.Stack, error) {
	stackEntity, err := dbClient.Stack.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return &structures.Stack{
		Id: stackEntity.ID,
		Instance: structures.EC2Instance{
			Id:        stackEntity.Instance,
			VpcId:     stackEntity.VpcID,
			PublicIp:  stackEntity.PublicIP,
			PublicDns: &stackEntity.PublicDNS,
			Username:  stackEntity.Username,
		},
		SecurityGroup: structures.SecurityGroup{
			Id: stackEntity.SecurityGroup,
		},
		KeyPair: structures.KeyPair{
			Id:         stackEntity.KeyPair,
			Name:       stackEntity.Name,
			PrivateKey: stackEntity.PrivateKey,
			FilePath:   &stackEntity.Filepath,
		},
	}, nil
}

func SaveStack(dbClient *ent.Client, ctx context.Context, stack *structures.Stack) (*structures.Stack, error) {
	stackCreator := dbClient.Stack.Create().
		SetID(stack.Id).
		SetInstance(stack.Instance.Id).
		SetVpcID(stack.Instance.VpcId).
		SetPublicIP(stack.Instance.PublicIp).
		SetUsername(stack.Instance.Username).
		SetSecurityGroup(stack.SecurityGroup.Id).
		SetKeyPair(stack.KeyPair.Id).
		SetName(stack.KeyPair.Name).
		SetPrivateKey(stack.KeyPair.PrivateKey)
	if stack.Instance.PublicDns != nil {
		stackCreator = stackCreator.SetPublicDNS(*stack.Instance.PublicDns)
	}
	if stack.KeyPair.FilePath != nil {
		stackCreator = stackCreator.SetFilepath(*stack.KeyPair.FilePath)
	}
	_, err := stackCreator.Save(ctx)
	if err != nil {
		return nil, err
	}
	return stack, nil
}

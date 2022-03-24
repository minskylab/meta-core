package database

import (
	"context"

	"github.com/minskylab/meta-core/ent/task"

	"github.com/minskylab/meta-core/ent/credential"

	process2 "github.com/minskylab/meta-core/ent/process"

	"github.com/minskylab/meta-core/services/structures/processstate"

	"github.com/minskylab/meta-core/ent"
	"github.com/minskylab/meta-core/services/structures"
)

// TODO: Implement ALL

func SaveProcess(ctx context.Context, dbClient *ent.Client, process *structures.Process) (*structures.Process, error) {
	processEntityCreator := dbClient.Process.
		Create().
		SetID(process.Id).
		SetToken(process.Token).
		SetInstanceType(process.Definition.StackDefinition.InstanceType).
		SetAmiID(process.Definition.StackDefinition.AmiId).
		SetResourcePrefix(process.Definition.StackDefinition.ResourcePrefix).
		SetExpiration(process.Definition.StackDefinition.Expiration).
		SetState(process2.State(process.State))

	if process.Definition.Name != nil {
		processEntityCreator = processEntityCreator.SetName(*process.Definition.Name)
	}

	if process.Definition.Timeout != nil {
		processEntityCreator = processEntityCreator.SetTimeout(*process.Definition.Timeout)
	}

	if process.Stack != nil {
		stackEntity, err := dbClient.Stack.Get(ctx, process.Stack.Id)
		if err != nil {
			stackCreator := dbClient.Stack.Create().
				SetID(process.Stack.Id).
				SetInstance(process.Stack.Instance.Id).
				SetVpcID(process.Stack.Instance.VpcId).
				SetPublicIP(process.Stack.Instance.PublicIp).
				SetUsername(process.Stack.Instance.Username).
				SetSecurityGroup(process.Stack.SecurityGroup.Id).
				SetKeyPair(process.Stack.KeyPair.Id).
				SetName(process.Stack.KeyPair.Name).
				SetPrivateKey(process.Stack.KeyPair.PrivateKey)
			if process.Stack.Instance.PublicDns != nil {
				stackCreator = stackCreator.SetPublicDNS(*process.Stack.Instance.PublicDns)
			}
			if process.Stack.KeyPair.FilePath != nil {
				stackCreator = stackCreator.SetFilepath(*process.Stack.KeyPair.FilePath)
			}
			stackEntity, err = stackCreator.Save(ctx)
			if err != nil {
				return nil, err
			}
		}
		processEntityCreator = processEntityCreator.SetStack(stackEntity)
	}

	for _, credentialValue := range process.Definition.Credentials {
		credentialEntity, err := dbClient.Credential.Query().Where(credential.Username(credentialValue.Username), credential.Registry(credentialValue.Registry)).First(ctx)
		if err != nil {
			credentialEntity, err = dbClient.Credential.Create().
				SetRegistry(credentialValue.Registry).
				SetUsername(credentialValue.Username).
				SetPassword(credentialValue.Password).
				Save(ctx)
			if err != nil {
				return nil, err
			}
		}
		processEntityCreator = processEntityCreator.AddCredentials(credentialEntity)
	}

	for _, taskValue := range process.Definition.Tasks {
		taskEntity, err := dbClient.Task.Query().Where(task.Image(taskValue.Image)).First(ctx)
		if err != nil {
			taskCreator := dbClient.Task.Create().
				SetImage(taskValue.Image).
				SetEnvironment(taskValue.Environment).
				SetPorts(taskValue.Ports).
				SetVolumes(taskValue.Volumes).
				SetDetached(taskValue.Detached).
				SetSecurityOpt(taskValue.SecurityOpt).
				SetCapAdd(taskValue.CapAdd)
			if taskValue.Name != nil {
				taskCreator = taskCreator.SetName(*taskValue.Name)
			}
			if taskValue.Cmd != nil {
				taskCreator = taskCreator.SetCmd(*taskValue.Cmd)
			}
			if taskValue.Restart != nil {
				taskCreator = taskCreator.SetRestart(*taskValue.Restart)
			}
			if taskValue.Timeout != nil {
				taskCreator = taskCreator.SetTimeout(*taskValue.Timeout)
			}
			taskEntity, err = taskCreator.Save(ctx)
			if err != nil {
				return nil, err
			}
		}
		processEntityCreator = processEntityCreator.AddTasks(taskEntity)
	}

	_, err := processEntityCreator.Save(ctx)
	if err != nil {
		return nil, err
	}

	return process, nil
}

func ObtainProcessById(ctx context.Context, dbClient *ent.Client, id string) (*structures.Process, error) {
	processEntity, err := dbClient.Process.Query().WithStack().WithTasks().WithCredentials().Where(process2.ID(id)).First(ctx)
	if err != nil {
		return nil, err
	}

	var stack *structures.Stack
	var tasks []structures.TaskDescription
	var credentials []structures.DockerCredential

	for _, taskEntity := range processEntity.Edges.Tasks {
		tasks = append(tasks, structures.TaskDescription{
			Image:       taskEntity.Image,
			Timeout:     &taskEntity.Timeout,
			Name:        &taskEntity.Name,
			Cmd:         &taskEntity.Cmd,
			Detached:    taskEntity.Detached,
			Environment: taskEntity.Environment,
			Volumes:     taskEntity.Volumes,
			Ports:       taskEntity.Ports,
			Restart:     &taskEntity.Restart,
			SecurityOpt: taskEntity.SecurityOpt,
			CapAdd:      taskEntity.CapAdd,
		})
	}

	for _, taskEntity := range processEntity.Edges.Credentials {
		credentials = append(credentials, structures.DockerCredential{
			Username: taskEntity.Username,
			Password: taskEntity.Password,
			Registry: taskEntity.Registry,
		})
	}

	if processEntity.Edges.Stack != nil {
		stack = &structures.Stack{
			Id: processEntity.Edges.Stack.ID,
			Instance: structures.EC2Instance{
				Id:        processEntity.Edges.Stack.Instance,
				VpcId:     processEntity.Edges.Stack.VpcID,
				PublicIp:  processEntity.Edges.Stack.PublicIP,
				PublicDns: &processEntity.Edges.Stack.PublicDNS,
				Username:  processEntity.Edges.Stack.Username,
			},
			SecurityGroup: structures.SecurityGroup{
				Id: processEntity.Edges.Stack.SecurityGroup,
			},
			KeyPair: structures.KeyPair{
				Id:         processEntity.Edges.Stack.KeyPair,
				Name:       processEntity.Edges.Stack.Name,
				PrivateKey: processEntity.Edges.Stack.PrivateKey,
				FilePath:   &processEntity.Edges.Stack.Filepath,
			},
		}
	}

	return &structures.Process{
		Id:    processEntity.ID,
		Token: processEntity.Token,
		State: processstate.ProcessState(processEntity.State),
		Definition: structures.ProcessDefinition{
			Tasks:       tasks,
			Credentials: credentials,
			StackDefinition: structures.StackDefinition{
				InstanceType:   processEntity.InstanceType,
				AmiId:          processEntity.AmiID,
				ResourcePrefix: processEntity.ResourcePrefix,
				Expiration:     processEntity.Expiration,
			},
			Name:    &processEntity.Name,
			Timeout: &processEntity.Timeout,
		},
		Stack: stack,
	}, nil
}

func UpdateProcessState(ctx context.Context, dbClient *ent.Client, id string, newState processstate.ProcessState) (*structures.Process, error) {
	err := dbClient.Process.UpdateOneID(id).SetState(process2.State(newState)).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return ObtainProcessById(ctx, dbClient, id)
}

func AddStackToProcess(ctx context.Context, dbClient *ent.Client, id string, stack *structures.Stack) (*structures.Process, error) {
	stackEntity, err := dbClient.Stack.Get(ctx, stack.Id)
	if err != nil {
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
		stackEntity, err = stackCreator.Save(ctx)
		if err != nil {
			return nil, err
		}
	}
	err = dbClient.Process.UpdateOneID(id).SetStack(stackEntity).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return ObtainProcessById(ctx, dbClient, id)
}

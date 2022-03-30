package helpers

import (
	"context"
	"encoding/base64"
	"fmt"
	"math/rand"
	"time"

	"github.com/minskylab/meta-core/services/structures/callbackauthmethod"

	"github.com/minskylab/meta-core/database"

	"github.com/minskylab/meta-core/services/structures/processstate"

	log "github.com/sirupsen/logrus"

	"github.com/oklog/ulid/v2"

	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/minskylab/meta-core/config"
	"github.com/minskylab/meta-core/ent"
	"github.com/minskylab/meta-core/services/structures"
)

func NewProcess(ctx context.Context, dbClient *ent.Client, definition *structures.ProcessDefinition) (*structures.Process, error) {
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	id := ulid.MustNew(ulid.Timestamp(t), entropy).String()

	tokenBytes := make([]byte, 48)
	_, err := rand.Read(tokenBytes)
	if err != nil {
		return nil, err
	}
	token := base64.StdEncoding.EncodeToString(tokenBytes)

	log.Debugf("-->  PID: %s", id)

	process := &structures.Process{
		Id:         id,
		Token:      token,
		State:      processstate.PENDING,
		Definition: *definition,
		UpdatedAt:  time.Time{},
	}

	return database.SaveProcess(ctx, dbClient, process)
}

func StartNewProcess(ctx context.Context, dbClient *ent.Client, conf *config.Config, ec2Client *ec2.EC2, definition *structures.ProcessDefinition, sync bool) (*structures.Process, error) {
	proto, err := NewProcess(ctx, dbClient, definition)
	if err != nil {
		return nil, err
	}

	process, err := database.UpdateProcessState(ctx, dbClient, proto.Id, processstate.PROVISIONING)
	if err != nil {
		return nil, err
	}

	stack, err := DeployResourceStack(conf, ec2Client, definition.StackDefinition)
	if err != nil {
		return nil, err
	}

	stack, err = database.SaveStack(dbClient, ctx, stack)
	if err != nil {
		return nil, err
	}

	process, err = database.UpdateProcessState(ctx, dbClient, proto.Id, processstate.DEPLOYING)
	if err != nil {
		err2 := PurgeStack(ctx, dbClient, conf, ec2Client, stack.Id)
		if err2 != nil {
			log.Errorln(err2)
		}
		return nil, err
	}

	process, err = database.AddStackToProcess(ctx, dbClient, proto.Id, stack)
	if err != nil {
		err2 := PurgeStack(ctx, dbClient, conf, ec2Client, stack.Id)
		if err2 != nil {
			log.Errorln(err2)
		}
		return nil, err
	}

	client, err := ConnectRemote(conf, &stack.Instance, &stack.KeyPair, "ubuntu")
	if err != nil {
		err2 := PurgeStack(ctx, dbClient, conf, ec2Client, stack.Id)
		if err2 != nil {
			log.Errorln(err2)
		}
		return nil, err
	}

	time.Sleep(30 * time.Second)

	timeout := 360

	if process.Definition.Timeout != nil {
		timeout = *process.Definition.Timeout
	}

	var tasks []structures.DockerTask

	tasks = append(tasks, config.DozzleTask.CompleteDefaults().ToDockerTask(), config.NetdataTask.CompleteDefaults().ToDockerTask())

	for _, task := range process.Definition.Tasks {
		tasks = append(tasks, task.CompleteDefaults().ToDockerTask())
	}

	workerInputContext := structures.InputContext{
		Id:                   process.Id,
		UrlCallback:          conf.UrlCallback,
		CallbackAuthMethod:   callbackauthmethod.BASIC_AUTH,
		CallbackAuthUsername: &conf.BasicAuthUsername,
		CallbackAuthPassword: &conf.BasicAuthPassword,
		Timeout:              timeout,
		Interval:             conf.DefaultHeartBeatInterval,
		Tasks:                tasks,
		Credentials:          process.Definition.Credentials,
	}

	workerB64InputContext, err := InputContextToB64(workerInputContext)
	if err != nil {
		err2 := PurgeStack(ctx, dbClient, conf, ec2Client, stack.Id)
		if err2 != nil {
			log.Errorln(err2)
		}
		return nil, err
	}

	workerImageName := "jmacazana/worker:latest"
	rCmd := fmt.Sprintf("\nsudo docker run -e INPUT_DATA_B64=%s -v /var/run/docker.sock:/var/run/docker.sock -d %s python -u main.py\n", workerB64InputContext, workerImageName)

	ExecuteRemote(conf, client, rCmd)

	process, err = database.UpdateProcessState(ctx, dbClient, process.Id, processstate.RUNNING)
	if err != nil {
		err2 := PurgeStack(ctx, dbClient, conf, ec2Client, stack.Id)
		if err2 != nil {
			log.Errorln(err2)
		}
		err2 = client.Close()
		if err2 != nil {
			log.Errorln(err2)
		}
		return nil, err
	}
	err = client.Close()
	if err != nil {
		log.Errorln(err)
	}

	if sync {
		err2 := PurgeStack(ctx, dbClient, conf, ec2Client, stack.Id)
		if err2 != nil {
			log.Errorln(err2)
		}
	}
	return process, nil
}

func PurgeStack(ctx context.Context, dbClient *ent.Client, config *config.Config, ec2Client *ec2.EC2, id string) error {
	stack, err := database.ObtainStack(dbClient, ctx, id)
	if err != nil {
		return err
	}

	err = DeleteSecurityGroup(config, ec2Client, stack.SecurityGroup.Id)
	if err != nil {
		return err
	}
	err = TerminateInstance(config, ec2Client, stack.Instance.Id)
	if err != nil {
		return err
	}
	err = DeleteKeyPair(config, ec2Client, stack.KeyPair.Name, stack.KeyPair.Id)
	if err != nil {
		return err
	}

	return nil
}

func StopAndDestroyProcess(ctx context.Context, dbClient *ent.Client, config *config.Config, ec2Client *ec2.EC2, id string) (*structures.Process, error) {
	process, err := database.ObtainProcessById(ctx, dbClient, id)
	if err != nil {
		return nil, err
	}

	if process.Stack != nil {
		err := PurgeStack(ctx, dbClient, config, ec2Client, process.Stack.Id)
		if err != nil {
			return nil, err
		}
	}

	return process, nil
}

package helpers

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/aws/aws-sdk-go/service/ec2"

	log "github.com/sirupsen/logrus"

	"github.com/oklog/ulid/v2"

	"github.com/minskylab/meta-core/config"
	"github.com/minskylab/meta-core/services/structures"
)

func DeployResourceStack(conf *config.Config, ec2Client *ec2.EC2, definition structures.StackDefinition) (*structures.Stack, error) {
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	id := ulid.MustNew(ulid.Timestamp(t), entropy).String()
	prefix := fmt.Sprintf("%s_%s", definition.ResourcePrefix, id)
	keyPairName := fmt.Sprintf("%s_key", prefix)
	instanceName := fmt.Sprintf("%s_instance", prefix)
	groupName := fmt.Sprintf("%s_sec_group", prefix)

	if conf.Debug {
		log.Debug("Creating")
		log.Debug(fmt.Sprintf("-->  key_pair_name=%s", keyPairName))
		log.Debug(fmt.Sprintf("-->  group_name=%s", groupName))
		log.Debug(fmt.Sprintf("-->  instance_name=%s", instanceName))
	}

	vpcId, subnetId, err := GetNetworkResources(ec2Client)
	if err != nil {
		log.Errorln(err)
		return nil, err
	}

	keyPair, err := CreateKeyPair(ec2Client, keyPairName, true)
	if err != nil {
		log.Errorln(err)
		return nil, err
	}

	secGroup, err := CreateSecurityGroup(ec2Client, groupName, *vpcId)
	if err != nil {
		_ = DeleteKeyPair(conf, ec2Client, keyPair.Name, keyPair.Id)
		log.Errorln(err)
		return nil, err
	}

	instance, err := CreateEC2Instance(ec2Client, instanceName, definition.AmiId, keyPairName, definition.InstanceType, secGroup.Id, *subnetId, GetDefaultUserdata())
	if err != nil {
		_ = DeleteKeyPair(conf, ec2Client, keyPair.Name, keyPair.Id)
		log.Errorln(err)
		return nil, err
	}

	if conf.Debug {
		log.Debugln("Waiting")
	}

	time.Sleep(50 * time.Second)

	if conf.Debug {
		log.Debugln("Created")
		log.Debugln(fmt.Sprintf("-->  key_pair_id=%s", keyPair.Id))
		log.Debugln(fmt.Sprintf("-->  group_id=%s", secGroup.Id))
		log.Debugln(fmt.Sprintf("-->  instance_id=%s", instance.Id))
		log.Debugln("Done. Resources created and ready to use")
	}

	return &structures.Stack{
		Id:            id,
		Instance:      *instance,
		SecurityGroup: *secGroup,
		KeyPair:       *keyPair,
	}, nil
}

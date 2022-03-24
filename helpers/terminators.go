package helpers

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/minskylab/meta-core/config"
	log "github.com/sirupsen/logrus"
)

func TerminateInstance(conf *config.Config, ec2Client *ec2.EC2, instanceId string) error {
	response, err := ec2Client.TerminateInstances(&ec2.TerminateInstancesInput{
		InstanceIds: []*string{&instanceId},
	})
	if err != nil {
		return err
	}

	if conf.Debug {
		log.Debugln(response)
	}

	return nil
}

func DeleteSecurityGroup(conf *config.Config, ec2Client *ec2.EC2, groupId string) error {
	response, err := ec2Client.DeleteSecurityGroup(&ec2.DeleteSecurityGroupInput{
		GroupId: aws.String(groupId),
	})
	if err != nil {
		return err
	}

	if conf.Debug {
		log.Debugln(response)
	}

	return nil
}

func DeleteKeyPair(conf *config.Config, ec2Client *ec2.EC2, keyPairName, keyPairId string) error {
	response, err := ec2Client.DeleteKeyPair(&ec2.DeleteKeyPairInput{
		DryRun:  aws.Bool(false),
		KeyName: &keyPairName,
	})
	if err != nil {
		return err
	}

	if conf.Debug {
		log.Debugln(response)
	}

	return nil
}

package helpers

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/aws/aws-sdk-go/aws"

	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/minskylab/meta-core/services/structures"
)

func CreateKeyPair(ec2Client *ec2.EC2, name string, saveAsFile bool) (*structures.KeyPair, error) {
	var filePath *string = nil
	keyPair, err := ec2Client.CreateKeyPair(&ec2.CreateKeyPairInput{KeyName: &name})
	if err != nil {
		return nil, err
	}

	privateKey := keyPair.KeyMaterial
	keyId := keyPair.KeyPairId

	if saveAsFile {
		fileName := fmt.Sprintf("/tmp/%s.pem", name)
		err := ioutil.WriteFile(fileName, []byte(*privateKey), 0o666)
		if err != nil {
			return nil, err
		}
		filePath = &fileName
	}

	return &structures.KeyPair{
		Id:         *keyId,
		Name:       name,
		PrivateKey: *privateKey,
		FilePath:   filePath,
	}, nil
}

func GetDefaultUserdata() string {
	return fmt.Sprintf("#!/usr/bin/env bash\nsudo apt update && sudo apt install -y docker.io")
}

func CreateEC2Instance(ec2Client *ec2.EC2, name, imageId, keyPairName, instanceType, securityGroupId, subnetId, userdata string) (*structures.EC2Instance, error) {
	instances, err := ec2Client.RunInstances(&ec2.RunInstancesInput{
		ImageId:          &imageId,
		InstanceType:     &instanceType,
		KeyName:          &keyPairName,
		SecurityGroupIds: []*string{&securityGroupId},
		SubnetId:         &subnetId,
		MaxCount:         aws.Int64(1),
		MinCount:         aws.Int64(1),
		TagSpecifications: []*ec2.TagSpecification{
			{
				ResourceType: aws.String("instance"),
				Tags: []*ec2.Tag{
					{
						Key:   aws.String("Name"),
						Value: &name,
					},
				},
			},
		},
		UserData: aws.String(base64.StdEncoding.EncodeToString([]byte(userdata))),
	})
	if err != nil {
		return nil, err
	}

	netInterface := instances.Instances[0].NetworkInterfaces[0]
	vpcId := netInterface.VpcId

	instanceId := instances.Instances[0].InstanceId

	time.Sleep(3 * time.Second)

	out, err := ec2Client.DescribeInstances(&ec2.DescribeInstancesInput{
		DryRun:      nil,
		Filters:     nil,
		InstanceIds: []*string{instanceId},
		MaxResults:  nil,
		NextToken:   nil,
	})
	if err != nil {
		return nil, err
	}

	re := out.Reservations

	if len(re) < 1 {
		return nil, errors.New("reservations not found")
	}

	if len(re[0].Instances) < 1 {
		return nil, errors.New("instances not found in reservations")
	}

	rInstance := re[0].Instances[0]

	if rInstance.PublicIpAddress == nil || rInstance.PublicDnsName == nil {
		return nil, errors.New("ip or dns not found in current reservation")
	}

	publicIp := rInstance.PublicIpAddress
	publicDns := rInstance.PublicDnsName

	return &structures.EC2Instance{
		Id:        *instanceId,
		VpcId:     *vpcId,
		PublicIp:  *publicIp,
		PublicDns: publicDns,
	}, nil
}

/*
	TODO: En el Input de Create Security Group falta un Tag Specification
		TagSpecifications = [
			{
				"ResourceType": "security-group",
				"Tags": [{"Key": "Name", "Value": name}],
			},
		],
*/
func CreateSecurityGroup(ec2Client *ec2.EC2, name, vpcId string) (*structures.SecurityGroup, error) {
	response, err := ec2Client.CreateSecurityGroup(&ec2.CreateSecurityGroupInput{
		Description: aws.String("ChagasWorker"),
		GroupName:   &name,
		VpcId:       &vpcId,
	})
	if err != nil {
		return nil, err
	}

	securityGroupId := response.GroupId

	_, err = ec2Client.AuthorizeSecurityGroupIngress(&ec2.AuthorizeSecurityGroupIngressInput{
		GroupId: securityGroupId,
		IpPermissions: []*ec2.IpPermission{
			{
				FromPort:   aws.Int64(80),
				IpProtocol: aws.String("tcp"),
				ToPort:     aws.Int64(80),
				IpRanges: []*ec2.IpRange{
					{
						CidrIp: aws.String("0.0.0.0/0"),
					},
				},
			},
			{
				FromPort:   aws.Int64(8080),
				IpProtocol: aws.String("tcp"),
				ToPort:     aws.Int64(8080),
				IpRanges: []*ec2.IpRange{
					{
						CidrIp: aws.String("0.0.0.0/0"),
					},
				},
			},
			{
				FromPort:   aws.Int64(22),
				IpProtocol: aws.String("tcp"),
				ToPort:     aws.Int64(22),
				IpRanges: []*ec2.IpRange{
					{
						CidrIp: aws.String("0.0.0.0/0"),
					},
				},
			},
		},
	})

	return &structures.SecurityGroup{Id: *securityGroupId}, nil
}

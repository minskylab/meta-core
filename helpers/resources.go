package helpers

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"

	"github.com/aws/aws-sdk-go/service/ec2"
)

func GetNetworkResources(ec2Client *ec2.EC2) (*string, *string, error) {
	var subnet *ec2.Subnet

	vpcs, err := ec2Client.DescribeVpcs(nil)
	if err != nil {
		return nil, nil, err
	}
	if len(vpcs.Vpcs) == 0 {
		return nil, nil, errors.New("no VPCs found to associate security group with")
	}

	vpc := vpcs.Vpcs[0]

	subnets, err := ec2Client.DescribeSubnets(&ec2.DescribeSubnetsInput{
		Filters: []*ec2.Filter{
			{
				Name:   aws.String("vpc-id"),
				Values: []*string{vpc.VpcId},
			},
		},
	})
	if err != nil {
		return nil, nil, err
	}
	if len(subnets.Subnets) == 0 {
		return nil, nil, errors.New("no Subnets found to associate security group with")
	}

	for _, snet := range subnets.Subnets {
		if *snet.AvailabilityZone != "us-east-1e" {
			subnet = snet
			break
		}
	}

	return vpc.VpcId, subnet.SubnetId, nil
}

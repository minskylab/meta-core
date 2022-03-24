package routes

import (
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/minskylab/meta-core/config"
	"github.com/minskylab/meta-core/ent"
	"github.com/minskylab/meta-core/server"

	"github.com/aws/aws-sdk-go/aws/session"
)

type service struct {
	Conf       *config.Config
	AwsSession *session.Session
	Ec2Client  *ec2.EC2
	DbClient   *ent.Client
}

func NewService(conf *config.Config, awsSession *session.Session, dbClient *ent.Client) server.Service {
	return service{
		Conf:       conf,
		AwsSession: awsSession,
		Ec2Client:  ec2.New(awsSession),
		DbClient:   dbClient,
	}
}

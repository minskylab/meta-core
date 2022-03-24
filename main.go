package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/minskylab/meta-core/ent"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"

	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/minskylab/meta-core/helpers"

	"github.com/RussellLuo/kun/pkg/httpcodec"
	"github.com/RussellLuo/kun/pkg/httpoption2"
	"github.com/minskylab/meta-core/config"
	"github.com/minskylab/meta-core/server"
	"github.com/minskylab/meta-core/services/routes"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	var err error

	var conf *config.Config
	{
		conf = config.NewConfig()
		if conf.Debug {
			log.SetLevel(log.DebugLevel)
		}
	}

	var awsSession *session.Session
	{
		awsSession, err = session.NewSession(&aws.Config{
			Region:      &conf.AwsDefaultRegion,
			Credentials: credentials.NewStaticCredentials(conf.AwsAccessKey, conf.AwsSecretKey, ""),
		})
		if err != nil {
			log.Fatalln("Error at creating an AWS Session")
		}
	}

	var dbClient *ent.Client
	{
		dbClient, err = ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
		if err != nil {
			log.Fatalf("Failed opening connection to sqlite: %v\n", err)
		}
		defer dbClient.Close()

		if err := dbClient.Schema.Create(context.Background()); err != nil {
			log.Fatalf("failed creating schema resources: %v", err)
		}
	}

	var service server.Service
	{
		service = routes.NewService(conf, awsSession, dbClient)
	}

	errorCollector := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errorCollector <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		helpers.PrintLogo(conf.Host + ":" + conf.Port)
		s := server.NewHTTPRouter(service, httpcodec.NewDefaultCodecs(nil), httpoption.RequestValidators())
		serv := cors.AllowAll().Handler(s)
		errorCollector <- http.ListenAndServe(conf.Host+":"+conf.Port, serv)
	}()

	log.Errorln("exit: ", <-errorCollector)
}

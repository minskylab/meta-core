package helpers

import (
	ssh "github.com/helloyi/go-sshclient"
	"github.com/minskylab/meta-core/config"
	"github.com/minskylab/meta-core/services/structures"
	log "github.com/sirupsen/logrus"
)

func ConnectRemote(conf *config.Config, instance *structures.EC2Instance, keyPair *structures.KeyPair, username string) (*ssh.Client, error) {
	host := instance.PublicIp + ":22"

	if conf.Debug {
		log.Debugf("ssh -i %s %s@%s", *keyPair.FilePath, username, host)
	}

	client, err := ssh.DialWithKey(host, username, *keyPair.FilePath)
	if err != nil {
		return nil, err
	}

	log.Debugln("MEOW")

	return client, nil
}

func ExecuteRemote(conf *config.Config, client *ssh.Client, cmd string) string {
	out, err := client.Cmd(cmd).SmartOutput()
	if err != nil {
		// the 'out' is stderr output
		if conf.Debug {
			log.Debugln("error: ", out)
		}
		return string(out)
	}

	if conf.Debug {
		log.Debugln("output: ", out)
	}
	// the 'out' is stdout output
	return string(out)
}

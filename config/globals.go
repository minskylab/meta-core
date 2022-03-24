package config

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/minskylab/meta-core/services/structures"
)

var DozzleTask = &structures.TaskDescription{
	Image:    "amir20/dozzle",
	Name:     aws.String("amir20_dozzle"),
	Detached: true,
	Volumes:  []string{"/var/run/docker.sock:/var/run/docker.sock"},
	Ports:    []string{"8080:8080"},
}

var NetdataTask = &structures.TaskDescription{
	Image:    "netdata/netdata",
	Name:     aws.String("netdata_netdata"),
	Detached: true,
	Volumes: []string{
		"netdataconfig:/etc/netdata",
		"netdatalib:/var/lib/netdata",
		"netdatacache:/var/cache/netdata",
		"/etc/passwd:/host/etc/passwd:ro",
		"/etc/group:/host/etc/group:ro",
		"/proc:/host/proc:ro",
		"/sys:/host/sys:ro",
		"/etc/os-release:/host/etc/os-release:ro",
	},
	Ports:       []string{"80:19999"},
	Restart:     aws.String("unless-stopped"),
	SecurityOpt: []string{"apparmor=unconfined"},
	CapAdd:      []string{"SYS_ADMIN"},
}

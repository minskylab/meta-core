package structures

import (
	"time"

	"github.com/minskylab/meta-core/services/structures/processstate"
)

type ProcessDefinition struct {
	Tasks           []TaskDescription  `json:"tasks"`
	Credentials     []DockerCredential `json:"credentials"`
	StackDefinition StackDefinition    `json:"stack_definition"`
	Name            *string            `json:"name"`
	Timeout         *int               `json:"timeout"`
}

type ProcessUpdater struct {
	Id    string                    `json:"id"`
	State processstate.ProcessState `json:"state"`
}

type ProcessIdentity struct {
	Id    string `json:"id"`
	Token string `json:"token"`
}

type Process struct {
	Id         string                    `json:"id"`
	Token      string                    `json:"token"`
	State      processstate.ProcessState `json:"state"`
	Definition ProcessDefinition         `json:"definition"`
	UpdatedAt  time.Time                 `json:"updated_at"`
	Logs       *map[time.Time]string     `json:"logs"`
	Stack      *Stack                    `json:"stack"`
}

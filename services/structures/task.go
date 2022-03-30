package structures

type TaskDescription struct {
	Image       string                 `json:"image"`
	Timeout     *int                   `json:"timeout"`
	Name        *string                `json:"name"`
	Cmd         *string                `json:"cmd"`
	Detached    bool                   `json:"detached"`
	Environment map[string]interface{} `json:"environment"`
	Volumes     []string               `json:"volumes"`
	Ports       []string               `json:"ports"`
	Restart     *string                `json:"restart"`
	SecurityOpt []string               `json:"security_opt"`
	CapAdd      []string               `json:"cap_add"`
}

func (task TaskDescription) ToDockerTask() DockerTask {
	name := ""
	if task.Name != nil {
		name = *task.Name
	}
	return DockerTask{
		Timeout:     task.Timeout,
		Name:        name,
		Image:       task.Image,
		Cmd:         task.Cmd,
		Detached:    task.Detached,
		Environment: task.Environment,
		Volumes:     task.Volumes,
		Ports:       task.Ports,
		Restart:     task.Restart,
		SecurityOpt: task.SecurityOpt,
		CapAdd:      task.CapAdd,
	}
}

func (task TaskDescription) CompleteDefaults() TaskDescription {
	if task.Environment == nil {
		task.Environment = map[string]interface{}{}
	}
	if task.Volumes == nil {
		task.Volumes = []string{}
	}
	if task.Ports == nil {
		task.Ports = []string{}
	}
	return task
}

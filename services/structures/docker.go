package structures

import "github.com/minskylab/meta-core/services/structures/callbackauthmethod"

type DockerTask struct {
	Timeout     *int                   `json:"timeout"`
	Name        string                 `json:"name"`
	Image       string                 `json:"image"`
	Cmd         *string                `json:"cmd"`
	Detached    bool                   `json:"detached"`
	Environment map[string]interface{} `json:"environment"`
	Volumes     []string               `json:"volumes"`
	Ports       []string               `json:"ports"`
	Restart     *string                `json:"restart"`
	SecurityOpt []string               `json:"security_opt"`
	CapAdd      []string               `json:"cap_add"`
}

type DockerCredential struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Registry string `json:"registry"`
}

type InputContext struct {
	Id                   string                                `json:"id"`
	UrlCallback          string                                `json:"url_callback"`
	CallbackAuthMethod   callbackauthmethod.CallbackAuthMethod `json:"callback_auth_method"`
	CallbackAuthUsername *string                               `json:"callback_auth_username"`
	CallbackAuthPassword *string                               `json:"callback_auth_password"`
	Timeout              int                                   `json:"timeout"`
	Interval             float64                               `json:"interval"`
	Tasks                []DockerTask                          `json:"tasks"`
	Credentials          []DockerCredential                    `json:"credentials"`
}

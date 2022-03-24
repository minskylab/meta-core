package structures

type StackDefinition struct {
	InstanceType   string `json:"instance_type"`
	AmiId          string `json:"ami_id"`
	ResourcePrefix string `json:"resource_prefix"`
	Expiration     int    `json:"expiration"`
}

type StackIdentity struct {
	Id string `json:"id"`
}

type KeyPair struct {
	Id         string  `json:"id"`
	Name       string  `json:"name"`
	PrivateKey string  `json:"private_key"`
	FilePath   *string `json:"filepath"`
}

type EC2Instance struct {
	Id        string  `json:"id"`
	VpcId     string  `json:"vpc_id"`
	PublicIp  string  `json:"public_ip"`
	PublicDns *string `json:"public_dns"`
	Username  string  `json:"username"`
}

type SecurityGroup struct {
	Id string `json:"id"`
}

type Stack struct {
	Id            string        `json:"id"`
	Instance      EC2Instance   `json:"instance"`
	SecurityGroup SecurityGroup `json:"security_group"`
	KeyPair       KeyPair       `json:"key_pair"`
}

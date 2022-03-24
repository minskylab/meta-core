package config

import (
	"os"
	"strconv"
)

type Config struct {
	Host                     string
	Port                     string
	DBHost                   string
	DBPort                   string
	DBUsername               string
	DBPassword               string
	DbName                   string
	UrlCallback              string
	BasicAuthUsername        string
	BasicAuthPassword        string
	AwsAccessKey             string
	AwsSecretKey             string
	AwsDefaultRegion         string
	DefaultInstanceType      string
	DefaultAmiID             string
	DefaultPrefixResource    string
	Debug                    bool
	DefaultHeartBeatInterval float64
}

func NewConfig() *Config {
	debug, err := strconv.ParseBool(envOrDefault("DEBUG", "true"))
	if err != nil {
		debug = true
	}
	defaultHeartBeatInterval, err := strconv.ParseFloat(envOrDefault("DEFAULT_HEARTBEAT_INTERVAL", "15.0"), 32)
	if err != nil {
		defaultHeartBeatInterval = 15.0
	}
	return &Config{
		Host:                     envOrDefault("HOSTNAME", "127.0.0.1"),
		Port:                     envOrDefault("PORT", "8080"),
		DBHost:                   envOrDefault("DB_HOST", ""),
		DBPort:                   envOrDefault("DB_PORT", ""),
		DBUsername:               envOrDefault("DB_USERNAME", ""),
		DBPassword:               envOrDefault("DB_PASSWORD", ""),
		DbName:                   envOrDefault("DB_NAME", ""),
		UrlCallback:              envOrDefault("URL_CALLBACK", "https://httpbin.org/post"),
		BasicAuthUsername:        envOrDefault("BASIC_AUTH_USERNAME", ""),
		BasicAuthPassword:        envOrDefault("BASIC_AUTH_PASSWORD", ""),
		AwsAccessKey:             envOrDefault("AWS_ACCESS_KEY", ""),
		AwsSecretKey:             envOrDefault("AWS_SECRET_KEY", ""),
		AwsDefaultRegion:         envOrDefault("AWS_DEFAULT_REGION", ""),
		DefaultInstanceType:      envOrDefault("DEFAULT_INSTANCE_TYPE", "t2.small"),
		DefaultAmiID:             envOrDefault("DEFAULT_AMI_ID", "ami-0cc995a23da39b9ba"),
		DefaultPrefixResource:    envOrDefault("DEFAULT_PREFIX_RESOURCE", "amet"),
		Debug:                    debug,
		DefaultHeartBeatInterval: defaultHeartBeatInterval,
	}
}

func envOrDefault(key string, def string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return def
}

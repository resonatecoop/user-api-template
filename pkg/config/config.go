package config

import (
	"fmt"
	"io/ioutil"
	"time"

	yaml "gopkg.in/yaml.v2"
)

// Load loads the configuration file from the given path
func Load(path string) (*Configuration, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error reading config file, %s", err)
	}

	var cfg Configuration
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("unable to decode config into struct, %v", err)
	}

	return &cfg, nil
}

// Configuration holds application configuration data
type Configuration struct {
	Server       Server       `yaml:"server,omitempty"`
	RefreshToken RefreshToken `yaml:"refreshtoken,omitempty"`
	DB           DatabaseEnv  `yaml:"database,omitempty"`
	Access       Access       `yaml:"access,omitempty"`
	App          Application  `yaml:"application,omitempty"`
	OpenAPI      OpenAPI      `yaml:"openapi,omitempty"`
	Storage      Storage      `yaml:"storage,omitempty"`
}

// DatabaseEnv holds dev and test database data
type DatabaseEnv struct {
	Dev  Database `yaml:"dev,omitempty"`
	Test Database `yaml:"test,omitempty"`
}

// Database holds data necessery for database configuration
type Database struct {
	PSN            string `yaml:"psn,omitempty"`
	LogQueries     bool   `yaml:"log_queries,omitempty"`
	TimeoutSeconds int    `yaml:"timeout_seconds,omitempty"`
}

// Server holds data necessery for server configuration
type Server struct {
	CertName            string `yaml:"cert_name,omitempty"`
	ReadTimeoutSeconds  int    `yaml:"read_timeout_seconds,omitempty"`
	WriteTimeoutSeconds int    `yaml:"write_timeout_seconds,omitempty"`
}

type RefreshToken struct {
	Lifetime int `yaml:"lifetime_seconds,omitempty"`
}

// Access holds service access configuration data
type Access struct {
	NoTokenMethods string `yaml:"no_token_methods,omitempty"`
	PublicMethods  string `yaml:"public_methods,omitempty"`
	WriteMethods   string `yaml:"write_methods,omitempty"`
}

// Application represents application specific configuration
type Application struct {
	MinPasswordStrength int `yaml:"min_password_strength,omitempty"`
}

// OpenAPI holds username password for viewing api docs
type OpenAPI struct {
	Username string `yaml:"username,omitempty"`
	Password string `yaml:"password,omitempty"`
}

// Storage holds data necessary for backblaze configuration in track-server-api
type Storage struct {
	AccountId      string        `yaml:"account_id,omitempty"`
	Key            string        `yaml:"key,omitempty"`
	AuthEndpoint   string        `yaml:"auth_endpoint,omitempty"`
	FileEndpoint   string        `yaml:"file_endpoint,omitempty"`
	UploadEndpoint string        `yaml:"upload_endpoint,omitempty"`
	BucketId       string        `yaml:"bucket_id,omitempty"`
	Timeout        time.Duration `yaml:"timeout,omitempty"`
}

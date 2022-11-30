package models

import (
	"github.com/mitchellh/go-homedir"
	"os"
	"path"
)

const AuthAccessToken = "auth.access_token"
const AuthExpiredAt = "auth.expired_at"
const AuthRefreshToken = "auth.refresh_token"
const AuthUserId = "auth.user_id"

// TODO delete auth
var DefaultBuckets = []string{"folders", "contexts", "tasks", "auth", "account", "goals"}

// ToodledoConfigEnvironment ...
type ToodledoConfigEnvironment struct {
	Name    string `mapstructure:"name"`
	Folder  string `mapstructure:"folder"`
	Context string `mapstructure:"context"`
	Goal    string `mapstructure:"goal"`
}

// ToodledoConfig is configuration for toodledo
// mapstructure docs in https://github.com/spf13/viper/issues/258#issuecomment-253730493
type ToodledoConfig struct {
	UserID       string `mapstructure:"user_id"`
	ClientId     string `mapstructure:"client_id" yaml:"client_id"`
	ClientSecret string `mapstructure:"client_secret" yaml:"client_secret"`
	AccessToken  string `mapstructure:"access_token" yaml:"access_token"`
	ExpiredAt    string `mapstructure:"expired_at" yaml:"expired_at"`
	RefreshToken string `mapstructure:"refresh_token" yaml:"refresh_token"`
}

// ToodledoCliConfig is configuration for toodledo cli
type ToodledoCliConfig struct {
	Endpoint       string                                `mapstructure:"endpoint" yaml:"host"`
	Auth           ToodledoConfig                        `mapstructure:"auth" yaml:"auth"`
	Database       ToodledoConfigDatabase                `mapstructure:"database omitempty" yaml:"database omitempty"`
	Environment    map[string]*ToodledoConfigEnvironment `mapstructure:"environments  omitempty" yaml:"environment omitempty"`
	DefaultContext string                                `mapstructure:"default-environment" yaml:"defaultContext omitempty"`
}

func NewToodledoCliConfig() ToodledoCliConfig {
	return ToodledoCliConfig{
		Endpoint:       "https://api.toodledo.com",
		Auth:           ToodledoConfig{},
		Database:       NewDefaultToodledoConfigDatabase(),
		Environment:    map[string]*ToodledoConfigEnvironment{},
		DefaultContext: "",
	}
}

type ToodledoConfigDatabase struct {
	DataFile string   `mapstructure:"data_file omitempty" yaml:"data_file omitempty"`
	Buckets  []string `mapstructure:"-" yaml:"-"`
}

func NewDefaultToodledoConfigDatabase() ToodledoConfigDatabase {
	home, _ := homedir.Dir()
	return ToodledoConfigDatabase{
		DataFile: path.Join(home, ".config", "toodledo", "data", "data.db"),
		Buckets:  DefaultBuckets,
	}
}

func NewToodledoConfigDatabaseFromToodledoCliConfig(config ToodledoCliConfig) ToodledoConfigDatabase {
	return ToodledoConfigDatabase{
		DataFile: config.Database.DataFile,
		Buckets:  config.Database.Buckets,
	}
}

func CleanDatabase(db ToodledoConfigDatabase) error {
	file := db.DataFile
	// remove files
	_ = os.Remove(file)
	return nil
}

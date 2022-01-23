package models

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
	AccessToken  string `mapstructure:"access_token" yaml:"access_token"`
	ClientId     string `mapstructure:"client_id" yaml:"client_id"`
	ClientSecret string `mapstructure:"client_secret" yaml:"client_secret"`
	ExpiredAt    string `mapstructure:"expired_at" yaml:"expired_at"`
	RefreshToken string `mapstructure:"refresh_token" yaml:"refresh_token"`
}

// ToodledoCliConfig ...
type ToodledoCliConfig struct {
	Auth           ToodledoConfig                        `mapstructure:"auth" yaml:"auth"`
	Database       ToodledoConfigDatabase                `mapstructure:"database" yaml:"database"`
	Environment    map[string]*ToodledoConfigEnvironment `mapstructure:"environments"`
	DefaultContext string                                `mapstructure:"default-environment"`
}

var DefaultBuckets = []string{"folders", "contexts", "tasks", "auth"}

// ToodledoConfigDatabase ...
type ToodledoConfigDatabase struct {
	DataFile string   `mapstructure:"data_file" yaml:"data_file"`
	Buckets  []string `mapstructure:"-" yaml:"-"`
}

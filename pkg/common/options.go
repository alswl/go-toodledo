package common

type ToodledoConfigAuth struct {
	AccessToken  string `mapstructure:"access_token" yaml:"access_token"`
	ClientId     string `mapstructure:"client_id" yaml:"client_id"`
	ClientSecret string `mapstructure:"client_secret" yaml:"client_secret"`
	ExpiredAt    string `mapstructure:"expired_at" yaml:"expired_at"`
	RefreshToken string `mapstructure:"refresh_token" yaml:"refresh_token"`
}

type ToodledoConfigDatabase struct {
	DataFile string   `mapstructure:"data_file" yaml:"data_file"`
	Buckets  []string `mapstructure:"buckets" yaml:"buckets"`
}

type ToodledoConfig struct {
	Auth     ToodledoConfigAuth     `mapstructure:"auth" yaml:"auth"`
	Database ToodledoConfigDatabase `mapstructure:"database" yaml:"database"`
}

func NewToodledoConfig(configs Configs) ToodledoConfig {
	// TODO dirty, but it works
	cfg := *configs.Get()
	// TODO fix consts
	cfg.Database.Buckets = []string{
		"folders",
		"contexts",
	}
	return cfg
}

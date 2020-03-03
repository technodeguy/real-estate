package config

import (
	"strings"

	"github.com/technodeguy/real-estate/api/consts"

	"github.com/spf13/viper"
)

type FileWhiteListType []string

type ServerConfig struct {
	Host string
	Port int
}

type DbConfig struct {
	Uri string
}

type RedisConfig struct {
	Uri string
}

type AwsConfig struct {
	BucketName     string `mapstructure:"bucket_name"`
	AccessKeyId    string `mapstructure:"access_key_id"`
	SecretAccesKey string `mapstructure:"secret_access_key"`
}

type FileStoreConfig struct {
	FileWhiteList FileWhiteListType
}

type Config struct {
	Server    ServerConfig
	Db        DbConfig
	Redis     RedisConfig
	Aws       AwsConfig
	FileStore FileStoreConfig
}

func LoadConfig(filename string) (*Config, error) {

	v := viper.New()

	v.SetConfigName(filename)
	v.AddConfigPath(consts.CONFIG_DIR)
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	var err error

	if err = v.ReadInConfig(); err != nil {
		return nil, err
	}

	config := new(Config)

	if err = v.Unmarshal(config); err != nil {
		return nil, err
	}

	return config, nil
}

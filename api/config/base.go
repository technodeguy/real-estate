package config

import (
	"strings"

	"github.com/spf13/viper"
)

type ServerConfig struct {
	Host string
	Port int16
}

type DbConfig struct {
	Uri string
}

type AwsConfig struct {
	BucketName     string `mapstructure:"bucket_name"`
	AccessKeyId    string `mapstructure:"access_key_id"`
	SecretAccesKey string `mapstructure:"secret_access_key"`
}

type FileStoreConfig struct {
	FileWhiteList []string
}

type Config struct {
	Server    ServerConfig
	Db        DbConfig
	Aws       AwsConfig
	FileStore FileStoreConfig
}

func ReadConfig(filename string) (*Config, error) {

	v := viper.New()

	v.SetConfigName(filename)
	v.AddConfigPath("./api/config")
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

package config

import (
	"github.com/spf13/viper"
	"strings"
	"time"
)

//go:generate mockgen -source=$GOFILE -package=mock_config -destination=../../test/mock/config/$GOFILE

type ConfigStore interface {
	Get(key string) interface{}
	GetBool(key string) bool
	GetFloat64(key string) float64
	GetInt(key string) int
	GetIntSlice(key string) []int
	GetString(key string) string
	GetStringMap(key string) map[string]interface{}
	GetStringMapString(key string) map[string]string
	GetStringSlice(key string) []string
	GetTime(key string) time.Time
	GetDuration(key string) time.Duration
	IsSet(key string) bool
	IsBatchMode() bool
}

type DefaultConfigStore struct {
	viper *viper.Viper
}

func (dcs *DefaultConfigStore) Get(key string) interface{} {
	return dcs.viper.Get(key)
}

func (dcs *DefaultConfigStore) GetBool(key string) bool {
	return dcs.viper.GetBool(key)
}

func (dcs *DefaultConfigStore) GetFloat64(key string) float64 {
	return dcs.viper.GetFloat64(key)
}

func (dcs *DefaultConfigStore) GetInt(key string) int {
	return dcs.viper.GetInt(key)
}

func (dcs *DefaultConfigStore) GetIntSlice(key string) []int {
	return dcs.viper.GetIntSlice(key)
}

func (dcs *DefaultConfigStore) GetString(key string) string {
	return dcs.viper.GetString(key)
}

func (dcs *DefaultConfigStore) GetStringMap(key string) map[string]interface{} {
	return dcs.viper.GetStringMap(key)
}

func (dcs *DefaultConfigStore) GetStringMapString(key string) map[string]string {
	return dcs.viper.GetStringMapString(key)
}

func (dcs *DefaultConfigStore) GetStringSlice(key string) []string {
	return dcs.viper.GetStringSlice(key)
}

func (dcs *DefaultConfigStore) GetTime(key string) time.Time {
	return dcs.viper.GetTime(key)
}

func (dcs *DefaultConfigStore) GetDuration(key string) time.Duration {
	return dcs.viper.GetDuration(key)
}

func (dcs *DefaultConfigStore) IsSet(key string) bool {
	return dcs.viper.IsSet(key)
}

func (dcs *DefaultConfigStore) IsBatchMode() bool {
	return dcs.GetString("jobName") > ""
}

func NewConfigManager() (ConfigStore, error) {
	intViper := viper.New()
	intViper.SetConfigName("config")
	intViper.SetConfigType("yaml")
	intViper.AddConfigPath("./configs")
	intViper.AddConfigPath(".")    // optionally look for config in the working directory
	err := intViper.ReadInConfig() // Find and read the config file
	if err != nil {                // Handle errors reading the config file
		return nil, err
	}
	intViper.SetEnvPrefix("chat-jobsity")
	intViper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	intViper.AutomaticEnv()

	return &DefaultConfigStore{viper: intViper}, nil
}

package config

import (
	"io"

	"github.com/spf13/viper"
)

type CollectionConfig struct {
	Key      string `mapstructure:"key"`
	Type     string `mapstructure:"type"`
	Pattern  string `mapstructure:"pattern"`
	Index    int    `mapstructure:"index"`
	Unique   bool   `mapstructure:"unique"`
	Required bool   `mapstructure:"required"`
}

type ApiConfig struct {
	Uri        string   `mapstructure:"uri"`
	Collection string   `mapstructure:"collection"`
	Method     string   `mapstructure:"method"`
	First      bool     `mapstructure:"first"`
	Many       bool     `mapstructure:"many"`
	Searchable []string `mapstructure:"searchable"`
	Modifiable []string `mapstructure:"modifable"`
	Retain     []string `mapstructure:"retain"`
}

type Config struct {
	ApiPrefix   string                      `mapstructure:"api_prefix"`
	Collections map[string]CollectionConfig `mapstructure:"collections"`
	Api         []ApiConfig                 `mapstructure:"api"`
}

func Load(in io.Reader, configType string) {
	config := new(Config)

	v := viper.New()
	v.SetConfigType(configType)

	err := v.ReadConfig(in)
	if err != nil {
		return nil, err
	}

	config := new(Config)
	err = v.Unmarshal(config)
}

package config

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

var (
	// b pkg/common/config
	_, b, _, _ = runtime.Caller(0)
	root       = filepath.Join(filepath.Dir(b), "../../..")
)

// Config 全局的配置信息
var Config config

type config struct {
	Redis struct {
		DBAddress     string `yaml:"dbAddress"`
		DBMaxIdle     int    `yaml:"dbMaxIdle"`
		DBIdleTimeout int    `yaml:"dbIdleTimeout"`
		DBPassword    string `yaml:"dbPassword"`
	} `yaml:"redis"`
}

func init() {
	cfgName := root + "/config/config.yaml"
	viper.SetConfigFile(cfgName)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err.Error())
	}
	bytes, err := os.ReadFile(cfgName)
	if err != nil {
		panic(err.Error())
	}
	if err = yaml.Unmarshal(bytes, &Config); err != nil {
		panic(err.Error())
	}
	fmt.Println("load config:", Config)
}

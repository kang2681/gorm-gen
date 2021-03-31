package configs

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"sync/atomic"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"

	"github.com/kang2681/common/drivers"
)

var (
	Root       string
	ConfigPath string
	config     = &ConfigYaml{}
	av         atomic.Value
)

type ConfigYaml struct {
	Mysql []drivers.MysqlConfig `yaml:"mysql"`
}

func LoadConfig(configPath string) {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	if root, err := filepath.Abs(pwd); err == nil {
		Root = root + string(os.PathSeparator)
	} else {
		panic(err)
	}
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		panic(err)
	}
	if err := yaml.Unmarshal(data, config); err != nil {
		panic(err)
	}
	ConfigPath = configPath
	av.Store(config)
	logrus.Infof("%+v", config)
}

func SaveConfig() error {
	conf := Data()
	data, err := yaml.Marshal(conf)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(ConfigPath, data, os.ModePerm)
}

func Append(mcf drivers.MysqlConfig) error {
	conf := *Data()
	conf.Mysql = append(conf.Mysql, mcf)
	if err := SaveConfig(); err != nil {
		return err
	}
	av.Store(conf)
	return nil
}

func Data() *ConfigYaml {
	return av.Load().(*ConfigYaml)
}

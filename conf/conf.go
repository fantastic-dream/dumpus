package conf

import (
	"encoding/json"
	"github.com/pengcheng789/dumpus/log"
	"github.com/pengcheng789/dumpus/util"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

const (
	configName        = "conf/dumpus.conf"
	defaultIpAddr = "localhost"
	defaultServerPort = 14869
)

var (
	config Config
	logger = log.Logger
)

func Init() {
	logger.Info("Loading configuration.")

	var err error
	config, err = loadConfig()
	if err != nil {
		logger.WithError(err).Warn("Load configuration failure, use default configuration")
	}

	logger.Info("Loaded configuration: ", config)
}

func GetConfig() Config {
	return config
}

type Config struct {
	Server struct {
		IpAddr string `json:"ipAddr" yaml:"ipAddr"`
		Port   int    `json:"port" yaml:"port"`
	} `json:"server" yaml:"server"`
}

// 设置默认值
func newConfig() Config {
	config := Config{}

	config.Server.IpAddr = defaultIpAddr
	config.Server.Port = defaultServerPort

	return config
}

func loadConfig() (Config, error) {
	runtimeDir, err := util.RuntimeDir()
	if err != nil {
		return newConfig(), err
	}

	configFilePath := filepath.Join(runtimeDir, configName)
	data, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		return newConfig(), errors.Wrap(err, "Read config file failure.")
	}

	config := newConfig()
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return newConfig(), errors.Wrap(err, "Unmarshal config failure.")
	}

	return config, nil
}

func (c Config) String() string {
	s, _ := json.Marshal(c)
	return string(s)
}

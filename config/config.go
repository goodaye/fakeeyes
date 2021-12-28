package config

import (
	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
)

// DBConfig config for database
type DBConfig struct {
	Datasource  string `json:"Datasource"`
	MaxIdleConn int    `json:"MaxIdleConn"`
	MaxOpenConn int    `json:"MaxOpenConn"`
	SQLLog      bool   `json:"SQLLog"`
}

type Config struct {
	HTTP struct {
		Port int `json:"Port"`
	} `json:"HTTP"`
	MetaDB DBConfig `json:"MetaDB"`
}

var (
	// 全局配置中心地址
	GlobalConfig Config
	// defaultConfigFile 默认的配置文件地址
	DefaultConfigFile string = "../bin/config.toml"
	//defaultLogPath 默认的log目录
	DefaultLogsPath string = "../bin/logs"
)

// Loggers logger集合
var Loggers = struct {
	WebLogger    *logrus.Logger
	AccessLogger *logrus.Logger
}{}

// EnvName 环境变量名
var EnvName = struct {
	TemplatesPath  string
	ConfigFilePath string
	LogsPath       string
}{
	TemplatesPath:  "TemplatesPath",
	ConfigFilePath: "ConfigFilePath",
	LogsPath:       "LogsPath",
}

// return default config
func NewConfig() Config {

	config := Config{}
	// 初始化默认值
	config.HTTP.Port = 80
	return config
}

// LoadConfigFile load config from file
func LoadConfigFile(filepath string) error {
	//解析rainbow配置
	_, err := toml.DecodeFile(filepath, &GlobalConfig)
	return err
}

// SetConfigFile 设置配置文件路径
func SetConfigFile(fp string) {
	DefaultConfigFile = fp
}

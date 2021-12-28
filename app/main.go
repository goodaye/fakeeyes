package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/goodaye/fakeeyes/config"
)

var (
	// 配置文件地址
	configFilePath string = "./config.toml"
	//casbin文件地址
	templatesPath string = "./templates"
	logsPath      string = "./logs"
	// 执行的命令
	command string
)

func main() {
	var err error
	_ParseOption()
	config.SetConfigFile(configFilePath)
	os.Setenv(config.EnvName.LogsPath, logsPath)

	fmt.Println("command : ", command)
	switch command {
	case "syncdb":
		err = SyncDB()
	case "dropdb":
		err = DropDB()
	case "startweb":
		err = StartWeb()
	// case "syncdb":
	// 	err = SyncDB()
	// case "dropdb":
	// 	err = DropDB()

	default:
		err = fmt.Errorf("unspport command :%s", command)
	}
	if err != nil {
		fmt.Println(err.Error())
	}
}

// _ParseOption  解析命令行参数
func _ParseOption() {
	flag.Usage = Usage
	flag.StringVar(&configFilePath, "f", "", "Config File Path")
	flag.StringVar(&templatesPath, "t", "", "Templates Dir Path")
	flag.StringVar(&logsPath, "l", "", "Log File Path")
	flag.Parse()
	command = flag.Arg(0)
}

//Usage 打印CommandLine Usage
func Usage() {
	helpHeader := `fakeeye cmdline 
Options:
fakeeye command [ options ]
command : 
	startweb : 启动web服务
`
	fmt.Println(helpHeader)
	flag.PrintDefaults()
}

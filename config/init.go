package config

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/goodaye/wire"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func init() {
	wire.Append(svc{})
}

type svc struct {
	wire.BaseService
}

func (s svc) Init() error {

	var err error

	GlobalConfig = NewConfig()
	fileinfo, err := os.Stat(DefaultConfigFile)
	if err != nil {
		return errors.Wrap(err, DefaultConfigFile)
	}
	if !fileinfo.Mode().IsRegular() {
		return ErrorPathIsNotRegularFile
	}
	absfilepath, err := filepath.Abs(DefaultConfigFile)
	if err != nil {
		return err
	}
	fmt.Printf("Loading  Config File  '%s' \n", absfilepath)

	err = LoadConfigFile(DefaultConfigFile)
	if err != nil {
		return err
	}

	// INIT Log
	err = initlogger()
	return err
}

func initlogger() error {
	var err error
	logenvpath := os.Getenv(EnvName.LogsPath)
	if logenvpath != "" {
		DefaultLogsPath = logenvpath
	}
	fileinfo, err := os.Stat(DefaultLogsPath)
	if os.IsNotExist(err) {
		err = os.MkdirAll(DefaultLogsPath, os.ModePerm)
		if err != nil {
			return err
		}
	} else {
		if !fileinfo.IsDir() {
			return fmt.Errorf("logs template is not dir ")
		}
	}
	fmt.Println("default Logs Path:", DefaultLogsPath)
	// weblog
	Loggers.WebLogger, err = CreateLogger("web.log")
	if err != nil {
		return err
	}
	// accesslog
	Loggers.AccessLogger, err = CreateLogger("access.log")
	if err != nil {
		return err
	}
	return nil
}

//CreateLogger 创建日志句柄
func CreateLogger(filename string) (*logrus.Logger, error) {

	writer, err := rotatelogs.New(
		path.Join(DefaultLogsPath, filename+".%Y-%m-%d"),
		rotatelogs.WithLinkName(path.Join(DefaultLogsPath, filename)),
		// 每24 小时轮转一次, 最大保留5天的日志数据
		rotatelogs.WithMaxAge(24*7*time.Hour),
		rotatelogs.WithRotationTime(time.Duration(1)*time.Hour),
		rotatelogs.WithRotationSize(100*1024*1024),
	)
	if err != nil {
		return nil, err
	}
	logger := logrus.New()
	logger.Out = writer
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetReportCaller(true)
	return logger, nil
}

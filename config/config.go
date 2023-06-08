package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"go.uber.org/zap"
)

var C = &Config{}

type Config struct {
	Database DatabaseConf `json:"database" toml:"database"`
	Logger   LoggerConf   `json:"logger" toml:"logger"`
}

type DatabaseConf struct {
	DSN string `json:"dsn" toml:"dsn"` //db dsn
}

type LoggerConf struct {
	Env    string `json:"env" toml:"env"`
	Level  string `json:"level" toml:"level"`
	Output string `json:"output" toml:"output"`
}

func LoadConfigFromToml(cfgFile string) error {
	if _, err := toml.DecodeFile(cfgFile, C); err != nil {
		return err
	}

	zapLevel := zap.NewAtomicLevel()
	if err := zapLevel.UnmarshalText([]byte(C.Logger.Level)); err != nil {
		panic(err.Error())
	}
	var zapConf = zap.Config{}

	if env := C.Logger.Env; env == "dev" {
		zapConf = zap.NewDevelopmentConfig()
	} else {
		zapConf = zap.NewProductionConfig()
	}
	zapConf.Level = zapLevel
	if C.Logger.Output != "" {
		zapConf.OutputPaths = []string{C.Logger.Output}
		zapConf.ErrorOutputPaths = []string{C.Logger.Output}
	}
	if logger, err := zapConf.Build(); err != nil {
		panic(err.Error())
	} else {
		zap.RedirectStdLog(logger)
		zap.ReplaceGlobals(logger)
	}
	zap.L().Info(fmt.Sprintf("load config: %s", cfgFile))
	return nil
}

package skloglib

import (
	"io/ioutil"

	"github.com/uber-go/zap"
	//	"github.com/natefinch/lumberjack"
	//	"github.com/DaishinUehara/suikin/skconflib"
)

const (
	defaultConfFile = "./conf/log.yml"
)

// SkLogger is Suikin Logger.
type SkLogger struct {
	logger    zap.Logger
	logconfig zap.Config
}

// GetLogger is get logger.
func (skl *SkLogger) GetLogger() ( zap.Logger , error) {
	var conffile string
	skl.logger, err = zap.NewProduction()
	if err != nil {
		return skl.logger, err
	}
	conffile = conf.GetLogConfig("CONFIG_YAML")
	if conffile == "" {
		conffile = defaultConfFile
	} 
	configYaml, err := ioutil.ReadFile(conffile)
	if err != nil {
		return skl.logger, err
	}
	if err = yaml.Unmarshal(configYaml, &skl.logconfig); err != nil {
		return skl.logger, err
	}
	skl.logger, err := myConfig.Build()
	if err != nil {
		return skl.logger, err
	}
	return skl.logger, err

}

package skloglib

import (
	"io/ioutil"

	"github.com/DaishinUehara/suikin/skconflib"
	"go.uber.org/zap"

	//	"github.com/natefinch/lumberjack"
	yaml "gopkg.in/yaml.v2"
)

const (
	defaultConfFile = "conf/log.yml"
)

// SkLogger is Suikin Logger.
type SkLogger struct {
	logger    *zap.Logger
	logconfig zap.Config
}

var conf skconflib.SkConf

// GetLogger is get logger.
func (skl *SkLogger) GetLogger() (*zap.Logger, error) {
	var conffile string
	var err error
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
	skl.logger, err = skl.logconfig.Build()
	if err != nil {
		return skl.logger, err
	}
	return skl.logger, err

}

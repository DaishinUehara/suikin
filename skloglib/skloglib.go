package skloglib

import (
	"io/ioutil"

	"path/filepath"

	"github.com/DaishinUehara/suikin/skconflib"
	"github.com/DaishinUehara/suikin/skerrlib"
	"go.uber.org/zap"

	//	"github.com/natefinch/lumberjack"
	yaml "gopkg.in/yaml.v2"
)

const (
	defaultConfFile = "../conf/log.yml"
	// defaultConfFile = "C:\\work\\02_go\\go\\src\\github.com\\DaishinUehara\\suikin\\conf\\log.yml"
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
	var fpath string
	err = nil
	skl.logger, err = zap.NewProduction()
	if err != nil {
		return nil, skerrlib.ErrLoggerGenerate{Err: err, StackTrace: skerrlib.PrintCallStack()}
	}
	conffile = conf.GetLogConfig("CONFIG_YAML")
	if conffile == "" {
		conffile = defaultConfFile
	}
	fpath, err = filepath.Abs(conffile)
	if err != nil {
		return nil, skerrlib.ErrGetAbsolutePath{Err: err, StackTrace: skerrlib.PrintCallStack()}
	}

	configYaml, err := ioutil.ReadFile(fpath)
	if err != nil {
		return nil, skerrlib.ErrReadFile{Err: err, StackTrace: skerrlib.PrintCallStack()}
	}
	logconfig := &skl.logconfig
	if err = yaml.Unmarshal(configYaml, logconfig); err != nil {
		return nil, skerrlib.ErrYamlMapping{Err: err, StackTrace: skerrlib.PrintCallStack()}
	}
	skl.logger, err = logconfig.Build()
	if err != nil {
		return nil, err
	}
	return skl.logger, err

}

package skloglib

import (
	"github.com/DaishinUehara/suikin/skcallstacklib"
	"github.com/DaishinUehara/suikin/skcmnlib"
	"github.com/DaishinUehara/suikin/skconflib"
	"github.com/DaishinUehara/suikin/skerrlib"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	lumberjack "gopkg.in/natefinch/lumberjack.v2"
	yaml "gopkg.in/yaml.v2"
)

const (
	defaultLogFormatConf   = "../conf/log_format.yml"
	defalutLogRotationConf = "../conf/log_rotation.yml"
)

// SkLog is Logger for Suikin.
var SkLog SkLogger

// SkLogger is Suikin Logger Struct.
type SkLogger struct {
	logger    *zap.Logger
	logconfig zap.Config
	lmblogger *lumberjack.Logger
}

var conf skconflib.SkConf

// GetLogger is get logger.
func (skl *SkLogger) GetLogger() (*zap.Logger, error) {
	var formatconf string
	var err error
	var rotationconf string
	err = nil
	if skl.logger != nil {
		return skl.logger, err
	}

	formatconf = conf.GetLogConfig("FORMAT_YAML")
	if formatconf == "" {
		formatconf = defaultLogFormatConf
	}
	configYaml, err := skcmnlib.ReadByteFile(formatconf)
	if err != nil {
		// 設定ファイルの読み込みに失敗した場合
		return nil, err
	}

	if err = yaml.Unmarshal(configYaml, &skl.logconfig); err != nil {
		// 設定ファイルを構造体にセットできなかった場合
		return nil, &skerrlib.ErrYamlMapping{Err: err, StackTrace: skcallstacklib.PrintCallStack()}
	}

	enc := zapcore.NewJSONEncoder(skl.logconfig.EncoderConfig)

	rotationconf = conf.GetLogConfig("ROTATION_YAML")
	if rotationconf == "" {
		rotationconf = defalutLogRotationConf
	}

	rotateYaml, err := skcmnlib.ReadByteFile(rotationconf)
	if err != nil {
		// 設定ファイルの読み込みに失敗した場合
		return nil, err
	}
	if err = yaml.Unmarshal(rotateYaml, &skl.lmblogger); err != nil {
		// 設定ファイルを構造体にセットできなかった場合
		return nil, &skerrlib.ErrYamlMapping{Err: err, StackTrace: skcallstacklib.PrintCallStack()}
	}

	sink := zapcore.AddSync(skl.lmblogger)

	skl.logger = zap.New(zapcore.NewCore(enc, sink, skl.logconfig.Level))

	return skl.logger, err

}

// ErrLogOutput Output Log "Index Error"
func ErrLogOutput(e error) error {
	logger, err := SkLog.GetLogger()
	if err != nil {
		return err
	}
	logger.Error(e.Error())
	return err
}

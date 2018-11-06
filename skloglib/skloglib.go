package skloglib

import (
	"io/ioutil"
	"path/filepath"

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

// SkLogger is Suikin Logger.
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

	formatconf = conf.GetLogConfig("FORMAT_YAML")
	if formatconf == "" {
		formatconf = defaultLogFormatConf
	}
	configYaml, err := ReadByteFile(formatconf)
	if err != nil {
		// 設定ファイルの読み込みに失敗した場合
		return nil, err
	}

	if err = yaml.Unmarshal(configYaml, &skl.logconfig); err != nil {
		// 設定ファイルを構造体にセットできなかった場合
		return nil, skerrlib.ErrYamlMapping{Err: err, StackTrace: skerrlib.PrintCallStack()}
	}

	enc := zapcore.NewJSONEncoder(skl.logconfig.EncoderConfig)

	rotationconf = conf.GetLogConfig("ROTATION_YAML")
	if rotationconf == "" {
		rotationconf = defalutLogRotationConf
	}

	rotateYaml, err := ReadByteFile(rotationconf)
	if err != nil {
		// 設定ファイルの読み込みに失敗した場合
		return nil, err
	}
	if err = yaml.Unmarshal(rotateYaml, &skl.lmblogger); err != nil {
		// 設定ファイルを構造体にセットできなかった場合
		return nil, skerrlib.ErrYamlMapping{Err: err, StackTrace: skerrlib.PrintCallStack()}
	}

	sink := zapcore.AddSync(skl.lmblogger)

	skl.logger = zap.New(zapcore.NewCore(enc, sink, skl.logconfig.Level))

	return skl.logger, err

}

// ReadByteFile is ReadFile absolute path or relative path.
func ReadByteFile(path string) ([]byte, error) {
	rconf, err := filepath.Abs(path)
	if err != nil {
		// 設定ファイルの絶対パス取得に失敗した場合
		return nil, skerrlib.ErrGetAbsolutePath{Err: err, StackTrace: skerrlib.PrintCallStack()}
	}
	data, err := ioutil.ReadFile(rconf)
	if err != nil {
		// 設定ファイルの読み込みに失敗した場合
		return nil, skerrlib.ErrReadFile{Err: err, StackTrace: skerrlib.PrintCallStack()}
	}
	return data, err
}

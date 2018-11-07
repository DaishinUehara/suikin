package skloglib_test

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/DaishinUehara/suikin/skcmnlib"
	"github.com/DaishinUehara/suikin/skloglib"
	"go.uber.org/zap"
	yaml "gopkg.in/yaml.v2"
)

var loglib skloglib.SkLogger

const (
	logfilepath = "../logs/suikin.log"
)

type LogData struct {
	Level string   `json:"Level"`
	Time  string   `json:"Time"`
	Msg   string   `json:"Msg"`
	Key   string   `json:"key"`
	Now   string   `json:"now"`
	Stack []string `json:"stack"`
}

func TestGetLogger(t *testing.T) {
	var logd LogData

	abslog, err := filepath.Abs(logfilepath)
	if _, err := os.Stat(abslog); err == nil {
		if err := os.Remove(abslog); err != nil {
			t.Errorf("[NG]:err=%v,file not found %v\n", err, abslog)
		}
	}

	logger, err := loglib.GetLogger()
	if err != nil {
		t.Errorf("[NG]:err=%v\n", err)
	} else {
		defer logger.Sync()
		arr := []string{"abc", "def", "hij"}
		var now time.Time
		now = time.Now()
		logger.Info("Hello zap", zap.String("key", "value"), zap.Time("now", now), zap.Strings("stack", arr))
		// Parse Log File
		logdata, _ := skcmnlib.ReadByteFile(logfilepath)
		if err = yaml.Unmarshal(logdata, &logd); err != nil {
			// 設定ファイルを構造体にセットできなかった場合
			t.Errorf("[NG]:err=%v\n", err)
		} else {
			if logd.Level == "INFO" && logd.Msg == "Hello zap" && logd.Key == "value" && logd.Stack[0] == "abc" && logd.Stack[1] == "def" && logd.Stack[2] == "hij" {
				t.Logf("[OK]:Level=%s,Msg=%s,key=%s,logd.Stack=%v\n", logd.Level, logd.Msg, logd.Key, logd.Stack)
			} else {
				t.Errorf("[NG]:Level=%s,Msg=%s,key=%s,logd.Stack=%v\n", logd.Level, logd.Msg, logd.Key, logd.Stack)
			}
		}

	}
}

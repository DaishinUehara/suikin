package skloglib_test

import (
	"testing"
	"time"

	"encoding/json"

	"github.com/DaishinUehara/suikin/skcmnlib"
	"github.com/DaishinUehara/suikin/skfilelib"
	"github.com/DaishinUehara/suikin/skloglib"
	"go.uber.org/zap"
)

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
	var logger *zap.Logger
	var err error
	// var abslog string

	err = skfilelib.RmFile(logfilepath)
	if err != nil {
		t.Errorf("[NG]:rmLog:err=%v\n", err)
	}

	logger, err = skloglib.SkLog.GetLogger()
	//	logger, err = loglib.GetLogger()
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
		if err = json.Unmarshal(logdata, &logd); err != nil {
			// 設定ファイルを構造体にセットできなかった場合
			t.Errorf("[NG]:err=%v\n", err)
		} else {
			if logd.Level == "INFO" && logd.Msg == "Hello zap" && logd.Key == "value" && logd.Stack[0] == "abc" && logd.Stack[1] == "def" && logd.Stack[2] == "hij" {
				t.Logf("[OK]:Level=%s,Msg=%s,key=%s,logd.Stack=%v\n", logd.Level, logd.Msg, logd.Key, logd.Stack)
			} else {
				t.Errorf("[NG]:Level=%s,Msg=%s,key=%s,logd.Stack=%v\n", logd.Level, logd.Msg, logd.Key, logd.Stack)
			}
		}

		logger, err = skloglib.SkLog.GetLogger() // 再度のlogger取得
		// logger, err = loglib.GetLogger() // 再度のlogger取得
		if err != nil {
			t.Errorf("[NG]:err=%v\n", err)
		}
		if err = json.Unmarshal(logdata, &logd); err != nil {
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

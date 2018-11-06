package skloglib_test

import (
	"bytes"
	"io"
	"os"
	"testing"
	"time"

	"github.com/DaishinUehara/suikin/skloglib"
	"go.uber.org/zap"
)

var loglib skloglib.SkLogger

func TestGetLogger(t *testing.T) {
	logger, err := loglib.GetLogger()
	if err != nil {
		t.Errorf("[NG]:err=%v\n", err)
	} else {
		defer logger.Sync()
		arr := []string{"abc", "def", "hij"}
		stdout := os.Stdout // backup os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w // connect
		logger.Info("Hello zap", zap.String("key", "value"), zap.Time("now", time.Now()), zap.Strings("stack", arr))
		os.Stdout = stdout
		w.Close()
		var buf bytes.Buffer
		io.Copy(&buf, r)
		str := buf.String()
		t.Logf("[OK]:%s\n", str)
	}
}

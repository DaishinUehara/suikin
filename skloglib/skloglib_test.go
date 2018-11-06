package skloglib_test

import (
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
		logger.Info("Hello zap", zap.String("key", "value"), zap.Time("now", time.Now()), zap.Strings("stack", arr))
		t.Logf("[OK]:\n")
	}
}

package skloglib_test

import (
	"testing"

	"github.com/DaishinUehara/suikin/skloglib"
)

var loglib skloglib.SkLogger

func TestGetLogger(t *testing.T) {
	logger, err := loglib.GetLogger()
	if err != nil {
		t.Errorf("[NG]:err=%v\n", err)
	} else {
		defer logger.Sync()
		t.Logf("[OK]:\n")
	}
}

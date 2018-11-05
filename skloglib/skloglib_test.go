package skloglib_test

import (
	"testing"

	"github.com/DaishinUehara/suikin/skloglib"
)

var loglib skloglib.SkLogger

func TestGetLogger(t *testing.T) {
	logger, err := loglib.GetLogger()
	defer logger.Sync()

}

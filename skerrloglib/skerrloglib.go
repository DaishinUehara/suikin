package skerrloglib

import (
	"github.com/DaishinUehara/suikin/skerrlib"
	"github.com/DaishinUehara/suikin/skloglib"
	"go.uber.org/zap"
)

// ErrOutOfIndex Error out of index.
type ErrOutOfIndex skerrlib.ErrOutOfIndex

// LogOutput Output Log "Index Error"
func (e *ErrOutOfIndex) LogOutput() error {
	logger, err := skloglib.SkLog.GetLogger()
	if err != nil {
		return err
	}
	logger.Error("Out of Index Error.", zap.String("ArrayName", e.ArrayName), zap.Int("Index", e.Index), zap.Strings("StackTrace", e.StackTrace))
	return err
}

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

// ErrNotInitialized Error out of index.
type ErrNotInitialized skerrlib.ErrNotInitialized

// LogOutput Output Log "Index Error"
func (e *ErrNotInitialized) LogOutput() error {
	logger, err := skloglib.SkLog.GetLogger()
	if err != nil {
		return err
	}
	logger.Error("Not Initialized Error.", zap.String("NoInitializedItem", e.NoInitializedItem), zap.Strings("StackTrace", e.StackTrace))
	return err
}

// ErrFlushBuffer Flush Buffer Error.
type ErrFlushBuffer skerrlib.ErrFlushBuffer

// LogOutput Output Log "Flush Buffer Error"
func (e *ErrFlushBuffer) LogOutput() error {
	logger, err := skloglib.SkLog.GetLogger()
	if err != nil {
		return err
	}
	logger.Error("Flush Buffer Error.", zap.String("ErrorItem", e.ErrorItem), zap.Strings("StackTrace", e.StackTrace))
	return err
}

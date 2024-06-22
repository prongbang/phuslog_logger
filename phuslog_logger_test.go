package phuslog_logger_test

import (
	"testing"

	"github.com/prongbang/phuslog_logger"
)

type data struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func TestLogger(t *testing.T) {
	logs := phuslog_logger.NewPhuslogLogger()
	d := &data{
		Key:   "Key",
		Value: "Value",
	}
	logs.Info("test", "data", d)
	logs.Warn("test", "data", d)
	logs.Debug("test", "data", d)
	logs.Error("test", "data", d)
}

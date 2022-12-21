package logging

import (
	"chat-jobsity/internal/config"
	"encoding/json"
	"go.uber.org/zap"
)

type SimpleLogger interface {
	Debug(msg string, keysAndValues ...interface{})
	Info(msg string, keysAndValues ...interface{})
	Error(msg string, keysAndValues ...interface{})
	Warn(msg string, keysAndValues ...interface{})
}

type SimpleLoggerImpl struct {
	log *zap.Logger
}

func (s SimpleLoggerImpl) Debug(msg string, keysAndValues ...interface{}) {
	s.log.Debug(msg)
}

func (s SimpleLoggerImpl) Info(msg string, keysAndValues ...interface{}) {
	s.log.Info(msg)
}

func (s SimpleLoggerImpl) Error(msg string, keysAndValues ...interface{}) {
	s.log.Error(msg)
}

func (s SimpleLoggerImpl) Warn(msg string, keysAndValues ...interface{}) {
	s.log.Warn(msg)
}

func NewSimpleLogger(config config.ConfigStore) SimpleLogger {
	rawJSON := []byte(`{
	  "level": "info",
	  "encoding": "json",
	  "outputPaths": ["stdout"],
	  "errorOutputPaths": ["stderr"],
	  "encoderConfig": {
	    "messageKey": "message",
	    "levelKey": "level",
	    "levelEncoder": "lowercase"
	  }
	}`)

	var cfg zap.Config
	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		panic(err)
	}
	logger := zap.Must(cfg.Build())
	defer logger.Sync()

	logger.Info("logger construction succeeded")

	return &SimpleLoggerImpl{
		log: logger,
	}
}

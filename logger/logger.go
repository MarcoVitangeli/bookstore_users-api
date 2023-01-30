package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	log *zap.Logger
)

func init() {
	logCfg := zap.Config{
		OutputPaths: []string{"stdout"},
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel), // Info, Warning, Error, but not Debug (Debug is below info)
		Encoding:    "json",                              // because we are going to use this in Elasticsearch
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "level",                       // with this, we say that we want a key-value pair like: "level":"info"
			TimeKey:      "time",                        // "time":"2019-11-10T23:00:23-0700"
			MessageKey:   "msg",                         // "msg":"this is the logging line"
			EncodeTime:   zapcore.ISO8601TimeEncoder,    // format for the time
			EncodeLevel:  zapcore.LowercaseLevelEncoder, // add colors for the output in the console
			EncodeCaller: zapcore.ShortCallerEncoder,    // shortness the caller path
		},
	}
	var err error
	if log, err = logCfg.Build(); err != nil {
		panic(err)
	}
}

func Info(msg string, tags ...zap.Field) {
	log.Info(msg, tags...)
	log.Sync()
}

func Error(msg string, err error, tags ...zap.Field) {
	if err != nil {
		tags = append(tags, zap.Error(err))
	}
	log.Error(msg, tags...)
	log.Sync()
}

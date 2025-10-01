package logger

import (
	"go.uber.org/zap"
	"workerPool1/internal/entity"
)

type Logger struct {
	sugar *zap.SugaredLogger
}

func New(environment string) (*Logger, error) {
	var l *zap.Logger
	var err error
	var cfg zap.Config

	if environment == entity.ProdEnvName {
		cfg = zap.NewProductionConfig()
	} else {
		cfg = zap.NewDevelopmentConfig()
	}

	cfg.OutputPaths = []string{"stdout"}
	cfg.ErrorOutputPaths = []string{"stderr"}
	l, err = cfg.Build(zap.AddCaller(), zap.AddCallerSkip(1))
	if err != nil {
		return nil, err
	}

	sugar := l.Sugar()
	return &Logger{sugar: sugar}, nil
}

func (l *Logger) Infof(template string, args ...interface{}) {
	l.sugar.Infof(template, args...)
}

func (l *Logger) Errorf(template string, args ...interface{}) {
	l.sugar.Errorf(template, args...)
}

func (l *Logger) Warnf(template string, args ...interface{}) {
	l.sugar.Warnf(template, args...)
}

func (l *Logger) Debugf(template string, args ...interface{}) {
	l.sugar.Debugf(template, args...)
}

func (l *Logger) Info(msg string, fields ...zap.Field) {
	l.sugar.Desugar().Info(msg, fields...)
}

func (l *Logger) Error(msg string, fields ...zap.Field) {
	l.sugar.Desugar().Error(msg, fields...)
}

func (l *Logger) Warn(msg string, fields ...zap.Field) {
	l.sugar.Desugar().Warn(msg, fields...)
}

func (l *Logger) Debug(msg string, fields ...zap.Field) {
	l.sugar.Desugar().Debug(msg, fields...)
}

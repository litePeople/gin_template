package logx

import "go.uber.org/zap"

type Logger interface {
	Debug(args ...interface{})
	Debugf(template string, args ...interface{})
	Debugx(msg string, args ...zap.Field)

	Info(args ...interface{})
	Infof(template string, args ...interface{})
	Infox(msg string, args ...zap.Field)

	Error(args ...interface{})
	Errorf(template string, args ...interface{})
	Errorx(msg string, args ...zap.Field)

	Panic(args ...interface{})
	Panicf(template string, args ...interface{})
	Panicx(msg string, args ...zap.Field)

	Fatal(args ...interface{})
	Fatalf(template string, args ...interface{})
	Fatalx(msg string, args ...zap.Field)

	Logger() *zap.Logger
	SugarLogger() *zap.SugaredLogger
}

type logxLogger struct {
	logger *zap.Logger
}

func (ll *logxLogger) Debug(args ...interface{}) {
	ll.logger.Sugar().Debug(args)
}

func (ll *logxLogger) Debugf(template string, args ...interface{}) {
	ll.logger.Sugar().Debugf(template, args)
}

func (ll *logxLogger) Debugx(msg string, args ...zap.Field) {
	ll.logger.Debug(msg, args...)
}

func (ll *logxLogger) Info(args ...interface{}) {
	ll.logger.Sugar().Info(args)
}

func (ll *logxLogger) Infof(template string, args ...interface{}) {
	ll.logger.Sugar().Infof(template, args)
}

func (ll *logxLogger) Infox(msg string, args ...zap.Field) {
	ll.logger.Info(msg, args...)
}

func (ll *logxLogger) Error(args ...interface{}) {
	ll.logger.Sugar().Error(args)
}

func (ll *logxLogger) Errorf(template string, args ...interface{}) {
	ll.logger.Sugar().Errorf(template, args)
}

func (ll *logxLogger) Errorx(msg string, args ...zap.Field) {
	ll.logger.Error(msg, args...)
}

func (ll *logxLogger) Panic(args ...interface{}) {
	ll.logger.Sugar().Panic(args)
}

func (ll *logxLogger) Panicf(template string, args ...interface{}) {
	ll.logger.Sugar().Panicf(template, args)
}

func (ll *logxLogger) Panicx(msg string, args ...zap.Field) {
	ll.logger.Panic(msg, args...)
}

func (ll *logxLogger) Fatal(args ...interface{}) {
	ll.logger.Sugar().Fatal(args)
}

func (ll *logxLogger) Fatalf(template string, args ...interface{}) {
	ll.logger.Sugar().Fatalf(template, args)
}

func (ll *logxLogger) Fatalx(msg string, args ...zap.Field) {
	ll.logger.Fatal(msg, args...)
}

func (ll *logxLogger) Logger() *zap.Logger {
	return ll.logger
}

func (ll *logxLogger) SugarLogger() *zap.SugaredLogger {
	return ll.logger.Sugar()
}

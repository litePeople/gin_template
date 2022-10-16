package logx

import "go.uber.org/zap"

func Debug(args ...interface{}) {
	cli.logger.Debug(args)
}

func Debugf(template string, args ...interface{}) {
	cli.logger.Debugf(template, args)
}

func Debugx(msg string, args ...zap.Field) {
	cli.logger.Debug(msg, args)
}

func Info(args ...interface{}) {
	cli.logger.Info(args)
}

func Infof(template string, args ...interface{}) {
	cli.logger.Infof(template, args)
}

func Infox(msg string, args ...zap.Field) {
	cli.logger.Infox(msg, args...)
}

func Error(args ...interface{}) {
	cli.logger.Error(args)
}

func Errorf(template string, args ...interface{}) {
	cli.logger.Errorf(template, args)
}

func Errorx(msg string, args ...zap.Field) {
	cli.logger.Errorx(msg, args...)
}

func Panic(args ...interface{}) {
	cli.logger.Panic(args)
}

func Panicf(template string, args ...interface{}) {
	cli.logger.Panicf(template, args)
}

func Panicx(msg string, args ...zap.Field) {
	cli.logger.Panicx(msg, args...)
}

func Fatal(args ...interface{}) {
	cli.logger.Fatal(args)
}

func Fatalf(template string, args ...interface{}) {
	cli.logger.Fatalf(template, args)
}

func Fatalx(msg string, args ...zap.Field) {
	cli.logger.Fatalx(msg, args...)
}

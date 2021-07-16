package log

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

var Log Logger

type Logger struct {
	*zap.Logger
}


func Init(logProd bool, outPath string, filed ...zap.Field)error{

	var config zap.Config
	if logProd {
		config = zap.NewProductionConfig()
	}else{
		config = zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeTime = func(time time.Time, encoder zapcore.PrimitiveArrayEncoder) {
			encoder.AppendString(time.Format("2006-01-02 15:04:05"))
		}
		if outPath != "" {
			config.Encoding = "json"
			config.OutputPaths = []string{outPath}
		}
	}

	logger, err := config.Build(zap.AddCallerSkip(1), zap.AddStacktrace(zap.FatalLevel))
	if err != nil {
		return err
	}
	Log.Logger = logger.With(filed...)
	return nil
}

func (log *Logger) Debugf(str string,args ...interface{})  {
	log.Logger.Debug(fmt.Sprintf(str,args...))
}

func (log *Logger) Infof(str string,args ...interface{})  {
	log.Logger.Info(fmt.Sprintf(str,args...))
}

func (log *Logger) Warnf(str string,args ...interface{})  {
	log.Logger.Warn(fmt.Sprintf(str,args...))
}

func (log *Logger) Errorf(str string,args ...interface{})  {
	log.Logger.Error(fmt.Sprintf(str,args...))
}

func (log *Logger) Fatalf(str string,args ...interface{})  {
	log.Logger.Fatal(fmt.Sprintf(str,args...))
}

func (log *Logger) Debug(args ...interface{}) {
	log.Logger.Debug(fmt.Sprint(args...))
}

func (log *Logger) Info(args ...interface{}) {
	log.Logger.Debug(fmt.Sprint(args...))
}

func (log *Logger) Warn(args ...interface{}) {
	log.Logger.Debug(fmt.Sprint(args...))
}

func (log *Logger) Error(args ...interface{}) {
	log.Logger.Debug(fmt.Sprint(args...))
}

func (log *Logger) Fatal(args ...interface{}) {
	log.Logger.Debug(fmt.Sprint(args...))
}

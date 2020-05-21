package main

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)


var logger *zap.Logger
func formatEncodeTime(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(fmt.Sprintf("%d%02d%02d_%02d%02d%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second()))
}

func FormateLog(args []interface{}) *zap.Logger {
	log := logger.With(ToJsonData(args))
	//fmt.Println(reflect.TypeOf(log))
	//fmt.Println(log)
	return log
}

func Level(msg string, args ...interface{}) {
	FormateLog(args).Sugar().Infof(msg)
}

func ToJsonData(args []interface{}) zap.Field {
	//fmt.Println(reflect.TypeOf(args))
	det := make([]string, 0)
	if len(args) > 0 {
		for _, v := range args {
			det = append(det, fmt.Sprintf("%+v", v))
		}
	}
	zap := zap.Any("detail", det) //k,v
	return zap
}
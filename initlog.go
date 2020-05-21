package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

//初始化日志

func InitZapLog() {
				cfg := zap.Config{
				Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
				Development: true,
				Encoding:    "json",
				EncoderConfig: zapcore.EncoderConfig{
				TimeKey:        "t", //输出时间的key名
				LevelKey:       "level",//输出日志级别的key名
				NameKey:        "logger",
				CallerKey:      "caller",
				MessageKey:     "msg",  //输入信息的key名
				StacktraceKey:  "trace",
				LineEnding:     zapcore.DefaultLineEnding, //每行的分隔符。基本zapcore.DefaultLineEnding 即"\n"
				EncodeLevel:    zapcore.LowercaseLevelEncoder, //基本zapcore.LowercaseLevelEncoder。将日志级别字符串转化为小写
				EncodeTime:     formatEncodeTime, //输出的时间格式
				EncodeDuration: zapcore.SecondsDurationEncoder, //一般zapcore.SecondsDurationEncoder,执行消耗的时间转化成浮点型的秒
				EncodeCaller:   zapcore.ShortCallerEncoder,  //一般zapcore.ShortCallerEncoder，以包/文件:行号 格式化调用堆栈
				},
				OutputPaths:      []string{"/tmp/gogitvisualize.log"},
				ErrorOutputPaths: []string{"/tmp/gogitvisualize.log"},
				InitialFields: map[string]interface{}{
				"app": "test",
		},
		}
		var err error
		logger, err = cfg.Build()
		if err != nil {
			panic("log init fail:" + err.Error())
		}
}



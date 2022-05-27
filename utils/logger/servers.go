package logger

import (
	"flag"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

var (
	Logger         *logrus.Logger
	logAppFileName = flag.String("app", "./storage/app/"+currentTime+".log", "App log file name")
)

func init() {
	//日志实例化
	Logger = logrus.New()
	//拼接写入文件路径
	file := *logAppFileName
	//设置输出到文件
	Logger.Out = os.Stdout
	//设置日志级别
	Logger.SetLevel(logrus.DebugLevel)
	//设置日志格式
	Logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	//设置日志切割 rotatelogs
	logWriter, _ := rotatelogs.New(
		file+".%Y%m%d.log",
		//生成软链 指向最新的日志文件
		rotatelogs.WithLinkName(file),
		//文件最大保存时间
		rotatelogs.WithMaxAge(7*24*time.Hour),
		//设置日志切割时间间隔(1天)(隔多久分割一次)
		rotatelogs.WithRotationTime(24*time.Hour),
	)

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}

	lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	Logger.AddHook(lfHook)

}

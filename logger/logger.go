package logger

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"time"
)

// 简单封装了一下logrus这个日志系统，之后如果想要记录日志的话，只需要先初始化配置一下
// 然后就可以直接调用的对应的方法，然后.info（）就可以。

var Logger *logrus.Logger

func InitLogrus(path string) {
	logger := logrus.New()
	logger.SetLevel(logrus.TraceLevel)
	content, _ := rotatelogs.New(
		path+"&Y%m%d%H%M%S",
		rotatelogs.WithLinkName(path),
		rotatelogs.WithRotationCount(30),
		//rotatelogs.WithMaxAge(720 *time.Hour)
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	logger.SetOutput(content)
	Logger = logger
}

//一般程序中输出的调试信息
func LogDebug(x map[string]interface{}) *logrus.Entry {
	return Logger.WithFields(x)
}

//关键操作，核心流程的日志；
func LogInfo(x map[string]interface{}) *logrus.Entry {
	return Logger.WithFields(x)
}

//警告信息，提醒程序员注意
func LogWarn(x map[string]interface{}) *logrus.Entry {
	return Logger.WithFields(x)
}

//错误日志，需要查看原因
func LogError(x map[string]interface{}) *logrus.Entry {
	return Logger.WithFields(x)
}

//致命错误，出现错误时程序无法正常运转。输出日志后，程序退出；
func LogFatal(x map[string]interface{}) *logrus.Entry {
	return Logger.WithFields(x)
}

//记录日志，然后panic
func LogPanic(x map[string]interface{}) *logrus.Entry {
	return Logger.WithFields(x)
}

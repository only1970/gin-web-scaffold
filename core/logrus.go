package core

import (
	"bytes"
	"fmt"
	"gvd_server/config"
	"os"
	"path"

	"github.com/sirupsen/logrus"
)

// 颜色
const (
	red    = 31
	yellow = 33
	blue   = 36
	gray   = 37
)

type LogRequest struct {
	AppName  string
	LogPath  string
	NoData   bool
	NoErr    bool
	NoGlobal bool
}

type LogFormatter struct{}

// DateHook 按天分割写入日志文件的hook
type DateHook struct {
	file     *os.File
	fileDate string
	appname  string
	logPath  string
}

func (DateHook) Levels() []logrus.Level {
	return logrus.AllLevels

}
func (hook DateHook) Fire(entry *logrus.Entry) error {
	timer := entry.Time.Format("2006-01-02")
	line, _ := entry.String()
	if hook.fileDate == timer {
		hook.file.Write([]byte(line))
		return nil
	}

	//时间不等
	hook.file.Close()
	os.MkdirAll(path.Join(hook.logPath), os.ModePerm)
	filename := path.Join(hook.logPath, fmt.Sprintf("%s-%s-%s", hook.appname, timer, "info.log"))

	hook.file, _ = os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
	hook.fileDate = timer
	hook.file.Write([]byte(line))
	return nil
}

// ErrHook 按日志级别分割写入日志文件的hook
type ErrHook struct {
	file     *os.File
	fileDate string
	appname  string
	logPath  string
}

func (ErrHook) Levels() []logrus.Level {
	return []logrus.Level{logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel}

}
func (hook ErrHook) Fire(entry *logrus.Entry) error {
	timer := entry.Time.Format("2006-01-02")
	line, _ := entry.String()
	if hook.fileDate == timer {
		hook.file.Write([]byte(line))
		return nil
	}

	//时间不等
	hook.file.Close()
	os.MkdirAll(path.Join(hook.logPath), os.ModePerm)
	filename := path.Join(hook.logPath, fmt.Sprintf("%s-%s-%s", hook.appname, timer, "error.log"))

	hook.file, _ = os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
	hook.fileDate = timer
	hook.file.Write([]byte(line))
	return nil
}

// Format 实现Formatter(entry *logrus.Entry) ([]byte, error)接口
func (t *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	//根据不同的level去展示颜色
	var levelColor int
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = gray
	case logrus.WarnLevel:
		levelColor = yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = red
	default:
		levelColor = blue
	}

	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	//自定义日期格式
	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	if entry.HasCaller() {
		//自定义文件路径
		funcVal := entry.Caller.Function
		fileVal := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)
		//自定义输出格式
		fmt.Fprintf(b, "[%s] \x1b[%dm[%s]\x1b[0m %s %s %s\n", timestamp, levelColor, entry.Level, fileVal, funcVal, entry.Message)
	} else {
		fmt.Fprintf(b, "[%s] \x1b[%dm[%s]\x1b[0m %s\n", timestamp, levelColor, entry.Level, entry.Message)
	}
	return b.Bytes(), nil
}

func InitLogger(request *config.Logrus) *logrus.Logger {
	// var request LogRequest
	// if len(requestList) > 0 {
	// 	request = requestList[0]
	// }
	if request.LogPath == "" {
		request.LogPath = "logs"
	}
	if request.AppName == "" {
		request.AppName = "gvd"
	}
	mLog := logrus.New()               //新建一个实例
	mLog.SetOutput(os.Stdout)          //设置输出类型
	mLog.SetReportCaller(true)         //开启返回函数名和行号
	mLog.SetFormatter(&LogFormatter{}) //设置自己定义的Formatter
	mLog.SetLevel(logrus.DebugLevel)   //设置最低的Level
	if !request.NoData {
		mLog.AddHook(&DateHook{
			appname: request.AppName,
			logPath: request.LogPath,
		})
	}
	if !request.NoErr {
		mLog.AddHook(&ErrHook{
			appname: request.AppName,
			logPath: request.LogPath,
		})
	}
	if !request.NoGlobal {
		InitDefaultLogger() //全局logrus生效
	}
	return mLog
}

func InitDefaultLogger() {
	// 全局log
	logrus.SetOutput(os.Stdout)          //设置输出类型
	logrus.SetReportCaller(true)         //开启返回函数名和行号
	logrus.SetFormatter(&LogFormatter{}) //设置自己定义的Formatter
	logrus.SetLevel(logrus.DebugLevel)   //设置最低的Level
}

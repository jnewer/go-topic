package main

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"time"
)

var log = logrus.New()

func init() {
	//log.Out = os.Stdout
	//log.Formatter = &logrus.JSONFormatter{}
	//logrus.SetLevel(logrus.InfoLevel)

	path := "logrus/message.log"

	writer, _ := rotatelogs.New(
		path+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(path),
		rotatelogs.WithMaxAge(time.Duration(180)*time.Second),
		rotatelogs.WithRotationTime(time.Duration(60)*time.Second),
	)

	log.SetOutput(writer)
}
func main() {
	//logrus.Info("go logrus")
	//logrus.SetLevel(logrus.InfoLevel)
	//logrus.SetLevel(logrus.ErrorLevel)

	//logrus.Info("info")
	//logrus.Debug("debug")
	//logrus.Warn("warn")
	//logrus.Error("error")
	//logrus.Panic("panic")
	//logrus.Fatal("fatal")

	//logrus.SetFormatter(&logrus.TextFormatter{})
	//logrus.SetFormatter(&logrus.JSONFormatter{})
	//logrus.Infof("info")

	//logrus.SetOutput(os.Stdout)
	//logrus.Info("info")
	//
	//file, err := os.OpenFile("./logrus/logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	//
	//if err == nil {
	//	logrus.SetOutput(file)
	//} else {
	//	logrus.Infof("日志文件创建失败")
	//}

	//logrus.Info("info")

	//var event = "下订单"
	//var topic = "order"
	//var key = 1001
	//
	//logrus.WithFields(logrus.Fields{
	//	"event": event,
	//	"topic": topic,
	//	"key":   key,
	//}).Fatal("事件发送失败")

	//logrus.SetReportCaller(true)
	//logrus.Infof("info")

	//log.WithFields(logrus.Fields{
	//	"event": event,
	//	"topic": topic,
	//	"key":   key,
	//}).Fatal("事件发送失败")

	for {
		log.Info("hello,world")
		time.Sleep(time.Duration(2) * time.Second)
	}
}

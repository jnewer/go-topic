package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetPrefix("order:")
	logFile, err := os.OpenFile("./log/test.log", os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Panic("打开日志文件异常")
	}

	log.SetOutput(logFile)
}
func TestLogger() {
	var (
		buf    bytes.Buffer
		logger = log.New(&buf, "logger: ", log.Lshortfile)
	)

	logger.Print("go log test")
	fmt.Print(&buf)
}

func TestLogPrint() {
	log.Println("TestLogPrint")
}

func TestLogPanic() {
	defer fmt.Println("Panic defer")
	log.Panicln("TestLogPanic")
}

func TestLogFatal() {
	defer fmt.Println("Fatal defer")
	log.Fatalln("TestLogFatal")
}
func main() {
	//log.Println("go-log")
	//
	//logger := log.Default()
	//logger.SetFlags(log.Llongfile)
	//logger.Output(0, "0 calldepth")
	//logger.Output(1, "1 calldepth")
	//logger.Output(2, "3 calldepth")

	//TestLogger()

	TestLogPrint()
	TestLogPanic()
	//TestLogFatal()
}

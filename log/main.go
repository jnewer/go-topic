package main

import (
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

func main() {
	log.Println("go-log")
}

package util

import (
	"log"
	"os"
	"time"
)

//日志文件目录
const logPathDefault = "/log"

func Log() *log.Logger {
	name := TimeFormatDateOnly(time.Now())
	filename := logPathDefault + "/" + name + ".log"
	//检查目录是否存在  不存在则创建
	err := IsDir(logPathDefault)
	if err != nil {
		log.Panic("创建日志文件失败")
		return nil
	}

	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open error log file:", err)
	}
	return log.New(file,
		"DEBUG: ",
		log.Ldate|log.Ltime|log.Llongfile)
}

//指定目录
func LogPath(p string) *log.Logger {
	name := TimeFormatDateOnly(time.Now())
	filename := p + "/" + name + ".log"
	//检查目录是否存在  不存在则创建
	err := IsDir(p)
	if err != nil {
		log.Panic("创建日志文件失败")
		return nil
	}

	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open error log file:", err)
	}
	return log.New(file,
		"DEBUG: ",
		log.Ldate|log.Ltime|log.Llongfile)
}

package main

import (
	"Database/configs"
	"Database/pkg/mysql"
	"fmt"
	"log"
	"os"
)

// 初始化日志信息
func init() {
	LogPath := configs.LogPath
	LogFile, err := os.OpenFile(LogPath,
		os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0664)
	if err != nil {
		fmt.Printf("创建日志路径%s失败, 错误信息%e\n", LogPath, err)
		return
	}
	log.SetOutput(LogFile)
	log.SetPrefix("[mysql] ")
	log.SetFlags(log.Ltime | log.Ldate | log.Lmicroseconds)
}

func main() {
	//mysql.InsertOperation()
	mysql.SelectOperation()
}

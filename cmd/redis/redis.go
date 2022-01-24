package main

import (
	"Database/configs"
	"Database/pkg/redis"
	"fmt"
	"log"
	"os"
	"strconv"
)

// 初始化日志信息
func init() {
	LogPath := configs.LogPathRedis
	LogFile, err := os.OpenFile(LogPath,
		os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0664)
	if err != nil {
		fmt.Printf("创建日志路径%s失败, 错误信息%e\n", LogPath, err)
		return
	}
	log.SetOutput(LogFile)
	log.SetPrefix("[redis] ")
	log.SetFlags(log.Ltime | log.Ldate | log.Lmicroseconds)
}

func main() {
	//redis.OperateWrite("name", 100)
	//redis.OperateGet("10")
	//s := "内部使用[]byte实现，不像直接运算符这种会产生很多临时的字符串，但是内部的逻辑比较复杂，有很多额外的判断，还用到了interface，所以性能也不是很好"
	for i := 1; i <= 2000000; i++ {
		// strconv.Itoa(i) int行转string
		//redis.OperateWrite(strconv.Itoa(i), strings.Join([]string{s, strconv.Itoa(i)}, ""))
		//redis.OperateWrite(strconv.Itoa(i), i)
		redis.OperateGet(strconv.Itoa(i))
		fmt.Println(strconv.Itoa(i), i)
	}
}

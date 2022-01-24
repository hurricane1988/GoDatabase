package redis

import (
	"Database/configs"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"log"
)

// Pool 定义redis链接变量
var pool *redis.Pool
var password string = "prodTsp2redis2021!"

// 初始化redis链接
func init() {
	// 实例化一个连接池
	pool = &redis.Pool{

		MaxIdle:   16, //最初的连接数量
		MaxActive: 0,  //连接池最大连接数量,不确定可以用0（0表示自动定义），按需分配
		// 连接超时时间
		IdleTimeout: 200, //连接关闭时间 300秒 （300秒不使用自动关闭）
		Dial: func() (redis.Conn, error) { //要连接的redis数据库
			conn, err := redis.Dial("tcp", configs.RedisHost)
			if err != nil {
				return nil, err
			}
			//return redis.Dial("tcp", configs.RedisHost)
			if password != "" {
				if _, err := conn.Do("AUTH", "prodTsp2redis2021!"); err != nil {
					err := conn.Close()
					if err != nil {
						return nil, err
					}
					return nil, err
				}
			}
			return conn, err
		},
	}
}

// OperateWrite  执行string写入操作
func OperateWrite(key string, value interface{}) {
	conn := pool.Get() // 从链接池中获取一个链接
	defer func(conn redis.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println("redis执行set操作失败")
		}
	}(conn) // 函数运行结束，回收连接到连接池
	// 写入操作
	_, err := conn.Do("Set", key, value)
	if err != nil {
		fmt.Println("执行redis操作失败,错误信息", err)
		log.Println("执行redis操作失败,错误信息", err)
		return
	}
	fmt.Printf("插入 {%s : %v}成功!\n", key, value)
	log.Printf("插入 {%s : %v}成功!\n", key, value)
	//pool.Close() //关闭连接池
}

// OperateGet 查询操作函数
func OperateGet(key string) {
	conn := pool.Get() // 从链接池中获取一个链接
	defer func(conn redis.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println("redis执行get操作失败")
		}
	}(conn) // 函数运行结束，回收连接到连接池
	result, err := redis.Int(conn.Do("Get", key))
	if err != nil {
		log.Println("查询key值失败,错误信息", err)
		return
	}
	fmt.Println(result)
	log.Println("查询执行成功!")
	_ = pool.Close() //关闭连接池
}

package configs

const LogPath = "./logs/mysql.log"

// MysqlHost 数据库配置连接信息变量
var (
	MysqlHost string = "root:root@tcp(127.0.0.1:3306)/test"
	//RedisHost string = "127.0.0.1:6379"
	//RedisHost string = "10.6.223.178:30010"
	RedisHost string = "10.120.106.28:30218"
)

const LogPathRedis = "./logs/redis.log"

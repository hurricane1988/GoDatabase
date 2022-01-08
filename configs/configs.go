package configs

const LogPath = "./logs/mysql.log"

// MysqlHost 数据库配置连接信息变量
var (
	MysqlHost string = "root:root@tcp(127.0.0.1:3306)/test"
)

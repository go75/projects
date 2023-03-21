package config

const (
	JwtSecret = "peerwise"
	GrpcAdrr = "127.0.0.1:10001"
	DbDsn = "root:1234@tcp(127.0.0.1:3306)/todolist?charset=utf8mb4&parseTime=True&loc=Local"
	EtcdAddr = "127.0.0.1:2379"
	RedisAddr = "127.0.0.1:6379"
)

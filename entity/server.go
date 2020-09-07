package entity

// Server 服务端配置
type Server struct {
	Type     string // client/server
	IP       string
	Port     int
	Password string
}

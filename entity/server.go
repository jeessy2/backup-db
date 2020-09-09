package entity

// Server 服务端配置
type Server struct {
	Type      string // client/server
	DBType    string // 数据库类型
	UploadURL string
}

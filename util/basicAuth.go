package util

import (
	"backup-db/entity"
	"bytes"
	"encoding/base64"
	"net/http"
	"strings"
)

// GetCurrentUser 获得当前登陆用户
func GetCurrentUser() *entity.User {
	conf, err := GetConfig()
	if err == nil {
		return &conf.User
	}

	return &entity.User{}
}

// ViewFunc func
type ViewFunc func(http.ResponseWriter, *http.Request)

// BasicAuth basic auth
func BasicAuth(f ViewFunc) ViewFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		basicAuthPrefix := "Basic "

		// 获取 request header
		auth := r.Header.Get("Authorization")
		// 如果是 http basic auth
		if strings.HasPrefix(auth, basicAuthPrefix) {
			// 解码认证信息
			payload, err := base64.StdEncoding.DecodeString(
				auth[len(basicAuthPrefix):],
			)
			if err == nil {
				pair := bytes.SplitN(payload, []byte(":"), 2)
				user := GetCurrentUser()
				if len(pair) == 2 &&
					bytes.Equal(pair[0], []byte(user.Username)) &&
					bytes.Equal(pair[1], []byte(user.Password)) {
					// 执行被装饰的函数
					f(w, r)
					return
				}
			}
		}

		// 认证失败，提示 401 Unauthorized
		// Restricted 可以改成其他的值
		w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
		// 401 状态码
		w.WriteHeader(http.StatusUnauthorized)
	}
}

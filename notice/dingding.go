package notice

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const dingWebHookPrefix = "https://oapi.dingtalk.com/robot/send?access_token="

// DingDing 钉钉通知
type DingDing struct {
	WebHook string // webhook地址
	Secret  string // 加签
}

// DingDingResp 钉钉通知返回消息
type DingDingResp struct {
	Errcode int
	errmsg  string
}

// CanBeSend 能否发送
func (ding *DingDing) CanBeSend() bool {
	return ding.WebHook != "" && ding.Secret != ""
}

// SendMessage 发送钉钉消息
func (ding *DingDing) SendMessage(title, message string) (err error) {

	if strings.HasPrefix(ding.WebHook, dingWebHookPrefix) {

		// sign
		h := hmac.New(sha256.New, []byte(ding.Secret))
		signData := fmt.Sprintf("%d\n%s", time.Now().Unix()*1000, ding.Secret)
		h.Write([]byte(signData))
		encrypt := url.QueryEscape(base64.StdEncoding.EncodeToString(h.Sum(nil)))

		// msg
		msg := `
		{
			"msgtype": "text", 
			"text": {
				"content": "` + message + `"
			}
		}
		`

		// url
		var addr = ding.WebHook
		addr += fmt.Sprintf("&timestamp=%d", time.Now().Unix()*1000)
		addr += fmt.Sprintf("&sign=%s", encrypt)

		// post
		resp, err := http.Post(addr, "application/json", strings.NewReader(msg))
		if err != nil {
			log.Println("请求钉钉接口失败。ERR: ", err)
			return err
		}
		defer resp.Body.Close()
		respBytes, _ := ioutil.ReadAll(resp.Body)

		dingResp := &DingDingResp{}
		err = json.Unmarshal(respBytes, dingResp)

		if err != nil {
			log.Println("解析钉钉返回数据失败")
			return err
		}

		if dingResp.Errcode == 0 {
			log.Printf("发送钉钉消息成功。\n")
			return nil
		}

		return fmt.Errorf("发送钉钉消息失败。errcode: %d, errmsg: %s", dingResp.Errcode, dingResp.errmsg)
	}

	return errors.New("没有webhook地址")
}

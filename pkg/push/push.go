// pkg/push/push.go
package push

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// PushResult 推送结果
type PushResult struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

// Push 推送消息
func Push(msg string) *PushResult {
	var (
		err error
		url = fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=%s", wechatConfig.AccessToken())
		req = map[string]interface{}{
			"touser":  wechatConfig.UserID,
			"msgtype": "text",
			"agentid": wechatConfig.AgentID,
			"text": map[string]string{
				"content": msg,
			},
		}
		body, _ = json.Marshal(req)
	)

	resp, err := http.Post(url, "application/json", bytes.NewReader(body))
	if err != nil {
		return &PushResult{ErrCode: -1, ErrMsg: err.Error()}
	}
	defer resp.Body.Close()

	var result PushResult
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return &PushResult{ErrCode: -1, ErrMsg: err.Error()}
	}

	return &result
}

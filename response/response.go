package response

import (
	"fmt"
)

//CodeAndMessage is the structure when fetch access token failed
type CodeAndMessage struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

//Format the msg
func (msg *CodeAndMessage) Format() string {
	return fmt.Sprintf("Error from wechat server: %d : %s", msg.ErrCode, msg.ErrMsg)
}

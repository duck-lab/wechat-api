package replyToWechat

import (
	"encoding/xml"
	"fmt"
	"testing"
)

func TestMarshal(t *testing.T) {
	xmlBytes, err := xml.Marshal(Model{
		UserOpenID: "123",
	})
	fmt.Println(string(xmlBytes), err)
}

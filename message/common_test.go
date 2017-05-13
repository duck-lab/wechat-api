package message

import (
	"encoding/xml"
	"fmt"
	"testing"
)

func TestMarshal(t *testing.T) {
	xmlBytes, err := xml.Marshal(Model{
		head: head{UserOpenID: "123"},
	})
	fmt.Println(string(xmlBytes), err)
}

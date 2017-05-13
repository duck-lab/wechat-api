package signature

import (
	"crypto/sha1"
	"io"
	"sort"
	"strings"
)

// Check if the signature is valid
func Check(echostr string, msgEncrypt string, nonce string, timestamp string,
	signature string) bool {
	strs := []string{echostr, msgEncrypt, nonce, timestamp}
	sort.Strings(strs)
	strToSign := strings.Join(strs, "")
	h := sha1.New()
	io.WriteString(h, strToSign)
	checkSum := string(h.Sum(nil))
	if signature == checkSum {
		return true
	}
	return false
}

// 参考文档1：https://en.wikipedia.org/wiki/Hash-based_message_authentication_code
// 参考文档2：http://www.jokecamp.com/blog/examples-of-creating-base64-hashes-using-hmac-sha256-in-different-languages/#go
package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

// MSigner 根据secret和文本进行hmac-sha1加密
func MSigner(message, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha1.New, key)
	h.Write([]byte(message))

	r := fmt.Sprintf("meiqia_sign:%v", base64.URLEncoding.EncodeToString([]byte(hex.EncodeToString(h.Sum(nil)))))
	return r
}

func main() {
	fmt.Println(MSigner("Message", "secret"))
}

package dzdecode

import (
	"bytes"
	"encoding/base64"
	"strings"
)

func Base64Decode(token string) string {
	decodeBytes, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		panic(err)
	}
	return string(decodeBytes)

}

func TokenDecode(token string) string {
	tkData := strings.Split(token, ".")[1]
	num := len(tkData) % 4
	if num > 0 {
		var buff bytes.Buffer
		buff.WriteString(tkData)

		for i := 0; i < num; i++ {
			buff.WriteString("=")
		}
		//fmt.Println(buff.String())
		return Base64Decode(buff.String())
	}
	return Base64Decode(tkData)

}

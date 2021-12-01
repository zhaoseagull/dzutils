package dzdecode

import (
	"encoding/json"
	"fmt"
	"time"
)

var v interface{}

func CheckExpire(jwtToken string) bool {
	data := TokenDecode(jwtToken)
	fmt.Println(data)
	json.Unmarshal([]byte(data), &v)
	dMap := v.(map[string]interface{})
	exp := dMap["exp"]
	//fmt.Println(exp.(float64))
	//对比时间戳，判断jwt是否过期
	jwtTime := int64(exp.(float64))
	nowTime := time.Now().Unix()
	//fmt.Printf("time now: %v, time token: %v", nowTime, jwtTime)
	return jwtTime > nowTime
}

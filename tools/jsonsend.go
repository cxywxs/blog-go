package tools

import (
	"encoding/json"
	"log"
)

//标准json返回结构体
type JosnStd struct {
	Code    int
	Message string
	Data    interface{}
}

const Ok = 200
const OkMsg = "request successful"

//写入string类型的data
func JsonSendOkString(data string) string {
	jsonstd := JosnStd{
		Code:    Ok,
		Message: OkMsg,
		Data:    data,
	}
	jsondata, err := json.Marshal(jsonstd)
	if err != nil {
		log.Println(err)
	}
	return string(jsondata)
}

//写入struct或数组类型的data
func JsonSendOkStruct(data interface{}) string {
	jsonstd := JosnStd{
		Code:    Ok,
		Message: OkMsg,
		Data:    data,
	}
	jsondata, err := json.Marshal(jsonstd)
	if err != nil {
		log.Println(err)
	}
	return string(jsondata)
}

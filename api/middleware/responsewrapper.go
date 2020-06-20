package middleware

import (
	"encoding/json"
	"fmt"
)

type ResponseCode int

// request success
const (
	_ ResponseCode = iota + 1000
	Success
	SuccessNodata
)

// request fail
const (
	_ ResponseCode = iota + 2000
	Failed
	FailedParamsError
	FailedUnknown ResponseCode = 2099
)

// request api error
const (
	_ ResponseCode = iota + 3000
	ApiNotExists
	ApiNotPer
)

// system error
const (
	SystemError ResponseCode = 9999
)

func (eCode ResponseCode) String() string {
	return responseMap[eCode]
}

var responseMap = map[ResponseCode]string{
	Success:           "查询成功",
	SuccessNodata:     "查询成功无记录",
	Failed:            "查询失败",
	FailedParamsError: "查询失败，参数为空或格式错误",
	FailedUnknown:     "查询失败，未知错误",
	ApiNotExists:      "访问的接口不存在",
	ApiNotPer:         "没有该接口的访问权限",
	SystemError:       "系统异常",
}

// Interface uniform return format
type ResponseWrapper struct {
	Code int         `json:"code"` // status code
	Msg  string      `json:"msg"`  // return message
	Data interface{} `json:"data"` // return data
}

// custom responseWrapper
func NewResponseWrapper(code int, msg string, data interface{}) *ResponseWrapper {
	rWrapper := new(ResponseWrapper)
	rWrapper.Code = code
	rWrapper.Msg = msg
	rWrapper.Data = data
	return rWrapper
}

// response to json
func (rWrapper *ResponseWrapper) ToJson() string {
	allBytes, err := json.Marshal(rWrapper)
	if err != nil {
		allBytes, _ = json.Marshal(MarkErrorUnknown("json.Marshal(data) error:" + err.Error()))
	}
	return string(allBytes)
}

// The parameter is empty or the format of the parameter is wrong
func MarkErrorParam(msg string) *ResponseWrapper {
	rWrapper := new(ResponseWrapper)
	rWrapper.Code = int(FailedParamsError)
	rWrapper.Msg = fmt.Sprintf("%s:%s", FailedParamsError.String(), msg)
	rWrapper.Data = nil
	return rWrapper
}

// select fail
func MarkError(msg string) *ResponseWrapper {
	rWrapper := new(ResponseWrapper)
	rWrapper.Code = int(Failed)
	rWrapper.Msg = fmt.Sprintf("%s:%s", Failed.String(), msg)
	rWrapper.Data = nil
	return rWrapper
}

// select fail and don`t error
func MarkErrorUnknown(msg string) *ResponseWrapper {
	rWrapper := new(ResponseWrapper)
	rWrapper.Code = int(FailedUnknown)
	rWrapper.Msg = fmt.Sprintf("%s:%s", FailedUnknown.String(), msg)
	rWrapper.Data = nil
	return rWrapper
}

// select success but don`t have data
func MarkSuccessNotData() *ResponseWrapper {
	rWrapper := new(ResponseWrapper)
	rWrapper.Code = int(SuccessNodata)
	rWrapper.Msg = SuccessNodata.String()
	rWrapper.Data = nil
	return rWrapper
}

// select success and have data
func MarkSuccess(data interface{}) *ResponseWrapper {
	rWrapper := new(ResponseWrapper)
	rWrapper.Code = int(Success)
	rWrapper.Msg = Success.String()
	rWrapper.Data = data
	return rWrapper
}

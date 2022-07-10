package main

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

// 测试errorWrapper

/**
 * 思路，调用errorWrapper，会返回一个函数，函数的参数为responsewriter和request，首先需要建立这两个参数
 * 本身测试的函数，入参为appHandler
 * 从response中获取code 和body进行对比查看是否正确
**/

func Test_errwra(t *testing.T) {
	tests := []struct {
		h       appHandler
		code    int
		message string
	}{
		{errPanic, 500, "Internal Server Error"},
		{errUserError, 400, "custom error trigger"},
		{errNotFound, 404, "Not Found"},
		{errPermisson, 403, "Forbidden"},
		{errDefault, 500, "Internal Server Error"},
		{errNIl, 200, "no error"},
	}

	for _, v := range tests {
		rlt := errWrapper(v.h) // rlt为一个函数，入参为responseWriter和request

		// 使用内置的包建立responsewriter和request
		responseWriter := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, "http://www/baidu.com", nil)

		rlt(responseWriter, request) // 错误信息都写到了resposneWriter里

		body, _ := ioutil.ReadAll(responseWriter.Body) // 读取所有的response.body 返回类型为byte

		body_str := strings.Trim(string(body), "\n")
		if responseWriter.Code != v.code || body_str != v.message {
			t.Errorf("期望的code为%d,得到的是%d", v.code, responseWriter.Code)
			t.Errorf("期望的内容为  %s,得到的是%s", v.message, body_str)
		}

	}

}

// 用于系统错误的触发
func errPanic(writer http.ResponseWriter, request *http.Request) error {
	panic(1)
	return nil
}

// 用于用户自定义错误的触发
//type userError string

// func (e userError) Error() string {
// 	return e.Message()
// }

// func (e userError) Message() string {
// 	return string(e)
// }

func errUserError(writer http.ResponseWriter, request *http.Request) error {
	return userError("custom error trigger")
}

// NotFound 情况
func errNotFound(writer http.ResponseWriter, request *http.Request) error {
	return os.ErrNotExist
}

// 权限不足error
func errPermisson(writer http.ResponseWriter, request *http.Request) error {
	return os.ErrPermission
}

// 其他类型
func errDefault(writer http.ResponseWriter, request *http.Request) error {
	return errors.New("abc")
}

// 其他类型
func errNIl(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

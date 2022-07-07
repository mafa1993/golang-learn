package main

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/gommon/log"
)

func main() {
	http.HandleFunc("/list/", func(writer http.ResponseWriter, request *http.Request) {
		path := request.URL.Path // url的path部分

		path = path[len("/list/"):] // 字符串切分，取list/后面的部分

		file, err := os.Open(path)

		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError) // 三个参数responsewriter   错误信息  http状态码，err.Error() 会将内部的错误信息暴露出去
			return
			//panic(err)
		}
		defer file.Close()

		all, err := ioutil.ReadAll(file)

		if err != nil {
			panic(err)
		}

		writer.Write(all)
	})

	http.HandleFunc("/list1/", errWrapper(handleList))

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		panic(err)
	}
}

// ===== 错误处理封装
type appHandler func(writer http.ResponseWriter, request *http.Request) error

// 用户自定义错误，一些内部错误，不能传递给前端展示，这个用来定义一些可以暴露给用户的错误
type customError interface {
	error
	Message() string
}

// 定义一个类型，用来实现customError接口
type userErr string

// 实现error的Error接口
func (e userErr) Error() string {
	return e.Message()
}

func (e userErr) Message() string {
	return string(e)
}

// 在这里进行error的处理，类似于catch
func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {

	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			r := recover() // recover 可以为nil的
			if r != nil {
				log.Panicf("%v", r)

				http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()

		err := handler(writer, request)
		log.Warn(err.Error()) // log 非内置库的，内置库的为log.Printf()
		if err != nil {
			code := http.StatusOK
			if customError, ok := err.(customError); ok {
				// 调用自定义的Message方法，返回错误信息
				http.Error(writer, customError.Message(), code)
				return
			}

			switch {
			case os.IsNotExist(err): // 目录、文件等不存在类型的报错
				code = http.StatusNotFound
			case os.IsPermission(err): //权限不足类型的错误
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}

			http.Error(writer, http.StatusText(code), code) // http.StatusText() 将状态码翻译成对应的文本
			return
		}
	}
}

// 所有error都直接return，在外层处理
const prefix = "/list1/"

func handleList(writer http.ResponseWriter, request *http.Request) error {
	path := request.URL.Path // url的path部分
	if strings.Index(path, prefix) != 0 {
		// 自定义测错误暴露不出去
		// return errors.New("必须以" + prefix + "开头")
		// 这里userErr相当于给string定义的别名，userErr() 相当于string()，这里返回error类型，所以userErr需要实现error接口
		return userErr("必须以" + prefix + "开头")
	}
	path = path[len("/list1/"):] // 字符串切分，取list/后面的部分

	file, err := os.Open(path)

	if err != nil {
		return err
	}
	defer file.Close()

	all, err := ioutil.ReadAll(file)

	if err != nil {
		return err
	}

	writer.Write(all)
	return nil
}

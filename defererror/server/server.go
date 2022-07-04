package main

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/labstack/gommon/log"
)

func main() {
	http.HandleFunc("/list/", func(writer http.ResponseWriter, request *http.Request) {
		path := request.URL.Path // url的path部分

		path = path[len("/list/"):] // 字符串切分，取list/后面的部分

		file, err := os.Open(path)
		defer file.Close()

		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError) // 三个参数responsewriter   错误信息  http状态码，err.Error() 会将内部的错误信息暴露出去
			return
			//panic(err)
		}

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

// 在这里进行error的处理，类似于catch
func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {

	return func(writer http.ResponseWriter, request *http.Request) {
		err := handler(writer, request)
		log.Warn(err.Error())  // log 非内置库的，内置库的为log.Printf()
		if err != nil {
			code := http.StatusOK
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
func handleList(writer http.ResponseWriter, request *http.Request) error {
	path := request.URL.Path // url的path部分

	path = path[len("/list1/"):] // 字符串切分，取list/后面的部分

	file, err := os.Open(path)
	defer file.Close()

	if err != nil {
		return err
	}

	all, err := ioutil.ReadAll(file)

	if err != nil {
		return err
	}

	writer.Write(all)
	return nil
}

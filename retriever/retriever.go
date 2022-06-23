/*
 * :Date: 2022-06-23 21:24:31
 * :LastEditTime: 2022-06-23 21:57:57
 * :Description:
 */
package retriever

type Retriever struct {
	Abc string
}

//只要其他类型实现了接口的方法，即实现了接口
func (r Retriever) Get(str string) string {
	return r.Abc
}

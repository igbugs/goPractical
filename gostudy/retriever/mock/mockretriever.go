package mock

import "fmt"

type Retriever struct {
	Contents string
}

func (r *Retriever) String() string {
	return fmt.Sprintf("Retriever: {Contents = %s}", r.Contents)
}

func (r *Retriever) Get(url string) string {
	return r.Contents
	// download 使用者只是定义了接口的样子，这里进行具体的接口所需要的Get方法的实现
}

func (r *Retriever) Post(url string, form map[string]string) string {
	r.Contents = form["contents"]
	return "OK"
}

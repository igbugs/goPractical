package filelisting

import (
	"net/http"
	"os"
	"io/ioutil"
	"strings"
)

const preurl = "/list/"

type userErrorA string

func (e userErrorA) Error() string {
	return e.Message()
}

func (e userErrorA) Message() string {
	return string(e)
}

func HandleFileList(writer http.ResponseWriter, request *http.Request) error {
	if strings.Index(request.URL.Path, preurl) != 0 {
		return userErrorA("path must start with " + preurl)
	}

	path := request.URL.Path[len(preurl):]
	//fmt.Println(path)
	// 例如: url 为 /list/fib.txt, 这行处理之后，才为真实的fib.txt 所在的路径
	file, err := os.Open(path)
	if err != nil {
		return err       // 所有的错误的不在进行处理，有errWrapper 包装后统一的处理
	}
	defer file.Close()

	all, err := ioutil.ReadAll(file)  // 读取文件的内容
	if err != nil {
		return err
	}

	writer.Write(all)		// 将读取的内容写回到 http.ResponseWriter中区
	return nil
}

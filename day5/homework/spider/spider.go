package main

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/satori/go.uuid"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"sync"
)

func getAllUrls() []string {
	var urls []string
	var url string
	for i := 0; i < 10; i++ {
		url = "http://www.meizitu.com/a/" + strconv.Itoa(i+1) + ".html"
		urls = append(urls, url)
	}
	return urls
}

func parseHtml(url string) int {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".postContent > p > img").Each(
		func(i int, s *goquery.Selection) {
			imgUrl, _ := s.Attr("src")
			// 启动协程下载图片
			wg.Add(1)
			os.Chdir(saveImagePath)
			go download(imgUrl)
		})

	//fmt.Printf("当前goroutine 数量: %d\n",runtime.NumGoroutine())

	return 0
}

// 下载图片
func download(imgUrl string) int {
	//fmt.Println(imgUrl)
	uid, _ := uuid.NewV1()
	filename := uid.String() + ".jpg"
	//fmt.Println(filename)

	resp, _ := http.Get(imgUrl)
	body, _ := ioutil.ReadAll(resp.Body)
	out, _ := os.Create(filename)
	io.Copy(out, bytes.NewReader(body))

	wg.Done()
	return 0
}

func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

var (
	saveImagePath = "./day5OriginImage"
	wg            sync.WaitGroup
)

func main() {
	if exist, _ := pathExists(saveImagePath); !exist {
		os.Mkdir(saveImagePath, os.ModePerm)
	}

	urls := getAllUrls()
	for _, url := range urls {
		parseHtml(url)
		fmt.Printf("当前goroutine 数量: %d\n", runtime.NumGoroutine())
	}

	wg.Wait()
}

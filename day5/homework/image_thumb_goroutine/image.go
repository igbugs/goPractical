package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var (
	imageChan chan *Task
	imageDir  string
)

// 此函数将取得 源图片的路径，生成保存文件的名字，组装成任务的结构体，分配任务 imagChan 的管道内，交由task.process() 函数进行处理
func procFile(path string, info os.FileInfo, err error) error {
	fmt.Printf("path:%s\n", path)
	if info.IsDir() {
		if err != nil {
			return err
		}
		return nil
	}

	if strings.HasSuffix(path, ".jpg") {
		baseFile := filepath.Base(path)
		pathDir := filepath.Dir(path)

		baseFileSeg := strings.Split(baseFile, ".")
		if strings.HasSuffix(baseFileSeg[0], "_thumb") {
			return err
		}

		saveFile := filepath.Join(pathDir,
			fmt.Sprintf("%s%s.%s", baseFileSeg[0], "_thumb", baseFileSeg[1]))

		fmt.Printf("save file:%s, path:%s dir:%s\n", saveFile, path, pathDir)
		task := &Task{
			imageFile: path,
			taskType:  TaskThumb,
			saveFile:  saveFile,
		}

		imageChan <- task
	}
	return err
}

func main() {
	//var imageDir string
	var threadNum int
	var chanSize int
	waitGroup := &sync.WaitGroup{}

	flag.StringVar(&imageDir, "dir", "", "--image dir")
	flag.IntVar(&threadNum, "c", 8, "--cocurrent thread num")
	flag.IntVar(&chanSize, "l", 1024, "--chanSize the image channel size")
	flag.Parse()

	var err error
	imageChan, err = initProgram(threadNum, chanSize, waitGroup)
	if err != nil {
		fmt.Printf("init program failed, err:%v\n", err)
		return
	}

	filepath.Walk(imageDir, procFile)
	close(imageChan)
	waitGroup.Wait()
}

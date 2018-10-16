package main

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	"image/png"
	"math"
	"os"

	"github.com/nfnt/resize"
	"path/filepath"
	"sync"
	"strings"
)

const DEFAULT_MAX_WIDTH float64 = 64
const DEFAULT_MAX_HEIGHT float64 = 64

// 计算图片缩放后的尺寸
func calculateRatioFit(srcWidth, srcHeight int) (int, int) {
	ratio := math.Min(DEFAULT_MAX_WIDTH/float64(srcWidth), DEFAULT_MAX_HEIGHT/float64(srcHeight))
	return int(math.Ceil(float64(srcWidth) * ratio)), int(math.Ceil(float64(srcHeight) * ratio))
}

// 生成缩略图
func makeThumbnail(imagePath, savePath string) error {
	file, _ := os.Open(imagePath)
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}

	b := img.Bounds()
	width := b.Max.X
	height := b.Max.Y

	w, h := calculateRatioFit(width, height)

	fmt.Println("width = ", width, " height = ", height)
	fmt.Println("w = ", w, " h = ", h)

	// 调用resize库进行图片缩放
	m := resize.Resize(uint(w), uint(h), img, resize.Lanczos3)

	// 需要保存的文件
	imgfile, err := os.Create(savePath)
	if err != nil {
		fmt.Printf("os create file, err:%v\n", err)
		return err
	}
	defer imgfile.Close()

	// 以PNG格式保存文件
	err = png.Encode(imgfile, m)
	if err != nil {
		return err
	}

	return nil
}

func getFilelist(path string) {
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		imagChan <- path
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
	close(imagChan)
	wg.Done()
}

var (
	wg sync.WaitGroup
	imagChan = make(chan string, 100)
)


func main() {
	wg.Add(1)
	go getFilelist("C:/GoProject/Go3Project/day5OriginImage")

	for imageFile := range imagChan {
		saveFile := strings.Split(imageFile, ".")[0] + "_thumb.png"
		wg.Add(1)

		go func(origin , resize string) {
			err := makeThumbnail(origin, resize)
			if err != nil {
				fmt.Printf("make thumbnail failed, err:%v\n", err)
				return
			}
			fmt.Printf("make thumbnail succ, path:%v\n", resize)
			wg.Done()
		}(imageFile, saveFile)
	}

	wg.Wait()
}

// 代码存在的问题:
// 1. 去除 91 行的wg.Add(1)的注释，处理 可能 40 张图片以下的可以完成处理，没有 deadlock 的发生
// 2. 但是 处理几百张的图片的时候， 91 行加注释，才能进行处理，不报死锁
// 3. 但是 处理几千张图片的时候，91 行加注释，还是会报死锁

// 我的 思路是 91 行的 wg.Add(1) 是因为 我在执行 getFilelist() 函数的时候，起了一个 goroutine 遍历文件夹的图片的名字，加入到 imagChan channel , 使用
// for ... range 从 imagChan 取出，每取出一个则起一个 goroutine 进行处理（起之前 wg.Add(1)），makeThumbnail() 处理完成后，wg.Done()，
// getFilelist() 在遍历完成目录之后，wg.Done()， 关闭getFilelist() 的计数。

// 我的疑惑:
// 1. 我感觉 wg.Add() 与 wg.Done() 都是一一对应的，为什么 还会死锁？

// 问题已经解决:
// 在进行wg.Done() 的时候 使用 defer wg.Done() 进行关闭，没有出现 死锁的情况，具体为何会出现 wg.Done() 不执行的问题，还是很疑惑？？

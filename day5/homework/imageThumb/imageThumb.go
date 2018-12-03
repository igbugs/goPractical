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
	"strings"
	"sync"
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

	//fmt.Println("width = ", width, " height = ", height)
	//fmt.Println("w = ", w, " h = ", h)

	// 调用resize库进行图片缩放
	m := resize.Resize(uint(w), uint(h), img, resize.Lanczos3)

	// 需要保存的文件
	imgfile, err := os.Create(savePath)
	if err != nil {
		//fmt.Printf("os create file, err:%v\n", err)
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
	defer wg.Done()
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
}

var (
	wg       sync.WaitGroup
	imagChan = make(chan string, 100)
)

func main() {
	wg.Add(1)
	go getFilelist("C:/GoProject/Go3Project/day5OriginImage")

	for imageFile := range imagChan {
		saveFile := strings.Split(imageFile, ".")[0] + "_thumb.png"
		wg.Add(1)

		go func(origin, resize string) {
			defer wg.Done()
			err := makeThumbnail(origin, resize)
			if err != nil {
				fmt.Printf("make thumbnail failed, err:%v\n", err)
				return
			}
			fmt.Printf("make thumbnail succ, path:%v\n", resize)
		}(imageFile, saveFile)

	}

	wg.Wait()
}

package qrcode

import (
	"github.com/boombuler/barcode/qr"
	"gin-blog/pkg/setting"
	"gin-blog/pkg/util"
	"gin-blog/pkg/file"
	"github.com/boombuler/barcode"
	"image/jpeg"
	)

type Qrcode struct {
	Url string
	Width int
	Height int
	Ext string
	Level qr.ErrorCorrectionLevel
	Mode qr.Encoding
}

const (
	EXT_JPG = ".jpg"
)

func NewQrcode(url string, width, height int, level qr.ErrorCorrectionLevel, mode qr.Encoding) *Qrcode {
	return &Qrcode{
		Url: url,
		Width: width,
		Height: height,
		Level: level,
		Mode: mode,
		Ext: EXT_JPG,
	}
}

func GetQrcodePath() string {
	return setting.AppSetting.QrcodeSavePath
}

func GetQrcodeFullPath() string {
	return setting.AppSetting.RuntimeRootPath + GetQrcodePath()
}

func GetQrcodeFullUrl(name string) string {
	return setting.AppSetting.PrefixUrl + "/" + GetQrcodePath() + name
}

func GetQrcodeFileName(value string) string {
	return util.EncodeMD5(value)
}

func (q *Qrcode) GetQrcodeExt() string {
	return q.Ext
}

func (q * Qrcode) CheckEncode(path string) bool {
	src := path + GetQrcodeFileName(q.Url) + q.GetQrcodeExt()
	if file.CheckNotExist(src) {
		return false
	}

	return true
}

func (q *Qrcode) Encode(path string) (string, string, error) {
	name := GetQrcodeFileName(q.Url) + q.GetQrcodeExt()
	src := path + name
	if file.CheckNotExist(src) {
		code, err := qr.Encode(q.Url, q.Level, q.Mode)
		if err != nil {
			return "", "", err
		}

		code, err = barcode.Scale(code, q.Width, q.Height)
		if err != nil {
			return "", "", err
		}

		f, err := file.MustOpen(name, path)
		if err != nil {
			return "", "", err
		}
		defer f.Close()

		err = jpeg.Encode(f, code, nil)
		if err != nil {
			return "", "", err
		}
	}

	return name, path, nil
}


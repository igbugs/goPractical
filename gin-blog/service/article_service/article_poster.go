package article_service

import (
	"gin-blog/pkg/qrcode"
	"gin-blog/pkg/file"
	"os"
	"image/jpeg"
	"image"
	"image/draw"
		"gin-blog/pkg/logging"
)

type ArticlePoster struct {
	PosterName string
	*Article
	Qr *qrcode.Qrcode
}

func NewArticlePoster(posterName string, article *Article, qr *qrcode.Qrcode) *ArticlePoster {
	return &ArticlePoster{
		PosterName: posterName,
		Article: article,
		Qr: qr,
	}
}

func GetPosterFlag() string {
	return "poster"
}

func (a *ArticlePoster) CheckMergedImage(path string) bool {
	if file.CheckNotExist(path + a.PosterName) {
		return false
	}

	return true
}

func (a *ArticlePoster) OpenMergedImage(path string) (*os.File, error) {
	f, err := file.MustOpen(a.PosterName, path)
	if err != nil {
		return nil, err
	}

	return f, nil
}

type Rect struct {
	Name string
	X0 int
	Y0 int
	X1 int
	Y1 int
}

type Pt struct {
	X int
	Y int
}

type ArticlePosterBg struct {
	Name string
	*ArticlePoster
	*Rect
	*Pt
}

func NewArticlePosterBg(name string, ap *ArticlePoster, rect *Rect, pt *Pt) *ArticlePosterBg {
	return &ArticlePosterBg{
		Name: name,
		ArticlePoster: ap,
		Rect: rect,
		Pt: pt,
	}
}

func (a *ArticlePosterBg) Generate() (string, string, error) {
	fullPath := qrcode.GetQrcodeFullPath()
	fileName, path, err := a.Qr.Encode(fullPath)
	if err != nil {
		return "", "", err
	}

	if !a.CheckMergedImage(path) {
		mergedF, err := a.OpenMergedImage(path)
		if err != nil {
			logging.Warn("mergedF open failed, err: %v", err)
			return "", "", err
		}
		defer mergedF.Close()

		bgF, err := file.MustOpen(a.Name, path)
		if err != nil {
			logging.Warn("bgF open failed, err: %v", err)
			return "", "", err
		}
		defer bgF.Close()

		qrF, err := file.MustOpen(fileName, path)
		if err != nil {
			logging.Warn("qrF open failed, err: %v", err)
			return "", "", err
		}
		defer qrF.Close()

		bgImage, err := jpeg.Decode(bgF)
		if err != nil {
			logging.Warn("bgImage Decode failed, err", err)
			return "", "", err
		}

		qrImage, err := jpeg.Decode(qrF)
		if err != nil {
			logging.Warn("qrImage Decode failed, err: %v", err)
			return "", "", err
		}

		jpg := image.NewRGBA(image.Rect(a.Rect.X0, a.Rect.Y0, a.Rect.X1, a.Rect.Y1))

		draw.Draw(jpg, jpg.Bounds(), bgImage, bgImage.Bounds().Min, draw.Over)
		draw.Draw(jpg, jpg.Bounds(), qrImage, qrImage.Bounds().Min.Sub(image.Pt(a.Pt.X, a.Pt.Y)), draw.Over)

		jpeg.Encode(mergedF, jpg, nil)
	}
	return fileName, path, nil
}
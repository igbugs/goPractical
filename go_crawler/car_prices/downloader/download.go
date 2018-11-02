package downloader

import (
	"io"
	"net/http"
	"logging"
	"go_crawler/car_prices/fake_agent"
	"github.com/axgle/mahonia"
)

func Get(url string) io.Reader {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		logging.Error("http.NewRequest err: %v", err)
	}

	req.Header.Add("User-Agent", fake.GetUserAgent())
	req.Header.Add("Referer", "https://car.autohome.com.cn")

	resq, err := client.Do(req)
	if resq == nil || err != nil {
		logging.Error("client.Do, resq:%v, err:%v", resq, err)
		return nil
	}

	mah := mahonia.NewDecoder("gbk")
	return mah.NewReader(resq.Body)
}
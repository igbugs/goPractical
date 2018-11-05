package downloader

import (
	"fmt"
	"net/url"
	"go_crawler/lagou_jobs/fake"
	"net/http"
	"strings"
	"logging"
	"go_crawler/lagou_jobs/pkg/uuid"
	"io/ioutil"
	"encoding/json"
)

var (
	jobsApiUrl = "https://www.lagou.com/jobs/positionAjax.json?city=%s&needAddtionalResult=false"
)

type ListResult struct {
	Success bool
	Msg     string
	Code    int
	Content Content
}

type Content struct {
	PageNo         int
	PageSize       int
	PositionResult PositionResult
}

type PositionResult struct {
	TotalCount int
	Result     []Result
}

type Result struct {
	City              string
	BusinessZones     []string
	CompanyFullName   string
	CompanyLabelList  []string
	CompanyShortName  string
	CompanySize       string
	CreateTime        string
	District          string
	Education         string
	FinanceStage      string
	FirstType         string
	IndustryField     string
	IndustryLables    []string
	JobNature         string
	Longitude         string
	Latitude          string
	PositionAdvantage string
	PositionId        int32
	PositionLables    []string
	PositionName      string
	Salary            string
	SecondType        string
	Stationname       string
	Subwayline        string
	Linestaion        string
	WorkYear          string
}

type jobService struct {
	City string
}

func NewJobService(city string) *jobService {
	return &jobService{City: city}
}

func (j *jobService) GetUrl() string {
	req := fmt.Sprintf(jobsApiUrl, j.City)
	u, _ := url.Parse(req)
	query := u.Query()
	u.RawQuery = query.Encode()

	return u.String()
}

func (j *jobService) GetJobs(pn int, kd string) (*ListResult, error) {
	//client := fake.ProxyAuth{License: "", SecretKey: ""}.GetProxyClient()
	client := http.Client{}
	// 构造POST 请求的 form-data 数据
	postReader := strings.NewReader(fmt.Sprintf("first=false&pn=%d&kd=%s", pn, kd))
	req, err := http.NewRequest("POST", j.GetUrl(), postReader)
	if err != nil {
		logging.Error("http.NewRequest err:%v", err)
		return nil, nil
	}

	//req.Header.Set("Proxy-Switch-Ip", "yes")

	req.Header.Add("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	req.Header.Add("Accept-Languag", "zh-CN,zh;q=0.9")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Content-Length", "25")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Add("Host", "www.lagou.com")
	req.Header.Add("Origin", "https://www.lagou.com")
	req.Header.Add("Referer", "https://www.lagou.com/jobs/list_golang?labelWords=&fromSearch=true&suginput=")
	req.Header.Add("User-Agent", fake.GetUserAgent())
	req.Header.Add("Cookie", "JSESSIONID=ABAAABAAAGGABCBDD7E20A1565F8EF040FBC5D3A2FC75A9; _ga=GA1.2.318521016.1540524890; user_trace_token=20181026113451-"+uuid.GetUUID()+"; LGUID=20181026113451-"+uuid.GetUUID()+"; index_location_city=%E5%85%A8%E5%9B%BD; Hm_lvt_4233e74dff0ae5bd0a3d81c6ccf756e6=1540524890,1541139077; _gid=GA1.2.1188536856.1541139077; TG-TRACK-CODE=index_search; Hm_lpvt_4233e74dff0ae5bd0a3d81c6ccf756e6=1541148082; LGRID=20181102164124-"+uuid.GetUUID()+"; SEARCH_ID=f6364459131c49af9ecb0682349544b7")

	resp, err := client.Do(req)
	if resp == nil || err != nil {
		logging.Error("request failed, resq: %v, err:%v", resp, err)
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logging.Error("read resp.Body failed, err: %v", err)
		return nil, err
	}

	var results ListResult
	err = json.Unmarshal([]byte(body), &results)
	if err != nil {
		logging.Error("json.Unmarshal failed, err: %v", err)
		return nil, err
	}

	return &results, nil
}

package express

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type baiduCE struct {
	ClientId     string
	ClientSecret string
}

func NewBaiCe(cid, csecret string) *baiduCE {
	return &baiduCE{
		ClientId:     cid,
		ClientSecret: csecret,
	}
}

func (b *baiduCE) AnalyzeAddr(str string) (BaiduAnalyzeResult, error) {
	var result BaiduAnalyzeResult
	url := "https://aip.baidubce.com/rpc/2.0/nlp/v1/address?access_token=" + b.GetAccessToken()
	payload := strings.NewReader(fmt.Sprintf(`{"text":"%s"}`, str))
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, payload)

	if err != nil {
		return result, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return result, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (b *baiduCE) GetAccessToken() string {
	url := "https://aip.baidubce.com/oauth/2.0/token"
	postData := fmt.Sprintf("grant_type=client_credentials&client_id=%s&client_secret=%s", b.ClientId, b.ClientSecret)
	resp, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(postData))
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	accessTokenObj := map[string]any{}
	_ = json.Unmarshal([]byte(body), &accessTokenObj)
	return accessTokenObj["access_token"].(string)
}

type BaiduAnalyzeResult struct {
	Lat          float64 `json:"lat"`
	Detail       string  `json:"detail"`
	Town         string  `json:"town"`
	Phonenum     string  `json:"phonenum"`
	CityCode     string  `json:"city_code"`
	Province     string  `json:"province"`
	Person       string  `json:"person"`
	Lng          float64 `json:"lng"`
	ProvinceCode string  `json:"province_code"`
	Text         string  `json:"text"`
	County       string  `json:"county"`
	City         string  `json:"city"`
	CountyCode   string  `json:"county_code"`
	TownCode     string  `json:"town_code"`
	LogId        int     `json:"log_id"`
}

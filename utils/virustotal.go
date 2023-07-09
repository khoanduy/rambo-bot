package utils

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/khoaji/rambo-bot/config"
)

type UrlAnalysisResp struct {
	Data struct {
		Attributes struct {
			Stats struct {
				Harmless   int `json:"harmless"`
				Malicious  int `json:"malicious"`
				Suspicious int `json:"suspicious"`
				Undetected int `json:"undetected"`
				Timeout    int `json:"timeout"`
			} `json:"stats"`
		} `json:"attributes"`
	} `json:"data"`
}

type ScanUrlResp struct {
	Data struct {
		Id string `json:"id"`
	} `json:"data"`
}

const vtApiUrl string = "https://www.virustotal.com/api/v3"

func IsSafeLink(link string) (bool, error) {
	scanResp, err := postScanUrl(link)
	if err != nil {
		return true, err
	}

	resp, err := getUrlAnalysis(scanResp.Data.Id)
	if err != nil {
		return true, err
	}

	if report, err := json.Marshal(resp); err != nil {
		log.Println(err)
	} else {
		log.Printf("Check link [%v], response: %v\n", link, string(report))
	}

	if resp.Data.Attributes.Stats.Malicious > 0 || resp.Data.Attributes.Stats.Suspicious > 0 {
		return false, nil
	}

	return true, nil
}

func postScanUrl(link string) (*ScanUrlResp, error) {
	payload := strings.NewReader("url=" + url.QueryEscape(link))
	req, err := http.NewRequest("POST", vtApiUrl+"/urls", payload)
	if err != nil {
		return nil, err
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("x-apikey", config.Env.VtApiKey)
	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var resp *ScanUrlResp
	if err = json.Unmarshal(body, &resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func getUrlAnalysis(id string) (*UrlAnalysisResp, error) {
	req, err := http.NewRequest("GET", vtApiUrl+"/analyses/"+id, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("x-apikey", config.Env.VtApiKey)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var resp *UrlAnalysisResp
	if err = json.Unmarshal(body, &resp); err != nil {
		return nil, err
	}

	return resp, nil
}

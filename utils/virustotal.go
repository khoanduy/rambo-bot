package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/samothrakii/rambo-bot/conf"
)

type VtUrlAnalysisResponse struct {
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

type VtScanUrlResponse struct {
	Data struct {
		Id string `json:"id"`
	} `json:"data"`
}

func CheckUnsafeLink(link string) (bool, error) {
	scanRes, err := vtScanUrl(link)
	if err != nil {
		return true, err
	}

	analysisRes, err := getVtUrlAnalysis(scanRes.Data.Id)
	if err != nil {
		return true, err
	}

	resText, _ := json.Marshal(analysisRes)
	log.Printf("Check link %v, response: %v\n", link, string(resText))

	if analysisRes.Data.Attributes.Stats.Malicious > 0 || analysisRes.Data.Attributes.Stats.Suspicious > 0 {
		return true, nil
	}

	return false, nil
}

func vtScanUrl(link string) (VtScanUrlResponse, error) {
	var response VtScanUrlResponse
	scanUrl := "https://www.virustotal.com/api/v3/urls"

	payload := strings.NewReader("url=" + url.QueryEscape(link))
	req, err := http.NewRequest("POST", scanUrl, payload)
	if err != nil {
		return response, fmt.Errorf("Could not create new request %v", err)
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("x-apikey", conf.BotConf.VirusTotalApiKey)
	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return response, fmt.Errorf("Request failed %v", err)
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	if err != nil {
		return response, fmt.Errorf("Could not read response body %v", err)
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return response, fmt.Errorf("Error parsing JSON: %v", err)
	}

	return response, nil
}

func getVtUrlAnalysis(id string) (VtUrlAnalysisResponse, error) {
	var response VtUrlAnalysisResponse
	analysesUrl := "https://www.virustotal.com/api/v3/analyses/" + id

	req, err := http.NewRequest("GET", analysesUrl, nil)
	if err != nil {
		return response, fmt.Errorf("Could not create new request %v", err)
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("x-apikey", conf.BotConf.VirusTotalApiKey)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return response, fmt.Errorf("Request failed %v", err)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return response, fmt.Errorf("Could not read response body %v", err)
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return response, fmt.Errorf("Error parsing JSON: %v", err)
	}

	return response, nil
}

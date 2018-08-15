package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// setWeatherAPIRequest 發送天氣API request
func setWeatherAPIRequest() []byte {

	spaceClient1 := http.Client{
		Timeout: time.Second * 5,
	}

	return apiUtil("GET", "http://weather.json.tw/api", "", nil, &spaceClient1)
}

// apiUtil API 執行單元
func apiUtil(method string, apiURL string, host string,
	parameter io.Reader, spaceClient *http.Client) []byte {

	// Step 1. 組裝要發的內容
	req, err := http.NewRequest(method, apiURL, parameter)

	if err != nil {
		log.Println(err)
	}

	// 視情況選擇設Host
	req.Host = host

	// Step 2. 實際執行
	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		fmt.Println(getErr)
	}

	// Step 3. 接收response
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		fmt.Println(readErr)
	}
	// 若沒有錯誤則關閉連線，避免memory leak
	defer res.Body.Close()

	return body
}

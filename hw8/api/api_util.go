package api

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func SetAPIRequest() {
	log.Println("hello main~")

	spaceClient1 := http.Client{
		Timeout: time.Second * 5,
	}
	spaceClient2 := http.Client{
		Timeout: time.Second * 5,
	}

	apiUtil("GET", "http://127.0.0.1:8080/get-hello1", "", nil, &spaceClient1)
	apiUtil("GET", "http://127.0.0.1:8080/get-hello2", "", nil, &spaceClient2)
}

func apiUtil(method string, apiURL string, host string,
	parameter io.Reader, spaceClient *http.Client) {

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

	var response string

	json.Unmarshal(body, &response)

	fmt.Printf("%s\n", string(body[:]))
}

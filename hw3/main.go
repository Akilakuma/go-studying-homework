package main

import (
	"encoding/json"
	"fmt"
)

type apiResponse struct {
	Ret []struct {
		ID               int    `json:"id"`
		ParentID         int    `json:"parent_id"`
		Parent           int    `json:"parent"`
		AllParents       []int  `json:"all_parents"`
		Domain           int    `json:"domain"`
		LastLogin        string `json:"last_login"`
		Currency         string `json:"currency"`
		PasswordExpireAt string `json:"password_expire_at"`
		PasswordReset    bool   `json:"password_reset"`
		LastBank         int    `json:"last_bank"`
		Username         string `json:"username"`
		Block            bool   `json:"block"`
		Bankrupt         bool   `json:"bankrupt"`
		Cash             struct {
			ID          int     `json:"id"`
			UserID      int     `json:"user_id"`
			Balance     float64 `json:"balance"`
			PreSub      int     `json:"pre_sub"`
			PreAdd      int     `json:"pre_add"`
			Currency    string  `json:"currency"`
			LastEntryAt string  `json:"last_entry_at"`
		} `json:"cash"`
		CashFake    interface{} `json:"cash_fake"`
		Card        interface{} `json:"card"`
		EnabledCard interface{} `json:"enabled_card"`
	} `json:"ret"`
	Result  string `json:"result"`
	Profile struct {
		ExecutionTime int    `json:"execution_time"`
		ServerName    string `json:"server_name"`
	} `json:"profile"`
}

type internetDetail struct {
	IP   string
	Host string
	Port int
}

func main() {
	Q1()
	Q2()
}

func Q1() {
	var response apiResponse
	apiData := getApiResult()
	json.Unmarshal(apiData, &response)
	fmt.Printf("Q1\n")
	fmt.Printf("id: ")
	fmt.Println(response.Ret[0].ID)
	fmt.Printf("domain: ")
	fmt.Println(response.Ret[0].Domain)
	fmt.Printf("username: ")
	fmt.Println(response.Ret[0].Username)
	fmt.Println("cash: ")
	fmt.Println(response.Ret[0].Cash)
}

func Q2() {
	/* Q2-1
	IP:192.17.55.3	Host:docs.google.com 		Port:80
	IP:192.17.33.17 Host:ts-phpadmin.vir999.com Port:80
	IP:192.17.99.52 Host:jsonviewer.stack.hu 	Port:7711
	*/

	var internet []internetDetail
	var temp internetDetail

	temp = internetDetail{
		"192.17.55.3", "docs.google.com", 80,
	}
	internet = append(internet, temp)

	temp = internetDetail{
		"192.17.33.17", "ts-phpadmin.vir999.com", 80,
	}
	internet = append(internet, temp)

	temp = internetDetail{
		"192.17.99.52", "jsonviewer.stack.hu", 7711,
	}
	internet = append(internet, temp)

	fmt.Printf("\n Q2-1 資料表內容 \n:")
	fmt.Println(internet)

	/* Q2-2
	IP:177.18.2.33 Host:github.com Port:80
	*/
	temp = internetDetail{
		"177.18.2.33", "Host:github.com ", 80,
	}
	internet = append(internet, temp)
	fmt.Printf("\n Q2-2 資料表內容 \n:")
	fmt.Println(internet)

	// Q2-3
	updateSetting(&internet)
	fmt.Printf("\n Q2-3 資料表內容 \n:")
	fmt.Println(internet)

}

// updateSetting
func updateSetting(internet *[]internetDetail) {
	a := *internet
	for k, v := range a {
		if v.Host == "jsonviewer.stack.hu" {
			a[k].Port = 80
		}
	}
}

func getApiResult() []byte {
	userData := `{"ret":[{"id":148500286,"parent_id":1414663,"parent":1414663,"all_parents":[1414663,59735,58971,24367,6],"domain":6,"last_login":"2018-06-12T09:17:51+0800","currency":"CNY","password_expire_at":"2020-09-04T18:19:02+0800","password_reset":false,"last_bank":54,"username":"wesley01","block":false,"bankrupt":false,"cash":{"id":71814701,"user_id":148500286,"balance":4923642.0545,"pre_sub":0,"pre_add":0,"currency":"CNY","last_entry_at":"20180612155900"},"cash_fake":null,"card":null,"enabled_card":null}],"result":"ok","profile":{"execution_time":14,"server_name":"ipl-web01.rd5.qa"}}`

	return []byte(userData)
}

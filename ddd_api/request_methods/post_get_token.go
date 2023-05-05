package request_methods

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type getTokenBody struct {
	AppCode   uint64 `json:"appCode"`
	AppSecret string `json:"appSecret"`
}

type data struct {
	token       string `json:"token"`
	expiredTime uint32 `json:"expiredTime"`
}

/*
type User struct {
	Name        string
	Age         int
	Active      bool
	lastLoginAt string
}
*/

func PostGetToken(appCode uint64, appSecret string) {
	/*
		u, err := json.Marshal(User{Name: "Bob", Age: 10, Active: true, lastLoginAt: "today"})
		if err != nil {
			panic(err)
		}
		fmt.Println(string(u)) // {"Name":"Bob","Age":10,"Active":true}
	*/

	url := "https://open.dingdanll.com/open/platform/pub/getToken"

	json_body := getTokenBody{AppCode: appCode, AppSecret: appSecret}

	jsonStr, _ := json.Marshal(json_body)
	//temp := string(jsonStr)
	//fmt.Println(temp)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		fmt.Println("Error:", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
	}
	defer resp.Body.Close()

	// Read the response body
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	//fmt.Println("Response status:", resp.Status)
	//fmt.Println("resp", resp)
	//fmt.Println("resp.body", responseBody)
	fmt.Println("resp.body", string(responseBody))
}

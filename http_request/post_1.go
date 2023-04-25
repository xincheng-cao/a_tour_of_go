package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func Request_post_method1() {
	url := "http://10.200.0.43:14000/albums"

	jsonStr := []byte(`
    {
        "id": "4",
        "title": "Chao Ge and his team",
        "artist": "Cong Ge",
        "price": 392.11
    }
	`)

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

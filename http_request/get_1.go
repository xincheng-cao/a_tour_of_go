package main

import (
	"fmt"
	"io"
	"net/http"
)

func Request_get_method1() {
	resp, err := http.Get("http://10.200.0.43:14000/albums")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(responseBody))
}

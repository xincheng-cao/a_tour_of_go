package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func request_get_method2(id int) {
	resp, err := http.Get("http://10.200.0.43:14000/albums/" + strconv.Itoa(id))
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

package main

import (
	"ddd_api/request_methods"
	"fmt"
	"os"
	"strconv"
)

func main() {

	APPCODE := os.Getenv("APPCODE")
	APPCODE_uint64, _ := strconv.ParseUint(APPCODE, 10, 64)
	APPSECRET := os.Getenv("APPSECRET")
	fmt.Println(APPCODE, APPCODE_uint64, APPSECRET)
	request_methods.PostGetToken(APPCODE_uint64, APPSECRET)
}

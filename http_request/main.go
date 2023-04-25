package main

func main() {
	/*
		https://stackoverflow.com/questions/28081486/how-can-i-go-run-a-project-with-multiple-files-in-the-main-package

		go run .

		in goland
		configuration:
			run kind :
				1. package : a_tour_of_go/http_request
				2. directory: /data1/shiba/a_tour_of_go/http_request
				3. file(not work!): /data1/shiba/a_tour_of_go/web-service-gin/main.go
	*/

	Request_get_method1()
	Request_post_method1()
	Request_get_method1()
	request_get_method2(4)
}

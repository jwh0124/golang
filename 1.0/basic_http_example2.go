package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type helloWorldResponse struct {
	Message string `json:"message"`
	// 필드를 출력하지 않는다.
	Author string `json:"-"`
	// 값이 비어 있으면 필드를 출력하지 않는다.
	Date string `json:",omitempty"`
	// 출력을 문자열로 변환하고 이름을 "id"로 바꾼다.
	//Id int `json:"id,string"`
}

func main() {
	port := 8080

	http.HandleFunc("/helloworld", helloWorldHandler)

	log.Printf("Server starting on port %v\n", 8080)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	response := helloWorldResponse{Message: "HelloWorld"}
	encoder := json.NewEncoder(w)
	encoder.Encode(&response)
}

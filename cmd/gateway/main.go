package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/apex/gateway"
	"github.com/go-resty/resty/v2"
)

const (
	_api    = "https://api.openai.com/v1/engines/chatgpt-4/completions"
	_apiKey = "OPENAI_API_KEY"
)

type Request struct {
	Prompt string `json:"prompt"`
	Length int    `json:"length"`
}

type Response struct {
	Completion string `json:"completion"`
}

/*
在这个文件中，我们定义了一个 handler 函数来处理传入的 HTTP 请求。
该函数会解码请求体，并使用 go-resty/resty/v2 包发送 POST 请求到 ChatGPT4 API。
响应数据将被反序列化为 Response 结构体，并通过 HTTP 响应返回。
*/
func handler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var reqBody Request
	err := decoder.Decode(&reqBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	url := _api
	apiKey := os.Getenv(_apiKey)

	client := resty.New()
	// respBody := Response{}
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+apiKey).
		SetBody(reqBody).
		// SetResult(&respBody).
		Post(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp.Body()) //Encode(respBody)
}

var (
	port = flag.Int("port", -1, "specify a port")
)

func main() {
	fmt.Println("server start")
	flag.Parse()

	http.HandleFunc("/api/chatgpt4", handler)
	listener := gateway.ListenAndServe
	portStr := "n/a"

	if *port != -1 {
		portStr = fmt.Sprintf(":%d", *port)
		listener = http.ListenAndServe
		http.Handle("/", http.FileServer(http.Dir("./public")))
	}

	log.Fatal(listener(portStr, nil))
}

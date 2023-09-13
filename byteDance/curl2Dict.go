package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type DictRequest struct {
	A string `json:"trans_type"`
	B string `json:"source"`
	C string `json:"user_id"`
}
type DictResponse struct {
	Rc   int `json:"rc"`
	Wiki struct {
	} `json:"wiki"`
	Dictionary struct {
		Prons struct {
			EnUs string `json:"en-us"`
			En   string `json:"en"`
		} `json:"prons"`
		Explanations []string      `json:"explanations"`
		Synonym      []string      `json:"synonym"`
		Antonym      []string      `json:"antonym"`
		WqxExample   [][]string    `json:"wqx_example"`
		Entry        string        `json:"entry"`
		Type         string        `json:"type"`
		Related      []interface{} `json:"related"`
		Source       string        `json:"source"`
	} `json:"dictionary"`
}

func main() {
	client := &http.Client{}
	data := DictRequest{A: "en2zh"}
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("请输入要翻译的文字：")
	scanner.Scan()
	data.B = scanner.Text()
	buf, _ := json.Marshal(data)
	request := bytes.NewReader(buf)
	req, err := http.NewRequest("POST", "https://api.interpreter.caiyunai.com/v1/dict", request)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("authority", "api.interpreter.caiyunai.com")
	req.Header.Set("sec-ch-ua", `" Not A;Brand";v="99", "Chromium";v="8"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.131 Safari/537.36 SLBrowser/8.0.1.4031 SLBChan/103")
	req.Header.Set("app-name", "xy")
	req.Header.Set("content-type", "application/json;charset=UTF-8")
	req.Header.Set("accept", "application/json, text/plain, */*")
	req.Header.Set("device-id", "4ae5fdcc0ab269714b499ea9a0d68a85")
	req.Header.Set("os-type", "web")
	req.Header.Set("x-authorization", "token:qgemv4jr1y38jyq6vhvi")
	req.Header.Set("origin", "https://fanyi.caiyunapp.com")
	req.Header.Set("sec-fetch-site", "cross-site")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("referer", "https://fanyi.caiyunapp.com/")
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		log.Fatal("bad statusCode: ", resp.StatusCode, "body", bodyText)
	}
	rsp := &DictResponse{}
	_ = json.Unmarshal(bodyText, rsp)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data.B, "UK:", rsp.Dictionary.Prons.En, "US:", rsp.Dictionary.Prons.EnUs)
	for _, item := range rsp.Dictionary.Explanations {
		fmt.Println(item)
	}
}

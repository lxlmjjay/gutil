package util

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetRequest(url string) (bytedata []byte) {
	client := &http.Client{}
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println("http.Do failed,[err=%s][url=%s]", err)
		return
	}
	defer resp.Body.Close()
	bytedata, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("http.Do failed,[err=%s][url=%s]", err)
	}
	return
}
func GetRequestUrlEncoded(url string) (bytedata []byte) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	fmt.Println(resp.StatusCode)
	if resp.StatusCode == 200 {
		return body
	}
	return
}

func PostRequest(url string, data string) (bytedata []byte) {
	client := &http.Client{}
	resp, err := client.Post(url, "application/json", bytes.NewReader([]byte(data)))
	if err != nil {
		fmt.Println("http.Do failed,[err=%s][url=%s]", err)
		return
	}
	defer resp.Body.Close()
	bytedata, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("http.Do failed,[err=%s][url=%s]", err)
	}
	return
}

func PostRequestUrlEncoded(url string, data string) (bytedata []byte) {
	client := &http.Client{}
	resp, err := client.Post(url, "application/x-www-form-urlencoded;charset=utf-8", bytes.NewReader([]byte(data)))
	if err != nil {
		fmt.Println("http.Do failed,[err=%s][url=%s]", err)
		return
	}
	defer resp.Body.Close()
	bytedata, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("http.Do failed,[err=%s][url=%s]", err)
	}
	return
}
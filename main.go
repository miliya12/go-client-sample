package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	token := "token"
	device := "device"

	url := "https://httpbin.org/post"

	jsonStr := `{"token":"` + token + `","device":"` + device + `"}`

	req, err := http.NewRequest(
		"POST",
		url,
		bytes.NewBuffer([]byte(jsonStr)),
	)
	if err != nil {
		fmt.Println(err.Error())
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Basic hogehoge")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer resp.Body.Close()

	val, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(bytes.NewBuffer(val))
}

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	for i := 1; i < 2; i++ {
		client := http.Client{}
		req, _ := http.NewRequest("GET", "https://api.apiopen.top/getJoke?count=1&type=text", nil)
		resp, _ := client.Do(req)
		fmt.Printf("%s\n", resp)

		all, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(all)
	}
}

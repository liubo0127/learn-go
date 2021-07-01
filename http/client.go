package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {
	request, err := http.NewRequest(http.MethodGet, "http://album.zhenai.com/u/106794356", nil)
	//resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Println(err)
	}

	request.Header.Add("User-Agent",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) "+
			"AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1")

	client := http.Client{
		Transport: nil,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println("Redirect: ", req)
			return nil
		},
		Jar:     nil,
		Timeout: 0,
	}

	response, err := client.Do(request)

	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()

	bytes, err := httputil.DumpResponse(response, true)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%s\n", bytes)
}

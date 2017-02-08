package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var urls = []string{
	"http://www.baidu.com/",
	"http://golang.org/",
	"http://blog.golang.org/",
}

func tryHTTPReq() {
	// tryHTTPHead()
	tryHTTPGet()
}

func tryHTTPHead() {
	// Execute an HTTP HEAD request for all url's
	// and returns the HTTP status string or an error string.
	for _, url := range urls {
		resp, err := http.Head(url)
		if err != nil {
			fmt.Println("Error:", url, err)
		}
		fmt.Print(url, ": ", resp.Status)
	}
}

func tryHTTPGet() {
	res, err := http.Get("http://www.google.com")
	checkError(err)
	data, err := ioutil.ReadAll(res.Body)
	checkError(err)
	fmt.Printf("Got: %q", string(data))
}

func checkError(err error) {
	if err != nil {
		log.Fatalf("Get : %v", err)
	}
}

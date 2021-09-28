package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const defaultTag = "go"
const defaultApiEndpoint = "https://api.stackexchange.com/2.2/search/advanced?order=desc&sort=activity&site=stackoverflow"

func main() {
	flag.Parse()

	if flag.NArg() < 1 {
		flag.Usage()
		os.Exit(1)
	}

	question := strings.Join(flag.Args(), " ")

	performSearch(question)

	//for result := range results {
	//	//fmt.Println(result)
	//}
}

func performSearch(question string) chan string {
	res := make(chan string, 1)

	u, _ := url.Parse(defaultApiEndpoint)
	q := u.Query()
	q.Set("tagged", defaultTag)
	q.Set("title", question)
	u.RawQuery = q.Encode()

	c := http.Client{}

	r, _ := c.Get(u.String())
	b, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(b))

	// Perform the request
	go func() {
		res <- "answer"
		close(res)
	}()

	return res
}

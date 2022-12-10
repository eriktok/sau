package main

import (
	"bufio"
	"fmt"
	"log"
	"net/url"
	"os"
)

/*
Type strcut
-host
- url path
-file path
*/

type Data struct {
	host string
	path string
}
type Datas struct {
	d []Data
}

func main() {
	urls := make([]string, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		urls = append(urls, scanner.Text())
	}

	if scanner.Err() != nil {
		log.Fatalf("Could not find file")
	}

	pu := parseUrls(urls)
	for _, v := range pu {
		fmt.Println(v)
	}

}

func parseUrls(urls []string) map[string][]string {
	m := make(map[string][]string)
	for _, u := range urls {
		parsedUrls, _ := url.Parse(u)
		m[parsedUrls.Host] = append(m[parsedUrls.Host], parsedUrls.Host+parsedUrls.Path)
	}
	return m
}

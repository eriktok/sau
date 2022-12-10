package main

import (
	"bufio"
	"log"
	"net/url"
	"os"
)

func main() {
	urls := make([]string, 0)

	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		urls = append(urls, scanner.Text())
	}

	if scanner.Err() != nil {
		log.Fatalf("Could not find file")
	}

	pu := parseUrls(urls)
	for k, v := range pu {
		createDirectoryIfNotExists(currentDir + "/" + k)
		createFile(currentDir+"/"+k+"/"+k, v)
	}
}

func createDirectoryIfNotExists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return os.MkdirAll(path, os.ModeDir|0755)
	}
	return nil
}

func createFile(path string, text []string) {
	file, err := os.Create(path + ".txt")
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	datawriter := bufio.NewWriter(file)
	for _, data := range text {
		_, _ = datawriter.WriteString(data + "\n")
	}
	datawriter.Flush()
	file.Close()
}

func parseUrls(urls []string) map[string][]string {
	m := make(map[string][]string)
	for _, u := range urls {
		parsedUrls, _ := url.Parse(u)
		m[parsedUrls.Host] = append(m[parsedUrls.Host], parsedUrls.Host+parsedUrls.Path)
	}
	return m
}

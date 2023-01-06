package main

import (
	"bufio"
	"flag"
	"log"
	"net/url"
	"os"
)

var originalUrl = flag.Bool("c", true, "Parses urls to return host and path if equals to false")
var outputDir = flag.String("o", "/js", " output directory name")

func main() {
	flag.Parse()
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Could not get current working directory: %v", err)
	}
	urls := make([]string, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		urls = append(urls, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading from stdin: %v", err)
	}
	parsedUrls := parseUrls(urls)
	for dirName, sortedUrlsBySub := range parsedUrls {
		dirPath := currentDir + "/" + *outputDir + "/" + dirName
		filePath := dirPath + "/" + dirName + ".txt"
		if err := createDirectoryAndFile(dirPath, filePath, sortedUrlsBySub); err != nil {
			log.Fatalf("Error creating directory and file: %v", err)
		}
	}
}

func parseUrls(urls []string) map[string][]string {
	data := make(map[string][]string)
	for _, u := range urls {
		parsedUrls, _ := url.Parse(u)
		if *originalUrl {
			data[parsedUrls.Host] = append(data[parsedUrls.Host], parsedUrls.String())
		} else {
			data[parsedUrls.Host] = append(data[parsedUrls.Host], parsedUrls.Host+parsedUrls.Path)
		}
	}
	return data
}

func createDirectoryAndFile(dir, file string, text []string) error {
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	for _, data := range text {
		if _, err := w.WriteString(data + "\n"); err != nil {
			return err
		}
	}
	return w.Flush()
}

package main

import (
	"bufio"
	"flag"
	"log"
	"net/url"
	"os"
)

var myflag = flag.Bool("c", false, "print whole url")

func main() {
	urls := make([]string, 0)
	flag.Parse()
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
	for dirName, sortedUrlsBySub := range pu {
		createDirectoryIfNotExists(currentDir + "/" + dirName)
		createFile(currentDir+"/"+dirName+"/"+dirName, sortedUrlsBySub)
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
		if *myflag {
			m[parsedUrls.Host] = append(m[parsedUrls.Host], parsedUrls.String())
		} else {
			m[parsedUrls.Host] = append(m[parsedUrls.Host], parsedUrls.Host+parsedUrls.Path)
		}
	}
	return m
}

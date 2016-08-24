package main

import (
	"flag"
	"fmt"
	"os"
	"sync"

	"github.com/galihrivanto/latihan/util"
)

var (
	rootUrl string
	domain  string
)

func main() {
	flag.StringVar(&rootUrl, "url", "", "root url")
	flag.StringVar(&domain, "domain", "", "main domain")
	flag.Parse()

	if rootUrl == "" {
		flag.Usage()
		os.Exit(-1)
	}

	// #1 step. fetch root child urls
	urls, err := util.GetChildURLs(rootUrl, domain)
	if err != nil {
		fmt.Println("failed to fetch url", rootUrl, "error:", err)
		return
	}

	// #2 step. fetch child urls
	// while prevent same url fetched twice
	urls = util.RemoveDuplicate(urls)

	clist := make(chan string)
	defer close(clist)

	wg := sync.WaitGroup{}
	for _, url := range urls {
		wg.Add(1)
		go func(u string, clist chan string, w *sync.WaitGroup) {
			defer w.Done()

			uurls, err := util.GetChildURLs(u)
			if err != nil {
				fmt.Println("failed to fetch", u)
			}

			fmt.Printf("url %s have %d child urls\n", u, len(uurls))
			for _, uu := range uurls {
				clist <- uu
			}

		}(url, clist, &wg)

	}

	// result aggregator
	result := []string{}
	go func() {
		for url := range clist {
			result = append(result, url)
		}
	}()

	wg.Wait()

	// remove dups
	result = util.RemoveDuplicate(result)

	for i, url := range result {
		fmt.Printf("%d - %s\n", i, url)
	}

	fmt.Println("total:", len(result))

}

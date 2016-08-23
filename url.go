package utils

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

// getChildURLs ambil child links dari alamat yg  diberikan
func getChildURLs(url string, domain ...string) ([]string, error) {

	// setup http client
	// pakai "GET" request
	client := &http.Client{}
	rb := &bytes.Buffer{}
	req, err := http.NewRequest("GET", url, rb)
	if err != nil {
		return nil, err
	}

	// pastikan response body close
	// dan "read all"
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// ambil pattern dari html anchor
	pattern, err := regexp.Compile(`<a\s+(?:[^>]*?\s+)?href="([^"]*)"`)
	if err != nil {
		return nil, err
	}

	urls := []string{}
	found := pattern.FindAllStringSubmatch(string(body), -1)

	// tidak ada link, tidak usah diteruskan
	if len(found) == 0 {
		return urls, nil
	}

	for _, f := range found {
		if len(f) == 2 {

			// jika parameter ada parameter domain
			u := f[1]
			if len(domain) > 0 {
				if strings.Contains(u, domain[0]) {
					urls = append(urls, u)
				}
			} else {
				urls = append(urls, u)
			}

		}
	}

	return urls, nil
}

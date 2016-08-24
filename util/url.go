package util

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

// getChildURLs ambil child links dari alamat yg  diberikan
func GetChildURLs(url string, domain ...string) ([]string, error) {

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

	p := `<a\s+(?:[^>]*?\s+)?href="([^"]*)"`
	if len(domain) > 0 && domain[0] != "" {
		p = fmt.Sprintf(`<a\s+(?:[^>]*?\s+)?href="((?:http|https)://[^"]*\.%s[^"]*)"`, domain[0])
	}

	// ambil pattern dari html anchor
	pattern, err := regexp.Compile(p)
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
			urls = append(urls, f[1])

		}
	}

	return urls, nil
}

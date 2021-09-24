/*
Copyright © 2021 Pierre Galvez <dev@pierre-galvez.fr>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package find

import (
	"bytes"
	"crypto/tls"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"runtime"
	"sync"
	"time"
)

//Execute receive a list of URLs and return a new list with the URLs that has Source Code Disclosure
func Execute(urls []url.URL) (result []url.URL) {
	// Limit the number of spare OS threads to just 3
	runtime.GOMAXPROCS(3)

	var wg sync.WaitGroup
	for _, u := range urls {
		wg.Add(1)
		go func(url url.URL) {
			defer wg.Done()
			if findGit(url) {
				result = append(result, url)
				fmt.Println(formatURL(url), "GIT")
				return
			}
			if findSvn(url) {
				result = append(result, url)
				fmt.Println(formatURL(url), "SVN")
			}
		}(u)
	}
	wg.Wait()
	return result
}

func doRequest(u url.URL, filePattern string) (resp *http.Response, body []byte, err error) {
	tr := &http.Transport{
		TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
		ResponseHeaderTimeout: 5 * time.Second,
	}
	client := &http.Client{
		Transport: tr,
	}

	uri := formatURL(u) + filePattern
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return resp, body, err
	}

	resp, err = client.Do(req)
	if err != nil {
		return resp, body, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return resp, body, errors.New("Http status not in 200")
	}

	body, err = ioutil.ReadAll(resp.Body)
	return resp, body, err
}

func checkPattern(body []byte, pattern string) bool {
	return bytes.Contains(body, []byte(pattern))
}

func formatURL(u url.URL) string {
	return u.Scheme + "://" + u.Host
}

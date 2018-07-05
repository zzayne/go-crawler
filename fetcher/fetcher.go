package fetcher

import (
	"bufio"
	"io/ioutil"
	"net/http"
)

// Fetch ...
func Fetch(url string) (content []byte, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	reader := bufio.NewReader(resp.Body)

	content, err = ioutil.ReadAll(reader)

	return content, err
}

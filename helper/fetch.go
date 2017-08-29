package helper

import (
	"io/ioutil"
	"net/http"
)

func Fetch(url string) ([]byte, error) {
	c := &http.Client{}
	res, err := c.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return b, err
}

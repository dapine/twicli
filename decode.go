package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func FromResponseJSON(res *http.Response, v interface{}) error {
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	json.Unmarshal(data, v)

	return nil
}

package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func FetchWebsite(w Website) (string, error) {
	log.Printf("fetchWebsite(%s)\n", w.URL)
	resp, err := http.Get(w.URL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	log.Printf("Return result for %s", w.URL)
	return string(body), nil
}

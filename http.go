// Copyright (c) 2023 Julian Müller (ChaoticByte)

package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func httpGet(url string, headers []http.Header, timeout time.Duration) ([]byte, error) {
	data := []byte{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return data, err
	}
	for _, h := range headers {
		for k, v := range h {
			req.Header.Set(k, v[0])
		}
	}
	client := &http.Client{Timeout: timeout}
	resp, err := client.Do(req)
	if err != nil {
		return data, err
	}
	data, err = io.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		return data, fmt.Errorf("status code %v while fetching %v", resp.StatusCode, url)
	}
	return data, err
}

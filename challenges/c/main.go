package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

func request(url string) ([]byte, error) {
	var (
		client http.Client
		res    *http.Response
		bin    []byte
		err    error
	)
	client = http.Client{
		Timeout: time.Second,
	}
	res, err = client.Get(url)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = res.Body.Close()
	}()
	bin, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, errors.New("invalid status code in response")
	}
	return bin, nil
}

func bulkRequest(urls []string) []string {
	var (
		channel   chan []byte
		wg        sync.WaitGroup
		responses []string
	)
	channel = make(chan []byte)
	for _, url := range urls {
		wg.Add(1)
		go func(
			url string,
			channel chan []byte,
			wg *sync.WaitGroup,
		) {
			defer (*wg).Done()
			res, err := request(url)
			if err != nil {
				return
			}
			channel <- res
		}(url, channel, &wg)
	}

	go func() {
		wg.Wait()
		close(channel)
	}()

	for response := range channel {
		responses = append(responses, string(response))
	}
	return responses
}

func main() {
	urls := []string{
		"https://jsonplaceholder.typicode.com/todos/1",
		"https://jsonplaceholder.typicode.com/todos/2",
		"https://jsonplaceholder.typicode.com/todos/3",
		"https://jsonplaceholder.typicode.com/todos/4",
	}
	res := bulkRequest(urls)
	fmt.Println(res)
}

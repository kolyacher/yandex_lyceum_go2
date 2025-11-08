/*
Напишите функцию FetchAPI(ctx context.Context, urls []string, timeout time.Duration) []*APIResponse, которая одновременно (concurrently) получает данные из переданных urls (метод GET). Используйте контекст, чтобы ограничить время запроса и отмены ожидания свыше timeout. В случае ошибки верните её в соответствующем объекте APIResponse. При превышении таймаута ожидания должна быть ошибка context.DeadlineExceeded. В коде должна быть структура:

type APIResponse struct { URL string // запрошенный URL Data string // тело ответа StatusCode int // код ответа Err error // ошибка, если возникла }
*/
package main

import (
	"context"
	"io/ioutil"
	"net/http"
	"time"
)

type APIResponse struct {
	URL        string // запрошенный URL
	Data       string // тело ответа
	StatusCode int    // код ответа
	Err        error  // ошибка, если возникла
}

func FetchAPI(ctx context.Context, urls []string, timeout time.Duration) []*APIResponse {

	results := make([]*APIResponse, len(urls))

	type result struct {
		index    int
		response *APIResponse
	}
	resultCh := make(chan result)

	for i, url := range urls {
		go func(i int, url string) {
			resp := &APIResponse{URL: url}
			ctx, cancel := context.WithTimeout(ctx, timeout)
			defer cancel()

			req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
			if err != nil {
				resp.Err = err
				resultCh <- result{i, resp}
				return
			}

			client := &http.Client{}
			httpResp, err := client.Do(req)
			if err != nil {
				resp.Err = err
				resultCh <- result{i, resp}
				return
			}
			defer httpResp.Body.Close()

			body, err := ioutil.ReadAll(httpResp.Body)
			if err != nil {
				resp.Err = err
				resultCh <- result{i, resp}
				return
			}

			resp.Data = string(body)
			resp.StatusCode = httpResp.StatusCode
			resultCh <- result{i, resp}
		}(i, url)
	}

	for i := 0; i < len(urls); i++ {
		res := <-resultCh
		results[res.index] = res.response
	}

	return results
}

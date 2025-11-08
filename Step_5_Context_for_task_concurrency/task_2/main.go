/*
Напишите функцию fetchAPI(ctx context.Context, url string, timeout time.Duration) (*APIResponse, error), которая запрашивает данные по адресу url (метод GET) и возвращает код ответа и само тело ответа. Используйте контекст для ограничения времени запроса и отмены ожидания свыше timeout. В случае ошибок возвращайте nil, error. При превышении таймаута ожидания — nil, context.DeadlineExceeded. В коде должна быть структура:

type APIResponse struct { Data string // тело ответа StatusCode int // код ответа }
*/
package main

import (
	"context"
	"io/ioutil"
	"net/http"
	"time"
)

type APIResponse struct {
	Data       string
	StatusCode int
}

func fetchAPI(ctx context.Context, url string, timeout time.Duration) (*APIResponse, error) {

	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)

	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return &APIResponse{
		Data:       string(body),
		StatusCode: resp.StatusCode,
	}, nil
}

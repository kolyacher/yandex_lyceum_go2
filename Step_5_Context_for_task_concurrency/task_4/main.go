/*Напишите функцию StartServer(maxTimeout time.duration), которая запускает веб-сервер по адресу http://localhost:8080. При обращении по URL http://localhost:8080/readSource сервер должен сделать запрос по другому адресу: http://localhost:8081/provideData (код запуска сервера localhost:8081 писать не нужно) и вернуть полученные данные. Используйте http.timeoutHandler, чтобы ограничить время ожидания данных с сервера localhost:8081. При превышении лимита maxTimeout пользователю должна вернуться ошибка с кодом StatusServiceUnavailable, иначе — полученные данные.*/
package main

import (
	"io/ioutil"
	"net/http"
	"time"
)

func StartServer(maxTimeout time.Duration) {
	http.HandleFunc("/readSource", func(w http.ResponseWriter, r *http.Request) {
		client := http.Client{
			Timeout: maxTimeout,
		}
		resp, err := client.Get("http://localhost:8081/provideData")
		if err != nil {
			http.Error(w, "Service Unavailable", http.StatusServiceUnavailable)
			return
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			http.Error(w, "Service Unavailable", http.StatusServiceUnavailable)
			return
		}
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, "Service Unavailable", http.StatusServiceUnavailable)
			return
		}
		w.Write(data)
	})
	http.ListenAndServe(":8080", nil)
}

package main

import (
	"testing"
)

func TestGetTasks(t *testing.T) {
	chatHistory := `
TICKET-12345_Паша Попов_Готово_2024-01-01
TICKET-12346_Иван Иванов_В работе_2024-01-02
TICKET-12347_Анна Смирнова_Не будет сделано_2024-01-03
TICKET-12348_Паша Попов_В работе_2024-01-04
`

	// Тест 1: поиск всех задач Паши Попова
	user := "Паша Попов"
	tasks := GetTasks(chatHistory, &user, nil)
	if len(tasks) != 2 {
		t.Errorf("Ожидалось 2 задачи для Паши Попова, найдено %d", len(tasks))
	}
	for _, task := range tasks {
		if task.User != "Паша Попов" {
			t.Errorf("Найденная задача не принадлежит Паше Попову: %v", task)
		}
	}

	// Тест 2: поиск всех задач со статусом "В работе"
	stat := "В работе"
	workTasks := GetTasks(chatHistory, nil, &stat)
	if len(workTasks) != 2 {
		t.Errorf("Ожидалось 2 задачи со статусом 'В работе', найдено %d", len(workTasks))
	}
	for _, task := range workTasks {
		if task.Status != "В работе" {
			t.Errorf("Найденная задача имеет неверный статус: %v", task)
		}
	}
}

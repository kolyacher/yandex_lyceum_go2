/*
Вас взяли на работу проджект-менеджером. На проекте полный хаос (иначе зачем вы там нужны). Пора спасать положение.

Нужно наладить работу и понять, кто какие задачи делает, что уже сделано и в какие сроки будет выполнена оставшаяся работа.

Но задачи в трекере никто не завёл, а людей, которые в курсе, в команде не оказалось.

Вы осмотрелись и поняли, что в командном чате велась переписка. И, о чудо, коллеги вели её в форматированном варианте.

Каждое сообщение, которое относилось к прогрессу по задачам, начиналось с фразы TICKET и имело следующий формат:

TICKET-<номер тикета>_<исполнитель тикета>_<статус задачи>_<дата обновления>

Пример: TICKET-12345_Паша Попов_Готово_2024-01-01

Возможные статусы задачи: «Готово», «В работе», «Не будет сделано».

Вы получили доступ к истории чата (как хорошо, что интернет всё помнит). Теперь вы должны написать себе инструмент (хочешь сделать хорошо — сделай сам), чтобы понять, сколько задач в каком статусе у какого исполнителя.

Все сообщения — это один большой текст. Каждое новое сообщение — на новой строке.

Напишите функцию:

	type Ticket struct {
	    Ticket string
	    User   string
	    Status string
	    Date   time.Time
	}

func GetTasks(text string, user *string, status *string) []Ticket

чтобы найти все задачи указанного пользователя и статуса. Если передать nil вместо пользователя или статуса, задачи по этим полям не будут фильтроваться.

# Примеры Входная строка

TICKET-12345_Паша Попов_Готово_2024-01-01

# Структура данных

timeStamp, _ := time.Parse("2006-01-02", "2024-01-01")

	ticket := Ticket{
	    Ticket: "TICKET-12345",
	    User:   "Паша Попов",
	    Status: "Готово",
	    Date:   timeStamp,
	}

# Обработка ошибок Если строка не соответствует правилам - просто пропкскайте эту строчку

Входная строка 12345_Паша Попов_Готово_2024-01-01 ошибочная - не начинается с Ticket - пропускаем ее. Входная строка TICKET-12345_Паша Попов_Готово_самый лучший день ошибочная - самый лучший день не является датой.

# Примеры

Входные данные:

TICKET-12345_Паша Попов_Готово_2024-01-01
TICKET-12346_Иван Иванов_В работе_2024-01-02
TICKET-12347_Анна Смирнова_Не будет сделано_2024-01-03
TICKET-12348_Паша Попов_В работе_2024-01-04

Так хотим найти все задачи пользователя Паша Попов (обратите внимание, что поле status тут равно nil)

user := "Паша Попов"
tasks := GetTasks(chatHistory, &user, nil)

Так хотим найти все задачи со статусом В работе (обратите внимание, что поле user тут равно nil)

status := "В работе"
tasks := GetTasks(chatHistory, nil, &status)
*/
package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Ticket struct {
	Ticket string
	User   string
	Status string
	Date   time.Time
}

func GetTasks(text string, user *string, status *string) []Ticket {
	lines := strings.Split(text, "\n")
	var tickets []Ticket

	for _, line := range lines {
		if len(line) > 0 {
			parts := strings.Split(line, "_")

			if len(parts) != 4 {
				continue
			}

			date, err := time.Parse("2006-01-02", parts[3])
			if err != nil {
				continue
			}

			tiknom := strings.Split(parts[0], "-")
			_, err = strconv.ParseInt(tiknom[1], 10, 64)
			if strings.ToUpper(tiknom[0]) != "TICKET" && err != nil {
				continue
			}

			matchUser := true
			if user != nil && *user != "" {
				matchUser = (*user == parts[1])
			}

			matchStatus := true
			if status != nil && *status != "" {
				validStatuses := map[string]bool{"Готово": true, "В работе": true, "Не будет сделано": true}
				if !validStatuses[*status] || *status != parts[2] {
					matchStatus = false
				}
			}

			if matchUser && matchStatus {
				tickets = append(tickets, Ticket{
					Ticket: parts[0],
					User:   parts[1],
					Status: parts[2],
					Date:   date,
				})
			}
		}
	}

	return tickets
}

func main() {
	text := `TICKET-12345_Паша Попов_Готово_2024-01-01
TICKET-12346_Иван Иванов_В работе_2024-01-02
TICKET-12347_Анна Смирнова_Не будет сделано_2024-01-03
TICKET-12348_Паша Попов_В работе_2024-01-04`

	user := "Иван Иванов"
	status := "В работе"
	tasks := GetTasks(text, &user, &status)

	for _, task := range tasks {
		fmt.Printf("Задача: %s, Пользователь: %s, Статус: %s, Дата: %v\n",
			task.Ticket, task.User, task.Status, task.Date.Format("2006-01-02"))
	}
}

/*
 --- FAIL: TestGetTasks (0.00s)
    source_test.go:143: Ожидалось 5 задач для текста с "мусорными" данными, найдено 4
    source_test.go:261: Ожидалось 2 задачи для Ивана Иванова со статусом 'В работе', найдено 0
    source_test.go:273: Ожидалось 2 задачи для Ивана Иванова с разными статусами, найдено 0
FAIL
*/

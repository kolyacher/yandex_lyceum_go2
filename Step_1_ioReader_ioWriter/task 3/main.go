/*
Создайте структуру UpperWriter с полем UpperString string

	type UpperWriter struct {
	    UpperString string
	}

и реализуйте интерфейс io.Writer. Метод Write должен переводить строку в верхний регистр и записывать данные в поле UpperString. Если возникает ошибка, функция должна возвращать её.
*/
package main

import "strings"

type UpperWriter struct {
	UpperString string
}

func (uw UpperWriter) Write(buf []byte) (int, error) {
	uw.UpperString = strings.ToUpper(string(buf))
	return len(buf), nil
}

/*
Напишите функцию
func CopyFilePart(inputFilename, outFileName string, startpos int) error
которая открывает файл с именем inputFilename на чтение, создаёт файл с именем outFileName и записывает содержимое файла inputFilename с позиции startPos и до конца в файл outFileName. Если возникнет ошибка, верните её. Если все операции прошли без ошибок, верните nil.
Не забудьте закрыть файлы после обработки.
*/
package main

import (
	"io"
	"os"
)

func CopyFilePart(inputFilename, outFileName string, startpos int) error {
	f, err := os.Open(inputFilename)
	if err != nil {
		return err
	}
	fout, err := os.Create(outFileName)
	if err != nil {
		return err
	}
	f.Seek(int64(startpos), 0)
	buffer := make([]byte, 100)
	for {

		n, err := f.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil
		}
		_, err = fout.Write(buffer[:n])
		if err != nil {
			return err
		}
	}
	f.Close()
	fout.Close()
	return nil
}

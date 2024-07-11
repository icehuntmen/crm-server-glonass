package middlewares

import (
	"bufio"
	"fmt"
	"os"
)

var Version string

func init() {
	file, err := os.Open("./VERSION")
	if err != nil {
		fmt.Println("Ошибка при открытии файла VERSION:", err)
		return
	}
	defer file.Close()

	// Чтение первой строки из файла
	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		Version = scanner.Text() // Присваиваем считанную версию переменной Version
	} else {
		fmt.Println("Файл VERSION пуст или пусты строка")
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка при чтении файла VERSION:", err)
		return
	}
}

func GetVersion() string {
	return Version
}

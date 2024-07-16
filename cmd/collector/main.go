package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func main() {
	// URL, куда будем отправлять POST-запрос
	url := "https://hosting.glonasssoft.ru/api/v3/auth/login"

	// Тело POST-запроса в формате JSON
	jsonBody := []byte(`{
  		"login": "demo",
  		"password": "%%demo%%"
	}`)

	// Создаем новый запрос типа POST
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Устанавливаем заголовок Content-Type для указания типа содержимого
	req.Header.Set("Content-Type", "application/json")

	// Создаем HTTP-клиент
	client := &http.Client{}

	// Отправляем запрос и получаем ответ
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Выводим статус код ответа
	fmt.Println("Response Status:", resp.Status)
}

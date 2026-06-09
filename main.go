package main

import (
	"log"
	"net/http"
)

func main() {
	// Создаем свой ServeMux - явный и изолированный, а не используем глобальный. Пустой маршрутизатор
	mux := http.NewServeMux()

	// Регистрируем handler на корневой путь "/"
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// w.Write отправляет байты в тело ответа
		w.Write([]byte("kafka_exporter is active\n"))
	})

	// Описываем сервер, слушаем на порту 9308, используем наш mux.
	server := &http.Server{
		Addr:    ":9308",
		Handler: mux,
	}

	log.Println("Listening on :9308")

	// ListenAndServe блокирует выполнение, пока сервер работает. При возврате ошибки (например, порт занят) - логируем.
	log.Fatal(server.ListenAndServe())
}
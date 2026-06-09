package main

import (
	"log"
	"net/http"
)

const (
	listenAddr = ":9308"
	metricsPath = "/metrics"
)

func main() {
	setup(listenAddr, metricsPath)
}

func setup(listenAddr, metricsPath string) {
	// Создаем свой ServeMux - явный и изолированный, а не используем глобальный. Пустой маршрутизатор
	mux := http.NewServeMux()

	// Регистрируем handler на корневой путь "/"
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("kafka_exporter is active\n"))
	})

	// эндпоинт /healthz - это проба, сервис живой. Например, k8s проверяет - жив ли процесс. Пока просто отвечаем "ok"
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	// эндпоинт /metrics - здесь будут метрики. Пока заглушка в формате Prometheus.
	mux.HandleFunc(metricsPath, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; version=0.0.4")
		w.Write([]byte("# HELP my_exporter_up Is the exporter running.\n"))
		w.Write([]byte("# TYPE my_exporter_up gauge\n"))
		w.Write([]byte("my_exporter_up 1\n"))
	})


	// Описываем сервер, слушаем на порту 9308, используем наш mux.
	server := &http.Server{
		Addr:    listenAddr,
		Handler: mux,
	}

	log.Println("Listening on :9308")

	// ListenAndServe блокирует выполнение, пока сервер работает. При возврате ошибки (например, порт занят) - логируем.
	log.Fatal(server.ListenAndServe())
}
package main

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	// Регистрируем наш Collector. Теперь на каждый scrape Prometheus будет вызызать e.Collect()
	exporter := NewExporter()
	prometheus.MustRegister(exporter)

	setup(":9308", "/metrics")
}

func setup(listenAddr, metricsPath string) {
	// Создаем свой ServeMux - явный и изолированный, а не используем глобальный. Пустой маршрутизатор
	mux := http.NewServeMux()

	// эндпоинт /healthz - это проба, сервис живой. Например, k8s проверяет - жив ли процесс. Пока просто отвечаем "ok"
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	// эндпоинт /metrics теперь обслуживает promhttp. Он сам читает реестр по умолчанию и форматирует ответ.
	mux.Handle(metricsPath, promhttp.Handler())

	// Описываем сервер, слушаем на порту 9308, используем наш mux.
	server := &http.Server{
		Addr:    listenAddr,
		Handler: mux,
	}

	log.Println("Listening on :9308")
	log.Fatal(server.ListenAndServe())
}
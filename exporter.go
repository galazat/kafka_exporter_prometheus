package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

const namespace = "kafka" // для имени метрик

// Это наши описания метрик - Descriptors. Мы их готовим заранее и переиспользуем
var (
	clusterBrokers = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "brokers"),
		"Number of brokers in the Kafka cluster.",
		nil, nil, // нет переменных лейблов, нет постоянных
	)
	topicPartitions = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "topic", "partitions"),
		"Number of partitions for this topic.",
		[]string{"topic"}, // один переменный лейбл: topic
		nil,
	)
)

// Exporter реализует интерфейс prometheus.Collector
type Exporter struct {
	// Сюда позже добавим клиент Kafka. Пока оставим пустым
}

func NewExporter() *Exporter {
	return &Exporter{}
}

// Describe: перечисляем все метрики, которые умеем отдавать
func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- clusterBrokers
	ch <- topicPartitions
}

// Collect: вызывается на каждый scrape. Происходит сбор данных
func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	// Пока тестовые данные. В Части 2 заменим на реальные из Kafka

	// Метрика без лейблов: число брокеров = 3
	ch <- prometheus.MustNewConstMetric(
		clusterBrokers, prometheus.GaugeValue, float64(3),
	)

	// Метрики с лейблом topic. Эмулируем два топика
	fakeTopics := map[string]int{"orders": 6, "payments": 3}
	for topic, partitions := range fakeTopics {
		ch <- prometheus.MustNewConstMetric(
			topicPartitions,
			prometheus.GaugeValue,
			float64(partitions),
			topic, // значение лейбла "topic"
		)
	}
}
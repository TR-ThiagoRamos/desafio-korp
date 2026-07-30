package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)
// Registra o contador de requisições HTTP para o Prometheus
var (
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Volume total de requisicoes HTTP recebidas",
		},
		[]string{"path"},
	)
)
// Registra as métricas na inicialização
func init() {
	prometheus.MustRegister(httpRequestsTotal)
}
// Estrutura JSON exigido pelo desafio
type Response struct {
	Nome    string `json:"nome"`
	Horario string `json:"horario"`
}
// Handler do caminho /projeto-korp: adiciona os dados no JSON com o horário atual UTC
func projetoKorpHandler(w http.ResponseWriter, r *http.Request) {
	httpRequestsTotal.WithLabelValues("/projeto-korp").Inc()

	res := Response{
		Nome:    "Projeto Korp",
		Horario: time.Now().UTC().Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
// Definindo os endpoints da aplicação e do endpoint de métricas
func main() {
	http.HandleFunc("/projeto-korp", projetoKorpHandler)
	http.Handle("/metrics", promhttp.Handler())

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
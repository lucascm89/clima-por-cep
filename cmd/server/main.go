package main

import (
	"fmt"
	"log"
	"net/http"

	"clima-por-cep/internal/config"
	"clima-por-cep/internal/infra/viacep"
	"clima-por-cep/internal/infra/weatherapi"
	"clima-por-cep/internal/usecase"
	"clima-por-cep/internal/web"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	viaCEPClient := viacep.NewClient()
	weatherAPIClient := weatherapi.NewClient(cfg.WeatherAPIKey)

	getWeatherByZipcodeUseCase := usecase.NewGetWeatherByZipcodeUseCase(
		viaCEPClient,
		weatherAPIClient,
	)

	handler := web.NewHandler(getWeatherByZipcodeUseCase)

	mux := http.NewServeMux()

	mux.HandleFunc("/", healthCheckHandler)
	mux.HandleFunc("/weather/", handler.GetWeatherByZipcode)

	fmt.Println("Server running on port", cfg.WebServerPort)

	if err := http.ListenAndServe(":"+cfg.WebServerPort, mux); err != nil {
		log.Fatal(err)
	}
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("OK"))
}

# Clima por CEP

Projeto desenvolvido em Go para consultar a temperatura atual de uma cidade a partir de um CEP brasileiro.

A aplicação recebe um CEP de 8 dígitos, consulta a cidade através da API ViaCEP e, em seguida, consulta a temperatura atual através da WeatherAPI.

O retorno contém a temperatura em Celsius, Fahrenheit e Kelvin.

## URL publicada no Google Cloud Run

```text
https://clima-por-cep-441837748287.southamerica-east1.run.app
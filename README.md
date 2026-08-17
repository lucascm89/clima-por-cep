# Clima por CEP

Projeto desenvolvido em Go para consultar a temperatura atual de uma cidade a partir de um CEP brasileiro.

A aplicação recebe um CEP de 8 dígitos, consulta a cidade através da API ViaCEP e, em seguida, consulta a temperatura atual através da WeatherAPI.

O retorno contém a temperatura em Celsius, Fahrenheit e Kelvin.

## URL publicada no Google Cloud Run

```text
https://clima-por-cep-441837748287.southamerica-east1.run.app
```

## Endpoint principal

```http
GET /weather/{cep}
```

Exemplo:

```bash
curl -i https://clima-por-cep-441837748287.southamerica-east1.run.app/weather/01001000
```

Resposta esperada:

```json
{
  "temp_C": 28.5,
  "temp_F": 83.3,
  "temp_K": 301.65
}
```

A temperatura pode variar, pois é consultada em tempo real na WeatherAPI.

## Cenários de erro

### CEP inválido

```bash
curl -i https://clima-por-cep-441837748287.southamerica-east1.run.app/weather/01001-000
```

Resposta esperada:

```text
HTTP/2 422

invalid zipcode
```

### CEP não encontrado

```bash
curl -i https://clima-por-cep-441837748287.southamerica-east1.run.app/weather/99999999
```

Resposta esperada:

```text
HTTP/2 404

can not find zipcode
```

---

## Como rodar os testes

Na raiz do projeto, execute:

```bash
go test ./...
```

Esse comando executa todos os testes automatizados do projeto, incluindo testes de validação de CEP, conversão de temperatura, use case, clients externos e handlers HTTP.

---

## Como rodar localmente via Docker

### 1. Criar arquivo `.env`

Na raiz do projeto, crie um arquivo `.env` com base no `.env.example`.

Exemplo:

```env
WEB_SERVER_PORT=8080
WEATHER_API_KEY=sua_chave_da_weatherapi
```

A variável `WEATHER_API_KEY` deve conter uma chave válida da WeatherAPI.

### 2. Subir a aplicação com Docker Compose

Na raiz do projeto, execute:

```bash
docker compose up --build
```

A aplicação ficará disponível em:

```text
http://localhost:8080
```

### 3. Testar a aplicação local

CEP válido:

```bash
curl -i http://localhost:8080/weather/01001000
```

Resposta esperada:

```json
{
  "temp_C": 28.5,
  "temp_F": 83.3,
  "temp_K": 301.65
}
```

CEP inválido:

```bash
curl -i http://localhost:8080/weather/01001-000
```

Resposta esperada:

```text
HTTP/1.1 422 Unprocessable Entity

invalid zipcode
```

CEP não encontrado:

```bash
curl -i http://localhost:8080/weather/99999999
```

Resposta esperada:

```text
HTTP/1.1 404 Not Found

can not find zipcode
```

### 4. Parar a aplicação

```bash
docker compose down
```

---

## Rodando localmente sem Docker

Também é possível rodar diretamente com Go:

```bash
go run ./cmd/server
```

Depois teste:

```bash
curl -i http://localhost:8080/weather/01001000
```

---

## Variáveis de ambiente

| Variável | Descrição |
|---|---|
| `WEB_SERVER_PORT` | Porta local da aplicação |
| `WEATHER_API_KEY` | Chave da WeatherAPI |

O Cloud Run utiliza automaticamente a variável `PORT`, por isso a aplicação está preparada para rodar tanto localmente quanto no Google Cloud Run.

---

## Tecnologias utilizadas

- Go
- Docker
- Docker Compose
- Google Cloud Run
- ViaCEP
- WeatherAPI
# Calculadora de Notas em Golang

Este projeto implementa um serviço em Golang que calcula a quantidade mínima de notas necessárias para formar um valor monetário específico. As notas disponíveis são de 200, 100, 50, 20, 10 e 5 reais.

## Estrutura do Projeto

```plaintext
projeto/
├── main.go
├── handlers/
│   └── notas.go
├── services/
│   └── calcular.go
└── models/
    └── response.go
```

## Funcionalidades
Recebe um valor monetário via HTTP request.
Calcula a quantidade mínima de notas necessárias.
Retorna o resultado em formato JSON.
Informa se o valor solicitado não pode ser formado com as notas disponíveis.

## Instalação
Clone o repositório:
```plaintext 
git clone https://github.com/vinialeixo/calculadora
```


## Testando o Endpoint
Você pode testar o endpoint usando ferramentas como Insomnia ou Postman.

Exemplo de Requisição
Endpoint: GET /notas
Parâmetro: valor (o valor monetário a ser calculado)

URL Exemplo
http://localhost:8080/notas?valor=770


Exemplo de Resposta
Para um valor de 770 reais:
```plaintext 
{
  "notas": {
    "200": 3,
    "100": 1,
    "50": 1,
    "20": 1
  }
}
```
Para um valor indisponível (por exemplo, 3 reais):
```plaintext 
{
  "mensagem": "Valor indisponível para saque."
}
```

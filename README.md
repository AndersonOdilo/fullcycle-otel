# fullcycle-otel

## Portas Utilizadas
- Servidor web: 8080
- Zipkin: 9411

## Enpoints disponivel

- Local: POST http://localhost:8080/

## Como rodar o projeto

1. Rodar o comando para iniciar o banco de dados , rodar migration e iniciar o servidor
``` shell
docker-compose up
```

## Como testar o projeto

1. Teste a aplicação REST API server
    - faça as chamadas usando o arquivo [api.http](api/api.http)

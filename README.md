# fullcycle-google-cloud-run

## Portas Utilizadas
- Servidor web: 8080


## Enpoints disponivel

- Local: localhost:8080/temp/{cep}
- Google Cloud Run: https://fullcycle-deploy-cloud-run-287183371119.us-central1.run.app/temp/{cep}

## Como rodar o projeto

1. Rodar o comando para iniciar o banco de dados , rodar migration e iniciar o servidor
``` shell
docker-compose up
```

## Como testar o projeto

1. Teste a aplicação REST API server
    - faça as chamadas usando o arquivo [api.http](api/api.http)
    - rode os testes unitarios da aplicação
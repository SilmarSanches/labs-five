# labs-five

Projeto de conclusão de pós-graduação

## Indice
1. [Build-Dev](#build-dev)
2. [Docker](#docker)
3. [Testes](#testes)
4. [Swagger](#swagger)


# Build-Dev

Para rodar em dev e testar:
```bash
go run ./cmd/main.go --url=http://www.google.com --requests=100 --concurrency=10
```

# Docker

Para gerar a imagem: 

```bash
docker build -t labs-five .
````

Para executar o teste no docker:

```bash
docker run labs-five --url=http://google.com --requests=1000 --concurrency=10
```
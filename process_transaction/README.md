

* **Dockerfile**
```dockerfile
FROM golang:1.17

WORKDIR /go/src


# Esse comando serve pra manter o container rodando.
# É usado pra programar em um container Docker, sem precisar instalar o Golang mas serve pra outras linguagens
CMD ["tail", "-f", "/dev/null"]
```

* **Docker Compose**
```yaml
version: "3"

services:
  app:
  # Constroi a imagem a partir do Dockerfile localizado no Diretório atual em .
    build: .

    # Criamos esse volume pra que qualquer arquivo criado dentro do container em /go/src, seja
    # espelhado no diretório atual e vice versa, usado pra programar dentro do Docker
    volumes:
      - .:/go/src/

```


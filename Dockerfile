# Use a imagem oficial do Golang como base
FROM golang:1.21-alpine as base
RUN rm -rf /var/cache/apk/* && rm -rf /tmp/*

# Defina o diretório de trabalho dentro do contêiner
WORKDIR /app

# Copie os módulos de dependência primeiro para melhorar o caching das camadas
COPY app/go.mod go.mod
COPY app/go.sum go.sum
RUN go mod download

# Copie o resto dos arquivos de código fonte para o contêiner
COPY app/ /app

FROM base AS dev
# Instale o gin para hotreload em dev
RUN go install github.com/codegangsta/gin@latest


# Instale o swag
RUN go install github.com/swaggo/swag/cmd/swag@latest
# Gere a documentação Swagger
RUN swag init
ENTRYPOINT sh -c "gin --appPort 8080 --bin build/gin-bin run main.go"

FROM base AS build
WORKDIR /app

# Copie o script de testes
COPY scripts/test.sh /scripts/test.sh
RUN chmod +x /scripts/test.sh


# # Executar os testes na fase de build
# RUN /scripts/test.sh


# Etapa final
FROM golang:1.21-alpine
WORKDIR /app
COPY --from=build /app /app

# Comando padrão: executar a aplicação
CMD ["./app"]
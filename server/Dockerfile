# Определите базовый образ для сборки приложения (Go 1.23 на Alpine Linux)
FROM golang:1.23-alpine as builder

# Установите необходимые зависимости, включая OpenSSL и инструменты разработки
# RUN apk add --update --no-cache openssl curl g++ gcc libxslt-dev

# Установите рабочую директорию внутри контейнера сборки
WORKDIR /build

# Скопируйте исходный код в контейнер сборки
COPY . .

# Соберите приложение Go для Linux (amd64) без использования CGO
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server .

# Создайте новую стадию для финального образа приложения (на основе Alpine Linux)
FROM alpine:3.18 as hoster

# Установите необходимые зависимости, включая cURL и инструменты разработки
# RUN apk add --update --no-cache openssl curl g++ gcc libxslt-dev

# Скопируйте конфигурационные файлы и собранное приложение из стадии сборки
COPY --from=builder /build/.env ./.env
#COPY --from=builder /build/assets ./assets
#COPY --from=builder /build/templates ./templates
COPY --from=builder /build/server ./server

# Определите точку входа для финального образа приложения
ENTRYPOINT ["./server"]
# Используем официальный образ Go
FROM golang:1.23-alpine

# Устанавливаем рабочую директорию
WORKDIR /app

# Устанавливаем аргументы сборки
ARG MODULE_PATH
ARG SERVICE_PATH

# Копируем go.mod и go.sum из указанного пути
COPY ${MODULE_PATH}/go.mod ./
COPY ${MODULE_PATH}/go.sum ./
RUN go mod download

# Копируем остальные файлы из указанного пути
COPY ${SERVICE_PATH}/ ./

COPY .env ./

# RUN chmod +x /app/main

# Собираем приложение
RUN go build -o main .



# Запускаем приложение
CMD ["/app/main"]

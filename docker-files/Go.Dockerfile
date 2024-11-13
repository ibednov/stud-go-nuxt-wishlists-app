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

# Устанавливаем swag
RUN go install github.com/swaggo/swag/cmd/swag@latest

# Добавляем путь к исполняемым файлам Go
ENV PATH="/go/bin:${PATH}"

# Копируем остальные файлы из указанного пути
COPY ${SERVICE_PATH}/ ./

# Проверяем содержимое директории после копирования
# RUN ls -l /app

# Копируем .env файл
COPY .env ./

# Собираем приложение
RUN go build -o main . || { echo 'Build failed'; exit 1; }  # Проверяем на ошибки сборки

# Проверяем наличие файла main
RUN ls -l /app

# Запускаем приложение
CMD ["/app/main"]

FROM golang:1.23.0-alpine
# Устанавливаем рабочую директорию
WORKDIR /app
# Копируем файл go.mod (зависимости)
COPY go.mod ./
# Подтягиваем зависимости проекта
RUN go mod download
# Копируем код проекта
COPY . .
# Выполянем команду сборки исполнительного файла
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bot .
# Открываем порт бэкенда
EXPOSE 8080
# 14 строка не запустится при сборке
CMD ["/app/bot"]
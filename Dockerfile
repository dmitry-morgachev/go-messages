# Используем официальный образ Golang на основе Debian в качестве базового
FROM golang:1.22

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Установка необходимых пакетов
RUN apt-get update && apt-get install -y \
    build-essential \
    wget \
    git \
    pkg-config \
    libssl-dev \
    libcurl4-openssl-dev \
    libsasl2-dev \
    zlib1g-dev \
    librdkafka-dev

# Копируем go.mod и go.sum для установки зависимостей
COPY go.mod go.sum ./

# Устанавливаем зависимости
RUN go mod download

# Копируем весь исходный код
COPY . .

# Компилируем приложение
RUN go build -tags dynamic -o /message-service cmd/main.go

# Устанавливаем рабочую директорию для выполнения
WORKDIR /

# Устанавливаем переменные окружения для подключения к базе данных
ENV DB_HOST=postgres
ENV DB_PORT=5432
ENV DB_USER=postgres
ENV DB_PASSWORD=yourpassword
ENV DB_NAME=go_messages
ENV KAFKA_BROKER=kafka:9092
ENV KAFKA_TOPIC=messages
ENV SERVER_ADDRESS=:8080

# Определяем команду для запуска приложения
CMD ["/message-service"]

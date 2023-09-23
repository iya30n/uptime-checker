FROM golang:latest as go

WORKDIR /app

COPY . .

RUN go build -mod=vendor cmd/uptime/main.go
RUN go build -mod=vendor cmd/database/Migrate.go

EXPOSE 7000

CMD ["sh", "-c", "./Migrate && ./main"]
FROM golang:latest as go

WORKDIR /app

COPY . .

RUN go build -mod=vendor cmd/uptime/main.go && go build -mod=vendor -o migrate cmd/database/Migrate.go

EXPOSE 7000
CMD ["./main", "./migrate"]
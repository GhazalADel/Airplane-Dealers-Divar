FROM golang:1.20.5

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

EXPOSE 8080
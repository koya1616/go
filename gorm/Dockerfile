# docker exec -it go-app /bin/sh
FROM golang:1.23.4-alpine

WORKDIR /api

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . ./

CMD ["go", "run", "main.go"]


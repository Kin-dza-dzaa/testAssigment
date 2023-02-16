FROM golang:1.18 AS build

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN CGO_ENABLED=0 go build -o /usr/src/app/bin/app ./cmd/

FROM alpine:latest

WORKDIR /usr/src/app

COPY --from=build /usr/src/app/bin/app .

EXPOSE 8001
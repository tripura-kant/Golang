FROM golang:1.18-alpine as go-builder

WORKDIR /app

COPY . .


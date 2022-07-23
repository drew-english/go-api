FROM golang:1.18-alpine AS build

WORKDIR /go-build
COPY . .

RUN go mod download && go mod verify
RUN go build -o /go-build/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=build /go-build .

CMD "./main.go"

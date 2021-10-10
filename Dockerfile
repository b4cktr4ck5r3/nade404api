FROM golang:1.15 AS build

WORKDIR /go/src/nade404api

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o main .

FROM alpine:latest

WORKDIR /app

COPY --from=build /go/src/nade404api .

EXPOSE 3000

CMD ["./main"]
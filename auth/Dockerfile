FROM golang:1.24.0-alpine AS build
WORKDIR /app
ADD . /app

RUN go mod download \
    && go build -o /bin/main ./cmd/main.go

FROM alpine:3.9
COPY ./config/config.json ./config/config.json
COPY --from=build /bin/main /bin/main
ENTRYPOINT ["/bin/main"]

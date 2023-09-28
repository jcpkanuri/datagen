#build stage
FROM golang:1.21.1-alpine3.18 as builder
WORKDIR /app
ADD . /app
RUN go mod tidy
RUN go build

# Deploy
FROM alpine:latest
RUN apk add file
WORKDIR /app
COPY --from=builder /app/datagen .
ENTRYPOINT [ "/app/datagen" ]

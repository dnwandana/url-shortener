FROM golang:1.17-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o main .
WORKDIR /dist
RUN cp /app/main .

FROM alpine:3.14
COPY --from=builder /dist/main /
CMD [ "/main" ]

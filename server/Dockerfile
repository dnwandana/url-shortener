FROM golang:1.16.6-alpine AS builder
WORKDIR /server
COPY . .
RUN go mod download
RUN go build -o main .
WORKDIR /dist
RUN cp /server/main .

FROM alpine:3.14
COPY --from=builder /dist/main /
CMD [ "/main" ]

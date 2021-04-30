FROM golang:1.16-alpine AS builder
WORKDIR /server
COPY . .
RUN go mod download
RUN go build -o main .
WORKDIR /dist
RUN cp /server/main .

FROM alpine:3.13
COPY --from=builder /dist/main /
CMD [ "/main" ]

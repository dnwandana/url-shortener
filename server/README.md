# Server-side Application

This is server-side application part of URL Shortener. Built using Golang, Fiber, and mongoDB.

## API Spec

Read more at [api-spec.yaml](./api-spec.yaml)

## Project Setup

### Environment Variables

- MONGO_URI
- MONGO_MIN_POOL
- MONGO_MAX_POOL
- MONGO_MAX_CONN_IDLE
- MONGO_DATABASE
- URL_COLLECTION
- DOMAIN

For example, you can see [`.env.example`](.env.example)

### Downloading Packages

```bash
go mod download
```

### Running The Application

```bash
go run main.go
```

### Compile The Application

```bash
go build main.go
```

## Dockerize The Application (Development)

1.  Build docker image
    ```bash
    docker build -t url-server:1.0 .
    ```
2.  Run docker image
    ```bash
    docker run -d --name url-server \
    -p 5000:5000 \
    -e MONGO_URI="mongodb://username:password@host:port" \
    -e MONGO_MIN_POOL="10" \
    -e MONGO_MAX_POOL="100" \
    -e MONGO_MAX_CONN_IDLE="60" \
    -e MONGO_DATABASE="urlShortener" \
    -e URL_COLLECTION="urls" \
    -e DOMAIN="http://localhost:5000" \
    url-server:1.0
    ```

## Useful Links

To learn more about this project, take a look at the following resources:

- [Fiber Documentation](https://docs.gofiber.io/)
- [MongoDB Go Driver](https://github.com/mongodb/mongo-go-driver)
- [Golang Package Validator](https://github.com/go-playground/validator)

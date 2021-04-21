# Server-side Application

This is server-side application part of URL Shortener. Built using Golang, Fiber, and mongoDB.

## API Spec

Read more at [API_SPEC.md](./API_SPEC.md)

## Project Setup

### Environment Variables

- MONGO_URI
- MONGO_MIN_POOL
- MONGO_MAX_POOL
- MONGO_MAX_CONN_IDLE
- MONGO_DATABASE
- URL_COLLECTION
- USER_COLLECTION
- JWT_SECRET
- JWT_LIFE
- STAGE

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

## Dockerize The Application

1.  Build docker image
    ```bash
    docker build -t url-server:1.0.1 .
    ```
2.  Run docker image
    ```bash
    docker run -d --name url-server \
    -p 5000:5000 \
    -e MONGO_URI="mongodb://localhost:27017" \
    -e MONGO_MIN_POOL=10 \
    -e MONGO_MAX_POOL=100 \
    -e MONGO_MAX_CONN_IDLE=60 \
    -e MONGO_DATABASE="urlShortener" \
    -e URL_COLLECTION="urls" \
    -e USER_COLLECTION="users" \
    -e JWT_SECRET="SECRET" \
    -e JWT_LIFE=6 \
    -e STAGE="DEVELOPMENT" \
    url-server:1.0.1
    ```

## Reference

- [Fiber Documentation](https://docs.gofiber.io/ "Fiber Documentation")
- [MongoDB Go Driver](https://github.com/mongodb/mongo-go-driver "MongoDB Go Driver")
- [Golang Package Validator](https://github.com/go-playground/validator "Golang Package Validator")

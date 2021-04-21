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

```go
go mod download
```

### Running The Application

```go
go run main.go
```

### Compile The Application

```go
go build main.go
```

## Reference

- [Fiber Documentation](https://docs.gofiber.io/ "Fiber Documentation")
- [MongoDB Go Driver](https://github.com/mongodb/mongo-go-driver "MongoDB Go Driver")
- [Golang Package Validator](https://github.com/go-playground/validator "Golang Package Validator")

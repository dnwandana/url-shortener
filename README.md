# URL Shortener

Fullstack URL Shortener, built with MongoDB, Golang GoFiber, and TypeScript Next.js

## Server Application

Read more at [server/README.md](server/README.md)

## Client Application

Read more at [client/README.md](client/README.md)

## Dockerize Application

### Setup MongoDB TTL Indexes

1. Login to your mongoDB cluster using [mongoDB shell](https://docs.mongodb.com/mongodb-shell/)
2. Add TTL Index at document field named `expireAt`
   ```bash
   db.<collectionName>.createIndex( { "expireAt": 1 }, { expireAfterSeconds: 0 } )
   ```
   Read more about [Expire Documents at a Specific Clock Time at mongoDB Documentation](https://docs.mongodb.com/manual/tutorial/expire-data/#expire-documents-at-a-specific-clock-time)

### Builds, (re)creates, starts, and attaches to containers for a service

```bash
docker-compose up -d
```

### Stops and removes containers, networks, volumes, and images created by up.

```bash
docker-compose down
```

### Delete and removes containers, networks, volumes, and images created by up.

```bash
docker-compose down --rmi all -v
```

## Useful Links

To learn more about this project, take a look at the following resources:

- [Overview of Docker Compose](https://docs.docker.com/compose/)
- [NGINX Reverse Proxy](https://docs.nginx.com/nginx/admin-guide/web-server/reverse-proxy/)
- [MongoDB Shell (mongosh)](https://docs.mongodb.com/mongodb-shell/)
- [MongoDB manual about TTL Indexes](https://docs.mongodb.com/manual/core/index-ttl/)

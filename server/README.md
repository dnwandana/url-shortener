# Server-side Application

## API Spec

### Shorten URL

Request:

- Method: `POST`
- Endpoint: `/go`
- Headers:
  - Content-Type: application/json
  - Accept: application/json
- Body:
  ```json
  {
    "url": "string"
  }
  ```

Success Response:

- Http status code: `201`
- Body:
  ```json
  {
    "data": {
      "longUrl": "integer",
      "shortUrl": "string"
    }
  }
  ```

Error Response:

- Http status code: `400`
- Body:
  ```json
  {
    "error": {
      "message": "string"
    }
  }
  ```

## Project Setup

### Install Depedencies

```bash
yarn install
```

### Setup Environment Variables

- DB_URI
- DOMAIN

for example, you can see [`.env.example`](.env.example)

### Compiles and Hot-reloads for Development

```bash
yarn dev
```

### Build and Run

```bash
# Compiles Source Code
yarn build

# Run the Application
yarn start
```

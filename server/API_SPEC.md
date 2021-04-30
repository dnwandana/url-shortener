# API Specification

## User Sign Up

- Method: `POST`
- Endpoint: `/go/sign-up`
- Headers:
  - Content-Type: application/json
  - Accept: application/json
- Body:
  ```json
  {
    "fullname": "string",
    "email": "string",
    "password": "string",
    "confirmationPassword": "string"
  }
  ```

Success Response:

- Http status code: `201`
- Body:
  ```json
  {
    "statusCode": "integer",
    "user": {
      "id": "string",
      "fullname": "string"
    }
  }
  ```

Error Response:

- Http status code: `400`
- Body:
  ```json
  {
    "statusCode": "integer",
    "error": ["string"]
  }
  ```

## User Sign In

- Method: `POST`
- Endpoint: `/go/sign-in`
- Headers:
  - Content-Type: application/json
  - Accept: application/json
- Body:
  ```json
  {
    "email": "string",
    "password": "string"
  }
  ```

Success Response:

- Http status code: `200`
- Headers:
  - userId Cookie
  - JWT Cookie
- Body:
  `OK`

Error Response:

- Http status code: `400`
- Body:
  ```json
  {
    "statusCode": "integer",
    "error": ["string"]
  }
  ```

## List All Short URLs

Request:

- Method: `GET`
- Endpoint: `/go`
- Headers:
  - userId Cookie
  - JWT Cookie
  - Accept: application/json

Success Response:

- Http status code: `200`
- Body:
  ```json
  {
    "statusCode": "integer",
    "url": [
      {
        "id": "string",
        "title": "string",
        "url": "string",
        "createdAt": "datetime",
        "updatedAt": "datetime"
      }
    ]
  }
  ```

Error Response:

- Http status code: `400`
- Body:
  ```json
  {
    "statusCode": "integer",
    "error": "string"
  }
  ```

## Shorten URL

Request:

- Method: `POST`
- Endpoint: `/go`
- Headers:
  - userId Cookie (Optional)
  - JWT Cookie (Optional)
  - Content-Type: application/json
  - Accept: application/json
- Body:
  ```json
  {
    "id": "optional string",
    "title": "optional string",
    "url": "string"
  }
  ```

Success Response:

- Http status code: `201`
- Body:
  ```json
  {
    "statusCode": "integer",
    "url": {
      "id": "string",
      "title": "string",
      "url": "string",
      "createdAt": "datetime",
      "updatedAt": "datetime"
    }
  }
  ```

Error Response:

- Http status code: `400`
- Body:
  ```json
  {
    "statusCode": "integer",
    "error": "string"
  }
  ```

## Get Specific Short URL

Request:

- Method: `GET`
- Endpoint: `/go/{id}`

Success Response:

- Redirected to `url`

Error Response:

- Redirected to `/404`

## Update Short URL

Request:

- Method: `PUT`
- Endpoint: `/go/{id}`
- Headers:
  - userId Cookie
  - JWT Cookie
  - Content-Type: application/json
  - Accept: application/json
- Body:
  ```json
  {
    "id": "string",
    "title": "string",
    "url": "string"
  }
  ```

Success Response:

- Http status code: `200`
- Body:
  ```json
  {
    "statusCode": "integer",
    "url": {
      "id": "string",
      "title": "string",
      "url": "string",
      "createdAt": "datetime",
      "updatedAt": "datetime"
    }
  }
  ```

Error Response:

- Http status code: `400`
- Body:
  ```json
  {
    "statusCode": "integer",
    "error": "string"
  }
  ```

## Delete Short URL

Request:

- Method: `DELETE`
- Endpoint: `/go/{id}`
- Headers:
  - userId Cookie
  - JWT Cookie

Success Response:

- Http status code: `204`

Error Response:

- Http status code: `400`
- Body:
  ```json
  {
    "statusCode": "integer",
    "error": "string"
  }
  ```

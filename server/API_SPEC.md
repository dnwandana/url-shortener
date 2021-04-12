# API Specification

## List All Short URLs

Request:

- Method: `GET`
- Endpoint: `/go`
- Headers:
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
        "url": "string"
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
    "statusCode": "integer",
    "url": {
      "id": "string",
      "title": "string",
      "url": "string"
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
      "url": "string"
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

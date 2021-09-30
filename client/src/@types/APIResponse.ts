type HttpCode = 200 | 201 | 400 | 500

type HttpStatus = "OK" | "CREATED" | "BAD_REQUEST" | "INTERNAL_SERVER_ERROR"

type ResponseMessage = {
  code: HttpCode
  status: HttpStatus
  message: string
}

type ResponseData = {
  code: HttpCode
  status: HttpStatus
  data: {
    id: string
    long_url: string
    short_url: string
    secret_key: string
    expire_at: string
  }
}

export type { ResponseMessage, ResponseData }

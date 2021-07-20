type SuccessResponse = {
  statusCode: 200 | 201
  data: {
    id: string
    title: string
    url: string
    createdAt: string
    updatedAt: string
  }
}

type ErrorResponse = {
  statusCode: 400
  error: string
}

type UserInformationResponse = {
  statusCode: 200
  data: {
    fullname: string
    email: string
  }
}

export type { SuccessResponse, ErrorResponse, UserInformationResponse }

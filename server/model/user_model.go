package model

// UserSignIn struct represents form data when trying to sign-in.
type UserSignIn struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

// UserSignUp struct represents form data when trying to sign-up.
type UserSignUp struct {
	Fullname             string `json:"fullname" validate:"required"`
	Email                string `json:"email" validate:"required,email"`
	Password             string `json:"password" validate:"required,min=8"`
	ConfirmationPassword string `json:"confirmationPassword" validate:"required,min=8,eqfield=Password"`
}

// UserInformation struct represents user information data.
type UserInformation struct {
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
}

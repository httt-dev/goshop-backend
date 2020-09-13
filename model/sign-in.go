package model
type SignIn struct {
	//FullName string `json:"fullName,omitempty" validate:"required"`
	Password string `json:password,omitempty`
	Email string `json:email,omitempty`

}

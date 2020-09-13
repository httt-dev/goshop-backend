package model

type SignUp struct {
	//FullName string `json:"fullName,omitempty" validate:"required"`
	FullName string `json:"fullName,omitempty"`
	Phone string `json:phone,oimitempty`
	Password string `json:password,omitempty`
	Email string `json:email,omitempty`
	Address string `json:address,omitempty`
	Avatar string `json:avatar,omitempty`

}


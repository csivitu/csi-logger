package schemas

type UserCreateSchema struct {
	Name            string `json:"name" validate:"required,max=25"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,min=8"`
	ConfirmPassword string `json:"confirmPassword" validate:"required,min=8"`
}

type UserLoginSchema struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type MakeAdminSchema struct {
	Email string `json:"email" validate:"required,email"`
}
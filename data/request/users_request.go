package request

type CreateUsersRequest struct {
	Username string `validate:"required,min=2,max=100" json:"username"`
	Email    string `validate:"required,min=2,max=100" json:"email"`
	Password string `validate:"required,min=2,max=100" json:"password"`
}

type UpdateUsersRequest struct {
	Id       int    `validate:"required"`
	Username string `validate:"required,max=200,min=2" json:"username"`
	Email    string `validate:"required,min=2,max=100" json:"email"`
	Password string `validate:"required,min=2,max=100" json:"password"`
}

type LoginRequest struct {
	Username string `validate:"required,max=200,min=2" json:"username"`
	Password string `validate:"required,min=2,max=100" json:"password"`
}

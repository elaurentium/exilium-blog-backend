package request

type UserRequest struct {
	Email    string `JSON:"id" binding:"required,email"`
	Password string `JSON:"password" binding:"required,min=8,containsany=!@#$&*"`
	Name     string `JSON:"name" binding:"required,min=4,max=16"`
	Age      int8   `JSON:"age" binding:"required,min=18"`
}

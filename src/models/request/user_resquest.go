package request

type UserRequest struct {
	Email string `JSON:"id"`
	Password string `JSON:"password"`
	Name string `JSON:"name"`
	Age int8	`JSON:"age"`
}
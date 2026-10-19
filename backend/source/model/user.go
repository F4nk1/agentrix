package model

type User struct {
	Id			int 	`json:"id,omitempty" db:"id"`
	FirstNames	string 	`json:"first_names,omitempty" db:"first_names"`
	LastNames	string 	`json:"last_names,omitempty" db:"last_names"`
	Email		string 	`json:"email,omitempty" db:"email"`
	PhoneNumber string 	`json:"phone_number,omitempty" db:"phone_number"`
	Password	string 	`json:"password,omitempty" db:"password"`
	RoleId		string	`json:"role_id,omitempty" db:"role_id"`
	Active		bool 	`json:"active,omitempty" db:"active"`
}
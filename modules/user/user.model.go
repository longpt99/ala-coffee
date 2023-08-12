package user

import "time"

type LoginUserReq struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type SignUpUserReq struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	Name     string `json:"name" validate:"required"`
}

type CreateUserReq struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type UpdateUserReq struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type User struct {
	ID        string     `json:"id" db:"id"`
	Name      string     `json:"name" db:"name"`
	Email     string     `json:"email" db:"email"`
	DOB       string     `json:"dob" db:"dob"`
	Status    string     `json:"status" db:"status"`
	CreatedAt *time.Time `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
	DeletedAt *time.Time `json:"-" db:"deleted_at"`
}

type UpdateUserProfileReq struct {
	Name  string `json:"name"`
	Email string `json:"email" validate:"email"`
	DOB   string `json:"dob" validate:"date_string"`
}

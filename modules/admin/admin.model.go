package admin

import "time"

type CreateAdminReq struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required"`
}

type CreateAdminParams struct {
	CreateAdminReq
	Password string `json:"password"`
}

type LoginAdminReq struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UpdateAdminReq struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Admin struct {
	ID        string     `json:"id" db:"id"`
	Name      string     `json:"name" db:"name"`
	Email     string     `json:"email" db:"email"`
	Password  string     `json:"-" db:"password"`
	DeletedAt *time.Time `json:"-" db:"deleted_at"`
	CreatedAt *time.Time `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
}

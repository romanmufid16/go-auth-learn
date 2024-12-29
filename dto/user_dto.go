package dto

type RegisterUser struct {
	Name     string `json:"name" validate:"required,min=1,max=100"` // Validasi bahwa nama tidak boleh kosong dan memiliki panjang antara 1 dan 100
	Email    string `json:"email" validate:"required,email"`        // Validasi email harus ada dan formatnya valid
	Password string `json:"password" validate:"required,min=6"`     // Validasi password harus ada dan panjangnya minimal 6 karakter
}

type LoginUser struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type UserResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

type UpdateUser struct {
	Name     string `json:"name" validate:"omitempty,min=1,max=100"`
	Password string `json:"password" validate:"omitempty,min=6,max=100"`
}

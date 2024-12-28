package models

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"not null"`        // Menambahkan validasi agar field ini tidak kosong
	Email     string    `gorm:"unique;not null"` // Kolom email unik dan tidak boleh kosong
	Password  string    `gorm:"not null"`        // Menambahkan validasi agar password tidak kosong
	CreatedAt time.Time `gorm:"autoCreateTime"`  // GORM mengatur CreatedAt secara otomatis
	UpdatedAt time.Time `gorm:"autoUpdateTime"`  // GORM mengatur UpdatedAt secara otomatis
}

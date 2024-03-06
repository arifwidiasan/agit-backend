package model

import "time"

type Karyawan struct {
	ID           int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Nama         string    `gorm:"not null" json:"nama"`
	NIP          string    `json:"nip"`
	TempatLahir  string    `json:"tempat_lahir"`
	TanggalLahir time.Time `json:"tanggal_lahir"`
	Umur         uint      `json:"umur"`
	Alamat       string    `json:"alamat"`
	Agama        string    `json:"agama"`
	JenisKelamin string    `json:"jenis_kelamin"`
	NoHP         string    `json:"no_hp"`
	Email        string    `json:"email"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	DeletedAt    time.Time `gorm:"index" json:"deleted_at"`
}

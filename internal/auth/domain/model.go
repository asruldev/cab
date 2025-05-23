// 📦 Package domain
// Folder/Package domain biasanya berisi definisi entitas bisnis dan kontrak (interface). Ini adalah lapisan paling inti dari aplikasi, dan tidak boleh bergantung pada detail teknis (seperti database, framework, dll).
package domain

// Ini adalah entitas User, mewakili data pengguna. Entitas ini hanya menyimpan informasi, tanpa tahu dari mana data itu berasal (database, API, dll).
type User struct {
	ID       string
	Email    string
	Password string // Simpel: plaintext untuk contoh (jangan di prod)
}

// Ini adalah kontrak repositori yang menjelaskan cara untuk:
// mencari user berdasarkan email
// menyimpan user baru
// Repositori ini nanti akan di-implementasikan oleh layer database, misalnya menggunakan PostgreSQL, MongoDB, atau lainnya.
// Kenapa pakai interface?
// Agar logika bisnis tidak tergantung pada database tertentu. Kita bisa swap Mongo ↔ PostgreSQL tanpa mengubah logic di usecase.
// setelah ini maka kerjakan repository -->
type AuthRepository interface {
	FindByEmail(email string) (*User, error)
	CreateUser(user *User) error // ini digunakan pada repo
}

// Ini adalah kontrak untuk logika bisnis autentikasi, yaitu:
// Login (menghasilkan token dan refreshToken)
// Register (mendaftarkan user baru)
// Bagian ini biasanya di-implementasikan dalam package usecase, dan di situ kamu akan menemukan logika seperti:
// validasi email/password
// pengecekan apakah user sudah ada
// hash password
// generate JWT token
// setelah ini maka kerjakan usecase -->
type AuthUsecase interface {
	Login(email, password string) (token string, refreshToken string, err error)
	Register(email, password string) (usr string, err error)
}

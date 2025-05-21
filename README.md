```
cmd/
└── server/
    ├── main.go           // ⛩️ Entry point, start server
    ├── router.go         // 🔀 Routing antar modul
internal/
└── auth/                 // 🔐 Modul auth (mandiri)
    ├── domain/           // 🧠 Model: struct + interface
    │   └── model.go
    ├── usecase/          // 💡 ViewModel: logika bisnis
    │   └── usecase.go
    ├── delivery/
    │   └── http/         // 🖥 View/Controller: HTTP Handler
    │       └── handler.go
    ├── repository/       // 🗄 Data: memory/DB/API
    │   └── memory_repo.go
    └── dto/              // 📦 Optional: Request/Response object
        └── dto.go
pkg/
├── config/               // ⚙️ App config & env
├── middleware/           // 🛡 JWT, logging, etc.
├── logger/               // 📋 Logrus/zap setup
└── utils/                // 🛠 Helper umum (hash, convert)
```


Clean Architecture adalah pendekatan desain arsitektur perangkat lunak yang **memisahkan tanggung jawab berdasarkan level kepentingan dan dependensi**. Dalam konteks kode Go yang kamu tunjukkan, alur clean architecture umumnya mengikuti struktur seperti ini:

---

## 🧱 Struktur Clean Architecture

```
.
├── internal/
│   └── auth/
│       ├── delivery/     ← Interface layer (HTTP, gRPC, dsb)
│       │   └── http/
│       │       └── handler.go
│       ├── usecase/      ← Business logic layer (Application rules)
│       │   └── usecase.go
│       └── repository/   ← Data access layer (Database, API eksternal)
│           └── postgres.go
├── pkg/
│   └── config/           ← Konfigurasi global (.env, DB connection)
└── server/
    └── router.go         ← Inisialisasi router dan DI antar layer
```

---

## 🔄 Alur Clean Architecture dalam Kode Tadi

Berikut alur aliran data dan dependensinya berdasarkan fungsi `SetupRouter()`:

### 1. **Delivery Layer (`auth/delivery/http`)**

* File: `handler.go`
* Fungsi: Menerima HTTP request dari client (misalnya `/login`).
* Handler `Login` akan:

  * Membaca request body (username/password).
  * Memvalidasi input.
  * Meneruskan ke `usecase`.

```go
handler := authDelivery.New(uc)
mux.HandleFunc("/login", handler.Login)
```

---

### 2. **Usecase Layer (`auth/usecase`)**

* File: `usecase.go`
* Fungsi: Menyimpan logika bisnis. Contoh:

  * Validasi user.
  * Generate token.
  * Cek akses.
* Mengandalkan `repository` untuk interaksi ke data.

```go
uc := authUsecase.New(repo)
```

---

### 3. **Repository Layer (`auth/repository`)**

* File: `postgres.go`
* Fungsi: Mengakses database PostgreSQL.

  * Query user.
  * Simpan atau update token.
* Diinisialisasi dengan koneksi DB dari `config`.

```go
repo := authRepo.NewPostgresRepo(db)
```

---

### 4. **Config Layer (`pkg/config`)**

* File: `config.go`
* Fungsi:

  * Load variabel dari `.env`.
  * Inisialisasi koneksi database.

```go
config.LoadEnv()
db := config.ConnectPostgres()
```

---

## 🔁 Dependency Rule (aturan utama Clean Architecture)

> **Kode di inner layer tidak boleh tahu apa-apa tentang outer layer. Tapi outer layer boleh bergantung ke inner layer.**

Artinya:

* `delivery` tergantung ke `usecase`.
* `usecase` tergantung ke `repository` (interface).
* `repository` tidak tahu apa-apa tentang `usecase` atau `delivery`.

---

## 🔌 Dependency Injection

Clean Architecture biasanya memakai **dependency injection (DI)** manual, seperti di fungsi `SetupRouter()`:

```go
repo -> usecase -> handler
```

Tanpa menggunakan framework DI seperti di Java/Spring, karena Go lebih mengutamakan eksplisit dan sederhana.

---

## 🧠 Manfaat Clean Architecture

1. **Mudah di-test**: logika bisnis bisa dites tanpa HTTP atau DB.
2. **Mudah diubah**: bisa ganti DB (misal: Postgres ke Mongo) tanpa ubah usecase.
3. **Terstruktur**: memisahkan tanggung jawab antar komponen.
4. **Maintainable**: mudah dirawat dan scalable untuk tim besar.

---

## Swagger
Install Global
```
go install github.com/swaggo/swag/cmd/swag@latest
```

lakukan ini jika komentar sudah lengkap
```
swag init -g cmd/main.go --output docs
```

## Mock
Install global
```
go install go.uber.org/mock/mockgen@latest
```

generate mock usecase
```bash
mockgen -source=internal/auth/domain/model.go -destination=internal/auth/mocks/mock_usecase.go -package=mocks
```

Terus buat unit test
```
auth/delivery/http/handler_test.go
```

Jika interface diubah, generate ulang
```
mockgen -source=internal/auth/domain/model.go -destination=internal/auth/mocks/mock_usecase.go -package=mocks
```

Jalankan test
```
go test -v ./internal/auth/delivery/http
```


### Alur Menambah Fitur Baru

1. **Definisikan Fitur di Domain (`internal/[module]/domain/`)**

   * Buat atau tambahkan model/data struct di `model.go` yang merepresentasikan data fitur baru.
   * Tambahkan interface repository (jika perlu) untuk akses data yang dibutuhkan fitur baru.

2. **Tambahkan DTO di Layer DTO (`internal/[module]/dto/`)**

   * Definisikan data transfer object (DTO) untuk input/output fitur baru.
   * Biasanya dipakai untuk validasi dan mapping dari/ke domain model.

3. **Implementasikan Usecase (`internal/[module]/usecase/`)**

   * Tambahkan fungsi bisnis utama di `usecase.go` untuk fitur baru.
   * Usecase ini mengorkestrasi logika bisnis, berinteraksi dengan repository dan domain model.

4. **Buat Repository (jika ada interaksi DB) (`internal/[module]/repository/`)**

   * Buat/mutakhirkan fungsi akses database di file repository (`postgres_repo.go` atau repo lain).
   * Repository bertugas komunikasi langsung dengan DB.

5. **Buat Delivery Handler HTTP (`internal/[module]/delivery/http/handler.go`)**

   * Tambahkan handler HTTP untuk endpoint fitur baru (misal POST /employee atau GET /auth/profile).
   * Handler ini menerima request, memanggil usecase, dan mengembalikan response.

6. **Daftarkan Route Baru (`cmd/server/router.go`)**

   * Tambahkan route untuk handler fitur baru.
   * Pastikan middleware dan authorization sudah sesuai kalau perlu.

7. **Update Konfigurasi dan Dependency Injection (jika ada)**

   * Kalau fitur baru butuh dependency baru, inject di `main.go` atau package inisialisasi lain.
   * Update `go.mod` kalau ada dependency eksternal baru.

8. **Testing**

   * Buat unit test untuk domain, usecase, dan handler.
   * Buat integration test untuk endpoint baru.

---

### Contoh Singkat (Misal fitur menambah employee baru)

* `internal/employee/domain/model.go`
  Buat `Employee` struct baru.

* `internal/employee/dto/dto.go`
  Buat `EmployeeCreateRequest` dan `EmployeeResponse`.

* `internal/employee/usecase/usecase.go`
  Tambahkan fungsi `CreateEmployee(input dto.EmployeeCreateRequest) (dto.EmployeeResponse, error)`.

* `internal/employee/repository/postgres_repo.go`
  Implementasikan fungsi simpan employee ke DB.

* `internal/employee/delivery/http/handler.go`
  Tambahkan handler HTTP `CreateEmployeeHandler`.

* `cmd/server/router.go`
  Tambahkan route baru misal `POST /employees`.

* `cmd/main.go`
  Pastikan usecase dan repository di-inject ke handler.

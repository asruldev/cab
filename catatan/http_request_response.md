
## 🧾 1. Apa itu `http.Request`?

`http.Request` adalah representasi dari permintaan (request) yang dikirim oleh client (misalnya browser atau Postman) ke server kamu.

### ✅ Properti Penting:

| Properti           | Kegunaan                                                                    |
| ------------------ | --------------------------------------------------------------------------- |
| `Method`           | Metode HTTP (GET, POST, PUT, DELETE, dsb)                                   |
| `URL.Path`         | Path dari URL (misal `/user/123`)                                           |
| `URL.Query()`      | Ambil query param (misalnya `?id=10`)                                       |
| `Header`           | Semua header yang dikirim client (misalnya `Authorization`, `Content-Type`) |
| `Body`             | Isi body request (biasanya di POST/PUT)                                     |
| `Form`, `PostForm` | Isi form dari body form-url-encoded / multipart                             |
| `RemoteAddr`       | IP Address dari client                                                      |

---

### ✨ Contoh:

```go
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method:", r.Method)
	fmt.Println("Path:", r.URL.Path)
	fmt.Println("Query id:", r.URL.Query().Get("id"))
	fmt.Println("Header User-Agent:", r.Header.Get("User-Agent"))

	// Baca body (hati-hati, hanya bisa dibaca sekali)
	body, _ := io.ReadAll(r.Body)
	fmt.Println("Body:", string(body))
}
```

---

## 📤 2. Apa itu `http.ResponseWriter`?

`http.ResponseWriter` adalah objek yang digunakan untuk mengirim **respons** dari server ke client.

### ✅ Fungsi Umum:

| Fungsi                 | Kegunaan                                       |
| ---------------------- | ---------------------------------------------- |
| `w.Write([]byte(...))` | Menulis isi body response                      |
| `w.WriteHeader(404)`   | Menetapkan status code (harus sebelum `Write`) |
| `w.Header().Set(...)`  | Menetapkan header seperti `Content-Type`       |

---

### ✨ Contoh:

```go
func handler(w http.ResponseWriter, r *http.Request) {
	// Set header
	w.Header().Set("Content-Type", "application/json")

	// Set status code
	w.WriteHeader(http.StatusOK)

	// Kirim response body
	w.Write([]byte(`{"message": "hello"}`))
}
```

> ⚠️ Kalau kamu memanggil `w.WriteHeader` **setelah** `w.Write`, maka status code default `200` akan dikirim otomatis.

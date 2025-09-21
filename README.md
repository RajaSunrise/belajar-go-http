# Belajar Go-HTTP: Contoh Aplikasi API

[![Go Report Card](https://goreportcard.com/badge/github.com/go-lang-id/belajar-go-http)](https://goreportcard.com/report/github.com/go-lang-id/belajar-go-http)

Selamat datang di repositori contoh untuk belajar membangun aplikasi web dan API dengan Go (`Golang`) menggunakan paket `net/http` dari pustaka standar. Proyek ini adalah implementasi praktis dari [TUTORIAL.md](./TUTORIAL.md), sebuah panduan lengkap untuk menguasai `net/http`.

Repositori ini berisi kode untuk **API CRUD (Create, Read, Update, Delete) Todo** yang sederhana dan fungsional.

## ✨ Fitur

- **API RESTful Sederhana**: Endpoint yang jelas untuk mengelola daftar tugas (todos).
- **Router `chi`**: Menggunakan router pihak ketiga yang ringan dan kuat (`go-chi/chi`) untuk routing yang ekspresif.
- **Middleware**: Termasuk contoh middleware untuk logging, recovery dari panic, dan lainnya.
- **Struktur Proyek yang Baik**: Mengikuti praktik terbaik untuk struktur proyek Go yang skalabel (`cmd`, `internal`).
- **Graceful Shutdown**: Implementasi shutdown yang aman untuk memastikan tidak ada request yang terputus saat server dimatikan.
- **Tanpa Database Eksternal**: Menggunakan penyimpanan di memori (*in-memory store*) agar mudah dijalankan tanpa perlu setup database.

## 🚀 Memulai

### Prasyarat

- [Go](https://golang.org/dl/) versi 1.18 atau lebih tinggi.

### Instalasi & Menjalankan Server

1.  **Clone repositori ini:**
    ```sh
    git clone https://github.com/go-lang-id/belajar-go-http.git
    cd belajar-go-http
    ```

2.  **Download dependencies:**
    ```sh
    go mod tidy
    ```

3.  **Jalankan server:**
    ```sh
    go run ./cmd/api
    ```

4.  Server sekarang berjalan di `http://localhost:8080`. Anda akan melihat log di terminal Anda:
    ```
    Server starting on port :8080
    ```

##  API Endpoints

Anda bisa menggunakan `curl`, Postman, atau Insomnia untuk berinteraksi dengan API.

| Metode | Endpoint          | Deskripsi                    | Contoh Body Request                               |
| :------- | :---------------- | :--------------------------- | :------------------------------------------------ |
| `GET`    | `/todos`          | Mendapatkan semua tugas      | -                                                 |
| `POST`   | `/todos`          | Membuat tugas baru           | `{"task": "Belajar Go net/http"}`                 |
| `GET`    | `/todos/{id}`     | Mendapatkan satu tugas       | -                                                 |
| `PUT`    | `/todos/{id}`     | Memperbarui tugas            | `{"task": "Selesai belajar", "completed": true}` |
| `DELETE` | `/todos/{id}`     | Menghapus tugas              | -                                                 |

### Contoh Penggunaan dengan `curl`

- **Membuat tugas baru:**
  ```sh
  curl -X POST -H "Content-Type: application/json" -d '{"task": "Membaca TUTORIAL.md"}' http://localhost:8080/todos
  ```
  *(Respons akan berisi tugas yang baru dibuat beserta ID-nya)*

- **Melihat semua tugas:**
  ```sh
  curl http://localhost:8080/todos
  ```

- **Memperbarui tugas dengan ID 1:**
  ```sh
  curl -X PUT -H "Content-Type: application/json" -d '{"task": "Selesai membaca TUTORIAL.md", "completed": true}' http://localhost:8080/todos/1
  ```

- **Menghapus tugas dengan ID 1:**
  ```sh
  curl -X DELETE http://localhost:8080/todos/1
  ```

## 📂 Struktur Proyek

```
.
├── cmd/api/
│   └── main.go         # Titik masuk aplikasi, setup server, graceful shutdown.
├── internal/
│   ├── handler/
│   │   ├── todo.go     # Handler HTTP untuk setiap endpoint.
│   │   └── routes.go   # Definisi semua rute API menggunakan `chi`.
│   └── model/
│       └── todo.go     # Definisi struct `Todo` dan logika penyimpanan data (in-memory store).
├── go.mod              # Definisi modul Go dan dependensi.
├── go.sum              # Checksum dari dependensi.
├── README.md           # File ini.
└── TUTORIAL.md         # Panduan lengkap dan mendalam tentang `net/http` (sumber dari proyek ini).
```

## 📚 Belajar Lebih Lanjut

Untuk pemahaman yang mendalam tentang setiap konsep yang digunakan dalam proyek ini, silakan baca [**TUTORIAL.md**](./TUTORIAL.md). Dokumen tersebut adalah panduan komprehensif yang mencakup semua aspek `net/http` dari dasar hingga topik lanjutan.

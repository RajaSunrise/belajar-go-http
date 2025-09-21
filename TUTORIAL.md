
# Belajar Go-HTTP: Panduan Lengkap Menguasai `net/http`

![Go Gopher](https://blog.golang.org/gopher/gopher.png)

Selamat datang di panduan paling lengkap untuk mempelajari dan menguasai paket `net/http` dari bahasa pemrograman Go. Panduan ini dirancang untuk membawa Anda dari seorang pemula yang belum pernah menyentuh web development di Go, hingga menjadi seorang developer yang percaya diri dalam membangun aplikasi web, API, dan layanan mikro yang tangguh dan efisien.

`net/http` adalah salah satu permata di dalam pustaka standar (standard library) Go. Tanpa memerlukan framework eksternal, Anda sudah dibekali dengan semua perangkat yang dibutuhkan untuk membangun server web berkinerja tinggi. Filosofi Go tentang kesederhanaan dan komposisi bersinar terang di dalam paket ini.

Dokumen ini sengaja dibuat sangat panjang dan mendetail. Tujuannya adalah untuk menjadi satu-satunya sumber yang Anda butuhkan untuk memahami `net/http` secara mendalam.

## Daftar Isi

1.  [**Pendahuluan: Kenapa `net/http`?**](#1-pendahuluan-kenapa-nethttp)
    *   [Filosofi di Balik `net/http`](#filosofi-di-balik-nethttp)
    *   [Tanpa Sihir (No Magic)](#tanpa-sihir-no-magic)
    *   [Kinerja Luar Biasa](#kinerja-luar-biasa)

2.  [**Bab 1: Server HTTP Pertama Anda - "Hello, World!"**](#2-bab-1-server-http-pertama-anda---hello-world)
    *   [Kode Minimalis](#kode-minimalis)
    *   [Membedah Kode](#membedah-kode)
    *   [Menjalankan Server](#menjalankan-server)

3.  [**Bab 2: Konsep Inti - Memahami Blok Bangunan `net/http`**](#3-bab-2-konsep-inti---memahami-blok-bangunan-nethttp)
    *   [`http.ResponseWriter`: Sang Penulis Respons](#httpresponsewriter-sang-penulis-respons)
    *   [`*http.Request`: Sang Pembawa Pesan](#httprequest-sang-pembawa-pesan)
    *   [`http.Handler`: Jantung dari Server Go](#httphandler-jantung-dari-server-go)
    *   [`http.HandlerFunc`: Adaptor Cerdas](#httphandlerfunc-adaptor-cerdas)
    *   [`http.ServeMux`: Sang Manajer Rute](#httpservmux-sang-manajer-rute)
    *   [`http.ListenAndServe`: Menghidupkan Mesin](#httplistenandserve-menghidupkan-mesin)

4.  [**Bab 3: Routing - Mengarahkan Lalu Lintas**](#4-bab-3-routing---mengarahkan-lalu-lintas)
    *   [Routing Dasar dengan `DefaultServeMux`](#routing-dasar-dengan-defaultservemux)
    *   [Pola URL di `ServeMux`](#pola-url-di-servemux)
    *   [Keterbatasan `ServeMux` dan Kapan Harus Menggunakan Router Pihak Ketiga](#keterbatasan-servemux-dan-kapan-harus-menggunakan-router-pihak-ketiga)
    *   [Contoh dengan Router Populer: `chi`](#contoh-dengan-router-populer-chi)

5.  [**Bab 4: Bekerja dengan Data - Request dan Response**](#5-bab-4-bekerja-dengan-data---request-dan-response)
    *   [Membaca Query Parameters](#membaca-query-parameters)
    *   [Membaca Headers](#membaca-headers)
    *   [Membaca Request Body](#membaca-request-body)
    *   [Menangani JSON API](#menangani-json-api)
    *   [Menangani Formulir (Form Data)](#menangani-formulir-form-data)
    *   [Menangani Upload File (Multipart/Form-Data)](#menangani-upload-file-multipartform-data)
    *   [Mengirim Respons: Status Code, Headers, dan Body](#mengirim-respons-status-code-headers-dan-body)

6.  [**Bab 5: Middleware - Kekuatan Komposisi dalam Aksi**](#6-bab-5-middleware---kekuatan-komposisi-dalam-aksi)
    *   [Apa itu Middleware?](#apa-itu-middleware)
    *   [Pola Middleware Standar di Go](#pola-middleware-standar-di-go)
    *   [Contoh 1: Middleware untuk Logging](#contoh-1-middleware-untuk-logging)
    *   [Contoh 2: Middleware untuk Autentikasi Sederhana](#contoh-2-middleware-untuk-autentikasi-sederhana)
    *   [Merangkai Middleware (Chaining)](#merangkai-middleware-chaining)

7.  [**Bab 6: Topik Lanjutan**](#7-bab-6-topik-lanjutan)
    *   [Konfigurasi Server Tingkat Lanjut dengan `http.Server`](#konfigurasi-server-tingkat-lanjut-dengan-httpserver)
    *   [Mengelola Timeouts](#mengelola-timeouts)
    *   [`context.Context` dalam `net/http`](#contextcontext-dalam-nethttp)
    *   [HTTPS/TLS: Menjalankan Server yang Aman](#httpstls-menjalankan-server-yang-aman)
    *   [Graceful Shutdown: Mematikan Server dengan Anggun](#graceful-shutdown-mematikan-server-dengan-anggun)
    *   [Menyajikan File Statis (Static File Serving)](#menyajikan-file-statis-static-file-serving)
    *   [Menggunakan `html/template` untuk Server-Side Rendering](#menggunakan-htmltemplate-untuk-server-side-rendering)

8.  [**Bab 7: Sisi Klien - Membuat HTTP Request**](#8-bab-7-sisi-klien---membuat-http-request)
    *   [Request Sederhana dengan `http.Get` & `http.Post`](#request-sederhana-dengan-httpget--httppost)
    *   [Kontrol Penuh dengan `http.Client`](#kontrol-penuh-dengan-httpclient)
    *   [Mengatur Timeout pada Klien](#mengatur-timeout-pada-klien)
    *   [Mengirim JSON ke Server Lain](#mengirim-json-ke-server-lain)

9.  [**Bab 8: Praktik Terbaik & Struktur Proyek**](#9-bab-8-praktik-terbaik--struktur-proyek)
    *   [Hindari Global State (`DefaultServeMux` dan `http.HandleFunc`)](#hindari-global-state-defaultservemux-dan-httphandlefunc)
    *   [Dependency Injection pada Handler](#dependency-injection-pada-handler)
    *   [Struktur Proyek yang Dianjurkan](#struktur-proyek-yang-dianjurkan)

10. [**Bab 9: Contoh Proyek Lengkap - API CRUD Sederhana**](#10-bab-9-contoh-proyek-lengkap---api-crud-sederhana)
    *   [Struktur File Proyek](#struktur-file-proyek)
    *   [Definisi Model dan Penyimpanan (In-Memory)](#definisi-model-dan-penyimpanan-in-memory)
    *   [Membuat Handler](#membuat-handler)
    *   [Merakit Semuanya di `main.go`](#merakit-semuanya-di-maingo)
    *   [Menjalankan dan Menguji Proyek](#menjalankan-dan-menguji-proyek)

11. [**Kesimpulan & Langkah Selanjutnya**](#11-kesimpulan--langkah-selanjutnya)

---

## 1. Pendahuluan: Kenapa `net/http`?

Di dunia pengembangan web, kita sering dihadapkan pada pilihan framework yang tak terhitung jumlahnya: Express.js di Node, Django/Flask di Python, Laravel di PHP, Ruby on Rails, dan sebagainya. Framework-framework ini menawarkan banyak kemudahan dan "sihir" untuk mempercepat pengembangan. Go mengambil pendekatan yang berbeda.

Pustaka standar Go, khususnya `net/http`, dirancang untuk menjadi fondasi yang kokoh, bukan sebuah bangunan jadi. Ini memberi Anda kekuatan dan fleksibilitas tanpa mengorbankan pemahaman tentang apa yang sebenarnya terjadi di bawah permukaan.

### Filosofi di Balik `net/http`

*   **Eksplisit dan Jelas**: Kode yang Anda tulis dengan `net/http` sangat mudah dibaca dan dipahami. Tidak ada *magic* atau konfigurasi tersembunyi. Alur request-response sangat transparan.
*   **Komposabilitas**: `net/http` dibangun di atas `interface` sederhana, terutama `http.Handler`. Ini memungkinkan Anda untuk "menyusun" atau "membungkus" fungsionalitas (seperti logging, autentikasi, kompresi) dengan cara yang elegan dan dapat digunakan kembali, sebuah konsep yang dikenal sebagai **middleware**.
*   **Performa**: Go adalah bahasa yang dikompilasi (compiled language) dengan konkurensi sebagai warga kelas satu. Server `net/http` secara default menangani setiap request dalam *goroutine* terpisah, memungkinkannya melayani ribuan request secara bersamaan dengan efisiensi yang sangat tinggi.

### Tanpa Sihir (No Magic)

Ketika Anda menggunakan `net/http`, Anda akan memahami setiap langkah:
1.  Sebuah request masuk.
2.  Router (ServeMux) mencocokkan URL request dengan sebuah handler.
3.  Handler menerima `ResponseWriter` (untuk menulis respons) dan `Request` (untuk membaca detail permintaan).
4.  Anda secara eksplisit membaca data dari request dan menulis data ke response writer.

Tidak ada *dependency injection container* yang otomatis, tidak ada *Object-Relational Mapper (ORM)* yang terintegrasi, tidak ada sistem template yang dipaksakan. Anda memiliki kebebasan penuh untuk memilih dan merakit komponen yang Anda butuhkan. Ini mungkin terasa lebih banyak pekerjaan di awal, tetapi akan memberikan pemahaman yang lebih dalam dan kontrol yang lebih besar atas aplikasi Anda.

### Kinerja Luar Biasa

Server `net/http` Go terkenal sangat cepat. Karena setiap koneksi ditangani dalam goroutine yang ringan, ia dapat dengan mudah menangani "C10k problem" (menangani 10.000 koneksi bersamaan) pada perangkat keras modern. Ini menjadikannya pilihan yang sangat baik untuk membangun API berkinerja tinggi dan layanan mikro.

Mari kita mulai perjalanan kita dengan kode pertama.

---

## 2. Bab 1: Server HTTP Pertama Anda - "Hello, World!"

Tidak ada cara yang lebih baik untuk belajar selain dengan langsung mencoba. Mari kita buat server web paling sederhana yang bisa dibuat dengan Go.

### Kode Minimalis

Buat sebuah file bernama `main.go` dan tulis kode berikut:

```go
package main

import (
	"fmt"
	"log"
	"net/http"
)

// helloHandler adalah fungsi yang akan menangani request ke endpoint "/hello"
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Menulis string "Hello, World!" ke body respons
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	// Mendaftarkan helloHandler untuk menangani semua request ke path URL "/hello"
	http.HandleFunc("/hello", helloHandler)

	// Mulai server di port 8080.
	// Jika terjadi error (misalnya port sudah digunakan), program akan berhenti dan mencatat error.
	log.Println("Server starting on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
```

### Membedah Kode

Mari kita pecah kode di atas baris per baris.

1.  **`package main`**: Mendeklarasikan bahwa ini adalah program yang dapat dieksekusi.
2.  **`import (...)`**: Mengimpor paket-paket yang kita butuhkan.
    *   `fmt`: Untuk memformat output, dalam kasus ini untuk menulis ke `ResponseWriter`.
    *   `log`: Untuk mencatat pesan, seperti status server atau error.
    *   `net/http`: Paket inti yang berisi semua fungsionalitas server dan klien HTTP.
3.  **`func helloHandler(w http.ResponseWriter, r *http.Request)`**: Ini adalah *fungsi handler*.
    *   Fungsi ini memiliki tanda tangan (signature) yang spesifik: ia menerima dua argumen.
    *   `w http.ResponseWriter`: Ini adalah sebuah `interface` yang kita gunakan untuk **menulis respons** kembali ke klien. Kita bisa mengatur status code, header, dan menulis body respons melaluinya.
    *   `r *http.Request`: Ini adalah sebuah *pointer* ke *struct* `http.Request`. Objek ini berisi **semua informasi tentang request** yang masuk dari klien: URL, method (GET, POST, dll.), header, body, dan lainnya.
    *   `fmt.Fprintf(w, "Hello, World!")`: Di sini kita menggunakan fungsi dari paket `fmt` untuk menulis string "Hello, World!" ke `w` (si penulis respons).
4.  **`func main()`**: Titik masuk program kita.
    *   **`http.HandleFunc("/hello", helloHandler)`**: Ini adalah bagian krusial. Kita memberitahu paket `net/http` untuk mendaftarkan `helloHandler`. Setiap kali ada request masuk dengan path URL `/hello`, panggil fungsi `helloHandler`.
    *   **`http.ListenAndServe(":8080", nil)`**: Ini adalah fungsi yang memulai server HTTP.
        *   Argumen pertama, `":8080"`, adalah alamat jaringan tempat server akan mendengarkan. Tanda titik dua di depan berarti mendengarkan di semua antarmuka jaringan pada port 8080.
        *   Argumen kedua adalah `handler`. Di sini kita memberikan `nil`. Ketika `nil` diberikan, Go akan menggunakan *handler default* yang disebut `DefaultServeMux`. `http.HandleFunc` yang kita panggil sebelumnya mendaftarkan handler kita ke `DefaultServeMux` ini. Kita akan membahas ini lebih detail nanti.
    *   **`log.Fatal(...)`**: `http.ListenAndServe` adalah fungsi yang *memblokir*. Ia akan terus berjalan sampai server dimatikan atau terjadi error. Jika terjadi error saat memulai (misalnya, port 8080 sudah dipakai oleh aplikasi lain), ia akan mengembalikan error. Kode ini akan menangkap error tersebut, mencatatnya, dan menghentikan program.

### Menjalankan Server

1.  Buka terminal Anda.
2.  Navigasi ke direktori tempat Anda menyimpan `main.go`.
3.  Jalankan perintah: `go run main.go`.
4.  Anda akan melihat output: `Server starting on port 8080...`
5.  Sekarang, buka browser web Anda dan kunjungi `http://localhost:8080/hello`.
6.  Anda akan melihat teks "Hello, World!" di browser Anda.

Selamat! Anda baru saja membuat dan menjalankan server web fungsional pertama Anda dengan Go.

---

## 3. Bab 2: Konsep Inti - Memahami Blok Bangunan `net/http`

Sekarang setelah kita melihat contoh praktis, mari kita selami lebih dalam komponen-komponen utama yang membuat semuanya bekerja. Memahami konsep-konsep ini adalah kunci untuk menjadi mahir dengan `net/http`.

### `http.ResponseWriter`: Sang Penulis Respons

Bayangkan `ResponseWriter` (yang biasa disingkat `w`) sebagai selembar kertas kosong dan sebuah pena yang diberikan kepada Anda setiap kali ada surat (request) masuk. Tugas Anda adalah menulis balasan di kertas itu.

`ResponseWriter` adalah sebuah `interface` yang memiliki beberapa method penting:

*   **`Header() http.Header`**: Mengembalikan sebuah `map` tempat Anda bisa mengatur header respons. Header harus diatur *sebelum* Anda menulis body.
    ```go
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("X-Custom-Header", "Nilai Kustom")
    ```
*   **`WriteHeader(statusCode int)`**: Mengirim header respons HTTP bersama dengan status code yang diberikan. Misalnya, `http.StatusOK` (200), `http.StatusNotFound` (404), `http.StatusInternalServerError` (500). Jika Anda tidak memanggil ini secara eksplisit, Go akan secara otomatis mengirim status `200 OK` saat Anda pertama kali memanggil `Write`.
    ```go
    w.WriteHeader(http.StatusCreated) // Mengirim status 201 Created
    ```
*   **`Write(data []byte) (int, error)`**: Menulis sepotong data (`[]byte`) ke body respons. Ini adalah method inti dari `interface io.Writer` yang di-embed oleh `ResponseWriter`, yang membuatnya sangat serbaguna. Banyak fungsi di Go yang bisa menulis ke `io.Writer`, termasuk `fmt.Fprintf`, `json.NewEncoder`, dan `template.Execute`.

**Penting:** Urutan operasinya adalah:
1.  Set Headers (opsional, tapi sering dilakukan).
2.  Panggil `WriteHeader` (opsional, akan default ke 200 jika tidak dipanggil).
3.  Panggil `Write` untuk mengirim body.

Setelah `WriteHeader` atau `Write` pertama kali dipanggil, Anda tidak bisa lagi mengubah header.

### `*http.Request`: Sang Pembawa Pesan

Jika `ResponseWriter` adalah alat untuk membalas, `*http.Request` (biasa disingkat `r`) adalah pesan itu sendiri. Ini adalah *struct* yang sangat kaya dan berisi semua yang perlu Anda ketahui tentang permintaan klien.

Beberapa field dan method yang paling sering digunakan:

*   **`r.Method` (string)**: Metode HTTP yang digunakan, mis. `"GET"`, `"POST"`, `"PUT"`, `"DELETE"`.
*   **`r.URL` (*url.URL)**: Sebuah struct yang berisi informasi terurai tentang URL request.
    *   `r.URL.Path`: Bagian path dari URL, mis. `/users/123`.
    *   `r.URL.Query()`: Mengembalikan nilai-nilai query string sebagai `map[string][]string`. Mis. untuk `/search?q=golang&lang=id`, `r.URL.Query().Get("q")` akan mengembalikan `"golang"`.
*   **`r.Header` (http.Header)**: Sebuah `map` yang berisi semua header yang dikirim oleh klien.
    *   `r.Header.Get("Content-Type")` akan mengembalikan nilai header `Content-Type`.
*   **`r.Body` (io.ReadCloser)**: Body dari request. Ini adalah `io.Reader`, artinya Anda dapat membacanya sebagai aliran (stream) data. Ini sangat efisien untuk menangani body yang besar. Penting untuk diingat: **Anda hanya bisa membaca `r.Body` sekali**, dan Anda bertanggung jawab untuk menutupnya dengan `r.Body.Close()`.
*   **`r.ParseForm()` dan `r.ParseMultipartForm()`**: Method untuk mengurai data form. Setelah memanggil ini, Anda bisa mengakses nilainya melalui `r.Form` atau `r.PostForm`.
*   **`r.Context()` (context.Context)**: Mengembalikan konteks yang terkait dengan request ini. Konteks sangat berguna untuk menangani pembatalan, timeout, dan meneruskan nilai-nilai request-scoped. Kita akan membahas ini lebih lanjut di topik lanjutan.

### `http.Handler`: Jantung dari Server Go

Ini adalah `interface` paling fundamental dalam `net/http`. Definisi `interface`-nya sangat sederhana:

```go
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
```

Apapun—benar-benar *apapun*—yang memiliki method bernama `ServeHTTP` dengan tanda tangan `(http.ResponseWriter, *http.Request)` secara implisit memenuhi `interface http.Handler`.

Ini adalah ide yang sangat kuat. Ini berarti Anda bisa membuat `struct` Anda sendiri yang bertindak sebagai handler. Ini sangat berguna untuk mengelola state atau dependensi (seperti koneksi database).

Contoh:

```go
type MyServer struct {
    message string
}

func (s *MyServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, s.message)
}

func main() {
    // Membuat instance dari handler kita
    myServerHandler := &MyServer{message: "Halo dari struct!"}

    // Mendaftarkan handler kita ke path "/"
    http.Handle("/", myServerHandler)

    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

Perhatikan kita menggunakan `http.Handle` bukan `http.HandleFunc`. `http.Handle` menerima sesuatu yang memenuhi `interface http.Handler`.

### `http.HandlerFunc`: Adaptor Cerdas

Anda mungkin bertanya, "Tadi di 'Hello World', fungsi `helloHandler` saya tidak punya method `ServeHTTP`, tapi kok bisa bekerja?".

Di sinilah `http.HandlerFunc` berperan. Ini adalah sebuah tipe khusus di Go:

```go
type HandlerFunc func(ResponseWriter, *Request)
```

Ini adalah tipe fungsi yang tanda tangannya persis sama dengan method `ServeHTTP`. Tipe ini kemudian memiliki method `ServeHTTP` sendiri!

```go
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
    f(w, r)
}
```

Ini adalah trik yang brilian. Ia "mengangkat" atau "mengadaptasi" sebuah fungsi biasa menjadi sebuah `http.Handler`.

Jadi, ketika Anda melakukan:

```go
http.HandleFunc("/hello", helloHandler)
```

Yang sebenarnya terjadi di balik layar adalah sesuatu seperti ini:

```go
// 1. helloHandler adalah func(w http.ResponseWriter, r *http.Request)
// 2. Go mengubahnya menjadi tipe HandlerFunc
handlerAsHandlerFunc := http.HandlerFunc(helloHandler)
// 3. Sekarang handlerAsHandlerFunc memiliki method ServeHTTP dan memenuhi interface http.Handler
// 4. http.HandleFunc mendaftarkannya menggunakan http.Handle
http.Handle("/hello", handlerAsHandlerFunc)
```

`http.HandleFunc` hanyalah sebuah fungsi pembantu yang membuat hidup kita lebih mudah.

### `http.ServeMux`: Sang Manajer Rute

`ServeMux` (atau multiplexer) adalah router HTTP. Tugasnya adalah melihat URL dari request yang masuk dan memutuskan handler mana yang harus dieksekusi.

Ketika kita menggunakan `http.HandleFunc` atau `http.Handle`, kita sebenarnya mendaftarkan rute ke sebuah instance global dari `ServeMux` yang disebut `DefaultServeMux`.

Anda juga bisa membuat `ServeMux` Anda sendiri untuk kontrol yang lebih baik (dan ini adalah praktik yang lebih baik untuk menghindari state global):

```go
func main() {
    // Membuat ServeMux baru, bukan menggunakan yang default
    mux := http.NewServeMux()

    // Mendaftarkan handler ke mux yang baru kita buat
    mux.HandleFunc("/hello", helloHandler)
    // Anda juga bisa mendaftarkan handler lain
    mux.HandleFunc("/goodbye", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Goodbye!")
    })

    log.Println("Server starting on port 8080...")
    // Sekarang, kita berikan mux kita sendiri ke ListenAndServe, bukan nil
    err := http.ListenAndServe(":8080", mux)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
```

Menggunakan `ServeMux` Anda sendiri lebih disukai dalam aplikasi yang lebih besar karena membuat kode lebih terisolasi, lebih mudah diuji, dan menghindari potensi konflik jika Anda menggunakan library pihak ketiga yang mungkin juga mendaftarkan rute ke `DefaultServeMux`.

### `http.ListenAndServe`: Menghidupkan Mesin

Fungsi ini melakukan dua hal:
1.  **Listen**: Membuka soket jaringan pada alamat dan port yang ditentukan dan menunggu koneksi TCP masuk.
2.  **Serve**: Ketika sebuah koneksi diterima, ia memulai *loop* tak terbatas untuk:
    a. Menerima request HTTP dari koneksi tersebut.
    b. Membuat *goroutine* baru untuk menangani request tersebut.
    c. Di dalam goroutine baru, ia memanggil `ServeHTTP` dari handler yang Anda berikan (atau `DefaultServeMux` jika `nil`).
    d. Kembali menunggu request berikutnya pada koneksi yang sama (untuk HTTP keep-alive).

Fakta bahwa setiap request ditangani dalam goroutine-nya sendiri adalah alasan utama mengapa server HTTP Go sangat konkuren dan berperforma tinggi.

---

## 4. Bab 3: Routing - Mengarahkan Lalu Lintas

Routing adalah proses memetakan URL request ke kode (handler) yang akan menanganinya. `net/http` menyediakan `http.ServeMux` untuk ini, yang cukup untuk aplikasi sederhana, tetapi memiliki beberapa keterbatasan.

### Routing Dasar dengan `DefaultServeMux`

Seperti yang telah kita lihat, Anda menggunakan `http.HandleFunc` atau `http.Handle` untuk mendaftarkan rute.

```go
http.HandleFunc("/", rootHandler) // Menangani root path
http.HandleFunc("/users", usersHandler) // Menangani /users
http.HandleFunc("/users/create", createUserHandler) // Menangani /users/create
```

### Pola URL di `ServeMux`

`ServeMux` mencocokkan URL request dengan pola yang terdaftar. Aturan pencocokannya adalah **pencocokan prefiks terpanjang (longest prefix match)**.

*   **Pencocokan Tepat**: Pola tanpa garis miring di akhir (`/users`) akan cocok hanya dengan path tersebut.
*   **Pencocokan Sub-pohon**: Pola dengan garis miring di akhir (`/static/`) akan cocok dengan path tersebut dan semua path di bawahnya. Misalnya, `/static/` akan cocok dengan `/static/css/style.css` dan `/static/js/app.js`.
*   **Pencocokan Root**: Pola `/` akan cocok dengan semua path yang tidak cocok dengan pola lain yang lebih spesifik.

Contoh:

```go
mux := http.NewServeMux()
mux.HandleFunc("/api/v1/users/", usersAPIHandler) // Pola A
mux.HandleFunc("/api/v1/", apiRootHandler)     // Pola B
mux.HandleFunc("/", catchAllHandler)          // Pola C

// Request ke "/api/v1/users/123" akan cocok dengan Pola A (paling panjang)
// Request ke "/api/v1/products" akan cocok dengan Pola B
// Request ke "/about" akan cocok dengan Pola C
```

### Keterbatasan `ServeMux` dan Kapan Harus Menggunakan Router Pihak Ketiga

`ServeMux` bawaan sengaja dibuat sederhana. Ia memiliki beberapa keterbatasan yang signifikan untuk aplikasi web modern:

1.  **Tidak bisa menangani parameter di path (Path Parameters)**: Anda tidak bisa membuat rute seperti `/users/{id}` untuk menangkap ID pengguna. Anda harus melakukannya secara manual dengan mengurai string path, yang merepotkan dan rawan kesalahan.
2.  **Tidak bisa membedakan berdasarkan Metode HTTP**: Rute `/users` akan menangani `GET /users`, `POST /users`, `PUT /users`, dll. di handler yang sama. Anda harus menggunakan `if r.Method == "POST"` di dalam handler Anda, yang bisa membuat handler menjadi berantakan.
3.  **Tidak ada dukungan untuk middleware per-rute**: Anda tidak bisa dengan mudah menerapkan middleware (mis. autentikasi) hanya untuk sekelompok rute tertentu (mis. `/admin/*`).

Untuk alasan ini, hampir semua aplikasi Go yang serius menggunakan router pihak ketiga. Router-router ini dibangun di atas `net/http` (mereka semua mengimplementasikan `interface http.Handler`), jadi semua yang telah Anda pelajari tetap berlaku.

Beberapa router populer:
*   [**chi**](https://github.com/go-chi/chi): Ringan, cepat, dan sangat idiomatis dengan Go. Pilihan yang sangat baik.
*   [**gorilla/mux**](https://github.com/gorilla/mux): Salah satu yang tertua dan paling banyak digunakan. Sangat kuat dan kaya fitur.
*   [**httprouter**](https://github.com/julienschmidt/httprouter): Dikenal karena performanya yang sangat tinggi.

### Contoh dengan Router Populer: `chi`

Mari kita lihat betapa lebih ekspresifnya routing dengan `chi`.

Pertama, instal `chi`:
`go get -u github.com/go-chi/chi/v5`

Sekarang, mari kita tulis ulang server kita dengan `chi`:

```go
package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/go-chi/chi/v5" // Impor chi
)

func main() {
	// Buat router baru dari chi
	r := chi.NewRouter()

	// Mendaftarkan rute dengan metode HTTP yang spesifik
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the homepage!"))
	})

	// Rute dengan parameter path
	r.Get("/users/{userID}", func(w http.ResponseWriter, r *http.Request) {
		// chi.URLParam akan mengambil nilai dari placeholder {userID}
		userID := chi.URLParam(r, "userID")
		fmt.Fprintf(w, "Hello, user %s!", userID)
	})

	// Pengelompokan rute (grouping)
	r.Route("/articles", func(r chi.Router) {
		r.Get("/", listArticles)       // GET /articles
		r.Post("/", createArticle)     // POST /articles

		// Sub-router
		r.Route("/{articleID}", func(r chi.Router) {
			r.Get("/", getArticle)       // GET /articles/123
			r.Put("/", updateArticle)    // PUT /articles/123
			r.Delete("/", deleteArticle) // DELETE /articles/123
		})
	})

	log.Println("Server starting on port 8080...")
	// Gunakan router chi sebagai handler utama
	http.ListenAndServe(":8080", r)
}

// Dummy handlers untuk contoh di atas
func listArticles(w http.ResponseWriter, r *http.Request)    { w.Write([]byte("List of articles")) }
func createArticle(w http.ResponseWriter, r *http.Request)   { w.Write([]byte("Create a new article")) }
func getArticle(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "articleID")
	fmt.Fprintf(w, "Getting article %s", id)
}
func updateArticle(w http.ResponseWriter, r *http.Request)   { w.Write([]byte("Update an article")) }
func deleteArticle(w http.ResponseWriter, r *http.Request)   { w.Write([]byte("Delete an article")) }
```

Lihat betapa bersih dan terstrukturnya kode routing kita sekarang? Kita bisa mendefinisikan rute berdasarkan metode HTTP dan dengan mudah mengekstrak parameter dari URL. Ini adalah alasan kuat untuk menggunakan router pihak ketiga di proyek Anda.

---

## 5. Bab 4: Bekerja dengan Data - Request dan Response

Aplikasi web pada dasarnya adalah tentang menerima data, memprosesnya, dan mengirimkan data kembali. Mari kita jelajahi cara melakukannya dengan `net/http`.

### Membaca Query Parameters

Query parameters adalah bagian dari URL setelah tanda tanya (`?`), misalnya `/search?q=golang&page=2`.

```go
func searchHandler(w http.ResponseWriter, r *http.Request) {
	// Ambil semua query params
	query := r.URL.Query()

	// Ambil nilai spesifik. Get() mengembalikan string kosong jika tidak ditemukan.
	q := query.Get("q")
	// Anda juga bisa memeriksa keberadaannya
	page, ok := query["page"] // query adalah map[string][]string

	if !ok || len(page[0]) < 1 {
		fmt.Fprintf(w, "Query: %s. Halaman tidak dispesifikkan.", q)
		return
	}

	fmt.Fprintf(w, "Anda mencari '%s' di halaman %s.", q, page[0])
}

// di main.go:
// mux.HandleFunc("/search", searchHandler)
// Akses: http://localhost:8080/search?q=go+net/http&page=1
```

### Membaca Headers

Membaca header request sama mudahnya.

```go
func headerHandler(w http.ResponseWriter, r *http.Request) {
	// Membaca header 'User-Agent'
	userAgent := r.Header.Get("User-Agent")

	// Membaca header 'Accept-Language'
	acceptLang := r.Header.Get("Accept-Language")

	fmt.Fprintf(w, "User-Agent: %s\n", userAgent)
	fmt.Fprintf(w, "Accept-Language: %s\n", acceptLang)
}
```

### Membaca Request Body

Untuk metode seperti `POST` atau `PUT`, klien sering mengirim data di body request. `r.Body` adalah `io.ReadCloser`. Anda harus membacanya dan kemudian menutupnya.

```go
import "io/ioutil" // atau "io" di Go 1.16+

func echoHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        w.WriteHeader(http.StatusMethodNotAllowed)
        return
    }

    // Pastikan body ditutup setelah fungsi selesai
    defer r.Body.Close()

    // Baca seluruh body
    // Untuk Go 1.16+, gunakan io.ReadAll
    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "Gagal membaca body", http.StatusInternalServerError)
        return
    }

    // Tulis kembali body yang sama sebagai respons
    w.Write(body)
}
```

**Peringatan:** Membaca seluruh body ke memori (`ioutil.ReadAll`) bisa berbahaya jika klien mengirim file yang sangat besar. Untuk kasus seperti itu, lebih baik memproses body sebagai stream.

### Menangani JSON API

JSON adalah format data paling umum untuk API modern. Paket `encoding/json` di Go membuatnya sangat mudah untuk bekerja dengan JSON.

**1. Menerima (Decoding) JSON dari Request**

Misalkan klien mengirim JSON seperti ini: `{"name": "John Doe", "age": 30}`

```go
import "encoding/json"

// Definisikan struct yang cocok dengan struktur JSON
type User struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
    var user User

    // Buat decoder yang membaca langsung dari body request
    // Ini lebih efisien daripada membaca semua ke memori dulu
    decoder := json.NewDecoder(r.Body)
    err := decoder.Decode(&user)
    if err != nil {
        http.Error(w, "JSON tidak valid", http.StatusBadRequest)
        return
    }

    // Sekarang Anda punya data di struct `user`
    log.Printf("Menerima user baru: Name=%s, Age=%d\n", user.Name, user.Age)

    // ... (logika untuk menyimpan user ke database) ...

    fmt.Fprintf(w, "User %s berhasil dibuat!", user.Name)
}
```

**2. Mengirim (Encoding) JSON sebagai Response**

Sekarang, mari kita kirim respons dalam format JSON.

```go
type ResponseData struct {
    Status  string `json:"status"`
    Message string `json:"message"`
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
    data := ResponseData{
        Status:  "OK",
        Message: "Server berjalan dengan baik",
    }

    // 1. Set header Content-Type. Ini SANGAT PENTING!
    w.Header().Set("Content-Type", "application/json")

    // 2. (Opsional) Set status code jika bukan 200 OK
    w.WriteHeader(http.StatusOK)

    // 3. Buat encoder yang menulis langsung ke ResponseWriter
    encoder := json.NewEncoder(w)
    err := encoder.Encode(data)
    if err != nil {
        // Sulit untuk mengirim error HTTP di sini karena header mungkin sudah terkirim
        log.Printf("Gagal meng-encode JSON: %v", err)
    }
}
```

### Menangani Formulir (Form Data)

Ini biasanya digunakan untuk submit form dari halaman web tradisional (`application/x-www-form-urlencoded`).

```go
func formHandler(w http.ResponseWriter, r *http.Request) {
    // r.ParseForm() harus dipanggil sebelum mengakses data form
    if err := r.ParseForm(); err != nil {
        http.Error(w, "Gagal parse form", http.StatusBadRequest)
        return
    }

    // Ambil nilai dari form
    name := r.PostFormValue("name") // Cara mudah untuk satu nilai
    email := r.PostForm.Get("email") // Cara lain

    fmt.Fprintf(w, "Menerima form: Name=%s, Email=%s", name, email)
}
```

### Menangani Upload File (Multipart/Form-Data)

Ketika form berisi input `type="file"`, browser akan mengirimnya sebagai `multipart/form-data`.

```go
import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	// 1. Parse multipart form. 10 << 20 berarti maks. 10MB memori digunakan.
	// Sisanya akan disimpan ke file temporer di disk.
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "File terlalu besar", http.StatusBadRequest)
		return
	}

	// 2. Ambil file dari form. "myFile" adalah nama dari input field <input type="file" name="myFile">
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "Gagal mengambil file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	log.Printf("Uploaded File: %+v\n", handler.Filename)
	log.Printf("File Size: %+v\n", handler.Size)
	log.Printf("MIME Header: %+v\n", handler.Header)

	// 3. Buat file di server untuk menyimpan file yang di-upload
	dst, err := os.Create(handler.Filename)
	if err != nil {
		http.Error(w, "Gagal membuat file di server", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	// 4. Salin isi file yang di-upload ke file tujuan
	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, "Gagal menyimpan file", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "File %s berhasil di-upload!", handler.Filename)
}
```

### Mengirim Respons: Status Code, Headers, dan Body

Kita sudah melihat potongan-potongan ini, tapi mari kita gabungkan. Siklus hidup respons yang lengkap adalah:

```go
func completeResponseHandler(w http.ResponseWriter, r *http.Request) {
    // 1. Logika bisnis Anda
    // ...
    // Misalkan logika bisnis menghasilkan data ini
    responseData := map[string]string{"message": "Operasi berhasil"}

    // 2. Atur Headers
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("X-Powered-By", "Go")

    // 3. Tulis Status Code HTTP
    // Ini harus dilakukan SETELAH mengatur header dan SEBELUM menulis body.
    w.WriteHeader(http.StatusCreated) // 201 Created

    // 4. Tulis Body
    // Panggilan Write pertama secara implisit akan memanggil WriteHeader(http.StatusOK) jika belum dipanggil.
    json.NewEncoder(w).Encode(responseData)
}
```

---

## 6. Bab 5: Middleware - Kekuatan Komposisi dalam Aksi

Middleware (atau terkadang disebut "chaining handlers") adalah salah satu pola paling kuat dalam pengembangan web di Go.

### Apa itu Middleware?

Middleware adalah sebuah bagian kode yang "membungkus" sebuah handler HTTP. Ia bisa melakukan pekerjaan **sebelum** atau **sesudah** handler utama dieksekusi.

Kegunaan umum middleware:
*   **Logging**: Mencatat detail setiap request (metode, URL, waktu eksekusi).
*   **Autentikasi/Autorisasi**: Memeriksa token atau sesi sebelum mengizinkan akses ke handler.
*   **Kompresi**: Meng-compress body respons (mis. dengan Gzip).
*   **CORS**: Menambahkan header Cross-Origin Resource Sharing.
*   **Recovery**: Menangkap `panic` di dalam handler agar server tidak crash.
*   **Menambahkan data ke konteks**: Menambahkan ID request atau informasi user ke `r.Context()`.

### Pola Middleware Standar di Go

Middleware di Go biasanya adalah sebuah fungsi yang:
1.  Menerima `http.Handler` sebagai argumen (handler "selanjutnya" dalam rantai).
2.  Mengembalikan `http.Handler` baru (handler yang sudah "dibungkus").

Tanda tangan umumnya terlihat seperti ini:

`func(next http.Handler) http.Handler`

Handler yang dikembalikan biasanya berupa `http.HandlerFunc` yang berisi logika middleware.

Mari kita lihat contohnya agar lebih jelas.

### Contoh 1: Middleware untuk Logging

Kita ingin mencatat setiap request yang masuk, beserta waktu yang dibutuhkan untuk memprosesnya.

```go
import (
    "log"
    "net/http"
    "time"
)

// loggingMiddleware adalah middleware kita
func loggingMiddleware(next http.Handler) http.Handler {
    // Mengembalikan http.HandlerFunc yang akan menjadi handler baru
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // --- Logika yang dieksekusi SEBELUM handler utama ---
        start := time.Now()
        log.Printf("Mulai: %s %s", r.Method, r.URL.Path)

        // Panggil handler selanjutnya dalam rantai.
        // Ini adalah titik di mana kontrol diserahkan ke middleware berikutnya
        // atau ke handler utama.
        next.ServeHTTP(w, r)

        // --- Logika yang dieksekusi SETELAH handler utama selesai ---
        log.Printf("Selesai: %s %s dalam %v", r.Method, r.URL.Path, time.Since(start))
    })
}

// Handler utama kita
func indexHandler(w http.ResponseWriter, r *http.Request) {
    time.Sleep(100 * time.Millisecond) // Simulasi pekerjaan
    w.Write([]byte("Ini adalah halaman utama."))
}

func main() {
    // Buat handler utama
    finalHandler := http.HandlerFunc(indexHandler)

    // Bungkus handler utama dengan middleware kita
    wrappedHandler := loggingMiddleware(finalHandler)

    mux := http.NewServeMux()
    // Daftarkan handler yang sudah dibungkus
    mux.Handle("/", wrappedHandler)

    log.Println("Server starting on port 8080...")
    http.ListenAndServe(":8080", mux)
}
```

Jika Anda menjalankan ini dan mengakses `http://localhost:8080/`, Anda akan melihat log seperti:

```
2023/10/27 10:30:00 Mulai: GET /
2023/10/27 10:30:00 Selesai: GET / dalam 100.5ms
```

### Contoh 2: Middleware untuk Autentikasi Sederhana

Middleware ini akan memeriksa header `Authorization`. Jika tidak ada atau salah, ia akan menghentikan request dan mengirim respons error. Jika benar, ia akan melanjutkan ke handler utama.

```go
func authMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Ambil token dari header
        token := r.Header.Get("Authorization")

        // Logika autentikasi yang sangat sederhana
        if token != "Bearer my-secret-token" {
            // Jika tidak valid, kirim error dan JANGAN panggil next.ServeHTTP
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return // Penting untuk return di sini!
        }

        // Jika valid, lanjutkan ke handler berikutnya
        log.Println("Autentikasi berhasil")
        next.ServeHTTP(w, r)
    })
}

// Handler yang dilindungi
func secretHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Ini adalah area rahasia!"))
}

func main() {
    // ... (kode main dari contoh logging)
    secretH := http.HandlerFunc(secretHandler)

    // Terapkan middleware auth HANYA untuk secretHandler
    protectedHandler := authMiddleware(secretH)

    mux.Handle("/secret", protectedHandler)
    // ... (sisanya sama)
}
```

Untuk menguji ini, Anda bisa menggunakan `curl`:
*   `curl http://localhost:8080/secret` -> Akan mengembalikan `Unauthorized`
*   `curl -H "Authorization: Bearer my-secret-token" http://localhost:8080/secret` -> Akan mengembalikan `Ini adalah area rahasia!`

### Merangkai Middleware (Chaining)

Apa yang terjadi jika Anda ingin menerapkan `loggingMiddleware` DAN `authMiddleware` ke satu handler? Anda tinggal membungkusnya secara berurutan.

```go
func main() {
    mux := http.NewServeMux()

    // Handler asli
    secretH := http.HandlerFunc(secretHandler)

    // Rantai middleware:
    // Request -> loggingMiddleware -> authMiddleware -> secretHandler
    // urutan dari luar ke dalam
    finalHandler := loggingMiddleware(authMiddleware(secretH))

    mux.Handle("/secret", finalHandler)

    log.Println("Server starting on port 8080...")
    http.ListenAndServe(":8080", mux)
}
```

Pola ini sangat kuat karena setiap middleware fokus pada satu tugas saja (Prinsip Tanggung Jawab Tunggal) dan dapat digabungkan dalam urutan apapun yang Anda butuhkan. Banyak library router (seperti `chi`) menyediakan cara yang lebih mudah untuk merangkai middleware.

---

## 7. Bab 6: Topik Lanjutan

Setelah menguasai dasar-dasar, mari kita lihat beberapa fitur dan pola yang lebih canggih untuk membangun aplikasi production-grade.

### Konfigurasi Server Tingkat Lanjut dengan `http.Server`

`http.ListenAndServe` adalah fungsi pembantu yang nyaman. Namun, untuk aplikasi produksi, Anda memerlukan kontrol lebih. Ini dicapai dengan mengkonfigurasi `struct http.Server` secara manual.

```go
func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello from a custom server!"))
    })

    // Buat instance http.Server
    server := &http.Server{
        Addr:    ":8080",      // Alamat server
        Handler: mux,          // Handler utama (router Anda)
        // Pengaturan timeout sangat penting untuk server produksi!
        ReadTimeout:  5 * time.Second,
        WriteTimeout: 10 * time.Second,
        IdleTimeout:  120 * time.Second,
    }

    log.Println("Server starting on port 8080...")
    // Panggil ListenAndServe dari struct Server, bukan dari paket http
    err := server.ListenAndServe()
    if err != nil {
        log.Fatal("Server error: ", err)
    }
}
```

### Mengelola Timeouts

Timeout sangat penting untuk mencegah server Anda kehabisan sumber daya karena klien yang lambat atau serangan jahat (seperti Slowloris).

*   **`ReadTimeout`**: Waktu maksimum untuk membaca seluruh request, termasuk body.
*   **`WriteTimeout`**: Waktu maksimum untuk menulis respons.
*   **`IdleTimeout`**: Waktu maksimum untuk menunggu request berikutnya pada koneksi keep-alive.

Mengatur timeout ini adalah praktik keamanan dan keandalan yang fundamental.

### `context.Context` dalam `net/http`

Sejak Go 1.7, setiap `*http.Request` memiliki sebuah `context.Context` yang dapat diakses melalui `r.Context()`.

Konteks memiliki dua tujuan utama dalam server HTTP:

1.  **Sinyal Pembatalan (Cancellation Signal)**: Jika klien menutup koneksi di tengah-tengah request, konteks akan dibatalkan. Kode Anda bisa memeriksa pembatalan ini untuk menghentikan pekerjaan yang tidak lagi diperlukan (misalnya, query database yang lama).

    ```go
    func longRunningTaskHandler(w http.ResponseWriter, r *http.Request) {
        ctx := r.Context()
        log.Println("Memulai tugas panjang...")

        select {
        case <-time.After(5 * time.Second):
            // Tugas selesai
            w.Write([]byte("Tugas selesai!"))
            log.Println("Tugas selesai tanpa gangguan.")
        case <-ctx.Done():
            // Klien membatalkan request (mis. menutup browser)
            err := ctx.Err()
            log.Printf("Request dibatalkan: %v", err)
            // Tidak perlu menulis respons, koneksi sudah putus
        }
    }
    ```

2.  **Membawa Nilai Request-Scoped**: Middleware dapat menambahkan nilai ke konteks, dan handler di hilir dapat membacanya. Ini adalah cara yang aman untuk meneruskan data antar lapisan tanpa mencemari tanda tangan fungsi handler.

    *Contoh Middleware yang menambahkan user ID ke konteks:*
    ```go
    type contextKey string
    const userKey contextKey = "userID"

    func contextMiddleware(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            // (Asumsikan kita sudah mendapatkan userID dari token atau sesi)
            userID := "user-123"

            // Buat konteks baru dengan nilai userID
            ctx := context.WithValue(r.Context(), userKey, userID)

            // Buat request baru dengan konteks yang sudah diperbarui
            // dan teruskan ke handler berikutnya
            next.ServeHTTP(w, r.WithContext(ctx))
        })
    }

    *Handler yang menggunakan nilai dari konteks:*
    ```go
    func profileHandler(w http.ResponseWriter, r *http.Request) {
        // Ambil userID dari konteks
        userID, ok := r.Context().Value(userKey).(string)
        if !ok {
            http.Error(w, "Gagal mendapatkan userID", http.StatusInternalServerError)
            return
        }

        fmt.Fprintf(w, "Ini adalah profil untuk user %s", userID)
    }
    ```

### HTTPS/TLS: Menjalankan Server yang Aman

Menjalankan server melalui HTTPS sangat mudah. Anda hanya perlu file sertifikat dan file kunci privat.

```go
// Di dalam main()
// ...
// Ganti ListenAndServe dengan ListenAndServeTLS
err := http.ListenAndServeTLS(":8443", "cert.pem", "key.pem", mux)
if err != nil {
    log.Fatal("ListenAndServeTLS: ", err)
}
```
Port standar untuk HTTPS adalah 443. Saya menggunakan 8443 untuk contoh agar tidak memerlukan hak akses root.

Untuk pengembangan lokal, Anda bisa membuat sertifikat self-signed dengan `openssl` atau dengan paket Go `crypto/tls`.

### Graceful Shutdown: Mematikan Server dengan Anggun

Ketika Anda perlu me-restart server (misalnya untuk deployment baru), Anda tidak ingin mematikan koneksi yang sedang aktif secara tiba-tiba. *Graceful shutdown* adalah proses di mana server:
1.  Berhenti menerima koneksi baru.
2.  Menunggu semua request yang sedang berjalan untuk selesai.
3.  Baru kemudian mematikan dirinya.

Ini dicapai dengan menangani sinyal sistem operasi (seperti `SIGINT` yang dikirim saat Anda menekan `Ctrl+C`).

```go
import (
    "context"
    "os"
    "os/signal"
    "syscall"
    // ...
)

func main() {
    // ... (setup mux dan server seperti di contoh http.Server)

    // Jalankan server di goroutine agar tidak memblokir
    go func() {
        log.Println("Server starting on", server.Addr)
        if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatalf("could not listen on %s: %v\n", server.Addr, err)
        }
    }()

    // Channel untuk mendengarkan sinyal shutdown
    stop := make(chan os.Signal, 1)
    signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

    // Tunggu sinyal
    <-stop
    log.Println("Shutting down server...")

    // Buat konteks dengan timeout untuk shutdown
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    // Panggil Shutdown()
    if err := server.Shutdown(ctx); err != nil {
        log.Fatalf("Server Shutdown Failed:%+v", err)
    }

    log.Println("Server gracefully stopped")
}
```

### Menyajikan File Statis (Static File Serving)

`net/http` memiliki handler bawaan untuk menyajikan file dari sistem file.

```go
// Direktori 'assets' berisi file seperti 'style.css', 'app.js'
fs := http.FileServer(http.Dir("./assets"))

// Buat rute yang menghapus prefiks "/static/" sebelum mencarinya di direktori 'assets'
http.Handle("/static/", http.StripPrefix("/static/", fs))

// Request ke /static/style.css akan dilayani dari file ./assets/style.css
```

### Menggunakan `html/template` untuk Server-Side Rendering

Jika Anda membangun aplikasi web tradisional, bukan hanya API, Anda perlu me-render HTML. Paket `html/template` adalah cara yang aman dan kuat untuk melakukannya. Ia secara otomatis melakukan *escaping* terhadap data untuk mencegah serangan XSS (Cross-Site Scripting).

```go
import "html/template"

// Definisikan data yang akan dikirim ke template
type PageData struct {
    Title   string
    Content string
}

func renderTemplateHandler(w http.ResponseWriter, r *http.Request) {
    // Parse file template. Sebaiknya dilakukan sekali saat startup.
    tmpl, err := template.ParseFiles("template.html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    data := PageData{
        Title:   "Halaman Dinamis",
        Content: "Ini adalah konten yang datang dari Go!",
    }

    // Execute template, menulis hasilnya ke ResponseWriter
    err = tmpl.Execute(w, data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}
```

File `template.html`:
```html
<!DOCTYPE html>
<html>
<head>
    <title>{{.Title}}</title>
</head>
<body>
    <h1>{{.Title}}</h1>
    <p>{{.Content}}</p>
</body>
</html>
```

---

## 8. Bab 7: Sisi Klien - Membuat HTTP Request

Paket `net/http` tidak hanya untuk membuat server, tapi juga untuk bertindak sebagai klien untuk berkomunikasi dengan server lain.

### Request Sederhana dengan `http.Get` & `http.Post`

Untuk request sederhana, Anda bisa menggunakan fungsi-fungsi pembantu ini.

```go
// GET request
resp, err := http.Get("https://api.example.com/data")
if err != nil {
    // Handle error
}
defer resp.Body.Close()
body, _ := ioutil.ReadAll(resp.Body)
fmt.Println(string(body))


// POST request dengan form data
resp, err := http.PostForm("https://api.example.com/submit", url.Values{
    "name": {"John Doe"},
    "email": {"john@example.com"},
})
// ... (handle response)
```

### Kontrol Penuh dengan `http.Client`

Sama seperti `http.ListenAndServe` yang merupakan shortcut untuk `http.Server`, fungsi `http.Get` dan `http.Post` adalah shortcut yang menggunakan klien default (`http.DefaultClient`). Untuk kontrol lebih, Anda harus membuat `http.Client` Anda sendiri.

```go
// Buat klien kustom
client := &http.Client{
    Timeout: 10 * time.Second, // Sangat penting untuk mengatur timeout!
}

// Buat request baru
req, err := http.NewRequest("GET", "https://api.github.com/users/golang", nil)
if err != nil {
    // handle error
}

// Tambahkan header jika perlu
req.Header.Add("Accept", "application/vnd.github.v3+json")

// Kirim request menggunakan klien kustom
resp, err := client.Do(req)
if err != nil {
    // handle error
}
defer resp.Body.Close()

// ... proses respons ...
```

### Mengatur Timeout pada Klien

**Selalu** atur timeout pada `http.Client` Anda. Jika tidak, sebuah request bisa menggantung selamanya, menghabiskan sumber daya di aplikasi Anda.

### Mengirim JSON ke Server Lain

```go
import "bytes"

// Data yang akan dikirim
postData := map[string]interface{}{
    "title":  "Judul Post Baru",
    "body":   "Ini adalah isi dari post.",
    "userId": 1,
}

// Marshal data ke JSON
jsonData, _ := json.Marshal(postData)

// Buat request dengan body JSON
req, _ := http.NewRequest("POST", "https://jsonplaceholder.typicode.com/posts", bytes.NewBuffer(jsonData))
req.Header.Set("Content-Type", "application/json")

// Kirim request
client := &http.Client{Timeout: 5 * time.Second}
resp, err := client.Do(req)
// ... handle response ...
```

---

## 9. Bab 8: Praktik Terbaik & Struktur Proyek

Menulis kode yang berfungsi itu satu hal, menulis kode yang terawat, teruji, dan skalabel itu hal lain.

### Hindari Global State (`DefaultServeMux` dan `http.HandleFunc`)

Di awal, kita menggunakan `http.HandleFunc` yang mendaftarkan rute ke `DefaultServeMux` global. Ini nyaman untuk contoh kecil, tetapi buruk untuk aplikasi besar karena:
*   Bisa terjadi konflik jika beberapa bagian kode atau library mencoba mendaftarkan rute yang sama.
*   Membuat pengujian menjadi sulit karena Anda tidak bisa dengan mudah membuat instance server yang terisolasi.

**Solusi**: Selalu buat `ServeMux` atau router Anda sendiri di `main` dan teruskan ke `http.Server`.

```go
// Buruk (menggunakan global)
func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}

// Baik (terisolasi)
func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", handler)
    server := &http.Server{Addr: ":8080", Handler: mux}
    server.ListenAndServe()
}
```

### Dependency Injection pada Handler

Handler Anda seringkali membutuhkan akses ke dependensi lain, seperti koneksi database, logger, atau konfigurasi. Jangan gunakan variabel global untuk ini!

**Pola yang benar adalah Dependency Injection**. Buat sebuah struct yang menampung dependensi Anda, dan buat handler Anda sebagai method dari struct tersebut.

```go
// Definisikan dependensi Anda, mis. koneksi database
type DBConnection struct { /* ... */ }

// Buat struct Server/Application yang menampung semua dependensi
type Application struct {
    db     *DBConnection
    logger *log.Logger
}

// Buat handler sebagai method dari Application
func (app *Application) usersHandler(w http.ResponseWriter, r *http.Request) {
    // Di sini, Anda punya akses ke app.db dan app.logger
    // users, err := app.db.GetAllUsers()
    // ...
    app.logger.Println("Menangani request pengguna")
    w.Write([]byte("Daftar Pengguna"))
}

func main() {
    // Inisialisasi dependensi
    logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
    db := &DBConnection{} // Inisialisasi koneksi DB asli di sini

    // Buat instance dari Application
    app := &Application{
        db:     db,
        logger: logger,
    }

    // Daftarkan method handler
    mux := http.NewServeMux()
    mux.HandleFunc("/users", app.usersHandler)

    // ... jalankan server dengan mux ...
}
```
Pola ini membuat kode Anda sangat modular dan mudah diuji. Anda bisa dengan mudah membuat `Application` palsu dengan *mock* database untuk pengujian unit.

### Struktur Proyek yang Dianjurkan

Untuk proyek yang lebih besar dari satu file, Anda perlu struktur yang baik. Berikut adalah struktur yang umum dan skalabel:

```
my-web-app/
├── cmd/                # Titik masuk utama aplikasi
│   └── web/
│       └── main.go     # Inisialisasi server, dependensi, dan rute
├── internal/           # Kode privat aplikasi (tidak bisa diimpor oleh proyek lain)
│   ├── handler/        # Berisi handler HTTP (mis. user.go, product.go)
│   │   └── routes.go   # Fungsi untuk setup semua rute aplikasi
│   ├── model/          # Berisi struct data/domain (mis. user.go -> type User struct)
│   ├── store/          # Berisi logika akses data (database, dll)
│   │   └── user_store.go
│   └── service/        # (Opsional) Berisi logika bisnis
├── ui/                 # File UI statis
│   ├── html/           # Template HTML
│   └── static/         # CSS, JS, Gambar
├── go.mod
└── go.sum
```

---

## 10. Bab 9: Contoh Proyek Lengkap - API CRUD Sederhana

Mari kita gabungkan semua yang telah kita pelajari untuk membuat API CRUD (Create, Read, Update, Delete) sederhana untuk manajemen "Tugas" (Todo). Kita akan menggunakan router `chi` dan penyimpanan di memori (in-memory) agar tetap sederhana.

### Struktur File Proyek

```
todo-api/
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   ├── handler/
│   │   ├── todo.go
│   │   └── routes.go
│   └── model/
│       └── todo.go
└── go.mod
```

### Definisi Model dan Penyimpanan (In-Memory)

**`internal/model/todo.go`**
```go
package model

import "sync"

type Todo struct {
	ID        int    `json:"id"`
	Task      string `json:"task"`
	Completed bool   `json:"completed"`
}

// TodoStore adalah penyimpanan in-memory yang aman untuk konkurensi.
type TodoStore struct {
	sync.Mutex
	todos  map[int]Todo
	nextID int
}

func NewTodoStore() *TodoStore {
	return &TodoStore{
		todos:  make(map[int]Todo),
		nextID: 1,
	}
}

// Implementasikan method-method CRUD di sini (Create, Get, Update, Delete) ...
// (Ini akan kita tambahkan di handler)
```

### Membuat Handler

**`internal/handler/todo.go`**
```go
package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"sync"
	"todo-api/internal/model"

	"github.com/go-chi/chi/v5"
)

type TodoHandler struct {
	Store *model.TodoStore
}

// ... (helper functions untuk respons JSON, dll)

func (h *TodoHandler) ListTodos(w http.ResponseWriter, r *http.Request) {
	h.Store.Lock()
	defer h.Store.Unlock()

	var todos []model.Todo
	for _, todo := range h.Store.todos {
		todos = append(todos, todo)
	}
	// ... respond with JSON ...
}

func (h *TodoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	// ... decode JSON body ...
	// ... create new todo in store ...
	// ... respond with new todo ...
}

func (h *TodoHandler) GetTodo(w http.ResponseWriter, r *http.Request) {
    idStr := chi.URLParam(r, "id")
    id, _ := strconv.Atoi(idStr)
    // ... get todo from store ...
    // ... respond with JSON ...
}

// Implementasikan UpdateTodo dan DeleteTodo
```
*(Catatan: Kode handler CRUD lengkap akan sangat panjang, jadi saya berikan kerangkanya. Ide utamanya adalah handler ini memiliki akses ke `TodoStore`)*

**`internal/handler/routes.go`**
```go
package handler

import "github.com/go-chi/chi/v5"

func (h *TodoHandler) Routes() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/todos", h.ListTodos)
	r.Post("/todos", h.CreateTodo)
	r.Get("/todos/{id}", h.GetTodo)
	r.Put("/todos/{id}", h.UpdateTodo)
	r.Delete("/todos/{id}", h.DeleteTodo)

	return r
}
```

### Merakit Semuanya di `main.go`

**`cmd/api/main.go`**
```go
package main

import (
	"log"
	"net/http"
	"todo-api/internal/handler"
	"todo-api/internal/model"
)

func main() {
	// 1. Inisialisasi dependensi
	store := model.NewTodoStore()
	todoHandler := &handler.TodoHandler{
		Store: store,
	}

	// 2. Setup router
	router := todoHandler.Routes()

	// 3. Buat dan konfigurasikan server
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Println("Starting Todo API server on :8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
```

### Menjalankan dan Menguji Proyek

1.  Jalankan `go mod init todo-api` dan `go mod tidy` di root proyek.
2.  Jalankan server: `go run cmd/api/main.go`.
3.  Gunakan `curl` atau Postman untuk menguji endpoint Anda:
    *   `POST /todos` dengan body JSON untuk membuat tugas baru.
    *   `GET /todos` untuk melihat semua tugas.
    *   `GET /todos/1` untuk melihat tugas dengan ID 1.
    *   Dan seterusnya.

---

## 11. Kesimpulan & Langkah Selanjutnya

Anda telah menempuh perjalanan yang panjang! Dari "Hello, World!" sederhana hingga membangun API terstruktur dengan praktik terbaik, Anda sekarang memiliki fondasi yang sangat kuat dalam menggunakan paket `net/http` Go.

**Poin-poin Kunci untuk Diingat:**

*   `net/http` dibangun di atas `interface` sederhana seperti `http.Handler`, membuatnya sangat fleksibel dan dapat disusun.
*   Pahami peran `ResponseWriter` dan `Request` - ini adalah inti dari setiap handler.
*   Middleware adalah pola yang sangat kuat di Go untuk fungsionalitas lintas-fungsi (cross-cutting concerns).
*   Gunakan router pihak ketiga seperti `chi` untuk aplikasi yang lebih dari sekadar "Hello, World!".
*   Selalu konfigurasikan `http.Server` dengan timeout untuk aplikasi produksi.
*   Gunakan pola Dependency Injection (handler sebagai method pada struct) untuk mengelola dependensi dan membuat kode yang dapat diuji.

**Langkah Selanjutnya:**

*   **Pengujian**: Pelajari cara menulis pengujian unit dan integrasi untuk handler HTTP Anda menggunakan paket `net/http/httptest`.
*   **Database**: Ganti penyimpanan in-memory dengan database sungguhan seperti PostgreSQL atau MySQL.
*   **Autentikasi**: Terapkan sistem autentikasi yang lebih kuat, misalnya menggunakan JWT (JSON Web Tokens).
*   **Deployment**: Pelajari cara men-deploy aplikasi Go Anda, misalnya menggunakan Docker dan platform cloud seperti DigitalOcean, AWS, atau Google Cloud.
*   **Observability**: Tambahkan logging terstruktur, metrik (mis. dengan Prometheus), dan tracing terdistribusi.

Paket `net/http` adalah alat yang luar biasa. Ia memberikan Anda kekuatan dan performa dari Go tanpa menyembunyikan cara kerja web yang sebenarnya. Teruslah berlatih, membangun proyek, dan menjelajahi ekosistem Go yang luas.

Selamat coding

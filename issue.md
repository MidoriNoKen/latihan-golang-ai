# Issue: Implementasi Boilerplate Clean Architecture pada REST API Go (Gin & GORM)

## Deskripsi Tugas
Tugas ini bertujuan untuk merestrukturisasi dan mengimplementasikan struktur proyek Go menggunakan konsep **Clean Architecture**. Proyek saat ini menggunakan Gin Gonic sebagai router dan GORM untuk database MySQL. Implementasi baru harus memisahkan tanggung jawab (Separation of Concerns) ke dalam layer-layer yang terdefinisi dengan jelas dan menyediakan format response API JSON yang terstandarisasi.

---

## 1. Struktur Folder yang Diusulkan

Restrukturisasi proyek agar mengikuti pola Clean Architecture berikut:

```text
├── config/             # Konfigurasi aplikasi & database (sudah ada)
├── domain/             # Layer inti: Definisi Struct/Entity & Interface (Kontrak)
├── repository/         # Layer database: Query & interaksi langsung ke database (GORM)
├── service/            # Layer logika bisnis: Validasi & pemrosesan data utama
├── controller/         # Layer presentasi: HTTP Handler, parsing input, & trigger service
├── pkg/                # Utility & helper yang dapat digunakan kembali secara global
│   └── response/       # Helper format response JSON standar
├── routes/             # Perutean API (menghubungkan route ke controller) (sudah ada)
└── main.go             # Entry point aplikasi (sudah ada)
```

---

## 2. Standarisasi Format Response JSON

Semua API controller harus menggunakan struktur response JSON yang konsisten.

### A. Response Sukses (Success Response)
Digunakan ketika request berhasil diproses.
```json
{
  "success": true,
  "message": "Pesan sukses yang deskriptif",
  "data": { ... } // Dapat berupa object atau array, null jika tidak ada data
}
```

### B. Response Gagal/Error (Error Response)
Digunakan ketika terjadi error (validasi input, database error, not found, dll).
```json
{
  "success": false,
  "message": "Pesan error utama (misal: 'Internal Server Error')",
  "errors": "Detail error atau array error validasi" // Bisa string atau array string
}
```

---

## 3. Langkah Implementasi (High-Level Checklist)

### [ ] Langkah 1: Buat Standard Response Helper
*   Buat package `pkg/response`.
*   Buat helper function untuk menyederhanakan pengiriman response sukses (`JSONSuccess`) dan response error (`JSONError`) menggunakan *Gin Context*.

### [ ] Langkah 2: Definisikan Domain (Entity & Interface)
*   Pilih satu entitas contoh sederhana (misalnya `User` atau `Product`).
*   Buat file di dalam folder `domain/` untuk mendefinisikan:
    *   **Struct Entity** (representasi tabel database).
    *   **Repository Interface** (kontrak method database yang dibutuhkan).
    *   **Service Interface** (kontrak business logic yang disediakan).

### [ ] Langkah 3: Implementasikan Repository Layer
*   Buat implementasi dari `Repository Interface` di folder `repository/`.
*   Layer ini hanya menerima koneksi database (`*gorm.DB`) dan melakukan query langsung (Create, Read, Update, Delete).

### [ ] Langkah 4: Implementasikan Service/Usecase Layer
*   Buat implementasi dari `Service Interface` di folder `service/`.
*   Layer ini menerima repository via constructor (Dependency Injection).
*   Implementasikan logika bisnis, validasi, dan penanganan error di layer ini sebelum dikembalikan ke controller.

### [ ] Langkah 5: Implementasikan Controller Layer
*   Buat handler HTTP di folder `controller/`.
*   Layer ini menerima service via constructor (Dependency Injection).
*   Tugas controller: membaca request (JSON/Query), memanggil service, dan mengembalikan response menggunakan standard response helper.

### [ ] Langkah 6: Hubungkan Komponen (Dependency Injection & Routing)
*   Di dalam `routes/routes.go` atau `main.go`, lakukan inisialisasi:
    1.  Repository (inject `config.DB`).
    2.  Service (inject Repository).
    3.  Controller (inject Service).
*   Daftarkan endpoint API baru ke router Gin dan hubungkan ke method controller yang sesuai.

### [ ] Langkah 7: Verifikasi dan Pengujian
*   Jalankan server lokal (`go run main.go`).
*   Uji endpoint yang telah dibuat menggunakan curl atau Postman.
*   Pastikan format JSON yang dihasilkan sesuai dengan standar sukses dan error yang didefinisikan.

---

## Prinsip Penting yang Harus Diikuti
1.  **Dependency Flow**: Ketergantungan hanya boleh mengalir ke dalam. Controller memanggil Service, Service memanggil Repository. Repository TIDAK boleh memanggil Service, dan Service TIDAK boleh tahu tentang HTTP Context (Gin).
2.  **No Low-Level Logic in Controller**: Controller tidak boleh mengandung logika bisnis atau query SQL langsung.
3.  **Strict Typing & Interface**: Gunakan interface untuk decoupling agar kode mudah di-unit test di masa mendatang.

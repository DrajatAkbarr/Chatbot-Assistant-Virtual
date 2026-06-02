# MindFlow - Asisten Virtual Kesehatan Mental & Produktivitas

Halo! Di sini aku buat project bernama MindFlow, sebuah aplikasi berbasis terminal yang fokus buat bantu mantau suasana hati (mood) sekaligus ngatur daftar tugas harian. Aku sengaja desain tampilannya serapi mungkin pakai tabel dan ASCII art biar nyaman waktu dipakai.

## Fitur Utama

Jadi nanti programnya bakal punya beberapa fitur utama yang bisa langsung dipilih dari menu. Berikut rinciannya:

### 1. Tampilan Awal yang Rapi

Begitu program dijalankan, bakal langsung muncul ASCII Art bertuliskan "ASISTEN VIRTUAL KESEHATAN" yang lumayan besar, dilanjut sama kotak sapaan. Semua menu aku buat di dalam format tabel biar enak dibaca dan nggak berantakan.

### 2. Kelola Catatan Suasana Hati (Mood)

Di menu ini bisa mencatat perasaan tiap hari. Aku taruh fungsi buat tambah, ubah, hapus, dan lihat semua catatan. Waktu nambah data, bisa masukin tanggal (format DD-MM-YYYY), tugas apa yang lagi berkaitan sama mood saat itu, skor emosi dari 1 sampai 10, dan deskripsi perasaan.

### 3. Kelola Daftar Tugas Harian

Biar kerjaan lebih teratur, gunain ini buat nyatat tugas. Bisa masukin nama tugas, tingkat prioritas (1-5), estimasi waktu pengerjaan (durasi dalam menit), dan status selesai atau belum. Tentu saja, data tugas ini bebas diubah atau dihapus kapan aja.

### 4. Pencarian Data (Searching)

Kalau catatan atau tugas udah menumpuk, cari data manual pasti repot. Makanya aku bikin fitur pencarian dengan dua metode dari materi:

- **Sequential Search:** Bisa cari data tugas atau mood pakai kata kunci bebas atau tanggal. Program bakal ngecek satu per satu dan nampilin semua yang cocok.
- **Binary Search:** Khusus buat cari data berdasarkan tanggal pasti. Metode ini jauh lebih cepat buat data yang udah urut.

### 5. Pengurutan Tugas (Sorting)

Biar gampang nentuin mana yang harus dikerjain duluan, aku taruh fungsi pengurutan:

- **Berdasarkan Prioritas (Selection Sort):** Mengurutkan tugas dari prioritas paling tinggi ke paling rendah (Descending).
- **Berdasarkan Durasi (Insertion Sort):** Mengurutkan tugas dari yang butuh waktu paling sebentar sampai paling lama (Ascending).

### 6. Statistik & Visualisasi

Biar kelihatan perubahannya, program bisa ngeluarin rekap data otomatis:

- **Tren Mood Mingguan:** Menampilkan total catatan, rata-rata skor, skor tertinggi dan terendah. Ada juga diagram batang harian sederhana pakai simbol `#`.
- **Tingkat Penyelesaian Tugas:** Menghitung berapa banyak tugas yang udah beres, yang belum, rata-rata durasi, dan persentase kelar yang divisualisasikan pakai _progress bar_.

## Cara Menjalankan Program

Pastikan komputer udah terpasang Golang. Buat jalanin aplikasinya, cukup buka terminal, arahkan ke folder tempat file ini berada, terus jalankan perintah:

```bash
go run main.go
```

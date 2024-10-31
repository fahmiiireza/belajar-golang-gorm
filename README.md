# Dokumen Spesifikasi Perangkat Lunak

## Library Management System

**Nama**: Fahmi Fahreza  
**Nomor Mahasiswa**: 0706012214006  
**Email**: ffahreza@student.ciputra.ac.id  
**Dokumen**: Dokumen Spesifikasi Perangkat Lunak  

---

## Daftar Isi

1. [Latar Belakang](#latar-belakang) ............................ 2  
2. [Analisis Kebutuhan](#analisis-kebutuhan) ..................... 3  
    2.1 [Daftar Stakeholder](#daftar-stakeholder) .................. 3  
    2.2 [Kebutuhan Stakeholder](#kebutuhan-stakeholder) ............. 4  
3. [Spesifikasi Kebutuhan](#spesifikasi-kebutuhan) .............. 5  
    3.1 [Use Case Diagram](#use-case-diagram) ...................... 5  
    3.2 [Kebutuhan Fungsional](#kebutuhan-fungsional) ............... 6  
    3.3 [Kebutuhan Non-fungsional](#kebutuhan-non-fungsional) ....... 7  

---

## 1. Latar Belakang

Perpustakaan di Universitas Greenfield mengalami tantangan dalam mengelola inventaris buku, mengatur akses pengguna, dan menjaga agar data peminjaman tetap akurat dan aman. Saat ini, proses peminjaman, pengelolaan stok, dan pengecekan status ketersediaan buku dilakukan secara manual, yang sering mengakibatkan keterlambatan dalam pelayanan serta kurangnya visibilitas bagi mahasiswa dan staf.

Library Management System ini dirancang untuk mengatasi masalah-masalah tersebut dengan menyediakan sistem otomatis yang memungkinkan staf perpustakaan untuk mengelola buku secara efisien, serta memberikan kemudahan akses bagi mahasiswa dalam mencari dan memesan buku yang mereka butuhkan. Sistem ini juga memungkinkan admin untuk mengelola peran dan hak akses, sehingga keamanan data dapat lebih terjamin.

---

## 2. Analisis Kebutuhan

### 2.1 Daftar Stakeholder

| Nama            | Deskripsi                          | Jenis       |
|-----------------|------------------------------------|-------------|
| Admin           | Pengelola sistem utama, bertanggung jawab atas pengaturan akun pengguna dan peran dalam sistem. | Non-User |
| Pustakawan      | Pengelola buku di perpustakaan, bertanggung jawab atas penambahan, pengeditan, dan penghapusan data buku. | User |
| Mahasiswa       | Pengguna perpustakaan, mengakses sistem untuk mencari dan memesan buku yang tersedia. | User |
| Teknisi IT      | Penanggung jawab terhadap infrastruktur sistem, melakukan pemeliharaan dan pembaruan sistem. | Non-User |

### 2.2 Kebutuhan Stakeholder

| Stakeholder | Kebutuhan                                                                 |
|-------------|---------------------------------------------------------------------------|
| Admin       | Dapat mengelola akun pustakawan, menetapkan peran, serta mengatur keamanan dan kontrol akses. |
| Pustakawan  | Dapat mengelola inventaris buku, termasuk penambahan, pengeditan, dan penghapusan buku dalam sistem. |
| Mahasiswa   | Dapat mencari buku, melihat status ketersediaan, dan memesan buku secara online. |
| Teknisi IT  | Mengelola infrastruktur sistem dan memastikan sistem dapat berjalan secara optimal. |

---

## 3. Spesifikasi Kebutuhan

### 3.1 Use Case Diagram

[Diagram](https://lucid.app/lucidchart/4b2b7614-91b7-47b5-862c-897428dbe623/edit?viewport_loc=-1055%2C-319%2C4744%2C2587%2C0_0&invitationId=inv_7101167f-0671-4264-b981-9aea21f6770f)

### 3.2 Kebutuhan Fungsional

1. **Otentikasi Pengguna**  
   Sistem harus menyediakan fitur otentikasi menggunakan token JWT untuk menjaga keamanan akses.

2. **Manajemen Akun Pengguna**  
   Admin harus dapat membuat, mengubah, dan menghapus akun pustakawan serta menetapkan peran yang sesuai.

3. **Pengelolaan Inventaris Buku**  
   - Pustakawan dapat menambahkan buku baru, mengedit informasi buku, dan menghapus buku dari katalog.
   - Sistem harus memungkinkan pustakawan untuk memantau status ketersediaan setiap buku.

4. **Akses Peminjaman Buku oleh Mahasiswa**  
   Mahasiswa dapat melihat daftar buku yang tersedia, mencari berdasarkan kriteria tertentu, dan memesan buku yang mereka butuhkan.

### 3.3 Kebutuhan Non-fungsional

1. **Keamanan**  
   Data sensitif harus dienkripsi dan otentikasi pengguna dilakukan melalui token JWT untuk memastikan keamanan akses.

2. **Ketersediaan**  
   Sistem harus memiliki waktu operasional 99.9%, sehingga mahasiswa dan staf perpustakaan dapat mengaksesnya kapan saja.

3. **Skalabilitas**  
   Sistem dirancang untuk menangani setidaknya 5.000 pengguna aktif secara bersamaan dan mampu menambah jumlah data buku tanpa penurunan kinerja.

4. **Kinerja**  
   Sistem harus mampu merespon permintaan dalam waktu kurang dari 1 detik untuk 95% permintaan pengguna.

5. **Pemeliharaan dan Ekstensibilitas**  
   Setiap microservice dapat diperbarui atau dikembangkan secara independen tanpa memengaruhi bagian lain dari sistem.

---


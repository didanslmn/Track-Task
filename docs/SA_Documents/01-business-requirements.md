# Business Requirements

## BR-001 User Account

Sistem harus memungkinkan user membuat akun dan login agar dapat mengakses fitur project dan task.

### Rationale

Data project dan task harus terkait dengan user yang valid.

### Priority

Must Have

## BR-002 Secure Authentication

Sistem harus melindungi endpoint utama menggunakan authentication berbasis token.

### Rationale

Data project bersifat private dan tidak boleh diakses oleh user anonim.

### Priority

Must Have

## BR-003 Project Ownership

Sistem harus menetapkan pembuat project sebagai owner project.

### Rationale

Owner diperlukan untuk mengatur member dan menjadi pemilik hak akses tertinggi dalam project.

### Priority

Must Have

## BR-004 Project Membership

Sistem harus mendukung relasi antara user dan project melalui membership.

### Rationale

Tidak semua user boleh mengakses semua project.

### Priority

Must Have

## BR-005 Member Invitation by Existing User

Owner project harus dapat menambahkan registered user sebagai member project.

### Rationale

Project dikerjakan oleh lebih dari satu orang, sehingga sistem perlu mendukung kolaborasi.

### Priority

Must Have

## BR-006 Task Creation

Member project harus dapat membuat task di dalam project.

### Rationale

Task adalah objek utama untuk mencatat pekerjaan.

### Priority

Must Have

## BR-007 Task Assignment

Task harus dapat diberikan kepada member project.

### Rationale

Assignment membuat penanggung jawab pekerjaan menjadi jelas.

### Priority

Must Have

## BR-008 Task Status Tracking

Task harus memiliki status agar progress pekerjaan bisa dilacak.

### Rationale

Project lead dan member perlu mengetahui kondisi pekerjaan saat ini.

### Priority

Must Have

## BR-009 Access Isolation

Sistem harus memastikan user hanya dapat mengakses project dan task yang dia ikuti.

### Rationale

Ini adalah requirement keamanan utama untuk mencegah kebocoran data antar project.

### Priority

Must Have

## BR-010 API Documentation

Sistem harus memiliki dokumentasi endpoint yang dapat digunakan developer lain.

### Rationale

API akan dikonsumsi frontend/mobile di masa depan.

### Priority

Should Have

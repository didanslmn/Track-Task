# Product Requirements Document

## 1. Background

TeamTask API dibuat sebagai backend MVP untuk aplikasi task management tim kecil. Pada fase pertama, produk hanya menyediakan API tanpa frontend.

## 2. Goals

- Membuat backend API untuk project dan task management.
- Menyediakan authentication dan authorization dasar.
- Menyediakan dokumentasi API yang bisa digunakan developer frontend.
- Mendeploy API agar bisa diuji secara publik.

## 3. User Personas

### Project Lead

Project Lead membuat project, mengatur member, membuat task, dan memantau progress.

### Team Member

Team Member melihat task dalam project, menerima assignment, dan memperbarui status task.

## 4. Functional Requirements

### Authentication

- User dapat register menggunakan nama, email, dan password.
- User dapat login menggunakan email dan password.
- Sistem mengembalikan access token setelah login berhasil.
- Endpoint protected harus menolak request tanpa token valid.

### Project Management

- Authenticated user dapat membuat project.
- Pembuat project otomatis menjadi owner.
- User dapat melihat daftar project yang dia ikuti.
- User dapat melihat detail project jika dia adalah member project.

### Project Membership

- Owner project dapat menambahkan user lain sebagai member.
- Non-owner tidak dapat menambahkan member.
- User di luar project tidak dapat melihat data project.

### Task Management

- Member project dapat membuat task di dalam project.
- Task memiliki title, description, status, assignee, dan due date opsional.
- Task default memiliki status `todo`.
- Member project dapat melihat daftar task dalam project.
- Assignee atau owner dapat mengubah status task.

## 5. Non-Functional Requirements

- API menggunakan JSON.
- Password harus di-hash.
- Token menggunakan JWT.
- Database menggunakan PostgreSQL.
- API memiliki health check endpoint.
- API memiliki test untuk flow utama.
- API dapat dijalankan menggunakan Docker.

## 6. MVP Acceptance

MVP dianggap selesai jika:

- Semua endpoint MVP selesai dibuat.
- Endpoint protected membutuhkan JWT.
- Authorization project member berjalan.
- Dokumentasi API tersedia.
- Test berjalan di local dan CI.
- API berhasil dideploy.

## 7. Constraints

- Fase pertama hanya backend.
- Tidak ada realtime update pada MVP.
- Tidak ada frontend pada MVP.
- Tidak ada payment atau multi-tenant organization pada MVP.

## 8. Future Enhancements

- Comment pada task
- Attachment pada task
- Notification
- Activity log
- Role lebih detail seperti admin, maintainer, viewer
- Workspace atau organization
- Search dan filter task lanjutan

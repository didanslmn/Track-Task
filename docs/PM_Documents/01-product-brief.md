# Product Brief

## Nama Produk

TeamTask API

## Problem Statement

Tim kecil sering mengelola task lewat chat, spreadsheet, atau catatan terpisah. Akibatnya status pekerjaan sulit dilacak, penanggung jawab tidak jelas, dan project lead kesulitan mengetahui progress terbaru.

## Product Vision

Menyediakan backend API yang sederhana, aman, dan mudah dikembangkan untuk membantu tim kecil mengelola project dan task secara terpusat.

## Target User

### Primary User

Project lead, founder, atau koordinator tim kecil yang perlu membagi dan memantau pekerjaan.

### Secondary User

Anggota tim yang perlu melihat task miliknya dan memperbarui status pekerjaan.

## User Needs

- User perlu login dengan aman.
- Project lead perlu membuat project.
- Project lead perlu menambahkan member ke project.
- Member perlu membuat task dalam project.
- Task perlu bisa diberikan kepada member tertentu.
- Member perlu mengubah status task.
- User perlu melihat daftar task dalam project.

## Value Proposition

TeamTask API menyediakan fondasi backend yang jelas untuk aplikasi task management, dengan fitur inti yang cukup untuk workflow project kecil dan struktur yang siap dikembangkan.

## MVP Feature List

- Register user
- Login user
- Create project
- List user's projects
- Add member to project
- Create task inside project
- Assign task to project member
- Update task status
- List tasks by project
- Basic authorization
- API documentation

## Non-Functional Goals

- API response konsisten
- Error message mudah dipahami
- Password disimpan dalam bentuk hash
- Endpoint dilindungi dengan authentication
- Business rule penting dilindungi dengan authorization
- Struktur project mudah dites dan dikembangkan

## Product Risks

- Scope melebar terlalu cepat sebelum MVP selesai
- Authorization antar project tidak cukup kuat
- Dokumentasi API tidak sinkron dengan implementasi
- Deployment terlambat karena konfigurasi environment tidak disiapkan sejak awal

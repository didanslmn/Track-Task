# TeamTask API - Project Overview

## Ringkasan

TeamTask API adalah backend service untuk membantu tim kecil mengelola project, member, dan task secara terstruktur. Produk ini menyediakan REST API untuk authentication, project management, membership, task assignment, dan tracking status task.

## Tujuan Proyek

Membangun backend API yang siap digunakan oleh aplikasi frontend atau mobile untuk kebutuhan task management tim kecil.

## Scope MVP

MVP fokus pada fitur inti:

- User registration dan login
- Project creation dan project listing
- Project membership
- Task creation, assignment, update, dan listing
- Authorization berbasis project member
- API documentation
- Automated test dasar
- Deployment ke cloud platform

## Out of Scope Untuk MVP

Fitur berikut tidak dikerjakan pada fase MVP:

- Realtime notification
- File attachment
- Comment thread pada task
- Calendar view
- Payment/subscription
- Organization-level workspace
- Advanced analytics

## Target Pengguna

- Tim kecil berisi 2-10 orang
- Founder atau project lead yang ingin tracking task sederhana
- Developer atau freelancer yang bekerja dalam project kecil

## Success Metrics

Keberhasilan MVP diukur dari:

- Semua endpoint utama tersedia dan terdokumentasi
- User dapat membuat project dan task lewat API
- Authorization mencegah akses lintas project
- Test untuk flow utama berjalan di CI
- API berhasil dideploy dan memiliki health check publik

## Asumsi

- Frontend belum dibuat pada fase ini
- API akan dikonsumsi oleh frontend/mobile di masa depan
- Database utama menggunakan PostgreSQL
- Authentication menggunakan JWT
- Deployment awal menggunakan platform sederhana seperti Render, Railway, atau Fly.io

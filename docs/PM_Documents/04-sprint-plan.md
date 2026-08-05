# Sprint Plan

## Sprint Format

- Durasi sprint: 1 minggu
- Total sprint MVP: 4 sprint
- Metode: Scrum ringan
- Board: Backlog, Ready, In Progress, Code Review, Testing, Done

## Sprint 1 - Foundation and Auth

### Goal

Backend dapat berjalan lokal dan user dapat register/login.

### Scope

- Initialize backend project
- Setup environment config
- Setup PostgreSQL connection
- Setup migration
- Create health check endpoint
- Create register endpoint
- Create login endpoint
- Create JWT middleware

### Deliverables

- API berjalan lokal
- Database terkoneksi
- Auth endpoint bekerja
- Auth test dasar tersedia

## Sprint 2 - Project and Membership

### Goal

User dapat membuat project dan owner dapat menambahkan member.

### Scope

- Create project
- List user's projects
- Get project detail
- Add project member
- Project owner rule
- Project member authorization

### Deliverables

- Project CRUD minimal
- Membership berjalan
- Unauthorized access ditolak

## Sprint 3 - Task Management

### Goal

Member project dapat mengelola task.

### Scope

- Create task
- List tasks by project
- Update task
- Update task status
- Assign task
- Validate assignee is project member

### Deliverables

- Task workflow berjalan
- Task authorization berjalan
- Integration test untuk flow task tersedia

## Sprint 4 - Quality, Docs, and Deploy

### Goal

API siap ditunjukkan sebagai portfolio.

### Scope

- Complete test coverage untuk flow utama
- Create API documentation
- Create README
- Create Dockerfile
- Create CI workflow
- Deploy API
- Prepare release notes

### Deliverables

- API docs tersedia
- Test berjalan di CI
- API deployed
- README portfolio-ready

## Definition of Done

Sebuah task dianggap selesai jika:

- Kode selesai dibuat
- Test relevan ditambahkan atau diperbarui
- Manual test berhasil
- Error response ditangani
- Dokumentasi diperbarui jika endpoint berubah
- Tidak ada secret hardcoded
- Pull request sudah direview

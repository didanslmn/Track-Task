# User Stories

## Authentication

### US-001 Register User

Sebagai user baru, saya ingin membuat akun agar saya dapat menggunakan TeamTask API.

Acceptance:

- User dapat register dengan name, email, dan password.
- Email harus unik.
- Password disimpan dalam bentuk hash.
- Response tidak mengembalikan password.

### US-002 Login User

Sebagai registered user, saya ingin login agar saya mendapat access token.

Acceptance:

- User dapat login dengan email dan password.
- Sistem mengembalikan JWT jika credential valid.
- Sistem menolak credential salah.

## Project

### US-003 Create Project

Sebagai authenticated user, saya ingin membuat project agar saya dapat mengelola task di dalamnya.

Acceptance:

- User harus login.
- Project memiliki name dan description opsional.
- Pembuat project otomatis menjadi owner.

### US-004 List My Projects

Sebagai user, saya ingin melihat daftar project yang saya ikuti agar saya tahu project yang bisa saya akses.

Acceptance:

- User harus login.
- Response hanya berisi project tempat user menjadi member.

### US-005 View Project Detail

Sebagai member project, saya ingin melihat detail project agar saya memahami informasi project.

Acceptance:

- User harus login.
- User harus menjadi member project.
- Non-member mendapat response forbidden.

## Membership

### US-006 Add Project Member

Sebagai project owner, saya ingin menambahkan user lain ke project agar mereka dapat ikut mengerjakan task.

Acceptance:

- User harus login.
- Hanya owner yang dapat menambahkan member.
- Target user harus sudah terdaftar.
- User tidak boleh ditambahkan dua kali ke project yang sama.

## Task

### US-007 Create Task

Sebagai member project, saya ingin membuat task agar pekerjaan dapat dicatat.

Acceptance:

- User harus login.
- User harus menjadi member project.
- Title wajib diisi.
- Status default adalah `todo`.

### US-008 Assign Task

Sebagai member project, saya ingin assign task ke member agar penanggung jawab jelas.

Acceptance:

- Assignee harus menjadi member project.
- Non-member tidak dapat menjadi assignee.
- Task tetap valid jika assignee kosong.

### US-009 List Project Tasks

Sebagai member project, saya ingin melihat daftar task dalam project agar saya dapat memantau pekerjaan.

Acceptance:

- User harus login.
- User harus menjadi member project.
- Response hanya menampilkan task dari project tersebut.

### US-010 Update Task Status

Sebagai assignee atau owner, saya ingin mengubah status task agar progress pekerjaan tercatat.

Acceptance:

- User harus login.
- User harus memiliki permission update status.
- Status hanya boleh `todo`, `in_progress`, atau `done`.

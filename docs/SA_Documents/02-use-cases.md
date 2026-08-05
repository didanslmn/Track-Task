# Use Cases

## UC-001 Register User

### Actor

Guest

### Preconditions

- Guest belum login.
- Email belum terdaftar.

### Main Flow

1. Guest mengirim name, email, dan password.
2. Sistem memvalidasi input.
3. Sistem mengecek uniqueness email.
4. Sistem melakukan hash password.
5. Sistem menyimpan user baru.
6. Sistem mengembalikan data user tanpa password.

### Alternative Flows

- Email sudah digunakan: sistem mengembalikan 409.
- Input tidak valid: sistem mengembalikan 400.

### Postconditions

- User baru tersimpan di database.

## UC-002 Login User

### Actor

Guest

### Preconditions

- User sudah terdaftar.

### Main Flow

1. Guest mengirim email dan password.
2. Sistem mencari user berdasarkan email.
3. Sistem membandingkan password dengan password hash.
4. Sistem membuat JWT.
5. Sistem mengembalikan token dan data user ringkas.

### Alternative Flows

- Email tidak ditemukan: sistem mengembalikan 401.
- Password salah: sistem mengembalikan 401.

### Postconditions

- User mendapat token untuk akses endpoint protected.

## UC-003 Create Project

### Actor

Authenticated User

### Preconditions

- User sudah login.

### Main Flow

1. User mengirim name dan description opsional.
2. Sistem memvalidasi token.
3. Sistem memvalidasi input.
4. Sistem membuat project.
5. Sistem membuat membership dengan role `owner`.
6. Sistem mengembalikan data project.

### Alternative Flows

- Token tidak valid: sistem mengembalikan 401.
- Name kosong: sistem mengembalikan 400.

### Postconditions

- Project baru tersimpan.
- Pembuat project menjadi owner.

## UC-004 Add Project Member

### Actor

Project Owner

### Preconditions

- Actor sudah login.
- Actor adalah owner project.
- Target user sudah terdaftar.

### Main Flow

1. Owner mengirim project ID dan user ID atau email target.
2. Sistem memvalidasi token.
3. Sistem mengecek actor adalah owner project.
4. Sistem mengecek target user valid.
5. Sistem mengecek target belum menjadi member project.
6. Sistem membuat membership baru dengan role `member`.
7. Sistem mengembalikan data membership.

### Alternative Flows

- Actor bukan owner: sistem mengembalikan 403.
- Target user tidak ditemukan: sistem mengembalikan 404.
- Target sudah menjadi member: sistem mengembalikan 409.

### Postconditions

- Target user menjadi member project.

## UC-005 Create Task

### Actor

Project Member

### Preconditions

- Actor sudah login.
- Actor adalah member project.

### Main Flow

1. Member mengirim title, description, assignee, dan due date opsional.
2. Sistem memvalidasi token.
3. Sistem mengecek actor adalah member project.
4. Sistem memvalidasi title.
5. Jika assignee diisi, sistem mengecek assignee adalah member project.
6. Sistem membuat task dengan status default `todo`.
7. Sistem mengembalikan data task.

### Alternative Flows

- Actor bukan member: sistem mengembalikan 403.
- Title kosong: sistem mengembalikan 400.
- Assignee bukan member project: sistem mengembalikan 400.

### Postconditions

- Task baru tersimpan di project.

## UC-006 List Project Tasks

### Actor

Project Member

### Preconditions

- Actor sudah login.
- Actor adalah member project.

### Main Flow

1. Member mengirim request daftar task berdasarkan project ID.
2. Sistem memvalidasi token.
3. Sistem mengecek actor adalah member project.
4. Sistem mengambil task dari project.
5. Sistem mengembalikan daftar task.

### Alternative Flows

- Actor bukan member: sistem mengembalikan 403.
- Project tidak ditemukan: sistem mengembalikan 404.

### Postconditions

- Tidak ada perubahan data.

## UC-007 Update Task Status

### Actor

Task Assignee atau Project Owner

### Preconditions

- Actor sudah login.
- Task sudah ada.
- Actor adalah assignee task atau owner project.

### Main Flow

1. Actor mengirim status baru.
2. Sistem memvalidasi token.
3. Sistem mengambil task dan project.
4. Sistem mengecek permission actor.
5. Sistem memvalidasi nilai status.
6. Sistem memperbarui status task.
7. Sistem mengembalikan data task terbaru.

### Alternative Flows

- Actor tidak punya permission: sistem mengembalikan 403.
- Status tidak valid: sistem mengembalikan 400.
- Task tidak ditemukan: sistem mengembalikan 404.

### Postconditions

- Status task berubah.

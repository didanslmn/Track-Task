# Business Rules

## User Rules

| ID | Rule |
| --- | --- |
| UR-001 | Email user harus unik. |
| UR-002 | Password tidak boleh disimpan dalam plain text. |
| UR-003 | Response API tidak boleh mengembalikan password hash. |
| UR-004 | User harus login untuk mengakses project dan task. |

## Project Rules

| ID | Rule |
| --- | --- |
| PR-001 | Project harus memiliki name. |
| PR-002 | User yang membuat project otomatis menjadi owner. |
| PR-003 | Satu project minimal memiliki satu owner. |
| PR-004 | User hanya dapat melihat project jika menjadi member project tersebut. |
| PR-005 | Owner dapat menambahkan member ke project. |
| PR-006 | Non-owner tidak dapat menambahkan member ke project. |

## Membership Rules

| ID | Rule |
| --- | --- |
| MR-001 | Kombinasi project dan user pada membership harus unik. |
| MR-002 | Role membership MVP hanya `owner` dan `member`. |
| MR-003 | User yang bukan member project tidak boleh mengakses task project. |
| MR-004 | Target user harus registered user sebelum dapat ditambahkan sebagai member. |

## Task Rules

| ID | Rule |
| --- | --- |
| TR-001 | Task harus memiliki title. |
| TR-002 | Task dibuat dengan status default `todo`. |
| TR-003 | Status task hanya boleh `todo`, `in_progress`, atau `done`. |
| TR-004 | Task harus selalu terkait dengan satu project. |
| TR-005 | Assignee task bersifat opsional. |
| TR-006 | Jika assignee diisi, assignee harus member dari project task tersebut. |
| TR-007 | Member project dapat membuat task. |
| TR-008 | Status task hanya dapat diubah oleh assignee atau project owner. |

## Audit Rules For MVP

| ID | Rule |
| --- | --- |
| AR-001 | Semua tabel utama harus memiliki `created_at`. |
| AR-002 | Data yang dapat diperbarui harus memiliki `updated_at`. |
| AR-003 | Hard delete boleh digunakan untuk MVP, soft delete dapat ditambahkan pada fase berikutnya. |

# Tech Stack Decision

## Language: Go

### Reason

Go cocok untuk backend API karena:

- Cepat dan ringan.
- Standar library HTTP kuat.
- Cocok untuk service yang dideploy sebagai single binary.
- Banyak digunakan di backend, platform, infrastructure, dan cloud engineering.
- Mudah dibaca untuk portfolio.

## Router: Chi

### Reason

Chi dipilih karena:

- Ringan.
- Idiomatik untuk Go.
- Middleware-friendly.
- Tidak terlalu opinionated.
- Cocok untuk REST API kecil sampai menengah.

## Database: PostgreSQL

### Reason

PostgreSQL dipilih karena:

- Relational model cocok untuk users, projects, members, dan tasks.
- Constraint seperti unique key dan foreign key kuat.
- Banyak dipakai di industri.
- Mudah dideploy di banyak cloud provider.

## SQL Access: pgx

### Reason

pgx dipilih karena:

- Driver PostgreSQL modern untuk Go.
- Performa baik.
- Mendukung connection pool.
- Memberi kontrol SQL yang jelas.

## Migration: golang-migrate

### Reason

Migration dibutuhkan agar schema database versioned dan bisa dijalankan ulang di environment berbeda.

## Authentication: JWT

### Reason

JWT cocok untuk MVP karena:

- Stateless.
- Mudah digunakan oleh frontend/mobile.
- Umum dipakai pada REST API.

## Password Hashing: bcrypt

### Reason

Password tidak boleh disimpan sebagai plain text. bcrypt adalah pilihan umum dan aman untuk password hashing.

## Configuration: Environment Variables

### Reason

Environment variables membuat aplikasi bisa berjalan di local, CI, dan production tanpa hardcoded secret.

## Rejected Alternatives

| Alternative | Reason Not Selected |
| --- | --- |
| Gin | Bagus, tapi Chi lebih minimal dan idiomatik untuk portfolio ini |
| ORM penuh seperti GORM | SQL eksplisit lebih bagus untuk belajar database dan authorization |
| MySQL | PostgreSQL lebih kaya fitur dan umum untuk backend modern |
| Session-based auth | JWT lebih mudah untuk API frontend/mobile |
| Microservices | Terlalu kompleks untuk MVP |

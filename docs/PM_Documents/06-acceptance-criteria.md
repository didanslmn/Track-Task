# Acceptance Criteria

## Authentication

### Register

Given user belum memiliki akun, when user mengirim name, email, dan password valid, then sistem membuat akun baru dan mengembalikan data user tanpa password.

Given email sudah digunakan, when user register dengan email yang sama, then sistem mengembalikan error 409.

Given password kurang dari minimum length, when user register, then sistem mengembalikan error 400.

### Login

Given credential valid, when user login, then sistem mengembalikan access token.

Given credential salah, when user login, then sistem mengembalikan error 401.

## Project

### Create Project

Given user sudah login, when user membuat project dengan name valid, then sistem membuat project dan menjadikan user sebagai owner.

Given user tidak login, when user membuat project, then sistem mengembalikan error 401.

### List Projects

Given user sudah login, when user meminta daftar project, then sistem hanya mengembalikan project yang user ikuti.

### Project Detail

Given user adalah member project, when user meminta detail project, then sistem mengembalikan detail project.

Given user bukan member project, when user meminta detail project, then sistem mengembalikan error 403.

## Membership

### Add Member

Given user adalah owner project, when user menambahkan registered user sebagai member, then sistem membuat membership baru.

Given user bukan owner project, when user menambahkan member, then sistem mengembalikan error 403.

Given target user sudah menjadi member, when owner menambahkan user yang sama, then sistem mengembalikan error 409.

## Task

### Create Task

Given user adalah member project, when user membuat task dengan title valid, then sistem membuat task dengan status `todo`.

Given user bukan member project, when user membuat task, then sistem mengembalikan error 403.

Given title kosong, when user membuat task, then sistem mengembalikan error 400.

### Assign Task

Given assignee adalah member project, when task dibuat atau diperbarui dengan assignee tersebut, then sistem menyimpan assignee.

Given assignee bukan member project, when task dibuat atau diperbarui, then sistem mengembalikan error 400.

### Update Task Status

Given user adalah assignee task, when user mengubah status task ke nilai valid, then sistem memperbarui status task.

Given user adalah owner project, when user mengubah status task ke nilai valid, then sistem memperbarui status task.

Given status tidak valid, when user mengubah status task, then sistem mengembalikan error 400.

Given user tidak punya permission, when user mengubah status task, then sistem mengembalikan error 403.

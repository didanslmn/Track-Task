# TeamTask API - System Analysis Overview

## Purpose

Dokumen System Analysis ini menerjemahkan kebutuhan produk TeamTask API menjadi rancangan sistem yang cukup detail untuk dikerjakan oleh Backend Developer.

## Source Documents

Dokumen ini dibuat berdasarkan artefak Product/Project Manager:

- Product brief
- PRD
- Roadmap
- Sprint plan
- User stories
- Acceptance criteria

## System Scope

TeamTask API adalah backend REST API untuk:

- Authentication
- Project management
- Project membership
- Task management
- Task assignment
- Task status tracking

## Primary Actors

- Guest
- Authenticated User
- Project Owner
- Project Member
- Task Assignee

## Main System Modules

- Auth module
- User module
- Project module
- Membership module
- Task module
- Authorization module

## Key Analysis Outputs

- Business requirements
- Use cases
- Business rules
- Entity relationship design
- Data dictionary
- API contract
- Authorization matrix
- Validation and error rules
- Traceability matrix

## Analysis Assumptions

- API menggunakan JSON.
- Authentication menggunakan JWT.
- Database menggunakan PostgreSQL.
- ID menggunakan UUID atau auto-increment integer, keputusan final ditentukan saat technical design.
- Frontend belum tersedia pada fase MVP.
- API harus bisa diuji menggunakan Postman, Swagger, atau HTTP client lain.

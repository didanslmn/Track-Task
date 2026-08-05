# Risk Log

| ID | Risk | Impact | Probability | Mitigation | Owner |
| --- | --- | --- | --- | --- | --- |
| R-001 | Scope MVP melebar | High | Medium | Kunci scope MVP dan pindahkan fitur tambahan ke future enhancement | Product Manager |
| R-002 | Authorization antar project salah | High | Medium | Buat test khusus untuk non-member access dan role owner | Backend Developer |
| R-003 | Dokumentasi API tidak sinkron | Medium | Medium | Update docs setiap endpoint berubah dan gunakan Swagger/OpenAPI | Backend Developer |
| R-004 | Deployment terlambat | Medium | Medium | Siapkan Docker dan environment config sejak sprint awal | Project Manager |
| R-005 | Database design berubah besar di tengah jalan | Medium | Low | Review ERD sebelum coding fitur utama | System Analyst |
| R-006 | Test coverage terlalu minim | Medium | Medium | Prioritaskan integration test untuk auth, project, membership, dan task | Backend Developer |
| R-007 | Secret/token bocor ke repository | High | Low | Gunakan `.env`, `.gitignore`, dan environment variable di platform deploy | Backend Developer |

## Risk Review Cadence

Risk log direview pada akhir setiap sprint. Risk baru ditambahkan jika muncul blocker, perubahan scope, atau masalah teknis yang memengaruhi timeline.

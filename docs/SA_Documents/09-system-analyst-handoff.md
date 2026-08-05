# System Analyst Handoff

## Handoff Purpose

Dokumen ini merangkum keputusan System Analyst yang perlu dipakai Backend Developer saat mulai development.

## MVP Modules

- Auth
- Users
- Projects
- Project Members
- Tasks

## Recommended Development Order

1. Setup backend project
2. Setup database and migration
3. Implement users table
4. Implement auth register/login
5. Implement JWT middleware
6. Implement projects table and create project
7. Implement project_members table and owner assignment
8. Implement project access authorization
9. Implement add member
10. Implement tasks table
11. Implement create/list/update task
12. Implement task status authorization
13. Add tests
14. Add API docs
15. Deploy

## Key Decisions

- Project access is controlled by `project_members`.
- Project creator becomes `owner`.
- Membership role values for MVP: `owner`, `member`.
- Task status values for MVP: `todo`, `in_progress`, `done`.
- Assignee must be a project member.
- JWT handles authentication only.
- Resource-level authorization must be checked against database.

## Open Questions For Technical Design

| Question | Recommended Default |
| --- | --- |
| UUID or integer ID? | UUID for public API safety |
| Soft delete or hard delete? | Hard delete for MVP |
| Refresh token needed? | Not for MVP |
| Pagination needed? | Optional for MVP, recommended for tasks/projects |
| Swagger generated or manually written? | Generated if framework supports it |

## Backend Developer Checklist

- Do not return password hash in response.
- Do not allow cross-project access.
- Validate assignee membership before saving task.
- Use consistent error response format.
- Add integration tests for authorization.
- Keep environment values outside source code.
- Prepare local Docker setup early.

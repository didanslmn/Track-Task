# Requirement Traceability Matrix

| Business Requirement | User Story | Use Case | API Endpoint | Notes |
| --- | --- | --- | --- | --- |
| BR-001 User Account | US-001, US-002 | UC-001, UC-002 | POST /auth/register, POST /auth/login | Account and login flow |
| BR-002 Secure Authentication | US-002 | UC-002 | Protected endpoints | JWT required |
| BR-003 Project Ownership | US-003 | UC-003 | POST /projects | Creator becomes owner |
| BR-004 Project Membership | US-004, US-005 | UC-003, UC-006 | GET /projects, GET /projects/{project_id} | Access based on membership |
| BR-005 Member Invitation | US-006 | UC-004 | POST /projects/{project_id}/members | Owner only |
| BR-006 Task Creation | US-007 | UC-005 | POST /projects/{project_id}/tasks | Member only |
| BR-007 Task Assignment | US-008 | UC-005 | POST /projects/{project_id}/tasks, PATCH /tasks/{task_id} | Assignee must be project member |
| BR-008 Task Status Tracking | US-010 | UC-007 | PATCH /tasks/{task_id}/status | Status enum |
| BR-009 Access Isolation | US-005, US-009 | UC-006, UC-007 | All project/task endpoints | Prevent cross-project access |
| BR-010 API Documentation | All | All | API docs | Should be produced before release |

## Purpose

Traceability matrix membantu memastikan setiap business requirement memiliki:

- User story
- Use case
- API implementation target
- Test coverage target

## Test Planning Notes

Minimal test harus mencakup:

- Register success and duplicate email
- Login success and invalid credential
- Create project success
- Non-member cannot view project
- Owner can add member
- Non-owner cannot add member
- Member can create task
- Non-member cannot create task
- Assignee can update task status
- Unauthorized user cannot update task status

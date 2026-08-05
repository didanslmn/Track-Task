# API Contract

## General

Base URL untuk local development:

```text
http://localhost:8080
```

Format request dan response:

```text
Content-Type: application/json
```

Protected endpoints menggunakan header:

```text
Authorization: Bearer <access_token>
```

## Standard Success Response

```json
{
  "data": {}
}
```

## Standard Error Response

```json
{
  "error": {
    "code": "VALIDATION_ERROR",
    "message": "Validation failed",
    "details": []
  }
}
```

## Health

### GET /health

Response 200:

```json
{
  "data": {
    "status": "ok"
  }
}
```

## Authentication

### POST /auth/register

Request:

```json
{
  "name": "Budi Santoso",
  "email": "budi@example.com",
  "password": "password123"
}
```

Response 201:

```json
{
  "data": {
    "id": "user-id",
    "name": "Budi Santoso",
    "email": "budi@example.com",
    "created_at": "2026-06-25T10:00:00Z"
  }
}
```

### POST /auth/login

Request:

```json
{
  "email": "budi@example.com",
  "password": "password123"
}
```

Response 200:

```json
{
  "data": {
    "access_token": "jwt-token",
    "token_type": "Bearer",
    "user": {
      "id": "user-id",
      "name": "Budi Santoso",
      "email": "budi@example.com"
    }
  }
}
```

## Projects

### POST /projects

Auth: Required

Request:

```json
{
  "name": "Website Redesign",
  "description": "Project untuk redesign website company profile"
}
```

Response 201:

```json
{
  "data": {
    "id": "project-id",
    "name": "Website Redesign",
    "description": "Project untuk redesign website company profile",
    "role": "owner",
    "created_at": "2026-06-25T10:00:00Z"
  }
}
```

### GET /projects

Auth: Required

Response 200:

```json
{
  "data": [
    {
      "id": "project-id",
      "name": "Website Redesign",
      "description": "Project untuk redesign website company profile",
      "role": "owner",
      "created_at": "2026-06-25T10:00:00Z"
    }
  ]
}
```

### GET /projects/{project_id}

Auth: Required

Response 200:

```json
{
  "data": {
    "id": "project-id",
    "name": "Website Redesign",
    "description": "Project untuk redesign website company profile",
    "role": "owner",
    "created_at": "2026-06-25T10:00:00Z",
    "updated_at": "2026-06-25T10:00:00Z"
  }
}
```

## Membership

### POST /projects/{project_id}/members

Auth: Required

Permission: Project Owner

Request:

```json
{
  "user_email": "member@example.com"
}
```

Response 201:

```json
{
  "data": {
    "id": "membership-id",
    "project_id": "project-id",
    "user": {
      "id": "member-user-id",
      "name": "Member User",
      "email": "member@example.com"
    },
    "role": "member",
    "created_at": "2026-06-25T10:00:00Z"
  }
}
```

### GET /projects/{project_id}/members

Auth: Required

Permission: Project Member

Response 200:

```json
{
  "data": [
    {
      "id": "membership-id",
      "user": {
        "id": "user-id",
        "name": "Budi Santoso",
        "email": "budi@example.com"
      },
      "role": "owner"
    }
  ]
}
```

## Tasks

### POST /projects/{project_id}/tasks

Auth: Required

Permission: Project Member

Request:

```json
{
  "title": "Create login endpoint",
  "description": "Implement register and login flow",
  "assignee_id": "user-id",
  "due_date": "2026-07-05"
}
```

Response 201:

```json
{
  "data": {
    "id": "task-id",
    "project_id": "project-id",
    "title": "Create login endpoint",
    "description": "Implement register and login flow",
    "status": "todo",
    "assignee": {
      "id": "user-id",
      "name": "Budi Santoso",
      "email": "budi@example.com"
    },
    "due_date": "2026-07-05",
    "created_at": "2026-06-25T10:00:00Z"
  }
}
```

### GET /projects/{project_id}/tasks

Auth: Required

Permission: Project Member

Optional query:

```text
status=todo
assignee_id=user-id
```

Response 200:

```json
{
  "data": [
    {
      "id": "task-id",
      "title": "Create login endpoint",
      "status": "todo",
      "assignee": {
        "id": "user-id",
        "name": "Budi Santoso"
      },
      "due_date": "2026-07-05"
    }
  ]
}
```

### PATCH /tasks/{task_id}

Auth: Required

Permission: Project Member

Request:

```json
{
  "title": "Create auth endpoint",
  "description": "Implement register and login",
  "assignee_id": "user-id",
  "due_date": "2026-07-06"
}
```

Response 200:

```json
{
  "data": {
    "id": "task-id",
    "title": "Create auth endpoint",
    "description": "Implement register and login",
    "status": "todo",
    "due_date": "2026-07-06",
    "updated_at": "2026-06-25T11:00:00Z"
  }
}
```

### PATCH /tasks/{task_id}/status

Auth: Required

Permission: Task Assignee or Project Owner

Request:

```json
{
  "status": "in_progress"
}
```

Response 200:

```json
{
  "data": {
    "id": "task-id",
    "status": "in_progress",
    "updated_at": "2026-06-25T11:00:00Z"
  }
}
```

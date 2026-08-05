# Validation and Error Rules

## Standard HTTP Status Codes

| Status | Meaning | Usage |
| --- | --- | --- |
| 200 | OK | Successful read or update |
| 201 | Created | Successful resource creation |
| 400 | Bad Request | Validation error |
| 401 | Unauthorized | Missing or invalid token, invalid login |
| 403 | Forbidden | Authenticated but no permission |
| 404 | Not Found | Resource not found |
| 409 | Conflict | Duplicate resource or state conflict |
| 500 | Internal Server Error | Unexpected server error |

## Error Codes

| Code | HTTP Status | Description |
| --- | --- | --- |
| VALIDATION_ERROR | 400 | Request body or parameter is invalid |
| UNAUTHORIZED | 401 | Token missing, token invalid, or login failed |
| FORBIDDEN | 403 | User does not have permission |
| NOT_FOUND | 404 | Resource does not exist |
| EMAIL_ALREADY_EXISTS | 409 | Email already registered |
| MEMBER_ALREADY_EXISTS | 409 | User already member of project |
| INTERNAL_ERROR | 500 | Unexpected server error |

## Field Validation

### Register

| Field | Rule |
| --- | --- |
| name | Required, max 100 characters |
| email | Required, valid email format, max 255 characters, unique |
| password | Required, minimum 8 characters |

### Login

| Field | Rule |
| --- | --- |
| email | Required, valid email format |
| password | Required |

### Create Project

| Field | Rule |
| --- | --- |
| name | Required, max 150 characters |
| description | Optional |

### Add Member

| Field | Rule |
| --- | --- |
| user_email | Required, valid email format, must exist as registered user |

### Create/Update Task

| Field | Rule |
| --- | --- |
| title | Required on create, max 200 characters |
| description | Optional |
| assignee_id | Optional, must be project member if provided |
| due_date | Optional, date format YYYY-MM-DD |

### Update Task Status

| Field | Rule |
| --- | --- |
| status | Required, allowed values: todo, in_progress, done |

## Error Response Examples

### Validation Error

```json
{
  "error": {
    "code": "VALIDATION_ERROR",
    "message": "Validation failed",
    "details": [
      {
        "field": "title",
        "message": "Title is required"
      }
    ]
  }
}
```

### Forbidden

```json
{
  "error": {
    "code": "FORBIDDEN",
    "message": "You do not have permission to access this resource"
  }
}
```

### Conflict

```json
{
  "error": {
    "code": "MEMBER_ALREADY_EXISTS",
    "message": "User is already a member of this project"
  }
}
```

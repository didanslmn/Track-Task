# Authorization Matrix

## Roles

| Role | Description |
| --- | --- |
| Guest | User yang belum login |
| Authenticated User | User yang sudah login |
| Project Member | User yang menjadi member project |
| Project Owner | User dengan role owner pada project |
| Task Assignee | User yang ditugaskan pada task |

## Matrix

| Action | Guest | Authenticated User | Project Member | Project Owner | Task Assignee |
| --- | --- | --- | --- | --- | --- |
| Register | Yes | No Need | No Need | No Need | No Need |
| Login | Yes | No Need | No Need | No Need | No Need |
| Create project | No | Yes | Yes | Yes | Yes |
| List own projects | No | Yes | Yes | Yes | Yes |
| View project detail | No | No | Yes | Yes | Yes, if member |
| Add project member | No | No | No | Yes | No |
| List project members | No | No | Yes | Yes | Yes, if member |
| Create task | No | No | Yes | Yes | Yes, if member |
| List project tasks | No | No | Yes | Yes | Yes, if member |
| Update task fields | No | No | Yes | Yes | Yes, if member |
| Update task status | No | No | No | Yes | Yes |

## Authorization Rules Detail

### Project Access

User can access project data only if a `project_members` record exists for:

```text
project_id = requested project ID
user_id = current user ID
```

### Owner Permission

User is project owner if membership role is:

```text
role = owner
```

### Task Access

User can access task if:

- Task belongs to a project.
- User is member of that project.

### Task Status Update Permission

User can update task status if:

- User is project owner, or
- User is the assignee of the task.

## Security Notes

- Never trust project ID or user ID from request body without checking database membership.
- JWT only proves identity. It does not prove project permission.
- Authorization must be checked per resource, especially for project and task endpoints.

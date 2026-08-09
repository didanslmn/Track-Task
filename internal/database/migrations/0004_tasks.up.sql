-- BE-012: Create tasks table

CREATE TABLE IF NOT EXISTS tasks(
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title           VARCHAR(150) NOT NULL,
    description     TEXT,
    status          VARCHAR(20) NOT NULL DEFAULT 'todo',
    project_id      UUID NOT NULL REFERENCES project(id) ON DELETE CASCADE,
    assignee_id     UUID REFERENCES users(id) ON DELETE SET NULL,
    due_date        DATE,
    created_at      TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT chk_task_status CHECK (status IN ('todo', 'in_progress', 'done'))
);

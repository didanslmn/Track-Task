CREATE TYPE project_member_role AS ENUM ('owner', 'member');
ALTER TABLE project_member ADD COLUMN IF NOT EXISTS role project_member_role NOT NULL DEFAULT 'member';

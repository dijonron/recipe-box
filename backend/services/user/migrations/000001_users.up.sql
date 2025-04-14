CREATE TABLE IF NOT EXISTS users (
  email TEXT PRIMARY KEY,
  name TEXT,
  password_hash TEXT NOT NULL,
  tenant_id TEXT,
  role TEXT NOT NULL DEFAULT 'user',
  last_login TIMESTAMP,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP,
  UNIQUE(tenant_id, email)
);


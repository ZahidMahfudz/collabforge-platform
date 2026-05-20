-- =========================================
-- USERS TABLE
-- =========================================

CREATE TABLE IF NOT EXISTS users (
    id VARCHAR(20) PRIMARY KEY,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100),
    mid_name VARCHAR(100),
    username VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash TEXT,
    provider VARCHAR(20) NOT NULL DEFAULT 'local',
    provider_id TEXT,
    bio TEXT,
    avatar_url TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    
    CONSTRAINT users_provider_check CHECK (
    provider IN ('local', 'google')
)
);

-- =========================================
-- INDEXES
-- =========================================

CREATE INDEX idx_users_email ON users (email);

CREATE INDEX idx_users_username ON users (username);

-- koneksi postgre url = "postgres://postgres:zahid1@localhost:5432/colabforge_db?sslmode=disable"
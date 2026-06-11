-- tabel file untuk menyimpan file yang disimpan di minio
create table files (
    id VARCHAR(20) PRIMARY KEY,
    uploader_id VARCHAR(20) NOT NULL,
    bucket_name VARCHAR(255) NOT NULL,
    object_key TEXT NOT NULL,
    original_name TEXT NOT NULL,
    mime_type VARCHAR(255),
    file_size BIGINT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_files_uploader
    FOREIGN KEY(uploader_id)
    REFERENCES users(id)
    ON DELETE CASCADE,

    CONSTRAINT unique_object_key UNIQUE (object_key)
);

-- INDEX UNTUK CEPAT MENGAMBIL FILE BERDASARKAN UPLOADER_ID
CREATE INDEX idx_files_uploader_id ON files (uploader_id);
-- INDEX UNTUK CEPAT MENGAMBIL FILE BERDASARKAN OBJECT_KEY
CREATE INDEX idx_files_object_key ON files (object_key);

-- PERUBAHAN TABEL USERS
ALTER TABLE users
DROP COLUMN IF EXISTS bio,
DROP COLUMN IF EXISTS avatar_url;

-- TABEL USER_PROFILE
CREATE TABLE user_profiles (
    id VARCHAR(20) PRIMARY KEY,
    user_id VARCHAR(20) NOT NULL UNIQUE,
    headline VARCHAR(255),
    bio TEXT,
    avatar_file_id VARCHAR(20),
    cover_file_id VARCHAR(20),
    phone_number VARCHAR(20),
    country VARCHAR(100),
    province VARCHAR(100),
    city VARCHAR(100),
    website_url TEXT,
    linkedin_url TEXT,
    github_url TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_user_profiles_user
    FOREIGN KEY(user_id)
    REFERENCES users(id)
    ON DELETE CASCADE,

    CONSTRAINT fk_user_profiles_avatar_file
    FOREIGN KEY(avatar_file_id)
    REFERENCES files(id)
    ON DELETE SET NULL,

    CONSTRAINT fk_user_profiles_cover_file
    FOREIGN KEY(cover_file_id)
    REFERENCES files(id)
);

-- INDEX UNTUK CEPAT MENGAMBIL USER_PROFILE BERDASARKAN USER_ID
CREATE INDEX idx_user_profiles_user_id ON user_profiles (user_id);

-- TABEL USER_FOLLOWS
CREATE TABLE user_follows (
    id VARCHAR(20) PRIMARY KEY,
    follower_id VARCHAR(20) NOT NULL,
    following_id VARCHAR(20) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_user_follows_follower
    FOREIGN KEY(follower_id)
    REFERENCES users(id)
    ON DELETE CASCADE,

    CONSTRAINT fk_user_follows_following
    FOREIGN KEY(following_id)
    REFERENCES users(id)
    ON DELETE CASCADE,

    CONSTRAINT unique_follow UNIQUE (follower_id, following_id),

    CONSTRAINT user_follows_no_self_follow CHECK (follower_id <> following_id)
);

-- INDEX UNTUK CEPAT MENGAMBIL FOLLOWER BERDASARKAN FOLLOWING_ID
CREATE INDEX idx_user_follows_following_id ON user_follows (following_id);

-- INDEX UNTUK CEPAT MENGAMBIL FOLLOWING BERDASARKAN FOLLOWER_ID
CREATE INDEX idx_user_follows_follower_id ON user_follows (follower_id);

-- TABEL SKILLS
CREATE TABLE skills (
    id VARCHAR(20) PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- INDEX UNTUK CEPAT MENGAMBIL SKILL BERDASARKAN NAMA
CREATE INDEX idx_skills_name ON skills (name);

-- TABEL USER_SKILLS
CREATE TABLE user_skills (
    id VARCHAR(20) PRIMARY KEY,
    user_id VARCHAR(20) NOT NULL,
    skill_id VARCHAR(20) NOT NULL,
    level VARCHAR(50),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_user_skills_user
    FOREIGN KEY(user_id)
    REFERENCES users(id)
    ON DELETE CASCADE,

    CONSTRAINT fk_user_skills_skill
    FOREIGN KEY(skill_id)
    REFERENCES skills(id)
    ON DELETE CASCADE,

    CONSTRAINT unique_user_skill UNIQUE (user_id, skill_id),
    
    CONSTRAINT user_skills_level_check CHECK (level IN ('Beginner', 'Intermediate', 'Advanced', 'Expert'))
);

-- INDEX UNTUK CEPAT MENGAMBIL SKILL BERDASARKAN USER_ID
CREATE INDEX idx_user_skills_user_id ON user_skills (user_id);

-- INDEX UNTUK CEPAT MENGAMBIL USER BERDASARKAN SKILL_ID
CREATE INDEX idx_user_skills_skill_id ON user_skills (skill_id);

-- TABEL LANGUAGES
CREATE TABLE languages (
    id VARCHAR(20) PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- INDEX UNTUK CEPAT MENGAMBIL LANGUAGE BERDASARKAN NAMA
CREATE INDEX idx_languages_name ON languages (name);

-- TABEL USER_LANGUAGES
CREATE TABLE user_languages (
    id VARCHAR(20) PRIMARY KEY,
    user_id VARCHAR(20) NOT NULL,
    language_id VARCHAR(20) NOT NULL,
    proficiency VARCHAR(50),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_user_languages_user
    FOREIGN KEY(user_id)
    REFERENCES users(id)
    ON DELETE CASCADE,

    CONSTRAINT fk_user_languages_language
    FOREIGN KEY(language_id)
    REFERENCES languages(id)
    ON DELETE CASCADE,

    CONSTRAINT unique_user_language UNIQUE (user_id, language_id),

    CONSTRAINT user_languages_proficiency_check CHECK (proficiency IN ('Beginner', 'Intermediate', 'Advanced', 'Native'))
);

-- INDEX UNTUK CEPAT MENGAMBIL LANGUAGE BERDASARKAN USER_ID
CREATE INDEX idx_user_languages_user_id ON user_languages (user_id);

-- INDEX UNTUK CEPAT MENGAMBIL USER BERDASARKAN LANGUAGE_ID
CREATE INDEX idx_user_languages_language_id ON user_languages (language_id);

-- TABEL USER_EXPERIENCES
CREATE TABLE user_experiences (
    id VARCHAR(20) PRIMARY KEY,
    user_id VARCHAR(20) NOT NULL,
    company_name VARCHAR(255) NOT NULL,
    position VARCHAR(255) NOT NULL,
    employment_type VARCHAR(50),
    location VARCHAR(255),
    description TEXT,
    start_date DATE,
    end_date DATE,
    is_current BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_user_experiences_user
    FOREIGN KEY(user_id)
    REFERENCES users(id)
    ON DELETE CASCADE
);

-- INDEX UNTUK CEPAT MENGAMBIL EXPERIENCE BERDASARKAN USER_ID
CREATE INDEX idx_user_experiences_user_id ON user_experiences (user_id);

-- INDEX UNTUK MENGAMBIL EXPERIENCE TERBARU BERDASARKAN USER_ID DAN START_DATE
CREATE INDEX idx_user_experiences_user_start_date
ON user_experiences(user_id, start_date DESC);

-- TABEL USER_EDUCATIONS
CREATE TABLE user_educations (
    id VARCHAR(20) PRIMARY KEY,
    user_id VARCHAR(20) NOT NULL,
    institution_name VARCHAR(255) NOT NULL,
    degree VARCHAR(255),
    field_of_study VARCHAR(255),
    description TEXT,
    start_date DATE,
    end_date DATE,
    is_current BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_user_educations_user
    FOREIGN KEY(user_id)
    REFERENCES users(id)
    ON DELETE CASCADE
);

-- INDEX UNTUK CEPAT MENGAMBIL EDUCATION BERDASARKAN USER_ID
CREATE INDEX idx_user_educations_user_id ON user_educations (user_id);

-- INDEX UNTUK MENGAMBIL EDUCATION TERBARU BERDASARKAN USER_ID DAN START_DATE
CREATE INDEX idx_user_educations_user_start_date
ON user_educations(user_id, start_date DESC);

-- TABEL USER_CERTIFICATIONS
CREATE TABLE user_certifications (
    id VARCHAR(20) PRIMARY KEY,
    user_id VARCHAR(20) NOT NULL,
    title VARCHAR(255) NOT NULL,
    issuer VARCHAR(255),
    issue_date DATE,
    expiration_date DATE,
    credential_id VARCHAR(255),
    credential_url TEXT,
    certification_file_id VARCHAR(20),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_user_certifications_user
    FOREIGN KEY(user_id)
    REFERENCES users(id)
    ON DELETE CASCADE,

    constraint fk_user_certifications_file
    FOREIGN KEY(certification_file_id)
    REFERENCES files(id)
    ON DELETE SET NULL
);

-- INDEX UNTUK CEPAT MENGAMBIL CERTIFICATION BERDASARKAN USER_ID
CREATE INDEX idx_user_certifications_user_id ON user_certifications (user_id);

-- INDEX UNTUK MENGAMBIL CERTIFICATION TERBARU BERDASARKAN USER_ID DAN ISSUE_DATE
CREATE INDEX idx_user_certifications_issue_date
ON user_certifications(user_id, issue_date DESC);




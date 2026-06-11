-- =========================================
-- DROP USER CERTIFICATIONS
-- =========================================

DROP INDEX IF EXISTS idx_user_certifications_user_id;
DROP INDEX IF EXISTS idx_user_certifications_issue_date;

DROP TABLE IF EXISTS user_certifications;

-- =========================================
-- DROP USER EDUCATIONS
-- =========================================

DROP INDEX IF EXISTS idx_user_educations_user_id;
DROP INDEX IF EXISTS idx_user_educations_user_start_date;

DROP TABLE IF EXISTS user_educations;

-- =========================================
-- DROP USER EXPERIENCES
-- =========================================

DROP INDEX IF EXISTS idx_user_experiences_user_id;
DROP INDEX IF EXISTS idx_user_experiences_user_start_date;
DROP TABLE IF EXISTS user_experiences;

-- =========================================
-- DROP USER LANGUAGES
-- =========================================

DROP INDEX IF EXISTS idx_user_languages_language_id;
DROP INDEX IF EXISTS idx_user_languages_user_id;

DROP TABLE IF EXISTS user_languages;

-- =========================================
-- DROP LANGUAGES
-- =========================================

DROP INDEX IF EXISTS idx_languages_name;

DROP TABLE IF EXISTS languages;

-- =========================================
-- DROP USER SKILLS
-- =========================================

DROP INDEX IF EXISTS idx_user_skills_skill_id;
DROP INDEX IF EXISTS idx_user_skills_user_id;

DROP TABLE IF EXISTS user_skills;

-- =========================================
-- DROP SKILLS
-- =========================================

DROP INDEX IF EXISTS idx_skills_name;

DROP TABLE IF EXISTS skills;

-- =========================================
-- DROP USER FOLLOWS
-- =========================================

DROP INDEX IF EXISTS idx_user_follows_follower_id;
DROP INDEX IF EXISTS idx_user_follows_following_id;

DROP TABLE IF EXISTS user_follows;

-- =========================================
-- DROP USER PROFILES
-- =========================================

DROP INDEX IF EXISTS idx_user_profiles_user_id;

DROP TABLE IF EXISTS user_profiles;

-- =========================================
-- DROP FILES
-- =========================================

DROP INDEX IF EXISTS idx_files_object_key;
DROP INDEX IF EXISTS idx_files_uploader_id;

DROP TABLE IF EXISTS files;

-- =========================================
-- RESTORE USERS TABLE
-- =========================================

ALTER TABLE users
ADD COLUMN IF NOT EXISTS bio TEXT;

ALTER TABLE users
ADD COLUMN IF NOT EXISTS avatar_url TEXT;
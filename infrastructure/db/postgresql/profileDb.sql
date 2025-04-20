-- Enable psql command echoing
\set ECHO all

-- Drop tables in reverse order of creation to handle dependencies
DROP TABLE IF EXISTS project_media;
DROP TABLE IF EXISTS project_skills;
DROP TABLE IF EXISTS testimonials;
DROP TABLE IF EXISTS education;
DROP TABLE IF EXISTS experiences;
DROP TABLE IF EXISTS projects;
DROP TABLE IF EXISTS skills;
DROP TABLE IF EXISTS user_profile;


-- 1. UserProfile Table (or AboutMe)
CREATE TABLE user_profile (
    user_id BIGSERIAL PRIMARY KEY,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    tagline TEXT,
    bio TEXT,
    email TEXT UNIQUE NOT NULL, -- Email should likely be unique
    phone TEXT,
    location TEXT,
    profile_picture_url TEXT,
    resume_url TEXT,
    linkedin_url TEXT,
    github_url TEXT,
    portfolio_url TEXT,
    other_social_links JSONB -- Using JSONB is generally preferred in Postgres
);

COMMENT ON TABLE user_profile IS 'Stores the main profile information for the portfolio owner.';
COMMENT ON COLUMN user_profile.other_social_links IS 'Stores additional social media links as JSON, e.g., {"twitter": "url", "behance": "url"}';


-- 2. Skills Table
CREATE TABLE skills (
    skill_id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE, -- Skill names should be unique
    type TEXT, -- e.g., 'Language', 'Framework', 'Tool', 'Soft Skill'
    icon_url TEXT,
    icon_class TEXT -- e.g., for font awesome class names
);

COMMENT ON TABLE skills IS 'Stores individual skills.';
COMMENT ON COLUMN skills.type IS 'Category of the skill (e.g., Language, Framework, Tool, Soft Skill).';


-- 3. Projects Table
CREATE TABLE projects (
    project_id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES user_profile(user_id) ON DELETE RESTRICT, -- Restrict deletion of user if projects exist
    title TEXT NOT NULL,
    description TEXT,
    start_date DATE,
    end_date DATE,
    project_url TEXT,
    repo_url TEXT,
    featured_image_url TEXT,
    is_featured BOOLEAN DEFAULT FALSE,
    sort_order INTEGER
);

COMMENT ON TABLE projects IS 'Stores details about individual portfolio projects.';
COMMENT ON COLUMN projects.is_featured IS 'Flag to highlight key projects (default false).';


-- 4. Experiences Table (Work History)
CREATE TABLE experiences (
    experience_id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES user_profile(user_id) ON DELETE RESTRICT,
    job_title TEXT NOT NULL,
    company_name TEXT NOT NULL,
    location TEXT,
    start_date DATE NOT NULL,
    end_date DATE, -- NULL if current position
    description TEXT,
    company_url TEXT
);

COMMENT ON TABLE experiences IS 'Stores professional work experience entries.';
COMMENT ON COLUMN experiences.end_date IS 'End date of the employment (NULL if current position).';


-- 5. Education Table
CREATE TABLE education (
    education_id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES user_profile(user_id) ON DELETE RESTRICT,
    institution_name TEXT NOT NULL,
    degree_or_certification TEXT NOT NULL,
    field_of_study TEXT,
    start_date DATE NOT NULL,
    end_date DATE, -- Graduation date or expected end date
    description TEXT -- Relevant coursework, thesis, honors, etc.
);

COMMENT ON TABLE education IS 'Stores educational background entries.';
COMMENT ON COLUMN education.end_date IS 'Graduation date or expected end date.';


-- 6. Testimonials Table (Optional)
CREATE TABLE testimonials (
    testimonial_id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES user_profile(user_id) ON DELETE RESTRICT,
    quote TEXT NOT NULL,
    author_name TEXT NOT NULL,
    author_title TEXT, -- e.g., "CEO at Company X"
    author_company TEXT,
    date_received DATE
);

COMMENT ON TABLE testimonials IS 'Stores quotes or recommendations.';


-- 7. ProjectSkills Table (Junction Table)
CREATE TABLE project_skills (
    project_skill_id BIGSERIAL PRIMARY KEY,
    project_id BIGINT NOT NULL REFERENCES projects(project_id) ON DELETE CASCADE, -- Cascade delete if project is deleted
    skill_id BIGINT NOT NULL REFERENCES skills(skill_id) ON DELETE CASCADE, -- Cascade delete if skill is deleted
    UNIQUE (project_id, skill_id) -- Ensure a skill isn't linked twice to the same project
);

COMMENT ON TABLE project_skills IS 'Links projects to the skills used in them (Many-to-Many relationship).';


-- 8. ProjectMedia Table (Optional)
CREATE TABLE project_media (
    media_id BIGSERIAL PRIMARY KEY,
    project_id BIGINT NOT NULL REFERENCES projects(project_id) ON DELETE CASCADE, -- Cascade delete if project is deleted
    media_url TEXT NOT NULL,
    media_type TEXT NOT NULL, -- e.g., 'Image', 'Video'
    caption TEXT,
    alt_text TEXT, -- Important for accessibility
    sort_order INTEGER
);

COMMENT ON TABLE project_media IS 'Stores multiple media items (images, videos) associated with a project.';
COMMENT ON COLUMN project_media.media_type IS 'Type of media, e.g., Image, Video.';
COMMENT ON COLUMN project_media.alt_text IS 'Alternative text for images, important for accessibility.';


-- --- Indexes ---
-- PostgreSQL automatically creates indexes for PRIMARY KEY and UNIQUE constraints.
-- It's highly recommended to create indexes on foreign key columns manually for performance.

-- Indexes for Foreign Keys in main tables
CREATE INDEX idx_projects_user_id ON projects(user_id);
CREATE INDEX idx_experiences_user_id ON experiences(user_id);
CREATE INDEX idx_education_user_id ON education(user_id);
CREATE INDEX idx_testimonials_user_id ON testimonials(user_id);

-- Indexes for Foreign Keys in junction/dependent tables
CREATE INDEX idx_project_skills_project_id ON project_skills(project_id);
CREATE INDEX idx_project_skills_skill_id ON project_skills(skill_id);
CREATE INDEX idx_project_media_project_id ON project_media(project_id);

-- Optional: Indexes on frequently queried columns
-- CREATE INDEX idx_projects_title ON projects(title); -- Example if searching by title often
-- CREATE INDEX idx_skills_name ON skills(name); -- Already covered by UNIQUE constraint


-- --- End of Script ---
SELECT 'Database schema created successfully.' AS status;

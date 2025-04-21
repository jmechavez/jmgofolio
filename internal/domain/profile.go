package domain

import (
	"encoding/json" // For potential JSON field
	"time"

	"github.com/jmechavez/jmgofolio/errors"
	"github.com/jmechavez/jmgofolio/internal/dto"
)

// UserProfile corresponds to the 'UserProfile' table
type UserProfile struct {
	UserID            int64           `json:"user_id"                       db:"user_id"`
	FirstName         string          `json:"first_name"                    db:"first_name"`
	LastName          string          `json:"last_name"                     db:"last_name"`
	Tagline           string          `json:"tagline"                       db:"tagline"`
	Bio               string          `json:"bio"                           db:"bio"`
	Email             string          `json:"email"                         db:"email"`
	Phone             *string         `json:"phone,omitempty"               db:"phone"` // Pointer for optional field
	Location          string          `json:"location"                      db:"location"`
	ProfilePictureURL *string         `json:"profile_picture_url,omitempty" db:"profile_picture_url"` // Pointer for optional field
	ResumeURL         *string         `json:"resume_url,omitempty"          db:"resume_url"`          // Pointer for optional field
	LinkedInURL       *string         `json:"linkedin_url,omitempty"        db:"linkedin_url"`        // Pointer for optional field
	GithubURL         *string         `json:"github_url,omitempty"          db:"github_url"`          // Pointer for optional field
	PortfolioURL      *string         `json:"portfolio_url,omitempty"       db:"portfolio_url"`       // Pointer for optional field
	OtherSocialLinks  json.RawMessage `json:"other_social_links,omitempty"  db:"other_social_links"`  // Use json.RawMessage for flexibility or string/map[string]interface{}
	// If using an ORM like Gorm, you might add relationship fields here, e.g.:
	// Projects    []Project    `json:"projects,omitempty" gorm:"foreignKey:UserID"`
	// Experiences []Experience `json:"experiences,omitempty" gorm:"foreignKey:UserID"`
	// Education   []Education  `json:"education,omitempty" gorm:"foreignKey:UserID"`
	// Skills      []Skill      `json:"skills,omitempty" gorm:"many2many:user_skills;"` // Example if tracking skills per user
}

// Project corresponds to the 'Projects' table
type Project struct {
	ProjectID        int64      `json:"project_id"                   db:"project_id"`
	UserID           int64      `json:"user_id"                      db:"user_id"` // Foreign Key to UserProfile
	Title            string     `json:"title"                        db:"title"`
	Description      string     `json:"description"                  db:"description"`
	StartDate        *time.Time `json:"start_date,omitempty"         db:"start_date"`         // Pointer for optional field
	EndDate          *time.Time `json:"end_date,omitempty"           db:"end_date"`           // Pointer for optional field
	ProjectURL       *string    `json:"project_url,omitempty"        db:"project_url"`        // Pointer for optional field
	RepoURL          *string    `json:"repo_url,omitempty"           db:"repo_url"`           // Pointer for optional field
	FeaturedImageURL *string    `json:"featured_image_url,omitempty" db:"featured_image_url"` // Pointer for optional field
	IsFeatured       bool       `json:"is_featured"                  db:"is_featured"`
	SortOrder        *int       `json:"sort_order,omitempty"         db:"sort_order"` // Pointer for optional field
	// For related data loaded separately or via ORM:
	Skills []Skill        `json:"skills,omitempty"                                     gorm:"many2many:project_skills;"` // Example for GORM
	Media  []ProjectMedia `json:"media,omitempty"                                      gorm:"foreignKey:ProjectID"`      // Example for GORM
}

// Skill corresponds to the 'Skills' table
type Skill struct {
	SkillID   int64   `json:"skill_id"             db:"skill_id"`
	Name      string  `json:"name"                 db:"name"`
	Type      *string `json:"type,omitempty"       db:"type"`       // Pointer for optional field
	IconURL   *string `json:"icon_url,omitempty"   db:"icon_url"`   // Pointer for optional field
	IconClass *string `json:"icon_class,omitempty" db:"icon_class"` // Pointer for optional field
}

// Experience corresponds to the 'Experiences' table
type Experience struct {
	ExperienceID int64      `json:"experience_id"         db:"experience_id"`
	UserID       int64      `json:"user_id"               db:"user_id"` // Foreign Key to UserProfile
	JobTitle     string     `json:"job_title"             db:"job_title"`
	CompanyName  string     `json:"company_name"          db:"company_name"`
	Location     *string    `json:"location,omitempty"    db:"location"` // Pointer for optional field
	StartDate    time.Time  `json:"start_date"            db:"start_date"`
	EndDate      *time.Time `json:"end_date,omitempty"    db:"end_date"` // Pointer for nullable (current job)
	Description  string     `json:"description"           db:"description"`
	CompanyURL   *string    `json:"company_url,omitempty" db:"company_url"` // Pointer for optional field
}

// Education corresponds to the 'Education' table
type Education struct {
	EducationID           int64      `json:"education_id"             db:"education_id"`
	UserID                int64      `json:"user_id"                  db:"user_id"` // Foreign Key to UserProfile
	InstitutionName       string     `json:"institution_name"         db:"institution_name"`
	DegreeOrCertification string     `json:"degree_or_certification"  db:"degree_or_certification"`
	FieldOfStudy          *string    `json:"field_of_study,omitempty" db:"field_of_study"` // Pointer for optional field
	StartDate             time.Time  `json:"start_date"               db:"start_date"`
	EndDate               *time.Time `json:"end_date,omitempty"       db:"end_date"`    // Pointer for optional/expected
	Description           *string    `json:"description,omitempty"    db:"description"` // Pointer for optional field
}

// --- Linking Table Struct ---

// ProjectSkill corresponds to the 'ProjectSkills' junction table
type ProjectSkill struct {
	ProjectSkillID int64 `json:"project_skill_id" db:"project_skill_id"`
	ProjectID      int64 `json:"project_id"       db:"project_id"` // Foreign Key to Projects
	SkillID        int64 `json:"skill_id"         db:"skill_id"`   // Foreign Key to Skills
}

// --- Optional Table Structs ---

// Testimonial corresponds to the 'Testimonials' table
type Testimonial struct {
	TestimonialID int64      `json:"testimonial_id"           db:"testimonial_id"`
	UserID        int64      `json:"user_id"                  db:"user_id"` // Foreign Key to UserProfile
	Quote         string     `json:"quote"                    db:"quote"`
	AuthorName    string     `json:"author_name"              db:"author_name"`
	AuthorTitle   string     `json:"author_title"             db:"author_title"`
	AuthorCompany *string    `json:"author_company,omitempty" db:"author_company"` // Pointer for optional field
	DateReceived  *time.Time `json:"date_received,omitempty"  db:"date_received"`  // Pointer for optional field
}

// ProjectMedia corresponds to the 'ProjectMedia' table
type ProjectMedia struct {
	MediaID   int64   `json:"media_id"             db:"media_id"`
	ProjectID int64   `json:"project_id"           db:"project_id"` // Foreign Key to Projects
	MediaURL  string  `json:"media_url"            db:"media_url"`
	MediaType string  `json:"media_type"           db:"media_type"` // e.g., "Image", "Video"
	Caption   *string `json:"caption,omitempty"    db:"caption"`    // Pointer for optional field
	AltText   *string `json:"alt_text,omitempty"   db:"alt_text"`   // Pointer for optional field
	SortOrder *int    `json:"sort_order,omitempty" db:"sort_order"` // Pointer for optional field
}

func (p UserProfile) UserProfile() dto.UserProfileResponse {
	return dto.UserProfileResponse{
		UserID:            p.UserID,
		FirstName:         p.FirstName,
		LastName:          p.LastName,
		Tagline:           p.Tagline,
		Bio:               p.Bio,
		Email:             p.Email,
		Phone:             p.Phone,
		Location:          p.Location,
		ProfilePictureURL: p.ProfilePictureURL,
		ResumeURL:         p.ResumeURL,
		LinkedInURL:       p.LinkedInURL,
		GithubURL:         p.GithubURL,
		PortfolioURL:      p.PortfolioURL,
	}
}

type PortfolioRepository interface {
	MyProfile() (*UserProfile, *errors.AppError)
}

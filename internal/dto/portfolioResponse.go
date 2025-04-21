package dto

import "encoding/json"

type UserProfileResponse struct {
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

package ports

import (
	"github.com/jmechavez/jmgofolio/errors"
	"github.com/jmechavez/jmgofolio/internal/domain"
	"github.com/jmechavez/jmgofolio/internal/dto"
)

type PortfolioService interface {
	MyPortfolio() (*dto.UserProfileResponse, *errors.AppError)
}

type DefaultPortfolioService struct {
	repo domain.PortfolioRepository
}

func (s DefaultPortfolioService) MyPortfolio() (*dto.UserProfileResponse, *errors.AppError) {
	profile, err := s.repo.MyProfile()
	if err != nil {
		return nil, err
	}

	response := dto.UserProfileResponse{
		UserID:            profile.UserID,
		FirstName:         profile.FirstName,
		LastName:          profile.LastName,
		Tagline:           profile.Tagline,
		Bio:               profile.Bio,
		Email:             profile.Email,
		Phone:             profile.Phone,
		Location:          profile.Location,
		ProfilePictureURL: profile.ProfilePictureURL,
		ResumeURL:         profile.ResumeURL,
		LinkedInURL:       profile.LinkedInURL,
		GithubURL:         profile.GithubURL,
		PortfolioURL:      profile.PortfolioURL,
	}

	return &response, nil
}

func NewPortfolioService(repository domain.PortfolioRepository) DefaultPortfolioService {
	return DefaultPortfolioService{repository}
}

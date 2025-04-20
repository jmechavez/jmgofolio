package db

import (
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"

	"github.com/jmechavez/jmgofolio/errors"
	"github.com/jmechavez/jmgofolio/infrastructure/logger"
	"github.com/jmechavez/jmgofolio/internal/domain"
)

type PorfolioRepository struct {
	emailDb *sqlx.DB
}

func (r PorfolioRepository) MyProfile() ([]domain.UserProfile, *errors.AppError) {
	var profile []domain.UserProfile
	query := "SELECT * FROM user_profile"
	err := r.emailDb.Get(&profile, query)
	if err != nil {
		logger.Error("Database error while fetching user_profile", zap.Error(err))
		return nil, errors.NewUnExpectedError("Unexpected database error")
	}
	logger.Info("Successfully fetched user_profile", zap.Int("count", len(profile)))
	return profile, nil
}

package db

import (
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"

	"github.com/jmechavez/jmgofolio/errors"
	"github.com/jmechavez/jmgofolio/infrastructure/logger"
	"github.com/jmechavez/jmgofolio/internal/domain"
	_ "github.com/lib/pq"
)

type PortfolioRepository struct {
	db *sqlx.DB
}

func (r PortfolioRepository) MyProfile() (*domain.UserProfile, *errors.AppError) {
	var profile domain.UserProfile
	query := "SELECT * FROM user_profile"
	err := r.db.Get(&profile, query)
	if err != nil {
		logger.Error("Database error while fetching user_profile", zap.Error(err))
		return nil, errors.NewUnExpectedError("Unexpected database error")
	}
	logger.Info("Successfully fetched user_profile")
	return &profile, nil
}

func NewPortfolioRepository(db *sqlx.DB) PortfolioRepository {
	logger.Info("Initializing Database")
	return PortfolioRepository{db}
}

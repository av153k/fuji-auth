package services

import (
	"fmt"
	"fuji-auth/pkg/config"
	"fuji-auth/pkg/models"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabaseService() (DatabaseService, error) {
	dbConfig := config.LoadConfig()
	db, err := gorm.Open(postgres.Open(fmt.Sprintf(`
	host=%s
	user=%s
	password=%s
	dbname=%s`,
		dbConfig.DbHost,
		dbConfig.DbUser,
		dbConfig.DbPassword,
		dbConfig.DbName,
	)), &gorm.Config{})
	if err != nil {
		log.Error("Failed to connect to database: %v", err)
		return DatabaseService{}, err
	}

	db.AutoMigrate(&models.User{})

	return DatabaseService{DB: db}, nil
}

type DatabaseService struct {
	*gorm.DB
}

func (s *DatabaseService) GetUserByEmail(email string) (models.User, error) {
	return models.User{}, nil
}

func (s *DatabaseService) GetUserByPhone(email string) (models.User, error) {
	return models.User{}, nil
}

func (s *DatabaseService) GetUserByID(email string) (models.User, error) {
	return models.User{}, nil
}

func (s *DatabaseService) CreateUser(user *models.User) error {
	log.Debug(fmt.Sprintf("Creating user: %+v", user))
	result := s.DB.Create(user)
	return result.Error
}

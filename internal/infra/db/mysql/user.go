package mysql

import (
	"time"

	"github.com/google/uuid"
	"github.com/raita876/gotask/internal/domain"
	"gorm.io/gorm"
)

type DBUser struct {
	ID        uuid.UUID `gorm:"column:id;type:char(36);primaryKey"`
	Name      string    `gorm:"column:name;not null"`
	Email     string    `gorm:"column:email;not null"`
	Password  string    `gorm:"column:password;not null"`
	CreatedAt time.Time `gorm:"column:created_at;not null"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null"`
}

func (DBUser) TableName() string {
	return "users"
}

func toDBUser(user *domain.User) *DBUser {
	return &DBUser{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func fromDBUser(dbUser *DBUser) *domain.User {
	return &domain.User{
		ID:        dbUser.ID,
		Name:      dbUser.Name,
		Email:     dbUser.Email,
		Password:  dbUser.Password,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
	}
}

type DBUserRepository struct {
	db *gorm.DB
}

func NewDBUserRepository(db *gorm.DB) domain.UserRepository {
	return &DBUserRepository{
		db: db,
	}
}

func (repo *DBUserRepository) Create(user *domain.User) (*domain.User, error) {
	dbUser := toDBUser(user)

	if err := repo.db.Create(dbUser).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (repo *DBUserRepository) FindByID(id uuid.UUID) (*domain.User, error) {
	var dbUser DBUser
	if err := repo.db.First(&dbUser, id).Error; err != nil {
		return nil, err
	}

	return fromDBUser(&dbUser), nil
}

func (repo *DBUserRepository) Update(id uuid.UUID, name, password string) (*domain.User, error) {
	var dbUser DBUser

	// TODO: １つの関数につき発行する SQL は１つとしたい
	if err := repo.db.First(&dbUser, id).Error; err != nil {
		return nil, err
	}

	dbUser.Name = name
	dbUser.Password = password

	if err := repo.db.Model(&DBUser{}).Where("id = ?", dbUser.ID).Updates(&dbUser).Error; err != nil {
		return nil, err
	}

	return fromDBUser(&dbUser), nil
}

func (repo *DBUserRepository) Delete(id uuid.UUID) error {
	return repo.db.Delete(&DBUser{}, id).Error
}

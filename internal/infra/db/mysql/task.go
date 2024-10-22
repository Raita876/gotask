package mysql

import (
	"time"

	"github.com/google/uuid"
	"github.com/raita876/gotask/internal/domain"
	"gorm.io/gorm"
)

type DBTask struct {
	ID        uuid.UUID `gorm:"column:id;type:char(36);primaryKey"`
	Name      string    `gorm:"column:name;not null"`
	Status    uint16    `gorm:"column:status;not null"`
	UserID    uuid.UUID `gorm:"column:user_id;type:char(36);not null"`
	CreatedAt time.Time `gorm:"column:created_at;not null"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null"`
}

func (DBTask) TableName() string {
	return "tasks"
}

func toDBTask(task *domain.Task) *DBTask {
	return &DBTask{
		ID:        task.ID,
		Name:      task.Name,
		Status:    task.Status,
		UserID:    task.UserID,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}
}

func fromDBTask(dbTask *DBTask) *domain.Task {
	return &domain.Task{
		ID:        dbTask.ID,
		Name:      dbTask.Name,
		Status:    dbTask.Status,
		UserID:    dbTask.UserID,
		CreatedAt: dbTask.CreatedAt,
		UpdatedAt: dbTask.UpdatedAt,
	}
}

type DBTaskRepository struct {
	db *gorm.DB
}

func NewDBTaskRepository(db *gorm.DB) domain.TaskRepository {
	return &DBTaskRepository{
		db: db,
	}
}

func (repo *DBTaskRepository) Create(task *domain.Task) (*domain.Task, error) {
	dbTask := toDBTask(task)

	if err := repo.db.Create(dbTask).Error; err != nil {
		return nil, err
	}

	return task, nil
}

func (repo *DBTaskRepository) FindByID(id uuid.UUID) (*domain.Task, error) {
	var dbTask DBTask
	if err := repo.db.First(&dbTask, id).Error; err != nil {
		return nil, err
	}

	return fromDBTask(&dbTask), nil
}

func (repo *DBTaskRepository) FindByUserID(userID uuid.UUID) ([]*domain.Task, error) {
	var dbTasks []DBTask

	if err := repo.db.Find(&dbTasks, userID).Error; err != nil {
		return nil, err
	}

	tasks := make([]*domain.Task, len(dbTasks))
	for i, dbTask := range dbTasks {
		tasks[i] = fromDBTask(&dbTask)
	}

	return tasks, nil
}

func (repo *DBTaskRepository) Update(id uuid.UUID, name string, status uint16) (*domain.Task, error) {
	var dbTask DBTask

	// TODO: １つの関数につき発行する SQL は１つとしたい
	if err := repo.db.First(&dbTask, id).Error; err != nil {
		return nil, err
	}

	dbTask.Name = name
	dbTask.Status = status

	if err := repo.db.Model(&DBTask{}).Where("id = ?", dbTask.ID).Updates(&dbTask).Error; err != nil {
		return nil, err
	}

	return fromDBTask(&dbTask), nil
}

func (repo *DBTaskRepository) Delete(id uuid.UUID) error {
	return repo.db.Delete(&DBTask{}, id).Error
}

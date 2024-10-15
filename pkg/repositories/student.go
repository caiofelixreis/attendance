package repository

import (
	"github.com/caiofelixreis/attendance/pkg/models"
	"gorm.io/gorm"
)

type StudentRepository interface {
	Create(student *models.Student) error
	// FindById(id string) (*StudentModel, error)
	// Update(student *StudentModel) error
	// Delete(id string) error
	// List() ([]*StudentModel, error)
}

type studentRepository struct {
	db *gorm.DB
}

func NewStudentRepository(db *gorm.DB) StudentRepository {
	return &studentRepository{db: db}
}

func (sr *studentRepository) Create(student *models.Student) error {
	creation_response := sr.db.Create(student)
	return creation_response.Error
}

package controllers

import (
	"fmt"
	"sre-goapi/models"
	logger "sre-goapi/utils"

	"gorm.io/gorm"
)

type StudentController struct {
	db *gorm.DB // dependency injection
}

func NewStudentController(appDB *gorm.DB) *StudentController {
	return &StudentController{
		db: appDB,
	}
}

func (s *StudentController) CreateStudent(student *models.Student) (*models.Student, error) {
	res := s.db.Create(student)
	if res.Error != nil {
		logger.Error(res.Error)
		return nil, res.Error
	}

	return student, nil
}

func (s *StudentController) UpdateStudent(id int, student *models.Student) (*models.Student, error) {
	if err := s.db.First(student); err != nil {
		logger.Error(err.Error)
		return nil, err.Error
	}

	if err := s.db.Save(student); err != nil {
		logger.Error(err.Error)
		return nil, err.Error
	}
	return student, nil
}

func (s *StudentController) GetAllStudents() ([]models.Student, error) {
	var students []models.Student
	res := s.db.Find(&students)
	if res.Error != nil {
		logger.Error(res.Error)
		return nil, res.Error
	}

	return students, nil
}

func (s *StudentController) GetStudentById(id int) (*models.Student, error) {
	var student models.Student
	if err := s.db.First(&student, id); err != nil {
		logger.Error(err)
		return nil, fmt.Errorf("ERROR: cannot find student")
	}

	return &student, nil
}

func (s *StudentController) DeleteStudent(studentId int) error {
	if err := s.db.Delete(&models.Student{}, studentId); err != nil {
		logger.Error(err.Error)
		return err.Error
	}

	return nil
}

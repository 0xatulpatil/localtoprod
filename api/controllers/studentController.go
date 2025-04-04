package controllers

import (
	"sre-goapi/db"
	"sre-goapi/models"

	"gorm.io/gorm"
)

type StudentController struct {
	db *gorm.DB // dependency injection
}

func NewStudentController() *StudentController {
	appDB := db.GetAppDB()
	return &StudentController{
		db: appDB,
	}
}

func (s *StudentController) CreateStudent(name string, roll_no int) (*models.Student, error) {
	student := &models.Student{
		Name:   name,
		RollNo: roll_no,
	}

	res := s.db.Create(student)
	if res.Error != nil {
		return nil, res.Error
	}

	return student, nil
}

func (s *StudentController) UpdateStudent(id string, student *models.Student) (*models.Student, error) {
	if err := s.db.First(student); err != nil {
		return nil, err.Error
	}

	if err := s.db.Save(student); err != nil {
		return nil, err.Error
	}
	return student, nil
}

func (s *StudentController) GetAllStudents() ([]models.Student, error) {
	var students []models.Student
	res := s.db.Find(&students)
	if res.Error != nil {
		return nil, res.Error
	}

	return students, nil
}

func (s *StudentController) DeleteStudent(studentId string) error {
	if err := s.db.Delete(&models.Student{}, studentId); err != nil {
		return err.Error
	}

	return nil
}

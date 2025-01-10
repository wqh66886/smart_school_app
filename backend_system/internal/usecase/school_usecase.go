package usecase

import (
	"github.com/wqh/smart/school/system/internal/domain"
	"gorm.io/gorm"
)

type schoolUseCase struct {
	db *gorm.DB
}

func NewSchoolUseCase(db *gorm.DB) *schoolUseCase {
	return &schoolUseCase{
		db: db,
	}
}

func (s *schoolUseCase) GetSchoolByCode(code string) (*domain.School, error) {
	var school domain.School
	if err := s.db.Where("code = ?", code).First(&school).Error; err != nil {
		return nil, err
	}
	return &school, nil
}

func (s *schoolUseCase) SearchAllSchool() ([]domain.School, error) {
	var schools []domain.School
	if err := s.db.Find(&schools).Error; err != nil {
		return nil, err
	}
	return schools, nil
}

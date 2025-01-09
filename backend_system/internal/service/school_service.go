package service

import (
	"github.com/wqh/smart/school/system/internal/domain"
	"github.com/wqh/smart/school/system/internal/initiate"
	"github.com/wqh/smart/school/system/internal/repository"
)

type SchoolService struct {
	schoolRepository repository.SchoolRepository
}

func NewSchoolService(schoolRepository repository.SchoolRepository) *SchoolService {
	return &SchoolService{
		schoolRepository: schoolRepository,
	}
}

func (s *SchoolService) GetSchoolInfoByCode(code string) (*domain.School, error) {
	if code == "" {
		return nil, nil
	}
	school, err := s.schoolRepository.GetSchoolByCode(code)
	if err != nil {
		return nil, initiate.INNER_ERROR
	}
	if school == nil {
		return nil, initiate.NOT_FOUND
	}
	return school, nil
}

func (s *SchoolService) GetAllSchoolInfo() ([]map[string]interface{}, error) {
	schools, err := s.schoolRepository.SearchAllSchool()
	if err != nil {
		return nil, initiate.INNER_ERROR
	}
	mp := make([]map[string]interface{}, 0)
	for _, school := range schools {
		mp = append(mp, map[string]interface{}{
			"id":         school.Id,
			"name":       school.Name,
			"address":    school.Address,
			"phone":      school.Phone,
			"email":      school.Email,
			"website":    school.Website,
			"logo":       school.Logo,
			"createTime": school.CreateTime,
			"updateTime": school.UpdateTime,
		})
	}
	return mp, nil
}

package repository

import "github.com/wqh/smart/school/system/internal/domain"

/**
* description:
* author: wqh
* date: 2025/1/8
 */

type SchoolRepository interface {
	CreateSchool(school *domain.School) error
	GetSchoolByCode(code string) (*domain.School, error)
}

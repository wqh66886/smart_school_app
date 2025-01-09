package repository

import "github.com/wqh/smart/school/system/internal/domain"

/**
* description:
* author: wqh
* date: 2025/1/8
 */

type SchoolRepository interface {
	GetSchoolByCode(code string) (*domain.School, error)
	SearchAllSchool() ([]domain.School, error)
}

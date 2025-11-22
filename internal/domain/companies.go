package domain

import (
	"time"

	"github.com/google/uuid"
)

type CompanyType string

const (
	CompanyTypeCorporation        CompanyType = "Corporation"
	CompanyTypeNonProfit          CompanyType = "NonProfit"
	CompanyTypeCooperative        CompanyType = "Cooperative"
	CompanyTypeSoleProprietorship CompanyType = "Sole Proprietorship"
)

type Company struct { //nolint:decorder // consts sitting right after their type definition
	ID            *uuid.UUID
	Name          *string
	Description   *string
	EmployeeCount *int
	Registered    *bool
	Type          *CompanyType
	CreatedAt     *time.Time
	UpdatedAt     *time.Time
	CreatedBy     *uuid.UUID
	UpdatedBy     *uuid.UUID
}

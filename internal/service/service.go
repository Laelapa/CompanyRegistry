package service

import "time"

type Service struct {
	User    *UserService
	Company *CompanyService
}

const defaultEventPublishTimeout = 5 * time.Second

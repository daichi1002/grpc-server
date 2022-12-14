package usecase

import "clean-architecture/domain"

// インターフェース
type UserRepository interface {
	FindAll() (*domain.Users, error)
	Create(*domain.User) error
	FindById(id string) (*domain.User, error)
}

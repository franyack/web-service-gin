package gateway

import "example/web-service-gin/internal/business/domain"

var authRepository AuthRepository

func RegisterAuthRepository(repository AuthRepository) {
	authRepository = repository
}

func NewAuthRepository() AuthRepository {
	if authRepository == nil {
		panic("authRepository not found")
	}
	return authRepository
}

type AuthRepository interface {
	CreateAuth(userid uint64, tokenDetails *domain.TokenDetails) error
	FetchAuth(authD *domain.AccessDetails) (uint64, error)
	DeleteAuth(givenUuid string) (int64, error)
}

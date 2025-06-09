package user

import (
    domain "passkey-practice/backend/internal/domain/user"
    "passkey-practice/backend/internal/infra/repository"
    "github.com/google/uuid"
)

type RegisterUserInput struct {
    Name     string
    Email    string
    Password string
}

type RegisterUserUseCase struct {
    repo repository.UserRepository
}

func NewRegisterUserUseCase(repo repository.UserRepository) *RegisterUserUseCase {
    return &RegisterUserUseCase{repo: repo}
}

func (uc *RegisterUserUseCase) Register(input RegisterUserInput) (*domain.User, error) {
    email, err := domain.NewEmail(input.Email)
    if err != nil {
        return nil, err
    }
    password, err := domain.NewPasswordFromPlain(input.Password)
    if err != nil {
        return nil, err
    }
    user := domain.NewUser(uuid.New(), input.Name, email, password)
    if err := uc.repo.Save(user); err != nil {
        return nil, err
    }
    return user, nil
} 
package exampletocopy

import "context"

type Usecase struct {
	repo Repository
}

type Repository interface {
	FindAll(ctx context.Context) ([]User, error)
}

func NewUsecase(r Repository) *Usecase {
	return &Usecase{repo: r}
}

func (u *Usecase) GetAll(ctx context.Context) ([]User, error) {
	return u.repo.FindAll(ctx)
}

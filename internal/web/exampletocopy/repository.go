package exampletocopy

import "context"

type User struct {
	ID    int
	Name  string
	Email string
}

type MemoryRepo struct{}

func NewMemoryRepo() *MemoryRepo {
	return &MemoryRepo{}
}

func (r *MemoryRepo) FindAll(ctx context.Context) ([]User, error) {
	return []User{
		{1, "Iyan", "iyan@mail.com"},
		{2, "Budi", "budi@mail.com"},
	}, nil
}

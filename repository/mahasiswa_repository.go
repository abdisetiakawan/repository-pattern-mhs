package repository

import (
	"context"

	"github.com/abdisetiakawan/repository-pattern-mhs/entity"
)

type MahasiswaRepository interface {
	Insert(ctx context.Context, mhs entity.Mahasiswa) (entity.Mahasiswa, error)
	FindById(ctx context.Context, nim string) (entity.Mahasiswa, error)
	FindAll(ctx context.Context) ([]entity.Mahasiswa, error)
	Update(ctx context.Context, mhs entity.Mahasiswa) (entity.Mahasiswa, error)
	Delete(ctx context.Context, nim string) error
}
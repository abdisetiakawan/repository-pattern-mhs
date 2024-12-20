package repository

import (
	"context"
	"fmt"
	"testing"
	"time"

	repository_pattern_mhs "github.com/abdisetiakawan/repository-pattern-mhs"
	"github.com/abdisetiakawan/repository-pattern-mhs/entity"
)

func TestMhsInsert(t *testing.T) {
	mhsRepository := NewMhsRepositoryImpl(repository_pattern_mhs.GetConnection())
	ctx := context.Background()

	birthdate, _ := time.Parse("2006-01-02", "1998-03-06")
	mhs := entity.Mahasiswa{
		Nim:       "2200015112",
		Email:     "H2uZD@example.com",
		Nama:      "Erna Damayanti",
		Jurusan:   "Sistem Informasi",
		Birthdate: birthdate,
	}
	result, err := mhsRepository.Insert(ctx, mhs)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestMhsGetAll(t *testing.T) {
	mhsRepository := NewMhsRepositoryImpl(repository_pattern_mhs.GetConnection())
	ctx := context.Background()

	result, err := mhsRepository.FindAll(ctx)
	if err != nil {
		panic(err)
	}
	for _, mhs := range result {
		fmt.Println(mhs)
	}
}

func TestMhsGetById(t *testing.T) {
	mhsRepository := NewMhsRepositoryImpl(repository_pattern_mhs.GetConnection())
	ctx := context.Background()

	result, err := mhsRepository.FindById(ctx, "2200015112")
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestMhsUpdate(t *testing.T) {
	mhsRepository := NewMhsRepositoryImpl(repository_pattern_mhs.GetConnection())
	ctx := context.Background()

	birthdate, _ := time.Parse("2006-01-02", "1998-03-06")
	mhs := entity.Mahasiswa{
		Nim:       "2200015112",
		Email:     "H2uZD@example.com",
		Nama:      "Erna Damayanti - Updated",
		Jurusan:   "Sistem Informasi",
		Birthdate: birthdate,
	}
	result, err := mhsRepository.Update(ctx, mhs)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestMhsDelete(t *testing.T) {
	mhsRepository := NewMhsRepositoryImpl(repository_pattern_mhs.GetConnection())
	ctx := context.Background()

	err := mhsRepository.Delete(ctx, "2200015112")
	if err != nil {
		panic(err)
	}
	fmt.Println("Deleted")
}


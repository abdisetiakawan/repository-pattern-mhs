package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/abdisetiakawan/repository-pattern-mhs/entity"
)

type mhsRepositoryImpl struct {
	DB *sql.DB
}

func NewMhsRepositoryImpl(db *sql.DB) *mhsRepositoryImpl {
	return &mhsRepositoryImpl{DB: db}
}

func (repository *mhsRepositoryImpl) Insert(ctx context.Context, mhs entity.Mahasiswa) (entity.Mahasiswa, error) {
	script := "INSERT INTO mahasiswa (nim, email, nama, birthdate, jurusan) VALUES (?, ?, ?, ?, ?)"
	result, err := repository.DB.ExecContext(ctx, script, mhs.Nim, mhs.Email, mhs.Nama, mhs.Birthdate, mhs.Jurusan)
	if err != nil {
		return mhs, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return mhs, err
	}
	mhs.Id = int32(id)
	return mhs, nil
}

func (repository *mhsRepositoryImpl) FindById(ctx context.Context, nim string) (entity.Mahasiswa, error) {
	script := "SELECT id, nim, email, nama, birthdate, jurusan FROM mahasiswa WHERE nim = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, script, nim)
	if err != nil {
		return entity.Mahasiswa{}, err
	}
	defer rows.Close()
	mhs := entity.Mahasiswa{}
	if rows.Next() {
		rows.Scan(&mhs.Id, &mhs.Nim, &mhs.Email, &mhs.Nama, &mhs.Birthdate, &mhs.Jurusan)
		return mhs, nil
	} else {
		return mhs, errors.New("Nim " + nim + " not found")
	}
}

func (repository *mhsRepositoryImpl) FindAll(ctx context.Context) ([]entity.Mahasiswa, error) {
	script := "SELECT id, nim, email, nama, birthdate, jurusan FROM mahasiswa"
	rows, err := repository.DB.QueryContext(ctx, script)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var mhsList []entity.Mahasiswa
	for rows.Next() {
		mhs := entity.Mahasiswa{}
		rows.Scan(&mhs.Id, &mhs.Nim, &mhs.Email, &mhs.Nama, &mhs.Birthdate, &mhs.Jurusan)
		mhsList = append(mhsList, mhs)
	}
	return mhsList, nil
}

func (repository *mhsRepositoryImpl) Update(ctx context.Context, mhs entity.Mahasiswa) (entity.Mahasiswa, error) {
	script := "UPDATE mahasiswa SET email = ?, nama = ?, birthdate = ?, jurusan = ? WHERE nim = ?"
	_, err := repository.DB.ExecContext(ctx, script, mhs.Email, mhs.Nama, mhs.Birthdate, mhs.Jurusan, mhs.Nim)
	if err != nil {
		return mhs, err
	}
	return mhs, nil
}

func (repository *mhsRepositoryImpl) Delete(ctx context.Context, nim string) error {
	script := "DELETE FROM mahasiswa WHERE nim = ?"
	_, err := repository.DB.ExecContext(ctx, script, nim)
	if err != nil {
		return err
	}
	return nil
}
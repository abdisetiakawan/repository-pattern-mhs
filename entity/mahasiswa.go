package entity

import "time"

type Mahasiswa struct {
	Id        int32
	Nim       string
	Email     string
	Nama      string
	Birthdate time.Time
	Jurusan   string
}
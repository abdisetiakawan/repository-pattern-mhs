# Repository Pattern Mahasiswa

Ini adalah proyek untuk mempelajari pola repository dalam konteks manajemen data mahasiswa menggunakan bahasa pemrograman Go.

## Struktur Proyek

```
repository-pattern-mhs/
├── database.go
├── entity/
│   └── mahasiswa.go
├── go.mod
├── go.sum
├── repository/
│   ├── mahasiswa_repository.go
│   ├── mahasiswa_repository_impl.go
│   └── mahasiswa_repository_impl_test.go
```

## Deskripsi File

- `database.go`: Mengatur koneksi ke database MySQL.
- `entity/mahasiswa.go`: Mendefinisikan struktur data `Mahasiswa`.
- `repository/mahasiswa_repository.go`: Mendefinisikan interface `MahasiswaRepository`.
- `repository/mahasiswa_repository_impl.go`: Implementasi dari interface `MahasiswaRepository`.
- `repository/mahasiswa_repository_impl_test.go`: Unit test untuk implementasi repository `Mahasiswa`.

## Cara Menggunakan

1. Clone repository ini.
2. Pastikan Anda memiliki Go dan MySQL terinstal di sistem Anda.
3. Buat database MySQL dengan nama `repositorypatternmhs`.
4. Jalankan perintah berikut untuk mengunduh dependensi:
   ```sh
   go mod tidy
   ```
5. Jalankan unit test untuk memastikan semuanya berjalan dengan baik:
   ```sh
   go test ./repository
   ```

## Koneksi Database

Koneksi ke database diatur di file `database.go`. Pastikan konfigurasi koneksi sesuai dengan pengaturan database Anda.

```go
func GetConnection() *sql.DB {
    db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/repositorypatternmhs?parseTime=true")
    if err != nil {
        panic(err)
    }
    db.SetMaxIdleConns(10)
    db.SetMaxOpenConns(100)
    db.SetConnMaxIdleTime(5 * time.Minute)
    db.SetConnMaxLifetime(60 * time.Minute)

    return db
}
```

## Struktur Data Mahasiswa

Struktur data `Mahasiswa` didefinisikan di file `entity/mahasiswa.go`.

```go
type Mahasiswa struct {
    Id        int32
    Nim       string
    Email     string
    Nama      string
    Birthdate time.Time
    Jurusan   string
}
```

## Interface Repository

Interface `MahasiswaRepository` didefinisikan di file `repository/mahasiswa_repository.go`.

```go
type MahasiswaRepository interface {
    Insert(ctx context.Context, mhs entity.Mahasiswa) (entity.Mahasiswa, error)
    FindById(ctx context.Context, nim string) (entity.Mahasiswa, error)
    FindAll(ctx context.Context) ([]entity.Mahasiswa, error)
    Update(ctx context.Context, mhs entity.Mahasiswa) (entity.Mahasiswa, error)
    Delete(ctx context.Context, nim string) error
}
```

## Implementasi Repository

Implementasi dari interface `MahasiswaRepository` terdapat di file `repository/mahasiswa_repository_impl.go`.

## Unit Test

Unit test untuk implementasi repository terdapat di file `repository/mahasiswa_repository_impl_test.go`.

## Lisensi

Proyek ini dilisensikan di bawah lisensi MIT. Lihat file [LICENSE](LICENSE) untuk informasi lebih lanjut.

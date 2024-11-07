package models

import (
    "database/sql"
    "log"
    "fmt"
)

type Mahasiswa struct {
    Nama     string
    NIM      string
    Jurusan  string
    Angkatan int
    IPK      float64
}

func GetMahasiswaList() []Mahasiswa {
    return []Mahasiswa{
        {
            Nama:     "Fatih Fikry Oktavianto",
            NIM:      "1202223093",
            Jurusan:  "Sistem Informasi",
            Angkatan: 2022,
            IPK:      3.95,
        },
        {
            Nama:     "Azizi Shafa Asadel",
            NIM:      "1202223094",
            Jurusan:  "Tata Boga",
            Angkatan: 2023,
            IPK:      3.96,
        },
        {
            Nama:     "Rizky Kurniawan",
            NIM:      "1202223095",
            Jurusan:  "Teknik Telekomunikasi",
            Angkatan: 2024,
            IPK:      3.94,
        },
    }
}

// MigrasiDataSQL melakukan migrasi data dari aplikasi Go ke database SQL
func MigrasiDataSQL(db *sql.DB) {
    // Implementasikan logika migrasi data ke dalam tabel Mahasiswa di sini
    // Misalnya, menjalankan query untuk membuat tabel Mahasiswa
    _, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS Mahasiswa (
            Nama VARCHAR(255),
            NIM VARCHAR(10) PRIMARY KEY,
            Jurusan VARCHAR(100),
            Angkatan INT,
            IPK DOUBLE
        )
    `)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Migrasi data berhasil!")
}

package controllers

import (
    "database/sql"
    "log"
    "github.com/FatihZee/backend-golang/models"
)

func InsertMahasiswa(db *sql.DB) {
    mahasiswaList := models.GetMahasiswaList()
    
    for _, mahasiswa := range mahasiswaList {
        query := "INSERT INTO mahasiswa (nama, nim, jurusan, angkatan, ipk) VALUES (?, ?, ?, ?, ?)"
        _, err := db.Exec(query, mahasiswa.Nama, mahasiswa.NIM, mahasiswa.Jurusan, mahasiswa.Angkatan, mahasiswa.IPK)
        if err != nil {
            log.Fatal(err)
        }
    }

    log.Println("Data mahasiswa berhasil dimasukkan ke database.")
}

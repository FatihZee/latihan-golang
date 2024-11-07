package main

import (
	"database/sql"
	"log"

	"github.com/FatihZee/backend-golang/models"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
    // Buat koneksi ke MySQL menggunakan pengguna root dan tidak menggunakan kata sandi
    db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/mahasiswa_db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Uji koneksi
    if err := db.Ping(); err != nil {
        log.Fatal(err)
    }

    log.Println("Koneksi ke MySQL berhasil!")

    // Panggil fungsi migrasi data
    models.MigrasiDataSQL(db)

    // Dapatkan data mahasiswa dari fungsi GetMahasiswaList
    mahasiswaList := models.GetMahasiswaList()

    // Iterasi melalui daftar mahasiswa dan masukkan ke dalam tabel
    for _, mahasiswa := range mahasiswaList {
        _, err := db.Exec("INSERT INTO Mahasiswa (Nama, NIM, Jurusan, Angkatan, IPK) VALUES (?, ?, ?, ?, ?)",
            mahasiswa.Nama, mahasiswa.NIM, mahasiswa.Jurusan, mahasiswa.Angkatan, mahasiswa.IPK)
        if err != nil {
            log.Fatal(err)
        }
    }

    log.Println("Data mahasiswa berhasil ditambahkan!")
}

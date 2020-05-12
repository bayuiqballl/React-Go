package models

import (
	"database/sql"
	"log"
)

type SiswaMysql struct {
	DB *sql.DB
}

func NewSiswaStoreMysql() SiswaStore {
	dsn := "root:password@tcp(localhost:3307)/echo_go?parseTime=true&clientFoundRows=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	return &SiswaMysql{DB: db}
}

func (store *SiswaMysql) All() []Siswa {
	siswas := []Siswa{}
	rows, err := store.DB.Query("SELECT * FROM siswa")
	if err != nil {
		return siswas
	}

	siswa := Siswa{}
	for rows.Next() {
		rows.Scan(&siswa.ID, &siswa.Nama, &siswa.Nisn, &siswa.Pendidikan)
		siswas = append(siswas, siswa)
	}

	return siswas
}

func (store *SiswaMysql) Save(siswa *Siswa) error {
	result, err := store.DB.Exec(`INSERT INTO siswa(nama,nisn,pendidikan) VALUES(?,?,?)`, siswa.Nama, siswa.Nisn, siswa.Pendidikan)

	if err != nil {
		return err
	}
	_, err = result.RowsAffected()
	if err != nil {
		return err
	}
	// new id
	lastID, err := result.LastInsertId()

	siswa.ID = int(lastID)

	return nil
}

func (store *SiswaMysql) Find(id int) *Siswa {
	siswa := Siswa{}

	err := store.DB.QueryRow(`SELECT * FROM siswa WHERE id=?`, id).Scan(
		&siswa.ID,
		&siswa.Nama,
		&siswa.Nisn,
		&siswa.Pendidikan,
	)

	if err != nil {
		log.Fatal(err)
		return nil
	}

	return &siswa
}

func (store *SiswaMysql) Update(siswa *Siswa) error {

	result, err := store.DB.Exec(`
    UPDATE siswa SET nama = ?, nisn = ?,pendidikan = ? WHERE id = ?`,
		siswa.Nama,
		siswa.Nisn,
		siswa.Pendidikan,
		siswa.ID,
	)
	if err != nil {
		return err
	}

	_, err = result.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func (store *SiswaMysql) Delete(siswa *Siswa) error {
	result, err := store.DB.Exec(`DELETE FROM siswa WHERE id = ?`, siswa.ID)
	if err != nil {
		return err
	}

	_, err = result.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

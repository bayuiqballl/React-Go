package models

type Siswa struct {
	ID                     int
	Nama, Nisn, Pendidikan string
}

func CreateSiswa(nama, nisn, pendidikan string) (*Siswa, error) {
	return &Siswa{
		Nama:       nama,
		Nisn:       nisn,
		Pendidikan: pendidikan,
	}, nil
}

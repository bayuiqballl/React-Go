package models

type SiswaStore interface {
	All() []Siswa
	Save(*Siswa) error
	Find(int) *Siswa
	Update(*Siswa) error
	Delete(siswa *Siswa) error
}

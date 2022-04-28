package models

// import "gorm.io/gorm"

type Disease struct {
	// gorm.Model
	Nama 	string 	`json:"nama" gorm:"primaryKey"`
	Dna_seq	string 	`json:"dna_seq" gorm:"unique"`
	// Records []Record `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:Nama;references:Nama"`
}
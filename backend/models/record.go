package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Record struct {
	gorm.Model
	ID 				uint 			`json:"id" gorm:"primaryKey;autoIncrement"`
	Tanggal			datatypes.Date 	`json:"tanggal"`
	Nama_pasien 	string			`json:"nama_pasien"`
	Nama		 	string			`json:"nama"`
	Status 			string			`json:"status"`
	Diseases		[]Disease `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:Nama;references:Nama"`
}
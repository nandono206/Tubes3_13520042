package models

import (
	// "gorm.io/datatypes"
	// "gorm.io/gorm"
)

type Record struct {
	// gorm.Model
	ID 				uint 			`json:"id" gorm:"primaryKey;autoIncrement"`
	Tanggal			string 			`json:"tanggal"`
	Nama_pasien 	string			`json:"nama_pasien"`
	Nama		 	string			`json:"nama"`
	Status 			string			`json:"status"`
	Presentase		float64			`json:"presentase"`
	// Diseases		[]Disease `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:Nama;references:Nama"`
}

// mysql://b157d1d08e40b9:f736ea04@eu-cdbr-west-02.cleardb.net/heroku_dfe21377e5cd9a7?
package controller

import (
	"time"
	"github.com/gofiber/fiber/v2"
	"github.com/nandono206/FinalSeasonPart3/backend/algorithm"
	"github.com/nandono206/FinalSeasonPart3/backend/database"
	"github.com/nandono206/FinalSeasonPart3/backend/models"
	"gorm.io/datatypes"
)


func AddPenyakit(c *fiber.Ctx) error{
	var toAdd map[string]string

	
	


	if err := c.BodyParser(&toAdd); err != nil{
		return err
	}

	if (!algorithm.IsDNAValid(toAdd["dna_seq"])){
		return c.JSON(fiber.Map{
			"message": "invalid DNA sequence",
		})
	}


	penyakit := models.Disease{
		Nama: toAdd["nama"],
		Dna_seq: toAdd["dna_seq"],
	}

	database.DB.Create(&penyakit)

	

	return c.JSON(fiber.Map{
		"message": "Data penyakit berhasil ditambahkan",
	})


}

func IsSakit(c *fiber.Ctx) error{
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.Disease

	// database.DB.Where("dna_seq = ?", data["dna_seq"]).First(&userInput)

	if (!algorithm.IsDNAValid(data["dna_seq"])){
		return c.JSON(fiber.Map{
			"message": "invalid DNA sequence",
		})
	}

	result := database.DB.First(&user, "nama = ?", data["nama"])

	if (result.RowsAffected==0){
		return c.JSON(fiber.Map{
			"message": "Nama Penyakit Not Found",
		})

	}
	if (!algorithm.BmMatch(data["dna_seq"], user.Dna_seq)){
		record := models.Record{
			Nama: data["nama"],
			Nama_pasien: data["nama_pasien"],
			Status: "False",
			Tanggal: datatypes.Date(time.Now()),
		}
	
		database.DB.Create(&record)
		return c.JSON(fiber.Map{
			"message": "Tidak Sakit",
		})
	}

	record := models.Record{
		Nama: data["nama"],
		Nama_pasien: data["nama_pasien"],
		Status: "True",
		Tanggal: datatypes.Date(time.Now()),
	}

	database.DB.Create(&record)

	return c.JSON(fiber.Map{
		"message": "Sakit!",
	})
}
package controller

import (
	// "time"
	// "bufio"
	// "os"
	// "crypto/rc4"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/nandono206/FinalSeasonPart3/backend/algorithm"
	"github.com/nandono206/FinalSeasonPart3/backend/database"
	"github.com/nandono206/FinalSeasonPart3/backend/models"
	// "gorm.io/datatypes"
	// "fmt"
	// "net/http"
	// "github.com/nandono206/FinalSeasonPart3/backend/request"
	//"io/ioutil"
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

	// rows, errDB := database.RawDB.Query("SELECT * FROM diseases WHERE nama = ? ", toAdd["nama"])

	// if errDB != nil{
	// 	return c.JSON(fiber.Map{
	// 		"message": "error connecting to DB",
	// 	})
	// }

	// defer rows.Close()

	// var recs []models.Disease

	// for rows.Next(){
	// 	var rec models.Disease
	// 	if err := rows.Scan(&rec.Nama, &rec.Dna_seq ); err != nil{
	// 		return c.JSON(fiber.Map{
	// 			"message": "error connecting to DB",
	// 		})
	// 	}
	// 	// fmt.Println(rec.ID)
	// 	recs = append(recs, rec)
	// }

	// if len

	// fmt.Println(recs)

	// return c.JSON(recs)
	var user models.Disease

	result := database.DB.First(&user, "nama = ?", toAdd["nama"])

	if (result.RowsAffected!=0){
		return c.JSON(fiber.Map{
			"message": "Penyakit sudah ada di database! tidak bisa mengupdate data",
		})

	}

	var newuser models.Disease

	resultnew := database.DB.First(&newuser, "dna_seq = ?", toAdd["dna_seq"])

	if (resultnew.RowsAffected!=0){
		return c.JSON(fiber.Map{
			"message": "Sequence DNA yang sama dengan sudah terdapat dalam database!",
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
			"ket" : "error",
		})
	}

	result := database.DB.First(&user, "nama = ?", data["nama"])

	if (result.RowsAffected==0){
		return c.JSON(fiber.Map{
			"message": "Nama Penyakit Not Found",
			"ket" : "error",
		})

	}
	if (!algorithm.BmMatch(data["dna_seq"], user.Dna_seq) &&!algorithm.KmpMatch(data["dna_seq"], user.Dna_seq)){
		
		persen := (float64(algorithm.LCS(data["dna_seq"], user.Dna_seq))/float64(len([]rune (user.Dna_seq)))) * float64(100)
		record := models.Record{
			Nama: data["nama"],
			Nama_pasien: data["nama_pasien"],
			Status: "False",
			Tanggal: time.Now().Format("2006-01-02"),
			Presentase: persen,
		}
	
		database.DB.Create(&record)
		return c.JSON(fiber.Map{
			"message": time.Now().Format("2006-01-02") +" - " + data["nama_pasien"] + " - " +  data["nama"] + " - " + "False",
			"presentase": persen,
			"ket" : "succsess",
		})
	}

	// t := time.Now()
	// tForm := t.Format("2006-01-02")
	persen2 := (float64(algorithm.LCS(data["dna_seq"], user.Dna_seq))/float64(len([]rune (user.Dna_seq)))) * float64(100)
	record := models.Record{
		Nama: data["nama"],
		Nama_pasien: data["nama_pasien"],
		Status: "True",
		Tanggal: time.Now().Format("2006-01-02"),
		Presentase: persen2,
	}

	database.DB.Create(&record)

	return c.JSON(fiber.Map{
		"message": time.Now().Format("2006-01-02") +" - " + data["nama_pasien"] + " - " +  data["nama"] + " - " + "True",
		"presentase": 100.00,
		"ket" : "succsess",
	})
}


func CekHistory(c *fiber.Ctx) error{
    
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	norm := algorithm.Filter(data["datas"])

	// var records models.Record

	fmt.Println(norm["tanggal"])
	fmt.Println(norm["penyakit"])

	//fmt.Println(algorithm.Filter("28-4-2022 dead"))
	

	if norm["tanggal"] != "" && norm["penyakit"] != ""{
		rows, errDB := database.RawDB.Query("SELECT * FROM records WHERE nama = ? AND tanggal = ?", norm["penyakit"], norm["tanggal"])

		

		if errDB != nil{
			return c.JSON(fiber.Map{
				"message": "error connecting to DB",
			})
		}

		defer rows.Close()

		var recs []models.Record

		for rows.Next(){
			var rec models.Record
			if err := rows.Scan(&rec.ID, &rec.Tanggal,  &rec.Nama_pasien, &rec.Nama, &rec.Status, &rec.Presentase ); err != nil{
				return c.JSON(fiber.Map{
					"message": "error connecting to DB",
				})
			}
			// fmt.Println(rec.ID)
			recs = append(recs, rec)
		}

		// fmt.Println(recs)

		return c.JSON(recs)


	} else if norm["tanggal"] != "" {
		rows, errDB := database.RawDB.Query("SELECT * FROM records WHERE tanggal = ?", norm["tanggal"])

		if errDB != nil{
			return c.JSON(fiber.Map{
				"message": "error connecting to DB",
			})
		}

		defer rows.Close()

		var recs []models.Record

		for rows.Next(){
			var rec models.Record
			if err := rows.Scan(&rec.ID, &rec.Tanggal,  &rec.Nama_pasien, &rec.Nama, &rec.Status, &rec.Presentase); err != nil{
				return c.JSON(fiber.Map{
					"message": "error connecting to DB",
				})
			}
			// fmt.Println(rec.ID)
			recs = append(recs, rec)
		}

		// fmt.Println(recs)

		return c.JSON(recs)


	} else if norm["penyakit"] != ""{
		rows, errDB := database.RawDB.Query("SELECT * FROM records WHERE Nama = ?", norm["penyakit"])

		

		if errDB != nil{
			return c.JSON(fiber.Map{
				"message": "error connecting to DB",
			})
		}

		defer rows.Close()

		var recs []models.Record

		for rows.Next(){
			var rec models.Record
			if err := rows.Scan(&rec.ID, &rec.Tanggal,  &rec.Nama_pasien, &rec.Nama, &rec.Status, &rec.Presentase); err != nil{
				return c.JSON(fiber.Map{
					"message": "error connecting to DB",
				})
			}
			// fmt.Println(rec.ID)
			recs = append(recs, rec)
		}

		// fmt.Println(recs)

		return c.JSON(recs)


	}


	

	return c.JSON(fiber.Map{
		"message": "Data penyakit berhasil ditambahkan",
	})



}

package connection

import (
	"log"
	"training_golang_keempat/structs"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB  *gorm.DB
	Err error
)

func Connect() {
	DB, Err = gorm.Open("mysql", "dicky:kqhWHoN5hhpvc3JG@tcp(167.172.73.736:9092)/dicky")
	// Untuk mysql merupakan database yang akan digunakan
	// Untuk root merupakan username dari database apabila ada password maka menjadi root:nama_password
	// @tcp(127.0.0.1:3306) IP localhost dan 3306 port mysql
	// training_kedua merupakan nama database

	if Err != nil {
		log.Println("Connection failed", Err)
	} else {
		log.Println("Server up and running")
	}

	DB.AutoMigrate(&structs.Users{})
	// tbl_user didapat dari nama type structs

	DB.AutoMigrate(&structs.Product{})
	// tbl_product didapat dari nama type structs
}

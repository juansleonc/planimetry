package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	fmt.Println("PRUEBA")
	fmt.Println("DB_HOST", os.Getenv("DB_HOST"))
	fmt.Println("DB_USER", os.Getenv("DB_USER"))
	fmt.Println("DB_NAME", os.Getenv("DB_NAME"))
	fmt.Println("DB_PASSWORD", os.Getenv("DB_PASSWORD"))

	db, err := gorm.Open("postgres", "host="+os.Getenv("DB_HOST")+" user="+os.Getenv("DB_USER")+" dbname="+os.Getenv("DB_NAME")+" sslmode=disable password="+os.Getenv("DB_PASSWORD"))
	fmt.Println("failed to connect DB: ", err)
	if err != nil {
		panic("failed to connect database!!!")
	}
	defer db.Close()

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Conectado a Postgres con éxito!",
		})
	})

	r.Run()
}

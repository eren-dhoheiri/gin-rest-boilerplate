package main

import (
	"go-gin-api/handler"
	"go-gin-api/mobil"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
var Router *gin.Engine

func main() {
	dsn := "root@tcp(127.0.0.1:3306)/test_mobil?charset=utf8mb4&parseTime=True&loc=Local"
	
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	
	
	if err != nil {
		log.Fatal(err.Error())
	}
	
	db.AutoMigrate(&mobil.Mobil{}, &mobil.BrandMobil{})
	// db.Migrator().CreateConstraint(&mobil.Mobil{}, "TypeMobil")
	// db.Migrator().CreateConstraint(&mobil.Mobil{}, "fk_mobil_type_mobil")

	if err != nil {
		panic("migration failed")
	}

	mobilRepository := mobil.NewRepository(db)
	mobilService := mobil.NewService(mobilRepository)
	mobilHandler := handler.NewMobilHandler(mobilService)


	Router = gin.Default()
	api := Router.Group("/api/v1")

	api.POST("/brand", mobilHandler.CreateBrand)
	api.GET("/brand", mobilHandler.GetBrand)
	api.GET("/brand/:id", mobilHandler.GetBrandByID)
	api.DELETE("/brand/:id", mobilHandler.DeleteBrand)


	api.POST("/mobil", mobilHandler.CreateMobil)
	api.GET("/mobil", mobilHandler.GetMobil)
	api.GET("/mobil/:id", mobilHandler.GetMobilByID)
	api.PUT("/mobil/:id", mobilHandler.UpdateMobil)
	api.DELETE("/mobil/:id", mobilHandler.DeleteMobil)

	Router.Run(":5000")
}
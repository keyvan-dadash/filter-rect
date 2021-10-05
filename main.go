package main

import (
	"context"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sod-lol/filter-rect/core/repos/rect_repo"
	"github.com/sod-lol/filter-rect/routers"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	router := gin.New()

	root := context.Background()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tehran",
		os.Getenv("PS_HOST"),
		os.Getenv("PS_USERNAME"),
		os.Getenv("PS_PASSWORD"),
		os.Getenv("PS_DBNAME"),
		os.Getenv("PS_PORT"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	rectRepo := rect_repo.RectanglePostgresDB{}
	rectRepo.SetDB(root, db)
	rectRepo.Migrate(root)

	ctxWithRepo := rect_repo.SetRectangleRepoInContext(root, &rectRepo)

	routers.InitRoutes(ctxWithRepo, &router.RouterGroup)

	router.Run(":8080")
}

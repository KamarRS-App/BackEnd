package main

import (
	"kamarRS/config"
	"kamarRS/utils/database/mysql"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := config.GetConfig()
	db := mysql.InitDB(cfg)

	mysql.DBMigration(db)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())

	e.Start(":8080")

}

package main

import (

	"github.com/KamarRS-App/KamarRS-App/config"
	"github.com/KamarRS-App/KamarRS-App/factory"
	"github.com/KamarRS-App/KamarRS-App/utils/database/mysql"


	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	//----------main---------
	cfg := config.GetConfig()
	db := mysql.InitDB(cfg)

	mysql.DBMigration(db)

	e := echo.New()

	factory.InitFactory(e, db)

	e.Use(middleware.Logger())
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())

	e.Start(":8080")

}

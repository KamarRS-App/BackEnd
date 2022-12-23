package main

import (
	"kamarRS/config"
	"kamarRS/utils/database/mysql"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := config.GetConfig()
	db := mysql.InitDB(cfg)
	// db := posgresql.InitDB(cfg)

	mysql.DBMigration(db)

	e := echo.New()
	e.Use(middleware.Logger())
	e.GET("/hello", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "hello world")
	})
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())

	e.Start(":8080")

}

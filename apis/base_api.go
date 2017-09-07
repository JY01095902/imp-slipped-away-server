package apis

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	services "imp-slipped-away-server/services"
)

func Init(){
	services.Init()
	
	e := echo.New()
	e.Use(middleware.CORS())

	InitUserAPIS(e)
	InitWebSocket(e)

	go handleMessages()

	e.Logger.Fatal(e.Start(":9000"))
}
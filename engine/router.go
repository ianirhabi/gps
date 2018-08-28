package engine

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gps.com/gps/src/inputiuran"
	"gps.com/gps/src/lihatsms"
	"gps.com/gps/src/user"
)

func Router() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	enpoinregis := e.Group("/gps/user")
	enpoinregis.POST("/regis/:username/:status", user.Regis)

	enpoinuser := e.Group("/gps/user")
	enpoinuser.POST("/login", user.Login)
	enpoinuser.GET("", user.Getuser)

	epinputiuran := e.Group("/gps/iuran")
	epinputiuran.POST("", inputiuran.Inputiuran)
	epinputiuran.GET("/:bulan/:tahun/anggota/:user/:minggu", inputiuran.Getiuran)

	endpoin_lihatsms := e.Group("/gps/sms")
	endpoin_lihatsms.POST("/:username", lihatsms.Postsms)
	e.Logger.Fatal(e.Start(":4100"))
}

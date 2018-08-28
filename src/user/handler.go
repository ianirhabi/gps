package user

import (
	"net/http"

	"github.com/labstack/echo"
)

func Regis(c echo.Context) (e error) {
	var r Register
	status := c.Param("status")
	username := c.Param("username")
	if err := c.Bind(&r); err == nil {
		if cc, m := PostRegis(r, username, status); m == nil {
			return c.JSON(http.StatusOK, &cc)
		}
	}
	return e
}

func Login(c echo.Context) (e error) {
	var r RequestLogin
	if err := c.Bind(&r); err == nil {
		if cc, m := Postlogin(r); m == nil {
			return c.JSON(http.StatusOK, &cc)
		}
	}
	return e
}

func Getuser(c echo.Context) (e error) {

	if cc, m := Get(); m == nil {
		return c.JSON(http.StatusOK, &cc)
	}
	return e
}

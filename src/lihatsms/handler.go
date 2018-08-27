package lihatsms

import (
	"net/http"

	"github.com/labstack/echo"
)

func Postsms(c echo.Context) (e error) {
	var r Requestsms
	username := c.Param("username")
	if err := c.Bind(&r); err == nil {
		if cc, m := Kirimsms(r, username); m == nil {
			return c.JSON(http.StatusOK, &cc)
		}
	}
	return e
}

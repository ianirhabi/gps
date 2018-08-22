package inputiuran

import (
	"net/http"

	"github.com/labstack/echo"
)

func Inputiuran(c echo.Context) (e error) {
	var r RequestIuran
	if err := c.Bind(&r); err == nil {
		if cc, m := PostIuran(r); m == nil {
			return c.JSON(http.StatusOK, &cc)
		}
	}
	return e
}

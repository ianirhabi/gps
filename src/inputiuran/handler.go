package inputiuran

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func Inputiuran(c echo.Context) (e error) {
	var r RequestIuran
	if err := c.Bind(&r); err == nil {
		if cc, m := PostIuran(r); m == nil {
			return c.JSON(http.StatusOK, &cc)
		}
	} else {
		fmt.Println("debug error  === ", err)
	}
	return e
}
func Getiuran(c echo.Context) (e error) {
	bulan := c.Param("bulan")
	tahun := c.Param("tahun")
	user := c.Param("user")
	minggu := c.Param("minggu")

	if dd, m := IuranAnggotinfo(bulan, tahun, user, minggu); m == nil {
		return c.JSON(http.StatusOK, &dd)
	}

	return e
}

func Deleteiuran(c echo.Context) (e error) {
	bulan := c.Param("bulan")
	tahun := c.Param("tahun")
	user := c.Param("user")
	minggu := c.Param("minggu")

	if dd, m := IuranAnggoDelete(bulan, tahun, user, minggu); m == nil {
		return c.JSON(http.StatusOK, &dd)
	}
	return e
}

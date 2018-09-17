package download

import "github.com/labstack/echo"

func GetApk(c echo.Context) (e error) {
	return c.File("/home/ubuntu/go/src/gps.com/gps/download/gps-1.0.apk")
}

package main

import (
	"net/http"

	"github.com/k0kubun/pp"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET(":tag", func(c echo.Context) error {
		//return c.String(http.StatusOK, c.Param("%E3%81%93%E3%82%93%E3%81%AB%E3%81%A1%E3%82%8F"))
		return c.String(http.StatusOK, c.Param("tag"))
		//return c.String(http.StatusOK, "%e3%81%93%e3%82%93%e3%81%ab%e3%81%a1%e3%82%8f")
	})
	pp.Println("hoge")
	e.Logger.Fatal(e.Start(":1323"))
}

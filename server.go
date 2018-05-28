package main

import (
	"net/http"

	"github.com/labstack/echo"
)

type User struct {
	Name  string `json:"name" form:"name" query:"name"`
	Email string `json:"email" form:"email" query:"email"`
}

func main() {
	e := echo.New()

	e.POST("/users", func(c echo.Context) (err error) {
		u := new(User)
		if err = c.Bind(u); err != nil {
			return
		}
		return c.JSON(http.StatusOK, u)
	})

	e.GET("/users", func(c echo.Context) (err error) {
		u := User{"Me", "her"}
		// if err = c.Bind(u); err != nil {
		// 	return
		// }
		return c.JSON(http.StatusOK, u)
	})

	e.Logger.Fatal(e.Start(":8000"))
}

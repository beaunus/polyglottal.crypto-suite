package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo"
)

// User implements a simple User.
type User struct {
	Name  string `json:"name" form:"name" query:"name"`
	Email string `json:"email" form:"email" query:"email"`
}

// Template contains a pointer to templates.
type Template struct {
	templates *template.Template
}

// Render renders the given template.
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

// Hello renders the hello template with "World"
func Hello(c echo.Context) error {
	return c.Render(http.StatusOK, "hello", "World")
}

func main() {
	e := echo.New()

	t := &Template{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}

	e.Renderer = t
	e.GET("/hello", Hello)

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

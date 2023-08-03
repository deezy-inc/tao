package main

import (
	"errors"
	"html/template"
	"io"
	"os"

	"github.com/deezy-inc/tao/handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

// Implement e.Renderer interface
func (t *TemplateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	tmpl, ok := t.templates[name]
	if !ok {
		err := errors.New("Template not found -> " + name)
		return err
	}
	return tmpl.ExecuteTemplate(w, "base.html", data)
}

// Define the template registry struct
type TemplateRegistry struct {
	templates map[string]*template.Template
}

func main() {
	e := echo.New()

	templates := make(map[string]*template.Template)
	templates["index.html"] = template.Must(template.ParseFiles("templates/index.html", "templates/base.html"))
	templates["signup.html"] = template.Must(template.ParseFiles("templates/signup.html", "templates/base.html"))
	templates["address.html"] = template.Must(template.ParseFiles("templates/address.html", "templates/base.html"))
	templates["send.html"] = template.Must(template.ParseFiles("templates/send.html", "templates/base.html"))
	templates["asset.html"] = template.Must(template.ParseFiles("templates/asset.html", "templates/base.html"))
	templates["proof.html"] = template.Must(template.ParseFiles("templates/proof.html", "templates/base.html"))
	templates["user.html"] = template.Must(template.ParseFiles("templates/user.html", "templates/base.html"))
	templates["login.html"] = template.Must(template.ParseFiles("templates/login.html", "templates/base.html"))
	templates["dashboard.html"] = template.Must(template.ParseFiles("templates/dashboard.html", "templates/base.html"))
	templates["post.html"] = template.Must(template.ParseFiles("templates/post.html", "templates/base.html"))
	e.Renderer = &TemplateRegistry{
		templates: templates,
	}

	e.Logger.SetLevel(log.ERROR)
	e.Use(middleware.Logger())
	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		TokenLookup: "cookie:" + echo.HeaderAuthorization,
		SigningKey:  []byte(handler.Key),
		Skipper: func(c echo.Context) bool {
			// Skip authentication for requests
			if c.Path() != "/dashboard" {
				return true
			}
			return false
		},
	}))

	// Initialize handler
	h := &handler.Handler{}

	// Routes
	e.GET("/", h.Index)
	e.GET("/address", h.ViewAddress)
	e.GET("/send", h.ViewSend)
	e.GET("/asset/:name", h.FetchAsset)
	e.GET("/asset/:name/proof", h.FetchAssetProof)
	e.GET("/dashboard", h.Dashboard)
	e.GET("/list", h.ListUsers)
	e.GET("/:username", h.FetchUser)
	e.GET("/signup", h.ViewSignup)
	e.GET("/login", h.ViewLogin)

	e.PUT("/dashboard", h.UpdateUser)

	e.POST("/address", h.Address)
	e.POST("/send", h.Send)
	e.POST("/post", h.CheckPost)
	e.POST("/signup", h.Signup)
	e.POST("/login", h.Login)
	e.POST("/posts", h.CreatePost)
	e.POST("/dashboard", h.CreateWithdrawal)

	e.Static("/static", "static")
	e.File("/favicon.ico", "static/images/favicon.ico")

	// Start server
	port := "1323"
	if value, ok := os.LookupEnv("PORT"); ok {
		port = value
	}
	e.Logger.Fatal(e.Start(":" + port))
}

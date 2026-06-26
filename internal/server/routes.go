package server

import (
	"net/http"

	"imvinhnguyen/cmd/web"
	"imvinhnguyen/content"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// HSTS when behind HTTPS proxy (DigitalOcean App Platform sets this header)
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if c.Request().Header.Get("X-Forwarded-Proto") == "https" {
				c.Response().Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
			}
			return next(c)
		}
	})

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"https://*", "http://*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// Serve static files
	fileServer := http.FileServer(http.FS(web.Files))
	e.GET("/assets/*", echo.WrapHandler(fileServer))

	page := templ.Handler(web.QuickLinks(content.Site))

	e.GET("/", echo.WrapHandler(page))
	e.GET("/links", echo.WrapHandler(page))
	e.GET("/health", s.healthHandler)

	// redirect all other requests to home page
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		if c.Response().Committed {
			return
		}
		if c.Request().Method == http.MethodGet {
			c.Redirect(http.StatusMovedPermanently, "/")
		} else {
			c.JSON(http.StatusNotFound, map[string]string{"message": "Page not found"})
		}
	}

	return e
}

func (s *Server) healthHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, s.db.Health())
}

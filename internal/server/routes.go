package server

import (
	"net/http"

	"imvinhnguyen/cmd/web"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	handlerTag        = "imvinhnguyen"
	domain            = "imvinhnguyen.com"
	profilePictureUrl = "https://pbs.twimg.com/profile_images/1401961547887374338/0-W1k3N__400x400.jpg"
	youtubeVideoUrl   = "https://www.youtube.com/embed/vZEKF-0tA74?autoplay=1&mute=1"
	socialLinks       = []web.Links{
		{Label: "Email", Link: "mailto:business@imvinhnguyen.com", Icon: "fa-regular fa-envelope"},
		{Label: "YouTube", Link: "https://www.youtube.com/@imvinhnguyen", Icon: "fa-brands fa-youtube"},
		{Label: "Twitter", Link: "https://twitter.com/imvinhnguyen", Icon: "fa-brands  fa-twitter"},
		{Label: "Instagram", Link: "https://www.instagram.com/imvinhnguyen/", Icon: "fa-brands  fa-instagram"},
		{Label: "Snapchat", Link: "https://www.snapchat.com/add/djvinhii", Icon: "fa-brands  fa-snapchat"},
		{Label: "IMDb", Link: "https://www.imdb.com/name/nm12372318/", Icon: "fa-brands fa-imdb"},
	}
	quickLinks = []web.Links{
		{Label: "My Ebay Store", Link: "https://www.ebay.com/usr/vinhsellstuff", Icon: ""},
	}
)

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

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

	quickLinks := templ.Handler(web.QuickLinks(
		domain,
		handlerTag,
		profilePictureUrl,
		youtubeVideoUrl,
		socialLinks,
		quickLinks,
	))

	// for now until we create proper home page
	e.GET("/", echo.WrapHandler(quickLinks))
	e.GET("/links", echo.WrapHandler(quickLinks))
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

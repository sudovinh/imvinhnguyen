package server

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"imvinhnguyen/cmd/web"
)

var (
	handlerTag        = "imvinhnguyen"
	domain           = "imvinhnguyen.com"
	profilePictureUrl = "https://pbs.twimg.com/profile_images/1401961547887374338/0-W1k3N__400x400.jpg"
	youtubeVideoUrl  = "https://www.youtube.com/embed/vZEKF-0tA74?autoplay=1&mute=1"
	socialLinks = []web.Links{
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

	fileServer := http.FileServer(http.FS(web.Files))
	e.GET("/assets/*", echo.WrapHandler(fileServer))

	// e.GET("/web", echo.WrapHandler(templ.Handler(web.HelloForm())))
	// e.POST("/hello", echo.WrapHandler(http.HandlerFunc(web.HelloWebHandler)))
	e.GET("/links", echo.WrapHandler(templ.Handler(web.QuickLinks(
		domain,
		handlerTag,
		profilePictureUrl,
		youtubeVideoUrl,
		socialLinks,
		quickLinks,
	))))

	e.GET("/", s.HelloWorldHandler)

	e.GET("/health", s.healthHandler)

	return e
}

func (s *Server) HelloWorldHandler(c echo.Context) error {
	resp := map[string]string{
		"message": "Hello World",
	}

	return c.JSON(http.StatusOK, resp)
}

func (s *Server) healthHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, s.db.Health())
}

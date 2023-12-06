package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/MatheusAbdias/bookings/pkg/config"
	"github.com/MatheusAbdias/bookings/pkg/handlers"
	"github.com/MatheusAbdias/bookings/pkg/render"
	scs "github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var app config.AppConfig

var session *scs.SessionManager

func main() {
	app.Debug = true

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = !app.Debug

	app.Session = session
	tc, err := render.CreateTemplateCache()
	if err != nil {
		slog.Error(fmt.Sprintf("cannot create template cache:%s", err))
	}

	app.TemplateCache = tc
	app.UseCache = false
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))

	server := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = server.ListenAndServe()
	slog.Error(fmt.Sprintf("cannot start server: %s", err))
}

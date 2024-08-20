package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/snaypssnp/bookings/pkg/config"
	"github.com/snaypssnp/bookings/pkg/handlers"
	"github.com/snaypssnp/bookings/pkg/render"
	"log"
	"net/http"
	"time"
)

var portNumber = ":8080"
var app config.AppConfig
var session *scs.SessionManager

func main() {
	tc, err := render.CreateTemplateCache()

	app.InProduction = false
	app.TemplateCache = tc
	app.UseCache = true

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	if err != nil {
		log.Fatal("cannot create template cache")
	}

	render.NewTemplates(&app)

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	fmt.Printf("Starting application on port %s\n", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}

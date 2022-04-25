package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/mehmath/booking/pkg/config"
	"github.com/mehmath/booking/pkg/handlers"
	"github.com/mehmath/booking/pkg/render"

	"github.com/alexedwards/scs/v2"
)

const (
	port = ":8080"
)

var app config.AppConfig
var session *scs.SessionManager

func main() {

	//productionda bunu true yap
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache", err)
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println("Aplication started...")
	// _ = http.ListenAndServe(port, nil)

	srv := &http.Server{
		Addr:    port,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}

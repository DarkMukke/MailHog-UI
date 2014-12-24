package main

import (
	"flag"
	"os"

	"github.com/ian-kent/Go-MailHog/http"
	"github.com/ian-kent/go-log/log"
	gotcha "github.com/ian-kent/gotcha/app"
	"github.com/mailhog/MailHog-Server/config"
	"github.com/mailhog/MailHog-UI/assets"
	"github.com/mailhog/MailHog-UI/web"
)

var conf *config.Config
var exitCh chan int

func configure() {
	config.RegisterFlags()
	flag.Parse()
	conf = config.Configure()
}

func main() {
	configure()

	// FIXME need to make API URL configurable

	exitCh = make(chan int)
	cb := func(app *gotcha.App) {
		web.CreateWeb(conf, app)
	}
	go http.Listen(conf, assets.Asset, exitCh, cb)

	for {
		select {
		case <-exitCh:
			log.Printf("Received exit signal")
			os.Exit(0)
		}
	}
}

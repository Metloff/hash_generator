package main

import (
	"log"
	"os"

	bugsnag "github.com/bugsnag/bugsnag-go"
	cli "github.com/jawher/mow.cli"

	"github.com/metloff/hash_generator/api"
)

// Флаги.
var (
	app = cli.App("cv-listener", "CV listener to work with CV.")

	bugsnagKey = app.StringOpt("bugsnag-key", "", "Specify bugsnag API key")
	bugsnagRS  = app.StringOpt("bugsnag-release-stage", "", "Specify bugsnag release stage")
	publicAddr = app.StringOpt("public-addr", ":4184", "Public API address")
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	app.Action = appAction
}

func main() {
	if err := app.Run(os.Args); err != nil {
		log.Fatalln("cli:", err)
	}
}

// Флаги начинают работать только в данной функции.
func appAction() {
	defer bugsnag.Recover()

	bugsnag.Configure(bugsnag.Configuration{
		APIKey:       *bugsnagKey,
		ReleaseStage: *bugsnagRS,
		PanicHandler: func() {},
	})

	// Публичное API.
	publicAPIManager := api.NewManager()
	panic(publicAPIManager.Listen(*publicAddr))
}

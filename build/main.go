package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

var (
	TOCAS_PATH = os.Getenv("TOCAS_PATH") // path to `/tocas`
	TOCAS_DOCS = os.Getenv("TOCAS_DOCS") // path to `/tocas-docs`
	LANG       = ""
)

func main() {
	if TOCAS_PATH == "" || TOCAS_DOCS == "" {
		log.Fatal("env var `TOCAS_PATH`, `TOCAS_DOCS` are required")
	}

	app := &cli.App{
		Name:  "tocas-buildtool",
		Usage: "tools for compiling Tocas UI and docs",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "lang",
				Aliases: []string{"l"},
				Value:   "zh-tw",
				Usage:   "language code, e.g. `zh-tw`",
			},
			&cli.StringFlag{
				Name:  "screenshot-path",
				Usage: "local `file:///` to examples folder, e.g. `file:///C:/Users/yami/tocas/examples/`",
			},
		},
		Commands: []*cli.Command{
			{
				Name:   "build",
				Usage:  "compiles and builds the docs for the specified language",
				Action: build,
			},
			{
				Name:   "compile",
				Usage:  "compiles Tocas UI's CSS and JS files from TOCAS_SRC to TOCAS_DIST",
				Action: compile,
			},
			{
				Name:   "screenshot",
				Usage:  "screenshots the examples and save the images to /examples/screenshots",
				Action: screenshot,
			},
			{
				Name:   "fontawesome",
				Usage:  "update tocas icons from Font Awesome",
				Action: fontawesome,
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

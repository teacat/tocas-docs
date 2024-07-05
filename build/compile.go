package main

import (
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
	"github.com/tdewolff/minify/v2/js"
	cli "github.com/urfave/cli/v2"
)

var (
	RegexpImportCSS = regexp.MustCompile(`@import ".\/(.*?)";`)
	RegexpImportJS  = regexp.MustCompile(`// @import "(.*?)";`)
)

// compile compiles Tocas UI's CSS and JS files from TOCAS_SRC to TOCAS_DIST
func compile(c *cli.Context) error {
	// load `/src/tocas.css`
	b, err := os.ReadFile(Path("{src}/tocas.css"))
	if err != nil {
		log.Fatal(err)
	}
	tocas := string(b)

	var content string
	// resolve all `@import` files and replace `@import` in `tocas.css` with real content
	for _, v := range RegexpImportCSS.FindAllStringSubmatch(string(b), -1) {
		b, err := os.ReadFile(Path("{src}/" + v[1]))
		if err != nil {
			log.Fatal(err)
		}

		// remove this `@import` line from original `tocas.css`
		tocas = strings.ReplaceAll(tocas, v[0], "")
		content += string(b) + "\n"
	}

	// create `/dist` if not exists
	if err := os.MkdirAll(Path("{dist}"), 0777); err != nil {
		log.Fatal(err)
	}

	// save resolved `tocas.css` as `/dist/tocas.css`
	if err := os.WriteFile(Path("{dist}/tocas.css"), []byte(content+tocas), 0777); err != nil {
		log.Fatal(err)
	}

	m := minify.New()
	m.AddFunc("text/css", css.Minify)
	m.AddFunc("text/javascript", js.Minify)

	// minify the content of `/dist/tocas.css` and output to `/dist/tocas.min.css`
	s, err := m.String("text/css", content+tocas)
	if err != nil {
		panic(err)
	}

	if err := os.WriteFile(Path("{dist}/tocas.min.css"), []byte(s), 0644); err != nil {
		log.Fatal(err)
	}

	// load `/src/scripts/tocas.js`
	b, err = os.ReadFile(Path("{src}/scripts/tocas.js"))
	if err != nil {
		log.Fatal(err)
	}
	tocasjs := string(b)

	// resolve all `@import` files and replace `@import` in `tocas.js` with real content
	for _, v := range RegexpImportJS.FindAllStringSubmatch(string(b), -1) {
		b, err := os.ReadFile(Path("{src}/scripts/" + v[1]))
		if err != nil {
			log.Fatal(err)
		}
		// remove this `@import` line from original `tocas.js`
		tocasjs = strings.ReplaceAll(tocasjs, v[0], string(b)+"\n")
	}

	// save resolved `tocas.js` as `/dist/tocas.js`
	if err := os.WriteFile(Path("{dist}/tocas.js"), []byte(tocasjs), 0777); err != nil {
		log.Fatal(err)
	}
	// minify the content of `/dist/tocas.js` and output to `/dist/tocas.min.js`
	s, err = m.String("text/javascript", tocasjs)
	if err != nil {
		panic(err)
	}

	if err := os.WriteFile(Path("{dist}/tocas.min.js"), []byte(s), 0644); err != nil {
		log.Fatal(err)
	}

	// remove old `/dist/fonts/icons` and copy from `/src/fonts/icons`
	if err := os.RemoveAll(Path("{dist}/fonts/icons")); err != nil {
		log.Fatal(err)
	}

	if output, err := exec.Command("cp", "-rf", Path("{src}/fonts/icons"), Path("{dist}/fonts/icons")).CombinedOutput(); err != nil {
		log.Fatal(err.Error() + string(output))
	}
	// remove old `/dist/flags` and copy from `/src/flags`
	if err := os.RemoveAll(Path("{dist}/flags")); err != nil {
		log.Fatal(err)
	}
	if output, err := exec.Command("cp", "-rf", Path("{src}/flags"), Path("{dist}/flags")).CombinedOutput(); err != nil {
		log.Fatal(err.Error() + string(output))
	}
	return nil
}

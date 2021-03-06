package main

import (
	"flag"
	"log"
	"os"
	"strings"
	"time"

	"github.com/radovskyb/watcher"
	//"text/template"
)

var doDebug bool

func themeLog(format string, v ...interface{}) {
	log.Printf("mktheme: "+format, v...)
}

func themeDebug(format string, v ...interface{}) {
	if doDebug {
		themeLog(format, v...)
	}
}

func themeLogErr(format string, v ...interface{}) {
	log.Printf("mktheme error: "+format, v...)
}

func themeLogFatal(err error) {
	log.Fatalf("mktheme error: %s", err)
}

func main() {
	log.SetFlags(0)
	dbg := strings.ToLower(os.Getenv("CLARION_DEBUG"))
	switch dbg {
	case "y", "yes", "true", "1":
		doDebug = true
	}
	var watchFiles bool
	var makeScreenshots bool
	flag.BoolVar(&watchFiles, "watch", false, "watch files for changes and rebuild theme")
	flag.BoolVar(&makeScreenshots, "makeshots", false, "make the screenshots using applescript")
	if watchFiles && makeScreenshots {
		log.Fatalf("cannot use -watch and -makeshots together")
	}
	flag.Parse()
	if len(flag.Args()) != 2 {
		log.Fatalf("usage: mktheme <spec markdown file> <output directory>")
	}
	specPath := flag.Args()[0]
	outputPath := flag.Args()[1]
	spec, err := loadSpec(specPath)
	if err != nil {
		log.Fatalf("error loading specification: %s", err)
	}
	themeLog("building themes...")
	if err := buildThemes(spec, outputPath); err != nil {
		themeLogFatal(err)
	}
	themeLog("complete!")
	if watchFiles {
		themeLog("watching for changes...")
		w := watcher.New()
		w.SetMaxEvents(1)
		w.FilterOps(watcher.Write)
		w.Add("../SPEC.md")
		w.Add("template/clarion-color-theme.json")
		go func() {
			for {
				select {
				case <-w.Event:
					themeLog("rebuilding themes...")
					if err := buildThemes(spec, outputPath); err != nil {
						themeLogFatal(err)
					}
					themeLog("complete!")
				case err := <-w.Error:
					themeLogFatal(err)
				case <-w.Closed:
					return
				}
			}
		}()
		if err := w.Start(time.Millisecond * 500); err != nil {
			themeLogFatal(err)
		}
	}
	if makeScreenshots {
		if err := buildScreenshots(spec, outputPath); err != nil {
			themeLogFatal(err)
		}
	}
}

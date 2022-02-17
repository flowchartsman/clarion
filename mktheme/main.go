package main

import (
	"flag"
	"log"
	"time"

	"github.com/radovskyb/watcher"
	//"text/template"
)

func themeLog(format string, v ...interface{}) {
	log.Printf("mktheme: "+format, v...)
}

func themeLogErr(format string, v ...interface{}) {
	log.Printf("mktheme error: "+format, v...)
}

func themeLogFatal(err error) {
	log.Fatalf("mktheme error: %s", err)
}

func main() {
	log.SetFlags(0)
	var watchFiles bool
	flag.BoolVar(&watchFiles, "watch", false, "watch files for changes and rebuild theme")
	flag.Parse()
	if len(flag.Args()) != 2 {
		log.Fatalf("usage: mktheme <spec markdown file> <output directory>")
	}
	specPath := flag.Args()[0]
	outputPath := flag.Args()[1]
	themeLog("building themes...")
	if err := buildThemes(specPath, outputPath); err != nil {
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
					if err := buildThemes(specPath, outputPath); err != nil {
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
			log.Fatalln(err)
		}
	}
}

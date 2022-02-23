package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

type cmdRunner struct {
	cmdNumber int
	pause     time.Duration
	err       error
}

func newCmdRunner(pause time.Duration) *cmdRunner {
	return &cmdRunner{
		cmdNumber: 0,
		pause:     pause,
	}
}

func (c *cmdRunner) runWithInput(input []byte, cmd string, args ...string) *cmdRunner {
	if c.err == nil {
		ec := exec.Command(cmd, args...)
		inpipe, err := ec.StdinPipe()
		if err != nil {
			c.err = err
			return c
		}
		if input != nil {
			go func() {
				io.Copy(inpipe, bytes.NewReader(input))
				inpipe.Close()
			}()
		}
		if err := ec.Run(); err != nil {
			c.err = fmt.Errorf("command `%s %s` failed: %s", cmd, strings.Join(args, " "), err)
		} else {
			time.Sleep(c.pause)
		}
	}
	return c
}

func (c *cmdRunner) run(cmd string, args ...string) *cmdRunner {
	return c.runWithInput(nil, cmd, args...)
}

func (c *cmdRunner) Err() error {
	return c.err
}

func buildScreenshots(spec *spec, outputPath string) error {
	tmpl, err := template.New("").ParseFiles(`template/change_theme.scpt.tmpl`)
	if err != nil {
		return err
	}
	themescripts := make([][]byte, 0, len(spec.themeBases))
	for _, base := range spec.themeBases {
		// generate the individual keystrokes needed to change the theme
		themechars := strings.Split(base, "")
		var buf bytes.Buffer
		if err := tmpl.ExecuteTemplate(&buf, "change_theme.scpt.tmpl", map[string][]string{"themechars": themechars}); err != nil {
			return err
		}
		themescripts = append(themescripts, buf.Bytes())
	}
	themeLog("waiting for a moment for debug session to start...")
	time.Sleep(5 * time.Second)
	themeLog("prepping for screenshots...")
	c := newCmdRunner(2 * time.Second)
	c.run("osascript", "prep_screenshots.scpt")
	themeLog("generating screenshots...")
	readmePreviews := []map[string]string{}
	for i, s := range themescripts {
		c.runWithInput(s, "osascript")
		// need to run it twice for some reason sometimes. vscode debug extension bug maybe?
		c.runWithInput(s, "osascript")
		screenshotFilename := "img/Clarion-" + spec.themeBases[i] + ".jpg"
		c.run("screencapture", "-x", "-R0,23,900,900", filepath.Join(outputPath, screenshotFilename))
		readmePreviews = append(readmePreviews, map[string]string{
			"themename":  "Clarion " + spec.themeBases[i],
			"screenshot": screenshotFilename,
		})
	}
	if c.Err() != nil {
		return c.Err()
	}
	themeLog("regenerating readme...")
	sort.Slice(readmePreviews, func(i, j int) bool {
		if readmePreviews[i]["themename"] == "Clarion White" {
			return true
		}
		if readmePreviews[j]["themename"] == "Clarion White" {
			return false
		}
		return readmePreviews[i]["themename"] < readmePreviews[j]["themename"]
	})
	tmpl, err = template.New("").ParseFiles(`template/README.md.tmpl`)
	if err != nil {
		return err
	}
	readmeout, err := os.Create(filepath.Join(outputPath, "README.md"))
	if err != nil {
		return err
	}
	defer readmeout.Close()
	if err := tmpl.ExecuteTemplate(readmeout, "README.md.tmpl", readmePreviews); err != nil {
		return err
	}
	themeLog("done!")
	return nil
}

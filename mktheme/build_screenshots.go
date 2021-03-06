package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"os/exec"
	"path/filepath"
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
		cmdStr := cmd + " " + strings.Join(args, " ")
		themeDebug("run: %s", cmdStr)
		if err := ec.Run(); err != nil {
			output, _ := ec.CombinedOutput()
			outputStr := string(output)
			if outputStr == "" {
				outputStr = "<NO OUTPUT>"
			}
			c.err = fmt.Errorf("command `%s` failed: %s - %s", cmdStr, err, outputStr)
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
	tmpl, err := template.New("").ParseFiles(`screenshot_scripts/change_theme.scpt.tmpl`)
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
	c.run("osascript", "screenshot_scripts/prep_screenshots.scpt")
	themeLog("generating screenshots...")
	for i, s := range themescripts {
		c.runWithInput(s, "osascript")
		// need to run it twice for some reason sometimes. vscode debug extension bug maybe?
		c.runWithInput(s, "osascript")
		screenshotFilename := "img/Clarion-" + spec.themeBases[i] + ".jpg"
		c.run("screencapture", "-x", "-R0,23,1300,900", filepath.Join(outputPath, screenshotFilename))
	}
	c.run("osascript", "screenshot_scripts/close_window.scpt")
	if c.Err() != nil {
		return c.Err()
	}
	themeLog("done!")
	return nil
}

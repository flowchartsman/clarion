package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"gopkg.in/yaml.v3"
)

type MkthemeConfig struct {
	themeRoot    string
	themeVersion string
	Targets      []TemplateTarget `yaml:"targets"`
}

type templateMode string

const (
	ModeSimple templateMode = "simple"
	ModeMulti  templateMode = "multi"
	ModeTree   templateMode = "tree"
)

type TemplateTarget struct {
	Mode   templateMode `yaml:"mode"`
	Input  string       `yaml:"input"`
	Output string       `yaml:"output"`
}

func loadConfig() (*MkthemeConfig, error) {
	var config MkthemeConfig
	configFile, err := os.Open("mktheme_conf.yml")
	if err != nil {
		themeLogFatal(fmt.Errorf("loading config: %v", err))
		if err := yaml.NewDecoder(configFile).Decode(&config); err != nil {
			themeLogFatal(fmt.Errorf("invalid config: %v", err))
		}
		configFile.Close()
	}
	config.themeRoot = mustOutput("git", "rev-parse", "--show-toplevel")
	config.themeVersion = mustOutput("git", "describe", "--tags")
	return &config, nil
}

func mustOutput(command ...string) string {
	cmd := exec.Command(command[0], command[1:]...)
	cmdOut, err := cmd.CombinedOutput()
	if err != nil {
		themeLogFatal(fmt.Errorf(`error while running "%s": %v`, strings.Join(command, " "), err))
	}
	s := string(bytes.TrimSpace(cmdOut))
	if s == "" {
		themeLogFatal(fmt.Errorf(`output from "%s" is empty, expected result`, strings.Join(command, " ")))
	}
	return s
}

package main

import (
	"bufio"
	"fmt"
	"html/template"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"sort"
)

type ThemePkg struct {
	Version       string
	ThemeContribs []ThemeContrib
}

type ThemeContrib struct {
	Label string
	File  string
}

func buildThemes(config *MkthemeConfig, spec *spec) error {
	pkg := &ThemePkg{
		Version: config.themeVersion,
	}

	// get all theme themes
	themes, err := generateVariants(config, spec)
	if err != nil {
		return err
	}
	// for targets
	// do the different generation types here

	for _, theme := range themes {
		// do the different generation
		themeFilename := fmt.Sprintf("clarion-color-theme-%s.json", theme.Variant)
		pkg.ThemeContribs = append(pkg.ThemeContribs, ThemeContrib{
			Label: theme.ThemeName,
			File:  themeFilename,
		})
		outPath := filepath.Join(config.themeRoot, "themes", themeFilename)
		outFile, err := newFileWriter(outPath)
		if err != nil {
			return fmt.Errorf("unable to create output file %q: %v", outPath, err)
		}
		tmpl, err := template.New("").ParseFiles("templates/clarion-color-theme.json")
		if err != nil {
			return fmt.Errorf("template parse error: %v", err)
		}
		//
		if err := tmpl.ExecuteTemplate(outFile, "clarion-color-theme.json", theme); err != nil {
			return fmt.Errorf("template execution error: %v", err)
		}
		outFile.Close()
	}
	pkgPath := filepath.Join(config.themeRoot, "package.json")
	outPkg, err := os.Create(pkgPath)
	if err != nil {
		return fmt.Errorf("unable to create package output file %q: %v", pkgPath, err)
	}
	defer outPkg.Close()
	tmpl, err := template.New("").ParseFiles("templates/package.json.tmpl")
	if err != nil {
		return fmt.Errorf("package template parse error: %v", err)
	}
	if err := tmpl.ExecuteTemplate(outPkg, "package.json.tmpl", pkg); err != nil {
		return fmt.Errorf("package template execution error: %v", err)
	}
	// generate readme
	previews := []map[string]string{}
	for _, base := range spec.themeBases {
		previews = append(previews, map[string]string{
			"themename":  "Clarion " + base,
			"screenshot": fmt.Sprintf("Clarion-%s.jpg", base),
		})
	}
	sort.Slice(previews, func(i, j int) bool {
		if previews[i]["themename"] == "Clarion White" {
			return true
		}
		if previews[j]["themename"] == "Clarion White" {
			return false
		}
		return previews[i]["themename"] < previews[j]["themename"]
	})

	readmeData := map[string]interface{}{
		"previews":  previews,
		"baseTheme": themes[0],
	}
	tmpl, err = template.New("").ParseFiles(`templates/README.md.tmpl`)
	if err != nil {
		return err
	}
	readmeout, err := os.Create(filepath.Join(config.themeRoot, "README.md"))
	if err != nil {
		return err
	}
	defer readmeout.Close()
	if err := tmpl.ExecuteTemplate(readmeout, "README.md.tmpl", readmeData); err != nil {
		return err
	}
	return nil
}

var commentLine = regexp.MustCompile(`^\s*//`)

func newFileWriter(outputFile string) (io.WriteCloser, error) {
	output, err := os.Create(outputFile)
	if err != nil {
		return nil, err
	}
	pr, pw := io.Pipe()
	go func() {
		s := bufio.NewScanner(pr)
		for s.Scan() {
			if commentLine.MatchString(s.Text()) {
				continue
			}
			output.WriteString(s.Text() + "\n")
		}
		if err != nil {
			panic("scanner: " + err.Error())
		}
		if err := output.Close(); err != nil {
			panic("closing file: " + err.Error())
		}
	}()
	return pw, nil
}

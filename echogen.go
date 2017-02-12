package main

//go:generate go-bindata -ignore .DS_Store -pkg main -o bindata.go templates/...

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"gopkg.in/urfave/cli.v2"
)

var (
	app   *cli.App
	flags = []cli.Flag{
		&cli.StringFlag{
			Name:  "name",
			Usage: "Application's name",
		},
	}
)

func init() {
	app = &cli.App{
		Name:    "EchoGen, a Yeoman scaffolding for Echo web framework",
		Version: "0.0.1",
		Authors: []*cli.Author{
			&cli.Author{Name: "mdouchement", Email: "https://github.com/mdouchement"},
		},
		Flags:  flags,
		Action: action,
	}
}

func main() {
	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
	}
}

func action(c *cli.Context) error {
	projectName := c.String("name")
	if projectName == "" {
		return errors.New("--name argument must be specified")
	}

	pwd, err := os.Getwd()
	if err != nil {
		return err
	}

	data := map[string]interface{}{
		"project_name": projectName,
		"project_path": strings.Replace(pwd, os.ExpandEnv("$GOPATH")+"/src/", "", 1),
	}

	return generate(data, pwd, projectName)
}

func generate(data map[string]interface{}, pwd, projectName string) error {
	funcMap := template.FuncMap{
		"upcase": strings.ToUpper,
	}

	for _, name := range AssetNames() {
		asset, err := Asset(name)
		if err != nil {
			return err
		}

		p := strings.Replace(name, "templates/", "", 1)
		dir := filepath.Dir(p)
		base := strings.Replace(filepath.Base(p), filepath.Ext(p), "", 1)

		var filename string
		if base == "main.go" {
			filename = filepath.Join(pwd, projectName, dir, fmt.Sprintf("%s.go", data["project_name"]))
		} else {
			filename = filepath.Join(pwd, projectName, dir, base)
		}
		fmt.Println(filename)

		os.MkdirAll(filepath.Join(pwd, projectName, dir), 0755)

		f, err := os.Create(filename)
		if err != nil {
			return err
		}

		if strings.Contains(filename, "/views/") {
			f.Write(asset)
		} else {
			tmpl := template.Must(template.New(name).Funcs(funcMap).Parse(string(asset)))
			if err := tmpl.Execute(f, data); err != nil {
				return err
			}
		}

		f.Close()
	}

	fmt.Println(`
Run the following command to install dependencies:
	glide install

Run the server:
	make serve
`)

	return nil
}

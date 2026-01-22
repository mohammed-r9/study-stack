package module

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

const (
	basePath    = "internal/entities/app"
	templateDir = "internal/codegen/module/templates"
)

type Generator struct {
	Module string
	Path   string
}

func New(name string) *Generator {
	name = strings.ToLower(name)

	return &Generator{
		Module: name,
		Path:   filepath.Join(basePath, name),
	}
}

func (g *Generator) Generate() {
	mkdir(filepath.Join(g.Path, "internal/handler"))
	mkdir(filepath.Join(g.Path, "internal/service"))

	render("init.go.templ", filepath.Join(g.Path, "init.go"), g)
	render("handler.go.templ", filepath.Join(g.Path, "internal/handler/handler.go"), g)
	render("service.go.templ", filepath.Join(g.Path, "internal/service/service.go"), g)
}

func mkdir(path string) {
	if err := os.MkdirAll(path, 0755); err != nil {
		log.Fatal(err)
	}
}

func render(tmplName, outPath string, g *Generator) {
	tmplPath := filepath.Join(templateDir, tmplName)

	t, err := template.ParseFiles(tmplPath)
	if err != nil {
		log.Fatal("Template error:", err)
	}

	f, err := os.Create(outPath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if err := t.Execute(f, g); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"bytes"
	_ "embed"
	"strings"
	"text/template"
)

//go:embed rdc.sh
var bin string

type Template struct {
	DepsVar   string
	Deps      string
	UnsetDeps bool
}

func gencode(deps []string) string {
	var args args
	if len(deps) == 0 {
		return ""
	}

	rdc := Template{args.DepsName, "'" + strings.Join(deps, "' '") + "'", !args.DepsCode}
	tmpl, err := template.New("").Parse(bin)
	if err != nil {
		panic(err)
	}
	var buf bytes.Buffer
	tmpl.Execute(&buf, rdc)

	return buf.String()
}

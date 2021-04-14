package main

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"text/template"
)

func TestName(t *testing.T) {
	files, _ := filepath.Glob("*")
	fmt.Println(files) // contains a list of all files in the current directory

	type Inventory struct {
		Material string
		Count    uint
	}

	var Data = struct {
		What string
		A []string
	}{
		What: "1221",
		A: []string{"a", "b", "c"},
	}
	//var A = []string{"a", "b", "c"}
	tmpl, err := template.New("test").Parse("{{$Draft:=true}} {{$Draft}}\n{{range .A}}{{.}}{{end}}\n")
	tmpl.Execute(os.Stdout, Data)
	tmpl, err = template.New("test").Parse("\n{{23 -}} < {{- 45}}\n")
	tmpl.Execute(os.Stdout, nil)
	sweaters := Inventory{"wool", 17}
	tmpl, err = template.New("test").Parse("{{.Count}} items are made of {{.Material}}")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, sweaters)
	if err != nil {
		panic(err)
	}
}

//  "/*"
func Test1(t *testing.T) {
	//^/.+
	match, err := filepath.Match("/*", "/example")
	if err != nil {
		return
	}
	fmt.Println(match)
}

package main

import (
	"os"
	"text/template"
)

func main() {
	t1 := template.New("t1`")
	// 方式1
	t1, err := t1.Parse("Value is, {{.}}")
	if err != nil {
		panic(err)
	}

	// 方式2
	t1 = template.Must(t1.Parse("Value, {{.}}\n"))
	t1.Execute(os.Stdout, "some text")
	t1.Execute(os.Stdout, 5)
	t1.Execute(os.Stdout, []string{"Hello", "World"})
	t1.Execute(os.Stdout, map[string]string{"Name": "John"})
	// 类型一致
	Create := func(name, t string) *template.Template {
		return template.Must(template.New(name).Parse(t))
	}
	t2 := Create("t2", "name: {{.Name}}\n")
	t2.Execute(os.Stdout, struct{ Name string }{Name: "John"})
	t2.Execute(os.Stdout, map[string]string{"Name": "Doe"})
	t3 := Create("t3", "{{if . -}}yes {{else -}}  no {{end}}\n")
	t3.Execute(os.Stdout, "not empty")
	t3.Execute(os.Stdout, "")
	t4 := Create("t4", "{{range .}}{{.Name}} {{end}} \n")
	t4.Execute(os.Stdout, []struct{ Name string }{{Name: "John"}, {Name: "Doe"}})
}

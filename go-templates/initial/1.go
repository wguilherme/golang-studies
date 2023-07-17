// You can edit this code!
// Click here and start typing.
package main

import (
	"html/template"
	"os"
)

func sum(a, b int) int {
	return a + b
}

func main() {

	type Inventory struct {
		Material string
		Count    uint
	}

	sweaters := Inventory{"wool", 17}

	// map for custom template functions
	funcs := template.FuncMap{"sum": sum}

	// compile template with custom functions

	templateContent := "{{.Count}} items are made of {{ sum 3 1 }}"

	tmpl, err := template.New("test").Funcs(funcs).Parse(templateContent)

	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(os.Stdout, sweaters)
	if err != nil {
		panic(err)
	}

}

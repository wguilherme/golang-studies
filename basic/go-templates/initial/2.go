package main

import (
	"html/template"
	"os"
)

func main() {
	type Inventory struct {
		Material string
		Count    uint
	}

	alertsThreshold := []struct {
		ConnectionsSuccessRate float64
		RequestTime            int
	}{
		{ConnectionsSuccessRate: 0.5, RequestTime: 20},
	}

	sweaters := Inventory{"wool", 17}

	data := struct {
		Inventory       Inventory
		AlertsThreshold []struct {
			ConnectionsSuccessRate float64
			RequestTime            int
		}
	}{
		Inventory:       sweaters,
		AlertsThreshold: alertsThreshold,
	}

	templateContent := `
	{{.Inventory.Count}} items are made of {{.Inventory.Material}}
	{{ $connections_success_rate := or  }}
	`

	tmpl, err := template.New("test").Parse(templateContent)
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}

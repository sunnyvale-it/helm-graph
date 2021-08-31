package output

import (
	"encoding/json"
	"fmt"

	"github.com/sunnyvale-it/helm-graph/chart"
	"gopkg.in/yaml.v2"
)

func Write(format string, graph *chart.Chart) {

	switch format {
	case "json":
		json, err := json.Marshal(graph)
		if err != nil {
			fmt.Printf("An error occured: %v\n", err)
			return
		}
		fmt.Println(string(json))
	case "yaml":
		yaml, err := yaml.Marshal(graph)
		if err != nil {
			fmt.Printf("An error occured: %v\n", err)
			return
		}
		fmt.Println(string(yaml))
	default:
		fmt.Printf("Not a valid output format")
	}

}

package chart

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/sunnyvale-it/helm-graph/utils"
	"gopkg.in/yaml.v2"
)

type Chart struct {
	Name    string
	Version string
	Repo    string
	Deps    []Chart
}

func getChartDependencies(chart Chart) []Chart {

	type dependency struct {
		Name       string
		Repository string
		Version    string
		Tags       []string
	}

	tmpDir, err := os.MkdirTemp("", ".helm-graph-tmp-")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	archive := fmt.Sprintf("%s-%s.tgz", chart.Name, chart.Version)
	url := fmt.Sprintf("%s/%s", chart.Repo, archive)
	filePath := fmt.Sprintf("%s/%s", tmpDir, archive)

	// Get the data
	resp, err := http.Get(url)
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(fmt.Sprintf("%s", filePath))
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)

	gzipReader, err := os.Open(filePath)

	os.Mkdir(fmt.Sprintf("%s/%s", tmpDir, chart.Name), 0755)

	err = utils.UntarChartYaml(fmt.Sprintf("%s", tmpDir), gzipReader)

	yfile, err := ioutil.ReadFile(fmt.Sprintf("%s/%s/Chart.lock", tmpDir, chart.Name))

	data := make(map[interface{}]interface{})

	err = yaml.Unmarshal(yfile, &data)

	d, err := yaml.Marshal(data["dependencies"])

	deps := []dependency{}

	err = yaml.Unmarshal([]byte(string(d)), &deps)

	chartDeps := []Chart{}

	for _, v := range deps {
		tmpChart := &chart
		tmpChart.Name = v.Name
		tmpChart.Version = v.Version

		chartDeps = append(chartDeps, *tmpChart)

		tmpChart.Deps = chartDeps

	}

	return chartDeps

}

func (chart *Chart) Graph() {

	for _, v := range getChartDependencies(*chart) {
		chart.Deps = append(chart.Deps, v)
		if len(v.Deps) > 0 {
			v.Graph()
		}
	}

}
